[![ci](https://github.com/bridgewwater/gin-api-swagger-temple/actions/workflows/ci.yml/badge.svg)](https://github.com/bridgewwater/gin-api-swagger-temple/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/bridgewwater/gin-api-swagger-temple?label=go.mod)](https://github.com/bridgewwater/gin-api-swagger-temple)
[![GoDoc](https://godoc.org/github.com/bridgewwater/gin-api-swagger-temple?status.png)](https://godoc.org/github.com/bridgewwater/gin-api-swagger-temple)
[![goreportcard](https://goreportcard.com/badge/github.com/bridgewwater/gin-api-swagger-temple)](https://goreportcard.com/report/github.com/bridgewwater/gin-api-swagger-temple)

[![GitHub license](https://img.shields.io/github/license/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/tags)
[![GitHub release)](https://img.shields.io/github/v/release/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/releases)

### cli tools to init project fast

```bash
$ curl -L --fail https://raw.githubusercontent.com/bridgewwater/gin-api-swagger-temple/main/temp-gin-api-swagger
# let temp-gin-api-swagger file folder under $PATH
$ chmod +x temp-gin-api-swagger
# see how to use
$ temp-gin-api-swagger -h
```

## for what

- this project used to gin api server
- use this template, replace list below
- [ ] rename `github.com/bridgewwater/gin-api-swagger-temple` to your api package name
    - [ ] rename `bridgewwater` to your project owner name
    - [ ] rename `gin-api-swagger-temple` to your project name
    - [ ] rename `34565` to your service port

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [X] config file as [viper](https://github.com/spf13/viper)
- [X] api tracking and version control
  by [convention-change-log](https://github.com/convention-change/convention-change-log)
    - [X] embed file `package.json` by `convention-change-log` kit
    - [X] middleware `AppVersion` will add api version for Tracking
    - [X] middleware [gin-correlation-id](https://github.com/bar-counter/gin-correlation-id) can tracking this server
      each api request
- [X] log by [zap](https://github.com/uber-go/zap) and support rotate log file
    - [X] access log at different file, and can change by `zap.rotate.AccessFilename`
    - [X] api log file, and can change by config `zap.Api.**`
- [X] server status [monitor](https://github.com/bar-counter/monitor), for help DevOps tracking server status
- [X] `major version` api support
    - [X] `api/v1` this first version of major api
- [X] error management
    - basic error generate error by [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
    - http error at different api version use different error management, v1 use by `errdef`
- [X] generate swagger doc by [swag](https://github.com/swaggo/swag), and will auto remove at `runmode=release`
- [X] server handler Exit Signal by `ctrl+c` or `kill -15 [pid]` return code 0, for safe exit.
- [X] gin unit test case example, support [golden data test](https://github.com/sebdah/goldie), you can use `-update`
  test flag to update golden data.
- [X] local build management by [make](https://www.gnu.org/software/make/), also support windows, please
  see `make helpProjectRoot` to install windows need kit.
- [X] docker build support, see `make helpDocker`, Of course, it is more recommended to use docker-compose to build a
  local development environment.
- [X] github action CI workflow check.
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## env

- minimum go version: go 1.19
- change `go 1.19`, `^1.19`, `1.19.12-bullseye` `1.19.12` to new go version

### libs

| lib                                               | version    |
|:--------------------------------------------------|:-----------|
| https://github.com/gin-gonic/gin                  | v1.9.1     |
| https://github.com/swaggo/swag                    | v2.0.0-rc3 |
| https://github.com/swaggo/gin-swagger             | v1.6.0     |
| https://github.com/spf13/pflag                    | v1.0.5     |
| https://github.com/spf13/viper                    | v1.16.0    |
| https://github.com/json-iterator/go               | v1.1.12    |
| https://github.com/uber-go/zap                    | v1.25.0    |
| https://github.com/bar-counter/monitor            | v2.2.0     |
| https://github.com/bar-counter/gin-correlation-id | v1.2.0     |

more libs see [go.mod](go.mod)

## run

> if you want auto get local IP for fast develop, you can add evn `ENV_WEB_AUTO_HOST=true`

- each new swagger must rebuild swagger doc by task `make swagger`
- also use task `make dev` or `make run` also run task buildSwagger before.
- swagger tools use [swag](https://github.com/swaggo/swag)

```bash
$ go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc3
# this will install at: echo $(go env GOBIN)/bin
# this path must in your $PATH
```

- swagger doc format
  see [https://github.com/swaggo/swag#declarative-comments-format](https://github.com/swaggo/swag#declarative-comments-format)
- swagger example
  see [https://github.com/swaggo/swag/blob/master/example/basic/api/api.go](https://github.com/swaggo/swag/blob/master/example/basic/api/api.go)

### makefile usage

- `need go mod to management golang dependenceis`

```sh
# see project root help
$ make helpProjectRoot
# see full help
$ make help

# check this project dep
$ make dep
# run all test case
$ make test
# run test case with coverage and see report
$ make testCoverage testCoverageShow
# run test case with coverage and see report by browser
$ make testCoverageBrowser

# run server as dev
$ make dev

# check before, then push to CI build
$ make dep ci

## docker build support

# - first use can pull images
$ make dockerAllPull

# - test run container use ./build.dockerfile
$ make dockerTestBuildCheck

# - then can run as docker-compose build image and up
$ make dockerComposeUp
# - then see log as docker-compose
$ make dockerComposeFollowLogs
# - down as docker-compose will auto remove local image
$ make dockerComposeDown

# - prune test container and image
$ make dockerTestPruneLatest
```

most of the doc at [http://127.0.0.1:34565/swagger/index.html](http://127.0.0.1:34565/swagger/index.html)

## config

- config file is `config.yaml` demo see [conf/config.yaml](conf/config.yaml)

### log

- use log as: [https://github.com/uber-go/zap](https://github.com/uber-go/zap)

```yaml
# zap config
zap:
  AtomicLevel: -1 # DebugLevel:-1 InfoLevel:0 WarnLevel:1 ErrorLevel:2
  Api:
    PrefixPaths: "/api/v1/" # api path prefix list
    AtomicLevel: 0 # DebugLevel:-1 InfoLevel:0 WarnLevel:1 ErrorLevel:2 default 0
  FieldsAuto: false # is use auto Fields key set
  Fields:
    Key: key
    Val: val
  Development: true # is open file and line number
  Encoding: console # output format, only use console or json, default is console
  rotate:
    Filename: logs/template-gitea-gin-api.log # Log file path
    # AccessFilename: logs/access.log # Access log file path
    # ApiFilename: logs/api.log # api log file path
    MaxSize: 16 # Maximum size of each zlog file, Unit: M
    MaxBackups: 10 # How many backups are saved in the zlog file
    MaxAge: 7 # How many days can the file be keep, Unit: day
    Compress: true # need compress
  EncoderConfig:
    TimeKey: time
    LevelKey: level
    NameKey: logger
    CallerKey: caller
    MessageKey: msg
    StacktraceKey: stacktrace
    TimeEncoder: ISO8601TimeEncoder # ISO8601TimeEncoder EpochMillisTimeEncoder EpochNanosTimeEncoder EpochTimeEncoder default is ISO8601TimeEncoder
    EncodeDuration: SecondsDurationEncoder # NanosDurationEncoder SecondsDurationEncoder StringDurationEncoder default is SecondsDurationEncoder
    EncodeLevel: CapitalColorLevelEncoder # CapitalLevelEncoder CapitalColorLevelEncoder LowercaseColorLevelEncoder LowercaseLevelEncoder default is CapitalLevelEncoder
    EncodeCaller: ShortCallerEncoder # ShortCallerEncoder FullCallerEncoder default is FullCallerEncoder

```

## folder-structure

Project file definition

```
.
├── LICENSE                     # license
├── .golangci.yaml              # golangci-lint config
├── Dockerfile                  # ci build
├── docker-compose.yml          # local development docker-compose
├── build.dockerfile            # local docker build enter
├── z-MakefileUtils             # Makefile tool library
├── Makefile                    # Makefile file, using make as a compilation tool
├── README.md
├── api                         # api management
│   ├── middleware                # api middleware directory
│   │   ├── app_version.go          # app version tracking middleware, use package.json to manage api version, header: X-App-Version
│   │   ├── header.go               # header middleware, include: options secure noCache etc.
│   │   ├── monitor.go              # monitor middleware, use https://github.com/bar-counter/monitor
│   │   └── usage.go                # usage middleware for Gin engine.
│   │
│   │                           # Each major version of api will be distinguished here, so in this directory, there will be multiple implementations of the same function.
│   └── v1                        # api /v1 directory
│       ├── auth                    # under api/v1 auth api, The authentication method of each major version is inconsistent.
│       ├── errdef                  # under api/v1 err define api, The error code of each major version is inconsistent.
│       ├── handler                 # under api/v1 api, Similar to C in MVC architecture, it is used to read input, forward the processing flow to the actual processing function, and finally return the result.
│       │   ├── biz                    # api group folder for /biz
│       │   ├── json.go                # json parse
│       │   └── jsonResponse.go        # universal response structure, The authentication method of each major version is inconsistent.
│       ├── model                   # under api/v1 model define api, The model of each major version is inconsistent.
│       └── main.go                 # swag generated file entrance, swag base info update here
├── cmd                         # cmd folder
│   └── gin-api-swagger-temple    # package of this web app
│       ├── main.go                 # app program entrance
│       └── main_test.go            # app integration test entrance
│
├── conf                        # Configuration files are stored in a unified directory
│   ├── config.yaml
│   ├── release                   # release configuration file
│   └── test                      # test configuration file
├── build                       # build directory, which is not in the git list
├── dist                        # Publish the directory, which is not in the git list
├── doc                         # API document directory
│   ├── README.md
│   ├── monitor.md
│   ├── supervisor.md
│   └── systemctl.md
├── docs                        # swagger build directory, which is generated by cli swag
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod                      # go.mod file
├── logs                        # log directory not in git management list
├── resource.go                 # embed file entrance, such as: html, js, css, image, json, etc.
└── internal                    # internal tool directory
    ├── config                    # Dedicated to handling configuration and configuration files Go package
    │   ├── baseConf.go
    │   ├── config.go
    │   ├── logConf.go
    │   └── watchConf.go
    ├── folder                    # OS path tools
    ├── parsehttp                 # parse http request
    ├── pkg                       # referenced package
    └── sys                       # system info tools
```

## development skills

### goland auto generate swagger doc

- open `Settings` -> `Tools` -> `File Watchers` -> `+` -> `Custom`
- new `File Watcher` as `swag api/v1`
    - Files to Watch
        - name `swag api/v1`
        - file type `Go files`
        - scope `Project Files`
    - Tools to Run on Changes
        - Program `$GOPATH$/bin/swag`
        - Arguments `i -g main.go -dir api/v1 --instanceName v1`
        - Working directory `$ProjectFileDir$`

![](https://github.com/bridgewwater/gin-api-swagger-temple/raw/main/doc/img/goland-swag-auto-v1.png)
