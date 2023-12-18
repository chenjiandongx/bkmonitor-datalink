RELEASE_PATH ?= bin
BUILD_NO ?= 0
NAME = bkmonitorbeat
VERSION = ''
COMMIT_ID = ''
PERF_MODE ?= 0

.PHONY: build
build:
	env VERSION=$(VERSION) RELEASE_PATH=$(RELEASE_PATH) PERF_MODE=$(PERF_MODE) ./script/gse_package.sh

.PHONY: test
test:
	go test ./...
