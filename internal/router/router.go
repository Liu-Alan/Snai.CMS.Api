package router

import (
	"net/http"

	"Snai.CMS.Api/common/config"
	"Snai.CMS.Api/internal/api"
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

	rj := r.Group("/api").Use(middleware.Jwt())
	rj.StaticFS("/static", http.Dir(config.AppConf.FileHost+"/file"))
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

	rja.GET("/module/list", api.ModulesHandler)
	rja.GET("/module/get", api.GetModuleHandler)
	rja.POST("/module/add", api.AddModuleHandler)
	rja.POST("/module/update", api.UpdateModuleHandler)
	rja.POST("/module/endisable", api.EnDisableModuleHandler)
	rja.POST("/module/batchendisable", api.BatchEnDisableModuleHandler)
	rja.GET("/module/delete", api.DeleteModuleHandler)
	rja.POST("/module/batchdelete", api.BatchDeleteModuleHandler)

	return r
}
