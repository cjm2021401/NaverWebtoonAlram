package cron

import "NaverWebtoonAlram/postgresql/model"
func CompareDataMonday(new []model.MONDAY_DB, existed []model.MONDAY_DB) ([]model.MONDAY_DB, []model.MONDAY_DB) {
	var NewToon []model.MONDAY_DB
	var EndToon []model.MONDAY_DB
	for i := range new {
		var count =0
		for j := range existed {
			if new[i].Toon==existed[j].Toon{count++}
		}
		if count==0{
			NewToon = append(NewToon, new[i])
		}
	}

	for i := range existed {
		var count =0
		for j := range new {
			if existed[i].Toon==new[j].Toon{count++}
		}
		if count==0{
			EndToon = append(EndToon, existed[i])
		}
	}
	return NewToon, EndToon
}
func CompareDataTuesday(new []model.TUESDAY_DB, existed []model.TUESDAY_DB) ([]model.TUESDAY_DB, []model.TUESDAY_DB) {
	var NewToon []model.TUESDAY_DB
	var EndToon []model.TUESDAY_DB
	for i := range new {
		var count =0
		for j := range existed {
			if new[i].Toon==existed[j].Toon{count++}
		}
		if count==0{
			NewToon = append(NewToon, new[i])
		}
	}

	for i := range existed {
		var count =0
		for j := range new {
			if existed[i].Toon==new[j].Toon{count++}
		}
		if count==0{
			EndToon = append(EndToon, existed[i])
		}
	}
	return NewToon, EndToon
}
func CompareDataWednesday(new []model.WEDNESDAY_DB, existed []model.WEDNESDAY_DB) ([]model.WEDNESDAY_DB, []model.WEDNESDAY_DB) {
	var NewToon []model.WEDNESDAY_DB
	var EndToon []model.WEDNESDAY_DB
	for i := range new {
		var count =0
		for j := range existed {
			if new[i].Toon==existed[j].Toon{count++}
		}
		if count==0{
			NewToon = append(NewToon, new[i])
		}
	}

	for i := range existed {
		var count =0
		for j := range new {
			if existed[i].Toon==new[j].Toon{count++}
		}
		if count==0{
			EndToon = append(EndToon, existed[i])
		}
	}
	return NewToon, EndToon
}
func CompareDataThursday(new []model.THURSDAY_DB, existed []model.THURSDAY_DB) ([]model.THURSDAY_DB, []model.THURSDAY_DB) {
	var NewToon []model.THURSDAY_DB
	var EndToon []model.THURSDAY_DB
	for i := range new {
		var count =0
		for j := range existed {
			if new[i].Toon==existed[j].Toon{count++}
		}
		if count==0{
			NewToon = append(NewToon, new[i])
		}
	}

	for i := range existed {
		var count =0
		for j := range new {
			if existed[i].Toon==new[j].Toon{count++}
		}
		if count==0{
			EndToon = append(EndToon, existed[i])
		}
	}
	return NewToon, EndToon
}
func CompareDataFirday(new []model.FRIDAY_DB, existed []model.FRIDAY_DB) ([]model.FRIDAY_DB, []model.FRIDAY_DB) {
	var NewToon []model.FRIDAY_DB
	var EndToon []model.FRIDAY_DB
	for i := range new {
		var count =0
		for j := range existed {
			if new[i].Toon==existed[j].Toon{count++}
		}
		if count==0{
			NewToon = append(NewToon, new[i])
		}
	}

	for i := range existed {
		var count =0
		for j := range new {
			if existed[i].Toon==new[j].Toon{count++}
		}
		if count==0{
			EndToon = append(EndToon, existed[i])
		}
	}
	return NewToon, EndToon
}
func CompareDataSatureday(new []model.SATURDAY_DB, existed []model.SATURDAY_DB) ([]model.SATURDAY_DB, []model.SATURDAY_DB) {
	var NewToon []model.SATURDAY_DB
	var EndToon []model.SATURDAY_DB
	for i := range new {
		var count =0
		for j := range existed {
			if new[i].Toon==existed[j].Toon{count++}
		}
		if count==0{
			NewToon = append(NewToon, new[i])
		}
	}

	for i := range existed {
		var count =0
		for j := range new {
			if existed[i].Toon==new[j].Toon{count++}
		}
		if count==0{
			EndToon = append(EndToon, existed[i])
		}
	}
	return NewToon, EndToon
}
func CompareDataSunday(new []model.SUNDAY_DB, existed []model.SUNDAY_DB) ([]model.SUNDAY_DB, []model.SUNDAY_DB) {
	var NewToon []model.SUNDAY_DB
	var EndToon []model.SUNDAY_DB
	for i := range new {
		var count =0
		for j := range existed {
			if new[i].Toon==existed[j].Toon{count++}
		}
		if count==0{
			NewToon = append(NewToon, new[i])
		}
	}

	for i := range existed {
		var count =0
		for j := range new {
			if existed[i].Toon==new[j].Toon{count++}
		}
		if count==0{
			EndToon = append(EndToon, existed[i])
		}
	}
	return NewToon, EndToon
}