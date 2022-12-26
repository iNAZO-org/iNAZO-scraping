package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/sclevine/agouti"
)

const (
	SEARCH_URL = "https://educate.academic.hokudai.ac.jp/seiseki/GradeDistSerch.aspx"
	RESULT_URL = "https://educate.academic.hokudai.ac.jp/seiseki/GradeDistResult11.aspx"
)

func validateGradeDistribution(gd *GradeDistributionItem) bool {
	sumStudentNumber := (gd.apCount + gd.aCount + gd.amCount +
		gd.bpCount + gd.bCount + gd.bmCount +
		gd.cpCount + gd.cCount +
		gd.dCount + gd.dCount +
		gd.fCount)
	if gd.studentCount != sumStudentNumber {
		return false
	}

	return true
}

func searchGradeDistribution(ctx *ScrapingContext) error {
	// 検索画面へ移動
	page := ctx.page
	page.Navigate(SEARCH_URL)
	time.Sleep(time.Second * 3)

	// 検索条件の入力
	selectItems := []SelectItem{
		{"ddlTerm", ctx.year + ctx.semester}, // 年度・学期
		{"ddlDiv", "02"},                     // 課程
		{"ddlFac", ctx.facultyID},            // 学部
		{"ddlDataKind", "1"},                 // データ種別
	}

	for _, item := range selectItems {
		selectId := item.id
		selectValue := item.value
		optionXPath := fmt.Sprintf(`./option[@value="%s"]`, selectValue)
		err := page.FindByID(selectId).FindByXPath(optionXPath).Click()
		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}

	// 検索ボタンを押下
	err := page.FindByID("btnSerch").Click()
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 3)

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
	time.Sleep(time.Second * 3)
	return nil
}

func fetchGradeDistribution(ctx *ScrapingContext) ([]GradeDistributionItem, error) {
	var result []GradeDistributionItem = make([]GradeDistributionItem, 0)
	page := ctx.page

	trs := page.AllByXPath(`//*[@id="gvResult"]/tbody/tr`)
	n, err := trs.Count()
	if err != nil {
		return nil, err
	}
	validTrsCount := (n - 2) / 2
	bar := pb.StartNew(validTrsCount)

	// 初めの2個, 偶数番目のtrは成績データと無関係なので無視する
	for i := 2; i < n; i += 2 {
		/*
			rowItem: [
				subject, subTitle, class, teacher, studentCount,
				ap(%), a(%), bp(%), am(%), b(%), bm(%),
				cp(%), c(%), d(%), dm(%), f(%),
				gpa
			]
		*/
		rowItem := make([]string, 0)
		const lowerBound = 1
		const upperBound = 17
		statisticsRowWords := []string{
			"合計", "統計", "総計",
		}

		tds := trs.At(i).All("td")
		for j := lowerBound; j <= upperBound; j++ {
			text, err := tds.At(j).Text()
			if err != nil {
				bar.Finish()
				return nil, err
			}
			rowItem = append(rowItem, text)
		}

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

		gd := GradeDistributionItem{
			subject:      rowItem[0],
			subTitle:     rowItem[1],
			class:        rowItem[2],
			teacher:      rowItem[3],
			studentCount: studentCount,
			gpa:          gpa,
			year:         ctx.year,
			semester:     ctx.semester,
			faculty:      ctx.facultyName,

			apCount: int(math.Round(apPercent * float64(studentCount) / 100)),
			aCount:  int(math.Round(aPercent * float64(studentCount) / 100)),
			amCount: int(math.Round(amPercent * float64(studentCount) / 100)),
			bpCount: int(math.Round(bpPercent * float64(studentCount) / 100)),
			bCount:  int(math.Round(bPercent * float64(studentCount) / 100)),
			bmCount: int(math.Round(bmPercent * float64(studentCount) / 100)),
			cpCount: int(math.Round(cpPercent * float64(studentCount) / 100)),
			cCount:  int(math.Round(cPercent * float64(studentCount) / 100)),
			dCount:  int(math.Round(dPercent * float64(studentCount) / 100)),
			dmCount: int(math.Round(dmPercent * float64(studentCount) / 100)),
			fCount:  int(math.Round(fPercent * float64(studentCount) / 100)),
		}
		if !validateGradeDistribution(&gd) {
			bar.Finish()
			return nil, fmt.Errorf("grade distribution validation error")
		}
		result = append(result, gd)
		bar.Increment()
	}
	bar.Finish()
	return result, nil
}

func scrapingGradeDistribution(ctx *ScrapingContext) ([]GradeDistributionItem, error) {
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