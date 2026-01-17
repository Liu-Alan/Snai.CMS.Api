package api

import (
	"strconv"
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
	otpCode, errO := strconv.Atoi(loginIn.OtpCode)
	if errO != nil {
		msg.Code = message.ValidParamsError
		msg.Msg = message.GetMsg(message.ValidParamsError)
		msg.Result = "Otp动态码为6位数字"
		response.ToErrorResponse(msg)
		return
	}

	admin, err := service.AdminLogin(&loginIn, ip)
	if err.Code == message.Success {
		otpVerify := app.OtpVerifyCode(admin.OtpSecret, int32(otpCode))
		if !otpVerify {
			msg.Code = message.Error
			msg.Msg = "Otp动态码错误"
			response.ToErrorResponse(msg)
			return
		}

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
				msg.Result = model.LoginOut{Token: token, UserName: admin.UserName}
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
	token, _ := utils.GetGinContextByKey(c, "token")
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
	user_name, _ := utils.GetGinContextByKey(c, "user_name")

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
			msg.Code = message.InvalidParams
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

func MenuHandler(c *gin.Context) {
	response := app.NewResponse(c)
	user_name, _ := utils.GetGinContextByKey(c, "user_name")

	admin, msg := service.GetAdmin(user_name)
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
	} else {
		roleModules, msg := service.GetRoleModules(admin.RoleID)
		if msg.Code != message.Success {
			response.ToErrorResponse(msg)
		} else {
			var ids []int
			for _, v := range roleModules {
				ids = append(ids, v.ModuleID)
			}
			modules, msg := service.GetModulesByIDs(ids)
			if msg.Code != message.Success {
				response.ToErrorResponse(msg)
			} else {
				var menus []*model.MenuOut
				for _, module := range modules {
					if module.State == 1 {
						menu := model.MenuOut{
							ID:       module.ID,
							ParentID: module.ParentID,
							Title:    module.Title,
							Name:     module.Name,
							Router:   module.Router,
							UIRouter: module.UIRouter,
							Menu:     module.Menu,
							Sort:     module.Sort,
						}
						menus = append(menus, &menu)
					}
				}
				msg.Result = menus
				response.ToResponse(msg)
			}
		}
	}
}

func RoleHandler(c *gin.Context) {
	response := app.NewResponse(c)

	roles, msg := service.GetRoles(0, 0)
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
	} else {

		var rolesOut []*model.RoleOut
		for _, role := range roles {
			if role.State == 1 {
				roleOut := model.RoleOut{
					ID:    role.ID,
					Title: role.Title,
				}
				rolesOut = append(rolesOut, &roleOut)
			}
		}
		msg.Result = rolesOut
		response.ToResponse(msg)
	}

}
