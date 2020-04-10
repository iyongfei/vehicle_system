package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Sample struct {
	gorm.Model
	SampleId            string // 采集样本id
	StartTime    	time.Time//启动时间
	RemainTime    uint32 //剩余时间
	TotalTime    uint32 //剩余时间

	Status       uint8 //采集状态
	TimeOut                uint32//超时时间

	Name          string //采集名称
	Introduce           string     //采集说明

	Check uint8

	VehicleId       			string
	StudyOriginId        string
}

type SampleItem struct {
	gorm.Model
	SampleItemId string//id

	SampleId string

	SrcMac    string//源mac sm
	SrcIp     string//源ip sip
	SrcPort   uint32//源端口 sp

	DstIp       string///////////////目标ip  dip
	DstPort     uint32//目标端口  dp
	Url          string/////////////////目标url u

	FetchTime  time.Time//访问时间tm
}
