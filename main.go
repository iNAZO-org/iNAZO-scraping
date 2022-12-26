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

	for _, faclutyID := range opts.facultyIDList {
		ctx := &ScrapingContext{
			year:        opts.Year,
			semester:    opts.Semester,
			facultyID:   faclutyID,
			facultyName: FACULTY_ID_TO_NAME[faclutyID],
		}

		fmt.Printf("scraping %s年%s学期 %s... 🚀\n", opts.Year, opts.Semester, ctx.facultyName)
		result, err := scrapingGradeDistribution(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		fmt.Printf("writing data/%s%s/%s.csv... 🚀\n", opts.Year, opts.Semester, ctx.facultyName)
		err = writeGradeDistibutionToCSV(ctx, result)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
	}

	fmt.Println("success ✅")
}
