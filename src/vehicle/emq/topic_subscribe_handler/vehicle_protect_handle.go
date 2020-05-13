package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/service/push"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleProtect(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	vehicleProtectParam := &protobuf.GWProtectInfoParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), vehicleProtectParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicle protect param err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal vehicle protect param err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal vehicle protect err:%s", util.RunFuncName(), err.Error())
	}

	logger.Logger.Print("%s handle_protect:%+v", util.RunFuncName(), *vehicleProtectParam)
	logger.Logger.Info("%s handle_protect:%+v", util.RunFuncName(), *vehicleProtectParam)
	//create
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	_, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", vehicleInfo.VehicleId)

	if recordNotFound {
		return fmt.Errorf("%s handleVehicleProtect vehicle:%s recordNotFound", util.RunFuncName(), vehicleId)
	} else {
		//更新
		attrs := map[string]interface{}{
			"protect_status": vehicleProtectParam.GetProtectStatus(),
		}
		if err := modelBase.UpdateModelsByCondition(attrs, "vehicle_id = ?", vehicleInfo.VehicleId); err != nil {
			return fmt.Errorf("%s update vehicle protect err:%s", util.RunFuncName(), err.Error())
		}
	}

	pushActionTypeName := protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]
	pushVehicleid := vehicleId
	pushData := map[string]interface{}{
		"protect_status": vehicleProtectParam.GetProtectStatus(),
	}
	fPushData := push.CreatePushData(pushActionTypeName, pushVehicleid, pushData)

	push.GetPushervice().SetPushData(fPushData).Write()

	return nil
}
