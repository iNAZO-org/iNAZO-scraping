package controllers

import (
	"fmt"

	"karintou8710/iNAZO-scraping/csv"
	"karintou8710/iNAZO-scraping/scraping"
	"karintou8710/iNAZO-scraping/setting"
	"karintou8710/iNAZO-scraping/utils"
)

type ScrapingCommand struct {
	Positional struct {
		Year      int
		Semester  int
		FacultyID string
	} `positional-args:"yes" required:"yes"`
}

var scrapingCommand ScrapingCommand

func ScrapingHandler(year int, semester int, facultyIDByCmd string) error {
	var facultyIDList []string // allに対応する

	if facultyIDByCmd == "all" {
		facultyIDList = utils.GetKeysFromMap(setting.FacultyIdToName)
	} else {
		facultyIDList = []string{facultyIDByCmd}
	}

	for _, faclutyID := range facultyIDList {

		facultyName := setting.FacultyIdToName[faclutyID]

		fmt.Printf("scraping %d年%d学期 %s... 🚀\n", year, semester, facultyName)
		result, err := scraping.ScrapingGradeDistribution(year, semester, faclutyID, facultyName)
		if err != nil {
			return err
		}

		fmt.Printf("writing data/%d%d/%s.csv... 🚀\n", year, semester, facultyName)
		err = csv.WriteGradeDistributionToCSV(year, semester, facultyName, result)
		if err != nil {
			return err
		}
	}

	return nil
}
