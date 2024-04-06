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

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "service run")
	})

	r.POST("/api/home/login", api.LoginHandler)

	rj := r.Group("/api").Use(middleware.Jwt())
	rj.StaticFS("/static", http.Dir(config.AppConf.FileHost+"/file"))

	rr := r.Group("/api").Use(middleware.Jwt(), middleware.Auth())
	rr.POST("/home/logout", api.LogoutHandler)
	rr.POST("/home/changepassword", api.ChangePasswordHandler)

	rr.GET("/admin/list", api.AdminListHandler)

	return r
}
