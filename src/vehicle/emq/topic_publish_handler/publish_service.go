package topic_publish_handler

import (
	"time"
	"vehicle_system/src/vehicle/conf"
)
var quitChanValue = 0
var publishService *PublishService
type PublishService struct {
	PublishTopicChan chan interface{}
}
func GetPublishService() *PublishService {
	if publishService == nil{
		publishService = &PublishService{
			PublishTopicChan: make(chan interface{},conf.PublishChanCapa),
		}
	}
	return publishService
}
func (p *PublishService)PutMsg2PublicChan(data interface{})  {
	quitChan:=make(chan int)
	go StartPutMsg2PublicChan(p,data,quitChan)
	quitChan <- quitChanValue
	go SelectPublishTopicChan(publishService.PublishTopicChan)
}
func StartPutMsg2PublicChan(p *PublishService,data interface{},quitChan chan int){
	p.PublishTopicChan <- data
	select {
		case <-quitChan:
			return
	}
}
/**
发信息
 */
func SelectPublishTopicChan(publishTopicChan chan interface{})  {
	timeCh := time.NewTicker(2 * time.Second)
	defer timeCh.Stop()
	for {
		select {
			case data := <-publishTopicChan:
				//common_util.Vfmtf(log_util.LOG_WEB, "EmqClient SelectPublishTopicChan <-publishTopicChan %+v:\n",data)
				//log_util.VlogInfo(log_util.LOG_WEB, "EmqClient SelectPublishTopicChan <-publishTopicChan %+v:\n",data)
				PublishTopicMsg(data)
				return
			case <-timeCh.C:
				//common_util.Vfmtf(log_util.LOG_WEB, "EmqClient SelectPublishTopicChan Free Chan\n")
				//log_util.VlogInfo(log_util.LOG_WEB, "EmqClient SelectPublishTopicChan Free Chan\n")
				return
		}
	}
}