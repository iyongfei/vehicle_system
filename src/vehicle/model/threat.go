package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

type Threat struct {
	gorm.Model
	ThreatId       					string
	VehicleId       						string
	SrcIP                         string
	Type                      int//威胁类型\名称
	Content                   string
	Status                    int //威胁状态
	AttactTime                time.Time //威胁产生时间
	IsReaded					bool

	DstIP                         string
	DestType                   int
}


func (u *Threat) InsertModel() error {
	return mysql.CreateModel(u)
}

func (u *Threat) GetModelByCondition(query interface{}, args ...interface{}) (error,bool) {
	err,recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(u,query,args...)
	if err!=nil{
		return err,true
	}
	if recordNotFound{
		return nil,true
	}
	return nil,false
}
func (u *Threat) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(u,values,query,queryArgs...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (u *Threat) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}

func (u *Threat) GetModelListByCondition(model interface{},query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
/**

func QueryModelPaginationByWhereCondition(model interface{},pageNum int, pageCount int, totalCount *int,
	paginModel interface{}, query interface{}, args ...interface{}) error {
 */

func (u *Threat) GetModelPaginationByCondition(pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, query interface{}, args ...interface{})(error){

	err := mysql.QueryModelPaginationByWhereCondition(u,pageIndex,pageSize,totalCount,paginModel,query,args...)

	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}



func (threat *Threat)  CreateModel(threatParams ...interface{})  interface{}{
	threatParam := threatParams[0].(*protobuf.ThreatParam_Item)
	threat.SrcIP = threatParam.GetSrcIp()
	threat.Type = int(threatParam.GetThreatType())
	threat.Content = threatParam.GetContent()
	threat.Status = int(threatParam.GetThreatStatus())
	if threatParam.GetAttactTime() == 0{
		threat.AttactTime = time.Now()
	}else{
		threat.AttactTime = util.StampUnix2Time(int64(threatParam.GetAttactTime()))
	}

	threat.IsReaded = false
	threat.DstIP = threatParam.GetDstIp()
	return threat
}

