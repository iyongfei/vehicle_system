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

func HandleVehicleInfo(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	vehicleParam := &protobuf.GwInfoParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), vehicleParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicleParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal vehicleParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal vehicleParam err:%s", util.RunFuncName(), err.Error())
	}

	logger.Logger.Print("%s unmarshal vehicleParam:%+v", util.RunFuncName(), vehicleParam)
	logger.Logger.Info("%s unmarshal vehicleParam:%+v", util.RunFuncName(), vehicleParam)
	//create
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	_, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", vehicleInfo.VehicleId)

	modelBase.CreateModel(vehicleParam)
	if recordNotFound {
		if err := modelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert vehicle err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		//更新 排除VehicleId,Name,ProtectStatus,LeaderId
		attrs := map[string]interface{}{
			"version":          vehicleInfo.Version,
			"start_time":       vehicleInfo.StartTime,
			"firmware_version": vehicleInfo.FirmwareVersion,
			"hardware_model":   vehicleInfo.HardwareModel,
			"module":           vehicleInfo.Module,
			"supply_id":        vehicleInfo.SupplyId,
			"up_router_ip":     vehicleInfo.UpRouterIp,

			"ip":                  vehicleInfo.Ip,
			"type":                vehicleInfo.Type,
			"mac":                 vehicleInfo.Mac,
			"time_stamp":          vehicleInfo.TimeStamp,
			"hb_timeout":          vehicleInfo.HbTimeout,
			"deploy_mode":         vehicleInfo.DeployMode,
			"flow_idle_time_slot": vehicleInfo.FlowIdleTimeSlot,
			"online_status":       vehicleInfo.OnlineStatus,
		}
		if err := modelBase.UpdateModelsByCondition(attrs, "vehicle_id = ?", vehicleInfo.VehicleId); err != nil {
			return fmt.Errorf("%s update vehicle err:%s", util.RunFuncName(), err.Error())
		}
	}

	pushActionTypeName := protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]
	pushVehicleid := vehicleId
	pushData := map[string]interface{}{
		"vehicle_info": vehicleInfo,
	}

	fPushData := push.CreatePushData(pushActionTypeName, pushVehicleid, pushData)

	push.GetPushervice().SetPushData(fPushData).Write()

	return nil
}
