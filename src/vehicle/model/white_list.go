package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/util"
)

type WhiteList struct {
	gorm.Model
	WhiteListId string
	DestIp       string///////////////目标ip  dip
	Url          string/////////////////目标url u

	SourceMac    string//源mac sm
	SourceIp     string//源ip sip
}


//func (w *WhiteList) GetModelListByCondition(model interface{},query interface{}, args ...interface{}) (error,bool) {
//	return nil,false
//}


func (w *WhiteList) InsertModel(model interface{}) error {
	return mysql.CreateModel(model)
}

func (w *WhiteList) GetModelByCondition(model interface{},query interface{}, args ...interface{}) (error,bool) {
	err,recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(model,query,args...)
	if err!=nil{
		return err,true
	}
	if recordNotFound{
		return nil,true
	}
	return nil,false
}
func (w *WhiteList) UpdateModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (w *WhiteList) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (w *WhiteList) GetModelListByCondition(model interface{},query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}


























