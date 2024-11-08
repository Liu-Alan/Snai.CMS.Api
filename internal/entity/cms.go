package entity

type Admins struct {
	ID              int    `gorm:"column:id"`
	UserName        string `gorm:"column:user_name"`
	Password        string `gorm:"column:password"`
	RoleID          int    `gorm:"column:role_id"`
	State           int8   `gorm:"column:state"`
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

type Tokens struct {
	ID         int    `gorm:"column:id"`
	Token      string `gorm:"column:token"`
	UserID     int    `gorm:"column:user_id"`
	State      int8   `gorm:"column:state"`
	CreateTime int    `gorm:"column:create_time"`
}

func (Tokens) TableName() string {
	return "tokens"
}

type Modules struct {
	ID       int    `gorm:"column:id"`
	ParentID int    `gorm:"column:parent_id"`
	Title    string `gorm:"column:title"`
	Name     string `gorm:"column:name"`
	Router   string `gorm:"column:router"`
	UIRouter string `gorm:"column:ui_router"`
	Menu     int8   `gorm:"column:menu"`
	Sort     int    `gorm:"column:sort"`
	State    int8   `gorm:"column:state"`
}

func (Modules) TableName() string {
	return "modules"
}

type Roles struct {
	ID    int    `gorm:"column:id"`
	Title string `gorm:"column:title"`
	State int8   `gorm:"column:state"`
}

func (Roles) TableName() string {
	return "roles"
}

type RoleModule struct {
	ID       int `gorm:"column:id"`
	RoleID   int `gorm:"column:role_id"`
	ModuleID int `gorm:"column:module_id"`
}

func (RoleModule) TableName() string {
	return "role_module"
}
