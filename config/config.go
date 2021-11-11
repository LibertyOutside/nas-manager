package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Configs struct {
	Path      string
	AppConfig AppConfig
}

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

// C 配置单例
var C Configs

// Init viper by env,unmarshal to instance
func (c *Configs) Init() error {
	if c.Path != "" {
		viper.SetConfigFile(c.Path)
	} else {
		viper.SetConfigFile(getConfFileByEnv())
	}
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}
	var appConf AppConfig
	err := viper.Unmarshal(&appConf)
	c.AppConfig = appConf
	if err != nil {
		log.Infof("unable to unmarshal yaml config,error: %v", err)
		return err
	}
	return nil
}

func getConfFileByEnv() string {
	appEnv := os.Getenv("PROJECT_ENV")
	if appEnv == "" {
		log.Printf("environment variable not detected! default value: dev1")
		appEnv = "dev1"
	}
	log.Printf("current env: %s", appEnv)
	return fmt.Sprintf("conf/conf-%s.yaml", appEnv)
}
