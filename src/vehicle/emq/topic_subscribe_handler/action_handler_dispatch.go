package topic_subscribe_handler

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"strings"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
)

var topicSubscribeHandler *TopicSubscribeHandler

type TopicSubscribeHandler struct {
}

func GetTopicSubscribeHandler() *TopicSubscribeHandler {
	if topicSubscribeHandler == nil {
		topicSubscribeHandler = new(TopicSubscribeHandler)
	}
	return topicSubscribeHandler
}


/**
(web)EmqClient LineStatus LineO Topic:$SYS/brokers/emqx@127.0.0.1/clients/tianqi-R201b-967E6D9A3001/disconnected,
Payload:{"clientid":"tianqi-R201b-967E6D9A3001","username":"undefined","reason":"keepalive_timeout","ts":1561184234}

(web)EmqClient LineStatus LineO Topic:$SYS/brokers/emqx@127.0.0.1/clients/tianqi-R201b-967E6D9A3001/connected,
Payload:{"clean_start":true,"clientid":"tianqi-R201b-967E6D9A3001","connack":0,"ipaddress":"192.168.18.2","keepalive":2,"proto_name":"MQTT","proto_ver":4,"ts":1561184268,"username":"undefined"}

$SYS/brokers/emqx@127.0.0.1/clients/vehicle_test/connected
 */
func (t *TopicSubscribeHandler) HanleSubscribeTopicData(topicMsg mqtt.Message) error {
	disconnected:=strings.HasSuffix(topicMsg.Topic(),"disconnected")
	fmt.Println("topicMsg::",topicMsg.Topic())
	//_=strings.HasSuffix(topicMsg.Topic(),"connected")

	if disconnected{
		err :=HanleSubscribeTopicLineData(topicMsg)
		return err
	}

	vehicleResult := protobuf.GWResult{}
	err := proto.Unmarshal(topicMsg.Payload(), &vehicleResult)

	if err != nil {
		return fmt.Errorf("hanleSubscribeTopicData unmarshal payload err:%s",err)
	}
	vehicleId := vehicleResult.GetGUID()

	//更新为在线
	//gwOnlineAttrs:=[]interface{}{"online_status",true}
	//_=mysql_util.UpdateModelOneColumn(&gw.GwInfo{},gwOnlineAttrs,"gw_id = ?",[]interface{}{gwId}...)
	//_=mysql_util.UpdateModelOneColumn(&gw.GwInfoUncerted{},gwOnlineAttrs,"gw_id = ?",[]interface{}{gwId}...)


	actionTypeName:=protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]

	logger.Logger.Print("hanleSubscribeTopicData action name:%s,vehicleId:%s",actionTypeName,vehicleId)
	logger.Logger.Info("hanleSubscribeTopicData action name:%s,vehicleId:%s",actionTypeName,vehicleId)

	var handGwResultError error
	switch actionType := vehicleResult.ActionType; actionType {
	case protobuf.GWResult_THREAT: //威胁
		handGwResultError = HandleVehicleThreat(vehicleResult)
	case protobuf.GWResult_GW_INFO: //小v
		handGwResultError = HandleVehicleInfo(vehicleResult)
	case protobuf.GWResult_STRATEGY: //策略
		handGwResultError = HandleVehicleStrategy(vehicleResult)
	default:
		logger.Logger.Error("vehicleId:%s action type err:%d",vehicleId,int32(vehicleResult.ActionType))
		logger.Logger.Print("vehicleId:%s action type err:%d",vehicleId,int32(vehicleResult.ActionType))
	}
	return  handGwResultError
}

func HanleSubscribeTopicLineData(topicMsg mqtt.Message) error {
	//topicSlice:=strings.Split(topicMsg.Topic(),"$SYS/brokers/emqx@127.0.0.1/clients/")
	subscribeLineTopic := "$SYS/brokers/emqx@127.0.0.1/clients/"
	topicSlice:=strings.Split(topicMsg.Topic(),subscribeLineTopic)
	topicSlice_1 :=topicSlice[1]
	vehicleId := strings.Split(topicSlice_1,"/")[0]
	err := HandleVehicleOnline(vehicleId,false)
	fmt.Println(err,err!=nil,"sjdfl")
	return err
}
