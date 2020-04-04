package topic_subscribe_handler

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"
	"strings"
	"vehicle_system/src/vehicle/emq/emq_client"
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
 */
func (t *TopicSubscribeHandler) HanleSubscribeTopicData(topicMsg mqtt.Message) error {
	disconnected:=strings.HasSuffix(topicMsg.Topic(),"disconnected")
	_=strings.HasSuffix(topicMsg.Topic(),"connected")

	if disconnected{
		//topicSlice:=strings.Split(topicMsg.Topic(),"$SYS/brokers/emqx@127.0.0.1/clients/")
		topicSlice:=strings.Split(topicMsg.Topic(),emq_client.SUBSCRIBE_LINE_TOPIC)
		topicSlice_1 :=topicSlice[1]
		gwId := strings.Split(topicSlice_1,"/")[0]
		err := HandleGwManageInfoOnline(gwId,false)
		//common_util.Vfmtf(log_util.LOG_WEB,"EmqClient LineStatus disconnected Gwid:%s\n",gwId)
		//log_util.VlogInfo(log_util.LOG_WEB,"EmqClient LineStatus disconnected Gwid:%s\n",gwId)
		return err
	}

	pushReq := vsubscribe.GWResult{}
	err := proto.Unmarshal(topicMsg.Payload(), &pushReq)

	if err != nil {
		log_util.VlogInfo(log_util.LOG_WEB, "HanleSubscribeTopicData proto unmarshal requestBody to GWResult error:[%v]\n", err)
		common_util.Vfmtf(log_util.LOG_WEB, "HanleSubscribeTopicData proto unmarshal requestBody to GWResult error:[%v]\n", err)
		return err
	}
	gwId := pushReq.GetGUID()

	//更新为在线
	gwOnlineAttrs:=[]interface{}{"online_status",true}
	_=mysql_util.UpdateModelOneColumn(&gw.GwInfo{},gwOnlineAttrs,"gw_id = ?",[]interface{}{gwId}...)
	_=mysql_util.UpdateModelOneColumn(&gw.GwInfoUncerted{},gwOnlineAttrs,"gw_id = ?",[]interface{}{gwId}...)

	var GWResult_ActionType_name = map[int32]string{
		0: "DEFAULT",
		1: "DEVICE",
		2: "THREAT",
		3: "GW_INFO",
		4: "SAMPLE",
		5: "PROTECT",
		6: "STRATEGY",
		7: "PORTREDIRECT",
		8: "DEPLOYER",
		9: "FIRMWARE",
	}

	//log_util.VlogInfo(log_util.LOG_WEB, "HanleSubscribeTopicData ActionType:[%s],CmdId:[%s],Gwid:[%s]\n", GWResult_ActionType_name[int32(pushReq.ActionType)],pushReq.CmdID,pushReq.GUID)
	//common_util.Vfmtf(log_util.LOG_WEB, "HanleSubscribeTopicData ActionType:[%s],CmdId:[%s],Gwid:[%s]\n", GWResult_ActionType_name[int32(pushReq.ActionType)],pushReq.CmdID,pushReq.GUID)

	var handGwResultError error
	switch actionType := pushReq.ActionType; actionType {
	case vsubscribe.GWResult_DEVICE: //设备
		handGwResultError = HandleAssetsManageInfoList(pushReq,gwId)
	case vsubscribe.GWResult_THREAT: //威胁
		handGwResultError = HandleThreatInfoList(pushReq,gwId)
	case vsubscribe.GWResult_GW_INFO: //小v
		handGwResultError = HandleGwManageInfoList(pushReq,gwId)
	case vsubscribe.GWResult_DEPLOYER: //小v负责人
		handGwResultError = HandleGwManageLeaderInfo(pushReq,gwId)
	case vsubscribe.GWResult_SAMPLE: //样本
		handGwResultError = HandleSampleManageInfo(pushReq,gwId)
	case vsubscribe.GWResult_PROTECT: //全局保护
		handGwResultError = HandleGwProtectInfo(pushReq,gwId)
	case vsubscribe.GWResult_STRATEGY: //策略
		handGwResultError = HandleGwStrategy(pushReq,gwId)
	case vsubscribe.GWResult_PORTREDIRECT: //映射
		handGwResultError = PushGwPortRedirect(pushReq,gwId)
	case vsubscribe.GWResult_FIRMWARE: //更新
		handGwResultError = PushFirware(pushReq,gwId)
	default:
		log_util.VlogInfo(log_util.LOG_WEB, "handleGwPushSvr invalidParam actionType is not exist,gwId:%s",gwId)
		common_util.Vfmtf(log_util.LOG_WEB, "handleGwPushSvr invalidParam actionType is not exist,gwId:%s",gwId)
	}
	return  handGwResultError
}
