package model

import (
	"github.com/jinzhu/gorm"
	"time"
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
