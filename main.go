package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"

	"karintou8710/iNAZO-scraping/database"
	"karintou8710/iNAZO-scraping/models"
	"karintou8710/iNAZO-scraping/setting"
)

var parser = flags.NewParser(&struct{}{}, flags.Default)

func main() {
	setting.Init()
	database.Init(&models.GradeDistribution{})

	_, err := parser.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Println("success ✅")
}
