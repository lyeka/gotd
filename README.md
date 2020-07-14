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

## 客户端-终端 GUI

### 相关库
GUI: [gocui](https://github.com/jroimartin/gocui)
存储: [badger](https://github.com/dgraph-io/badger)

### 原型图
![](https://github.com/lyeka/gotd/blob/master/ui.png)