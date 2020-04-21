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

func HandleVehicleAsset(vehicleResult protobuf.GWResult) error {
	//parse
	assetParam := &protobuf.DeviceParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), assetParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal assetParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal assetParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal assetParam err:%s", util.RunFuncName(), err.Error())
	}
	//vehicleId
	vehicleId := vehicleResult.GetGUID()

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
	for _, assetItem := range assetParam.GetDeviceItem() {
		asset := &model.Asset{
			VehicleId: vehicleId,
			AssetId:   assetItem.GetMac(),
		}

		modelBase := model_base.ModelBaseImpl(asset)
		modelBase.CreateModel(assetItem)
		_, recordNotFound := modelBase.GetModelByCondition("asset_id = ?", asset.AssetId)
		if recordNotFound {
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
				//return fmt.Errorf("%s update flow err:%s",util.RunFuncName(),err.Error())
				continue
			}
		}
	}

	return nil
}
