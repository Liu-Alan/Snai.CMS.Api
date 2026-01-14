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
	if err := _cmsdb.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func ModifyAdmin(admin *entity.Admins) error {
	if err := _cmsdb.Save(admin).Error; err != nil {
		return err
	}
	return nil
}

func UpdateAdminState(ids []int, state int8, updateTime int) error {
	result := _cmsdb.Model(&entity.Admins{}).Where("id IN ?", ids).Updates(entity.Admins{State: state, UpdateTime: updateTime})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UnlockAdmin(id int) error {
	result := _cmsdb.Model(&entity.Admins{}).Where("id = ?", id).Update("lock_time", 0)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteAdmin(admin *entity.Admins) error {
	if err := _cmsdb.Delete(admin).Error; err != nil {
		return err
	}
	return nil
}

func BatchDeleteAdmin(ids []int) error {
	var admin entity.Admins
	if err := _cmsdb.Delete(&admin, ids).Error; err != nil {
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
	if err := _cmsdb.Create(token).Error; err != nil {
		return err
	}
	return nil
}

func ModifyToken(token *entity.Tokens) error {
	if err := _cmsdb.Save(token).Error; err != nil {
		return err
	}
	return nil
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

func GetModuleByID(id int) (*entity.Modules, error) {
	tx := _cmsdb.Where("id = ? ", id)
	var module entity.Modules
	err := tx.First(&module).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &module, nil
}

func GetModuleByTitle(parentID int, title string) (*entity.Modules, error) {
	tx := _cmsdb.Where("parent_id = ? And title = ?", parentID, title)
	var module entity.Modules
	err := tx.First(&module).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &module, nil
}

func GetModulesByIDs(ids []int) ([]*entity.Modules, error) {
	tx := _cmsdb.Where("id in ? ", ids)
	var modules []*entity.Modules
	err := tx.Order("sort, id").Find(&modules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return modules, nil
}

func GetModules(pageOffset, pageSize int) ([]*entity.Modules, error) {
	tx := _cmsdb
	if pageOffset >= 0 && pageSize > 0 {
		tx = tx.Offset(pageOffset).Limit(pageSize)
	}
	var modules []*entity.Modules
	err := tx.Find(&modules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return modules, nil
}

func GetModuleMenus() ([]*entity.Modules, error) {
	tx := _cmsdb.Where("menu = 1")
	var modules []*entity.Modules
	err := tx.Find(&modules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return modules, nil
}

func GetModuleCount() (int64, error) {
	tx := _cmsdb
	var count int64
	var admin entity.Modules
	err := tx.Model(&admin).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

func AddModule(module *entity.Modules) error {
	if err := _cmsdb.Create(module).Error; err != nil {
		return err
	}
	return nil
}

func ModifyModule(module *entity.Modules) error {
	if err := _cmsdb.Save(module).Error; err != nil {
		return err
	}
	return nil
}

func UpdateModuleState(ids []int, state int8) error {
	result := _cmsdb.Model(&entity.Modules{}).Where("id IN ?", ids).Update("state", state)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteModule(module *entity.Modules) error {
	if err := _cmsdb.Delete(module).Error; err != nil {
		return err
	}
	return nil
}

func BatchDeleteModule(ids []int) error {
	var module entity.Modules
	if err := _cmsdb.Delete(&module, ids).Error; err != nil {
		return err
	}
	return nil
}

func GetRoleByID(roleID int) (*entity.Roles, error) {
	tx := _cmsdb.Where("id = ? ", roleID)
	var role entity.Roles
	err := tx.First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &role, nil
}

func GetRoleByTitle(title string) (*entity.Roles, error) {
	tx := _cmsdb.Where("title = ?", title)
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

func GetRoleCount() (int64, error) {
	tx := _cmsdb
	var count int64
	var role entity.Roles
	err := tx.Model(&role).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

func AddRole(role *entity.Roles) error {
	if err := _cmsdb.Create(role).Error; err != nil {
		return err
	}
	return nil
}

func ModifyRole(role *entity.Roles) error {
	if err := _cmsdb.Save(role).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRoleState(ids []int, state int8) error {
	result := _cmsdb.Model(&entity.Roles{}).Where("id IN ?", ids).Update("state", state)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteRole(role *entity.Roles) error {
	if err := _cmsdb.Delete(role).Error; err != nil {
		return err
	}
	return nil
}

func BatchDeleteRole(ids []int) error {
	var role entity.Roles
	if err := _cmsdb.Delete(&role, ids).Error; err != nil {
		return err
	}
	return nil
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

func DeleteRoleModules(roleID int) error {
	var roleModule entity.RoleModule
	if err := _cmsdb.Delete(&roleModule, "role_id = ?", roleID).Error; err != nil {
		return err
	}
	return nil
}

func AddRoleModules(roleModules []*entity.RoleModule) error {
	if err := _cmsdb.Create(roleModules).Error; err != nil {
		return err
	}
	return nil
}
