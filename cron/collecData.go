package cron

import (
	"NaverWebtoonAlram/postgresql/model"
	"NaverWebtoonAlram/postgresql/query"
	"log"
	_ "log"
	"time"
)

func ReadWriteAllday(){
	ReadWriteMonday()
	ReadWriteTuesday()
	ReadWriteWednesday()
	ReadWriteThursday()
	ReadWriteFriday()
	ReadWriteSaturday()
	ReadWriteSunday()
}



func DailydataCheck() {
	today:=time.Now()
	yesterday:=time.Date(today.Year(),today.Month(),today.Day()-1,0,0,0,0,today.Location())
	switch int(yesterday.Weekday()) {
		case 0: {
			var existed_data []model.SUNDAY_DB
			err := query.Read_sunday(&existed_data)
			if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}
			new_data:=ReadSunday()

			var existed []string
			var new []string
			for i := range existed_data {
				existed = append(existed, existed_data[i].Toon)
			}
			for i := range new_data{
				new = append(new ,new_data[i].Toon)
			}
			NewWebtoon, EndWebtoon:=CompareData(new, existed)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}

			break
		}
		case 1:{
			var existed_data []model.MONDAY_DB
			err := query.Read_monday(&existed_data)
			if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}
			new_data:=ReadMonday()

			var existed []string
			var new []string
			for i := range existed_data {
				existed = append(existed, existed_data[i].Toon)
			}
			for i := range new_data{
				new = append(new ,new_data[i].Toon)
			}
			NewWebtoon, EndWebtoon:=CompareData(new, existed)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}

			break
		}
		case 2: {
			var existed_data []model.TUESDAY_DB
			err := query.Read_tuesday(&existed_data)
			if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}
			new_data:=ReadTuesday()

			var existed []string
			var new []string
			for i := range existed_data {
				existed = append(existed, existed_data[i].Toon)
			}
			for i := range new_data{
				new = append(new ,new_data[i].Toon)
			}
			NewWebtoon, EndWebtoon:=CompareData(new, existed)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}

			break
		}
		case 3: {
			var existed_data []model.WEDNESDAY_DB
			err := query.Read_wednesday(&existed_data)
			if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}
			new_data:=ReadWednesday()

			var existed []string
			var new []string
			for i := range existed_data {
				existed = append(existed, existed_data[i].Toon)
			}
			for i := range new_data{
				new = append(new ,new_data[i].Toon)
			}
			NewWebtoon, EndWebtoon:=CompareData(new, existed)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}

			break
		}
		case 4: {
			var existed_data []model.THURSDAY_DB
			err := query.Read_thursday(&existed_data)
			if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}
			new_data:=ReadThursday()

			var existed []string
			var new []string
			for i := range existed_data {
				existed = append(existed, existed_data[i].Toon)
			}
			for i := range new_data{
				new = append(new ,new_data[i].Toon)
			}
			NewWebtoon, EndWebtoon:=CompareData(new, existed)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}

			break
		}
		case 5: {
			var existed_data []model.FRIDAY_DB
			err := query.Read_friday(&existed_data)
			if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}
			new_data:=ReadFriday()

			var existed []string
			var new []string
			for i := range existed_data {
				existed = append(existed, existed_data[i].Toon)
			}
			for i := range new_data{
				new = append(new ,new_data[i].Toon)
			}
			NewWebtoon, EndWebtoon:=CompareData(new, existed)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}

			break
		}
		case 6: {
			var existed_data []model.SATURDAY_DB
			err := query.Read_saturday(&existed_data)
			if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}
			new_data:=ReadSaturday()

			var existed []string
			var new []string
			for i := range existed_data {
				existed = append(existed, existed_data[i].Toon)
			}
			for i := range new_data{
				new = append(new ,new_data[i].Toon)
			}
			NewWebtoon, EndWebtoon:=CompareData(new, existed)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}

			break
		}
	}

}

type New_EndToon struct {
	id int
	toon string
}

func CompareData(new []string, existed []string) ([]New_EndToon, []New_EndToon) {
	var NewToon []New_EndToon
	var EndToon []New_EndToon
	for i := range new {
		var count =0
		var New New_EndToon
		for j := range existed {
			if new[i]==existed[j]{count++}
		}
		if count==0{
			New.id=i
			New.toon=new[i]
			NewToon = append(NewToon, New)
		}
	}

	for i := range existed {
		var count =0
		var End New_EndToon
		for j := range new {
			if existed[i]==new[j]{count++}
		}
		if count==0{
			End.id=i
			End.toon=new[i]
			EndToon = append(EndToon, End)
		}
	}
	return NewToon, EndToon
}

