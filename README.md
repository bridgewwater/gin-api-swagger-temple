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
```

# config

- config file is `config.yaml` demo see [conf/config.yaml](conf/config.yaml)
- upper `swagger_index` do not set!

## log

+ `writers`: 输出位置，有2个可选项：file,stdout。选择file会将日志记录到`logger_file`指定的日志文件中，选择stdout会将日志输出到标准输出，当然也可以两者同时选择
+ `logger_level`: 日志级别，DEBUG, INFO, WARN, ERROR, FATAL
+ `logger_file`: 日志文件
+ `log_format_text`: 日志的输出格式，json或者plaintext，`true`会输出成json格式，`false`会输出成非json格式
+ `rollingPolicy`: rotate依据，可选的有：daily, size。如果选daily则根据天进行转存，如果是size则根据大小进行转存
+ `log_rotate_date`: rotate转存时间，配合`rollingPolicy: daily`使用
+ `log_rotate_size`: rotate转存大小，配合`rollingPolicy: size`使用
+ `log_backup_count`:当日志文件达到转存标准时，log系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数。

# dev

- swagger tools use [swag](https://github.com/swaggo/swag)
```sh
go get -v -u github.com/swaggo/swag/cmd/swag
```

- swagger doc see [https://swaggo.github.io/swaggo.io/declarative_comments_format/](https://swaggo.github.io/swaggo.io/declarative_comments_format/)

## evn

```bash
go version go1.11.4 darwin/amd64
gin version 1.3.0
swag version v1.4.1
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
│   ├── sd                       # 健康检查handler
│   │   └── check.go
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
│   │   ├── code.go
│   │   └── errno.go
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
