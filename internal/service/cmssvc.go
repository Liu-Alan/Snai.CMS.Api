package service

import (
	"fmt"
	"strings"
	"time"

	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/config"
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

func GetAdminCount(userName string) int64 {
	count, _ := dao.GetAdminCount(userName)
	return count
}

func GetAdminList(userName string, page, pageSize int) ([]*entity.Admins, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	pageOffset := app.GetPageOffset(page, pageSize)
	admins, _ := dao.GetAdminList(userName, pageOffset, pageSize)
	if admins == nil || len(admins) <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	for k := range admins {
		admins[k].Password = ""
	}
	return admins, &err
}

func ModifyAdmin(admin *entity.Admins) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if admin != nil && admin.ID > 0 {
		if errm := dao.ModifyAdmin(admin); errm != nil {
			err.Code = 400
			err.Msg = "用户信息修改失败"
			return &err
		}

		return &err
	} else {
		err.Code = 400
		err.Msg = "用户不存在"
		return &err
	}
}

func AdminLogin(loginIn *model.LoginIn, ip string) (*entity.Admins, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}
	if strings.TrimSpace(loginIn.UserName) == "" || strings.TrimSpace(loginIn.Password) == "" {
		err.Code = 400
		err.Msg = "用户名或密码不能为空"
		return nil, &err
	}
	if len(loginIn.UserName) > 32 {
		err.Code = 400
		err.Msg = "用户名或密码不能为空"
		return nil, &err
	}

	admin, _ := GetAdmin(loginIn.UserName)
	if admin == nil {
		err.Code = 400
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

	role, _ := GetRole(admin.RoleID)
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
			err.Code = 400
			err.Msg = "保存Token失败"
			return &err
		}

		return &err
	} else {
		err.Code = 400
		err.Msg = "保存Token失败"
		return &err
	}
}

func ModifyToken(token *entity.Tokens) *message.Message {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if token != nil && token.ID > 0 {
		if errm := dao.ModifyToken(token); errm != nil {
			err.Code = 400
			err.Msg = "Tokek修改失败"
			return &err
		}

		return &err
	} else {
		err.Code = 400
		err.Msg = "Tokek不存在"
		return &err
	}
}

func GetRole(roleID int) (*entity.Roles, *message.Message) {
	err := message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}

	if roleID <= 0 {
		return nil, &message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
	}
	role, _ := dao.GetRole(roleID)
	if role == nil || role.ID <= 0 {
		return nil, &message.Message{Code: message.RecordNotFound, Msg: message.GetMsg(message.RecordNotFound)}
	}
	return role, &err
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

	role, err := GetRole(admin.RoleID)
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
