package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type FirmwareInfo struct {
	gorm.Model
	System string//操作系統
	Version string//版本
	Soctype string //cpu
	HardwareModel string //硬件
	Md5 string//md5

	FirmwareType string //固件類型sysupgrade
	FirmwareName string//固件文件名
	Size uint64//大小
	UpdateInfo string//更新內容
	UploadTime time.Time//上傳時間
	BinPath string//二進制文件路徑
}



type FirmwareUpdate struct {
	gorm.Model
	DeployId		string
	VehicleId   string
	UpdateVersion string//版本校驗
	UpgradeTimestamp uint32 //開始下載時間戳
	UpgradeStatus uint8//下載狀態
	Timeout uint32
}

