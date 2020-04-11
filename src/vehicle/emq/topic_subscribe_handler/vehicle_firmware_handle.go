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

func HandleVehicleFirmware(vehicleResult protobuf.GWResult) error {
	cmdDeployId:=vehicleResult.GetCmdID()
	//parse
	vehicleFirmwareParam := &protobuf.FirwareParam{}
	err:=proto.Unmarshal(vehicleResult.GetParam(),vehicleFirmwareParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicle firmware param err:%s",util.RunFuncName(),err.Error())
		logger.Logger.Error("%s unmarshal vehicle firmware param err:%s",util.RunFuncName(),err.Error())
		return fmt.Errorf("%s unmarshal vehicle firmware err:%s",util.RunFuncName(),err.Error())
	}
	//vehicleId
	vehicleId:=vehicleResult.GetGUID()

	logger.Logger.Print("%s unmarshal vehicle firmware param:%+v",util.RunFuncName(),vehicleFirmwareParam)
	logger.Logger.Info("%s unmarshal vehicle firmware param:%+v",util.RunFuncName(),vehicleFirmwareParam)
	//create
	firmwareUpdate:=&model.FirmwareUpdate{
		DeployId:cmdDeployId,
		VehicleId:vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(firmwareUpdate)

	_,recordNotFound :=modelBase.GetModelByCondition("deploy_id = ? and vehicle_id = ?",
		[]interface{}{firmwareUpdate.DeployId, firmwareUpdate.VehicleId}...)

	modelBase.CreateModel(vehicleFirmwareParam)

	if recordNotFound{
		if err:=modelBase.InsertModel();err!=nil{
			return fmt.Errorf("%s insert vehicle firmware err:%s",util.RunFuncName(),err.Error())
		}

	}else {
		//更新
		attrs := map[string]interface{}{
			"update_version": firmwareUpdate.UpdateVersion,
			"upgrade_timestamp": firmwareUpdate.UpgradeTimestamp,
			"upgrade_status": firmwareUpdate.UpgradeStatus,
			"timeout": firmwareUpdate.Timeout,
		}
		if err:=modelBase.UpdateModelsByCondition(attrs,"deploy_id = ? and vehicle_id = ?",
			[]interface{}{firmwareUpdate.DeployId, firmwareUpdate.VehicleId}...);err!=nil{
			return fmt.Errorf("%s update vehicle firmware err:%s",util.RunFuncName(),err.Error())
		}
	}
	return nil
}