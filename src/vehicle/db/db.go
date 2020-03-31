package db

import (
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/db/redis"
	"vehicle_system/src/vehicle/logger"
)

func init()  {
	if gormDb,err:= mysql.GetMysqlInstance().InitDataBase();err != nil{
		logger.Logger.Error("gorm db connect fail:%+v",gormDb)
		logger.Logger.Print("gorm db connect fail:%+v",gormDb)
	}

	if redisDb,err:= redis.GetRedisInstance().InitDataBase();err != nil{
		logger.Logger.Error("redis db connect fail:%v",redisDb)
		logger.Logger.Print("redis db connect fail:%v",redisDb)
	}
}
