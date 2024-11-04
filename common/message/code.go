package message

const (
	Success = 200
	Fail    = 100
	Error   = 500

	InvalidParams    = 406
	BindParamsError  = 407
	ValidParamsError = 408

	PermissionFailed = 403

	RequestError     = 400
	AuthCheckFail    = 401
	AuthNotExist     = 402
	AuthCheckTimeout = 404
	AuthFormatFail   = 405

	RecordNotFound = 600
)
