车载后端系统目前运用到如下的技术点：(后续看需求会继续有新东西引入)
gin、goredis、gorm、restfulapi、protobuf、emq、jwt、gomod、cron、swagger

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



