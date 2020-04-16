package topic_subscribe_handler

import (
	"fmt"
	"time"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleCommonAction(vehicleResult protobuf.GWResult) error {
	vehicleId :=vehicleResult.GetGUID()

	//分组
	areaGroup := &model.AreaGroup{
		AreaName:response.UnGroupName,
		AreaCode:util.RandomString(32),
	}
	areaGroupModelBase := model_base.ModelBaseImpl(areaGroup)
	err,areaGroupUnExist := areaGroupModelBase.GetModelByCondition("area_name = ?",areaGroup.AreaName)
	if areaGroupUnExist{
		if err := areaGroupModelBase.InsertModel();err!=nil{
			return fmt.Errorf("%s vehicleId %s insert group err:%+v",util.RunFuncName(),vehicleId,areaGroup)
		}
	}


	vehicleInfo:=&model.VehicleInfo{
		VehicleId:vehicleId,
		//StartTime:model_base.UnixTime(time.Now()),
		StartTime:time.Now(),
		GroupId:areaGroup.AreaCode,
	}
	modelBase := model_base.ModelBaseImpl(vehicleInfo)

	err,recordNotFound :=modelBase.GetModelByCondition("vehicle_id = ?",vehicleInfo.VehicleId)

	if err!=nil{
		return fmt.Errorf("%s vehicleId:%s not exist",util.RunFuncName(),vehicleId)
	}
	if recordNotFound{
		err:= modelBase.InsertModel()
		if err!=nil{
			return fmt.Errorf("%s insert vehicleId:%s,err:%s",util.RunFuncName(),vehicleId,err.Error())
		}
	}else {
		attrs := map[string]interface{}{
			"group_id": vehicleInfo.GroupId,
		}
		if err:=modelBase.UpdateModelsByCondition(attrs,"vehicle_id = ?",
			[]interface{}{vehicleInfo.GroupId}...);err!=nil{
			return fmt.Errorf("%s update vehicle err:%s",util.RunFuncName(),err.Error())
		}
	}
	return nil
}





