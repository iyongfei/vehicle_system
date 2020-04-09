package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

type VehicleInfo struct {
	gorm.Model
	VehicleId       			string 		`gorm:"unique"`				//小v ID
	Name       		string 				//小v名称
	Version         string
	StartTime    	uint32//启动时间
	FirmwareVersion string
	HardwareModel   string
	Module          string
	SupplyId        string
	UpRouterIp      string

	Ip         	string
	Type            uint8
	Mac        		string       		//Mac地址
	TimeStamp 	uint32	//最近活跃时间戳
	HbTimeout 	uint32	//最近活跃时间戳

	DeployMode 	uint8	//部署模式

	OnlineStatus      bool											//在线状态
	ProtectStatus   	uint8	//保护状态										//保护状态
}



func (u *VehicleInfo) InsertModel() error {
	return mysql.CreateModel(u)
}

func (u *VehicleInfo) GetModelByCondition(query interface{}, args ...interface{}) (error,bool) {
	err,recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(u,query,args...)
	if err!=nil{
		return err,true
	}
	if recordNotFound{
		return nil,true
	}
	return nil,false
}
func (u *VehicleInfo) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(u,values,query,queryArgs...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (u *VehicleInfo) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}

func (u *VehicleInfo) GetModelListByCondition(model interface{},query interface{}, args ...interface{}) (error) {
	return nil
}

func (vehicleInfo *VehicleInfo)  CreateModel(vehicleParam ...interface{})  interface{}{
	vehicleParams := vehicleParam[0].(*protobuf.GwInfoParam)
	//vehicleInfo:= &VehicleInfo{}
	vehicleInfo.Version = vehicleParams.GetVersion()
	vehicleInfo.FirmwareVersion = vehicleParams.GetFirmwareVersion()
	vehicleInfo.HardwareModel = vehicleParams.GetHardwareModel()
	vehicleInfo.BindIp = vehicleParams.GetIp()
	vehicleInfo.SupplyId = vehicleParams.GetSupplyId()
	vehicleInfo.UpRouterIp = vehicleParams.GetUpRouterIp()
	vehicleInfo.Type = uint8(vehicleParams.GetType())
	vehicleInfo.Mac = vehicleParams.GetMac()
	vehicleInfo.RecentActiveTime = uint64(vehicleParams.GetTimestamp())
	vehicleInfo.DeployMode = int32(vehicleParams.GetDeployMode())
	vehicleInfo.OnlineStatus = true

	mapModule:=make(map[string]string)
	for _,moduleItem := range vehicleParams.GetModule(){
		_,ok := mapModule[moduleItem.GetName()]
		if ok {
			continue
		}
		mapModule[moduleItem.GetName()] = moduleItem.GetVersion()
	}
	mapModuleMarshal,_ := json.Marshal(mapModule)
	vehicleInfo.Module = string(mapModuleMarshal)


	if vehicleParams.GetStartTime() == 0{
		vehicleInfo.StartTime = time.Now()
	}else {
		vehicleInfo.StartTime = util.StampUnix2Time(int64(vehicleParams.GetStartTime()))
	}
	return vehicleInfo
}

