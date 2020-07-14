package _type

type User struct {
	Id string // 为了兼容 mongodb 的 _id ，在此使用 字符串作为用户标识
	Email string // 邮箱
	NickName string // 用户昵称
	Password string // 密码
}
