package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
	"vehicle_system/src/vehicle_script/tool"
)

/*************************************Sample*******************************************/

type Sample struct {
	gorm.Model
	SampleId   string    // 采集样本id
	StartTime  time.Time //启动时间
	RemainTime uint32    //剩余时间
	TotalTime  uint32    //剩余时间

	Status  uint8  //采集状态
	Timeout uint32 //超时时间

	Name      string //采集名称
	Introduce string //采集说明

	Check uint8

	VehicleId     string
	StudyOriginId string
}

func (sample *Sample) InsertModel() error {
	return mysql.CreateModel(sample)
}
func (sample *Sample) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(sample, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (sample *Sample) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(sample, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (sample *Sample) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (sample *Sample) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	return nil
}
func (sample *Sample) CreateModel(sampleParams ...interface{}) interface{} {
	sampleParam := sampleParams[0].(*protobuf.SampleParam)

	sample.StartTime = tool.StampUnix2Time(int64(sampleParam.GetStartTime()))
	sample.RemainTime = sampleParam.GetTimeRemain()
	sample.Status = uint8(sampleParam.GetStatus())
	sample.Timeout = sampleParam.GetTimeout()

	return sample
}

/*************************************SampleItem*******************************************/

type SampleItem struct {
	gorm.Model
	SampleItemId string //id

	SampleId string

	SrcMac  string //源mac sm
	SrcIp   string //源ip sip
	SrcPort uint32 //源端口 sp

	DstIp   string ///////////////目标ip  dip
	DstPort uint32 //目标端口  dp
	Url     string /////////////////目标url u

	FetchTime time.Time //访问时间tm
}

func (item *SampleItem) InsertModel() error {
	return mysql.CreateModel(item)
}
func (item *SampleItem) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(item, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (item *SampleItem) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(item, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (item *SampleItem) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (item *SampleItem) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (item *SampleItem) CreateModel(sampleItemParams ...interface{}) interface{} {
	sampleItemParam := sampleItemParams[0].(*protobuf.SampleParam_Item)
	item.SrcMac = sampleItemParam.GetSm()
	item.SrcIp = sampleItemParam.GetSip()
	item.SrcPort = sampleItemParam.GetSp()
	item.DstIp = sampleItemParam.GetDip()
	item.DstPort = sampleItemParam.GetDp()
	item.Url = sampleItemParam.GetU()
	//item.FetchTime = sampleItemParam.GetTm()
	item.FetchTime = time.Now()

	return item
}
