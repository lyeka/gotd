# thinking
关于项目设计实现的一些思考

## 目录
- 生成唯一ID
- 密码加密问题
- JWT


### 生成唯一ID

### 密码加密
代码参见 `pkg/encrypt`。
同样明文每次加密后的密文都不一致, 但是使用比对方法却可以正确判定密码是否正确。算法？
ref
- https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
- https://segmentfault.com/a/1190000016748637

### JWT
[jwt介绍](https://jwt.io/introduction/)
[jwt介绍（中文）](https://www.jianshu.com/p/576dbf44b2ae)
[jwt-go](https://github.com/dgrijalva/jwt-go)

### token 为什么要加上 Bearer
jwt token 一般放置在 http header 中的 `Authorization` 字段，而且 token 前还会加上 `Bearer `， 为什么需要这个 `Bearer `
因为JWT是一种token格式，Bearer Token是一种鉴权方案。
HTTP的header项「Authorization」是在HTTP 1.0引入的，它的值的格式是类型+token，支持多种鉴权方案，bearer只是其中的一种。bearer方案中使用的token是JWT格式，这就是它们之间的关系。

ref
- http://0x3f.org/post/introduction-of-json-web-token/