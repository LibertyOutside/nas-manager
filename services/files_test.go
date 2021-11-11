package services

import (
	log "github.com/sirupsen/logrus"
	"nas-manager/config"
	"testing"
)

func TestShowFiles(t *testing.T) {
	config.InitLogger()
	files, _ := ShowFiles("/home/kylin/下载")
	for _, file := range files {
		if len(file.Files) > 0 {
			for _, f := range file.Files {
				log.Printf("f: %v", f)
			}
		}
	}
}
