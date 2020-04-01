package mysql

import (
	"fmt"
	"vehicle_system/src/vehicle/util"
)

func UpdateModelAllColumns(model interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Save(model).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}

/**
 使用`map`更新多个属性，只会更新这些更改的字段
query:="gw_name = ? and gw_mac = ?"
values:= map[string]interface{}{
	"gw_group":"modify",
	"gw_strategy":"modify",
}
mysql_util.UpdateModelColumnsByInSlice("gw_manage_infos",values,query,"slfjskdlfernn","wjeio")
*/
func UpdateModelByMapTable(table string, values interface{}, query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Table(table).Where(query, args...).Updates(values).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}


/**
model:= &gw_manage.GwManageInfo{}
query:="gw_name = ? and gw_mac = ?"
values:= map[string]interface{}{
	"gw_online_status": gorm.Expr("gw_online_status + ?", 100),
}
mysql_util.UpdateModelSomeColumns(model,values,query,"jj","wjeio")
*/

/**
query:="gw_name = ? and gw_mac = ?"
attrs:= map[string]interface{}{
	"gw_bind_ip":"olo",
	"gw_online_status":2323,
}
mysql_util.UpdateModelSomeColumns(&gw_manage.GwManageInfo{},attrs,query,"jj","wjeio")
*/
func UpdateModelByMapModel(model interface{}, values interface{}, query interface{}, queryArgs ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Debug().Model(model).Where(query, queryArgs...).Updates(values).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}


/**
 更新单个属性（如果更改）
db.Model(&user).Update("name", "hello")
query:="gw_name = ? and gw_mac = ?"
attrs:=[]interface{}{"gw_bind_ip","xxxxxxxmodifyip"}
mysql_util.UpdateModelOneColumn(&gw_manage.GwManageInfo{},attrs,query,"jj","wjeio")
*/
func UpdateModelOneColumn(model interface{},attrs []interface{},query interface{},queryArgs ...interface{}) error{
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Model(model).Where(query,queryArgs...).Update(attrs...).Error
	if err != nil{
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return  nil
}

func UpdateModelByMapModelUnscoped(model interface{}, values interface{}, query interface{}, queryArgs ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Model(model).Unscoped().Where(query, queryArgs...).Updates(values).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}



func UpdateModelOneColumnUnscoped(model interface{},attrs []interface{},query interface{},queryArgs ...interface{}) error{
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Model(model).Unscoped().Where(query,queryArgs...).Update(attrs...).Error
	if err != nil{
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return  nil
}
