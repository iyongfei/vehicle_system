package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

type FprintFlow struct {
	gorm.Model

	FlowId       uint32
	VehicleId    string
	AssetId      string
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

	//add
	SrcDstPackets uint64
	DstSrcPackets uint64

	HostName string
	Category uint32

	HasPassive       bool
	IatFlowAvg       float32
	IatFlowStddev    float32
	DataRatio        float32
	StrDataRatio     uint8
	PktlenCToSAvg    float32
	PktlenCToSStddev float32
	PktlenSToCAvg    float32
	PktlenSToCStddev float32
	TlsClientInfo    string
	Ja3c             string
}

//序列化为数字类型
func (flow *FprintFlow) MarshalJSON() ([]byte, error) {
	type FlowType FprintFlow
	return json.Marshal(&struct {
		Protocol  string
		CreatedAt int64
		*FlowType
	}{
		Protocol:  protobuf.GetFlowProtocols(int(flow.Protocol)),
		CreatedAt: flow.CreatedAt.Unix(),
		FlowType:  (*FlowType)(flow),
	})
}

////todo
////删除os,dst_port列
//改为text

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

func (f *FprintFlow) InsertModel() error {
	return mysql.CreateModel(f)
}

func (f *FprintFlow) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(f, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}

func (f *FprintFlow) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(f, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (f *FprintFlow) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(f, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (f *FprintFlow) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}

func (flow *FprintFlow) CreateModel(flowParam ...interface{}) interface{} {
	flowItemParams := flowParam[0].(*protobuf.FItem)
	flow.Hash = flowItemParams.GetHash()

	sipLittleEndian := util.BytesToLittleEndian(util.BigToBytes(flowItemParams.GetSrcIp()))
	flow.SrcIp = util.IpIntToString(int(sipLittleEndian))

	flow.SrcPort = flowItemParams.GetSrcPort()

	dipLittleEndian := util.BytesToLittleEndian(util.BigToBytes(flowItemParams.GetDstIp()))
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

	//add
	flow.SrcDstPackets = flowItemParams.GetSrc2DstPackets()
	flow.DstSrcPackets = uint64(flowItemParams.GetDst2SrcPackets())
	flow.HostName = flowItemParams.GetHostName()
	flow.Category = flowItemParams.GetCategory()
	flow.HasPassive = flowItemParams.GetHasPassive()
	flow.IatFlowAvg = flowItemParams.GetIatFlowAvg()
	flow.IatFlowStddev = flowItemParams.GetIatFlowStddev()
	flow.DataRatio = flowItemParams.GetDataRatio()
	flow.StrDataRatio = uint8(flowItemParams.GetStrDataRatio())
	flow.PktlenCToSAvg = flowItemParams.GetPktlenCToSAvg()
	flow.PktlenCToSStddev = flowItemParams.GetPktlenCToSStddev()
	flow.PktlenSToCAvg = flowItemParams.GetPktlenSToCAvg()
	flow.PktlenSToCStddev = flowItemParams.GetPktlenSToCStddev()
	flow.TlsClientInfo = flowItemParams.GetTlsClientInfo()
	flow.Ja3c = flowItemParams.GetJa3C()
	return flow
}

func (flow *FprintFlow) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, orderBy interface{}, query interface{}, args ...interface{}) error {

	err := mysql.QueryModelPaginationByWhereCondition(flow, pageIndex, pageSize, totalCount, paginModel, orderBy, query, args...)

	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
