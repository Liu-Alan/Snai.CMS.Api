package model

type AdminIn struct {
	UserName string `form:"user_name" validate:"max=32" label:"用户名"`
}

type AdminOut struct {
	Key       int    `json:"key"`
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	RoleID    int    `json:"role_id"`
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

type UpdateAdminIn struct {
	ID        int    `form:"id" validate:"gte=1" label:"ID"`
	UserName  string `form:"user_name" validate:"required,max=32" label:"用户名"`
	Password  string `form:"password" validate:"max=20,min=6" label:"密码"`
	Password2 string `form:"password2" validate:"eqfield=Password" label:"确认密码"`
	RoleID    int    `form:"role_id" validate:"gte=1" label:"角色"`
	State     int8   `form:"state" validate:"oneof=1 2" label:"状态"`
}

type EnDisableAdminIn struct {
	ID    int  `form:"id" validate:"gte=1" label:"ID"`
	State int8 `form:"state" validate:"oneof=1 2" label:"状态"`
}

type BatchEnDisableAdminIn struct {
	IDs   []int `form:"ids[]" validate:"min=1,dive,gte=1" label:"ids"`
	State int8  `form:"state" validate:"oneof=1 2" label:"状态"`
}

type BatchDeleteAdminIn struct {
	IDs []int `form:"ids[]" validate:"min=1,dive,gte=1" label:"ids"`
}
