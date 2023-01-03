package main

import (
	"github.com/sclevine/agouti"
)

type ScrapingContext struct {
	driver      *agouti.WebDriver
	page        *agouti.Page
	year        int
	semester    int
	facultyID   string
	facultyName string
}

type SelectItem struct {
	id    string
	value string
}
