package model

import "github.com/jinzhu/gorm"

type PortMap struct {
	gorm.Model
	PortMapId             string
	VehicleId  			  string //关联某设备
	SrcPort               string //车载
	DstPort               string//车载资产
	DestIp 				  string
	Switch                bool//是否开启端口映射
	ProtocolType          uint8//网络协议类型

}