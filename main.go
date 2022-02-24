package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"nas-manager/routes"
	"nas-manager/settings"
)

func main() {

	app := gin.Default()

	if settings.App.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	routes.SetRoutes(app)
	app.Use(routes.Cors())

	if err := app.Run(settings.App.Port); err != nil {
		log.Fatalf("web service start failed: %v", err)
	}

}
