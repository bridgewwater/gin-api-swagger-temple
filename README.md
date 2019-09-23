## for what

- this project used to gin api server
- [ ] rename `git.sinlov.cn/bridgewwater/temp-gin-api-self` to your api package name

## use

- `need dep to management golang dependenceis`, will change to go mod

```sh
$ make help
# check base dep
$ make init
# first run just use
$ make checkDepends
# change conf/config.yaml

# run server as dev
$ make dev
# run as docker contant
$ make dockerRun
# stop or remove docker
$ make dockerStop
$ make dockerRemove
```

most of doc at [http://127.0.0.1:39000/swagger/index.html](http://127.0.0.1:39000/swagger/index.html)

# dev

- swagger tools use [swag](https://github.com/swaggo/swag)
```sh
go get -v -u github.com/swaggo/swag/cmd/swag
```

- swagger doc see [https://swaggo.github.io/swaggo.io/declarative_comments_format/](https://swaggo.github.io/swaggo.io/declarative_comments_format/)
- swagger example see [https://github.com/swaggo/swag/blob/master/example/basic/api/api.go](https://github.com/swaggo/swag/blob/master/example/basic/api/api.go)

## evn

```bash
go version go1.11.4 darwin/amd64
gin version 1.3.0
swag version v1.4.1
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
├── Gopkg.lock                   # dep 依赖锁文件
├── Gopkg.toml                   # dep 依赖管理文件
├── admin.sh                     # 进程的start|stop|status|restart控制文件,用于 linux 集成
├── conf                         # 配置文件统一存放目录
│   ├── config.yaml              # 配置文件
│   ├── server.crt               # TLS配置文件
│   └── server.key
├── config                       # 专门用来处理配置和配置文件的Go package
│   └── config.go
├── db                           # 在部署新环境时数据库使用
│   ├── mongo.sh                 # 部署 mongo 数据库
│   └── db.sql                   # 可以登录MySQL客户端，执行source db.sql创建数据库和表
├── docs                         # swagger文档，执行 swag init 生成的, 不可自行修改
│   ├── docs.go
│   └── swagger
│       ├── swagger.json
│       └── swagger.yaml
├── handler                      # 类似MVC架构中的C，用来读取输入，并将处理流程转发给实际的处理函数，最后返回结果
│   ├── handler.go
│   ├── biz                      # 业务范例 handler service status check
│   │   └── biz.go
│   └── user                     # 核心：用户业务逻辑handler
│       ├── create.go            # 新增用户
│       ├── delete.go            # 删除用户
│       ├── get.go               # 获取指定的用户信息
│       ├── list.go              # 查询用户列表
│       ├── login.go             # 用户登录
│       ├── update.go            # 更新用户
│       └── user.go              # 存放用户handler公用的函数、结构体等
├── main.go                      # Go程序唯一入口
├── Makefile                     # Makefile文件，使用make来作为编译工具
├── model                        # 数据库相关的操作统一放在这里，包括数据库初始化和对表的增删改查
│   ├── init.go                  # 初始化和连接数据库
│   ├── model.go                 # 存放一些公用的go struct
│   └── user.go                  # 用户相关的数据库CURD操作
├── pkg                          # 引用的包
│   ├── auth                     # 认证包
│   │   └── auth.go
│   ├── constvar                 # 常量统一存放位置
│   │   └── constvar.go
│   ├── errdef                   # 错误码存放位置
│   │   ├── errcode.go           # 错误码添加
│   │   └── errdef.go            # 错误码定义及辅助函数
│   ├── token
│   │   └── token.go
│   └── version                  # 版本包
│       ├── base.go
│       ├── doc.go
│       └── version.go
├── README.md                    # 工程目录README
├── router                       # 路由相关处理
│   ├── middleware               # API服务器用的是Gin Web框架，Gin中间件存放位置
│   │   ├── auth.go
│   │   ├── header.go
│   │   ├── logging.go
│   │   └── requestid.go
│   └── router.go
├── service                      # 实际业务处理函数存放位置
│   └── service.go
├── util                         # 工具类函数存放目录
│   ├── util.go
│   └── util_test.go
└── vendor                         # vendor目录用来管理依赖包 这里使用 dep 管理
    ├── github.com
    ├── golang.org
    ├── gopkg.in
    └── vendor.json
```
