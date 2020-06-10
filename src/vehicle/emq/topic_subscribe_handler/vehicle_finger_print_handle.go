package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/db/mysql"
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

	//查询终端
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	_, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", vehicleInfo.VehicleId)

	if recordNotFound {
		return fmt.Errorf("%s insert asset vehicleId recordNotFound", util.RunFuncName())
	}

	//////////////////////////事务开始//////////////////////
	vgorm, err := mysql.GetMysqlInstance().GetMysqlDB()
	tx := vgorm.Begin()

	//插入asset
	asset := &model.Asset{
		VehicleId:  vehicleId,
		AssetId:    fingerprintParam.GetMac(),
		AssetGroup: vehicleInfo.GroupId,
	}
	isAssetInfoExist := tx.Where("asset_id = ?",
		[]interface{}{asset.AssetId}...).First(asset).RecordNotFound()

	if isAssetInfoExist {
		exist := checkoutAssetPrintInfos(asset.AssetId)
		asset.AccessNet = exist

		if err = tx.Create(asset).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("%s insert fprintInfo err:%s", util.RunFuncName(), err.Error())
		}
	}

	detect := fingerprintParam.GetActive()
	passive := fingerprintParam.GetPassive()

	fprintInfo := &model.FprintInfo{
		FprintInfoId: util.RandomString(32),
		DeviceMac:    fingerprintParam.GetMac(),
		VehicleId:    vehicleId,

		TradeMark: util.RrgsTrim(mac.GetOrgByMAC(fingerprintParam.GetMac())),
	}

	isFprintInfoExist := tx.Where("device_mac = ? and vehicle_id = ?",
		[]interface{}{fprintInfo.DeviceMac, fprintInfo.VehicleId}...).First(fprintInfo).RecordNotFound()

	if isFprintInfoExist {
		if err = tx.Create(fprintInfo).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("%s insert fprintInfo err:%s", util.RunFuncName(), err.Error())
		}
	}

	//添加被动
	fprintInfoPassive := &model.FprintInfoPassive{
		FprintInfoId: fprintInfo.FprintInfoId,
		DstPort:      passive.GetDstPort(),
	}

	isFprintInfoPassiveExist := tx.Where("fprint_info_id = ?",
		[]interface{}{fprintInfoPassive.FprintInfoId}...).First(fprintInfoPassive).RecordNotFound()

	fprintInfoPassive.CreateModel(passive)

	if isFprintInfoPassiveExist {
		if err = tx.Create(fprintInfoPassive).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("%s insert fprintInfoPassive err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		attrs := map[string]interface{}{
			"dst_port": fprintInfoPassive.DstPort,
		}

		if err = tx.Model(&model.FprintInfoPassive{}).Where("fprint_info_id = ?",
			[]interface{}{fprintInfoPassive.FprintInfoId}...).Updates(attrs).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("%s update fprintInfoPassive err:%s", util.RunFuncName(), err.Error())
		}
	}

	//被动信息采集列表
	passiveInfos := passive.GetItem()
	for _, passiveInfo := range passiveInfos {

		fprintPassiveInfo := &model.FprintPassiveInfo{
			FprintInfoId: fprintInfo.FprintInfoId,
			Hash:         passiveInfo.GetHash(),
		}

		isfprintPassiveInfoExist := tx.Where("fprint_info_id = ? and hash = ?",
			[]interface{}{fprintPassiveInfo.FprintInfoId, fprintPassiveInfo.Hash}...).First(fprintPassiveInfo).RecordNotFound()

		fprintPassiveInfo.CreateModel(passiveInfo)

		if isfprintPassiveInfoExist {
			if err = tx.Create(fprintPassiveInfo).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("%s insert fprintPassiveInfo err:%s", util.RunFuncName(), err.Error())
			}
		} else {
			attrs := map[string]interface{}{
				"hash":           fprintPassiveInfo.Hash,
				"src_ip":         fprintPassiveInfo.SrcIp,
				"src_port":       fprintPassiveInfo.SrcPort,
				"dst_ip":         fprintPassiveInfo.DstIp,
				"dst_port":       fprintPassiveInfo.DstPort,
				"protocol":       fprintPassiveInfo.Protocol,
				"flow_info":      fprintPassiveInfo.FlowInfo,
				"safe_type":      fprintPassiveInfo.SafeType,
				"safe_info":      fprintPassiveInfo.SafeInfo,
				"start_time":     fprintPassiveInfo.StartTime,
				"last_seen_time": fprintPassiveInfo.LastSeenTime,
				"src_dst_bytes":  fprintPassiveInfo.SrcDstBytes,
				"dst_src_bytes":  fprintPassiveInfo.DstSrcBytes,
				"stat":           fprintPassiveInfo.Stat,
			}

			if err = tx.Model(&model.FprintPassiveInfo{}).Where("fprint_info_id = ? and hash = ?",
				[]interface{}{fprintPassiveInfo.FprintInfoId, fprintPassiveInfo.Hash}...).Updates(attrs).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("%s update fprintPassiveInfo err:%s", util.RunFuncName(), err.Error())
			}
		}
	}

	//添加主动
	fprintInfoActive := &model.FprintInfoActive{
		FprintInfoId: fprintInfo.FprintInfoId,
		Os:           detect.GetOs(),
	}

	isFprintInfoActiveExist := tx.Where("fprint_info_id = ?",
		[]interface{}{fprintInfoActive.FprintInfoId}...).First(fprintInfoActive).RecordNotFound()

	fprintInfoActive.CreateModel(detect)

	if isFprintInfoActiveExist {
		if err = tx.Create(fprintInfoActive).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("%s insert fprintInfoActive err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		attrs := map[string]interface{}{
			"os": fprintInfoActive.Os,
		}

		if err = tx.Model(&model.FprintInfoPassive{}).Where("fprint_info_id = ?",
			[]interface{}{fprintInfoPassive.FprintInfoId}...).Updates(attrs).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("%s update fprintInfoActive err:%s", util.RunFuncName(), err.Error())
		}

	}

	tx.Commit()

	return nil
}
