package router

import (
	"net/http"

	"Snai.CMS.Api/common/config"
	"Snai.CMS.Api/internal/api"
	"Snai.CMS.Api/internal/common"
	"Snai.CMS.Api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "service run")
	})

	r.POST("/api/home/login", api.LoginHandler)

	rs := r.Group("/static").Use(middleware.Jwt())
	rs.StaticFS("/upload", common.NoListFS{FS: http.Dir(config.AppConf.FileHost + "/upload")})

	rj := r.Group("/api").Use(middleware.Jwt())
	rj.GET("/home/menu", api.MenuHandler)
	rj.GET("/home/role", api.RoleHandler)

	rja := r.Group("/api").Use(middleware.Jwt(), middleware.Auth())
	rja.POST("/home/logout", api.LogoutHandler)
	rja.POST("/home/changepassword", api.ChangePasswordHandler)

	rja.GET("/admin/list", api.AdminsHandler)
	rja.GET("/admin/get", api.GetAdminHandler)
	rja.POST("/admin/add", api.AddAdminHandler)
	rja.POST("/admin/update", api.UpdateAdminHandler)
	rja.POST("/admin/endisable", api.EnDisableAdminHandler)
	rja.POST("/admin/batchendisable", api.BatchEnDisableAdminHandler)
	rja.GET("/admin/unlock", api.UnlockAdminHandler)
	rja.GET("/admin/delete", api.DeleteAdminHandler)
	rja.POST("/admin/batchdelete", api.BatchDeleteAdminHandler)
	rja.GET("/admin/qrcode", api.GetAdminQrcodeHandler)

	rja.GET("/module/list", api.ModulesHandler)
	rja.GET("/module/getlist", api.GetModulesHandler)
	rja.GET("/module/get", api.GetModuleHandler)
	rja.POST("/module/add", api.AddModuleHandler)
	rja.POST("/module/update", api.UpdateModuleHandler)
	rja.POST("/module/endisable", api.EnDisableModuleHandler)
	rja.POST("/module/batchendisable", api.BatchEnDisableModuleHandler)
	rja.GET("/module/delete", api.DeleteModuleHandler)
	rja.POST("/module/batchdelete", api.BatchDeleteModuleHandler)

	rja.GET("/role/list", api.RolesHandler)
	rja.GET("/role/get", api.GetRoleHandler)
	rja.POST("/role/add", api.AddRoleHandler)
	rja.POST("/role/update", api.UpdateRoleHandler)
	rja.POST("/role/endisable", api.EnDisableRoleHandler)
	rja.POST("/role/batchendisable", api.BatchEnDisableRoleHandler)
	rja.GET("/role/delete", api.DeleteRoleHandler)
	rja.POST("/role/batchdelete", api.BatchDeleteRoleHandler)

	rja.GET("/role/rolemodules", api.RoleModulesHandler)
	rja.POST("/role/assignperm", api.AssignPermHandler)

	return r
}
