package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleAsset(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	assetParam := &protobuf.DeviceParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), assetParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal assetParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal assetParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal assetParam err:%s", util.RunFuncName(), err.Error())
	}

	logger.Logger.Print("%s unmarshal assetParam:%+v", util.RunFuncName(), assetParam)
	logger.Logger.Info("%s unmarshal assetParam:%+v", util.RunFuncName(), assetParam)
	//create
	vehicleInfo := &model.VehicleInfo{
		VehicleId: vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	_, recordNotFound := modelBase.GetModelByCondition("vehicle_id = ?", vehicleInfo.VehicleId)

	if recordNotFound {
		return fmt.Errorf("%s insert asset vehicleId recordNotFound", util.RunFuncName())
	}

	//初始化资产默认分组
	assetGroup := &model.AreaGroup{
		AreaName:       response.UnGroupName,
		AreaCode:       util.RandomString(32),
		ParentAreaCode: "",
		TreeAreaCode:   "",
	}

	assetGroupModelBase := model_base.ModelBaseImpl(assetGroup)

	_, assetGroupRecordNotFound := assetGroupModelBase.GetModelByCondition("area_name = ?",
		[]interface{}{assetGroup.AreaName}...)
	if assetGroupRecordNotFound {
		err := assetGroupModelBase.InsertModel()
		if err != nil {
			return fmt.Errorf("%s insert asset ungroup err:%s", err)
		}
	}

	for _, assetItem := range assetParam.GetDeviceItem() {
		asset := &model.Asset{
			VehicleId:  vehicleId,
			AssetId:    assetItem.GetMac(),
			AssetGroup: assetGroup.AreaCode,
		}

		modelBase := model_base.ModelBaseImpl(asset)

		_, recordNotFound := modelBase.GetModelByCondition("asset_id = ?", asset.AssetId)
		modelBase.CreateModel(assetItem)

		if recordNotFound {
			//检索白名单列表
			exist := checkoutAssetPrintInfos(asset.AssetId)
			asset.ProtectStatus = exist

			err := modelBase.InsertModel()
			if err != nil {
				continue
			}
		} else {
			//更新
			attrs := map[string]interface{}{
				"vehicle_id":       asset.VehicleId,
				"asset_id":         asset.AssetId,
				"ip":               asset.IP,
				"mac":              asset.Mac,
				"name":             asset.Name,
				"trade_mark":       asset.TradeMark,
				"online_status":    asset.OnlineStatus,
				"last_online":      asset.LastOnline,
				"internet_switch":  asset.InternetSwitch,
				"protect_status":   asset.ProtectStatus,
				"lan_visit_switch": asset.LanVisitSwitch,
				"asset_group":      asset.AssetGroup,
				"asset_leader":     asset.AssetLeader,
			}
			if err := modelBase.UpdateModelsByCondition(attrs, "asset_id = ?", asset.AssetId); err != nil {
				continue
			}
		}
	}

	return nil
}

func checkoutAssetPrintInfos(assetId string) bool {
	whiteAsset := &model.WhiteAsset{
		DeviceMac: assetId,
	}
	modelBase := model_base.ModelBaseImpl(whiteAsset)

	_, recordNotFound := modelBase.GetModelByCondition("device_mac = ?", whiteAsset.DeviceMac)

	if recordNotFound {
		return false
	}
	return true
}
