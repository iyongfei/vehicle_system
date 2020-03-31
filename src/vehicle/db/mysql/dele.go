package mysql

import (
	"fmt"
	"vehicle_system/src/vehicle/util"
)

//软删除A
/**
query:="gw_name = ? and gw_mac = ?"
mysql_util.SoftDeleteGwManageInfoA(&gw_manage.GwManageInfo{},query,"jj","wjeio")
 */
func SoftDeleteModelA(model interface{},query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Where(query, args...).Delete(model).Error
	if err!= nil{
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}
//软删除B
/**

query:="gw_name = ? and gw_mac = ?"
mysql_util.SoftDeleteGwManageInfoB(&gw_manage.GwManageInfo{},query,"slfjgskdlfe","wjeio")
 */
func SoftDeleteModelB(model interface{},where ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Delete(model, where...).Error
	if err!= nil{
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}


//硬删除
/**
query:="gw_name = ? and gw_mac = ?"
mysql_util.HardDeleteGwManageInfoA(&gw_manage.GwManageInfo{},query,"slfjgskdlfe","wjeio")
 */
func HardDeleteModelA(model interface{},where ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err =vgorm.Unscoped().Delete(model, where...).Error
	if err!= nil{
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}
/**
query:="gw_name = ? and gw_mac = ?"
mysql_util.HardDeleteGwManageInfoB(&gw_manage.GwManageInfo{},query,"jj","wjeio")
*/
func HardDeleteModelB(model interface{},query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Unscoped().Where(query, args...).Delete(model).Error
	if err!= nil{
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}


/**
db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
 */
func HardDeleteModelInSlice(model interface{},query interface{}, args interface{})error{
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Unscoped().Where(query, args).Delete(model).Error
	if err!= nil{
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}
