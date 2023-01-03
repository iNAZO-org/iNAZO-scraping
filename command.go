package main

import "karintou8710/iNAZO-scraping/controllers"

type ScrapingCommand struct {
	Positional struct {
		Year      int
		Semester  int
		FacultyID string
	} `positional-args:"yes" required:"yes"`
}
type UpdateGradeCommand struct{}

var updateGradeCommand UpdateGradeCommand
var scrapingCommand ScrapingCommand

func (cmd *ScrapingCommand) Execute(args []string) error {
	return controllers.ScrapingHandler(cmd.Positional.Year, cmd.Positional.Semester, cmd.Positional.FacultyID)
}

func (cmd *UpdateGradeCommand) Execute(args []string) error {
	return controllers.UpdateGradeHandler()
}

func init() {
	const scrapingDescription = "Scraping from North University's grade distribution site and saving as a CSV file."
	parser.AddCommand("scraping", scrapingDescription, scrapingDescription, &scrapingCommand)

	const updateGradeDescription = "Updating the grade distribution table in the database with CSV files."
	parser.AddCommand("updateGrade", updateGradeDescription, updateGradeDescription, &updateGradeCommand)
}
