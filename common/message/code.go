package message

const (
	Success = 200
	Fail    = 100
	Error   = 500

	InvalidParams    = 501
	BindParamsError  = 502
	ValidParamsError = 503

	PermissionFailed = 403

	RequestError     = 400
	AuthCheckFail    = 401
	AuthNotExist     = 402
	AuthCheckTimeout = 404
	AuthFormatFail   = 405

	RecordNotFound = 600
)
