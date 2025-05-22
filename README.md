[![ci](https://github.com/bridgewwater/gin-api-swagger-temple/actions/workflows/ci.yml/badge.svg)](https://github.com/bridgewwater/gin-api-swagger-temple/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/bridgewwater/gin-api-swagger-temple?label=go.mod)](https://github.com/bridgewwater/gin-api-swagger-temple)
[![GoDoc](https://godoc.org/github.com/bridgewwater/gin-api-swagger-temple?status.png)](https://godoc.org/github.com/bridgewwater/gin-api-swagger-temple)
[![goreportcard](https://goreportcard.com/badge/github.com/bridgewwater/gin-api-swagger-temple)](https://goreportcard.com/report/github.com/bridgewwater/gin-api-swagger-temple)

[![GitHub license](https://img.shields.io/github/license/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/tags)
[![GitHub release)](https://img.shields.io/github/v/release/bridgewwater/gin-api-swagger-temple)](https://github.com/bridgewwater/gin-api-swagger-temple/releases)

### cli tools to init project fast

```bash
$ v=1.2.0; curl -L --fail https://raw.githubusercontent.com/bridgewwater/gin-api-swagger-temple/v$v/temp-gin-api-swagger -o temp-gin-api-swagger
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

## dev

- see [dev.md](doc-dev/dev.md)

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
