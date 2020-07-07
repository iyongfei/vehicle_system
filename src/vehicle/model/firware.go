package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

//
type FirmwareInfo struct {
	gorm.Model
	System        string //操作系統
	Version       string //版本
	Soctype       string //cpu
	HardwareModel string //硬件
	Md5           string //md5

	FirmwareType string    //固件類型sysupgrade
	FirmwareName string    //固件文件名
	Size         uint64    //大小
	UpdateInfo   string    //更新內容
	UploadTime   time.Time //上傳時間
	BinPath      string    //二進制文件路徑
}

type FirmwareUpdate struct {
	gorm.Model
	DeployId         string
	VehicleId        string
	UpdateVersion    string //版本校驗
	UpgradeTimestamp uint32 //開始下載時間戳
	UpgradeStatus    uint8  //下載狀態
	Timeout          uint32
}

func (firmware *FirmwareUpdate) InsertModel() error {
	return mysql.CreateModel(firmware)
}
func (firmware *FirmwareUpdate) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(firmware, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (firmware *FirmwareUpdate) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(firmware, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (firmware *FirmwareUpdate) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (firmware *FirmwareUpdate) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	return nil
}

func (firmware *FirmwareUpdate) CreateModel(firmwareParams ...interface{}) interface{} {
	firmwareParam := firmwareParams[0].(*protobuf.FirwareParam)

	firmware.UpdateVersion = firmwareParam.GetVersion()
	firmware.UpgradeTimestamp = firmwareParam.GetUpgradeTimestamp()
	firmware.UpgradeStatus = uint8(firmwareParam.GetUpgradeStatus())
	firmware.Timeout = firmwareParam.GetTimeout()

	return firmware
}
