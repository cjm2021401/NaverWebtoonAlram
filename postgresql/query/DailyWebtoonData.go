package query

import (
	"NaverWebtoonAlram/config"
	"NaverWebtoonAlram/postgresql/model"
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