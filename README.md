[![ci](https://github.com/bridgewwater/gin-api-swagger-temple/actions/workflows/ci.yml/badge.svg)](https://github.com/bridgewwater/gin-api-swagger-temple/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/bridgewwater/gin-api-swagger-temple?label=go.mod)](https://github.com/bridgewwater/gin-api-swagger-temple)
[![GoDoc](https://godoc.org/github.com/bridgewwater/gin-api-swagger-temple?status.png)](https://godoc.org/github.com/bridgewwater/gin-api-swagger-temple)
[![goreportcard](https://goreportcard.com/badge/github.com/bridgewwater/gin-api-swagger-temple)](https://goreportcard.com/report/github.com/bridgewwater/gin-api-swagger-temple)

[![GitHub license](https://img.shields.io/github/license/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple)
[![codecov](https://codecov.io/gh/bridgewwater/gin-api-swagger-temple/branch/main/graph/badge.svg)](https://codecov.io/gh/bridgewwater/gin-api-swagger-temple)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/tags)
[![GitHub release)](https://img.shields.io/github/v/release/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/releases)

## for what

- this project used to gin api server
- [ ] rename `github.com/bridgewwater/gin-api-swagger-temple` to your api package name

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [X] server handler Exit Signal by `ctrl+c` or `kill -15 [pid]` return code 0, for safe exit.
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

- use this template, replace list below
  - `github.com/bridgewwater/gin-api-swagger-temple` to your package name
  - `bridgewwater` to your owner name
  - `gin-api-swagger-temple` to your project name

## env

- minimum go version: go 1.18
- change `go 1.18`, `^1.18`, `1.18.10` to new go version

### libs

| lib                                               | version    |
|:--------------------------------------------------|:-----------|
| https://github.com/gin-gonic/gin                  | v1.9.1     |
| https://github.com/swaggo/swag                    | v2.0.0-rc3 |
| https://github.com/swaggo/gin-swagger             | v1.6.0     |
| https://github.com/spf13/pflag                    | v1.0.5     |
| https://github.com/spf13/viper                    | v1.16.0    |
| https://github.com/json-iterator/go               | v1.1.12    |
| https://github.com/bar-counter/slog               | v1.4.0     |
| https://github.com/bar-counter/monitor            | v2.2.0     | 
| https://github.com/bar-counter/gin-correlation-id | v1.2.0     | 

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
$ make help
# check base dep
$ make init
# first run just use
$ make dep
# change conf/config.yaml

# run server as dev
$ make dev

# - first use can pull images
$ make dockerAllPull

# - test run container use ./build.dockerfile
$ make dockerTestBuildCheck

# - then can run as docker-compose
$ make dockerComposeUp
# - then see log as docker-compose
$ make dockerComposeFollowLogs
# - down as docker-compose
$ make dockerComposeDown

# - prune test container and image
$ make dockerTestPruneLatest
```

most of doc at [http://127.0.0.1:34567/swagger/index.html](http://127.0.0.1:34567/swagger/index.html)

## config

- config file is `config.yaml` demo see [conf/config.yaml](conf/config.yaml)

### log

- use log as: [http://github.com/bar-counter/slog](http://github.com/bar-counter/slog)

```yaml
log:
  writers: file,stdout            # file,stdout。`file` will let `logger_file` to file，`stdout` will show at std, most of time use bose
  logger_level: DEBUG             # log level: DEBUG, INFO, WARN, ERROR, FATAL
  logger_file: log/server.log     # log file setting
  log_hide_lineno: false # `true` will hide code line number, `false` will show code line number, default is false
  log_format_text: true # format_text `false` will format json, `true` will out stdout
  rolling_policy: size            # rotate policy, can choose as: daily, size. `daily` store as daily，`size` will save as max
  log_rotate_date: 1              # rotate date, coordinate `rollingPolicy: daily`
  log_rotate_size: 8              # rotate size，coordinate `rollingPolicy: size`
  log_backup_count: 7             # backup max count, log system will compress the log file when log reaches rotate set, this set is max file count
```

## folder-structure

Project file definition

```
.
├── LICENCE                     # licence
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
