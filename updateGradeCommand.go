package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	_ "github.com/lib/pq"
)

type UpdateGradeCommand struct{}

var updateGradeCommand UpdateGradeCommand

func existGradeDistributionRow(gd GradeDistributionItem) (bool, error) {
	var cnt int
	err := db.QueryRow(`
		SELECT COUNT(*) FROM grade_distribution
		WHERE
			subject=$1 AND sub_title=$2 AND class=$3 AND
			teacher=$4 AND year=$5 AND semester=$6 AND faculty=$7 AND student_count=$8 AND
			gpa=$9 AND ap_count=$10 AND a_count=$11 AND am_count=$12 AND bp_count=$13 AND
			b_count=$14 AND bm_count=$15 AND cp_count=$16 AND c_count=$17 AND d_count=$18 AND
			dm_count=$19 AND f_count=$20`,
		gd.subject, gd.subTitle, gd.class, gd.teacher, gd.year, gd.semester,
		gd.faculty, gd.studentCount, gd.gpa, gd.apCount, gd.aCount, gd.amCount,
		gd.bpCount, gd.bCount, gd.bmCount, gd.cpCount, gd.cCount, gd.dCount,
		gd.dmCount, gd.fCount,
	).Scan(&cnt)
	if err != nil {
		return false, err
	}
	existFlag := cnt > 0
	return existFlag, nil
}

func insertGradeDistributionList(gdList []GradeDistributionItem) error {
	for _, gd := range gdList {
		existFlag, err := existGradeDistributionRow(gd)
		if err != nil {
			return err
		}
		// 既に登録済みの項目は2重に保存しない
		if existFlag {
			continue
		}

		if _, err := db.Exec(`
		INSERT INTO grade_distribution(
			subject, sub_title, class, teacher, year,
			semester, faculty, student_count, gpa,
			ap_count, a_count, am_count,
			bp_count, b_count, bm_count,
			cp_count, c_count, d_count,
			dm_count, f_count
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`,
			gd.subject, gd.subTitle, gd.class, gd.teacher, gd.year, gd.semester,
			gd.faculty, gd.studentCount, gd.gpa, gd.apCount, gd.aCount, gd.amCount,
			gd.bpCount, gd.bCount, gd.bmCount, gd.cpCount, gd.cCount, gd.dCount,
			gd.dmCount, gd.fCount,
		); err != nil {
			return err
		}
	}

	return nil
}

func (cmd *UpdateGradeCommand) Execute(args []string) error {
	err := filepath.Walk("./data", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fmt.Printf("saving %s...\n", path)
		gdList, err := readGradeDistributionFromCSV(path)
		if err != nil {
			return err
		}
		err = insertGradeDistributionList(gdList)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func init() {
	const description = "Updating the grade distribution table in the database with CSV files."
	parser.AddCommand("updateGrade", description, description, &updateGradeCommand)
}
