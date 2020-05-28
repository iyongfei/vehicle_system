package model

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

type VehicleInfo struct {
	gorm.Model
	VehicleId string `gorm:"unique"` //小v ID
	Name      string //小v名称
	Version   string
	//StartTime       model_base.UnixTime //启动时间
	StartTime       time.Time //启动时间
	FirmwareVersion string
	HardwareModel   string
	Module          string
	SupplyId        string
	UpRouterIp      string

	Ip        string
	Type      uint8
	Mac       string //Mac地址
	TimeStamp uint32 //最近活跃时间戳
	HbTimeout uint32 //最近活跃时间戳

	DeployMode       uint8 //部署模式
	FlowIdleTimeSlot uint32

	OnlineStatus  bool   //在线状态
	ProtectStatus uint8  //保护状态										//保护状态
	LeaderId      string //保护状态 // 保护状态
	GroupId       string
}

type VehicleInfoT struct {
	//gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt int64
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	VehicleId       string `gorm:"unique"` //小v ID
	Name            string //小v名称
	Version         string
	StartTime       time.Time //启动时间
	FirmwareVersion string
	HardwareModel   string
	Module          string
	SupplyId        string
	UpRouterIp      string

	Ip        string
	Type      uint8
	Mac       string //Mac地址
	TimeStamp uint32 //最近活跃时间戳
	HbTimeout uint32 //最近活跃时间戳

	DeployMode       uint8 //部署模式
	FlowIdleTimeSlot uint32

	OnlineStatus  bool   //在线状态
	ProtectStatus uint8  //保护状态										//保护状态
	LeaderId      string //保护状态 // 保护状态
	GroupId       string
	FlowCount     int
}

func (vehicle *VehicleInfo) AfterCreate(tx *gorm.DB) error {
	logger.Logger.Print("%s afterCreate vehicle_id:%s", util.RunFuncName(), vehicle.VehicleId)
	logger.Logger.Info("%s afterCreate vehicle_id:%s", util.RunFuncName(), vehicle.VehicleId)

	//err := HandleVehicleStrategyInitAction(vehicle.VehicleId)
	//if err != nil {
	//	logger.Logger.Print("%s afterCreate vehicle_id:%s,init strategy err:%s", util.RunFuncName(), vehicle.VehicleId, err)
	//	logger.Logger.Info("%s afterCreate vehicle_id:%s,init strategy err:%s", util.RunFuncName(), vehicle.VehicleId, err)
	//}
	//
	////下发策略
	//initVehicleStrategy(vehicle.VehicleId)
	return nil
}

var InitVehicleStrategyChan = make(chan string, 5)

func initVehicleStrategy(vehicleId string) {
	InitVehicleStrategyChan <- vehicleId
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}

//序列化为数字类型
func (vehicle *VehicleInfo) MarshalJSON() ([]byte, error) {
	type VehicleType VehicleInfo
	return json.Marshal(&struct {
		StartTime int64
		*VehicleType
	}{
		StartTime:   time.Time(vehicle.StartTime).Unix(),
		VehicleType: (*VehicleType)(vehicle),
	})
}

func (vehicle *VehicleInfo) UnmarshalJSON(data []byte) error {
	type VehicleType VehicleInfo
	aux := &struct {
		StartTime int64
		*VehicleType
	}{
		VehicleType: (*VehicleType)(vehicle),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	vehicle.StartTime = time.Unix(aux.StartTime, 0)
	return nil
}

func (u *VehicleInfo) InsertModel() error {
	return mysql.CreateModel(u)
}
func (u *VehicleInfo) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(u, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (u *VehicleInfo) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(u, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (u *VehicleInfo) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(u, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (u *VehicleInfo) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (vehicleInfo *VehicleInfo) CreateModel(vehicleParam ...interface{}) interface{} {
	vehicleParams := vehicleParam[0].(*protobuf.GwInfoParam)
	vehicleInfo.Version = vehicleParams.GetVersion()
	if vehicleParams.GetStartTime() == 0 {
		//vehicleInfo.StartTime = model_base.UnixTime(time.Now())
		vehicleInfo.StartTime = time.Now()
	} else {
		//vehicleInfo.StartTime = model_base.UnixTime(util.StampUnix2Time(int64(vehicleParams.GetStartTime())))
		vehicleInfo.StartTime = util.StampUnix2Time(int64(vehicleParams.GetStartTime()))
	}

	vehicleInfo.FirmwareVersion = vehicleParams.GetFirmwareVersion()
	vehicleInfo.HardwareModel = vehicleParams.GetHardwareModel()
	mapModule := make(map[string]string)
	for _, moduleItem := range vehicleParams.GetModule() {
		_, ok := mapModule[moduleItem.GetName()]
		if ok {
			continue
		}
		mapModule[moduleItem.GetName()] = moduleItem.GetVersion()
	}
	mapModuleMarshal, _ := json.Marshal(mapModule)
	vehicleInfo.Module = string(mapModuleMarshal)

	vehicleInfo.SupplyId = vehicleParams.GetSupplyId()
	vehicleInfo.UpRouterIp = vehicleParams.GetUpRouterIp()
	vehicleInfo.Ip = vehicleParams.GetIp()
	vehicleInfo.Type = uint8(vehicleParams.GetType())
	vehicleInfo.Mac = vehicleParams.GetMac()
	vehicleInfo.TimeStamp = uint32(vehicleParams.GetTimestamp())
	vehicleInfo.HbTimeout = uint32(vehicleParams.GetHbTimeout())
	vehicleInfo.DeployMode = uint8(vehicleParams.GetDeployMode())
	vehicleInfo.FlowIdleTimeSlot = uint32(vehicleParams.GetFlowIdleTimeSlot())
	vehicleInfo.OnlineStatus = true
	return vehicleInfo
}

func (vehicleInfo *VehicleInfo) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(vehicleInfo, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
