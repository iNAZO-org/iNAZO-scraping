package main

import (
	"fmt"
)

type ScrapingCommand struct {
	Positional struct {
		Year      int
		Semester  int
		FacultyID string
	} `positional-args:"yes" required:"yes"`
}

var scrapingCommand ScrapingCommand

func (cmd *ScrapingCommand) Execute(args []string) error {
	var facultyIDList []string // allに対応する

	if cmd.Positional.FacultyID == "all" {
		facultyIDList = getKeysFromMap(FACULTY_ID_TO_NAME)
	} else {
		facultyIDList = []string{cmd.Positional.FacultyID}
	}

	for _, faclutyID := range facultyIDList {
		ctx := &ScrapingContext{
			year:        cmd.Positional.Year,
			semester:    cmd.Positional.Semester,
			facultyID:   faclutyID,
			facultyName: FACULTY_ID_TO_NAME[faclutyID],
		}

		fmt.Printf("scraping %d年%d学期 %s... 🚀\n", cmd.Positional.Year, cmd.Positional.Semester, ctx.facultyName)
		result, err := scrapingGradeDistribution(ctx)
		if err != nil {
			return err
		}

		fmt.Printf("writing data/%d%d/%s.csv... 🚀\n", cmd.Positional.Year, cmd.Positional.Semester, ctx.facultyName)
		err = writeGradeDistributionToCSV(ctx, result)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	const description = "Scraping from North University's grade distribution site and saving as a CSV file."
	parser.AddCommand("scraping", description, description, &scrapingCommand)
}
