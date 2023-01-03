package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"path"
	"strconv"

	"karintou8710/iNAZO-scraping/models"
)

func deserializationGradeDistribution(row []string) (*models.GradeDistribution, error) {
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

	return &models.GradeDistribution{
		Subject:      row[0],
		SubTitle:     row[1],
		Class:        row[2],
		Teacher:      row[3],
		Year:         year,
		Semester:     semester,
		Faculty:      row[6],
		StudentCount: studentCount,
		Gpa:          gpa,

		ApCount: apCount,
		ACount:  aCount,
		AmCount: amCount,
		BpCount: bpCount,
		BCount:  bCount,
		BmCount: bmCount,
		CpCount: cpCount,
		CCount:  cCount,
		DCount:  dCount,
		DmCount: dmCount,
		FCount:  fCount,
	}, nil
}

func ReadGradeDistributionFromCSV(filePath string) ([]*models.GradeDistribution, error) {
	var result []*models.GradeDistribution
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
		result = append(result, gd)
	}

	return result, nil
}

func WriteGradeDistributionToCSV(year, semester int, facultyName string, gd []*models.GradeDistribution) error {
	filePath := fmt.Sprintf("data/%d%d/%s.csv", year, semester, facultyName)
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
			record.Subject,
			record.SubTitle,
			record.Class,
			record.Teacher,
			strconv.Itoa(record.Year),
			strconv.Itoa(record.Semester),
			record.Faculty,
			strconv.Itoa(record.StudentCount),
			strconv.FormatFloat(record.Gpa, 'f', -1, 64),

			strconv.Itoa(record.ApCount),
			strconv.Itoa(record.ACount),
			strconv.Itoa(record.AmCount),
			strconv.Itoa(record.BpCount),
			strconv.Itoa(record.BCount),
			strconv.Itoa(record.BmCount),
			strconv.Itoa(record.CpCount),
			strconv.Itoa(record.CCount),
			strconv.Itoa(record.DCount),
			strconv.Itoa(record.DmCount),
			strconv.Itoa(record.FCount),
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
