SHELL := /bin/bash
BASEDIR = $(shell pwd)

# 可在make是带入参数进行替换
# eg: make SERVICE_NAME=chatgram build
SERVICE_NAME?=chatgram

# build with version infos
versionDir = "github.com/go-eagle/eagle/pkg/version"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

PROJECT_NAME := "github.com/go-microservice/chatgram"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

# proto
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && echo $$b)
API_PROTO_FILES=$(shell find api$(APP_RELATIVE_PATH) -name *.proto)
API_PROTO_PB_FILES=$(shell find api$(APP_RELATIVE_PATH) -name *.pb.go)

# init environment variables
export PATH        := $(shell go env GOPATH)/bin:$(PATH)
export GOPATH      := $(shell go env GOPATH)
export GO111MODULE := on

# make   make all
.PHONY: all
all: lint test build

.PHONY: build
# make build, Build the binary file
build: 
	GOOS=linux GOARCH=amd64 go build -v -ldflags ${ldflags} -o build/$(SERVICE_NAME) cmd/server/main.go cmd/server/wire_gen.go

.PHONY: run
# make run, run current project
run: wire
	go run cmd/server/main.go cmd/server/wire_gen.go

.PHONY: wire
# make wire, generate wire_gen.go
wire: 
	cd cmd/server && wire

.PHONY: dep
# make dep Get the dependencies
dep:
	@go mod download

.PHONY: dev
# make dev
dev:
	@air -c .air.toml

.PHONY: fmt
# make fmt
fmt:
	@gofmt -s -w .

.PHONY: lint
# make lint
lint:
	@golint -set_exit_status ${PKG_LIST}

.PHONY: ci-lint
# make ci-lint
ci-lint: prepare-lint
	${GOPATH}/bin/golangci-lint run ./...

.PHONY: prepare-lint
# make prepare-lint
prepare-lint:
	@if ! which golangci-lint &>/dev/null; then \
  		echo "Installing golangci-lint"; \
  		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest; \
  	fi

.PHONY: test
# make test
test: test-case vet
	@go test -short ${PKG_LIST}

.PHONY: test-case
# make test-case
test-case:
	@go test -cover ./... | grep -v vendor;true

.PHONY: vet
# make vet
vet:
	go vet ./... | grep -v vendor;true

.PHONY: cover
# make cover
cover: gen-coverage
	@go tool cover -html=coverage.txt

.PHONY: gen-coverage
# make gen-coverage
gen-coverage:
	@go test -short -coverprofile coverage.txt -covermode=atomic ${PKG_LIST}

.PHONY: docker
# make docker  生成docker镜像
docker:
	sh deploy/docker_image.sh $(GIT_TAG)

.PHONY: clean
# make clean
clean:
	@-rm -vrf eagle
	@-rm -vrf cover.out
	@-rm -vrf coverage.txt
	@go mod tidy
	@echo "clean finished"

.PHONY: graph
# make graph 生成交互式的可视化Go程序调用图(会在浏览器自动打开)
graph:
	@export GO111MODULE="on"
	@if ! which go-callvis &>/dev/null; then \
  		echo "downloading go-callvis"; \
  		go get -u github.com/ofabry/go-callvis; \
  	fi
	@echo "generating graph"
	@go-callvis github.com/go-eagle/eagle

.PHONY: mockgen
# make mockgen gen mock file
mockgen:
	@echo "downloading mockgen"
	@go get github.com/golang/mock/mockgen
	cd ./internal &&  for file in `egrep -rnl "type.*?interface" ./dao | grep -v "_test" `; do \
		echo $$file ; \
		cd .. && mockgen -destination="./internal/mock/$$file" -source="./internal/$$file" && cd ./internal ; \
	done

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install github.com/google/gnostic@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/favadi/protoc-go-inject-tag@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/gogo/protobuf/protoc-gen-gogo@latest
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	go install -v github.com/cosmtrek/air@latest

.PHONY: proto
# generate proto struct to pb.go only
proto:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=. --go_opt=paths=source_relative \
           $(API_PROTO_FILES)

.PHONY: grpc
# generate grpc code with remove omitempty from json tag
grpc:
# note: --gogofaster_out full replace --go_out=. --go_opt=paths=source_relative
	@for v in $(API_PROTO_FILES); do \
  		echo "./$$v"; \
		protoc --proto_path=. \
			   --proto_path=./third_party \
			   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
			   --gogofaster_out=. --gogofaster_opt=paths=source_relative\
			   "./$$v"; \
    done

.PHONY: http
# generate http code
http:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go-gin_out=. --go-gin_opt=paths=source_relative \
           $(API_PROTO_FILES)

.PHONY: tag
# add tag to pb struct for *.pb.go by using gotags, not include *_grpc.pb.go and *_gin.pb.go
tag: grpc http
	@for v in $(API_PROTO_PB_FILES); do \
  		echo "./$$v"; \
  		protoc-go-inject-tag -input="./$$v"; \
  	done

.PHONY: openapi
# generate openapi
openapi:
	protoc --proto_path=. \
          --proto_path=./third_party \
          --openapi_out=./docs \
          $(API_PROTO_FILES)

.PHONY: docs
# gen swagger doc
docs:
	@if ! which swag &>/dev/null; then \
  		echo "downloading swag"; \
  		go get -u github.com/swaggo/swag/cmd/swag; \
  	fi
	protoc --proto_path=. \
		  --proto_path=./third_party \
		  --openapiv2_out ./docs \
		  --openapiv2_opt logtostderr=true \
		  $(API_PROTO_FILES)
	@#swag init
	@echo "gen-docs done"
	@echo "see docs by: http://localhost:8080/swagger/index.html"

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m  %-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := all
