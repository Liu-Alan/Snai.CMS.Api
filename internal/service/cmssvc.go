package service

import (
	"fmt"
	"strings"
	"time"

	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/config"
	"Snai.CMS.Api/common/logging"
	"Snai.CMS.Api/common/message"
	"Snai.CMS.Api/common/utils"
	"Snai.CMS.Api/internal/dao"
	"Snai.CMS.Api/internal/entity"
	"Snai.CMS.Api/internal/model"
)

func GetAdmin(userName string) (*entity.Admins, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if strings.TrimSpace(userName) == "" {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	admin, _ := dao.GetAdmin(userName)
	if admin == nil || admin.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return admin, &err
}

func GetAdminByID(id int) (*entity.Admins, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if id <= 0 {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	admin, _ := dao.GetAdminByID(id)
	if admin == nil || admin.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return admin, &err
}

func GetAdminCount(userName string) int64 {
	count, _ := dao.GetAdminCount(userName)
	return count
}

func GetAdmins(userName string, page, pageSize int) ([]*entity.Admins, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	pageOffset := app.GetPageOffset(page, pageSize)
	admins, _ := dao.GetAdmins(userName, pageOffset, pageSize)
	if admins == nil || len(admins) <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	for k := range admins {
		admins[k].Password = ""
	}
	return admins, &err
}

func AddAdmin(admin *entity.Admins) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if admin != nil && strings.TrimSpace(admin.UserName) != "" {
		_, errA := GetAdmin(admin.UserName)
		if errA.Code == message.Success {
			err.Code = message.InvalidParams
			err.Msg = "当前用户名已存在"
			return &err
		}
		if errm := dao.AddAdmin(admin); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "保存Admin失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "保存Admin失败"
		return &err
	}
}

func ModifyAdmin(admin *entity.Admins) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if admin != nil && admin.ID > 0 {
		reAdmin, _ := GetAdmin(admin.UserName)
		if reAdmin != nil && admin.ID != reAdmin.ID {
			err.Code = message.InvalidParams
			err.Msg = "用户名重复"
			return &err
		} else {
			if errm := dao.ModifyAdmin(admin); errm != nil {
				logging.Error(errm.Error())
				err.Code = message.InvalidParams
				err.Msg = "用户信息修改失败"
				return &err
			}

			return &err
		}
	} else {
		err.Code = message.InvalidParams
		err.Msg = "用户不存在"
		return &err
	}
}

func UpdateAdminState(ids []int, state int8) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}
	updateTime := int(time.Now().Unix())
	if len(ids) > 0 {
		if errm := dao.UpdateAdminState(ids, state, updateTime); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "用户更新失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "没有选择任何列"
		return &err
	}
}

func UnlockAdmin(id int) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if id > 0 {
		if errm := dao.UnlockAdmin(id); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "用户解锁失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "用户不存在"
		return &err
	}
}

func DeleteAdmin(id int) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if id > 0 {
		admin := entity.Admins{
			ID: id,
		}

		if errm := dao.DeleteAdmin(&admin); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "用户删除失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "用户不存在"
		return &err
	}
}

func BatchDeleteAdmin(ids []int) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if len(ids) > 0 {
		if errm := dao.BatchDeleteAdmin(ids); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "用户删除失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "没有选择任何列"
		return &err
	}
}

func AdminLogin(loginIn *model.LoginIn, ip string) (*entity.Admins, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}
	if strings.TrimSpace(loginIn.UserName) == "" || strings.TrimSpace(loginIn.Password) == "" {
		err.Code = message.InvalidParams
		err.Msg = "用户名或密码不能为空"
		return nil, &err
	}
	if len(loginIn.UserName) > 32 {
		err.Code = message.InvalidParams
		err.Msg = "用户名或密码不能为空"
		return nil, &err
	}

	admin, _ := GetAdmin(loginIn.UserName)
	if admin == nil {
		err.Code = message.InvalidParams
		err.Msg = "用户名或密码不能为空"
		return nil, &err
	}

	if admin.State == 2 {
		err.Code = message.Error
		err.Msg = "用户名或密码不能为空"
		return nil, &err
	}

	nowtime := int(time.Now().Unix())
	if admin.LockTime > nowtime {
		err.Code = message.Error
		err.Msg = fmt.Sprintf("帐号已锁定，请%d分钟后再来登录", config.AppConf.LoginLockMinute)
		return nil, &err
	}

	role, _ := GetRoleByID(admin.RoleID)
	if role == nil || role.ID <= 0 || role.State == 2 {
		err.Code = message.Error
		err.Msg = "用户角色禁用"
		return nil, &err
	}

	pwd := strings.ToLower(utils.EncodeMD5(config.AppConf.PwdSalt + strings.TrimSpace(loginIn.Password)))
	if admin.Password != pwd {
		if admin.ErrorLogonTime+(config.AppConf.LoginLockMinute*60) < nowtime {
			admin.ErrorLogonTime = nowtime
			admin.ErrorLogonCount = 1
		} else {
			admin.ErrorLogonCount += 1
		}

		if admin.ErrorLogonCount >= config.AppConf.LoginErrorCount {
			admin.ErrorLogonTime = 0
			admin.ErrorLogonCount = 0
			admin.LockTime = nowtime + (config.AppConf.LoginLockMinute * 60)
			admin.UpdateTime = nowtime
			//锁定账号
			ModifyAdmin(admin)

			err.Code = message.Error
			err.Msg = fmt.Sprintf("帐号或密码在%d分钟内，错误%d次，锁定帐号%d分钟", config.AppConf.LoginLockMinute, config.AppConf.LoginErrorCount, config.AppConf.LoginLockMinute)
			return nil, &err
		} else {
			//更新错误登录信息
			ModifyAdmin(admin)

			err.Code = message.Error
			err.Msg = fmt.Sprintf("帐号或密码错误，如在%d分钟内，错误%d次，将锁定帐号%d分钟", config.AppConf.LoginLockMinute, config.AppConf.LoginErrorCount, config.AppConf.LoginLockMinute)
			return nil, &err
		}
	}
	admin.LastLogonTime = nowtime
	admin.ErrorLogonTime = 0
	admin.ErrorLogonCount = 0
	admin.LockTime = 0
	admin.LastLogonIP = ip
	//更新账号登录信息
	ModifyAdmin(admin)

	return admin, &err
}

func GetToken(token string) (*entity.Tokens, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if strings.TrimSpace(token) == "" {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	tk, _ := dao.GetToken(token)
	if tk == nil || tk.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return tk, &err
}

func AddToken(token *entity.Tokens) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if token != nil && token.UserID > 0 {
		if errm := dao.AddToken(token); errm != nil {
			err.Code = message.InvalidParams
			err.Msg = "保存Token失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "保存Token失败"
		return &err
	}
}

func ModifyToken(token *entity.Tokens) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if token != nil && token.ID > 0 {
		if errm := dao.ModifyToken(token); errm != nil {
			err.Code = message.InvalidParams
			err.Msg = "Tokek修改失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "Tokek不存在"
		return &err
	}
}

func GetModule(router string) (*entity.Modules, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if strings.TrimSpace(router) == "" {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	module, _ := dao.GetModule(router)
	if module == nil || module.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return module, &err
}

func GetModuleByID(id int) (*entity.Modules, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if id <= 0 {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	module, _ := dao.GetModuleByID(id)
	if module == nil || module.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return module, &err
}

func GetModuleByTitle(parentID int, title string) (*entity.Modules, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if parentID < 0 || strings.TrimSpace(title) == "" {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	module, _ := dao.GetModuleByTitle(parentID, title)
	if module == nil || module.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return module, &err
}

func GetModulesByIDs(ids []int) ([]*entity.Modules, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if ids == nil || len(ids) <= 0 {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	modules, _ := dao.GetModulesByIDs(ids)
	if modules == nil || len(modules) <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return modules, &err
}

func GetModules(page, pageSize int) ([]*entity.Modules, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	pageOffset := app.GetPageOffset(page, pageSize)
	modules, _ := dao.GetModules(pageOffset, pageSize)
	if modules == nil || len(modules) <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return modules, &err
}

func GetModuleMenus() ([]*entity.Modules, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	modules, _ := dao.GetModuleMenus()
	if modules == nil || len(modules) <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return modules, &err
}

func GetModuleCount() int64 {
	count, _ := dao.GetModuleCount()
	return count
}

func AddModule(module *entity.Modules) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if module != nil && strings.TrimSpace(module.Title) != "" {
		_, errA := GetModuleByTitle(module.ParentID, module.Title)
		if errA.Code == message.Success {
			err.Code = message.InvalidParams
			err.Msg = "当前模块名已存在"
			return &err
		}
		if errm := dao.AddModule(module); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "保存Module失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "保存Module失败"
		return &err
	}
}

func ModifyModule(module *entity.Modules) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if module != nil && module.ID > 0 {
		reModule, _ := GetModuleByTitle(module.ParentID, module.Title)
		if reModule != nil && module.ID != reModule.ID {
			err.Code = message.InvalidParams
			err.Msg = "模块名重复"
			return &err
		} else {
			if errm := dao.ModifyModule(module); errm != nil {
				logging.Error(errm.Error())
				err.Code = message.InvalidParams
				err.Msg = "模块修改失败"
				return &err
			}

			return &err
		}
	} else {
		err.Code = message.InvalidParams
		err.Msg = "模块不存在"
		return &err
	}
}

func UpdateModuleState(ids []int, state int8) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if len(ids) > 0 {
		if errm := dao.UpdateModuleState(ids, state); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "模块更新失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "没有选择任何列"
		return &err
	}
}

func DeleteModule(id int) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if id > 0 {
		module := entity.Modules{
			ID: id,
		}

		if errm := dao.DeleteModule(&module); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "模块删除失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "模块不存在"
		return &err
	}
}

func BatchDeleteModule(ids []int) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if len(ids) > 0 {
		if errm := dao.BatchDeleteModule(ids); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "模块删除失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "没有选择任何列"
		return &err
	}
}

func GetRoleByID(roleID int) (*entity.Roles, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if roleID <= 0 {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	role, _ := dao.GetRoleByID(roleID)
	if role == nil || role.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return role, &err
}

func GetRoleByTitle(title string) (*entity.Roles, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if strings.TrimSpace(title) == "" {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	role, _ := dao.GetRoleByTitle(title)
	if role == nil || role.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return role, &err
}

func GetRoles(page, pageSize int) ([]*entity.Roles, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	pageOffset := app.GetPageOffset(page, pageSize)
	roles, _ := dao.GetRoles(pageOffset, pageSize)
	if roles == nil || len(roles) <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return roles, &err
}

func GetRoleCount() int64 {
	count, _ := dao.GetRoleCount()
	return count
}

func AddRole(role *entity.Roles) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if role != nil && strings.TrimSpace(role.Title) != "" {
		_, errA := GetRoleByTitle(role.Title)
		if errA.Code == message.Success {
			err.Code = message.InvalidParams
			err.Msg = "当前角色名已存在"
			return &err
		}
		if errm := dao.AddRole(role); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "保存Role失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "保存Role失败"
		return &err
	}
}

func ModifyRole(role *entity.Roles) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if role != nil && role.ID > 0 {
		reRole, _ := GetRoleByTitle(role.Title)
		if reRole != nil && role.ID != reRole.ID {
			err.Code = message.InvalidParams
			err.Msg = "角色名重复"
			return &err
		} else {
			if errm := dao.ModifyRole(role); errm != nil {
				logging.Error(errm.Error())
				err.Code = message.InvalidParams
				err.Msg = "角色修改失败"
				return &err
			}

			return &err
		}
	} else {
		err.Code = message.InvalidParams
		err.Msg = "角色不存在"
		return &err
	}
}

func UpdateRoleState(ids []int, state int8) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if len(ids) > 0 {
		if errm := dao.UpdateRoleState(ids, state); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "角色更新失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "没有选择任何列"
		return &err
	}
}

func DeleteRole(id int) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if id > 0 {
		role := entity.Roles{
			ID: id,
		}

		if errm := dao.DeleteRole(&role); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "角色删除失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "角色不存在"
		return &err
	}
}

func BatchDeleteRole(ids []int) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if len(ids) > 0 {
		if errm := dao.BatchDeleteRole(ids); errm != nil {
			logging.Error(errm.Error())
			err.Code = message.InvalidParams
			err.Msg = "角色删除失败"
			return &err
		}

		return &err
	} else {
		err.Code = message.InvalidParams
		err.Msg = "没有选择任何列"
		return &err
	}
}

func GetRoleModule(roleID int, moduleID int) (*entity.RoleModule, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if roleID <= 0 || moduleID <= 0 {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	roleModule, _ := dao.GetRoleModule(roleID, moduleID)
	if roleModule == nil || roleModule.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return roleModule, &err
}

func GetRoleModules(roleID int) ([]*entity.RoleModule, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	roleModules, _ := dao.GetRoleModules(roleID)
	if roleModules == nil || len(roleModules) <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}

	return roleModules, &err
}

// 判断权限
func VerifyUserRole(userName string, router string) *message.Message {
	err := &message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	admin, err := GetAdmin(userName)
	if err.Code == message.RecordNotFound {
		return &message.Message{Code: message.RecordNotFound, Msg: "账号不存在"}
	}
	if admin == nil {
		return err
	}
	if admin.State == 2 {
		return &message.Message{Code: message.RecordNotFound, Msg: "账号已禁用"}
	}

	role, err := GetRoleByID(admin.RoleID)
	if err.Code == message.RecordNotFound {
		return &message.Message{Code: message.RecordNotFound, Msg: "角色不存在"}
	}
	if role == nil {
		return err
	}
	if role.State == 2 {
		return &message.Message{Code: message.RecordNotFound, Msg: "角色已禁用"}
	}

	module, err := GetModule(router)
	if err.Code == message.RecordNotFound {
		return &message.Message{Code: message.RecordNotFound, Msg: "模块不存在"}
	}
	if module == nil {
		return err
	}
	if module.State == 2 {
		return &message.Message{Code: message.RecordNotFound, Msg: "模块已禁用"}
	}

	// 为-1不验证权限
	if module.ParentID == -1 {
		return err
	}

	roleModule, err := GetRoleModule(role.ID, module.ID)
	if roleModule == nil {
		return &message.Message{Code: message.PermissionFailed, Msg: message.GetMsg(message.PermissionFailed)}
	}

	return err
}
