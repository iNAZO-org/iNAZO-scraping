package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/sclevine/agouti"
)

func readScriptFile() (string, error) {
	file, err := os.Open("script.js")
	if err != nil {
		return "", err
	}
	defer file.Close()

	script, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(script), nil
}

func validateGradeDistribution(gd *GradeDistribution) error {
	sumStudentNumber := (gd.ApCount + gd.ACount + gd.AmCount +
		gd.BpCount + gd.BCount + gd.BmCount +
		gd.CpCount + gd.CCount +
		gd.DCount + gd.DmCount +
		gd.FCount)
	if gd.StudentCount != sumStudentNumber {
		return fmt.Errorf("grade distribution validation error:\n %+v", gd)
	}

	return nil
}

func searchGradeDistribution(ctx *ScrapingContext) error {
	// 検索画面へ移動
	page := ctx.page
	page.Navigate(searchUrl)

	// 検索条件の入力
	selectItems := []SelectItem{
		{"ddlTerm", strconv.Itoa(ctx.year) + strconv.Itoa(ctx.semester)}, // 年度・学期
		{"ddlDiv", "02"},          // 課程
		{"ddlFac", ctx.facultyID}, // 学部
		{"ddlDataKind", "1"},      // データ種別
	}

	for _, item := range selectItems {
		selectId := item.id
		selectValue := item.value
		optionXPath := fmt.Sprintf(`./option[@value="%s"]`, selectValue)
		err := page.FindByID(selectId).FindByXPath(optionXPath).Click()
		if err != nil {
			return err
		}
	}

	// 検索ボタンを押下
	err := page.FindByID("btnSerch").Click()
	if err != nil {
		return err
	}

	return nil
}

func viewAllGradeDistribution(ctx *ScrapingContext) error {
	// 表示件数を全てにする
	page := ctx.page
	selectId := "ddlLine_ddl"
	optionXPath := "./option[position()=1]"
	err := page.FindByID(selectId).FindByXPath(optionXPath).Click()
	if err != nil {
		return err
	}
	return nil
}

func fetchGradeDistribution(ctx *ScrapingContext) ([]GradeDistribution, error) {
	var result []GradeDistribution = make([]GradeDistribution, 0)
	page := ctx.page

	script, err := readScriptFile()
	if err != nil {
		return nil, err
	}

	trs := page.AllByXPath(`//*[@id="gvResult"]/tbody/tr[count(td)=18]`)
	n, err := trs.Count()
	if err != nil {
		return nil, err
	}
	validTrsCount := n
	bar := pb.StartNew(validTrsCount)

	for i := 0; i < n; i++ {
		/*
			rowItem: [
				subject, subTitle, class, teacher, studentCount,
				ap(%), a(%), bp(%), am(%), b(%), bm(%),
				cp(%), c(%), d(%), dm(%), f(%),
				gpa
			]
		*/
		statisticsRowWords := []string{
			"合計", "統計", "総計",
		}

		/*
			各tdをText()で取得するとWebDriverとの通信がボトルネックになるので、
			JavaScriptで一度に要素を取得する。
			:===:で区切る。
		*/
		var scriptResult string
		err := page.RunScript(
			script,
			map[string]interface{}{"pos": i + 1},
			&scriptResult,
		)
		if err != nil {
			return nil, err
		}
		rowItem := strings.Split(scriptResult, ":---:")[1:]

		if rowItem[1] == statisticsRowWords[0] ||
			rowItem[1] == statisticsRowWords[1] ||
			rowItem[1] == statisticsRowWords[2] {
			bar.Increment()
			continue
		}

		studentCount, err := strconv.Atoi(rowItem[4])
		if err != nil {
			bar.Finish()
			return nil, err
		}
		apPercent, err := strconv.ParseFloat(rowItem[5], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		aPercent, err := strconv.ParseFloat(rowItem[6], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		amPercent, err := strconv.ParseFloat(rowItem[7], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		bpPercent, err := strconv.ParseFloat(rowItem[8], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		bPercent, err := strconv.ParseFloat(rowItem[9], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		bmPercent, err := strconv.ParseFloat(rowItem[10], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		cpPercent, err := strconv.ParseFloat(rowItem[11], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		cPercent, err := strconv.ParseFloat(rowItem[12], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		dPercent, err := strconv.ParseFloat(rowItem[13], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		dmPercent, err := strconv.ParseFloat(rowItem[14], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		fPercent, err := strconv.ParseFloat(rowItem[15], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}
		gpa, err := strconv.ParseFloat(rowItem[16], 64)
		if err != nil {
			bar.Finish()
			return nil, err
		}

		gd := GradeDistribution{
			Subject:      rowItem[0],
			SubTitle:     rowItem[1],
			Class:        rowItem[2],
			Teacher:      rowItem[3],
			StudentCount: studentCount,
			Gpa:          gpa,
			Year:         ctx.year,
			Semester:     ctx.semester,
			Faculty:      ctx.facultyName,

			ApCount: int(math.Round(apPercent * float64(studentCount) / 100)),
			ACount:  int(math.Round(aPercent * float64(studentCount) / 100)),
			AmCount: int(math.Round(amPercent * float64(studentCount) / 100)),
			BpCount: int(math.Round(bpPercent * float64(studentCount) / 100)),
			BCount:  int(math.Round(bPercent * float64(studentCount) / 100)),
			BmCount: int(math.Round(bmPercent * float64(studentCount) / 100)),
			CpCount: int(math.Round(cpPercent * float64(studentCount) / 100)),
			CCount:  int(math.Round(cPercent * float64(studentCount) / 100)),
			DCount:  int(math.Round(dPercent * float64(studentCount) / 100)),
			DmCount: int(math.Round(dmPercent * float64(studentCount) / 100)),
			FCount:  int(math.Round(fPercent * float64(studentCount) / 100)),
		}
		if err := validateGradeDistribution(&gd); err != nil {
			bar.Finish()
			return nil, err
		}
		result = append(result, gd)
		bar.Increment()
	}
	bar.Finish()
	return result, nil
}

func scrapingGradeDistribution(ctx *ScrapingContext) ([]GradeDistribution, error) {
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--disable-gpu",
		})

	driver := agouti.ChromeDriver(options)
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		return nil, err
	}

	page, err := driver.NewPage()
	page.SetImplicitWait(10)
	if err != nil {
		return nil, err
	}

	ctx.driver = driver
	ctx.page = page

	err = searchGradeDistribution(ctx)
	if err != nil {
		return nil, err
	}

	err = viewAllGradeDistribution(ctx)
	if err != nil {
		return nil, err
	}

	result, err := fetchGradeDistribution(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
