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

func HandleVehicleDeployer(vehicleResult protobuf.GWResult) error {
	//parse
	vehicleDeployerParam := &protobuf.DeployerParam{}
	err:=proto.Unmarshal(vehicleResult.GetParam(),vehicleDeployerParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal vehicle deployer param err:%s",util.RunFuncName(),err.Error())
		logger.Logger.Error("%s unmarshal vehicle deployer param err:%s",util.RunFuncName(),err.Error())
		return fmt.Errorf("%s unmarshal vehicle deployer err:%s",util.RunFuncName(),err.Error())
	}
	//vehicleId
	vehicleId:=vehicleResult.GetGUID()

	logger.Logger.Print("%s unmarshal vehicle deployer param:%+v",util.RunFuncName(),vehicleDeployerParam)
	logger.Logger.Info("%s unmarshal vehicle deployer param:%+v",util.RunFuncName(),vehicleDeployerParam)
	//create
	leaderInfo:=&model.VehicleLeader{
		LeaderId:util.RandomString(32),
		Name:vehicleDeployerParam.GetName(),
		Phone:vehicleDeployerParam.GetPhone(),
	}
	modelBase := model_base.ModelBaseImpl(leaderInfo)

	_,recordNotFound :=modelBase.GetModelByCondition("name = ? and phone = ?",
		[]interface{}{leaderInfo.Name, leaderInfo.Phone}...)

	modelBase.CreateModel(vehicleDeployerParam)


	if recordNotFound{
		if err:=modelBase.InsertModel();err!=nil{
			return fmt.Errorf("%s insert vehicle deployer err:%s",util.RunFuncName(),err.Error())
		}

	}else {
		//更新
		attrs := map[string]interface{}{
			"phone": leaderInfo.Phone,
			"name": leaderInfo.Name,
		}
		if err:=modelBase.UpdateModelsByCondition(attrs,"name = ? and phone = ?",
			[]interface{}{leaderInfo.Name,leaderInfo.Phone}...);err!=nil{
			return fmt.Errorf("%s update vehicle protect err:%s",util.RunFuncName(),err.Error())
		}
	}

	//为小V赋值
	vehicleInfo := &model.VehicleInfo{
		VehicleId:vehicleId,
	}
	vehicleModelBase := model_base.ModelBaseImpl(vehicleInfo)
	vehicleLeaderAttrs := map[string]interface{}{
		"leader_id": leaderInfo.LeaderId,
		"name": vehicleDeployerParam.GetName(),
	}

	err= vehicleModelBase.UpdateModelsByCondition(vehicleLeaderAttrs,"vehicle_id = ?",vehicleInfo.VehicleId)
	if err!=nil{
		return fmt.Errorf("%s update vehicle leaderid:%s err:%s",util.RunFuncName(),vehicleId,err.Error())
	}
	return nil
}