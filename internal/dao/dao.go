package dao

import (
	"time"

	"Snai.CMS.Api/common/config"
	"Snai.CMS.Api/common/logging"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

var (
	_cmsdb *gorm.DB
)

// 自定义sql日志
var newLogger = logger.New(
	logging.SlogWriter{},
	logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	},
)

// InitDB 初始化DB
func InitDB() {
	var err error
	//初始化cms
	_cmsdb, err = gorm.Open(mysql.Open(config.AppConf.CmsDB), &gorm.Config{Logger: newLogger})
	if err != nil {
		logging.Fatal("gorm.Open(%v): %v", config.AppConf.CmsDB, err)
	}
	_cmsdbM, _ := _cmsdb.DB()
	_cmsdb.SetConnMaxLifetime(time.Second * 300)
	_cmsdb.SetMaxIdleConns(5)
	_cmsdb.SetMaxOpenConns(100)
	logging.Info("cms库连接成功")
}

func GetCmsDB() *gorm.DB {
	return _cmsdb
}

type ReTotal struct {
	Total int `gorm:"column:total"`
}
