package middleware

import (
	"strings"

	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/message"
	"Snai.CMS.Api/internal/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 假设Token放在Header的Authorization中，并使用Bearer开头
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			authStr string
			ecode   = message.Message{Code: message.Success, Msg: message.GetMsg(message.Success)}
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			authStr = s
		} else {
			authStr = c.GetHeader("Authorization")
		}

		if authStr == "" {
			ecode = message.Message{Code: message.InvalidParams, Msg: message.GetMsg(message.InvalidParams)}
		} else {
			authParts := strings.SplitN(authStr, " ", 2)
			if !(len(authParts) == 2 && authParts[0] == "Bearer") {
				ecode = message.Message{Code: message.AuthFormatFail, Msg: message.GetMsg(message.AuthFormatFail)}
			} else {
				token := authParts[1]
				mc, err := app.ParseToken(token)
				if err != nil {
					switch err.(*jwt.ValidationError).Errors {
					case jwt.ValidationErrorExpired:
						ecode = message.Message{Code: message.AuthCheckTimeout, Msg: message.GetMsg(message.AuthCheckTimeout)}
					default:
						ecode = message.Message{Code: message.AuthCheckFail, Msg: message.GetMsg(message.AuthCheckFail)}
					}
				} else {
					// token是否已退出
					tk, _ := service.GetToken(token)
					if tk == nil || tk.State == 2 {
						ecode = message.Message{Code: message.AuthCheckFail, Msg: message.GetMsg(message.AuthCheckFail)}
					} else {
						c.Set("user_name", mc.UserName)
						c.Set("token", token)
					}
				}
			}
		}

		if ecode.Code != message.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(&ecode)
			c.Abort()
			return
		}

		c.Next()
	}
}
