package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleThreat(vehicleResult protobuf.GWResult) error {

	//parse
	vehicleParam := &protobuf.ThreatParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), vehicleParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicleParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal vehicleParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal vehicleParam err:%s", util.RunFuncName(), err.Error())
	}
	//vehicleId
	vehicleId := vehicleResult.GetGUID()

	logger.Logger.Print("%s unmarshal vehicleParam:%+v", util.RunFuncName(), vehicleParam)
	logger.Logger.Info("%s unmarshal vehicleParam:%+v", util.RunFuncName(), vehicleParam)
	//create
	vehicleInfo := &model.VehicleInfo{}
	vehicleInfo.VehicleId = vehicleId
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	_, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", vehicleInfo.VehicleId)

	if recordNotFound {
		return fmt.Errorf("%s insert threat vehicleId recordNotFound", util.RunFuncName())
	}
	for _, threatItem := range vehicleParam.GetThreatItem() {
		threat := &model.Threat{}
		threat.ThreatId = util.RandomString(32)
		threat.VehicleId = vehicleId

		modelBase := model_base.ModelBaseImpl(threat)

		_, recordNotFound := modelBase.GetModelByCondition("threat_id = ?", threat.ThreatId)
		if !recordNotFound {
			continue
		}

		modelBase.CreateModel(threatItem)
		err := modelBase.InsertModel()
		if err != nil {
			continue
		}
	}
	return nil
}
