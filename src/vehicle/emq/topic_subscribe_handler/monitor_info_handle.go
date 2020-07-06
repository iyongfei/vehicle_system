package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/service/push"
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
	//create
	var diskList []*model.Disk

	diskItems := monitorParam.GetDiskItem()
	for _, diskItem := range diskItems {

		logger.Logger.Print("%s handle_monitor:%+v", util.RunFuncName(), *diskItem)
		logger.Logger.Info("%s handle_monitor:%+v", util.RunFuncName(), *diskItem)

		disk := &model.Disk{
			MonitorId:  vehicleId,
			Path:       diskItem.Path,
			GatherTime: monitorParam.GetGatherTime(),
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
		diskList = append(diskList, disk)
	}

	//redis
	redisInfoParam := monitorParam.GetRedisInfo()
	var redisInfo *model.RedisInfo
	if redisInfoParam.GetActive() {
		logger.Logger.Print("%s handle_monitor:%+v", util.RunFuncName(), *redisInfoParam)
		logger.Logger.Info("%s handle_monitor:%+v", util.RunFuncName(), *redisInfoParam)

		redisInfo = &model.RedisInfo{
			MonitorId:  vehicleId,
			GatherTime: monitorParam.GatherTime,
		}
		redisInfoModelBase := model_base.ModelBaseImpl(redisInfo)

		_, redisRecordNotFound := redisInfoModelBase.GetModelByCondition(
			"gather_time = ?", []interface{}{redisInfo.GatherTime}...)

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
	var vhaloInfo *model.VhaloNets
	if vhaloNetsParam.GetActive() {
		logger.Logger.Print("%s handle_monitor:%+v", util.RunFuncName(), *vhaloNetsParam)
		logger.Logger.Info("%s handle_monitor:%+v", util.RunFuncName(), *vhaloNetsParam)

		vhaloInfo = &model.VhaloNets{
			MonitorId:  vehicleId,
			GatherTime: monitorParam.GatherTime,
		}
		vhaloModelBase := model_base.ModelBaseImpl(vhaloInfo)

		_, redisRecordNotFound := vhaloModelBase.GetModelByCondition(
			"gather_time = ?", []interface{}{vhaloInfo.GatherTime}...)

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

	//会话状态

	pushActionTypeName := protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]
	pushVehicleid := vehicleId
	pushData := map[string]interface{}{
		"disks": diskList,
		"redis": redisInfo,
		"vhalo": vhaloInfo,
	}

	fPushData := push.CreatePushData(pushActionTypeName, pushVehicleid, pushData)

	push.GetPushervice().SetPushData(fPushData).Write()

	return nil
}
