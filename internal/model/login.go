package model

type LoginIn struct {
	UserName string `form:"user_name" validate:"required,max=32" label:"用户名"`
	Password string `form:"password" validate:"required" label:"密码"`
}

type LoginOut struct {
	Token string `json:"token"`
}

type ChangePasswordIn struct {
	OldPassword string `form:"old_password" validate:"required" label:"旧密码"`
	Password    string `form:"password" validate:"max=20,min=6" label:"新密码"`
	Password2   string `form:"password2" validate:"eqfield=Password" label:"确认密码"`
}
