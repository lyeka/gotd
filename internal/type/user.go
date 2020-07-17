package _type

type User struct {
	ID       int64  // 为了兼容 mongodb 的 _id ，在此使用字符串作为用户标识
	Email    string // 邮箱
	Nickname string // 用户昵称
	Password string // 密码
}
