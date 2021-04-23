package cron

import (
	"NaverWebtoonAlram/postgresql/model"
	"NaverWebtoonAlram/postgresql/query"
	"NaverWebtoonAlram/slack"
	"fmt"
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
			}else{
				if len(EndWebtoon)!=0{
				for i := range EndWebtoon{
					slack.SendSlackEndWebtoon(EndWebtoon[i], "일요일")
				}
				err :=query.Delete_Endtoon_sunday(EndWebtoon)
				if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.SUNDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoon(NewWebtoon[i],"일요일")
						var data model.SUNDAY_DB
						data.Toon=NewWebtoon[i]
						dataset = append(dataset, data)
					}
					result, err:=query.Insert_sunday(dataset)
					if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
					if result!=nil {log.Println("new webtoon data insert success")}
				}
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
			}else{
				if len(EndWebtoon)!=0{
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoon(EndWebtoon[i], "월요일")
					}
					err :=query.Delete_Endtoon_monday(EndWebtoon)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.MONDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoon(NewWebtoon[i], "월요일")
						var data model.MONDAY_DB
						data.Toon=NewWebtoon[i]
						dataset = append(dataset, data)
					}
					result, err:=query.Insert_monday(dataset)
					if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
					if result!=nil {log.Println("new webtoon data insert success")}
				}
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
			}else{
				if len(EndWebtoon)!=0{
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoon(EndWebtoon[i], "화요일")
					}
					err :=query.Delete_Endtoon_tuesday(EndWebtoon)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.TUESDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoon(NewWebtoon[i], "화요일")
						var data model.TUESDAY_DB
						data.Toon=NewWebtoon[i]
						dataset = append(dataset, data)
					}
					result, err:=query.Insert_tuesday(dataset)
					if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
					if result!=nil {log.Println("new webtoon data insert success")}
				}
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
			}else{
				if len(EndWebtoon)!=0{
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoon(EndWebtoon[i], "수요일")
					}
					err :=query.Delete_Endtoon_wednesday(EndWebtoon)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.WEDNESDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoon(NewWebtoon[i], "수요일")
						var data model.WEDNESDAY_DB
						data.Toon=NewWebtoon[i]
						dataset = append(dataset, data)
					}
					result, err:=query.Insert_wednesday(dataset)
					if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
					if result!=nil {log.Println("new webtoon data insert success")}
				}
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
			}else{
				if len(EndWebtoon)!=0{
					err :=query.Delete_Endtoon_thursday(EndWebtoon)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoon(EndWebtoon[i], "목요일")
					}
					var dataset []model.THURSDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoon(NewWebtoon[i], "목요일")
						var data model.THURSDAY_DB
						data.Toon=NewWebtoon[i]
						dataset = append(dataset, data)
					}
					result, err:=query.Insert_thursday(dataset)
					if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
					if result!=nil {log.Println("new webtoon data insert success")}
				}
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
			}else{
				if len(EndWebtoon)!=0{
					fmt.Println(EndWebtoon)
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoon(EndWebtoon[i], "금요일")
					}
					err :=query.Delete_Endtoon_friday(EndWebtoon)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.FRIDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoon(NewWebtoon[i], "금요일")
						var data model.FRIDAY_DB
						data.Toon=NewWebtoon[i]
						dataset = append(dataset, data)
					}
					result, err:=query.Insert_friday(dataset)
					if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
					if result!=nil {log.Println("new webtoon data insert success")}
				}
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
			}else{
				if len(EndWebtoon)!=0{
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoon(EndWebtoon[i], "토요일")
					}
					err :=query.Delete_Endtoon_saturday(EndWebtoon)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.SATURDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoon(NewWebtoon[i], "토요일")
						var data model.SATURDAY_DB
						data.Toon=NewWebtoon[i]
						dataset = append(dataset, data)
					}
					result, err:=query.Insert_saturday(dataset)
					if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
					if result!=nil {log.Println("new webtoon data insert success")}
				}
			}
			break
		}
	}

}

func CompareData(new []string, existed []string) ([]string, []string) {
	var NewToon []string
	var EndToon []string
	for i := range new {
		var count =0
		for j := range existed {
			if new[i]==existed[j]{count++}
		}
		if count==0{
			NewToon = append(NewToon, new[i])
		}
	}

	for i := range existed {
		var count =0
		for j := range new {
			if existed[i]==new[j]{count++}
		}
		if count==0{
			EndToon = append(EndToon, existed[i])
		}
	}
	return NewToon, EndToon
}

