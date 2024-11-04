package api

import (
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
	var adminsIn model.AdminsIn

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
			var adminsOut []*model.AdminsOut
			roles, _ := service.GetRoles(0, 0)
			nowtime := int(time.Now().Unix())
			for _, admin := range admins {
				adminOut := model.AdminsOut{
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

func AddAdminHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var adminIn model.AddAdminIn

	msg := app.BindAndValid(c, &adminIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	adminIn.Password = strings.ToLower(utils.EncodeMD5(config.AppConf.PwdSalt + strings.TrimSpace(adminIn.Password)))
	admin := entity.Admins{
		UserName:   adminIn.UserName,
		Password:   adminIn.Password,
		RoleID:     adminIn.RoleID,
		State:      adminIn.State,
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
