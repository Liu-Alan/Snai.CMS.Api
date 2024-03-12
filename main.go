package main

import (
	"net/http"

	"Snai.CMS.Api/common/config"
	"Snai.CMS.Api/common/logger"
	"Snai.CMS.Api/internal/dao"
	"github.com/gin-gonic/gin"
)

func init() {
	// 补始化验证

	// 初始化配置
	config.InitConfig("./config.json")
	// 初始化log
	logger.InitLog(logger.LogLevel(config.GetConf.LogLevel), config.GetConf.LogTargets, config.GetConf.FileHost+"/logs")
	logger.InitSlog(logger.LogLevel(config.GetConf.LogLevel), config.GetConf.LogTargets, config.GetConf.FileHost+"/logs_sql")

	logger.Info("初始化config成功")
	logger.Info("初始化log成功")

	// 初始化数据库
	dao.InitDB()

	logger.Info("服务启动,监听端口: %s", config.GetConf.Port)
	logger.Warn("当前环境: %v", config.GetConf.Env)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "service run")
	})

	r.Run(":" + config.GetConf.Port)
}
