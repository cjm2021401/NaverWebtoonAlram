package cron

import (
	"NaverWebtoonAlram/config"
	"NaverWebtoonAlram/postgresql/model"
	"NaverWebtoonAlram/postgresql/query"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func ReadWebSite(dayofweek string)(*http.Response , error){

	resp, err := http.Get(config.Env.NaverWebtoon.Url+dayofweek)
	if err != nil {
		panic(err)
	}
	return resp, err
}

func ReadMonday()([]model.MONDAY_DB) {
	strVar1, err:=ReadWebSite("mon")
	if err!=nil{log.Fatalf("Can't run chromedp, %v", err)}
	mon_data := make([]model.MONDAY_DB, 0, 0)
	doc, _ := goquery.NewDocumentFromReader(strVar1.Body)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		var data model.MONDAY_DB
		width, ok := s.Attr("width")
		height,ok1 :=s.Attr("height")
		if ok &&ok1 {
			if width == "83"&& height=="90" {
				webtoonName, pass := s.Attr("alt")
				if pass{
						data.Toon = webtoonName
						mon_data=append(mon_data, data)
				}
			}
		}
	})
	return mon_data
}

func ReadWriteMonday(){
	mon_data:=ReadMonday()
	result, err:=query.Insert_monday(mon_data)
	if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
	if result!=nil {log.Println("monday webtoon data insert success")}

}

func ReadTuesday()([]model.TUESDAY_DB){
	strVar1, err:=ReadWebSite("tue")
	if err!=nil{log.Fatalf("Can't run chromedp, %v", err)}
	tue_data := make([]model.TUESDAY_DB, 0, 0)
	doc, _ := goquery.NewDocumentFromReader(strVar1.Body)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		var data model.TUESDAY_DB
		width, ok := s.Attr("width")
		height,ok1 :=s.Attr("height")
		if ok &&ok1 {
			if width == "83"&& height=="90" {
				webtoonName, pass := s.Attr("alt")
				if pass{
						data.Toon = webtoonName
						tue_data=append(tue_data, data)
				}
			}
		}
	})
	return tue_data
}

func ReadWriteTuesday(){
	tue_data:=ReadTuesday()
	result, err:=query.Insert_tuesday(tue_data)
	if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
	if result!=nil {log.Println("tuesday webtoon data insert success")}

}

func ReadWednesday()([]model.WEDNESDAY_DB){
	strVar1, err:=ReadWebSite("wed")
	if err!=nil{log.Fatalf("Can't run chromedp, %v", err)}
	wed_data := make([]model.WEDNESDAY_DB, 0, 0)
	doc, _ := goquery.NewDocumentFromReader(strVar1.Body)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		var data model.WEDNESDAY_DB
		width, ok := s.Attr("width")
		height,ok1 :=s.Attr("height")
		if ok &&ok1 {
			if width == "83" && height == "90" {
				webtoonName, pass := s.Attr("alt")
				if pass {
						data.Toon = webtoonName
						wed_data=append(wed_data, data)
				}
			}
		}
	})
	return wed_data
}

func ReadWriteWednesday(){
	wed_data:=ReadWednesday()
	result, err:=query.Insert_wednesday(wed_data)
	if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
	if result!=nil {log.Println("wednesday webtoon data insert success")}

}

func ReadThursday()([]model.THURSDAY_DB){
	strVar1, err:=ReadWebSite("thu")
	if err!=nil{log.Fatalf("Can't run chromedp, %v", err)}
	thu_data := make([]model.THURSDAY_DB, 0, 0)
	doc, _ := goquery.NewDocumentFromReader(strVar1.Body)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		var data model.THURSDAY_DB
		width, ok := s.Attr("width")
		height,ok1 :=s.Attr("height")
		if ok &&ok1 {
			if width == "83"&& height=="90" {
				webtoonName, pass := s.Attr("alt")
				if pass{
					if webtoonName!="" {
						data.Toon = webtoonName
						thu_data=append(thu_data, data)
					}
				}
			}
		}
	})
	return thu_data
}

func ReadWriteThursday(){
	thu_data:=ReadThursday()
	result, err:=query.Insert_thursday(thu_data)
	if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
	if result!=nil {log.Println("thursday webtoon data insert success")}
}


func ReadFriday()([]model.FRIDAY_DB){
	strVar1, err:=ReadWebSite("fri")
	if err!=nil{log.Fatalf("Can't run chromedp, %v", err)}
	fri_data := make([]model.FRIDAY_DB, 0, 0)
	doc, _ := goquery.NewDocumentFromReader(strVar1.Body)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		var data model.FRIDAY_DB
		width, ok := s.Attr("width")
		height,ok1 :=s.Attr("height")
		if ok &&ok1 {
			if width == "83"&& height=="90" {
				webtoonName, pass := s.Attr("alt")
				if pass{
					if webtoonName!="" {
						data.Toon = webtoonName
						fri_data=append(fri_data, data)
					}
				}
			}
		}
	})
	return fri_data
}

func ReadWriteFriday(){
	fri_data:=ReadFriday()
	result, err:=query.Insert_friday(fri_data)
	if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
	if result!=nil {log.Println("friday webtoon data insert success")}

}

func ReadSaturday()([]model.SATURDAY_DB){
	strVar1, err:=ReadWebSite("sat")
	if err!=nil{log.Fatalf("Can't run chromedp, %v", err)}
	sat_data := make([]model.SATURDAY_DB, 0, 0)
	doc, _ := goquery.NewDocumentFromReader(strVar1.Body)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		var data model.SATURDAY_DB
		width, ok := s.Attr("width")
		height,ok1 :=s.Attr("height")
		if ok &&ok1 {
			if width == "83"&& height=="90" {
				webtoonName, pass := s.Attr("alt")
				if pass{
					if webtoonName!="" {
						data.Toon = webtoonName
						sat_data=append(sat_data, data)
					}
				}
			}
		}
	})
	return ReadSaturday()
}

func ReadWriteSaturday(){
	sat_data:=ReadSaturday()
	result, err:=query.Insert_saturday(sat_data)
	if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
	if result!=nil {log.Println("saturday webtoon data insert success")}
}


func ReadSunday()([]model.SUNDAY_DB){
	strVar1, err:=ReadWebSite("sun")
	if err!=nil{log.Fatalf("Can't run chromedp, %v", err)}
	sun_data := make([]model.SUNDAY_DB, 0, 0)
	doc, _ := goquery.NewDocumentFromReader(strVar1.Body)
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		var data model.SUNDAY_DB
		width, ok := s.Attr("width")
		height,ok1 :=s.Attr("height")
		if ok &&ok1 {
			if width == "83" && height == "90" {
				webtoonName, pass := s.Attr("alt")
				if pass {
					if webtoonName != "" {
						data.Toon = webtoonName
						sun_data=append(sun_data, data)
					}
				}
			}
		}
	})
	return sun_data
}

func ReadWriteSunday(){
	sun_data:=ReadSunday()
	result, err:=query.Insert_sunday(sun_data)
	if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
	if result!=nil {log.Println("sunday webtoon data insert success")}

}

