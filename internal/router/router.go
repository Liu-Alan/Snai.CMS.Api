package router

import (
	"net/http"

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

	r.POST("/api/login", api.LoginHandler)

	rr := r.Group("/api").Use(middleware.Jwt(), middleware.Auth())
	rr.POST("/logout", api.LogoutHandler)
	rr.POST("/changepassword", api.ChangePasswordHandler)

	return r
}
