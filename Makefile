.PHONY: test check clean build dist all

TOP_DIR := $(shell pwd)

# ifeq ($(FILE), $(wildcard $(FILE)))
# 	@ echo target file not found
# endif

DIST_VERSION := 1.0.0
# linux windows darwin  list as: go tool dist list
DIST_OS := linux
DIST_ARCH := amd64

DIST_OS_DOCKER ?= linux
DIST_ARCH_DOCKER ?= amd64

ROOT_NAME ?= temp-gin-api-self
ROOT_DOCKER_SERVICE ?= $(ROOT_NAME)
ROOT_BUILD_PATH ?= ./build
ROOT_DIST ?= ./dist
ROOT_REPO ?= ./dist
ROOT_TEST_BUILD_PATH ?= $(ROOT_BUILD_PATH)/test/$(DIST_VERSION)
ROOT_TEST_DIST_PATH ?= $(ROOT_DIST)/test/$(DIST_VERSION)
ROOT_TEST_OS_DIST_PATH ?= $(ROOT_DIST)/$(DIST_OS)/test/$(DIST_VERSION)
ROOT_REPO_DIST_PATH ?= $(ROOT_REPO)/$(DIST_VERSION)
ROOT_REPO_OS_DIST_PATH ?= $(ROOT_REPO)/$(DIST_OS)/release/$(DIST_VERSION)

ROOT_LOG_PATH ?= ./log
ROOT_SWAGGER_PATH ?= ./docs

SERVER_TEST_SSH_ALIASE = aliyun-ecs
SERVER_TEST_FOLDER = /home/work/Document/
SERVER_REPO_SSH_ALIASE = temp-gin-web
SERVER_REPO_FOLDER = /home/ubuntu/$(ROOT_NAME)

# can use as https://goproxy.io/ https://gocenter.io https://mirrors.aliyun.com/goproxy/
INFO_GO_PROXY ?= https://mirrors.aliyun.com/goproxy/

checkEnvGo:
ifndef GOPATH
	@echo Environment variable GOPATH is not set
	exit 1
endif

init: checkEnvGo
	@echo "~> start init this project"
	@echo "-> check version"
	go version
	@echo "-> check env golang"
	go env
	@echo "~> you can use [ make help ] see more task"
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod vendor
	which swag
	swag --help
	@echo "~> you can use [ make help ] see more task"

checkDepends:
	# in GOPATH just use GO111MODULE=on go mod init to init after golang 1.12
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod verify

dependenciesVendor:
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod vendor

dep: checkDepends dependenciesVendor
	@echo "just check dependencies info below"

dependenciesInit:
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod init

dependenciesTidy:
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod tidy

dependenciesDownload:
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod download
	-GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod vendor

dependenciesGraph:
	GOPROXY="$(INFO_GO_PROXY)" GO111MODULE=on go mod graph

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

buildMain: buildSwagger
	@go build -o build/main main.go

buildARCH:
	@echo "-> start build OS:$(DIST_OS) ARCH:$(DIST_ARCH)"
	@GOOS=$(DIST_OS) GOARCH=$(DIST_ARCH) go build -o build/main main.go

buildDocker: checkDepends cleanBuild buildSwagger
	@echo "-> start build OS:$(DIST_OS_DOCKER) ARCH:$(DIST_ARCH_DOCKER)"
	@GOOS=$(DIST_OS_DOCKER) GOARCH=$(DIST_ARCH_DOCKER) go build -o build/main main.go

dev: buildMain
	-./build/main -c ./conf/config.yaml

runTest:  buildMain
	-./build/main -c ./conf/test/config.yaml

test: checkDepends buildMain checkTestDistPath
	mv ./build/main $(ROOT_TEST_DIST_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_DIST_PATH)
	@echo "=> pkg at: $(ROOT_TEST_DIST_PATH)"

testOS: checkDepends buildARCH checkTestOSDistPath
	@echo "=> Test at: $(DIST_OS) ARCH as: $(DIST_ARCH)"
	mv ./build/main $(ROOT_TEST_OS_DIST_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_OS_DIST_PATH)
	@echo "=> pkg at: $(ROOT_TEST_OS_DIST_PATH)"

testOSScp:
	@echo "=> must check below config of set for testOSScp"
	#scp -r $(ROOT_TEST_OS_DIST_PATH) $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)

release: checkDepends buildMain checkReleaseDistPath
	mv ./build/main $(ROOT_REPO_DIST_PATH)
	cp ./conf/release/config.yaml $(ROOT_REPO_DIST_PATH)
	@echo "=> pkg at: $(ROOT_REPO_DIST_PATH)"

releaseOS: checkDepends buildARCH checkReleaseOSDistPath
	@echo "=> Release at: $(DIST_OS) ARCH as: $(DIST_ARCH)"
	mv ./build/main $(ROOT_REPO_OS_DIST_PATH)
	cp ./conf/release/config.yaml $(ROOT_REPO_OS_DIST_PATH)
	@echo "=> pkg at: $(ROOT_REPO_OS_DIST_PATH)"

# just use test config and build as linux amd64
dockerRun: buildDocker checkTestBuildPath
	mv ./build/main $(ROOT_TEST_BUILD_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_BUILD_PATH)
#	cp -R ./static $(ROOT_TEST_BUILD_PATH)
#	cp -R ./views $(ROOT_TEST_BUILD_PATH)
	@echo "=> pkg at: $(ROOT_TEST_BUILD_PATH)"
	@echo "-> try run docker container $(ROOT_NAME)"
	ROOT_NAME=$(ROOT_NAME) DIST_VERSION=$(DIST_VERSION) docker-compose up -d
	-sleep 5
	@echo "=> container $(ROOT_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_NAME)
	docker logs $(ROOT_NAME)
	@echo "most of swagger see at http://127.0.0.1:39000/swagger/index.html"

dockerStop:
	ROOT_NAME=$(ROOT_NAME) DIST_VERSION=$(DIST_VERSION) docker-compose stop

dockerRemove: dockerStop
	ROOT_NAME=$(ROOT_NAME) DIST_VERSION=$(DIST_VERSION) docker-compose rm -f $(ROOT_DOCKER_SERVICE)
	docker network prune

scpDockerComposeTest:
	scp ./conf/test/docker-compose.yml $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)
	@echo "=> finish update docker compose at test"

help:
	@echo "make init - check base env of this project"
	@echo "make dep - check depends of project"
	@echo "make dependenciesGraph - see dependencies graph of project"
	@echo "make dependenciesTidy - tidy dependencies graph of project"
	@echo ""
	@echo "make clean - remove binary file and log files"
	@echo ""
	@echo "-- now build name: $(ROOT_NAME) version: $(DIST_VERSION)"
	@echo "-- testOS or releaseOS will out abi as: $(DIST_OS) $(DIST_ARCH) --"
	@echo "make test - build dist at $(ROOT_TEST_DIST_PATH)"
	@echo "make testOS - build dist at $(ROOT_TEST_OS_DIST_PATH)"
	@echo "make testOSTar - build dist at $(ROOT_TEST_OS_DIST_PATH) and tar"
	@echo "make release - build dist at $(ROOT_REPO_DIST_PATH)"
	@echo "make releaseOS - build dist at $(ROOT_REPO_OS_DIST_PATH)"
	@echo "make releaseOSTar - build dist at $(ROOT_REPO_OS_DIST_PATH) and tar"
	@echo ""
	@echo "make runTest - run server use conf/test/config.yaml"
	@echo "make dev - run server use conf/config.yaml"
	@echo "make dockerRun - run docker-compose server as $(ROOT_DOCKER_SERVICE) container-name at $(ROOT_NAME)"
	@echo "make dockerStop - stop docker-compose server as $(ROOT_DOCKER_SERVICE) container-name at $(ROOT_NAME)"
