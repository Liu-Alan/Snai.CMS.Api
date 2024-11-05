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

func GetAdminByID(id int) (*entity.Admins, error) {
	tx := _cmsdb.Where("id = ? ", id)
	var admin entity.Admins
	err := tx.First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &admin, nil
}

func GetAdminCount(userName string) (int64, error) {
	tx := _cmsdb
	if userName != "" {
		tx = _cmsdb.Where("user_name like ? ", "%"+userName+"%")
	}

	var count int64
	var admin entity.Admins
	err := tx.Model(&admin).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

func GetAdmins(userName string, pageOffset, pageSize int) ([]*entity.Admins, error) {
	tx := _cmsdb
	if userName != "" {
		tx = _cmsdb.Where("user_name like ? ", "%"+userName+"%")
	}
	if pageOffset >= 0 && pageSize > 0 {
		tx = tx.Offset(pageOffset).Limit(pageSize)
	}
	var admins []*entity.Admins
	err := tx.Find(&admins).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return admins, nil
}

func AddAdmin(admin *entity.Admins) error {
	if err := _cmsdb.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}

func ModifyAdmin(admin *entity.Admins) error {
	if err := _cmsdb.Save(&admin).Error; err != nil {
		return err
	}
	return nil
}

func GetToken(token string) (*entity.Tokens, error) {
	tx := _cmsdb.Where("token = ? ", token)
	var tk entity.Tokens
	err := tx.First(&tk).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &tk, nil
}

func AddToken(token *entity.Tokens) error {
	if err := _cmsdb.Create(&token).Error; err != nil {
		return err
	}
	return nil
}

func ModifyToken(token *entity.Tokens) error {
	if err := _cmsdb.Save(&token).Error; err != nil {
		return err
	}
	return nil
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

func GetRoles(pageOffset, pageSize int) ([]*entity.Roles, error) {
	tx := _cmsdb
	if pageOffset >= 0 && pageSize > 0 {
		tx = tx.Offset(pageOffset).Limit(pageSize)
	}
	var roles []*entity.Roles
	err := tx.Find(&roles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return roles, nil
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

func GetModules(ids []int) ([]*entity.Modules, error) {
	tx := _cmsdb.Where("id in ? ", ids)
	var modules []*entity.Modules
	err := tx.Order("sort, id").Find(&modules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return modules, nil
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

func GetRoleModules(roleID int) ([]*entity.RoleModule, error) {
	tx := _cmsdb.Where("role_id = ? ", roleID)
	var roleModules []*entity.RoleModule
	err := tx.Find(&roleModules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return roleModules, nil
}
