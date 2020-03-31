package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/logger"
)

var connectOnce sync.Once
var GormDb *gorm.DB

type MysqlConn struct {}

var MysqlConnInstance *MysqlConn

func GetMysqlInstance() *MysqlConn{
	connectOnce.Do(NewMysqlConn)
	return MysqlConnInstance
}
func NewMysqlConn()  {
	if MysqlConnInstance == nil{
		MysqlConnInstance = new(MysqlConn)
	}
}

func (m *MysqlConn) InitDataBase()(*gorm.DB,error){
	var err error
	GormDb, err = gorm.Open("mysql",getConnectParams())
	//GormDb.LogMode(true)
	if err != nil {
		logger.Logger.Error("gorm open error:%v", err)
		logger.Logger.Print("gorm open error:%v", err)
		log.Fatalf("gorm open error:%v\n", err)
	}
	GormDb.DB().SetMaxIdleConns(conf.MaxIdleConns)
	GormDb.DB().SetMaxOpenConns(conf.MaxOpenConns)
	return GormDb,err
}

func getConnectParams() string{
	return fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.MysqlUser, conf.MysqlPassword, conf.MysqlPort, conf.MysqlDbname)
}
func (m *MysqlConn)GetMysqlDB() (db_con *gorm.DB,err error) {
	if GormDb == nil{
		GormDb, err = gorm.Open("mysql",getConnectParams())

		if err != nil {
			logger.Logger.Error("gorm open error:%v", err)
			logger.Logger.Print("gorm open error:%v", err)
			log.Fatalf("gorm open error:%v\n", err)
		}
		return GormDb,err
	}
	return GormDb,err
}
