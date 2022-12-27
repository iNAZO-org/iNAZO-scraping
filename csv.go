package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"
)

func deserializationGradeDistribution(row []string) (*GradeDistributionItem, error) {
	year, err := strconv.Atoi(row[4])
	if err != nil {
		return nil, err
	}
	semester, err := strconv.Atoi(row[5])
	if err != nil {
		return nil, err
	}
	studentCount, err := strconv.Atoi(row[7])
	if err != nil {
		return nil, err
	}
	gpa, err := strconv.ParseFloat(row[8], 64)
	if err != nil {
		return nil, err
	}
	apCount, err := strconv.Atoi(row[9])
	if err != nil {
		return nil, err
	}
	aCount, err := strconv.Atoi(row[10])
	if err != nil {
		return nil, err
	}
	amCount, err := strconv.Atoi(row[11])
	if err != nil {
		return nil, err
	}
	bpCount, err := strconv.Atoi(row[12])
	if err != nil {
		return nil, err
	}
	bCount, err := strconv.Atoi(row[13])
	if err != nil {
		return nil, err
	}
	bmCount, err := strconv.Atoi(row[14])
	if err != nil {
		return nil, err
	}
	cpCount, err := strconv.Atoi(row[15])
	if err != nil {
		return nil, err
	}
	cCount, err := strconv.Atoi(row[16])
	if err != nil {
		return nil, err
	}
	dCount, err := strconv.Atoi(row[17])
	if err != nil {
		return nil, err
	}
	dmCount, err := strconv.Atoi(row[18])
	if err != nil {
		return nil, err
	}
	fCount, err := strconv.Atoi(row[19])
	if err != nil {
		return nil, err
	}

	return &GradeDistributionItem{
		subject:      row[0],
		subTitle:     row[1],
		class:        row[2],
		teacher:      row[3],
		year:         year,
		semester:     semester,
		faculty:      row[6],
		studentCount: studentCount,
		gpa:          gpa,

		apCount: apCount,
		aCount:  aCount,
		amCount: amCount,
		bpCount: bpCount,
		bCount:  bCount,
		bmCount: bmCount,
		cpCount: cpCount,
		cCount:  cCount,
		dCount:  dCount,
		dmCount: dmCount,
		fCount:  fCount,
	}, nil
}

func readGradeDistributionFromCSV(filePath string) ([]GradeDistributionItem, error) {
	var result []GradeDistributionItem
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		gd, err := deserializationGradeDistribution(row)
		if err != nil {
			return nil, err
		}
		result = append(result, *gd)
	}

	return result, nil
}

func writeGradeDistributionToCSV(ctx *ScrapingContext, gd []GradeDistributionItem) error {
	filePath := fmt.Sprintf("data/%d%d/%s.csv", ctx.year, ctx.semester, ctx.facultyName)
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
			strconv.Itoa(record.year),
			strconv.Itoa(record.semester),
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
