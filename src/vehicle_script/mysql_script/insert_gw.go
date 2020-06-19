package main

import (
	"flag"
	"fmt"
)

func main() {

	//获取命令行参数
	var defaultVehicleCount int
	flag.IntVar(&defaultVehicleCount, "count", 50, "插入小v数量")
	flag.Parse()

	fmt.Println(defaultVehicleCount, "defaultVehicleCount")
	//查询小v最大的gwid

	//获取未分组id

	////插入小v
	//for i := 0; i < defaultVehicleCount; i++ {
	//	finalVehicleId := strconv.Itoa(i)
	//	vehicleInfoModel := &model.VehicleInfo{
	//		VehicleId:       finalVehicleId,
	//		Name:            finalVehicleId,
	//		Version:         tools.GenVersion(),
	//		StartTime:       time.Now(),
	//		FirmwareVersion: tools.GenVersion(),
	//		HardwareModel:   tools.GenVersion(),
	//		Module:          tools.RandomString(8),
	//		SupplyId:        tools.RandomString(8),
	//		UpRouterIp:      tools.GenIpAddr(),
	//
	//		Ip:   tools.GenIpAddr(),
	//		Type: tools.GenBrandType(1),
	//
	//		Mac: finalVehicleId,
	//
	//		OnlineStatus:  tools.RandomAlternativeBool(1),
	//		ProtectStatus: tools.GenBrandType(1),
	//	}
	//
	//	_ = tools.CreateModel(&vehicleInfoModel)
	//}

}
