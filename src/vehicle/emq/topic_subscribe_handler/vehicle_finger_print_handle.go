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

	//if detect != nil {
	tradeMark := mac.GetOrgByMAC(fingerprintParam.GetMac())
	fTradeMark := util.RrgsTrim(tradeMark)
	fprintDetectInfo := &model.FprintDetectInfo{
		DetectInfoId: util.RandomString(32),
		DeviceMac:    fingerprintParam.GetMac(),
		VehicleId:    vehicleId,
		Os:           detect.GetOs(),
		TradeMark:    fTradeMark,
	}
	detectInfoModelBase := model_base.ModelBaseImpl(fprintDetectInfo)

	_, dRecordNotFound := detectInfoModelBase.GetModelByCondition("device_mac = ? and os = ?",
		[]interface{}{fprintDetectInfo.DeviceMac, fprintDetectInfo.Os}...)

	detectInfoModelBase.CreateModel(fingerprintParam)
	if dRecordNotFound {
		if err := detectInfoModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert fprintDetectInfo err:%s", util.RunFuncName(), err.Error())
		}
	}
	//}

	//if passive != nil {
	passiveTradeMark := mac.GetOrgByMAC(fingerprintParam.GetMac())
	passiveFTradeMark := util.RrgsTrim(passiveTradeMark)
	fprintPassiveInfo := &model.FprintPassiveInfo{
		PassiveInfoId: util.RandomString(32),
		DeviceMac:     fingerprintParam.GetMac(),
		VehicleId:     vehicleId,
		DstPort:       passive.GetDstPort(),
		TradeMark:     passiveFTradeMark,
	}
	passiveInfoModelBase := model_base.ModelBaseImpl(fprintPassiveInfo)

	_, recordNotFound := passiveInfoModelBase.GetModelByCondition("device_mac = ? and dst_port = ?",
		[]interface{}{fprintPassiveInfo.DeviceMac, fprintPassiveInfo.DstPort}...)

	passiveInfoModelBase.CreateModel(fingerprintParam)
	if recordNotFound {
		if err := passiveInfoModelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert fprintDetectInfo err:%s", util.RunFuncName(), err.Error())
		}
	}
	//}
	return nil
}
