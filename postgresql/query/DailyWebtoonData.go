package query

import (
	"NaverWebtoonAlram/config"
	"NaverWebtoonAlram/postgresql/model"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func Read_monday(models *[]model.MONDAY_DB) error {
	err := config.DB.Model(models).Select()
	return err
}

func Insert_monday(models []model.MONDAY_DB) (orm.Result, error) {
	result, err := config.DB.Model(&models).Insert()
	return result, err
}

func Read_tuesday(models *[]model.TUESDAY_DB) error {
	err := config.DB.Model(models).Select()
	return err
}

func Insert_tuesday(models []model.TUESDAY_DB) (orm.Result, error) {
	result, err := config.DB.Model(&models).Insert()
	return result, err
}



func Read_wednesday(models *[]model.WEDNESDAY_DB) error {
	err := config.DB.Model(models).Select()
	return err
}

func Insert_wednesday(models []model.WEDNESDAY_DB) (orm.Result, error) {
	result, err := config.DB.Model(&models).Insert()
	return result, err
}



func Read_thursday(models *[]model.THURSDAY_DB) error {
	err := config.DB.Model(models).Select()
	return err
}

func Insert_thursday(models []model.THURSDAY_DB) (orm.Result, error) {
	result, err := config.DB.Model(&models).Insert()
	return result, err
}



func Read_friday(models *[]model.FRIDAY_DB) error {
	err := config.DB.Model(models).Select()
	return err
}



func Insert_friday(models []model.FRIDAY_DB) (orm.Result, error) {
	result, err := config.DB.Model(&models).Insert()
	return result, err
}



func Read_saturday(models *[]model.SATURDAY_DB) error {
	err := config.DB.Model(models).Select()
	return err
}

func Insert_saturday(models []model.SATURDAY_DB) (orm.Result, error) {
	result, err := config.DB.Model(&models).Insert()
	return result, err
}



func Read_sunday(models *[]model.SUNDAY_DB) error {
	err := config.DB.Model(models).Select()
	return err
}

func Insert_sunday(models []model.SUNDAY_DB) (orm.Result, error) {
	result, err := config.DB.Model(&models).Insert()
	return result, err
}
func Delete_Endtoon_monday(endWebtoon []string) error{
	ids := pg.In(endWebtoon)
	_, err :=  config.DB.Model((*model.MONDAY_DB)(nil)).Where("toon IN (?)",ids).Delete()
	return err
}
func Delete_Endtoon_tuesday(endWebtoon []string) error{
	ids := pg.In(endWebtoon)
	_, err :=  config.DB.Model((*model.TUESDAY_DB)(nil)).Where("toon IN (?)",ids).Delete()
	return err
}
func Delete_Endtoon_wednesday(endWebtoon []string) error{
	ids := pg.In(endWebtoon)
	_, err :=  config.DB.Model((*model.WEDNESDAY_DB)(nil)).Where("toon IN (?)",ids).Delete()
	return err
}
func Delete_Endtoon_thursday(endWebtoon []string) error{
	ids := pg.In(endWebtoon)
	_, err :=  config.DB.Model((*model.THURSDAY_DB)(nil)).Where("toon IN (?)",ids).Delete()
	return err
}
func Delete_Endtoon_friday(endWebtoon []string) error{
	ids := pg.In(endWebtoon)
	_, err :=  config.DB.Model((*model.FRIDAY_DB)(nil)).Where("toon IN (?)",ids).Delete()
	return err
}
func Delete_Endtoon_saturday(endWebtoon []string) error{
	ids := pg.In(endWebtoon)
	_, err :=  config.DB.Model((*model.SATURDAY_DB)(nil)).Where("toon IN (?)",ids).Delete()
	return err
}
func Delete_Endtoon_sunday(endWebtoon []string) error{
	ids := pg.In(endWebtoon)
	_, err :=  config.DB.Model((*model.SUNDAY_DB)(nil)).Where("toon IN (?)",ids).Delete()
	return err
}