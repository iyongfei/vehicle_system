package tools

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle_script/tool"
)

const (
	mysqlUser   = "root"
	mysqlPwd    = "root"
	mysqlPort   = 3306
	mysqlDbname = "vehicle"

	maxIdleConns = 10
	maxOpenConns = 100
)

var GormDb *gorm.DB

var user_name string
var password string
var db_name string
var mysql_port string

func initConfIni() {
	apiConfigMap := tool.InitConfig("conf.ini")
	user_name = apiConfigMap["user_name"]
	password = "root"
	db_name = apiConfigMap["db_name"]
	mysql_port = apiConfigMap["mysql_port"]

	fmt.Println("ini_user_name-->", user_name, ",password-->", password, ",db_name-->", db_name, ",mysql_port-->", mysql_port)

}

//func GET() *gorm.DB {
//	GormDb, _ = gorm.Open("mysql", getConnectParams())
//	return GormDb
//}

func GetMysqlInstance() {
	//initConfIni()

	var err error
	GormDb, err = gorm.Open("mysql", getConnectParams())

	fmt.Println(err)
	GormDb.LogMode(true)

	GormDb.DB().SetMaxIdleConns(maxIdleConns)
	GormDb.DB().SetMaxOpenConns(maxOpenConns)

}

func CreateModel(model interface{}) error {
	err := GormDb.Create(model).Error
	return err
}

func getConnectParams() string {
	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlUser, mysqlPwd, mysqlPort, mysqlDbname)
}
func QueryPluckWhere(model interface{}, column string, result interface{}, query interface{}, args ...interface{}) error {
	err := GormDb.Model(model).Where(query, args...).Pluck(column, result).Error
	return err
}

func QueryPluck(model interface{}, column string, result interface{}) error {
	err := GormDb.Model(model).Pluck(column, result).Error
	return err
}

func QueryRand(model interface{}, column string, result interface{}, limit int, order string) error {
	err := GormDb.Model(model).Order(order).Limit(limit).Pluck(column, result).Error
	return err
}

func UpdateModelOneColumn(model interface{}, attrs []interface{}, query interface{}, queryArgs ...interface{}) error {
	err := GormDb.Model(model).Where(query, queryArgs...).Update(attrs...).Error
	return err
}

func QueryFirstModel(model interface{}, query interface{}, args ...interface{}) error {
	err := GormDb.Where(query, args...).First(model).Error
	return err
}

func QueryFirstModelRecord(model interface{}, query interface{}, args ...interface{}) bool {
	boolFlag := GormDb.Where(query, args...).First(model).RecordNotFound()
	return boolFlag
}

/**
按gw_infos表的gw_id排序，gw_id是字符串类型的数值
SELECT gw_id FROM gw_infos ORDER BY CAST(gw_id AS SIGNED)
*/
func OrderByVarcharConvertNum(model interface{}, column string, orderBy interface{}) error {
	err := GormDb.Select(column).Limit(1).Order(orderBy).Find(model).Error
	return err
}
