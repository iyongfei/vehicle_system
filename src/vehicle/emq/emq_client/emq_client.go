package emq_client

import (
	"github.com/eclipse/paho.mqtt.golang"
	"sync"
	"time"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/emq/subscribe_server_test"
	"vehicle_system/src/vehicle/emq/topic_router"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model/model_helper"
	"vehicle_system/src/vehicle/util"
)

/**
mosquitto_pub -h localhost -p 1883  -t topic -m 'sljdfklsjd' //211
mosquitto_sub -h localhost -p 1883 -t topic//211

mosquitto_sub -h 211.159.167.112 -p 1883 -t topic //本机
*/

/**
mosquitto_sub -h 192.168.1.103 -p 8843 --cafile ca.crt -t "a/b" --insecure

mosquitto_pub -h 192.168.1.103 -p 8843 -t "a/b" -m " hello" --cafile ca.crt

*/
const (
	SUBSCRIBE_SEEVER_TOPIC = "s/+/p"
	SUBSCRIBE_MAIN_TOPIC   = "+/s/p"
	SUBSCRIBE_LINE_TOPIC   = "$SYS/brokers/emqx@127.0.0.1/clients/+/+"
)

var EmqClient mqtt.Client

//单例
type EmqInstance struct{}

var emqInstance *EmqInstance

var EmqConnectOnce sync.Once

func GetEmqInstance() *EmqInstance {
	EmqConnectOnce.Do(NewClient)
	return emqInstance
}

func NewClient() {
	if emqInstance == nil {
		emqInstance = new(EmqInstance)
	}
}

func (m *EmqInstance) InitEmqClient() {
	if EmqClient != nil {
		EmqClient.Disconnect(250)
	}
	clientOptions := m.NewClientOptions()
	EmqClient = mqtt.NewClient(clientOptions)

	if token := EmqClient.Connect(); token.Wait() && token.Error() != nil {
		logger.Logger.Print("%s,mosquitto connect err:%s", util.RunFuncName(), token.Error())
		logger.Logger.Error("%s,mosquitto connect err:%s", util.RunFuncName(), token.Error())
		go EmqReConnectTokenError()
		return
	}
	if token := EmqClient.Subscribe(SUBSCRIBE_MAIN_TOPIC, 0, topic_router.MainTopicRouter); token.Wait() && token.Error() != nil {
		logger.Logger.Print("%s,err:%s", util.RunFuncName(), token.Error())
		logger.Logger.Error("%s,err:%s", util.RunFuncName(), token.Error())
	}
	//在离线
	if token := EmqClient.Subscribe(SUBSCRIBE_LINE_TOPIC, 0, topic_router.LineTopicRouter); token.Wait() && token.Error() != nil {
		logger.Logger.Print("%s,err:%s", util.RunFuncName(), token.Error())
		logger.Logger.Error("%s,err:%s", util.RunFuncName(), token.Error())
	}

	//command测试
	if token := EmqClient.Subscribe(SUBSCRIBE_SEEVER_TOPIC, 0, subscribe_server_test.SubscribeServerTest); token.Wait() && token.Error() != nil {
		logger.Logger.Print("%s,err:%s", util.RunFuncName(), token.Error())
		logger.Logger.Error("%s,err:%s", util.RunFuncName(), token.Error())
	}

	logger.Logger.Print("%s,emqClient init success:%v", util.RunFuncName(), &EmqClient)
	logger.Logger.Info("%s,emqClient init success:%v", util.RunFuncName(), &EmqClient)

	//a := model_helper.JudgeAssetCollectByteTotalRate("DfQWLAOw")
	//b := model_helper.JudgeAssetCollectTlsInfoRate("DfQWLAOw")
	//c := model_helper.JudgeAssetCollectHostNameRate("DfQWLAOw")
	//d := model_helper.JudgeAssetCollectProtoFlowRate("DfQWLAOw")
	//e := model_helper.JudgeAssetCollectTimeRate("DfQWLAOw")
	//
	//fmt.Println(a, b, c, d, e)
	//
	//fmt.Println("rate....", a+b+c+d+e)

	model_helper.GetAssetCateStdMark()

}

func (m *EmqInstance) NewClientOptions() *mqtt.ClientOptions {
	return mqtt.NewClientOptions().
		AddBroker(conf.EmqBrokerUrl).
		SetClientID(conf.EmqClientId).
		SetKeepAlive(conf.EmqKeepAlive * time.Second).
		SetAutoReconnect(conf.AutoReconnect).
		SetConnectTimeout(conf.ConnectTimeOut * time.Second).
		SetPingTimeout(conf.EmqPingTimeOut * time.Second).
		SetMaxReconnectInterval(20 * time.Second).
		SetConnectionLostHandler(Conconnlost).
		SetMaxReconnectInterval(conf.MaxReconnectInterval * time.Second).
		SetCleanSession(conf.EmqCleanSession).
		SetUsername(conf.EmqUser).
		SetPassword(conf.EmqPassword).SetTLSConfig(NewTLSConfig())
}

func (m *EmqInstance) GetEmqClient() (emqClient mqtt.Client) {
	if EmqClient == nil {
		m.InitEmqClient()
	}
	return EmqClient
}
