package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"nas-manager/config"
	"nas-manager/routes"
)

func main() {
	config.InitLogger()

	app := gin.Default()
	nasGroup := app.Group("/nas")
	nasGroup.GET("/files", routes.ShowFiles)

	if err := app.Run(); err != nil {
		log.Fatalf("web service start failed: %v", err)
	} // 监听并在 0.0.0.0:8080 上启动服务

}
