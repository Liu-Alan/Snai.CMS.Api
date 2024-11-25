package api

import (
	"strconv"
	"strings"
	"time"

	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/config"
	"Snai.CMS.Api/common/message"
	"Snai.CMS.Api/common/utils"
	"Snai.CMS.Api/internal/entity"
	"Snai.CMS.Api/internal/model"
	"Snai.CMS.Api/internal/service"
	"github.com/gin-gonic/gin"
)

func AdminsHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var adminsIn model.AdminIn

	msg := app.BindAndValid(c, &adminsIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	pager := app.ResponsePage{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows := service.GetAdminCount(adminsIn.UserName)

	if totalRows <= 0 {
		page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: totalRows}
		response.ResponsePage(&page)
	} else {
		admins, msg := service.GetAdmins(adminsIn.UserName, pager.Page, pager.PageSize)
		if msg.Code != message.Success {
			page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: 0}
			response.ResponsePage(&page)
		} else {
			var adminsOut []*model.AdminOut
			roles, _ := service.GetRoles(0, 0)
			nowtime := int(time.Now().Unix())
			for _, admin := range admins {
				adminOut := model.AdminOut{
					Key:      admin.ID,
					ID:       admin.ID,
					UserName: admin.UserName,
					State:    admin.State,
				}
				for _, role := range roles {
					if role.ID == admin.RoleID {
						adminOut.Role = role.Title
						break
					}
				}
				if admin.LockTime > nowtime {
					adminOut.LockState = 2
				} else {
					adminOut.LockState = 1
				}
				adminsOut = append(adminsOut, &adminOut)
			}
			page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: totalRows, List: adminsOut}
			response.ResponsePage(&page)
		}
	}
}

func GetAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id <= 0 {
		response.ToErrorResponse(&message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)})
		return
	}

	admin, msg := service.GetAdminByID(id)
	if msg.Code == message.Success {
		adminOut := model.AdminOut{
			ID:       admin.ID,
			UserName: admin.UserName,
			RoleID:   admin.RoleID,
			State:    admin.State,
		}
		msg.Result = adminOut
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func AddAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var adminIn model.AddAdminIn

	msg := app.BindAndValid(c, &adminIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	adminIn.Password = strings.ToLower(utils.EncodeMD5(config.AppConf.PwdSalt + strings.TrimSpace(adminIn.Password)))
	optSecret := app.OtpSecret()
	admin := entity.Admins{
		UserName:   adminIn.UserName,
		Password:   adminIn.Password,
		RoleID:     adminIn.RoleID,
		State:      adminIn.State,
		OtpSecret:  optSecret,
		CreateTime: int(time.Now().Unix()),
	}
	msgM := service.AddAdmin(&admin)
	if msgM.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "添加成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msgM)
	}
}

func UpdateAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var adminIn model.UpdateAdminIn

	msg := app.BindAndValid(c, &adminIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	admin, msg := service.GetAdminByID(adminIn.ID)
	if msg.Code == message.Success {
		if adminIn.Password != "******" && strings.TrimSpace(adminIn.Password) != "" {
			admin.Password = strings.ToLower(utils.EncodeMD5(config.AppConf.PwdSalt + strings.TrimSpace(adminIn.Password)))
		}
		admin.UserName = adminIn.UserName
		admin.RoleID = adminIn.RoleID
		admin.State = adminIn.State
		admin.UpdateTime = int(time.Now().Unix())

		msg := service.ModifyAdmin(admin)
		if msg.Code == message.Success {
			msg.Code = message.Success
			msg.Msg = "修改成功"
			msg.Result = nil
			response.ToResponse(msg)
		} else {
			response.ToErrorResponse(msg)
		}
	} else {
		response.ToErrorResponse(msg)
	}
}

func EnDisableAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var adminIn model.EnDisableAdminIn

	msg := app.BindAndValid(c, &adminIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}
	ids := []int{adminIn.ID}
	msgM := service.UpdateAdminState(ids, adminIn.State)
	if msgM.Code == message.Success {
		msgC := "启用成功"
		if adminIn.State == 2 {
			msgC = "禁用成功"
		}
		msg.Code = message.Success
		msg.Msg = msgC
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func BatchEnDisableAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var adminIn model.BatchEnDisableAdminIn

	msg := app.BindAndValid(c, &adminIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}
	msgM := service.UpdateAdminState(adminIn.IDs, adminIn.State)
	if msgM.Code == message.Success {
		msgC := "启用成功"
		if adminIn.State == 2 {
			msgC = "禁用成功"
		}
		msg.Code = message.Success
		msg.Msg = msgC
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func UnlockAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil || id <= 0 {
		response.ToErrorResponse(&message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)})
		return
	}

	msg := service.UnlockAdmin(id)
	if msg.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "解锁成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func DeleteAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil || id <= 0 {
		response.ToErrorResponse(&message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)})
		return
	}

	msg := service.DeleteAdmin(id)
	if msg.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "删除成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func BatchDeleteAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var adminIn model.BatchDeleteAdminIn

	msg := app.BindAndValid(c, &adminIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	msgM := service.BatchDeleteAdmin(adminIn.IDs)
	if msgM.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "删除成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func GetAdminQrcodeHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id <= 0 {
		c.Data(message.BindParamsError, "image/png", nil)
		return
	}

	admin, msg := service.GetAdminByID(id)
	if msg.Code == message.Success {
		if strings.TrimSpace(admin.OtpSecret) != "" {
			otpQr := app.OtpQrcode(config.AppConf.JwtIssuer, admin.UserName, admin.OtpSecret)
			otpBy := app.QrcodeEncode(otpQr, 200)
			if otpBy != nil {
				c.Data(message.Success, "image/png", otpBy)
			} else {
				c.Data(message.Error, "image/png", nil)
			}
		} else {
			c.Data(message.Error, "image/png", nil)
		}
	} else {
		c.Data(message.Error, "image/png", nil)
	}
}
