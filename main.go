package main

import (
	"fmt"
	"os"
)

func main() {
	ctx := &ScrapingContext{
		year:        "2022",
		semester:    "1",
		facultyID:   "02",
		facultyName: "工学部",
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
