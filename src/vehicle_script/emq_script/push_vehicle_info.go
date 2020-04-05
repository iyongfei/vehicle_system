package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strconv"
	"vehicle_system/src/vehicle_script/emp_service"
	"vehicle_system/src/vehicle_script/emp_service/protobuf"
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
	fmt.Println("defaultVehicleCount:",defaultVehicleCount)

	emqx:=emp_service.NewEmqx()
	for i:=0;i< defaultVehicleCount;i++{
		vid:=tool.RandomString(32)
		//vid:="dgzeKAoBGbl5E5ajqOq1phksMCVz8S7C"
		fmt.Println("vid:",vid)
		emqx.Publish(vid,createGwProbuf(vid))
	}
}


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