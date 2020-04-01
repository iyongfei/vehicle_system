package model

import (
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
)

type User struct {
	gorm.Model//创建时间在
	UserId  string
	UserName string//管理员账号
	Password string
	Type int //管理员类型

	Email string//管理员邮箱
	Phone string//手机
	Marks string//备注
}

func (u *User) InsertModel(model interface{}) error {
	return mysql.CreateModel(model)
}

func (u *User) GetModelByCondition(model interface{},query interface{}, args ...interface{}) (error,bool) {
	err,recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(model,query,args...)
	if err!=nil{
		return err,true
	}
	if recordNotFound{
		return nil,true
	}
	return nil,false
}
func (u *User) UpdateModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}
func (u *User) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	return nil
}

func (u *User) GetModelListByCondition(model interface{},query interface{}, args ...interface{}) (error) {
	return nil
}
