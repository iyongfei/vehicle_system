package topic_subscribe_handler

import (
	"errors"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"strings"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
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
	disconnected := strings.HasSuffix(topicMsg.Topic(), "disconnected")
	//_=strings.HasSuffix(topicMsg.Topic(),"connected")

	if disconnected {
		err := HanleSubscribeTopicLineData(topicMsg)
		return err
	}

	//parse
	vehicleResult := protobuf.GWResult{}
	err := proto.Unmarshal(topicMsg.Payload(), &vehicleResult)

	if err != nil {
		//$SYS/brokers/emqx@127.0.0.1/clients/vehicle_test/connected
		return fmt.Errorf("hanleSubscribeTopicData unmarshal payload err:%s", err)
	}
	//vehicleId null
	vehicleIdTrim := util.RrgsTrimEmptyTableEnter(vehicleResult.GetGUID())
	vehicleId := vehicleResult.GetGUID()

	logger.Logger.Print("hanleSubscribeTopicData vehicleIdTrim_length:%d,vehicleId_length:%d", len(vehicleIdTrim), len(vehicleId))
	logger.Logger.Info("hanleSubscribeTopicData vehicleIdTrim_length:%d,vehicleId_length:%d", len(vehicleIdTrim), len(vehicleId))

	fmt.Printf("vehicleResult vehicleId:%s", vehicleId)
	if util.RrgsTrimEmpty(vehicleId) {
		return fmt.Errorf("vehicleResult vehicle id null")
	}

	//vehicleId exist
	actionCommonErr := HandleVehicleCommonAction(vehicleResult, vehicleId)

	if actionCommonErr != nil {
		return actionCommonErr
	}

	////////////////////////////////////////////////////////////////////////////////////////
	//更新为在线
	//gwOnlineAttrs:=[]interface{}{"online_status",true}
	//_=mysql_util.UpdateModelOneColumn(&gw.GwInfo{},gwOnlineAttrs,"gw_id = ?",[]interface{}{gwId}...)
	//_=mysql_util.UpdateModelOneColumn(&gw.GwInfoUncerted{},gwOnlineAttrs,"gw_id = ?",[]interface{}{gwId}...)

	actionTypeName := protobuf.GWResult_ActionType_name[int32(vehicleResult.ActionType)]

	logger.Logger.Print("hanleSubscribeTopicData action name:%s,vehicleId:%s", actionTypeName, vehicleId)
	logger.Logger.Info("hanleSubscribeTopicData action name:%s,vehicleId:%s", actionTypeName, vehicleId)

	var handGwResultError error
	switch actionType := vehicleResult.ActionType; actionType {

	case protobuf.GWResult_GW_INFO: //GwInfoParam
		handGwResultError = HandleVehicleInfo(vehicleResult, vehicleId)

	case protobuf.GWResult_MONITORINFO: //MONITORINFO
		handGwResultError = HandleMonitorInfo(vehicleResult, vehicleId)

	case protobuf.GWResult_FLOWSTATISTIC: //FlowStatisticInfo
		handGwResultError = HandleFlowStatisticInfo(vehicleResult, vehicleId)

	case protobuf.GWResult_DEPLOYER: //DeployerParam
		handGwResultError = HandleVehicleDeployer(vehicleResult, vehicleId)

	case protobuf.GWResult_PROTECT: //GWProtectInfoParam
		handGwResultError = HandleVehicleProtect(vehicleResult, vehicleId)

		//////////////////////////////////////////////
	case protobuf.GWResult_FLOWSTAT: //FlowParam
		handGwResultError = HandleVehicleFlow(vehicleResult, vehicleId)

	case protobuf.GWResult_FIRMWARE: //FirwareParam
		handGwResultError = HandleVehicleFirmware(vehicleResult, vehicleId)

	case protobuf.GWResult_DEVICE: //DeviceParam
		handGwResultError = HandleVehicleAsset(vehicleResult, vehicleId)

	case protobuf.GWResult_THREAT: //ThreatParam
		handGwResultError = HandleVehicleThreat(vehicleResult, vehicleId)

	case protobuf.GWResult_SAMPLE: //SampleParam
		handGwResultError = HandleVehicleSample(vehicleResult, vehicleId)

	case protobuf.GWResult_STRATEGY: //StrawtegyParam
		handGwResultError = HandleVehicleStrategy(vehicleResult, vehicleId)

	case protobuf.GWResult_FLOWSTRATEGYSTAT: //flowStrawtegyParam
		handGwResultError = HandleVehicleFlowStrategy(vehicleResult, vehicleId)

	case protobuf.GWResult_PORTREDIRECT: //PortRedirectParam
		handGwResultError = HandleVehiclePortMap(vehicleResult, vehicleId)

	default:
		logger.Logger.Error("vehicleId:%s action type err:%d", vehicleId, int32(vehicleResult.ActionType))
		logger.Logger.Print("vehicleId:%s action type err:%d", vehicleId, int32(vehicleResult.ActionType))
		handGwResultError = errors.New("vehicleResult action type err")
		return handGwResultError
	}
	return handGwResultError
}

//$SYS/brokers/emqx@127.0.0.1/clients/+/+
func HanleSubscribeTopicLineData(topicMsg mqtt.Message) error {
	//$SYS/brokers/emqx@127.0.0.1/clients/tianqi-R201b-967E6D9A3001/disconnected

	subscribeLineTopic := "$SYS/brokers/emqx@127.0.0.1/clients/"

	var vehicleId string
	if strings.Contains(topicMsg.Topic(), subscribeLineTopic) {
		topicSlice := strings.Split(topicMsg.Topic(), subscribeLineTopic)
		topicSliceSuffix := topicSlice[1]

		if strings.HasSuffix(topicSliceSuffix, "disconnected") {
			vehicleId = strings.Split(topicSliceSuffix, "/")[0]
		}
	}

	var err error
	if util.RrgsTrimEmpty(vehicleId) {
		err = HandleVehicleOnline(vehicleId, false)
	}

	return err
}
