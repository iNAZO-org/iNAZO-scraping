package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"
)

func writeGradeDistibutionToCSV(ctx *ScrapingContext, gd []GradeDistributionItem) error {
	filePath := fmt.Sprintf("data/%s%s/%s.csv", ctx.year, ctx.semester, ctx.facultyName)
	folderPath := path.Dir(filePath)
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.Create(filePath)
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
			err := os.Remove(filePath)
			if err != nil {
				return err
			}

			return err
		}
	}
	w.Flush()

	return nil
}
