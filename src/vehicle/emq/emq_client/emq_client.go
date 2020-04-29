package emq_client

import (
	"github.com/eclipse/paho.mqtt.golang"
	"sync"
	"time"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/db/tdata"
	"vehicle_system/src/vehicle/emq/subscribe_server_test"
	"vehicle_system/src/vehicle/emq/topic_router"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

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
		logger.Logger.Print("%s,err:%s", util.RunFuncName(), token.Error())
		logger.Logger.Error("%s,err:%s", util.RunFuncName(), token.Error())
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

func EmqReConnectTokenError() {

	//var err error
	//if util.RrgsTrimEmpty(vehicleId) {
	//	err = HandleVehicleOfflineStatus(vehicleId, false)
	//}

	err := tdata.VehicleAssetCheck("", false)
	if err != nil {
		logger.Logger.Print("%s,emqReConnectTokenError update vehicle online status err:%+v", util.RunFuncName(), err)
		logger.Logger.Info("%s,emqReConnectTokenError update vehicle online status err:%+v", util.RunFuncName(), err)
	}

	t := time.NewTicker(time.Second * 10)
	select {
	case <-t.C:
		if !EmqClient.IsConnected() {
			logger.Logger.Print("%s,emqClient:%v", util.RunFuncName(), &EmqClient)
			logger.Logger.Info("%s,emqClient:%v", util.RunFuncName(), &EmqClient)
			GetEmqInstance().InitEmqClient()
		}
		t.Stop()
		return
	}
}

func (m *EmqInstance) GetEmqClient() (emqClient mqtt.Client) {
	if EmqClient == nil {
		m.InitEmqClient()
	}
	return EmqClient
}
