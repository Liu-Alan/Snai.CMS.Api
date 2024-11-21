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

func ModulesHandler(c *gin.Context) {
	response := app.NewResponse(c)

	pager := app.ResponsePage{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	totalRows := service.GetModuleCount()

	if totalRows <= 0 {
		page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: totalRows}
		response.ResponsePage(&page)
	} else {
		modules, msg := service.GetModules(pager.Page, pager.PageSize)
		if msg.Code != message.Success {
			page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: 0}
			response.ResponsePage(&page)
		} else {
			var modulesOut []*model.ModuleOut
			menus, _ := service.GetModuleMenus()
			for _, module := range modules {
				moduleOut := model.ModuleOut{
					Key:      module.ID,
					ID:       module.ID,
					ParentID: module.ParentID,
					Title:    module.Title,
					Name:     module.Name,
					UIRouter: module.UIRouter,
					Router:   module.Router,
					Menu:     module.Menu,
					Sort:     module.Sort,
					State:    module.State,
				}
				if module.ParentID > 0 {
					for _, menu := range menus {
						if module.ParentID == menu.ID {
							moduleOut.ParentTitle = menu.Title
							break
						}
					}
				}
				modulesOut = append(modulesOut, &moduleOut)
			}
			page := app.ResponsePage{Page: pager.Page, PageSize: pager.PageSize, Total: totalRows, List: modulesOut}
			response.ResponsePage(&page)
		}
	}
}

func GetModulesHandler(c *gin.Context) {
	response := app.NewResponse(c)

	modules, msg := service.GetModules(0, 0)
	if msg.Code == message.Success {
		var modulesOut []*model.ModuleOut
		for _, module := range modules {
			if module.State == 1 {
				moduleOut := model.ModuleOut{
					ID:       module.ID,
					ParentID: module.ParentID,
					Title:    module.Title,
					Name:     module.Name,
					Menu:     module.Menu,
				}
				modulesOut = append(modulesOut, &moduleOut)
			}
		}

		msg.Result = modulesOut
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func GetModuleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id <= 0 {
		response.ToErrorResponse(&message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)})
		return
	}

	module, msg := service.GetModuleByID(id)
	if msg.Code == message.Success {
		moduleOut := model.ModuleOut{
			ID:       module.ID,
			ParentID: module.ParentID,
			Title:    module.Title,
			Name:     module.Name,
			Router:   module.Router,
			UIRouter: module.UIRouter,
			Menu:     module.Menu,
			Sort:     module.Sort,
			State:    module.State,
		}
		msg.Result = moduleOut
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func AddModuleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var moduleIn model.AddModuleIn

	msg := app.BindAndValid(c, &moduleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	module := entity.Modules{
		ParentID: moduleIn.ParentID,
		Title:    moduleIn.Title,
		Name:     moduleIn.Name,
		Router:   moduleIn.Router,
		UIRouter: moduleIn.UIRouter,
		Menu:     moduleIn.Menu,
		Sort:     moduleIn.Sort,
		State:    moduleIn.State,
	}
	msgM := service.AddModule(&module)
	if msgM.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "添加成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msgM)
	}
}

func UpdateModuleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var moduleIn model.UpdateModuleIn

	msg := app.BindAndValid(c, &moduleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	module := entity.Modules{
		ID:       moduleIn.ID,
		ParentID: moduleIn.ParentID,
		Title:    moduleIn.Title,
		Name:     moduleIn.Name,
		Router:   moduleIn.Router,
		UIRouter: moduleIn.UIRouter,
		Menu:     moduleIn.Menu,
		Sort:     moduleIn.Sort,
		State:    moduleIn.State,
	}

	msg = service.ModifyModule(&module)
	if msg.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "修改成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func EnDisableModuleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var moduleIn model.EnDisableModuleIn

	msg := app.BindAndValid(c, &moduleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}
	ids := []int{moduleIn.ID}
	msgM := service.UpdateModuleState(ids, moduleIn.State)
	if msgM.Code == message.Success {
		msgC := "启用成功"
		if moduleIn.State == 2 {
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

func BatchEnDisableModuleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var moduleIn model.BatchEnDisableModuleIn

	msg := app.BindAndValid(c, &moduleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}
	msgM := service.UpdateModuleState(moduleIn.IDs, moduleIn.State)
	if msgM.Code == message.Success {
		msgC := "启用成功"
		if moduleIn.State == 2 {
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

func DeleteModuleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil || id <= 0 {
		response.ToErrorResponse(&message.Message{Code: message.BindParamsError, Msg: message.GetMsg(message.BindParamsError)})
		return
	}

	msg := service.DeleteModule(id)
	if msg.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "删除成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}

func BatchDeleteModuleHandler(c *gin.Context) {
	response := app.NewResponse(c)
	var moduleIn model.BatchDeleteModuleIn

	msg := app.BindAndValid(c, &moduleIn, "form")
	if msg.Code != message.Success {
		response.ToErrorResponse(msg)
		return
	}

	msgM := service.BatchDeleteModule(moduleIn.IDs)
	if msgM.Code == message.Success {
		msg.Code = message.Success
		msg.Msg = "删除成功"
		msg.Result = nil
		response.ToResponse(msg)
	} else {
		response.ToErrorResponse(msg)
	}
}
