package api

import (
	"github.com/gin-gonic/gin"
	"nas-manager/models"
	"nas-manager/services/download"
)

// TrGetTorrents todo:查询过慢优化
func TrGetTorrents(c *gin.Context) {
	result := models.NewResult(c)
	torrents, err := download.GetTrTorrents()
	if err != nil {
		result.Error(200, "Cannot get torrents!")
	}
	result.Success(torrents)
}

func GetTransmissionClients(c *gin.Context) {
	models.NewResult(c).Success(download.TrClient)
}
