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
			NewWebtoon, EndWebtoon:=CompareDataSunday(new_data, existed_data)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}else{
				if len(EndWebtoon)!=0{
					fmt.Println(EndWebtoon)
					var dataset []string
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoonImage(EndWebtoon[i].Toon,EndWebtoon[i].Image, "일요일")
						dataset = append(dataset, EndWebtoon[i].Toon)
					}
					err :=query.Delete_Endtoon_sunday(dataset)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.SUNDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoonImage(NewWebtoon[i].Toon, NewWebtoon[i].Image, "일요일")
						dataset = append(dataset, NewWebtoon[i])
					}
					result, err:=query.Insert_sunday(dataset)
					if err!=nil {log.Fatalf("Can't write data in DB, %v", err)}
					if result!=nil {log.Println("new webtoon data insert success")}
				}
			}
			break
		}
		case 1: {
			var existed_data []model.MONDAY_DB
			err := query.Read_monday(&existed_data)
			if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}
			new_data:=ReadMonday()
			NewWebtoon, EndWebtoon:=CompareDataMonday(new_data, existed_data)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}else{
				if len(EndWebtoon)!=0{
					fmt.Println(EndWebtoon)
					var dataset []string
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoonImage(EndWebtoon[i].Toon, EndWebtoon[i].Image, "월요일")
						dataset = append(dataset, EndWebtoon[i].Toon)
					}
					err :=query.Delete_Endtoon_monday(dataset)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.MONDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoonImage(NewWebtoon[i].Toon, NewWebtoon[i].Image,"월요일")
						dataset = append(dataset, NewWebtoon[i])
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
			NewWebtoon, EndWebtoon:=CompareDataTuesday(new_data, existed_data)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}else{
				if len(EndWebtoon)!=0{
					fmt.Println(EndWebtoon)
					var dataset []string
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoonImage(EndWebtoon[i].Toon,EndWebtoon[i].Image, "화요일")
						dataset = append(dataset, EndWebtoon[i].Toon)
					}
					err :=query.Delete_Endtoon_tuesday(dataset)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.TUESDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoonImage(NewWebtoon[i].Toon,NewWebtoon[i].Image, "화요일")
						dataset = append(dataset, NewWebtoon[i])
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
			NewWebtoon, EndWebtoon:=CompareDataWednesday(new_data, existed_data)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}else{
				if len(EndWebtoon)!=0{
					var dataset []string
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoonImage(EndWebtoon[i].Toon, EndWebtoon[i].Image,"수요일")
						dataset = append(dataset, EndWebtoon[i].Toon)
					}
					err :=query.Delete_Endtoon_wednesday(dataset)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.WEDNESDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoonImage(NewWebtoon[i].Toon,NewWebtoon[i].Image, "수요일")
						dataset = append(dataset, NewWebtoon[i])
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
			NewWebtoon, EndWebtoon:=CompareDataThursday(new_data, existed_data)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}else{
				if len(EndWebtoon)!=0{
					var dataset []string
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoonImage(EndWebtoon[i].Toon, EndWebtoon[i].Image, "목요일")
						dataset = append(dataset, EndWebtoon[i].Toon)
					}
					err :=query.Delete_Endtoon_thursday(dataset)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.THURSDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoonImage(NewWebtoon[i].Toon, NewWebtoon[i].Image,"목요일")
						dataset = append(dataset, NewWebtoon[i])
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
			NewWebtoon, EndWebtoon:=CompareDataFirday(new_data, existed_data)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}else{
				if len(EndWebtoon)!=0{
					fmt.Println(EndWebtoon)
					var dataset []string
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoonImage(EndWebtoon[i].Toon, EndWebtoon[i].Image,"금요일")
						dataset = append(dataset, EndWebtoon[i].Toon)
					}
					err :=query.Delete_Endtoon_friday(dataset)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.FRIDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoonImage(NewWebtoon[i].Toon, NewWebtoon[i].Image,"금요일")
						dataset = append(dataset, NewWebtoon[i])
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
			NewWebtoon, EndWebtoon:=CompareDataSatureday(new_data, existed_data)
			if (len(NewWebtoon)==0 && len(EndWebtoon)==0){
				log.Println("Nothing changed")
			}else{
				if len(EndWebtoon)!=0{
					fmt.Println(EndWebtoon)
					var dataset []string
					for i := range EndWebtoon{
						slack.SendSlackEndWebtoonImage(EndWebtoon[i].Toon, EndWebtoon[i].Image,"토요일")
						dataset = append(dataset, EndWebtoon[i].Toon)
					}
					err :=query.Delete_Endtoon_saturday(dataset)
					if err != nil {{log.Fatalf("Can't read data from DB, %v", err)}}else{log.Println("Success to delete endwebtoon data")}
				}
				if len(NewWebtoon)!=0{
					var dataset []model.SATURDAY_DB
					for i := range NewWebtoon {
						slack.SendSlackStartWebtoonImage(NewWebtoon[i].Toon, NewWebtoon[i].Image,"토요일")
						dataset = append(dataset, NewWebtoon[i])
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




