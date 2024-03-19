package model

type LoginIn struct {
	UserName string `form:"user_name" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type LoginOut struct {
	Token string `json:"token"`
}
