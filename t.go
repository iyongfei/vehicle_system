package main

import (
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/db/redis"
)

func main()  {

	//vgorm := InitDataBase()

	//strategyVehicleLearningResultJoins := []*model.StrategyVehicleLearningResultJoin{}
	//_ = vgorm.Debug().
	//	Table("strategies").
	//	Select("strategies.*,strategy_vehicles.vehicle_id").
	//	Where("strategy_vehicles.vehicle_id = ?","TDavCZX6IyE3NDa2OVRaMlt92pMOG3Hw").
	//	Joins("inner join strategy_vehicles ON strategies.strategy_id = strategy_vehicles.strategy_id").
	//	Order("strategies.created_at desc").
	//	Scan(&strategyVehicleLearningResultJoins).
	//	Error
	//for k,v:=range strategyVehicleLearningResultJoins{
	//	fmt.Println(k,v.StrategyId)
	//}






	//re,_:=redis.GetRedisInstance()
	//for k,v:=range re{
	//	fmt.Println(k,v,reflect.TypeOf(v))
	//}

	//redis.GetRedisInstance().VHGet()

	//redis.GetRedisInstance().Vkeys()

	redis.GetRedisInstance().VSet("sdf",[]byte("sdfsd"),0)
}




func InitDataBase()(*gorm.DB){
	GormDb, _ := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/vehicle?charset=utf8&parseTime=True&loc=Local")
	//GormDb.LogMode(true)

	GormDb.DB().SetMaxIdleConns(conf.MaxIdleConns)
	GormDb.DB().SetMaxOpenConns(conf.MaxOpenConns)
	return GormDb
}