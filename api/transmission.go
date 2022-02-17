package api

import (
	"github.com/gin-gonic/gin"
	"nas-manager/services/download"
)

// TrGetTorrents todo:统一错误处理
func TrGetTorrents(c *gin.Context) {
	torrents, err := download.GetTrTorrents()
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, torrents)
}

func GetTransmissionClients(c *gin.Context) {
	c.JSON(200, download.TrClient)
}
