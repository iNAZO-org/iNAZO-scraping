package main

import (
	"fmt"
	"os"
)

func main() {
	opts, err := parseCommand()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	ctx := &ScrapingContext{
		year:        opts.Year,
		semester:    opts.Semester,
		facultyID:   opts.FacultyID,
		facultyName: FACULTY_ID_TO_NAME[opts.FacultyID],
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
