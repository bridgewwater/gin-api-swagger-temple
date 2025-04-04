.PHONY: test check clean build dist all
# Makefile root
# can change this by env:ENV_CI_DIST_VERSION use and change by env:ENV_CI_DIST_MARK by CI
ENV_DIST_VERSION=latest
ENV_DIST_MARK=

ROOT_NAME?=gin-api-swagger-temple

## MakeDockerCompose.mk settings start
ROOT_DOCKER_CONTAINER_PORT =34565
ROOT_OWNER ?=bridgewwater
ROOT_PARENT_SWITCH_TAG ?=1.22.11
# for image local build
INFO_TEST_BUILD_DOCKER_PARENT_IMAGE ?=golang
# for image running
INFO_BUILD_DOCKER_FROM_IMAGE ?=alpine:3.17
INFO_BUILD_DOCKER_FILE ?=Dockerfile
INFO_TEST_BUILD_DOCKER_FILE ?=build.dockerfile
INFO_DOCKER_COMPOSE_DEFAULT_FILE ?=docker-compose.yml
## MakeDockerCompose.mk settings end

## run info start
ENV_RUN_INFO_HELP_ARGS=-h
ENV_RUN_INFO_ARGS=-c ./conf/config.yaml
## run info end

## build dist env start
# change to other build entrance
ENV_ROOT_BUILD_ENTRANCE=cmd/gin-api-swagger-temple/main.go
ENV_ROOT_BUILD_BIN_NAME=${ROOT_NAME}
ENV_ROOT_BUILD_PATH = build
ENV_ROOT_BUILD_BIN_PATH=${ENV_ROOT_BUILD_PATH}/${ENV_ROOT_BUILD_BIN_NAME}
ENV_ROOT_LOG_PATH=logs/
ENV_ROOT_SWAGGER_PATH=docs/
# linux windows darwin  list as: go tool dist list
ENV_DIST_GO_OS=linux
# amd64 386
ENV_DIST_GO_ARCH=amd64
# mark for dist and tag helper
ENV_ROOT_MANIFEST_PKG_JSON?=package.json
ENV_ROOT_MAKE_FILE?=Makefile
ENV_ROOT_CHANGELOG_PATH?=CHANGELOG.md
## build dist env end

## go test MakeGoTest.mk start
# ignore used not matching mode
# set ignore of test case like grep -v -E "vendor|go_fatal_error" to ignore vendor and go_fatal_error package
ENV_ROOT_TEST_INVERT_MATCH?="vendor|go_fatal_error|robotn|shirou"
ifeq ($(OS),Windows_NT)
ENV_ROOT_TEST_LIST?=./...
else
ENV_ROOT_TEST_LIST?=$$(go list ./... | grep -v -E ${ENV_ROOT_TEST_INVERT_MATCH})
endif
# test max time
ENV_ROOT_TEST_MAX_TIME:=1m
## go test MakeGoTest.mk end

include z-MakefileUtils/MakeBasicEnv.mk
include z-MakefileUtils/MakeDistTools.mk
include z-MakefileUtils/go-mod.mk
include z-MakefileUtils/go-test.mk
include z-MakefileUtils/go-test-integration.mk
include z-MakefileUtils/go-dist.mk
include z-MakefileUtils/go-list.mk
include z-MakefileUtils/MakeGoDistScp.mk
# include MakeDockerCompose.mk for docker run
include z-MakefileUtils/MakeDockerCompose.mk

all: env

env: envBasic
	@echo "== project env info start =="
	@echo ""
	@echo "ENV_DIST_VERSION                          ${ENV_DIST_VERSION}"
	@echo "ENV_DIST_MARK                             ${ENV_DIST_MARK}"
	@echo "ENV_DIST_CODE_MARK                        ${ENV_DIST_CODE_MARK}"
	@echo ""
	@echo "== project env info end =="

cleanBuild:
	@$(RM) -r ${ENV_ROOT_BUILD_PATH}
	@echo "~> finish clean path: ${ENV_ROOT_BUILD_PATH}"

cleanLog:
	@$(RM) -r ${ENV_ROOT_LOG_PATH}
	@echo "~> finish clean path: ${ENV_ROOT_LOG_PATH}"

cleanSwaggerDoc:
	@$(RM) -r ${ENV_ROOT_SWAGGER_PATH}
	@echo "~> finish clean swagger gen path: ${ENV_ROOT_SWAGGER_PATH}"

cleanTestGoldenData:
	$(info -> notes: remove folder [ testdata ] unable to match subdirectories)
	@$(RM) -r **/testdata
	@$(RM) -r **/**/testdata
	@$(RM) -r **/**/**/testdata
	@$(RM) -r **/**/**/**/testdata
	@$(RM) -r **/**/**/**/**/testdata
	@$(RM) -r **/**/**/**/**/**/testdata
	@$(RM) -r **/**/**/**/**/**/**/testdata
	@$(RM) -r **/**/**/**/**/**/**/**/testdata
	$(info -> finish clean folder [ testdata ])

cleanTestData:
	@$(RM) coverage.txt
	@$(RM) coverage.out
	@$(RM) profile.txt

clean: cleanTestData cleanBuild cleanLog
	@echo "~> clean finish"

cleanAll: clean cleanAllDist
	@echo "~> clean all finish"

.PHONY: init
init:
	@echo "~> start init this project"
	@echo "-> check version"
	go version
	@echo "-> check env golang"
	go env
	@echo "~> you can use [ make help ] see more task"
	-go mod verify
	@echo "~> as dev need kit https://pkg.go.dev/golang.org/x/tools/cmd/stringer"
	@echo "~> can fix by: go install golang.org/x/tools/cmd/stringer@latest"

.PHONY: zymosisGo
zymosisGo:
	$(info -> fix zymosis tools run as: go install -v github.com/convention-change/zymosis/cmd/zymosis@latest)
	@zymosis --version
	$(info -> generate res mark)
	@zymosis -g go

.PHONY: swagger
swagger: zymosisGo cleanSwaggerDoc
	$(info -> fix swag tools run as: go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc4)
	@swag --version
	$(info -> generate swagger doc v1 at path api/v1/main.go)
	$(info swag i -g main.go -dir api/v1 --instanceName v1)
	swag i -dir api/v1 --instanceName v1

.PHONY: dep
dep: swagger go.mod.verify go.mod.download go.mod.tidy

.PHONY: style
style: go.mod.verify go.mod.tidy go.mod.fmt go.mod.lint.run

.PHONY: test
test: test.go

.PHONY: ci
ci: style go.mod.vet test

.PHONY: ci.test.benchmark
ci.test.benchmark: test.go.benchmark

.PHONY: ci.coverage.show
ci.coverage.show: test.go.coverage.show

.PHONY: ci.all
ci.all: ci ci.test.benchmark ci.coverage.show

.PHONY: buildMain
buildMain: swagger
	@echo "-> start build local OS"
ifeq ($(OS),Windows_NT)
	@go build -ldflags "-X main.buildID=${ENV_DIST_CODE_MARK}" -o ${ENV_ROOT_BUILD_BIN_PATH}.exe ${ENV_ROOT_BUILD_ENTRANCE}
	@echo "-> finish build out path: $(subst /,\,${ENV_ROOT_BUILD_BIN_PATH}).exe"
else
	@go build -ldflags "-X main.buildID=${ENV_DIST_CODE_MARK}" -o ${ENV_ROOT_BUILD_BIN_PATH} ${ENV_ROOT_BUILD_ENTRANCE}
	@echo "-> finish build out path: ${ENV_ROOT_BUILD_BIN_PATH}"
endif

.PHONY: buildCross
buildCross: swagger
	@echo "-> start build OS:${ENV_DIST_GO_OS} ARCH:${ENV_DIST_GO_ARCH}"
ifeq ($(ENV_DIST_GO_OS),windows)
	@GOOS=$(ENV_DIST_GO_OS) GOARCH=$(ENV_DIST_GO_ARCH) go build \
	-a \
	-tags netgo \
	-ldflags '-w -s --extldflags "-static -fpic"' \
	-o $(subst /,\,${ENV_ROOT_BUILD_BIN_PATH}).exe ${ENV_ROOT_BUILD_ENTRANCE}
	@echo "-> finish build out path: $(subst /,\,${ENV_ROOT_BUILD_BIN_PATH}).exe"
else
	@GOOS=$(ENV_DIST_GO_OS) GOARCH=$(ENV_DIST_GO_ARCH) go build \
	-a \
	-tags netgo \
	-ldflags '-w -s --extldflags "-static -fpic"' \
	-o ${ENV_ROOT_BUILD_BIN_PATH} ${ENV_ROOT_BUILD_ENTRANCE}
	@echo "-> finish build out path: ${ENV_ROOT_BUILD_BIN_PATH}"
endif

.PHONY: dev
dev: export ENV_WEB_AUTO_HOST=true
dev: cleanBuild buildMain
ifeq ($(OS),windows)
	$(subst /,\,${ENV_ROOT_BUILD_BIN_PATH}).exe ${ENV_RUN_INFO_ARGS}
else
	${ENV_ROOT_BUILD_BIN_PATH} ${ENV_RUN_INFO_ARGS}
endif

.PHONY: run.help
run.help:
	go run -v ${ENV_ROOT_BUILD_ENTRANCE} ${ENV_RUN_INFO_HELP_ARGS}

.PHONY: run
run: export GIN_MODE=test
run: export ENV_WEB_LOG_LEVEL=INFO
run: export ENV_WEB_AUTO_HOST=true
run: cleanBuild buildMain
	@echo "=> run GIN_MODE=test start"
ifeq ($(OS),windows)
	$(subst /,\,${ENV_ROOT_BUILD_BIN_PATH}).exe ${ENV_RUN_INFO_ARGS}
else
	${ENV_ROOT_BUILD_BIN_PATH} ${ENV_RUN_INFO_ARGS}
endif

.PHONY: run.release
run.release: export GIN_MODE=release
run.release: export ENV_WEB_LOG_LEVEL=INFO
run.release: export ENV_WEB_AUTO_HOST=true
run.release: cleanBuild buildMain
	@echo "=> run GIN_MODE=release start"
ifeq ($(OS),windows)
	$(subst /,\,${ENV_ROOT_BUILD_BIN_PATH}).exe ${ENV_RUN_INFO_ARGS}
else
	${ENV_ROOT_BUILD_BIN_PATH} ${ENV_RUN_INFO_ARGS}
endif

.PHONY: cloc
cloc:
	@echo "see: https://stackoverflow.com/questions/26152014/cloc-ignore-exclude-list-file-clocignore"
	cloc --exclude-list-file=.clocignore .

.PHONY: helpProjectRoot
helpProjectRoot:
	@echo "Help: Project root Makefile"
ifeq ($(OS),Windows_NT)
	@echo ""
	@echo "warning: other install make cli tools has bug, please use: scoop install main/make"
	@echo " run will at make tools version 4.+"
	@echo "windows use this kit must install tools blow:"
	@echo ""
	@echo "https://scoop.sh/#/apps?q=busybox&s=0&d=1&o=true"
	@echo "-> scoop install main/busybox"
	@echo "and"
	@echo "https://scoop.sh/#/apps?q=shasum&s=0&d=1&o=true"
	@echo "-> scoop install main/shasum"
	@echo ""
endif
	@echo "-- now build name: ${ROOT_NAME} version: ${ENV_DIST_VERSION}"
	@echo "-- distTestOS or distReleaseOS will out abi as: ${ENV_DIST_GO_OS} ${ENV_DIST_GO_ARCH} --"
	@echo ""
	@echo "~> make test                 - run test fast"
	@echo "~> make ci.all               - run CI tasks all"
	@echo "~> make ci.test.benchmark    - run CI tasks as test benchmark"
	@echo "~> make ci.coverage.show     - run CI tasks as test coverage and show"
	@echo ""
	@echo "~> make env                  - print env of this project"
	@echo "~> make init                 - check base env of this project"
	@echo "~> make dep                  - check and install by go mod"
	@echo "~> make clean                - remove build binary file, log files, and testdata"
	@echo "~> make style                - run local code fmt and style check"
	@echo "~> make ci                   - run CI tools tasks"
	@echo ""
	@echo "~> make run.help             - run use ${ENV_RUN_INFO_HELP_ARGS}"
	@echo "~> make dev                  - run as develop mode"
	@echo "~> make run.release          - run as release mode"
	@echo "~> make run                  - run as test mode"
	@echo ""

.PHONY: help
help: helpProjectRoot
	@echo "== show more help"
	@echo ""
	@echo "$$ make helpGoDist"
	@echo ""
	@echo "$$ make help.test.go.integration"
	@echo "$$ make help.test.go"
	@echo "$$ make help.go.list"
	@echo "$$ make help.go.mod"
	@echo ""
	@echo "-- more info see Makefile include --"