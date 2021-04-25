package cron

import (
	"github.com/robfig/cron/v3"
	"time"
)

func SetupCron(location string) (*cron.Cron, error) {
	//ReadWriteAllday()  //처음 세팅할때만 돌림
	loc, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	c := cron.New(cron.WithLocation(loc))
	_,err=c.AddFunc("* * 8 * * *",  DailydataCheck)//매일 8시에 확인
	if err != nil {
		return nil, err
	}

	return c, nil
}