package utils

import "github.com/gin-gonic/gin"

func GetGinContextByKey(c *gin.Context, key string) (string, bool) {
	v, ok := c.Get(key)
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}
