package config

import (
	"embed"
	log "github.com/sirupsen/logrus"
	"os"
)

//go:embed *.yaml
var fs embed.FS

func GetFile() []byte {
	env := os.Getenv("PROJECT_ENV")
	var f string

	if env == "" {
		log.Info("env:prod")
		f = "application.yaml"
	} else {
		log.Infof("env:%s", env)
		f = "application-" + env + ".yaml"
	}

	config, _ := fs.ReadFile(f)
	return config
}
