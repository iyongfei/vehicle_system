package model

import "github.com/jinzhu/gorm"

type Flow struct {
	gorm.Model

	FlowId          string
	Hash            uint32
	SrcIp           uint32
	SrcPort         uint32
	DstIp           uint32
	DstPort         uint32
	Protocol        uint8
	FlowInfo        string
	SafeType        uint8
	SafeInfo        string
	StartTime       uint32
	LastSeenTime    uint32
	SrcDstBytes     uint64
	DstSrcBytes     uint64
}
