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

func HandleMonitorInfo(vehicleResult protobuf.GWResult) error {
	//parse
	monitorParam := &protobuf.MonitorInfoParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), monitorParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal monitorParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal monitorParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal monitorParam err:%s", util.RunFuncName(), err.Error())
	}
	//vehicleId
	vehicleId := vehicleResult.GetGUID()

	logger.Logger.Print("%s unmarshal monitorParam:%+v", util.RunFuncName(), monitorParam)
	logger.Logger.Info("%s unmarshal monitorParam:%+v", util.RunFuncName(), monitorParam)
	//create
	monitor := &model.Monitor{
		MonitorId: vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(monitor)

	_, recordNotFound := modelBase.GetModelByCondition("monitor_id = ?", monitor.MonitorId)

	modelBase.CreateModel(monitorParam)
	if recordNotFound {
		if err := modelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert monitor err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		//更新 排除VehicleId,Name,ProtectStatus,LeaderId
		attrs := map[string]interface{}{
			"cpu_rate":    monitor.CpuRate,
			"mem_rate":    monitor.MemRate,
			"gather_time": monitor.GatherTime,
		}
		if err := modelBase.UpdateModelsByCondition(attrs, "monitor_id = ?", monitor.MonitorId); err != nil {
			return fmt.Errorf("%s update monitor err:%s", util.RunFuncName(), err.Error())
		}
	}

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

	return nil
}
