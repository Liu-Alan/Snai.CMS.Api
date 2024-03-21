package model

type LoginIn struct {
	UserName string `form:"user_name" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type LoginOut struct {
	Token string `json:"token"`
}

type ChangePasswordIn struct {
	OldPassword string `form:"old_password" validate:"required"`
	Password    string `form:"password" validate:"max=20,min=6"`
	Password2   string `form:"password2" validate:"eqfield=Password"`
}
