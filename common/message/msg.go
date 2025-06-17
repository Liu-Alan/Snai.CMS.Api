package message

// MsgInfo 定义code对应的msg
var MsgInfo = map[int]string{
	Success: "ok",
	Fail:    "fail",
	Error:   "Error",

	InvalidParams:    "参数错误",
	BindParamsError:  "绑定参数错误",
	ValidParamsError: "校验参数错误",

	PermissionFailed: "没有权限",

	RequestError:     "请求错误",
	AuthCheckFail:    "Token校验失败",
	AuthNotExist:     "缺少Token",
	AuthCheckTimeout: "Token超时",
	AuthFormatFail:   "Token格式错误",

	RecordNotFound: "记录不存在",
}

// GetMsg 根据code获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgInfo[code]
	if ok {
		return msg
	}

	return MsgInfo[Error]
}
