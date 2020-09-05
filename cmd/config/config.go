package config

import (
	"daily/cmd/common"
	"github.com/Unknwon/goconfig"

	"daily/cmd/logger"
)

var Config *goconfig.ConfigFile

type PostgreSQL struct {
	Driver string
	Host   string
}

type FileServer struct {
	Photo       string
	PhotoPath   string
	Service     string
	ServicePath string
}

// init conf
func init() {
	con, err := goconfig.LoadConfigFile(common.ConfigPath)
	if err != nil {
		logger.Log.Error(logger.StrError, logger.Field("[conf]读配置文件失败", err))
	}
	Config = con
}

// get postgres conf
func GetPostgres() PostgreSQL {
	sec, _ := Config.GetSection("postgres")

	mySetting := PostgreSQL{}
	mySetting.Driver = sec["driver"]
	mySetting.Host = sec["host"]

	return mySetting
}

// get file server
func GetFileServer() FileServer {
	sec, err := Config.GetSection("file")
	if err != nil {
		logger.Log.Error(logger.StrError, logger.Field("[conf]读配置文件失败", err))
		panic(err)
	}
	mySetting := FileServer{}
	mySetting.Photo = sec["photo"]
	mySetting.PhotoPath = sec["photopath"]
	mySetting.Service = sec["service"]
	mySetting.ServicePath = sec["servicepath"]

	return mySetting
}

// get ip conf
func GetIP() string {
	url, err := Config.GetValue("server", "url")
	if err != nil {
		logger.Log.Error(logger.StrError, logger.Field("[conf]读配置文件失败", err))
		return ":10001"
	}
	return url
}
