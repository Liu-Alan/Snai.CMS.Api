package app

import (
	"strconv"

	"Snai.CMS.Api/common/config"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		return 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize <= 0 {
		return config.AppConf.DefaultPageSize
	}
	if pageSize > config.AppConf.MaxPageSize {
		return config.AppConf.MaxPageSize
	}

	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}

	return result
}
