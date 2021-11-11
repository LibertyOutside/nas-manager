package services

import (
	"io/ioutil"
	"nas-manager/models"
)

func ShowFiles(path string) (files []models.File, err error) {
	rd, err := ioutil.ReadDir(path)

	var result []models.File
	for _, fi := range rd {
		f := models.File{Name: fi.Name(), IsDir: fi.IsDir(), Dir: path, Path: path + "/" + fi.Name()}
		if fi.IsDir() {
			ffi, _ := ShowFiles(path + "/" + fi.Name())
			f.Files = append(f.Files, ffi...)
		} else {

		}
		result = append(result, f)
	}

	return result, nil
}
