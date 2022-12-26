package main

import (
	"github.com/sclevine/agouti"
)

type ScrapingContext struct {
	driver      *agouti.WebDriver
	page        *agouti.Page
	year        string
	semester    string
	facultyID   string
	facultyName string
}

type SelectItem struct {
	id    string
	value string
}

type GradeDistributionItem struct {
	subject      string
	subTitle     string
	class        string
	teacher      string
	year         string
	semester     string
	faculty      string
	studentCount int
	gpa          float64

	apCount int // A+の人数
	aCount  int // A
	amCount int // A-
	bpCount int // B+
	bCount  int // B
	bmCount int // B-
	cpCount int // C+
	cCount  int // C
	dCount  int // D
	dmCount int // D-
	fCount  int // F
}

type Options struct {
	Year      string `short:"y" long:"year" description:"Year of the grade you wish to obtain"`
	Semester  string `short:"s" long:"semester" description:"The semester of the grade you wish to obtain"`
	FacultyID string `short:"f" long:"facultyId" description:"Faculty ID of the grade you wish to obtain"`

	facultyIDList []string // allに対応するためのフィールド
}
