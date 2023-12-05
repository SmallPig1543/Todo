
### 项目结构
```shell
├─.idea
├─api
├─conf
├─log
├─middleware
├─pkg
│  ├─ctl
│  ├─e
│  └─util
├─repository
│  ├─cache
│  └─db
│      ├─dao
│      └─model
├─route
├─service
└─types
```

- api接口
- conf redis和mysql的配置
- log日志打印的地址
- middleware 存放中间件
- pkg/ctl 存放返回结构
- pkg/e 错误码
- pkg/util jwt和日志实现
- repository 存放redis和mysql的操作
- route 路由组
- service 接口函数实现
- types 请求结构体封装

### 接口文档
https://apifox.com/apidoc/shared-cac07b03-37fd-47b8-b0fb-0d4a8939adfc

##### 该项目可部署在docker上