package middleware

import (
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
