package controllers

import (
	"karintou8710/iNAZO-scraping/database"
	"karintou8710/iNAZO-scraping/models"
)

func MigrateHandler() error {
	db := database.GetDB()
	err := db.AutoMigrate(&models.GradeDistribution{})
	if err != nil {
		return err
	}

	return nil
}
