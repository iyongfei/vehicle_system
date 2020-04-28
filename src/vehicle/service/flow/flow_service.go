package flow

import (
	"encoding/json"
	"reflect"
	"time"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

//chanLength
var flowChanDefaultLength = 100
var writChanDuration = 2 * time.Second
var VFlowService *FlowService

func Setup() {
	flowService := CreatFlowService()
	flowImpl := FlowImpl(flowService)

	flowImpl.ReadFlow()
}

//get
func GetFlowService() (service *FlowService) {
	if VFlowService == nil {
		return CreatFlowService()
	}
	return VFlowService
}

type FlowService struct {
	FlowChan     chan map[string]interface{}
	FlowData     map[string]interface{}
	WriteTimeout time.Duration
	ReadIdleFlag bool
}

//creat
func CreatFlowService() (service *FlowService) {
	if VFlowService == nil {
		VFlowService = &FlowService{
			FlowChan:     make(chan map[string]interface{}, GetFlowChanLength(flowChanDefaultLength)),
			WriteTimeout: GetWriteChanDuration(2 * time.Second),
		}
	}
	return VFlowService
}

//setData
func (fservce *FlowService) SetFlowData(FlowData map[string]interface{}) *FlowService {
	fservce.FlowData = FlowData
	return fservce
}

//readFlowGo
func (fservce *FlowService) ReadFlow() {
	startReadFlowG(fservce)
}

//writeFlowGo
func (fservce *FlowService) WriteFlow() {
	startWriteFlowG(fservce)

}
func (f *FlowService) SendFlow(data interface{}) {
	logger.Logger.Print("%s send flow info %+v", util.RunFuncName(), data)

	switch data.(type) {
	case map[string]interface{}:

		postData := data.(map[string]interface{})

		actionType := postData[ActionType]
		url := getFlowReqUrl(actionType.(string))

		resp, _ := util.PostJson(url, postData, "")
		respJsonBys, _ := json.Marshal(resp)

		logger.Logger.Print("%s sendFlow resp json:%s, type:%+v", util.RunFuncName(), string(respJsonBys), reflect.TypeOf(data))
		logger.Logger.Info("%s sendFlow resp json:%s, type:%+v", util.RunFuncName(), string(respJsonBys), reflect.TypeOf(data))
	default:
		logger.Logger.Print("%s send flow info err type:%+v", util.RunFuncName(), reflect.TypeOf(data))
		logger.Logger.Info("%s send flow info err type:%+v", util.RunFuncName(), reflect.TypeOf(data))
	}
}
