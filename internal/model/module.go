package model

type ModuleOut struct {
	Key         int    `json:"key"`
	ID          int    `json:"id"`
	ParentID    int    `json:"parent_id"`
	ParentTitle string `json:"parent_title"`
	Title       string `json:"title"`
	Name        string `json:"name"`
	Router      string `json:"router"`
	UIRouter    string `json:"ui_router"`
	Menu        int8   `json:"menu"` // 是否菜单：1 是，2 否
	Sort        int    `json:"sort"`
	State       int8   `json:"state"`
}

type AddModuleIn struct {
	Title    string `form:"title" validate:"required,max=32" label:"名称"`
	Name     string `form:"name" validate:"max=32" label:"前端名称"`
	ParentID int    `form:"parent_id" validate:"gte=-1" label:"父模块"`
	Router   string `form:"router" label:"api路由"`
	UIRouter string `form:"ui_router" label:"前端路由"`
	Sort     int    `form:"sort" validate:"gte=1" label:"排序"`
	Menu     int8   `form:"menu" validate:"oneof=1 2" label:"菜单"`
	State    int8   `form:"state" validate:"oneof=1 2" label:"状态"`
}

type UpdateModuleIn struct {
	ID       int    `form:"id" validate:"gte=1" label:"ID"`
	Title    string `form:"title" validate:"required,max=32" label:"名称"`
	Name     string `form:"name" validate:"max=32" label:"前端名称"`
	ParentID int    `form:"parent_id" validate:"gte=-1" label:"父模块"`
	Router   string `form:"router" label:"api路由"`
	UIRouter string `form:"ui_router" label:"前端路由"`
	Sort     int    `form:"sort" validate:"gte=1" label:"排序"`
	Menu     int8   `form:"menu" validate:"oneof=1 2" label:"菜单"`
	State    int8   `form:"state" validate:"oneof=1 2" label:"状态"`
}

type EnDisableModuleIn struct {
	ID    int  `form:"id" validate:"gte=1" label:"ID"`
	State int8 `form:"state" validate:"oneof=1 2" label:"状态"`
}

type BatchEnDisableModuleIn struct {
	IDs   []int `form:"ids[]" validate:"min=1,dive,gte=1" label:"ids"`
	State int8  `form:"state" validate:"oneof=1 2" label:"状态"`
}

type BatchDeleteModuleIn struct {
	IDs []int `form:"ids[]" validate:"min=1,dive,gte=1" label:"ids"`
}
