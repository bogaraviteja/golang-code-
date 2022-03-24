package database

import (
	"project/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("project.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Employee{})
	return db
}
