package tdata

import (
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/model"
)

/**
检测数据库表
*/

func TableCheck() {
	tables := getTables()

	for _, table := range tables {
		err, _ := checkTable(table)
		if err != nil {
			err, _ = checkTable(table)
			continue
		}
	}
}

func checkTable(model interface{}) (error, int64) {
	err, rowsAffected := mysql.CreatTable(model)
	return err, rowsAffected
}

func getTables() []interface{} {
	//添加各个&model表
	alayerProto := &model.AlayerProto{}
	area := &model.AreaGroup{}
	assetArea := &model.AssetGroup{}
	assetLeader := &model.AssetLeader{}
	asset := &model.Asset{}
	automatedLearning := &model.AutomatedLearning{}
	automatedLearningResult := &model.AutomatedLearningResult{}
	category := &model.Category{}
	disk := &model.Disk{}
	fingerPrint := &model.FingerPrint{}
	firmwareInfo := &model.FirmwareInfo{}
	firmwareUpdate := &model.FirmwareUpdate{}

	flowStatistic := &model.FlowStatistic{}
	flow := &model.Flow{}
	flowStrategy := &model.Fstrategy{}
	flowStrategyVehicles := &model.FstrategyVehicle{}
	flowStrategyTtem := &model.FstrategyItem{}
	flowStrategyRelateItem := &model.FstrategyVehicleItem{}
	portMap := &model.PortMap{}
	redisInfo := &model.RedisInfo{}
	sampleItem := &model.SampleItem{}
	sample := &model.Sample{}
	strategy := &model.Strategy{}
	strategyVehicle := &model.StrategyVehicle{}
	strategyVehicleLearnings := &model.StrategyVehicleLearningResult{}
	strategyGroup := &model.StrategyGroup{}

	studyOrigin := &model.StudyOrigin{}

	tflow := &model.TempFlow{}
	threat := &model.Threat{}
	user := &model.User{}

	vehicleInfo := &model.VehicleInfo{}
	vehicleLeader := &model.VehicleLeader{}

	vhalonet := &model.VhaloNets{}
	whiteList := &model.WhiteList{}

	tables := []interface{}{
		alayerProto, area, assetArea, asset, assetLeader, category, fingerPrint, flow, tflow, flowStatistic, firmwareUpdate, firmwareInfo, vehicleInfo, vehicleLeader,
		user, disk, threat, whiteList, vhalonet, redisInfo, portMap, strategy, strategyGroup, strategyVehicle,
		flowStrategy, flowStrategyVehicles,
		flowStrategyTtem, flowStrategyRelateItem, sample, strategyVehicleLearnings,
		sampleItem, studyOrigin, automatedLearning, automatedLearningResult,
	}

	return tables
}
