// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package httpmiddleware

import (
	"net/http"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/collector/internal/utils"
)

type MiddlewareFunc func(http.Handler) http.Handler

var middlewares = map[string]func(string) MiddlewareFunc{}

func Register(name string, f func(opt string) MiddlewareFunc) {
	middlewares[name] = f
}

func Get(nameOpts string) MiddlewareFunc {
	name, opts := utils.NameOpts(nameOpts)
	f, ok := middlewares[name]
	if !ok {
		return nil
	}

	return f(opts)
}
