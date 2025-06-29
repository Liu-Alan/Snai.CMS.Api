package model

type LoginIn struct {
	UserName string `form:"user_name" validate:"required,max=32" label:"用户名"`
	Password string `form:"password" validate:"required" label:"密码"`
	OtpCode  string `form:"otp_code" validate:"required,len=6" label:"Otp动态码"`
}

type LoginOut struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
}

type ChangePasswordIn struct {
	OldPassword string `form:"old_password" validate:"required" label:"旧密码"`
	Password    string `form:"password" validate:"passwd" label:"新密码"`
	Password2   string `form:"password2" validate:"eqfield=Password" label:"确认密码"`
}

type MenuOut struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	Router   string `json:"router"`
	UIRouter string `json:"ui_router"`
	Menu     int8   `json:"menu"`
	Sort     int    `json:"sort"`
}
