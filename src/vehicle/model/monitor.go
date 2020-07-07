package model

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/util"
)

//type Monitor struct {
//	gorm.Model
//	MonitorId  string
//	CpuRate    float32
//	MemRate    float32
//	GatherTime uint32
//}

//type VehicleMonitorJoinItems struct {
//	Monitor
//
//	Path     string
//	DiskRate float32
//}
//
//func GetVehicleMonitorItems(query string, args ...interface{}) ([]*VehicleMonitorJoinItems, error) {
//	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
//	if err != nil {
//		return nil, fmt.Errorf("%s open grom err:%v", util.RunFuncName(), err.Error())
//	}
//	vehicleMonitorItems := []*VehicleMonitorJoinItems{}
//	err = vgorm.Debug().
//		Table("monitors").
//		Select("monitors.*,disks.path,disks.disk_rate").
//		Where(query, args...).
//		Joins("inner join disks ON monitors.monitor_id = disks.monitor_id").
//		Scan(&vehicleMonitorItems).
//		Error
//	return vehicleMonitorItems, err
//}

//
//func (monitor *Monitor) InsertModel() error {
//	return mysql.CreateModel(monitor)
//}
//func (monitor *Monitor) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
//	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(monitor, query, args...)
//	if err != nil {
//		return err, true
//	}
//	if recordNotFound {
//		return nil, true
//	}
//	return nil, false
//}
//func (monitor *Monitor) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
//	err := mysql.UpdateModelByMapModel(monitor, values, query, queryArgs...)
//	if err != nil {
//		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
//	}
//	return nil
//}
//func (monitor *Monitor) DeleModelsByCondition(query interface{}, args ...interface{}) error {
//	err := mysql.HardDeleteModelB(monitor, query, args...)
//	if err != nil {
//		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
//	}
//	return nil
//}
//func (monitor *Monitor) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
//	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
//	if err != nil {
//		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
//	}
//	return nil
//}
//func (monitor *Monitor) CreateModel(monitorParams ...interface{}) interface{} {
//	//monitorParam := monitorParams[0].(*protobuf.MonitorInfoParam)
//	//monitor.CpuRate = monitorParam.GetCpuRate()
//	//monitor.MemRate = monitorParam.GetMemRate()
//	//monitor.GatherTime = monitorParam.GetGatherTime()
//	return monitor
//}

///////////////////////////////////////Disk////////////////////////////////////////////////////

type VehicleMonitorItemsResponse struct {
	/////////////////////
	Disks     []*Disk
	RedisInfo RedisInfo
	VhaloNets VhaloNets
}

type Disk struct {
	gorm.Model
	MonitorId  string
	Path       string
	DiskRate   float32
	GatherTime uint64
}

func (tmp *Disk) MarshalJSON() ([]byte, error) {
	type tempType Disk
	return json.Marshal(&struct {
		CreatedAt int64
		*tempType
	}{
		CreatedAt: tmp.CreatedAt.Unix(),
		tempType:  (*tempType)(tmp),
	})
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

	parseDiskRate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", diskParam.GetDiskRate()), 64)

	disk.DiskRate = float32(parseDiskRate)
	return disk
}

///////////////////////////////////////RedisInfo////////////////////////////////////////////////////
type RedisInfo struct {
	gorm.Model
	MonitorId  string
	Active     bool
	CpuRate    float32
	MemRate    float32
	Mem        uint64
	GatherTime uint64
}

func (tmp *RedisInfo) MarshalJSON() ([]byte, error) {
	type tempType RedisInfo
	return json.Marshal(&struct {
		CreatedAt int64
		*tempType
	}{
		CreatedAt: tmp.CreatedAt.Unix(),
		tempType:  (*tempType)(tmp),
	})
}

func (redisInfo *RedisInfo) InsertModel() error {
	return mysql.CreateModel(redisInfo)
}
func (redisInfo *RedisInfo) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(redisInfo, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (redisInfo *RedisInfo) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(redisInfo, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (redisInfo *RedisInfo) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(redisInfo, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (redisInfo *RedisInfo) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (redisInfo *RedisInfo) CreateModel(redisParams ...interface{}) interface{} {
	redisParam := redisParams[0].(*protobuf.MonitorInfoParam_RedisInfo)

	redisInfo.Active = redisParam.GetActive()

	parseCpuRate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", redisParam.GetCpuRate()), 64)
	parseMemRate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", redisParam.GetMemRate()), 64)

	redisInfo.CpuRate = float32(parseCpuRate)
	redisInfo.MemRate = float32(parseMemRate)
	redisInfo.Mem = redisParam.GetMem()
	return redisInfo
}

///////////////////////////////////////VhaloNets////////////////////////////////////////////////////
type VhaloNets struct {
	gorm.Model
	MonitorId  string
	Active     bool
	CpuRate    float32
	MemRate    float32
	Mem        uint64
	GatherTime uint64
}

func (tmp *VhaloNets) MarshalJSON() ([]byte, error) {
	type tempType VhaloNets
	return json.Marshal(&struct {
		CreatedAt int64
		*tempType
	}{
		CreatedAt: tmp.CreatedAt.Unix(),
		tempType:  (*tempType)(tmp),
	})
}

func (vhaloNets *VhaloNets) InsertModel() error {
	return mysql.CreateModel(vhaloNets)
}
func (vhaloNets *VhaloNets) GetModelByCondition(query interface{}, args ...interface{}) (error, bool) {
	err, recordNotFound := mysql.QueryModelOneRecordIsExistByWhereCondition(vhaloNets, query, args...)
	if err != nil {
		return err, true
	}
	if recordNotFound {
		return nil, true
	}
	return nil, false
}
func (vhaloNets *VhaloNets) UpdateModelsByCondition(values interface{}, query interface{}, queryArgs ...interface{}) error {
	err := mysql.UpdateModelByMapModel(vhaloNets, values, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (vhaloNets *VhaloNets) DeleModelsByCondition(query interface{}, args ...interface{}) error {
	err := mysql.HardDeleteModelB(vhaloNets, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (vhaloNets *VhaloNets) GetModelListByCondition(model interface{}, query interface{}, args ...interface{}) error {
	err := mysql.QueryModelRecordsByWhereCondition(model, query, args...)
	if err != nil {
		return fmt.Errorf("%s err %s", util.RunFuncName(), err.Error())
	}
	return nil
}
func (vhaloNets *VhaloNets) CreateModel(vhaloParams ...interface{}) interface{} {
	vhaloParam := vhaloParams[0].(*protobuf.MonitorInfoParam_VHaloNets)
	vhaloNets.Active = vhaloParam.GetActive()

	parseCpuRate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", vhaloParam.GetCpuRate()), 64)
	parseMemRate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", vhaloParam.GetMemRate()), 64)
	vhaloNets.CpuRate = float32(parseCpuRate)
	vhaloNets.MemRate = float32(parseMemRate)
	vhaloNets.Mem = vhaloParam.GetMem()
	return vhaloNets
}
