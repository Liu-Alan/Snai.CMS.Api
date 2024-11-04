package model

type AdminsIn struct {
	UserName string `form:"user_name" validate:"max=32" label:"用户名"`
}

type AdminsOut struct {
	Key       int    `json:"key"`
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	Role      string `json:"role"`
	State     int8   `json:"state"`
	LockState int8   `json:"lock_state"` // 1 未锁定 2 锁定
}

type AddAdminIn struct {
	UserName  string `form:"user_name" validate:"required,max=32" label:"用户名"`
	Password  string `form:"password" validate:"max=20,min=6" label:"密码"`
	Password2 string `form:"password2" validate:"eqfield=Password" label:"确认密码"`
	RoleID    int    `form:"role_id" validate:"gte=1" label:"角色"`
	State     int8   `form:"state" validate:"oneof=1 2" label:"状态"`
}
