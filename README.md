# remarks_monitor

一款基于gozero的微服务分布式言论管控平台，搭载有完整的链路追踪、服务监控和日志管理系统

应用场景：在2077年，荒坂集团为了对夜之城的居民的言论进行严格的管控，针对 **用户名，访问的域名，违禁词，发送时间**进行监测并方便实施打击



README持续施工中。。



------

## 技术栈

| 功能               | 实现                                      |
| :----------------- | :---------------------------------------- |
| http框架           | gozero                                    |
| rpc框架            | gozero                                    |
| 数据库             | MongoDB、MySQL、Redis                     |
| 访问控制           | casbin                                    |
| 服务发现与配置中心 | consul                                    |
| 链路追踪           | jaeger                                    |
| 服务监控           | prometheus，grafana                       |
| 消息队列           | kafka                                     |
| 日志管理           | filebeat，go-stash，elasticsearch，kibana |
| 网关               | traefik                                   |

## 架构图

https://excalidraw.com/#json=fNEX_332YTuzOPvuiYCes,ArMrpKtpwo8e8LfNC_BR2A

![image-20230616215656498](C:\Users\ViolaPioggia\AppData\Roaming\Typora\typora-user-images\image-20230616215656498.png)

![image-20230616215713507](C:\Users\ViolaPioggia\AppData\Roaming\Typora\typora-user-images\image-20230616215713507.png)

## 结构介绍

├─app
│  ├─casbin **权限控制**
│  ├─input **输入**
│  │  ├─cmd
│  │  │  ├─api
│  │  │  │  ├─etc
│  │  │  │  └─internal
│  │  │  │      ├─config
│  │  │  │      ├─handler
│  │  │  │      ├─logic
│  │  │  │      ├─svc
│  │  │  │      └─types
│  │  │  └─rpc
│  │  │      ├─etc
│  │  │      ├─input
│  │  │      ├─internal
│  │  │      │  ├─config
│  │  │      │  ├─logic
│  │  │      │  ├─server
│  │  │      │  └─svc
│  │  │      └─pb
│  │  └─model
│  ├─map1 **map工作节点**
│  │  └─cmd
│  │      └─rpc
│  │          ├─etc
│  │          ├─internal
│  │          │  ├─config
│  │          │  ├─logic
│  │          │  ├─server
│  │          │  └─svc
│  │          ├─map1
│  │          └─mapwork
│  ├─map2
│  │  └─cmd
│  │      └─rpc
│  │          ├─etc
│  │          ├─internal
│  │          │  ├─config
│  │          │  ├─logic
│  │          │  ├─server
│  │          │  └─svc
│  │          ├─map2
│  │          └─mapwork
│  ├─map3
│  │  └─cmd
│  │      └─rpc
│  │          ├─etc
│  │          ├─internal
│  │          │  ├─config
│  │          │  ├─logic
│  │          │  ├─server
│  │          │  └─svc
│  │          ├─map3
│  │          └─mapwork
│  ├─master **数据处理管理节点**
│  │  ├─cmd
│  │  │  ├─api
│  │  │  │  ├─etc
│  │  │  │  └─internal
│  │  │  │      ├─config
│  │  │  │      ├─handler
│  │  │  │      │  └─master
│  │  │  │      ├─logic
│  │  │  │      │  └─master
│  │  │  │      ├─svc
│  │  │  │      └─types
│  │  │  └─rpc
│  │  │      ├─etc
│  │  │      ├─internal
│  │  │      │  ├─config
│  │  │      │  ├─logic
│  │  │      │  ├─server
│  │  │      │  └─svc
│  │  │      ├─master
│  │  │      └─pb
│  │  └─model
│  ├─reduce1 **reduce工作节点**
│  │  └─cmd
│  │      └─rpc
│  │          ├─etc
│  │          ├─internal
│  │          │  ├─config
│  │          │  ├─logic
│  │          │  ├─server
│  │          │  └─svc
│  │          ├─reduce1
│  │          └─reducework1
│  ├─reduce2
│  │  └─cmd
│  │      └─rpc
│  │          ├─etc
│  │          ├─internal
│  │          │  ├─config
│  │          │  ├─logic
│  │          │  ├─server
│  │          │  └─svc
│  │          ├─reduce2
│  │          └─reducework2
│  └─usercenter **用户中心**
│      ├─cmd
│      │  ├─api
│      │  │  ├─etc
│      │  │  └─internal
│      │  │      ├─config
│      │  │      ├─handler
│      │  │      │  └─user
│      │  │      ├─logic
│      │  │      │  └─user
│      │  │      ├─svc
│      │  │      └─types
│      │  └─rpc
│      │      ├─etc
│      │      ├─internal
│      │      │  ├─config
│      │      │  ├─logic
│      │      │  ├─server
│      │      │  └─svc
│      │      ├─pb
│      │      └─usercenter
│      └─model
├─common
│  ├─ctxdata **JWT相关**
│  └─tool **工具**
├─data
│  ├─elasticsearch **日志处理数据**
│  │  └─data
│  │      └─nodes
│  │          └─0
│  ├─grafana **监控数据**
│  │  └─data
│  │      ├─csv
│  │      ├─plugins
│  │      └─png
│  ├─prometheus **监控数据**
│  │  └─data
│  └─remarks_monitor **数据处理**
│      ├─input
│      ├─map_content
│      ├─map_domain
│      ├─map_username
│      └─reduce
│          ├─content
│          ├─domain
│          └─username
├─deploy **组件管理**
│  ├─filebeat
│  │  └─conf
│  ├─go-stash
│  │  └─etc
│  ├─prometheus
│  │  └─server
│  └─traefik
├─doc 
└─test **测试**

## 功能演示

### traefik

![image-20230616222308042](C:\Users\ViolaPioggia\AppData\Roaming\Typora\typora-user-images\image-20230616222308042.png)

### prometheus

![image-20230616222346318](C:\Users\ViolaPioggia\AppData\Roaming\Typora\typora-user-images\image-20230616222346318.png)

### jaeger

![image-20230616222416867](C:\Users\ViolaPioggia\AppData\Roaming\Typora\typora-user-images\image-20230616222416867.png)

### consul

![image-20230616222500317](C:\Users\ViolaPioggia\AppData\Roaming\Typora\typora-user-images\image-20230616222500317.png)