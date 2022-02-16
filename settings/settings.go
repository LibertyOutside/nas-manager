package settings

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"nas-manager/config"
)

type AppConfig struct {
	Debug      bool   `yaml:"debug"`
	Loglevel   string `yaml:"loglevel"`
	Port       string `yaml:"port"`
	SqlitePath string `yaml:"sqlite_path"`
}

var App AppConfig

func InitSettings() {
	err := yaml.Unmarshal(config.GetFile(), &App)
	if err != nil {
		log.Fatalf("failed to read project settings! because of: %v", err)
	}
	log.Infoln("Project Settings Init Successfully")
}

func init() {
	InitSettings()

}
