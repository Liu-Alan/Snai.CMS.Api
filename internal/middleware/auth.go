package middleware

import (
	"Snai.CMS.Api/common/msg"
	"Snai.CMS.Api/internal/service"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.MustGet("username").(string)
		router := c.Request.URL.Path

		err := service.VerifyUserRole(username, router)
		if err.Code != msg.Success {
			c.Abort()
			return
		}

		c.Set("user_name", username)
		c.Next()
	}
}
