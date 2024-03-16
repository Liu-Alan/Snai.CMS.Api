package dao

import (
	"Snai.CMS.Api/internal/entity"
	"gorm.io/gorm"
)

func GetAdmin(userName string) (*entity.Admins, error) {
	tx := _cmsdb.Where("user_name = ? ", userName)
	var admin entity.Admins
	err := tx.First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &admin, nil
}

func GetRole(roleID int) (*entity.Roles, error) {
	tx := _cmsdb.Where("id = ? ", roleID)
	var role entity.Roles
	err := tx.First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &role, nil
}

func GetModule(router string) (*entity.Modules, error) {
	tx := _cmsdb.Where("router = ? ", router)
	var module entity.Modules
	err := tx.First(&module).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &module, nil
}

func GetRoleModule(roleID int, moduleID int) (*entity.RoleModule, error) {
	tx := _cmsdb.Where("role_id = ? and module_id = ?", roleID, moduleID)
	var roleModule entity.RoleModule
	err := tx.First(&roleModule).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &roleModule, nil
}
