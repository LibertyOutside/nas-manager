package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	Debug        bool   `yaml:"debug"`
	Loglevel     string `yaml:"loglevel"`
	WebPort      string `yaml:"webPort"`
	SqlitePath   string `yaml:"sqlitePath"`
	DownloadPath string `yaml:"downloadPath"`
}

var App AppConfig

const (
	AppEnv    = "NM_ENV"
	AppEnvDev = "DEV"
)

func InitViper() {
	viper.SetConfigName("nasmanager")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/nasmanager/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("failed to read project settings: %v", err)
	}
	err = viper.Unmarshal(&App)
	if err != nil {
		log.Fatalf("failed to read project settings: %v", err)
	}
	viper.WatchConfig()
}

// DevMode 判断是否开发环境
func DevMode() bool {
	return os.Getenv(AppEnv) == AppEnvDev
}
