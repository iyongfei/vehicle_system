package main

import (
	"github.com/golang/protobuf/proto"
	"strconv"
	"vehicle_system/src/vehicle_script/emq_service"
	"vehicle_system/src/vehicle_script/emq_service/protobuf"
	"vehicle_system/src/vehicle_script/tool"
)

/**
添加车载信息
insert_vehicle_count
 */
const (
	Conf_Name = "conf.txt"
	InsertVehicleCount = "insert_vehicle_count"
)

func main()  {
	configMap := tool.InitConfig(Conf_Name)
	insertVehicleCount := configMap[InsertVehicleCount]
	defaultVehicleCount ,_ := strconv.Atoi(insertVehicleCount)

	emqx:= emq_service.NewEmqx()
	for i:=0;i< defaultVehicleCount;i++{
		vid:=tool.RandomString(32)
		//vid:="dgzeKAoBGbl5E5ajqOq1phksMCVz8S7C"
		emqx.Publish(vid,createGwProbuf(vid))
	}
}

/**
Version              string                    `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	StartTime            uint32                    `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	FirmwareVersion      string                    `protobuf:"bytes,3,opt,name=firmware_version,json=firmwareVersion,proto3" json:"firmware_version,omitempty"`
	HardwareModel        string                    `protobuf:"bytes,4,opt,name=hardware_model,json=hardwareModel,proto3" json:"hardware_model,omitempty"`
	Module               []*GwInfoParam_ModuleItem `protobuf:"bytes,5,rep,name=module,proto3" json:"module,omitempty"`
	SupplyId             string                    `protobuf:"bytes,6,opt,name=supply_id,json=supplyId,proto3" json:"supply_id,omitempty"`
	UpRouterIp           string                    `protobuf:"bytes,7,opt,name=up_router_ip,json=upRouterIp,proto3" json:"up_router_ip,omitempty"`
	Ip                   string                    `protobuf:"bytes,8,opt,name=ip,proto3" json:"ip,omitempty"`
	Type                 GwInfoParam_Type          `protobuf:"varint,9,opt,name=type,proto3,enum=protobuf.GwInfoParam_Type" json:"type,omitempty"`
	Mac                  string                    `protobuf:"bytes,10,opt,name=mac,proto3" json:"mac,omitempty"`
	Timestamp            uint32                    `protobuf:"varint,11,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	HbTimeout            int32                     `protobuf:"varint,12,opt,name=hb_timeout,json=hbTimeout,proto3" json:"hb_timeout,omitempty"`
	DeployMode           GwInfoParam_DeployMode    `protobuf:"varint,13,opt,name=deploy_mode,json=deployMode,proto3,enum=protobuf.GwInfoParam_DeployMode" json:"deploy_mode,omitempty"`
	FlowIdleTimeSlot     uint32
 */
func createGwProbuf(vId string)[]byte{
	pushReq:=&protobuf.GWResult{
		ActionType:protobuf.GWResult_GW_INFO,
		GUID:vId,
	}

	params := &protobuf.GwInfoParam{
		Version:tool.GenVersion(),
		StartTime:tool.TimeNowToUnix(),
		FirmwareVersion:tool.GenVersion(),
		HardwareModel:tool.GenBrand(1),
		SupplyId:tool.RandomString(5),
		UpRouterIp:tool.GenIpAddr(),
		Ip:tool.GenIpAddr(),
		Type:protobuf.GwInfoParam_DEFAULT,
		Mac:tool.RandomString(8),
		Timestamp:tool.TimeNowToUnix(),
		HbTimeout:30,
		DeployMode:protobuf.GwInfoParam_SWITCHMODE,
		FlowIdleTimeSlot:1212,
	}
	//module begin
	items:=[]*protobuf.GwInfoParam_ModuleItem{}
	moduleItem1 := &protobuf.GwInfoParam_ModuleItem{
		Name:tool.RandomString(10),
		Version:tool.GenVersion(),
	}
	moduleItem2 := &protobuf.GwInfoParam_ModuleItem{
		Name:tool.RandomString(10),
		Version:tool.GenVersion(),
	}
	items = append(items,moduleItem1,moduleItem2)
	//module end
	params.Module = items
	bys,_:=proto.Marshal(params)
	///////////////////////////////////

	pushReq.Param = bys
	ret,_:=proto.Marshal(pushReq)
	return  ret
}