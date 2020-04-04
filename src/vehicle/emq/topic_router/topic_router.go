package topic_router

import "github.com/eclipse/paho.mqtt.golang"

func MainTopicRouter(client mqtt.Client, msg mqtt.Message)  {

	GetTopicRouterData(client,msg).HandleTopicRouterData()

}

func LineTopicRouter(client mqtt.Client, msg mqtt.Message)  {

	GetTopicRouterData(client,msg).HandleTopicRouterData()

}