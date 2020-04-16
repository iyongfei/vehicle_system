package subscribe_server_test

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

func SubscribeServerTest(client mqtt.Client, msg mqtt.Message)  {
	serverTopic:=msg.Topic()

	logger.Logger.Info("%s serverTopic:%s",util.RunFuncName(),serverTopic)
	logger.Logger.Print("%s serverTopic:%s",util.RunFuncName(),serverTopic)

	command := &protobuf.Command{}
	_ = proto.Unmarshal(msg.Payload(), command)

	switch command.ItemType {
	case protobuf.Command_GW_SET:
		gwSetModel := &protobuf.GwSetParam{}
		param := command.GetParam()
		_ = proto.Unmarshal(param,gwSetModel)

		logger.Logger.Info("%s gwSetModel:%+v",util.RunFuncName(),gwSetModel)
		logger.Logger.Print("%s gwSetModel:%+v",util.RunFuncName(),gwSetModel)

	case protobuf.Command_DEVICE_SET:
		assetSetModel := &protobuf.DeviceSetParam{}
		param := command.GetParam()
		_ = proto.Unmarshal(param,assetSetModel)

		logger.Logger.Info("%s assetSetModel:%+v",util.RunFuncName(),assetSetModel)
		logger.Logger.Print("%s assetSetModel:%+v",util.RunFuncName(),assetSetModel)

	}
}