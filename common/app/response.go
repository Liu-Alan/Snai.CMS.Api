package app

import (
	"Snai.CMS.Api/common/message"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(result *message.Message) {
	data := gin.H{}
	if result != nil {
		data = gin.H{
			"code":   result.Code,
			"msg":    result.Msg,
			"result": result.Result,
		}
	}
	r.Ctx.JSON(message.Success, data)
}

func (r *Response) ToErrorResponse(err *message.Message) {
	response := gin.H{"code": err.Code, "msg": err.Msg, "result": err.Result}
	httpStatus := err.Code
	if err.Code > 500 {
		httpStatus = message.Error
	}

	r.Ctx.JSON(httpStatus, response)
}
