package settings

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"nas-manager/config"
)

type AppConfig struct {
	Mysql     MysqlStruct
	WhiteList []string
}

type MysqlStruct struct {
	User   string
	Passwd string
	Host   string
	Port   int
	Db     string
}

var App AppConfig

func InitSettings() {
	err := yaml.Unmarshal(config.GetFile(), &App)
	if err != nil {
		log.Fatalf("config init failed！ cause: %v", err)
	}
}
