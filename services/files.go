package services

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"nas-manager/models"
	"path/filepath"
	"strings"
)

func ShowFiles(path string) (files []models.File, err error) {
	rd, err := ioutil.ReadDir(path)

	var result []models.File
	for _, fi := range rd {
		f := models.File{Name: fi.Name(), IsDir: fi.IsDir(), Dir: path, Path: filepath.Join(path, fi.Name()), Size: fi.Size()}
		if fi.IsDir() {
			ffi, _ := ShowFiles(filepath.Join(path, fi.Name()))
			f.Files = append(f.Files, ffi...)
		} else {

		}
		result = append(result, f)
	}

	return result, err
}

func Purify(path string) error {
	files, err := ShowFiles(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.IsDir {
			ProcessFile(file)
		}
	}
	return nil
}

func suffix(filename string) string {
	s := strings.Split(filename, ".")
	return s[len(s)-1]
}

func ProcessFile(file models.File) {
	deleteSuffix := "jpg,png,txt,jpeg,gif,html,exe"
	d := strings.Split(deleteSuffix, ",")
	for _, sfx := range d {
		if suffix(file.Name) == sfx {
			log.Infof("not video,deleted: %s", file.Name)
			break
		}
	}
	if file.Size < 100*1024*1024 {
		log.Infof("file size small ,deleted:%s, size:%f", file.Name, float64(file.Size)/float64(1024*1024))
	}
}
