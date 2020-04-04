package topic_router

import "github.com/eclipse/paho.mqtt.golang"

type TopicRouterData struct {
	Client mqtt.Client
	Msg mqtt.Message
}

func GetTopicRouterData(client mqtt.Client, msg mqtt.Message) *TopicRouterData {
	return &TopicRouterData{
		Client : client,
		Msg : msg,
	}
}

func (emqTopicData *TopicRouterData)HandleTopicRouterData()  {
	HandleTopicRouterDataGo(emqTopicData)
}



