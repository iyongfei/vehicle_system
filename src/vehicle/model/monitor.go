package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

type Monitor struct {
	gorm.Model
	MonitorId  string
	CpuRate    float32
	MemRate    float32
	GatherTime uint32
}

type VehicleMonitorJoinItems struct {
	Monitor

	Path     string
	DiskRate float32
}

func GetVehicleMonitorItems(query string, args ...interface{}) ([]*VehicleMonitorJoinItems, error) {
	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
	if err != nil {
		return nil, fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
	}
	vehicleMonitorItems := []*VehicleMonitorJoinItems{}
	err = vgorm.Debug().
		Table("monitors").
		Select("monitors.*,disks.path,disks.disk_rate").
		Where(query, args...).
		Joins("inner join disks ON monitors.monitor_id = disks.monitor_id").
		Scan(&vehicleMonitorItems).
		Error
	return vehicleMonitorItems, err
}

type VehicleMonitorItemsResponse struct {
	Monitor
	/////////////////////
	VehicleMonitorItemList []Disk
}

func (monitor *Monitor) InsertModel() error {
	return mysql.CreateModel(monitor)
}
func (monitor *Monitor) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(monitor, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (monitor *Monitor) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(monitor, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (monitor *Monitor) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(monitor, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (monitor *Monitor) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (monitor *Monitor) CreateModel(monitorParams ...interface{}) interface{} {
	monitorParam := monitorParams[0].(*protobuf.MonitorInfoParam)
	monitor.CpuRate = monitorParam.GetCpuRate()
	monitor.MemRate = monitorParam.GetMemRate()
	monitor.GatherTime = monitorParam.GetGatherTime()
	return monitor
}

type Disk struct {
	gorm.Model
	MonitorId string
	Path      string
	DiskRate  float32
}

func (disk *Disk) InsertModel() error {
	return mysql.CreateModel(disk)
}
func (disk *Disk) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(disk, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (disk *Disk) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(disk, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (disk *Disk) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(disk, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (disk *Disk) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (disk *Disk) CreateModel(diskParams ...interface{}) interface{} {
	diskParam := diskParams[0].(*protobuf.MonitorInfoParam_DiskOverFlow)
	disk.DiskRate = diskParam.GetDiskRate()
	return disk
}
