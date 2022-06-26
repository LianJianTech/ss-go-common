package errno

var (
	OK          = &Errno{Code: 0, Msg: "ok"}
	AuthError   = &Errno{Code: 1, Msg: "auth error"}
	ParamError  = &Errno{Code: 2, Msg: "request param error"}
	HandleError = &Errno{Code: 3, Msg: "handle error"}
)
