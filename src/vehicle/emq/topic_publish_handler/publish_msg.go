package topic_publish_handler

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"vehicle_system/src/vehicle/emq/emq_client"
	"vehicle_system/src/vehicle/emq/emq_cmd"
)

func GetEmqClient()  mqtt.Client {
	return emq_client.GetEmqInstance().GetEmqClient()
}
//Pub 盒子发送的GUID/s/p   s/GUID/p(服务器发送)
//Sub 盒子接受的s/GUID/p   +/s/p(服务器订阅盒子)

func PublishTopicMsg(data interface{}){
	emqClient := GetEmqClient()

	var bys interface{}
	var vehicleId string

	switch data.(type) {
	case *emq_cmd.VehicleSetCmd:
		vehicleSetCmd := data.(*emq_cmd.VehicleSetCmd)
 		bys = vehicleSetCmd.CreateProtoTopicMsg()
		vehicleId = vehicleSetCmd.VehicleId
	default:
	}

	if token := emqClient.Publish(fmt.Sprintf("s/%s/p",vehicleId), 0, false, bys);
		token.Wait() && token.Error() != nil {
		//common_util.Vfmtf(log_util.LOG_WEB,"EmqClient PublishTopicMsg error:%s\n",token.Error())
		//log_util.VlogInfo(log_util.LOG_WEB,"EmqClient PublishTopicMsg error:%s\n",token.Error())
	}
}

