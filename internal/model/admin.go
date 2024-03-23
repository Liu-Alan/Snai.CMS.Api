package model

type AdminsIn struct {
	UserName string `form:"user_name" validate:"max=32" label:"用户名"`
}
