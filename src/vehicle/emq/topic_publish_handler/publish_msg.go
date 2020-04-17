package topic_publish_handler

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"vehicle_system/src/vehicle/emq/emq_client"
	"vehicle_system/src/vehicle/emq/emq_cmd"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

func GetEmqClient()  mqtt.Client {
	return emq_client.GetEmqInstance().GetEmqClient()
}
//Pub 盒子发送的GUID/s/p   s/GUID/p(服务器发送)
//Sub 盒子接受的s/GUID/p   +/s/p(服务器订阅盒子)

func PublishTopicMsg(data interface{}){
	emqClient := GetEmqClient()

	var payload interface{}
	var vehicleId string

	switch data.(type) {
	case *emq_cmd.VehicleSetCmd:
		vehicleSetCmd := data.(*emq_cmd.VehicleSetCmd)
		payload = vehicleSetCmd.CreateVehicleTopicMsg()
		vehicleId = vehicleSetCmd.VehicleId

	case *emq_cmd.AssetSetCmd:
		assetSetCmd := data.(*emq_cmd.AssetSetCmd)
		payload = assetSetCmd.CreateAssetTopicMsg()
		vehicleId = assetSetCmd.VehicleId

	case *emq_cmd.StrategySetCmd:
		strategySetCmd := data.(*emq_cmd.StrategySetCmd)
		payload = strategySetCmd.CreateStrategyTopicMsg()
		vehicleId = strategySetCmd.VehicleId

	case *emq_cmd.DeployerSetCmd:
		deployerSetCmd := data.(*emq_cmd.DeployerSetCmd)
		payload = deployerSetCmd.CreateDeployerTopicMsg()
		vehicleId = deployerSetCmd.VehicleId
	default:
	}

	if token := emqClient.Publish(fmt.Sprintf("s/%s/p",vehicleId), 0, false, payload);
		token.Wait() && token.Error() != nil {
		logger.Logger.Error("%s publishTopicMsg err:%s",util.RunFuncName(),token.Error())
		logger.Logger.Print("%s publishTopicMsg err:%s",util.RunFuncName(),token.Error())
	}
}