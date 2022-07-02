package errno

var (
	OK           = &Errno{Code: 0, Msg: "ok"}
	ParamError   = &Errno{Code: 1, Msg: "request param error"}
	TokenError   = &Errno{Code: 2, Msg: "token error"}
	AccountError = &Errno{Code: 3, Msg: "account or password error"}
	HandleError  = &Errno{Code: 4, Msg: "handle error"}
	OtherError   = &Errno{Code: 5, Msg: "other error"}
)
