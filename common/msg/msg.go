package msg

// MsgInfo 定义code对应的msg
var MsgInfo = map[int]string{
	Success:          "ok",
	Error:            "fail",
	InvalidParams:    "请求参数错误",
	BindParamsError:  "绑定参数错误",
	ValidParamsError: "校验参数错误",
	RequestError:     "请求错误",
	AuthNotExist:     "缺少头部参数Authorization",
	AuthCheckTimeout: "Authorization超时",
	AuthCheckFail:    "校验Authorization失败",
	AuthFormatFail:   "Authorization格式错误",
}

// GetMsg 根据code获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgInfo[code]
	if ok {
		return msg
	}

	return MsgInfo[Error]
}
