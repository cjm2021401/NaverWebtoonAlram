package main

import (
	"NaverWebtoonAlram/config"
	"NaverWebtoonAlram/cron"
	"NaverWebtoonAlram/gin"
	"github.com/go-pg/pg/v10"
	"log"
)

func main(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := config.GetEnvironmentVariable()
	if err != nil {
		log.Panic(err)
	}
	config.DB = pg.Connect(&pg.Options{
		Addr:     config.Env.Database.Url,
		User:     config.Env.Database.User,
		Password: config.Env.Database.Password,
		Database: config.Env.Database.Name,
	})
	defer config.DB.Close()
	c, err := cron.SetupCron("Asia/Seoul")
	if err != nil {
		log.Println(err)
	}
	c.Start()

	r := gin.SetupRouter()
	err = r.Run()
	if err != nil {
		log.Panic(err)

	}
}
