// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package runtimex

import (
	"fmt"
	"runtime"
	"time"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/logger"
)

// GetFuncName get function name
func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// GetCallerFuncName get caller function name
func GetCallerFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

var PanicHandlers = []func(any){
	logPanicHandler,
}

func logPanicHandler(r any) {
	const size = 64 << 10
	stacktrace := make([]byte, size)
	stacktrace = stacktrace[:runtime.Stack(stacktrace, false)]
	if _, ok := r.(string); ok {
		logger.Errorf("observed a panic: %s\n%s", r, stacktrace)
	} else {
		logger.Errorf("observed a panic: %#v (%v)\n%s", r, r, stacktrace)
	}
}

func HandleCrash() {
	if r := recover(); r != nil {
		for _, fn := range PanicHandlers {
			fn(r)
		}
	}
}

func HandleCrashToChan(errorReceiveChan chan<- error) {
	if r := recover(); r != nil {
		const size = 64 << 10
		stacktrace := make([]byte, size)
		stacktrace = stacktrace[:runtime.Stack(stacktrace, false)]

		if _, ok := r.(string); ok {
			err := fmt.Errorf("observed a panic: %s\n%s", r, stacktrace)
			select {
			case errorReceiveChan <- err:
				logger.Infof("send error to receiveChan")
			case <-time.After(5 * time.Second):
				// avoid errorReceiveChan blocking caused by multi-exceptions
				logger.Infof("send error to receiveChan timeout, this error(%s) will not be received", err)
			}

		} else {
			err := fmt.Errorf("observed a panic: %#v (%v)\n%s", r, r, stacktrace)
			select {
			case errorReceiveChan <- err:
				logger.Infof("send error to receiveChan")
			case <-time.After(5 * time.Second):
				logger.Infof("send error to receiveChan timeout, this error(%s) will not be received", err)
			}
		}
	}
}
