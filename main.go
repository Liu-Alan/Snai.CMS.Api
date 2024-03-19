package main

import (
	"Snai.CMS.Api/common/app"
	"Snai.CMS.Api/common/config"
	"Snai.CMS.Api/common/logging"
	"Snai.CMS.Api/internal/dao"
	"Snai.CMS.Api/internal/router"
	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化配置
	config.InitConfig("./config.json")
	// 初始化log
	logging.InitLog(logging.LogLevel(config.AppConf.LogLevel), config.AppConf.LogTargets, config.AppConf.FileHost+"/logs")
	logging.InitSlog(logging.LogLevel(config.AppConf.LogLevel), config.AppConf.LogTargets, config.AppConf.FileHost+"/logs_sql")

	logging.Info("初始化config成功")
	logging.Info("初始化log成功")

	// 初始化数据库
	dao.InitDB()

	// 补始化验证
	app.InitValidator()

	logging.Info("服务启动,监听端口: %s", config.AppConf.Port)
	logging.Warn("当前环境: %v", config.AppConf.Env)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := router.NewRouter()
	r.Run(":" + config.AppConf.Port)
}
