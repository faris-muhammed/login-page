package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/model"
)

var DB *gorm.DB

func DBconnect() { //database connecting and table creation

	dsn := `host=localhost user=postgres password=501417 dbname=test port=5432`

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	DB = db

	DB.AutoMigrate(&model.UserModel{}, &model.AdminModel{})
}
