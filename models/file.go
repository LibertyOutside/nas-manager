package models

type File struct {
	Name  string `json:"name"`
	Dir   string `json:"dir"`
	IsDir bool   `json:"is_dir"`
	Files []File `json:"files"`
	Path  string `json:"path"`
	Size  int64  `json:"size"`
}
