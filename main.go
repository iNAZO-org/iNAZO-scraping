package main

import (
	"fmt"
	"os"
)

func main() {
	year := "2022"
	semester := "1"
	facultyID := "02"
	ctx := &ScrapingContext{
		year:        year,
		semester:    semester,
		facultyID:   facultyID,
		facultyName: FACULTY_ID_TO_NAME[facultyID],
	}

	fmt.Println("scraping... 🚀")
	result, err := scrapingGradeDistribution(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Println("writing csv file... 🚀")
	err = writeGradeDistibutionToCSV(ctx, result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
}
