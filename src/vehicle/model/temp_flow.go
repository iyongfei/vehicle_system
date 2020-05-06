package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

type TempFlow struct {
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

type TempFlowDp struct {
	VehicleId string
	Dps       map[string][]uint32
}

//序列化为数字类型
func (tflow *TempFlow) MarshalJSON() ([]byte, error) {
	type FlowType TempFlow
	return json.Marshal(&struct {
		Protocol string
		*FlowType
	}{
		Protocol: protobuf.GetFlowProtocols(int(tflow.Protocol)),
		FlowType: (*FlowType)(tflow),
	})
}

//
//func (flow *Flow) UnmarshalJSON(data []byte) error {
//	type FlowType Flow
//	aux := &struct {
//		StartTime int64
//		*FlowType
//	}{
//		FlowType: (*FlowType)(flow),
//	}
//	if err := json.Unmarshal(data, &aux); err != nil {
//		return err
//	}
//	flow.StartTime = time.Unix(aux.StartTime, 0)
//	return nil
//}

func (f *TempFlow) InsertModel() error {
	return mysql.CreateModel(f)
}

func (f *TempFlow) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(f, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}

func (f *TempFlow) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(f, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (f *TempFlow) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(f, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (f *TempFlow) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (flow *TempFlow) CreateModel(flowParam ...interface{}) interface{} {
	flowItemParams := flowParam[0].(*protobuf.FlowParam_FItem)
	flow.Hash = flowItemParams.GetHash()

	sipLittleEndian := util.BytesToLittleEndian(util.UintToBytes(flowItemParams.GetSrcIp()))
	flow.SrcIp = util.IpIntToString(int(sipLittleEndian))

	flow.SrcPort = flowItemParams.GetSrcPort()

	dipLittleEndian := util.BytesToLittleEndian(util.UintToBytes(flowItemParams.GetDstIp()))
	flow.DstIp = util.IpIntToString(int(dipLittleEndian))

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

func (flow *TempFlow) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(flow, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
