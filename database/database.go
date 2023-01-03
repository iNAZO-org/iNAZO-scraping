package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := os.Getenv("DB_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
