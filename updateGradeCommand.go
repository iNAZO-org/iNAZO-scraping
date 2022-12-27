package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type UpdateGradeCommand struct{}

var updateGradeCommand UpdateGradeCommand

func (cmd *UpdateGradeCommand) Execute(args []string) error {
	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return nil
}

func init() {
	const description = "Updating the grade distribution table in the database with CSV files."
	parser.AddCommand("updateGrade", description, description, &updateGradeCommand)
}
