.PHONY: test check clean build dist all

TOP_DIR := $(shell pwd)

# ifeq ($(FILE), $(wildcard $(FILE)))
# 	@ echo target file not found
# endif

DIST_VERSION := v1.0.0
# linux windows darwin  list as: go tool dist list
DIST_OS := linux
DIST_ARCH := amd64

DIST_OS_DOCKER ?= linux
DIST_ARCH_DOCKER ?= amd64

ROOT_NAME ?= gin-api-swagger-temple

ROOT_BUILD_PATH ?= ./build
ROOT_DIST ?= ./dist
ROOT_REPO ?= ./dist
ROOT_LOG_PATH ?= ./log
ROOT_SWAGGER_PATH ?= ./docs

ROOT_TEST_BUILD_PATH ?= $(ROOT_BUILD_PATH)/test/$(DIST_VERSION)
ROOT_TEST_DIST_PATH ?= $(ROOT_DIST)/test/$(DIST_VERSION)
ROOT_TEST_OS_DIST_PATH ?= $(ROOT_DIST)/$(DIST_OS)/test/$(DIST_VERSION)
ROOT_REPO_DIST_PATH ?= $(ROOT_REPO)/$(DIST_VERSION)
ROOT_REPO_OS_DIST_PATH ?= $(ROOT_REPO)/$(DIST_OS)/release/$(DIST_VERSION)

ROOT_LOCAL_IP_V4_LINUX = $$(ifconfig enp8s0 | grep inet | grep -v inet6 | cut -d ':' -f2 | cut -d ' ' -f1)
ROOT_LOCAL_IP_V4_DARWIN = $$(ifconfig en0 | grep inet | grep -v inet6 | cut -d ' ' -f2)

SERVER_TEST_SSH_ALIASE = aliyun-ecs
SERVER_TEST_FOLDER = /home/work/Document/
SERVER_REPO_SSH_ALIASE = temp-gin-web
SERVER_REPO_FOLDER = /home/ubuntu/$(ROOT_NAME)

# can use as https://goproxy.io/ https://gocenter.io https://mirrors.aliyun.com/goproxy/
ENV_GO_PROXY ?= https://goproxy.cn/

ENV_GO_SWAG_VERSION ?= v1.6.2

# include MakeDockerRun.mk for docker run
include MakeGoMod.mk
include MakeDockerRun.mk

checkEnvGo:
ifndef GOPATH
	@echo Environment variable GOPATH is not set
	exit 1
endif

env:
	@go version
	@swag --version

# check must run environment
init:
	@echo "~> start init this project"
	@echo "-> check version"
	go version
	@echo "-> check env golang"
	go env
	@echo "if swag can not find can use [ make installTools ] to fix"
	which swag
	swag --version
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod download
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod vendor
	@echo "~> you can use [ make help ] see more task"

installTools:
	go get -u github.com/swaggo/swag/cmd/swag@${ENV_GO_SWAG_VERSION}

cleanBuild:
	@if [ -d ${ROOT_BUILD_PATH} ]; then rm -rf ${ROOT_BUILD_PATH} && echo "~> cleaned ${ROOT_BUILD_PATH}"; else echo "~> has cleaned ${ROOT_BUILD_PATH}"; fi

cleanDist:
	@if [ -d ${ROOT_DIST} ]; then rm -rf ${ROOT_DIST} && echo "~> cleaned ${ROOT_DIST}"; else echo "~> has cleaned ${ROOT_DIST}"; fi

cleanLog:
	@if [ -d ${ROOT_LOG_PATH} ]; then rm -rf ${ROOT_LOG_PATH} && echo "~> cleaned ${ROOT_LOG_PATH}"; else echo "~> has cleaned ${ROOT_LOG_PATH}"; fi

cleanSwaggerDoc:
	@if [ -d ${ROOT_SWAGGER_PATH} ]; then rm -rf ${ROOT_SWAGGER_PATH} && echo "~> cleaned ${ROOT_SWAGGER_PATH}"; else echo "~> has cleaned ${ROOT_SWAGGER_PATH}"; fi

clean: cleanBuild cleanLog cleanSwaggerDoc
	@echo "~> clean finish"

checkTestBuildPath:
	@if [ ! -d ${ROOT_TEST_BUILD_PATH} ]; then mkdir -p ${ROOT_TEST_BUILD_PATH} && echo "~> mkdir ${ROOT_TEST_BUILD_PATH}"; fi

checkTestDistPath:
	@if [ ! -d ${ROOT_TEST_DIST_PATH} ]; then mkdir -p ${ROOT_TEST_DIST_PATH} && echo "~> mkdir ${ROOT_TEST_DIST_PATH}"; fi

checkTestOSDistPath:
	@if [ ! -d ${ROOT_TEST_OS_DIST_PATH} ]; then mkdir -p ${ROOT_TEST_OS_DIST_PATH} && echo "~> mkdir ${ROOT_TEST_OS_DIST_PATH}"; fi

checkReleaseDistPath:
	@if [ ! -d ${ROOT_REPO_DIST_PATH} ]; then mkdir -p ${ROOT_REPO_DIST_PATH} && echo "~> mkdir ${ROOT_REPO_DIST_PATH}"; fi

checkReleaseOSDistPath:
	@if [ ! -d ${ROOT_REPO_OS_DIST_PATH} ]; then mkdir -p ${ROOT_REPO_OS_DIST_PATH} && echo "~> mkdir ${ROOT_REPO_OS_DIST_PATH}"; fi

buildSwagger:
	which swag
	swag --version
	@if [ -d ${ROOT_SWAGGER_PATH} ]; then rm -rf ${ROOT_SWAGGER_PATH} && echo "~> cleaned ${ROOT_SWAGGER_PATH}"; else echo "~> has cleaned ${ROOT_SWAGGER_PATH}"; fi
	swag init

buildMain: dep buildSwagger
	@echo "-> start build local OS"
	@go build -o build/main main.go

buildARCH: dep buildSwagger
	@echo "-> start build OS:$(DIST_OS) ARCH:$(DIST_ARCH)"
	@GOOS=$(DIST_OS) GOARCH=$(DIST_ARCH) go build -o build/main main.go

buildDocker: dep cleanBuild
	@echo "-> start build OS:$(DIST_OS_DOCKER) ARCH:$(DIST_ARCH_DOCKER)"
	@GOOS=$(DIST_OS_DOCKER) GOARCH=$(DIST_ARCH_DOCKER) go build -o build/main main.go

test:
	@echo "=> run test start"
	@go test -test.v ./...

testBenchmem:
	@echo "=> run test benchmem start"
	@go test -test.benchmem

dev: buildMain
	-ENV_WEB_AUTO_HOST=true ./build/main -c ./conf/config.yaml

runTest: buildMain
	-ENV_WEB_AUTO_HOST=true ./build/main -c ./conf/test/config.yaml

distTest: buildMain checkTestDistPath
	mv ./build/main $(ROOT_TEST_DIST_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_DIST_PATH)
	@echo "=> pkg at: $(ROOT_TEST_DIST_PATH)"

tarDistTest: distTest
	cd $(ROOT_DIST)/test && tar zcvf $(ROOT_NAME)-test-$(DIST_VERSION).tar.gz $(DIST_VERSION)

distTestOS: buildARCH checkTestOSDistPath
	@echo "=> Test at: $(DIST_OS) ARCH as: $(DIST_ARCH)"
	mv ./build/main $(ROOT_TEST_OS_DIST_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_OS_DIST_PATH)
	@echo "=> pkg at: $(ROOT_TEST_OS_DIST_PATH)"

distRelease: buildMain checkReleaseDistPath
	mv ./build/main $(ROOT_REPO_DIST_PATH)
	cp ./conf/release/config.yaml $(ROOT_REPO_DIST_PATH)
	@echo "=> pkg at: $(ROOT_REPO_DIST_PATH)"

distReleaseOS: buildARCH checkReleaseOSDistPath
	@echo "=> Release at: $(DIST_OS) ARCH as: $(DIST_ARCH)"
	mv ./build/main $(ROOT_REPO_OS_DIST_PATH)
	cp ./conf/release/config.yaml $(ROOT_REPO_OS_DIST_PATH)
	@echo "=> pkg at: $(ROOT_REPO_OS_DIST_PATH)"

tarDistReleaseOS: distReleaseOS
	@echo "=> start tar release as os $(DIST_OS) $(DIST_ARCH)"
	tar zcvf $(ROOT_DIST)/$(DIST_OS)/release/$(ROOT_NAME)-$(DIST_OS)-$(DIST_ARCH)-$(DIST_VERSION).tar.gz $(ROOT_REPO_OS_DIST_PATH)

scpTestOS:
	@echo "=> must check below config of set for testOSScp"
	#scp -r $(ROOT_TEST_OS_DIST_PATH) $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)

scpDockerComposeTest:
	scp ./conf/test/docker-compose.yml $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)
	@echo "=> finish update docker compose at test"

helpProjectRoot:
	@echo "Help: Project root Makefile"
	@echo " need https://github.com/swaggo/swag version: ${ENV_GO_SWAG_VERSION}"
	@echo "-- distTestOS or distReleaseOS will out abi as: $(DIST_OS) $(DIST_ARCH) --"
	@echo "~> make distTest         - build dist at $(ROOT_TEST_DIST_PATH) in local OS"
	@echo "~> make tarDistTest      - build dist at $(ROOT_TEST_OS_DIST_PATH) and tar"
	@echo "~> make distTestOS       - build dist at $(ROOT_TEST_OS_DIST_PATH) as: $(DIST_OS) $(DIST_ARCH)"
	@echo "~> make distRelease      - build dist at $(ROOT_REPO_DIST_PATH) in local OS"
	@echo "~> make distReleaseOS    - build dist at $(ROOT_REPO_OS_DIST_PATH) as: $(DIST_OS) $(DIST_ARCH)"
	@echo "~> make tarDistReleaseOS - build dist at $(ROOT_REPO_OS_DIST_PATH) as: $(DIST_OS) $(DIST_ARCH) and tar"
	@echo ""
	@echo "-- now build name: $(ROOT_NAME) version: $(DIST_VERSION)"
	@echo "~> make init         - check base env of this project"
	@echo "~> make clean        - remove binary file and log files"
	@echo "~> make test         - run test case all"
	@echo "~> make testBenchmem - run go test benchmem case all"
	@echo "~> make runTest      - run server use conf/test/config.yaml"
	@echo "~> make dev          - run server use conf/config.yaml"

help: helpGoMod helpDockerRun helpProjectRoot
	@echo ""
	@echo "-- more info see Makefile include: MakeGoMod.mk MakeDockerRun.mk --"
