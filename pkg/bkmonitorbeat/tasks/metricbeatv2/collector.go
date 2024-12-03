// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package metricbeatv2

import (
	"bytes"
	"io"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/protoparser/common"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/relabel"
	"github.com/prometheus/prometheus/prompb"
	"github.com/prometheus/prometheus/util/fmtutil"
)

type TextStreamParser struct {
	r        io.Reader
	capacity int

	relabelCfgs []*relabel.Config
	dstBuf      []byte
	tailBuf     []byte
	eof         bool
}

func New1(r io.Reader, capacity int, relabelConfigs ...*relabel.Config) *TextStreamParser {
	return &TextStreamParser{
		r:           r,
		capacity:    capacity,
		relabelCfgs: relabelConfigs,
		dstBuf:      make([]byte, 0, capacity),
		tailBuf:     make([]byte, 0, capacity),
	}
}

func (tsp *TextStreamParser) read() error {
	var err error
	tsp.dstBuf, tsp.tailBuf, err = common.ReadLinesBlockExt(tsp.r, tsp.dstBuf, tsp.tailBuf, tsp.capacity)
	if err != nil {
		if err == io.EOF {
			tsp.eof = true
			return nil
		}
	}
	return err
}

func (tsp *TextStreamParser) Read() ([]CustomMetrics, bool, error) {
	if err := tsp.read(); err != nil {
		return nil, false, err
	}

	buf := &bytes.Buffer{} // TODO: 池化
	buf.Write(tsp.dstBuf)
	buf.WriteByte('\n')

	//fmt.Println(bytes.Contains(buf.Bytes(), []byte("1\r\n")))

	wr, err := fmtutil.MetricTextToWriteRequest(buf, nil)
	if err != nil {
		return nil, false, err
	}

	var metrics []CustomMetrics
	seriess := wr.GetTimeseries()
	for i := 0; i < len(seriess); i++ {
		pblbs := seriess[i].Labels
		lbs := make(labels.Labels, 0, len(pblbs))
		for _, l := range pblbs {
			lbs = append(lbs, labels.Label{
				Name:  l.GetName(),
				Value: l.GetValue(),
			})
		}

		var keep bool
		if len(tsp.relabelCfgs) > 0 {
			lbs, keep = relabel.Process(lbs, tsp.relabelCfgs...)
			if !keep {
				continue
			}
		}
		metrics = append(metrics, tsp.castMetrics(lbs, seriess[i].Samples)...)
	}

	return metrics, tsp.eof, err
}

func (tsp *TextStreamParser) castMetrics(lbs labels.Labels, samples []prompb.Sample) []CustomMetrics {
	var metricName string
	dims := make(map[string]string)
	for _, l := range lbs {
		if l.Name != "__name__" {
			dims[l.Name] = l.Value
		} else {
			metricName = l.Name
		}
	}
	ret := make([]CustomMetrics, 0)
	for _, s := range samples {
		ret = append(ret, CustomMetrics{
			Metrics:   map[string]float64{metricName: s.Value},
			Target:    "",
			Timestamp: s.Timestamp,
			Dimension: dims,
		})
	}
	return ret
}

type CustomMetrics struct {
	Metrics   map[string]float64 `json:"metrics"`
	Target    string             `json:"target"`
	Timestamp int64              `json:"timestamp"`
	Dimension map[string]string  `json:"dimension"`
}
