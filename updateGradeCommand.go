package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type UpdateGradeCommand struct{}

var updateGradeCommand UpdateGradeCommand

func (cmd *UpdateGradeCommand) Execute(args []string) error {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	gradeList, err := readGradeDistributionFromCSV("data/20221/総合教育部.csv")
	if err != nil {
		return err
	}
	fmt.Println(gradeList)

	return nil
}

func init() {
	const description = "Updating the grade distribution table in the database with CSV files."
	parser.AddCommand("updateGrade", description, description, &updateGradeCommand)
}
