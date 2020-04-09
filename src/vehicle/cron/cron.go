package cron

import "github.com/robfig/cron/v3"

const (
	 PerMinuteSpec = "* * * * *"
)

func Setup()  {
	c := cron.New()
	c.AddFunc(PerMinuteSpec, perMinuteFun)
	c.Start()
}
