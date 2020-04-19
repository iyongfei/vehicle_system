package tdata

import (
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/model"
)

/**
检测数据库表
 */

func TableCheck()  {
	tables := getTables()

	for _,table := range tables{
		err,_ := checkTable(table)
		if err!=nil{
			err,_  = checkTable(table)
			continue
		}
	}
}

func checkTable(model interface{}) (error,int64) {
	err,rowsAffected:= mysql.CreatTable(model)
	return err,rowsAffected
}


func getTables() []interface{} {
	//添加各个&model表
	flow:= &model.Flow{}
	firmwareUpdate:= &model.FirmwareUpdate{}
	firmwareInfo:=&model.FirmwareInfo{}
	vehicleInfo:=&model.VehicleInfo{}
	vehicleLeader:=&model.VehicleLeader{}
	user:=&model.User{}
	asset:=&model.Asset{}
	threat:=&model.Threat{}
	whiteList:=&model.WhiteList{}
	portMap:=&model.PortMap{}
	strategy:=&model.Strategy{}
	strategyVehicle:=&model.StrategyVehicle{}
	strategyGroup:=&model.StrategyGroup{}
	strategyGroupsLearningResult:=&model.StrategyGroupLearningResult{}
	flowStrategy:=&model.Fstrategy{}
	flowStrategyVehicles:=&model.FstrategyVehicle{}
	flowStrategyTtem:=&model.FstrategyItem{}
	flowStrategyRelateItem:=&model.FstrategyVehicleItem{}
	sample:=&model.Sample{}
	sampleItem:=&model.SampleItem{}
	studyOrigin:=&model.StudyOrigin{}
	automatedLearning:=&model.AutomatedLearning{}
	automatedLearningResult:=&model.AutomatedLearningResult{}

	tables:=[]interface{}{
		flow,firmwareUpdate,firmwareInfo,vehicleInfo,vehicleLeader,
		user,asset,threat,whiteList,portMap,strategy,strategyVehicle,strategyGroup,
		strategyGroupsLearningResult,flowStrategy,flowStrategyVehicles,flowStrategyVehicles,
		flowStrategyTtem,flowStrategyRelateItem,sample,
		sampleItem,studyOrigin,automatedLearning,automatedLearningResult,
	}

	return tables
}