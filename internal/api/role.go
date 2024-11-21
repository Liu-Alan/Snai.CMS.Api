package api

import (
	"strconv"

	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/message"
	"Snai.CMS.Api/internal/entity"
	"Snai.CMS.Api/internal/model"
	"Snai.CMS.Api/internal/service"
	"github.com/gin-gonic/gin"
)

func RolesHandler(c *gin.Context) {
	response := app.NewResponse(c)

	pager := app.ResponsePage{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows := service.GetRoleCount()

	if totalRows <= 0 {
		page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: totalRows}
		response.ResponsePage(&page)
	} else {
		roles, msg := service.GetRoles(pager.Page, pager.PageSize)
		if msg.Code != message.Success {
			page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: 0}
			response.ResponsePage(&page)
		} else {
			var rolesOut []*model.RoleOut
			for _, role := range roles {
				roleOut := model.RoleOut{
					Key:   role.ID,
					ID:    role.ID,
					Title: role.Title,
					State: role.State,
				}

				rolesOut = append(rolesOut, &roleOut)
			}
			page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: totalRows, List: rolesOut}
			response.ResponsePage(&page)
		}
	}
}

func GetRoleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id <= 0 {
		response.ToErrorResponse(&message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)})
		return
	}

	role, msg := service.GetRoleByID(id)
	if msg.Code == message.Success {
		roleOut := model.RoleOut{
			ID:    role.ID,
			Title: role.Title,
			State: role.State,
		}
		msg.Result = roleOut
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func AddRoleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var roleIn model.AddRoleIn

	msg := app.BindAndValid(c, &roleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	role := entity.Roles{
		Title: roleIn.Title,
		State: roleIn.State,
	}
	msgM := service.AddRole(&role)
	if msgM.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "添加成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msgM)
	}
}

func UpdateRoleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var roleIn model.UpdateRoleIn

	msg := app.BindAndValid(c, &roleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	role := entity.Roles{
		ID:    roleIn.ID,
		Title: roleIn.Title,
		State: roleIn.State,
	}

	msg = service.ModifyRole(&role)
	if msg.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "修改成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func EnDisableRoleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var roleIn model.EnDisableRoleIn

	msg := app.BindAndValid(c, &roleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}
	ids := []int{roleIn.ID}
	msgM := service.UpdateRoleState(ids, roleIn.State)
	if msgM.Code == message.Success {
		msgC := "启用成功"
		if roleIn.State == 2 {
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

func BatchEnDisableRoleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var roleIn model.BatchEnDisableRoleIn

	msg := app.BindAndValid(c, &roleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}
	msgM := service.UpdateRoleState(roleIn.IDs, roleIn.State)
	if msgM.Code == message.Success {
		msgC := "启用成功"
		if roleIn.State == 2 {
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

func DeleteRoleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil || id <= 0 {
		response.ToErrorResponse(&message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)})
		return
	}

	msg := service.DeleteRole(id)
	if msg.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "删除成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func BatchDeleteRoleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var roleIn model.BatchDeleteRoleIn

	msg := app.BindAndValid(c, &roleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	msgM := service.BatchDeleteRole(roleIn.IDs)
	if msgM.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "删除成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func RoleModulesHandler(c *gin.Context) {
	response := app.NewResponse(c)
	roleID, err := strconv.Atoi(c.Query("role_id"))
	if err != nil || roleID <= 0 {
		response.ToErrorResponse(&message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)})
		return
	}

	roleModules, msg := service.GetRoleModules(roleID)
	if msg.Code == message.Success {
		var modules []int
		for _, module := range roleModules {
			modules = append(modules, module.ModuleID)
		}
		msg.Result = modules
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func AssignPermHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var permIn model.AssignPermIn

	msg := app.BindAndValid(c, &permIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	service.DeleteRoleModules(permIn.RoleID)
	msgM := service.AddRoleModules(permIn.RoleID, permIn.ModuleIDs)
	if msgM.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "分配权限成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}
