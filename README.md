### 技术点

车载后端系统目前运用到如下的技术点：(后续看需求会继续有新东西引入)

gin web框架、goredis 缓存、gorm 数据库orm

restfulapi api风格、protobuf 数据通讯协议、emq 长连接

jwt鉴权、gomod 本地管理依赖库、cron 定时任务、swagger 测试人员使用


### gomod用法：

set GO111MODULE=on			//打开包管理

set GOPROXY= https://goproxy.io		//设置下载包代理

go mod download: 下载依赖的 module 到本地 cache

go mod edit: 编辑 go.mod

go mod graph: 打印模块依赖图

go mod init: 在当前目录下初始化 go.mod(就是会新建一个 go.mod 文件)

go mod tidy: 整理依赖关系，会添加丢失的 module，删除不需要的 module

go mod vender: 将依赖复制到 vendor 下

go mod verify: 校验依赖

go mod why: 解释为什么需要依赖

### 项目目录层级结构：

```
.
├── assets 存放一些资源，比如数据库脚本
│   ├── doc_tips
│   └── vehicle.sql
├── bin
├── conf.ini  主服务的配置文件
├── conf.txt  emq测试各个配置文件
├── go.mod    gomod文件
├── go.sum
├── pkg
├── src
│   ├── vehicle  主服务包
│   │   ├── api_server  各个api包
│   │   ├── conf        配置文件读取
│   │   ├── cron        定时任务
│   │   ├── db          gorm，goredis封装目录
│   │   ├── docs        swagger文档
│   │   ├── emq         emq包
│   │   ├── logger      logger工具包
│   │   ├── main.go     主入口文件
│   │   ├── middleware  中间件包
│   │   ├── model       model层，baseModel接口统一封装
│   │   ├── response    返回对象封装
│   │   ├── router      路由分组，路由版本
│   │   ├── service     服务，比如jwt鉴权，cors，flow流
│   │   ├── timing      倒计时库
│   │   └── util        工具类
│   └── vehicle_script  服务脚本包
│       ├── api_script  对外接口的测试脚本
│       ├── emq_script  emq模拟脚本
│       ├── emq_service 
│       └── tool        服务脚本工具
└── vlog
    └── 2020413.log
```



