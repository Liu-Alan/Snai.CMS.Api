package app

import (
	"Snai.CMS.Api/common/msg"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(result *msg.Message) {
	data := gin.H{}
	if result != nil {
		data = gin.H{
			"code":   result.Code,
			"msg":    result.Msg,
			"result": result.Result,
		}
	}
	r.Ctx.JSON(msg.Success, data)
}

func (r *Response) ToErrorResponse(err *msg.Message) {
	response := gin.H{"code": err.Code, "msg": err.Msg}
	httpStatus := err.Code
	if err.Code > 500 {
		httpStatus = msg.Error
	}

	r.Ctx.JSON(httpStatus, response)
}
