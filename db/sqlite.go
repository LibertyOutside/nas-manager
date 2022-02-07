package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDatabase() {
	db, err := gorm.Open(sqlite.Open("nasmanager.db"), &gorm.Config{})
	if err != nil {
		println(db.Config)
	}
}
