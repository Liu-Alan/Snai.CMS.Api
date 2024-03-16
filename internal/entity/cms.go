package entity

type Admins struct {
	ID              int    `gorm:"column:id"`
	UserName        string `gorm:"column:user_name"`
	PassWord        string `gorm:"column:password"`
	RoleID          int    `gorm:"column:role_id"`
	State           int16  `gorm:"column:state"`
	CreateTime      int    `gorm:"column:create_time"`
	UpdateTime      int    `gorm:"column:update_time"`
	LastLogonTime   int    `gorm:"column:last_logon_time"`
	LastLogonIP     string `gorm:"column:last_logon_ip"`
	ErrorLogonTime  int    `gorm:"column:error_logon_time"`
	ErrorLogonCount int    `gorm:"column:error_logon_count"`
	LockTime        int    `gorm:"column:lock_time"`
}

func (Admins) TableName() string {
	return "admins"
}

type Roles struct {
	ID    int    `gorm:"column:id"`
	Title string `gorm:"column:title"`
	State int16  `gorm:"column:state"`
}

func (Roles) TableName() string {
	return "roles"
}

type Modules struct {
	ID       int    `gorm:"column:id"`
	ParentID int    `gorm:"column:parent_id"`
	Title    string `gorm:"column:title"`
	Router   string `gorm:"column:router"`
	Sort     int    `gorm:"column:sort"`
	State    int16  `gorm:"column:state"`
}

func (Modules) TableName() string {
	return "modules"
}

type RoleModule struct {
	ID       int `gorm:"column:id"`
	RoleID   int `gorm:"column:role_id"`
	ModuleID int `gorm:"column:module_id"`
}

func (RoleModule) TableName() string {
	return "role_module"
}
