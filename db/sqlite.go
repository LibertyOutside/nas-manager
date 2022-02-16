package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nas-manager/models"
)

var DB *gorm.DB

func CreateDatabase() {
	db, err := gorm.Open(sqlite.Open("nasmanager.db"), &gorm.Config{})
	if err != nil {
		log.Errorf("Failed To Init Database!%v\n", err)
	}
	DB = db
	Tables()
}

func Tables() {
	log.Infoln("Init Tables ...")
	err := DB.AutoMigrate(&models.TransmissionClient{})
	if err != nil {
		log.Errorf("AutoMigrate Failed!")
	}
}
