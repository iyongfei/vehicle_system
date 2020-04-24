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

func HandleVehiclePortMap(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	portMapParams := &protobuf.PortRedirectParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), portMapParams)
	if err != nil {
		logger.Logger.Print("%s unmarshal portmaps Param err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal portmaps err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal portmaps err:%s", util.RunFuncName(), err.Error())
	}

	for _, portMapParam := range portMapParams.GetPortRedirect() {
		srcPort := portMapParam.GetSrcPort()
		portMapModel := &model.PortMap{
			PortMapId: util.RandomString(32),
			VehicleId: vehicleId,
			SrcPort:   srcPort,
		}
		modelBase := model_base.ModelBaseImpl(portMapModel)
		_, recordNotFound := modelBase.GetModelByCondition("src_port = ? and vehicle_id = ?",
			[]interface{}{srcPort, vehicleId}...)

		modelBase.CreateModel(portMapParam)

		logger.Logger.Print("%s unmarshal portmaps Param portMapModel:%+v", util.RunFuncName(), portMapModel)

		if recordNotFound {
			if err := modelBase.InsertModel(); err != nil {
				//return fmt.Errorf("%s insert flow err:%s",util.RunFuncName(),err.Error())
				continue
			}
		} else {
			//update
			//更新 排除VehicleId,Name,ProtectStatus,LeaderId
			attrs := map[string]interface{}{
				"dst_port":      portMapModel.DstPort,
				"dst_ip":        portMapModel.DstIp,
				"swith":         portMapModel.Switch,
				"protocol_type": portMapModel.ProtocolType,
			}
			if err := modelBase.UpdateModelsByCondition(attrs, "src_port = ? and vehicle_id = ?",
				[]interface{}{srcPort, vehicleId}...); err != nil {
				//return fmt.Errorf("%s update flow err:%s",util.RunFuncName(),err.Error())
				continue
			}
		}
	}
	return nil
}
