package middleware

import (
	"strings"

	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/message"
	"Snai.CMS.Api/internal/service"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.MustGet("user_name").(string)
		token := c.MustGet("token").(string)
		router := c.Request.URL.Path
		router = replaceRoute(router)

		err := service.VerifyUserRole(username, router)
		if err.Code != message.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(err)
			c.Abort()
			return
		}

		c.Set("user_name", username)
		c.Set("token", token)
		c.Next()
	}
}

func replaceRoute(route string) string {
	var newRoute string
	if strings.Contains(route, "/admin/get") {
		newRoute = strings.Replace(route, "/admin/get", "/admin/list", 1)
	} else if strings.Contains(route, "/admin/batchdelete") {
		newRoute = strings.Replace(route, "/admin/batchdelete", "/admin/delete", 1)
	} else if strings.Contains(route, "/admin/batchendisable") {
		newRoute = strings.Replace(route, "/admin/batchendisable", "/admin/endisable", 1)
	} else {
		newRoute = route
	}
	return newRoute
}
