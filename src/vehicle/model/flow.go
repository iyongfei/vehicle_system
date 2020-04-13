package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

type Flow struct {
	gorm.Model

	FlowId          uint32
	VehicleId       string
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
	Stat        uint8
}



func (f *Flow) InsertModel() error {
	return mysql.CreateModel(f)
}

func (f *Flow) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(f, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}

func (f *Flow) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(f, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (f *Flow) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}

func (f *Flow) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}

func (flow *Flow) CreateModel(flowParam ...interface{}) interface{} {
	flowItemParams := flowParam[0].(*protobuf.FlowParam_FItem)
	flow.Hash = flowItemParams.GetHash()
	flow.SrcIp = flowItemParams.GetSrcIp()
	flow.SrcPort = flowItemParams.GetSrcPort()
	flow.DstIp = flowItemParams.GetDstIp()
	flow.DstPort = flowItemParams.GetDstPort()
	flow.Protocol = uint8(flowItemParams.GetProtocol())
	flow.FlowInfo = flowItemParams.GetFlowInfo()
	flow.SafeType = uint8(flowItemParams.GetSafeType())
	flow.SafeInfo = flowItemParams.GetSafeInfo()
	flow.StartTime = flowItemParams.GetStartTime()
	flow.LastSeenTime = flowItemParams.GetLastSeenTime()
	flow.SrcDstBytes = flowItemParams.GetSrc2DstBytes()
	flow.DstSrcBytes = flowItemParams.GetDst2SrcBytes()
	flow.Stat = uint8(flowItemParams.GetFlowStat())
	return flow
}




































