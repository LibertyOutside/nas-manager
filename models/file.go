package models

type File struct {
	Name  string
	Dir   string
	IsDir bool
	Files []File
	Path  string
}
