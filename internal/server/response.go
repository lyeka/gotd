package server

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 响应内容
}

const (
	CodeOK = iota
	CodeErr
	CodeInvalidParam
	CodeErrCreateTask
	CodeErrUpdateTask
	CodeErrGetUserTask
)
