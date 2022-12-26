package main

type UpdateGradeCommand struct{}

var updateGradeCommand UpdateGradeCommand

func (cmd *UpdateGradeCommand) Execute(args []string) error {
	return nil
}

func init() {
	const description = "Updating the grade distribution table in the database with CSV files."
	parser.AddCommand("updateGrade", description, description, &updateGradeCommand)
}
