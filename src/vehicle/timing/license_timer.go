package timing

import (
	"time"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

func licenseCron() {
	logger.Logger.Print("%s licenseCron:%v", util.RunFuncName(), time.Now())
	logger.Logger.Info("%s licenseCron:%v", util.RunFuncName(), time.Now())

}
