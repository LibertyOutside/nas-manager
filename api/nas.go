package api

import (
	"github.com/gin-gonic/gin"
	"nas-manager/services"
	"net/http"
)

// ShowFiles 列出指定文件夹的
func ShowFiles(c *gin.Context) {
	//todo:move to settings
	files, err := services.ShowFiles("/home/kylin/下载")
	if err != nil {
		c.Abort()
	}
	c.JSON(http.StatusOK, files)
}
