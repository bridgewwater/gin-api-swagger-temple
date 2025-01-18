## env

- minimum go version: go 1.22
- change `go 1.22`, `^1.22`, `1.22.11` to new go version
- change `golangci-lint@v1.59.1` from [golangci-lint version release](https://github.com/golangci/golangci-lint/releases) to new version
    - more info see [golangci-lint local-installation](https://golangci-lint.run/usage/install/#local-installation)
- change swag version `github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc4`
    - more info see [github.com/swaggo/swag/releases](https://github.com/swaggo/swag/releases)

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
$ go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc4
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
$ make test.go.coverage test.go.coverage.show
# run test case with coverage and see report by browser
$ make test.go.coverage.browser

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