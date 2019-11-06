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
ROOT_DOCKER_SERVICE_NAME ?= $(ROOT_NAME)
ROOT_DOCKER_SERVICE_PORT ?= 39000
ROOT_DOCKER_IMAGE_NAME ?= $(ROOT_NAME)
ROOT_DOCKER_IMAGE_TAG ?= $(DIST_VERSION)
ROOT_BUILD_PATH ?= ./build
ROOT_DIST ?= ./dist
ROOT_REPO ?= ./dist
ROOT_TEST_BUILD_PATH ?= $(ROOT_BUILD_PATH)/test/$(DIST_VERSION)
ROOT_TEST_DIST_PATH ?= $(ROOT_DIST)/test/$(DIST_VERSION)
ROOT_TEST_OS_DIST_PATH ?= $(ROOT_DIST)/$(DIST_OS)/test/$(DIST_VERSION)
ROOT_REPO_DIST_PATH ?= $(ROOT_REPO)/$(DIST_VERSION)
ROOT_REPO_OS_DIST_PATH ?= $(ROOT_REPO)/$(DIST_OS)/release/$(DIST_VERSION)

ROOT_LOCAL_IP_V4_LINUX = $$(ifconfig enp8s0 | grep inet | grep -v inet6 | cut -d ':' -f2 | cut -d ' ' -f1)
ROOT_LOCAL_IP_V4_DARWIN = $$(ifconfig en0 | grep inet | grep -v inet6 | cut -d ' ' -f2)

ROOT_LOG_PATH ?= ./log
ROOT_SWAGGER_PATH ?= ./docs

SERVER_TEST_SSH_ALIASE = aliyun-ecs
SERVER_TEST_FOLDER = /home/work/Document/
SERVER_REPO_SSH_ALIASE = temp-gin-web
SERVER_REPO_FOLDER = /home/ubuntu/$(ROOT_NAME)

# can use as https://goproxy.io/ https://gocenter.io https://mirrors.aliyun.com/goproxy/
ENV_GO_PROXY ?= https://goproxy.io/

checkEnvGo:
ifndef GOPATH
	@echo Environment variable GOPATH is not set
	exit 1
endif

# check must run environment
init:
	@echo "~> start init this project"
	@echo "-> check version"
	go version
	@echo "-> check env golang"
	go env
	@echo "~> you can use [ make help ] see more task"
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod vendor
	which swag
	swag --help
	@echo "~> you can use [ make help ] see more task"

# For Docker dev images init
initDockerDevImages:
	@echo "~> start init this project in docker"
	@echo "-> check version"
	go version
	@echo "-> check env golang"
	go env
	@echo "-> install swag"
	GOPROXY="$(ENV_GO_PROXY)" go get -u github.com/swaggo/swag/cmd/swag

checkDepends:
	# in GOPATH just use GO111MODULE=on go mod init to init after golang 1.12
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod verify

downloadDepends:
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod download
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod vendor

tidyDepends:
	-GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod tidy

# can check depends
dep: checkDepends
	@echo "just show dependencies info below"

dependenciesGraph:
	GOPROXY="$(ENV_GO_PROXY)" GO111MODULE=on go mod graph

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
	@echo "-> start build local OS"
	@go build -o build/main main.go

buildARCH:
	@echo "-> start build OS:$(DIST_OS) ARCH:$(DIST_ARCH)"
	@GOOS=$(DIST_OS) GOARCH=$(DIST_ARCH) go build -o build/main main.go

buildDocker: checkDepends cleanBuild
	@echo "-> start build OS:$(DIST_OS_DOCKER) ARCH:$(DIST_ARCH_DOCKER)"
	@GOOS=$(DIST_OS_DOCKER) GOARCH=$(DIST_ARCH_DOCKER) go build -o build/main main.go

dev: buildMain
	-ENV_WEB_AUTO_HOST=true ./build/main -c ./conf/config.yaml

runTest:  buildMain
	-ENV_WEB_AUTO_HOST=true ./build/main -c ./conf/test/config.yaml

test: checkDepends buildMain checkTestDistPath
	mv ./build/main $(ROOT_TEST_DIST_PATH)
	cp ./conf/test/config.yaml $(ROOT_TEST_DIST_PATH)
	@echo "=> pkg at: $(ROOT_TEST_DIST_PATH)"

testTar: test
	cd $(ROOT_DIST)/test && tar zcvf $(ROOT_NAME)-test-$(DIST_VERSION).tar.gz $(DIST_VERSION)

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

releaseOSTar: releaseOS
	@echo "=> start tar release as os $(DIST_OS) $(DIST_ARCH)"
	tar zcvf $(ROOT_DIST)/$(DIST_OS)/release/$(ROOT_NAME)-$(DIST_OS)-$(DIST_ARCH)-$(DIST_VERSION).tar.gz $(ROOT_REPO_OS_DIST_PATH)

dockerLocalImageInit:
	docker build --tag $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG) .

dockerLocalImageRebuild:
	docker image rm $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	docker build --tag $(ROOT_DOCKER_IMAGE_NAME):$(ROOT_DOCKER_IMAGE_TAG) .

localIPLinux:
	@echo "=> now run as docker with linux"
	@echo "local ip address is: $(ROOT_LOCAL_IP_V4_LINUX)"

dockerRunLinux: localIPLinux
	docker image inspect --format='{{ .Created}}' $(ROOT_DOCKER_SERVICE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	ENV_WEB_HOST=$(ROOT_LOCAL_IP_V4_LINUX) \
	ENV_WEB_PORT=$(ROOT_DOCKER_SERVICE_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_SERVICE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	docker-compose up -d
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_SERVICE_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_SERVICE_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_SERVICE_NAME) -f"

dockerRestartLinux: localIPLinux
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_SERVICE_NAME)
	ENV_WEB_HOST=$(ROOT_LOCAL_IP_V4_LINUX) \
	ENV_WEB_PORT=$(ROOT_DOCKER_SERVICE_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_SERVICE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	docker-compose up -d
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_SERVICE_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_SERVICE_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_SERVICE_NAME) -f"

localIPDarwin:
	@echo "=> now run as docker with darwin"
	@echo "local ip address is: $(ROOT_LOCAL_IP_V4_DARWIN)"

dockerRunDarwin: localIPDarwin
	docker image inspect --format='{{ .Created}}' $(ROOT_DOCKER_SERVICE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	ENV_WEB_HOST=$(ROOT_LOCAL_IP_V4_DARWIN) \
	ENV_WEB_PORT=$(ROOT_DOCKER_SERVICE_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_SERVICE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	docker-compose up -d
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_SERVICE_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_SERVICE_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_SERVICE_NAME) -f"

dockerRestartDarwin: localIPDarwin
	docker inspect --format='{{ .State.Status}}' $(ROOT_NAME)
	ENV_WEB_HOST=$(ROOT_LOCAL_IP_V4_DARWIN) \
	ENV_WEB_PORT=$(ROOT_DOCKER_SERVICE_PORT) \
	ROOT_NAME=$(ROOT_DOCKER_SERVICE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	docker-compose restart
	-sleep 5
	@echo "=> container $(ROOT_DOCKER_SERVICE_NAME) now status"
	docker inspect --format='{{ .State.Status}}' $(ROOT_DOCKER_SERVICE_NAME)
	@echo "=> see log with: docker logs $(ROOT_DOCKER_SERVICE_NAME) -f"

dockerStop:
	ROOT_NAME=$(ROOT_DOCKER_SERVICE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_WEB_PORT=$(ROOT_DOCKER_SERVICE_PORT) \
	docker-compose stop

dockerPrune: dockerStop
	ROOT_NAME=$(ROOT_DOCKER_SERVICE_NAME) \
	DIST_TAG=$(ROOT_DOCKER_IMAGE_TAG) \
	ENV_WEB_PORT=$(ROOT_DOCKER_SERVICE_PORT) \
	docker-compose rm -f $(ROOT_DOCKER_SERVICE_NAME)
	docker rmi -f $(ROOT_DOCKER_SERVICE_NAME):$(ROOT_DOCKER_IMAGE_TAG)
	docker network prune
	docker volume prune

scpDockerComposeTest:
	scp ./conf/test/docker-compose.yml $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)
	@echo "=> finish update docker compose at test"

scpTestOS:
	#scp -r $(ROOT_TEST_OS_DIST_PATH) $(SERVER_TEST_SSH_ALIASE):$(SERVER_TEST_FOLDER)
	@echo "=> must check below config of set for testOSScp"

help:
	@echo "make init - check base env of this project"
	@echo "make dep - check depends of project"
	@echo "make dependenciesGraph - see depends graph of project"
	@echo "make tidyDepends - tidy depends graph of project"
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
	@echo "make dockerRunLinux - run docker-compose server as $(ROOT_DOCKER_SERVICE_NAME) container-name at $(ROOT_DOCKER_SERVICE_NAME) in dockerRunLinux"
	@echo "make dockerRunDarwin - run docker-compose server as $(ROOT_DOCKER_SERVICE_NAME) container-name at $(ROOT_DOCKER_SERVICE_NAME) in macOS"
	@echo "make dockerStop - stop docker-compose server as $(ROOT_DOCKER_SERVICE_NAME) container-name at $(ROOT_DOCKER_SERVICE_NAME)"
