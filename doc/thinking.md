# thinking
关于项目设计实现的一些思考

## 目录
- 生成唯一ID
- 密码加密问题
- JWT
- 状态变量选status还是state？
- go mongo driver 的使用姿势？


### 生成唯一ID
MongoDB 提供了生成 objectID 的方法， 文档见 [ObjectId](https://docs.mongodb.com/manual/reference/method/ObjectId/)
这里 ObjectID 的组成已经随着版本更改了，主要是将主机名+进程号的组合改成了随机数, 可以参考[Mongo ObjectId 早就不用机器标识和进程号了](https://blog.wolfogre.com/posts/mongo-objectid-design/)

现在存在的问题是 
- 是否可以使用这个 ObjectID 作为唯一标识？作为唯一标识的话冲突情况如何？
- 如果不用 ObjectID 的话，其它生成 UUID 的选择？


目前项目的选择是用户与任务均使用 ObjectID


### 密码加密
代码参见 `pkg/encrypt`。
同样明文每次加密后的密文都不一致, 但是使用比对方法却可以正确判定密码是否正确。算法？
ref
- https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
- https://segmentfault.com/a/1190000016748637

### JWT
一些资料
[jwt介绍](https://jwt.io/introduction/)
[jwt介绍（中文）](https://www.jianshu.com/p/576dbf44b2ae)
[jwt-go](https://github.com/dgrijalva/jwt-go)


token 为什么要加上 Bearer
jwt token 一般放置在 http header 中的 `Authorization` 字段，而且 token 前还会加上 `Bearer `， 为什么需要这个 `Bearer `
因为JWT是一种token格式，Bearer Token是一种鉴权方案。
HTTP的header项「Authorization」是在HTTP 1.0引入的，它的值的格式是类型+token，支持多种鉴权方案，bearer只是其中的一种。bearer方案中使用的token是JWT格式，这就是它们之间的关系。

ref
- http://0x3f.org/post/introduction-of-json-web-token/

### 状态变量选status还是state？
目前我的理解是state表示可以迁移的状态， 如任务的开始——进行中——结束；
status表示不可迁移（或者说迁移没有依赖）的状态，如心情的悲伤——快乐——痛苦——平静

ref
- [程序代码中，怎么区分status和state？](https://www.zhihu.com/question/21994784)


### go mongo driver 的使用姿势？
todo