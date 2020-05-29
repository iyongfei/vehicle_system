package tools

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"venus_halo/common_util/log_util"
	//"venus_halo/common_util"
	//"venus_halo/common_util/log_util"
)

const (
	mysqlUser = "root"
	mysqlPwd  = "root"
	//mysqlPort = 3306
	mysqlPort    = 33066
	mysqlDbname  = "vehicle"
	maxIdleConns = 10
	maxOpenConns = 100
)

var GormDb *gorm.DB

func GetMysqlInstance() *gorm.DB {
	if GormDb == nil {
		GormDb, _ = gorm.Open("mysql", getConnectParams())
		GormDb.LogMode(true)

		GormDb.DB().SetMaxIdleConns(maxIdleConns)
		GormDb.DB().SetMaxOpenConns(maxOpenConns)
	}
	return GormDb
}

func CreateModel(model interface{}) error {
	vgorm := GetMysqlInstance()
	err := vgorm.Create(model).Error
	return err
}

func getConnectParams() string {
	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlUser, mysqlPwd, mysqlPort, mysqlDbname)
}

//GormDb.Exec("truncate table area_groups;")
func Exec(sql string) {
	GetMysqlInstance().Exec(sql)
}

func QueryPluckWhere(model interface{}, column string, result interface{}, query interface{}, args ...interface{}) error {
	err := GetMysqlInstance().Model(model).Where(query, args...).Pluck(column, result).Error
	return err
}

func QueryPluck(model interface{}, column string, result interface{}) error {
	err := GetMysqlInstance().Model(model).Pluck(column, result).Error
	return err
}

func QueryRand(model interface{}, column string, result interface{}, limit int, order string) error {
	err := GetMysqlInstance().Model(model).Order(order).Limit(limit).Pluck(column, result).Error
	return err
}

func UpdateModelOneColumn(model interface{}, attrs []interface{}, query interface{}, queryArgs ...interface{}) error {
	err := GetMysqlInstance().Model(model).Where(query, queryArgs...).Update(attrs...).Error
	return err
}

func QueryFirstModel(model interface{}, query interface{}, args ...interface{}) error {
	err := GetMysqlInstance().Where(query, args...).First(model).Error
	return err
}

/**
按gw_infos表的gw_id排序，gw_id是字符串类型的数值
SELECT gw_id FROM gw_infos ORDER BY CAST(gw_id AS SIGNED)
*/
func OrderByVarcharConvertNum(model interface{}, column string, orderBy interface{}) error {
	err := GetMysqlInstance().Select(column).Limit(1).Order(orderBy).Find(model).Error
	return err
}
