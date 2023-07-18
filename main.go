package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"nas-manager/config"
	"nas-manager/routes"
)

func main() {
	config.InitViper()
	if config.DevMode() == false {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	app.Use(routes.Cors())

	routes.SetRoutes(app)

	log.Infoln("initiation finished，starting web service")
	if err := app.Run(config.App.WebPort); err != nil {
		log.Fatalf("web service start failed: %v", err)
	}

}
