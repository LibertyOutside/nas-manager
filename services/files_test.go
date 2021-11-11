package services

import (
	"log"
	"testing"
)

func TestShowFiles(t *testing.T) {
	files, _ := ShowFiles("/home/kylin/下载/")
	for _, file := range files {
		if len(file.Files) > 0 {
			for _, f := range file.Files {
				log.Printf("f: %v", f)
			}
		}
	}
}
