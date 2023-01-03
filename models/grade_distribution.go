package models

import "gorm.io/gorm"

// gorm用
type GradeDistribution struct {
	gorm.Model

	Subject      string
	SubTitle     string
	Class        string
	Teacher      string
	Year         int
	Semester     int
	Faculty      string
	StudentCount int
	Gpa          float64

	ApCount int // A+の人数
	ACount  int // A
	AmCount int // A-
	BpCount int // B+
	BCount  int // B
	BmCount int // B-
	CpCount int // C+
	CCount  int // C
	DCount  int // D
	DmCount int // D-
	FCount  int // F
}
