package main

const (
	SEARCH_URL = "https://educate.academic.hokudai.ac.jp/seiseki/GradeDistSerch.aspx"
	RESULT_URL = "https://educate.academic.hokudai.ac.jp/seiseki/GradeDistResult11.aspx"
)

var FACULTY_ID_TO_NAME = map[string]string{
	"00": "全学教育",
	"02": "総合教育部",
	"05": "文学部",
	"07": "教育学部",
	"11": "現代日本学プログラム課程",
	"15": "法学部",
	"17": "経済学部",
	"22": "理学部",
	"25": "工学部",
	"34": "農学部",
	"36": "獣医学部",
	"38": "水産学部",
	"42": "医学部",
	"43": "歯学部",
	"44": "薬学部",
}
