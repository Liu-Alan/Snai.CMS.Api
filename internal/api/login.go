package api

import (
	"strings"
	"time"

	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/config"
	"Snai.CMS.Api/common/logging"
	"Snai.CMS.Api/common/message"
	"Snai.CMS.Api/common/utils"
	"Snai.CMS.Api/internal/entity"
	"Snai.CMS.Api/internal/model"
	"Snai.CMS.Api/internal/service"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var loginIn model.LoginIn

	msg := app.BindAndValid(c, &loginIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	ip := c.ClientIP()
	admin, err := service.AdminLogin(&loginIn, ip)
	if err.Code == message.Success {
		token, err := app.GenerateToken(loginIn.UserName)
		if err != nil {
			logging.Error("app.GenerateToken err: %v", err)
			msg.Code = message.Error
			msg.Msg = "登录失败Token"
			response.ToErrorResponse(msg)
			return
		} else {
			tk := entity.Tokens{
				Token:      token,
				UserID:     admin.ID,
				State:      1,
				CreateTime: int(time.Now().Unix()),
			}
			errt := service.AddToken(&tk)
			if errt.Code != message.Success {
				response.ToErrorResponse(errt)
			} else {
				msg.Code = message.Success
				msg.Msg = message.GetMsg(message.Success)
				msg.Result = model.LoginOut{Token: token}
				response.ToResponse(msg)
			}
		}
	} else {
		response.ToErrorResponse(err)
		return
	}
}

func LogoutHandler(c *gin.Context) {
	response := app.NewResponse(c)
	token := c.MustGet("token").(string)
	tk, msg := service.GetToken(token)
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
	} else {
		tk.State = 2
		service.ModifyToken(tk)

		msg.Code = message.Success
		msg.Msg = "已退出"
		msg.Result = nil
		response.ToResponse(msg)
	}
}

func ChangePasswordHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var passwordIn model.ChangePasswordIn
	user_name := c.MustGet("user_name").(string)

	msg := app.BindAndValid(c, &passwordIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	admin, msg := service.GetAdmin(user_name)
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
	} else {
		pwd := strings.ToLower(utils.EncodeMD5(config.AppConf.PwdSalt + strings.TrimSpace(passwordIn.OldPassword)))
		if admin.Password != pwd {
			msg.Code = message.Error
			msg.Msg = "原密码错误"
			response.ToErrorResponse(msg)
			return
		}

		admin.Password = strings.ToLower(utils.EncodeMD5(config.AppConf.PwdSalt + strings.TrimSpace(passwordIn.Password)))
		msgM := service.ModifyAdmin(admin)
		if msgM.Code == message.Success {
			msg.Code = message.Success
			msg.Msg = "修改成功"
			msg.Result = nil
			response.ToResponse(msg)
		} else {
			response.ToErrorResponse(msgM)
		}
	}
}
