package errno

var (
	OK           = &Errno{Code: 0, Msg: "成功"}
	ParamError   = &Errno{Code: 1, Msg: "请求参数错误"}
	TokenError   = &Errno{Code: 2, Msg: "token错误"}
	CodeError    = &Errno{Code: 3, Msg: "验证码错误"}
	AccountError = &Errno{Code: 4, Msg: "账号或密码错误"}
	HandleError  = &Errno{Code: 5, Msg: "服务处理错误"}
	OtherError   = &Errno{Code: 6, Msg: "其他错误"}
)
