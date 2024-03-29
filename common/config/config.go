package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var AppConf = &Config{}

// Config 系统配置
type Config struct {
	Env             string `json:"env"`
	CmsDB           string `json:"cms.db"`
	LogLevel        int    `json:"log.level"`
	LogTargets      int    `json:"log.targets"`
	LogDB           bool   `json:"log.db"`
	FileHost        string `json:"file.host"`
	Port            string `json:"port"`
	PwdSalt         string `json:"pwd.salt"`
	JwtSecret       string `json:"jwt.secret"`
	JwtIssuer       string `json:"jwt.issuer"`
	JwtExpire       int    `json:"jwt.expire"`
	LoginLockMinute int    `json:"login.lockminute"`
	LoginErrorCount int    `json:"login.errorcount"`
	DefaultPageSize int    `json:"defaultpagesize"`
	MaxPageSize     int    `json:"maxpagesize"`
}

// InitConfig 初始化配置文件
func InitConfig(path string) {
	// 先加载基础配置，再加载不同环境下的配置文件
	envs := []string{"develop", "test", "prerelease", "production"}
	dir, file := filepath.Split(path)

	for _, x := range envs {
		file = strings.Replace(file, "."+x, "", -1)
	}

	// basicPath 是去掉了环境变量的配置文件，例如: config.json
	basicPath := dir + file

	// 加载环境配置文件
	if basicPath != path {
		_, err := os.Stat(basicPath)
		if err == nil {
			readConfigFile(basicPath)
		}
	}

	readConfigFile(path)
}

func readConfigFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("readConfigFile: %v\n", err)
		return
	}

	err = json.Unmarshal(data, AppConf)
	if err != nil {
		log.Fatalf("readConfigFile: %v\n", err)
		return
	}
}
