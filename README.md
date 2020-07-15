# GOTD
Go 实现 TODO 类应用，C/S 架构。

## 服务端

## 特性
- 数据存储层支持切换
    - mongodb
    - redis（持久化）
    - mysql
- 数据存储、传输格式支持切换
    - protobuf
    - json
- websocket 通信
- 协同编辑同步

## 依赖
- Go
- MongoDB/MySQL/Redis 依赖于你选择什么数据存储引擎

### todo
- 撰写数据存储引擎切换时用于数据迁移的脚本
- swagger 管理 api？
- mongodb生成唯一id的算法
    - 依赖redis生成？
    - 分布式唯一id生成算法——snowflake?
- 补全测试&TDD
- 不登录也可以使用（功能受限）

## 客户端-终端 GUI

### 相关库
GUI: [gocui](https://github.com/jroimartin/gocui)
存储: [badger](https://github.com/dgraph-io/badger)

### 原型图
![](https://github.com/lyeka/gotd/blob/master/ui.png)