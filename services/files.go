package services

import (
	"io/ioutil"
	"nas-manager/models"
	"path/filepath"
)

func ShowFiles(path string) (files []models.File, err error) {
	rd, err := ioutil.ReadDir(path)

	var result []models.File
	for _, fi := range rd {
		f := models.File{Name: fi.Name(), IsDir: fi.IsDir(), Dir: path, Path: filepath.Join(path, fi.Name())}
		if fi.IsDir() {
			ffi, _ := ShowFiles(filepath.Join(path, fi.Name()))
			f.Files = append(f.Files, ffi...)
		} else {

		}
		result = append(result, f)
	}

	return result, err
}
