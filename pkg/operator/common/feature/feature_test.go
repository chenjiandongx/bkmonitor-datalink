// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package feature

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSelector(t *testing.T) {
	cases := []struct {
		input  string
		output map[string]string
	}{
		{
			input: "__meta_kubernetes_endpoint_address_target_name=^eklet-.*,__meta_kubernetes_endpoint_address_target_kind=Node",
			output: map[string]string{
				"__meta_kubernetes_endpoint_address_target_name": "^eklet-.*",
				"__meta_kubernetes_endpoint_address_target_kind": "Node",
			},
		},
		{
			input: "__meta_kubernetes_endpoint_address_target_name=^eklet-.*,,",
			output: map[string]string{
				"__meta_kubernetes_endpoint_address_target_name": "^eklet-.*",
			},
		},
		{
			input: "foo=bar, , ,k1=v1 ",
			output: map[string]string{
				"foo": "bar",
				"k1":  "v1",
			},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.output, parseSelector(c.input))
	}
}

func TestParseLabelJoinMatcher(t *testing.T) {
	cases := []struct {
		input string

		annotations       []string
		annotationsInput  map[string]string
		annotationsOutput map[string]string

		labels       []string
		labelsInput  map[string]string
		labelsOutput map[string]string
	}{
		{
			input:       "Pod://annotation:biz.service,annotation:biz.set,label:zone.key1,label:zone.key2",
			annotations: []string{"biz.service", "biz.set"},
			labels:      []string{"zone.key1", "zone.key2"},
			annotationsInput: map[string]string{
				"biz.service": "test-service",
				"biz.set":     "test-set",
			},
			annotationsOutput: map[string]string{
				"biz.service": "test-service",
				"biz.set":     "test-set",
			},
		},
		{
			input:       "Pod:// annotation: biz.service, annotation: biz.set, label: zone.key1, label: zone.key2",
			annotations: []string{"biz.service", "biz.set"},
			labels:      []string{"zone.key1", "zone.key2"},
		},
		{
			input:       "Pod://annotation({[0]['id']}):biz.service,annotation:biz.set,label:zone.key1,label:zone.key2",
			annotations: []string{"biz.service", "biz.set"},
			labels:      []string{"zone.key1", "zone.key2"},
			annotationsInput: map[string]string{
				"biz.service": `[{"id":"test-service","foo":"bar"}]`,
				"biz.set":     "test-set",
			},
		},
	}

	for _, c := range cases {
		matcher := parseLabelJoinMatcher(c.input)
		assert.Equal(t, c.annotations, matcher.Annotations)
		assert.Equal(t, c.labels, matcher.Labels)

		for _, s := range matcher.Annotations {
			v, ok := c.annotationsInput[s]
			if !ok {
				continue
			}

			f := matcher.AnnotationsFunc[s]
			if f != nil {
				t.Log(f(v))
			}
		}
	}
}
