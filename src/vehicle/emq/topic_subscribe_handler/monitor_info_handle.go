package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

func HandleMonitorInfo(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	monitorParam := &protobuf.MonitorInfoParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), monitorParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal monitorParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal monitorParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal monitorParam err:%s", util.RunFuncName(), err.Error())
	}
	//vehicleId
	logger.Logger.Print("%s unmarshal monitorParam:%+v", util.RunFuncName(), monitorParam)
	logger.Logger.Info("%s unmarshal monitorParam:%+v", util.RunFuncName(), monitorParam)
	//create
	//monitor := &model.Monitor{
	//	MonitorId: vehicleId,
	//}
	//modelBase := model_base.ModelBaseImpl(monitor)
	//
	//_, recordNotFound := modelBase.GetModelByCondition("monitor_id = ?", monitor.MonitorId)
	//
	//modelBase.CreateModel(monitorParam)
	//if recordNotFound {
	//	if err := modelBase.InsertModel(); err != nil {
	//		return fmt.Errorf("%s insert monitor err:%s", util.RunFuncName(), err.Error())
	//	}
	//} else {
	//	//更新 排除VehicleId,Name,ProtectStatus,LeaderId
	//	attrs := map[string]interface{}{
	//		"cpu_rate":    monitor.CpuRate,
	//		"mem_rate":    monitor.MemRate,
	//		"gather_time": monitor.GatherTime,
	//	}
	//	if err := modelBase.UpdateModelsByCondition(attrs, "monitor_id = ?", monitor.MonitorId); err != nil {
	//		return fmt.Errorf("%s update monitor err:%s", util.RunFuncName(), err.Error())
	//	}
	//}

	diskItems := monitorParam.GetDiskItem()
	for _, diskItem := range diskItems {
		disk := &model.Disk{
			MonitorId: vehicleId,
			Path:      diskItem.Path,
		}
		diskModelBase := model_base.ModelBaseImpl(disk)

		_, diskRecordNotFound := diskModelBase.GetModelByCondition(
			"monitor_id = ? and path = ?", []interface{}{disk.MonitorId, disk.Path}...)

		diskModelBase.CreateModel(diskItem)

		if diskRecordNotFound {
			if err := diskModelBase.InsertModel(); err != nil {
				return fmt.Errorf("%s insert disk err:%s", util.RunFuncName(), err.Error())
			}
		} else {
			attrs := map[string]interface{}{
				"disk_rate": disk.DiskRate,
			}
			if err := diskModelBase.UpdateModelsByCondition(attrs,
				"monitor_id = ? and path = ?", []interface{}{disk.MonitorId, disk.Path}...); err != nil {
				return fmt.Errorf("%s update monitor err:%s", util.RunFuncName(), err.Error())
			}
		}
	}

	//redis
	redisInfoParam := monitorParam.GetRedisInfo()
	if redisInfoParam.GetActive() {
		redisInfo := &model.RedisInfo{
			MonitorId:  vehicleId,
			GatherTime: monitorParam.GatherTime,
		}
		redisInfoModelBase := model_base.ModelBaseImpl(redisInfo)

		_, redisRecordNotFound := redisInfoModelBase.GetModelByCondition(
			"monitor_id = ?", []interface{}{redisInfo.MonitorId}...)

		redisInfoModelBase.CreateModel(redisInfoParam)
		if redisRecordNotFound {
			if err := redisInfoModelBase.InsertModel(); err != nil {
				return fmt.Errorf("%s insert redis_info err:%s", util.RunFuncName(), err.Error())
			}
		} else {
			attrs := map[string]interface{}{
				"active":      redisInfo.Active,
				"cpu_rate":    redisInfo.CpuRate,
				"mem_rate":    redisInfo.MemRate,
				"mem":         redisInfo.Mem,
				"gather_time": redisInfo.GatherTime,
			}
			if err := redisInfoModelBase.UpdateModelsByCondition(attrs,
				"monitor_id = ?", []interface{}{redisInfo.MonitorId}...); err != nil {
				return fmt.Errorf("%s update redis_info err:%s", util.RunFuncName(), err.Error())
			}
		}
	}

	//vhalonets
	vhaloNetsParam := monitorParam.GetVhaloInfo()
	if vhaloNetsParam.GetActive() {
		vhaloInfo := &model.VhaloNets{
			MonitorId:  vehicleId,
			GatherTime: monitorParam.GatherTime,
		}
		vhaloModelBase := model_base.ModelBaseImpl(vhaloInfo)

		_, redisRecordNotFound := vhaloModelBase.GetModelByCondition(
			"monitor_id = ?", []interface{}{vhaloInfo.MonitorId}...)

		vhaloModelBase.CreateModel(vhaloNetsParam)
		if redisRecordNotFound {
			if err := vhaloModelBase.InsertModel(); err != nil {
				return fmt.Errorf("%s insert vhalo_info err:%s", util.RunFuncName(), err.Error())
			}
		} else {
			attrs := map[string]interface{}{
				"active":      vhaloInfo.Active,
				"cpu_rate":    vhaloInfo.CpuRate,
				"mem_rate":    vhaloInfo.MemRate,
				"mem":         vhaloInfo.Mem,
				"gather_time": vhaloInfo.GatherTime,
			}
			if err := vhaloModelBase.UpdateModelsByCondition(attrs,
				"monitor_id = ?", []interface{}{vhaloInfo.MonitorId}...); err != nil {
				return fmt.Errorf("%s update vhalo_info err:%s", util.RunFuncName(), err.Error())
			}
		}
	}

	return nil
}
