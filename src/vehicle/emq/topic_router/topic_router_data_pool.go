package topic_router

import (
	"context"
	"github.com/eclipse/paho.mqtt.golang"
	"time"
	"vehicle_system/src/vehicle/conf"
	"vehicle_system/src/vehicle/emq/topic_subscribe_handler"
)

var topicRouterDataCh = make(chan mqtt.Message,conf.SubscribeChanCapa)

func HandleTopicRouterDataGo(emqTopicData *TopicRouterData) {
	go PutTopicRouterDataCh(emqTopicData)
	go FetchtopicRouterDataCh()
}

func PutTopicRouterDataCh(emqTopicData *TopicRouterData) {
	topicRouterDataCh <- emqTopicData.Msg
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}

func FetchtopicRouterDataCh() {
	timeCh := time.NewTicker(2*time.Second)
	defer timeCh.Stop()
	for {
		select {
		case topicMsg := <-topicRouterDataCh:
			err := topic_subscribe_handler.GetTopicSubscribeHandler().HanleSubscribeTopicData(topicMsg)
			if err!=nil{
				//log_util.VlogInfo(log_util.LOG_WEB, "EmqClient SelectSubscribeTopicData Err:%s\n",err)
				//common_util.Vfmtf(log_util.LOG_WEB, "EmqClient SelectSubscribeTopicData Err:%s\n",err)
				return
			}
			return
		case <-timeCh.C:
			return
		}
	}
}
