package mysql

import (
	"fmt"
	"vehicle_system/src/vehicle/util"
)

/**
将结果扫描到另一个结构中
slice
struct
 */
func QueryRawsqlScanStruct(sql string,arg interface{},obj interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Raw(sql, arg).Scan(obj).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}

/**
将结果扫描到变量
err := mysql_util.QueryRawSqlScanVariable("area_groups",
		"area_code",&groupId,"area_name = ?",conf.Ungrouped)
 */

func QueryRawsqlScanVariable(tname string,selectColumn string,variable  interface{},query string,args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Table(tname).Where(query, args...).Select(selectColumn).Row().Scan(variable) // (*sql.Row)
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}


/**
获取记录

//func PaginationT(whichPageNum int, perPageNum int, count *int, totalModel interface{},
//	paginModel interface{}, query interface{}, args interface{}) error {
//	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
//	if err != nil {
//		return errors.Wrapf(err, "Pagination open gorm error %v", err)
//	}
//
//	//这里判断的不严谨
//	if whichPageNum < 1 && perPageNum < 1 {
//		return err
//	}
//
//	err = vgorm.Offset((whichPageNum-1)*perPageNum).Limit(perPageNum).Where(query, args).Find(paginModel).
//		Offset(-1).Find(totalModel).Count(count).Error
//
//	return err
//}
 */
//db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)

func QueryModelPaginationByWhereCondition(model interface{},pageIndex int, pageSize int, totalCount *int,
	paginModel interface{}, query interface{}, args ...interface{}) error {

	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}

	err = vgorm.Offset((pageIndex-1)*pageSize).Limit(pageSize).Where(query, args...).Find(paginModel).
		Offset(-1).Model(model).Count(totalCount).Error

	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}



/**
获取记录
 */
func QueryModelRecordsByWhereCondition(models interface{}, query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Where(query, args...).Find(models).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}
/**
获取某记录
 */
func QueryModelOneRecordByWhereCondition(model interface{}, query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Where(query, args...).First(model).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}
/**
检查是否有记录
err = vgorm.Where(query, args...).Find(models).Error
*/
func QueryModelOneRecordIsExistByWhereCondition(model interface{}, query interface{}, args ...interface{}) (error, bool) {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error()),false
	}
	isExist := vgorm.Where(query, args...).First(model).RecordNotFound()

	return nil, isExist
}


/**
var ages []int64
db.Find(&users).Pluck("age", &ages)
 */

func QueryPluckByWhere(model interface{}, column string, result interface{},query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Where(query,args...).Find(model).Pluck(column, result).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}


func QueryPluckByModelWhere(model interface{}, column string, result interface{},query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}

	err = vgorm.Model(model).Where(query,args...).Pluck(column, result).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}

/**
Offset((whichPageNum-1) * perPageNum).
Limit(perPageNum).
 */

func QueryModelOneRecordOffsetCount(model interface{}, count int,query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Where(query, args...).Offset(count).First(model).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}






/**
models:=&[]gw_manage.GwManageInfo{}
slice:=[]string{"gw_id","gw_name"}
mysql_util.QueryModelRecordsBySelect(models,slice)
 */
func QueryModelRecordsBySelect(models interface{}, slice interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Select(slice).Find(models).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}

func GetModelCountByWhere(model interface{}, count interface{}, query interface{}, args ...interface{}) error {
	vgorm,err := GetMysqlInstance().GetMysqlDB()
	if err!= nil{
		return fmt.Errorf("%s open grom err:%v",util.RunFuncName(),err.Error())
	}
	err = vgorm.Model(model).Where(query, args...).Count(count).Error
	if err != nil {
		return fmt.Errorf("%s err %v",util.RunFuncName(),err.Error())
	}
	return nil
}























