package services

import (
	log "github.com/sirupsen/logrus"
	"io/fs"
	"io/ioutil"
)

func ShowFiles(path string) (files []fs.FileInfo, err error) {
	rd, err := ioutil.ReadDir(path)
	for _, fi := range rd {
		if fi.IsDir() {
			log.Infof("%#v", fi)
		} else {

		}
	}

	return rd, nil
}
