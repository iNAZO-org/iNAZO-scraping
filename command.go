package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

func parseCommand() (*Options, error) {
	var opts Options
	args, err := flags.Parse(&opts)
	if err != nil {
		return nil, err
	}

	if opts.Year == "" || opts.Semester == "" || opts.FacultyID == "" {
		if len(args) != 3 {
			err := fmt.Errorf("require three args\n")
			return nil, err
		}

		opts.Year = args[0]
		opts.Semester = args[1]
		opts.FacultyID = args[2]
	}

	return &opts, err
}
