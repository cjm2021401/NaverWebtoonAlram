package cron

import (
	"github.com/robfig/cron/v3"
	"time"
)

func SetupCron(location string) (*cron.Cron, error) {

	loc, err := time.LoadLocation(location)
	if err != nil {
		return nil, err
	}
	c := cron.New(cron.WithLocation(loc))
	_,err=c.AddFunc("0 0 1 1 *",  ReadWriteAllday)
	if err != nil {
		return nil, err
	}
	_,err=c.AddFunc("0/1 * * * *",  DailydataCheck)
	if err != nil {
		return nil, err
	}

	return c, nil
}