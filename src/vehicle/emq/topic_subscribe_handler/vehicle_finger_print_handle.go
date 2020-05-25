package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/mac"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleFingerPrint(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	fingerprintParam := &protobuf.FingerprintParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), fingerprintParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicleFingerPrint err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal vehicleFingerPrint err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal vehicleFingerPrint err:%s", util.RunFuncName(), err.Error())
	}
	detect := fingerprintParam.GetActive()
	passive := fingerprintParam.GetPassive()

	tradeMark := mac.GetOrgByMAC(fingerprintParam.GetMac())
	fTradeMark := util.RrgsTrim(tradeMark)

	fprintInfo := &model.FprintInfo{
		FprintInfoId: util.RandomString(32),
		DeviceMac:    fingerprintParam.GetMac(),
		VehicleId:    vehicleId,
		Os:           detect.GetOs(),
		TradeMark:    fTradeMark,
		DstPort:      passive.GetDstPort(),
	}
	detectInfoModelBase := model_base.ModelBaseImpl(fprintInfo)

	_, dRecordNotFound := detectInfoModelBase.GetModelByCondition("device_mac = ?",
		[]interface{}{fprintInfo.DeviceMac}...)

	detectInfoModelBase.CreateModel(fingerprintParam)
	if dRecordNotFound {
		if err := detectInfoModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert fprintInfo err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		attrs := map[string]interface{}{
			"os":       fprintInfo.Os,
			"dst_port": fprintInfo.DstPort,
		}
		if err := detectInfoModelBase.UpdateModelsByCondition(attrs, "device_mac = ?", fprintInfo.DeviceMac); err != nil {
			return fmt.Errorf("%s update vehicle finger print err:%s", util.RunFuncName(), err.Error())
		}
	}

	return nil
}
