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
		//更新 排除VehicleId,Name,ProtectStatus,LeaderId
		attrs:= map[string]interface{}{
			"version":vehicleInfo.Version,
			"start_time":vehicleInfo.StartTime,
			"firmware_version":vehicleInfo.FirmwareVersion,
			"hardware_model":vehicleInfo.HardwareModel,
			"module":vehicleInfo.Module,
			"supply_id":vehicleInfo.SupplyId,
			"up_router_ip":vehicleInfo.UpRouterIp,

			"ip":vehicleInfo.Ip,
			"type":vehicleInfo.Type,
			"mac":vehicleInfo.Mac,
			"time_stamp":vehicleInfo.TimeStamp,
			"hb_timeout":vehicleInfo.HbTimeout,
			"deploy_mode":vehicleInfo.DeployMode,
			"flow_idle_time_slot":vehicleInfo.FlowIdleTimeSlot,
			"online_status":vehicleInfo.OnlineStatus,
		}
		if err:=modelBase.UpdateModelsByCondition(attrs,"vehicle_id = ?",vehicleInfo.VehicleId);err!=nil{
			return fmt.Errorf("%s update vehicle err:%s",util.RunFuncName(),err.Error())
		}
	}
	return nil
}

/**
string version = 1;//软件版本
	uint32 start_time = 2;//启动时间
	string firmware_version = 3;//固件版本
	string hardware_model= 4;//硬件型号
	repeated ModuleItem module = 5;//模块信息
	string supply_id = 6;//渠道id
	string up_router_ip = 7;//上级路由ip
	string ip = 8;//盒子ip
	Type type = 9;//盒子类型
	string mac = 10;//盒子mac
	uint32 timestamp = 11;//时间戳
	int32 hb_timeout = 12; //心跳超时(秒),用于显示小V在线/离线
	DeployMode deploy_mode = 13; //网关部署模式
	uint32 flow_idle_time_slot = 14; //会话flow在时间段间隔之后做IDLE处理.
 */