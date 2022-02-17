package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"nas-manager/api"
	"nas-manager/db"
	"nas-manager/settings"
	"net/http"
)

func main() {
	//todo:move to init

	//settings.InitLogger()
	//settings.InitSettings()

	db.CreateDatabase()
	app := gin.Default()

	if settings.App.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	app.Use(Cors())
	nasGroup := app.Group("/nas")
	nasGroup.GET("/files", api.ShowFiles)
	downloadGroup := app.Group("/downloads")
	downloadGroup.GET("/torrents", api.TrGetTorrents)
	downloadGroup.GET("/clients", api.GetTransmissionClients)

	if err := app.Run(settings.App.Port); err != nil {
		log.Fatalf("web service start failed: %v", err)
	}

}

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
