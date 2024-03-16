package middleware

import (
	"strings"

	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/msg"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 假设Token放在Header的Authorization中，并使用Bearer开头
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authStr string
			mc      *app.Claims
			ecode   = msg.Message{Code: msg.Success, Msg: msg.GetMsg(msg.Success)}
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			authStr = s
		} else {
			authStr = c.GetHeader("Authorization")
		}

		if authStr == "" {
			ecode = msg.Message{Code: msg.InvalidParams, Msg: msg.GetMsg(msg.InvalidParams)}
		} else {
			authParts := strings.SplitN(authStr, " ", 2)
			if !(len(authParts) == 2 && authParts[0] == "Bearer") {
				ecode = msg.Message{Code: msg.AuthFormatFail, Msg: msg.GetMsg(msg.AuthFormatFail)}
			} else {
				token := authParts[1]
				mc, err := app.ParseToken(token)
				if err != nil {
					switch err.(*jwt.ValidationError).Errors {
					case jwt.ValidationErrorExpired:
						ecode = msg.Message{Code: msg.AuthCheckTimeout, Msg: msg.GetMsg(msg.AuthCheckTimeout)}
					default:
						ecode = msg.Message{Code: msg.AuthCheckFail, Msg: msg.GetMsg(msg.AuthCheckFail)}
					}
				}
			}
		}

		if ecode.Code != msg.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(&ecode)
			c.Abort()
			return
		}

		c.Set("user_name", mc.UserName)
		c.Next()
	}
}
