package timing

import (
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

func licenseCron() {
	logger.Logger.Print("%s license", util.RunFuncName())
	logger.Logger.Info("%s license", util.RunFuncName())

}
