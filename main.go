package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
	filename := fmt.Sprintf("%s%s.csv", ctx.year, ctx.semester)
	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	w := csv.NewWriter(f)

	for _, record := range result {
		err := w.Write([]string{
			record.subject,
			record.subTitle,
			record.class,
			record.teacher,
			record.year,
			record.semester,
			record.faculty,
			strconv.Itoa(record.studentCount),
			strconv.FormatFloat(record.gpa, 'f', -1, 64),

			strconv.Itoa(record.apCount),
			strconv.Itoa(record.aCount),
			strconv.Itoa(record.amCount),
			strconv.Itoa(record.bpCount),
			strconv.Itoa(record.bCount),
			strconv.Itoa(record.bmCount),
			strconv.Itoa(record.cpCount),
			strconv.Itoa(record.cCount),
			strconv.Itoa(record.dCount),
			strconv.Itoa(record.dmCount),
			strconv.Itoa(record.fCount),
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)

			err := os.Remove(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}

			return
		}
	}
	w.Flush()
}
