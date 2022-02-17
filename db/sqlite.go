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
	AddDefaultValue()
}

func Tables() {
	log.Infoln("Init Tables ...")
	err := DB.AutoMigrate(&models.TransmissionClient{})
	if err != nil {
		log.Errorf("AutoMigrate Failed!")
	}

}

func AddDefaultValue() {
	// transmission client
	DB.FirstOrCreate(&models.TransmissionClient{Alias: "default", Host: "127.0.0.1"})
}

func init() {
	CreateDatabase()
}
