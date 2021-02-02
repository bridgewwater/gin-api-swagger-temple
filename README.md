## for what

- this project used to gin api server
- [ ] rename `github.com/bridgewwater/gin-api-swagger-temple` to your api package name

## use

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
# run as docker contant
$ make dockerRunLinux
# if use macOS
$ make dockerRunDarwin
# stop or remove docker
$ make dockerStop
$ make dockerRemove
```

most of doc at [http://127.0.0.1:39000/swagger/index.html](http://127.0.0.1:39000/swagger/index.html)

# dev

> if want auto get local IP for fast develop, you can add evn `ENV_WEB_AUTO_HOST=true`

- each new swagger must rebuild swagger doc by task `make buildSwagger`
- also use task `make dev` or `make runTest` also run task buildSwagger before.
- swagger tools use [swag](https://github.com/swaggo/swag)
```sh
go get -v -u github.com/swaggo/swag/cmd/swag
```

| lib | version |
|:---------------------|:---|
| github.com/swaggo/swag | v1.6.2 |
| github.com/swaggo/gin-swagger | v1.3.0 |

- swagger doc see [https://swaggo.github.io/swaggo.io/declarative_comments_format/](https://swaggo.github.io/swaggo.io/declarative_comments_format/)
- swagger example see [https://github.com/swaggo/swag/blob/master/example/basic/api/api.go](https://github.com/swaggo/swag/blob/master/example/basic/api/api.go)

## evn

```bash
go version go1.15.6 darwin/amd64
swag version v1.6.2
gin version v1.4.0
```

# config

- config file is `config.yaml` demo see [conf/config.yaml](conf/config.yaml)

## log

```yaml
log:
  writers: file,stdout            # file,stdout。`file` will let `logger_file` to file，`stdout` will show at std, most of time use bose
  logger_level: DEBUG             # log level: DEBUG, INFO, WARN, ERROR, FATAL
  logger_file: log/server.log     # log file setting
  log_format_text: false          # format `false` will format json, `true` will show abs
  rollingPolicy: size             # rotate policy, can choose as: daily, size. `daily` store as daily，`size` will save as max
  log_rotate_date: 1              # rotate date, coordinate `rollingPolicy: daily`
  log_rotate_size: 8              # rotate size，coordinate `rollingPolicy: size`
  log_backup_count: 7             # backup max count, log system will compress the log file when log reaches rotate set, this set is max file count
```

## folder-Def

工程文件定义

```
.
├── Dockerfile
├── LIB.md
├── MakeDockerRun.mk
├── MakeGoMod.mk
├── Makefile                    # Makefile文件，使用make来作为编译工具
├── README.md
├── conf                        # 配置文件统一存放目录
│   ├── config.yaml
│   ├── release
│   └── test
├── config                      # 专门用来处理配置和配置文件的Go package
│   ├── baseConf.go
│   ├── config.go
│   ├── logConf.go
│   └── watchConf.go
├── dist                        # 发布目录，不在 git 管理列表中
├── doc                         # 工程文档目录
│   ├── README.md
│   ├── monitor.md
│   ├── supervisor.md
│   └── systemctl.md
├── docker-compose.yml
├── docs                        # swagger 生成目录
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── handler                     # 类似MVC架构中的C，用来读取输入，并将处理流程转发给实际的处理函数，最后返回结果
│   ├── biz
│   ├── json.go
│   └── jsonResponse.go
├── log                         # 日志目录，不在 git 管理列表中
├── main.go                     # Go程序唯一入口
├── model                       # 数据库相关的操作统一放在这里，包括数据库初始化和对表的增删改查
│   ├── biz
│   ├── dbMongoOffical.md
│   ├── dbRedis.md
│   └── response.go
├── pkg                         # 引用的包
│   ├── auth
│   └── errdef
├── router                      # 路由相关处理
│   ├── api.go                  # api 都放在这个目录
│   ├── middleware              # 中间件存放目录
│   ├── monitor.go              # github.com/bar-counter/monitor 实现目录
│   ├── router.go               # router 入口
│   └── swagger.go              # 参数化 swagger 实现
└── util                        # 工具目录
    ├── parsehttp
    └── sys
```
