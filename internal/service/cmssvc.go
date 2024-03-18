package service

import (
	"strings"

	"Snai.CMS.Api/common/msg"
	"Snai.CMS.Api/internal/dao"
	"Snai.CMS.Api/internal/entity"
)

func GetAdmin(userName string) (*entity.Admins, msg.Message) {
	err := msg.Message{Code: msg.Success, Msg: msg.GetMsg(msg.Success)}

	if strings.TrimSpace(userName) == "" {
		return nil, msg.Message{Code: msg.InvalidParams, Msg: msg.GetMsg(msg.InvalidParams)}
	}
	admin, _ := dao.GetAdmin(userName)
	if admin == nil || admin.ID <= 0 {
		return nil, msg.Message{Code: msg.RecordNotFound, Msg: msg.GetMsg(msg.RecordNotFound)}
	}
	return admin, err
}

func GetToken(token string) (*entity.Tokens, msg.Message) {
	err := msg.Message{Code: msg.Success, Msg: msg.GetMsg(msg.Success)}

	if strings.TrimSpace(token) == "" {
		return nil, msg.Message{Code: msg.InvalidParams, Msg: msg.GetMsg(msg.InvalidParams)}
	}
	tk, _ := dao.GetToken(token)
	if tk == nil || tk.ID <= 0 {
		return nil, msg.Message{Code: msg.RecordNotFound, Msg: msg.GetMsg(msg.RecordNotFound)}
	}
	return tk, err
}

func GetRole(roleID int) (*entity.Roles, msg.Message) {
	err := msg.Message{Code: msg.Success, Msg: msg.GetMsg(msg.Success)}

	if roleID <= 0 {
		return nil, msg.Message{Code: msg.InvalidParams, Msg: msg.GetMsg(msg.InvalidParams)}
	}
	role, _ := dao.GetRole(roleID)
	if role == nil || role.ID <= 0 {
		return nil, msg.Message{Code: msg.RecordNotFound, Msg: msg.GetMsg(msg.RecordNotFound)}
	}
	return role, err
}

func GetModule(router string) (*entity.Modules, msg.Message) {
	err := msg.Message{Code: msg.Success, Msg: msg.GetMsg(msg.Success)}

	if strings.TrimSpace(router) == "" {
		return nil, msg.Message{Code: msg.InvalidParams, Msg: msg.GetMsg(msg.InvalidParams)}
	}
	module, _ := dao.GetModule(router)
	if module == nil || module.ID <= 0 {
		return nil, msg.Message{Code: msg.RecordNotFound, Msg: msg.GetMsg(msg.RecordNotFound)}
	}
	return module, err
}

func GetRoleModule(roleID int, moduleID int) (*entity.RoleModule, msg.Message) {
	err := msg.Message{Code: msg.Success, Msg: msg.GetMsg(msg.Success)}

	if roleID <= 0 || moduleID <= 0 {
		return nil, msg.Message{Code: msg.InvalidParams, Msg: msg.GetMsg(msg.InvalidParams)}
	}
	roleModule, _ := dao.GetRoleModule(roleID, moduleID)
	if roleModule == nil || roleModule.ID <= 0 {
		return nil, msg.Message{Code: msg.RecordNotFound, Msg: msg.GetMsg(msg.RecordNotFound)}
	}
	return roleModule, err
}

// 判断权限
func VerifyUserRole(userName string, router string) msg.Message {
	err := msg.Message{Code: msg.Success, Msg: msg.GetMsg(msg.Success)}

	admin, err := GetAdmin(userName)
	if admin == nil {
		return err
	}
	if admin.State == 2 {
		return msg.Message{Code: msg.RecordNotFound, Msg: msg.GetMsg(msg.RecordNotFound)}
	}

	role, err := GetRole(admin.RoleID)
	if role == nil {
		return err
	}
	if role.State == 2 {
		return msg.Message{Code: msg.RecordNotFound, Msg: msg.GetMsg(msg.RecordNotFound)}
	}

	module, err := GetModule(router)
	if module == nil {
		return err
	}
	if module.State == 2 {
		return msg.Message{Code: msg.RecordNotFound, Msg: msg.GetMsg(msg.RecordNotFound)}
	}

	// 为-1不验证权限
	if module.ParentID == -1 {
		return err
	}

	roleModule, err := GetRoleModule(role.ID, module.ID)
	if roleModule == nil {
		return msg.Message{Code: msg.PermissionFailed, Msg: msg.GetMsg(msg.PermissionFailed)}
	}

	return err
}
