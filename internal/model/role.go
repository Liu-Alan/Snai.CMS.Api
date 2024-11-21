package model

type RoleOut struct {
	Key   int    `json:"key"`
	ID    int    `json:"id"`
	Title string `json:"title"`
	State int8   `json:"state"`
}

type AddRoleIn struct {
	Title string `form:"title" validate:"required,max=32" label:"名称"`
	State int8   `form:"state" validate:"oneof=1 2" label:"状态"`
}

type UpdateRoleIn struct {
	ID    int    `form:"id" validate:"gte=1" label:"ID"`
	Title string `form:"title" validate:"required,max=32" label:"名称"`
	State int8   `form:"state" validate:"oneof=1 2" label:"状态"`
}

type EnDisableRoleIn struct {
	ID    int  `form:"id" validate:"gte=1" label:"ID"`
	State int8 `form:"state" validate:"oneof=1 2" label:"状态"`
}

type BatchEnDisableRoleIn struct {
	IDs   []int `form:"ids[]" validate:"min=1,dive,gte=1" label:"ids"`
	State int8  `form:"state" validate:"oneof=1 2" label:"状态"`
}

type BatchDeleteRoleIn struct {
	IDs []int `form:"ids[]" validate:"min=1,dive,gte=1" label:"ids"`
}

type AssignPermIn struct {
	RoleID    int   `form:"role_id" validate:"gte=1" label:"角色id"`
	ModuleIDs []int `form:"module_ids[]" validate:"min=1,dive,gte=1" label:"模块ids"`
}
