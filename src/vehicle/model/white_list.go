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



func (w *WhiteList) InsertModel() error {
	return mysql.CreateModel(w)
}



func (w *WhiteList) GetModelByCondition(query interface{}, args ...interface{}) (error,bool) {
	err,recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(w,query,args...)
	if err!=nil{
		return err,true
	}
	if recordNotFound{
		return nil,true
	}
	return nil,false
}
func (w *WhiteList) UpdateModelsByCondition( values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(w,values,query,queryArgs...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}


func (w *WhiteList) DeleModelsByCondition(query interface{}, args ...interface{}) error {
err := mysql.HardDeleteModelB(w,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}
func (w *WhiteList) GetModelListByCondition(model interface{},query interface{}, args ...interface{}) (error) {
	err := mysql.QueryModelRecordsByWhereCondition(model,query,args...)
	if err!=nil{
		return fmt.Errorf("%s err %s",util.RunFuncName(),err.Error())
	}
	return nil
}

func (w *WhiteList)  CreateModel(params ...interface{})  interface{}{

	return nil
}








