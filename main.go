package main

import (
	"database/sql"
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
)

var parser = flags.NewParser(&struct{}{}, flags.Default)
var db *sql.DB

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func main() {
	connStr := os.Getenv("DB_URL")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	defer db.Close()

	_, err = parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Println("success ✅")
}
