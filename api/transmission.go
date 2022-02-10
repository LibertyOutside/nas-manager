package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"nas-manager/services/download"
)

// TrGetTorrents todo:统一错误处理
func TrGetTorrents(c *gin.Context) {
	client := download.TrClient
	torrents, err := client.TorrentGetAll(context.TODO())
	if err != nil {
		c.JSON(200, err)
	}
	c.JSON(200, torrents)
}
