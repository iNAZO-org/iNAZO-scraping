package controllers

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"karintou8710/iNAZO-scraping/csv"
	"karintou8710/iNAZO-scraping/models"
)

func existGradeDistributionRow(gd models.GradeDistribution) (bool, error) {
	var cnt int
	existFlag := cnt > 0
	return existFlag, nil
}

func insertGradeDistributionList(gdList []models.GradeDistribution) error {
	for _, gd := range gdList {
		existFlag, err := existGradeDistributionRow(gd)
		if err != nil {
			return err
		}
		// 既に登録済みの項目は2重に保存しない
		if existFlag {
			continue
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
