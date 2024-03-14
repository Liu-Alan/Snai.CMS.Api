package msg

const (
	Success          = 200
	Fail             = 100
	Error            = 500
	InvalidParams    = 400
	BindParamsError  = 401
	ValidParamsError = 402

	RequestError     = 10000
	AuthNotExist     = 10001
	AuthCheckTimeout = 10002
	AuthCheckFail    = 10003
	AuthFormatFail   = 10004
)
