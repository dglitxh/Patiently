package db

import (
	"log"

	"github.com/dglitxh/patiently/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(db_url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Patient{}, &models.User{}, &models.Records{}, &models.MedicalHx{})

	return db
}
