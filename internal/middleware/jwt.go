package middleware

import (
	"app"
	"msg"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = msg.Message{Code: msg.Success, msg.GetMsg(msg.Success)}
		)
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = msg.Message{Code: msg.InvalidParams, msg.GetMsg(msg.InvalidParams)}
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = msg.Message{Code: msg.AuthCheckTimeout, msg.GetMsg(msg.AuthCheckTimeout)}
				default:
					ecode = msg.Message{Code: msg.AuthCheckFail, msg.GetMsg(msg.AuthCheckFail)}
				}
			}
		}

		if ecode.Code != msg.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
