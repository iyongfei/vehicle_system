package db

import (
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/db/redis"
	"vehicle_system/src/vehicle/db/tdata"
	"vehicle_system/src/vehicle/logger"
)

func Setup() {
	if gormDb, err := mysql.GetMysqlInstance().InitDataBase(); err != nil {
		logger.Logger.Error("gorm db connect fail:%+v", gormDb)
		logger.Logger.Print("gorm db connect fail:%+v", gormDb)
	}

	if redisDb, err := redis.GetRedisInstance().InitDataBase(); err != nil {
		logger.Logger.Error("redis db connect fail:%v", redisDb)
		logger.Logger.Print("redis db connect fail:%v", redisDb)
	}

	//检测表
	tdata.TableCheck()

	//检测表数据
	err := tdata.TdataCheck()
	if err != nil {
		logger.Logger.Error("tdata group check err:%v", err.Error())
		logger.Logger.Print("tdata group check err:%v", err.Error())
	}

	//检测设备
	err = tdata.VehicleAssetCheck()
	if err != nil {
		logger.Logger.Error("tdata vehicle_asset check err:%v", err.Error())
		logger.Logger.Print("tdata vehicle_asset check err:%v", err.Error())
	}
}
