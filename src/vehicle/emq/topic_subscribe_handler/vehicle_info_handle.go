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

func HandleVehicleInfo(vehicleResult protobuf.GWResult) error {
	//parse
	vehicleParam := &protobuf.GwInfoParam{}
	err:=proto.Unmarshal(vehicleResult.GetParam(),vehicleParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicleParam err:%s",util.RunFuncName(),err.Error())
		logger.Logger.Error("%s unmarshal vehicleParam err:%s",util.RunFuncName(),err.Error())
		return fmt.Errorf("%s unmarshal vehicleParam err:%s",util.RunFuncName(),err.Error())
	}
	//vehicleId
	vehicleId:=vehicleResult.GetGUID()

	logger.Logger.Print("%s unmarshal vehicleParam:%+v",util.RunFuncName(),vehicleParam)
	logger.Logger.Info("%s unmarshal vehicleParam:%+v",util.RunFuncName(),vehicleParam)
	//create
	vehicleInfo:=&model.VehicleInfo{}
	vehicleInfo.VehicleId = vehicleId
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	_,recordNotFound :=modelBase.GetModelByCondition("vehicle_id = ?",vehicleInfo.VehicleId)

	modelBase.CreateModel(vehicleParam)
	if recordNotFound {
		if err:=modelBase.InsertModel();err!=nil{
			return fmt.Errorf("%s insert vehicle err:%s",util.RunFuncName(),err.Error())
		}
	}else {
		//更新
		attrs:= map[string]interface{}{
			"bind_ip":vehicleInfo.BindIp,
			"mac":vehicleInfo.Mac,
			"version":vehicleInfo.Version,
			"start_time":vehicleInfo.StartTime,
			"firmware_version":vehicleInfo.FirmwareVersion,
			"hardware_model":vehicleInfo.HardwareModel,
			"supply_id":vehicleInfo.SupplyId,
			"up_router_ip":vehicleInfo.UpRouterIp,
			"module":vehicleInfo.Module,
			"type":vehicleInfo.Type,
			"gw_timeout":vehicleInfo.GwTimeout,
			"online_status":vehicleInfo.OnlineStatus,
			"protect_status":vehicleInfo.ProtectStatus,
			"recent_active_time":vehicleInfo.RecentActiveTime,
			"deploy_mode":vehicleInfo.DeployMode,
		}
		if err:=modelBase.UpdateModelsByCondition(attrs,"vehicle_id = ?",vehicleInfo.VehicleId);err!=nil{
			return fmt.Errorf("%s update vehicle err:%s",util.RunFuncName(),err.Error())
		}
	}
	return nil
}

