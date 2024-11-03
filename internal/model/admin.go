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
