package model

import "github.com/jinzhu/gorm"

type FlowStatistic struct {
	gorm.Model

	FlowId       uint32
	VehicleId    string
	Hash         uint32
	SrcIp        string
	SrcPort      uint32
	DstIp        string
	DstPort      uint32
	Protocol     uint8
	FlowInfo     string
	SafeType     uint8
	SafeInfo     string
	StartTime    uint32
	LastSeenTime uint32
	SrcDstBytes  uint64
	DstSrcBytes  uint64
	Stat         uint8
}

//
////message FlowStatisticParam {
////string interface_name = 1; //网卡名称
////uint64 rx = 2; //接收总数据，字节
////uint64 tx = 3; //上传总数据，字节
////uint32 flow_count = 4; //探测的flow的总数
////uint32 pub = 5; //本次发布的flow数量
////uint32 notlocal = 6; //不是与本地相关的flow数
////uint32 white =7; //其中属于白名单的数量
//}
