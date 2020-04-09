package cron

import (
	"time"
	"vehicle_system/src/vehicle/logger"
)

func perMinuteFun()  {

	logger.Logger.Print("PerMinuteFun %v",time.Now())

}