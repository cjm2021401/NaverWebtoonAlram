package model

type MONDAY_DB struct {
	tableName   struct{} `pg:"webtoon_data.monday"`
	Id          int      `pg:"id"`
	Toon  string   `pg:"toon"`
}

type TUESDAY_DB struct {
	tableName   struct{} `pg:"webtoon_data.tuesday"`
	Id          int      `pg:"id"`
	Toon  string   `pg:"toon"`
}
type WEDNESDAY_DB struct {
	tableName   struct{} `pg:"webtoon_data.wednesday"`
	Id          int      `pg:"id"`
	Toon  string   `pg:"toon"`
}
type THURSDAY_DB struct {
	tableName   struct{} `pg:"webtoon_data.thursday"`
	Id          int      `pg:"id"`
	Toon  string   `pg:"toon"`
}
type FRIDAY_DB struct {
	tableName   struct{} `pg:"webtoon_data.friday"`
	Id          int      `pg:"id"`
	Toon  string   `pg:"toon"`
}
type SATURDAY_DB struct {
	tableName   struct{} `pg:"webtoon_data.saturday"`
	Id          int      `pg:"id"`
	Toon  string   `pg:"toon"`
}

type SUNDAY_DB struct {
	tableName   struct{} `pg:"webtoon_data.sunday"`
	Id          int      `pg:"id"`
	Toon  string   `pg:"toon"`
}