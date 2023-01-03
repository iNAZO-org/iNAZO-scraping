package controllers

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"karintou8710/iNAZO-scraping/csv"
	"karintou8710/iNAZO-scraping/database"
	"karintou8710/iNAZO-scraping/models"
)

func existGradeDistributionRow(gd *models.GradeDistribution) (bool, error) {
	var exists bool
	db := database.GetDB()
	err := db.Model(gd).
		Select("count(*) > 0").
		Where(gd).
		Find(&exists).
		Error
	return exists, err
}

func insertGradeDistributionList(gdList []*models.GradeDistribution) error {
	db := database.GetDB()
	for _, gd := range gdList {
		existFlag, err := existGradeDistributionRow(gd)
		if err != nil {
			return err
		}
		// 既に登録済みの項目は2重に保存しない
		if existFlag {
			continue
		}

		if err := db.Create(&gd).Error; err != nil {
			return err
		}
	}

	return nil
}

func UpdateGradeHandler() error {
	err := filepath.Walk("./data", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Printf("saving %s...\n", path)
		gdList, err := csv.ReadGradeDistributionFromCSV(path)
		if err != nil {
			return err
		}
		err = insertGradeDistributionList(gdList)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
