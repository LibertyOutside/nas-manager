package config

import (
	"embed"
	log "github.com/sirupsen/logrus"
	"os"
)

//go:embed *.yaml
var fs embed.FS

const (
	AppEnv    = "NM_ENV"
	AppEnvDev = "dev"
)

// GetFile 开发环境需要设置环境变量，没有设置就识别为生产环境
func GetFile() []byte {
	env := os.Getenv(AppEnv)
	var f string

	if env == AppEnvDev {

		log.Info("app env: dev")
		f = "application-" + env + ".yaml"
	} else {
		log.Info("app env: prod")
		f = "application.yaml"
	}

	config, _ := fs.ReadFile(f)
	return config
}

// DevMode 判断是否开发环境
func DevMode() bool {
	return os.Getenv(AppEnv) == AppEnvDev
}
