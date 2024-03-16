package router

import (
	"net/http"

	"Snai.CMS.Api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "service run")
	})

	rr := r.Group("/api").Use(middleware.Jwt(), middleware.Auth())

	return r
}
