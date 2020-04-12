package flow

import (
	"context"
	"encoding/json"
	"reflect"
	"time"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)
//chanLength
var flowChanDefaultLength = 100
var writChanDuration = 2* time.Second
var VFlowService *FlowService
func Setup()  {
	flowService := CreatFlowService()
	flowImpl:=FlowImpl(flowService)
	flowImpl.ReadFlow()
}
//get
func GetFlowService()  (service *FlowService) {
	if VFlowService == nil{
		return CreatFlowService()
	}
	return VFlowService
}
type FlowService struct {
	FlowChan chan map[string]interface{}
	FlowData map[string]interface{}
	Timeout time.Duration
}
//creat
func CreatFlowService()  (service *FlowService) {
	if VFlowService == nil{
		VFlowService = &FlowService{
			FlowChan: make(chan map[string]interface{},GetFlowChanLength(flowChanDefaultLength)),
			Timeout:GetWriteChanDuration(2*time.Second),
		}
	}
	return VFlowService
}
//chanLength
func GetFlowChanLength(chanLength int)int{
	if chanLength <= 0{
		return flowChanDefaultLength
	}
	return chanLength
}
//timeout
func GetWriteChanDuration(timer time.Duration)time.Duration{
	if timer <= 0{
		return timer*time.Second
	}
	return writChanDuration
}
//setData
func  (fservce *FlowService) SetFlowData(FlowData map[string]interface{})  *FlowService{
	fservce.FlowData = FlowData
	return VFlowService
}
//readFlowGo
func (fservce *FlowService) ReadFlow()  {
	go ReadFlowGo(fservce)

}
func ReadFlowGo(fService *FlowService)  {
	for {
		select {
		case flowData := <-fService.FlowChan:

			//发送请求
			fService.SendFlow(flowData)
		}
	}
}
//writeFlowGo
func (fservce *FlowService)  WriteFlow()  {
	go WriteFlowGo(fservce)
}
func WriteFlowGo(fservice *FlowService)  {
	fservice.FlowChan <- fservice.FlowData
	ctx, cancel := context.WithTimeout(context.Background(), fservice.Timeout)
	defer cancel()
	select {
	case <-ctx.Done():
		return
	}
}
func (f *FlowService)  SendFlow(data interface{})  {
	logger.Logger.Print("%s send flow info %+v",util.RunFuncName(),data)
	url:= getFlowReq()
	switch data.(type) {
	case map[string]interface{}:
		postData := data.(map[string]interface{})

		resp,_:= util.Post(url,postData)
		respJsonBys,_:=json.Marshal(resp)

		logger.Logger.Print("%s sendFlow resp json err%+v",util.RunFuncName(),string(respJsonBys),reflect.TypeOf(data))
		logger.Logger.Info("%s sendFlow resp json err%+v",util.RunFuncName(),string(respJsonBys),reflect.TypeOf(data))
	default:
		logger.Logger.Print("%s send flow info type err%+v",util.RunFuncName(),reflect.TypeOf(data))
		logger.Logger.Info("%s send flow info type err%+v",util.RunFuncName(),reflect.TypeOf(data))
	}
}

