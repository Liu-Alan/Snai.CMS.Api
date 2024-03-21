package message

const (
	Success          = 200
	Fail             = 100
	Error            = 500
	InvalidParams    = 400
	BindParamsError  = 401
	ValidParamsError = 402
	PermissionFailed = 403

	RequestError     = 10000
	AuthNotExist     = 10001
	AuthCheckTimeout = 10002
	AuthCheckFail    = 10003
	AuthFormatFail   = 10004

	RecordNotFound = 20000
)
