.PHONY: test check clean build dist all

TOP_DIR := $(shell pwd)

# ifeq ($(FILE), $(wildcard $(FILE)))
# 	@ echo target file not found
# endif

ROOT_BUILD_PATH ?= ./build
ROOT_LOG_PATH ?= ./log
ROOT_SWAGGER_PATH ?= ./docs

checkEnvGo:
ifndef GOPATH
	@echo Environment variable GOPATH is not set
	exit 1
endif

init: checkEnvGo
	@echo "~> start init this project"
	@echo "-> check version"
	-go version
	@echo "-> check env golang"
	-go env
	@echo "-> check env dep"
	-dep help
	@echo "~> you can use [ make help ] see more task"

checkDepends: checkEnvGo
	-dep ensure -v

cleanBuild:
	@if [ -d ${ROOT_BUILD_PATH} ]; then rm -rf ${ROOT_BUILD_PATH} && echo "~> cleaned ${ROOT_BUILD_PATH}"; else echo "~> has cleaned ${ROOT_BUILD_PATH}"; fi

cleanLog:
	@if [ -d ${ROOT_LOG_PATH} ]; then rm -rf ${ROOT_LOG_PATH} && echo "~> cleaned ${ROOT_LOG_PATH}"; else echo "~> has cleaned ${ROOT_LOG_PATH}"; fi

clean: cleanBuild cleanLog
	@echo "~> clean finish"

buildSwagger:
	which swag
	swag --version
	@if [ -d ${ROOT_SWAGGER_PATH} ]; then rm -rf ${ROOT_SWAGGER_PATH} && echo "~> cleaned ${ROOT_SWAGGER_PATH}"; else echo "~> has cleaned ${ROOT_SWAGGER_PATH}"; fi
	swag init

buildMain: buildSwagger
	@go build -o build/main main.go

dev: buildMain
	-./build/main -c ./conf/config.yaml

help:
	@echo "make init - check base env of this project"
	@echo "make checkDepends - check depends of project"
	@echo "make clean - remove binary file and log files"
	@echo "make dev - run server use conf/config.yaml"
