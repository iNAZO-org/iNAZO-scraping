package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func writeGradeDistibutionToCSV(ctx *ScrapingContext, gd []GradeDistributionItem) error {
	filename := fmt.Sprintf("%s%s.csv", ctx.year, ctx.semester)
	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		return err
	}
	w := csv.NewWriter(f)

	for _, record := range gd {
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
			err := os.Remove(filename)
			if err != nil {
				return err
			}

			return err
		}
	}
	w.Flush()

	return nil
}
