package emq_client

import (
	"sync"
	"github.com/eclipse/paho.mqtt.golang"
	"time"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/emq/topic_router"
)

const (
	SUBSCRIBESEEVERTOPIC = "s/+/p"
	SUBSCRIBE_MAIN_TOPIC = "+/s/p"
	SUBSCRIBE_LINE_TOPIC = "$SYS/brokers/emqx@127.0.0.1/clients/+/+"
)



var EmqClient mqtt.Client

//单例
type EmqInstance struct {}
var emqInstance *EmqInstance

var EmqConnectOnce sync.Once
func GetEmqInstance() *EmqInstance {
	EmqConnectOnce.Do(NewClient)
	return 	emqInstance
}

func NewClient()  {
	if emqInstance == nil{
		emqInstance = new(EmqInstance)
	}
}

func (m *EmqInstance) InitEmqClient()  {
	if EmqClient != nil {
		EmqClient.Disconnect(250)
	}
	clientOptions := m.NewClientOptions()
	EmqClient = mqtt.NewClient(clientOptions)

	if token := EmqClient.Connect(); token.Wait() && token.Error() != nil {
		//common_util.Vfmtf(log_util.LOG_WEB,"EmqClient Connect Error:[%s],EmqClient:[%v]\n",token.Error(),&EmqClient)
		//log_util.VlogInfo(log_util.LOG_WEB,"EmqClient Connect Error:[%s],EmqClient:[%v]\n",token.Error(),&EmqClient)
		//go EmqTokenError()
		//return
	}
	if token := EmqClient.Subscribe(SUBSCRIBE_MAIN_TOPIC, 0, topic_router.MainTopicRouter); token.Wait() && token.Error() != nil {
		//common_util.Vfmtf(log_util.LOG_WEB,"EmqClient Subscribe +/s/p Error:%s\n",token.Error())
		//log_util.VlogInfo(log_util.LOG_WEB,"EmqClient Subscribe +/s/p Error:%s\n",token.Error())
	}
	//在离线
	if token := EmqClient.Subscribe(SUBSCRIBE_LINE_TOPIC, 0, topic_router.LineTopicRouter); token.Wait() && token.Error() != nil {
		//common_util.Vfmtf(log_util.LOG_WEB,"EmqClient Subscribe OnOffLine Error:%s\n",token.Error())
		//log_util.VlogInfo(log_util.LOG_WEB,"EmqClient Subscribe OnOffLine Error:%s\n",token.Error())
	}
	//SetGWOnlineInfoWithEmqOffline(false)
	//common_util.Vfmtf(log_util.LOG_WEB, "EmqClient Init Success:[%v]\n",&EmqClient)
	//log_util.VlogInfo(log_util.LOG_WEB, "EmqClient Init Success:[%v]\n",&EmqClient)
}

func (m *EmqInstance) NewClientOptions()  *mqtt.ClientOptions{
	return  mqtt.NewClientOptions().
		AddBroker(conf.EmqBrokerUrl).
		SetClientID(conf.EmqClientId).
		SetKeepAlive(conf.EmqKeepAlive * time.Second).
		SetAutoReconnect(conf.AutoReconnect).
		SetConnectTimeout(conf.ConnectTimeOut *time.Second).
		SetPingTimeout(conf.EmqPingTimeOut * time.Second).
		SetMaxReconnectInterval(20 * time.Second).
		SetConnectionLostHandler(Conconnlost).
		SetMaxReconnectInterval(conf.MaxReconnectInterval * time.Second).
		SetCleanSession(conf.EmqCleanSession).
		SetUsername(conf.EmqUser).
		SetPassword(conf.EmqPassword).SetTLSConfig(NewTLSConfig())
}


//
//func (m *EmqInstance)GetEmqClient() (emqClient mqtt.Client) {
//	if EmqClient == nil{
//		m.InitEmqClient()
//	}
//	return EmqClient
//}

