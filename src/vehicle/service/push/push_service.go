package push

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
var VFlowService *PushService

func Setup() {
	flowService := GetPushervice()
	flowImpl := PushImpl(flowService)

	flowImpl.Read()
}

//get
func GetPushervice() (service *PushService) {
	if VFlowService == nil {
		return CreatPushService()
	}
	return VFlowService
}

type PushService struct {
	PushChan     chan map[string]interface{}
	FlowData     map[string]interface{}
	WriteTimeout time.Duration
	ReadIdleFlag bool
}

//creat
func CreatPushService() (service *PushService) {
	if VFlowService == nil {
		VFlowService = &PushService{
			PushChan:     make(chan map[string]interface{}, GetPushChanLength(flowChanDefaultLength)),
			WriteTimeout: GetWriteChanDuration(2 * time.Second),
		}
	}
	return VFlowService
}

//setData
func (fservce *PushService) SetPushData(FlowData map[string]interface{}) *PushService {
	fservce.FlowData = FlowData
	return fservce
}

//readFlowGo
func (fservce *PushService) Read() {
	startReadPushG(fservce)
}

//writeFlowGo
func (fservce *PushService) Write() {
	startWritePushG(fservce)

}
func (f *PushService) Send(data interface{}) {
	logger.Logger.Print("%s send flow info %+v", util.RunFuncName(), data)

	switch data.(type) {
	case map[string]interface{}:

		postData := data.(map[string]interface{})

		actionType := postData[ActionType]
		url := getPushReqUrl(actionType.(string))

		resp, _ := util.PostJson(url, postData, "")
		respJsonBys, _ := json.Marshal(resp)

		logger.Logger.Print("%s sendFlow resp json:%s, type:%+v", util.RunFuncName(), string(respJsonBys), reflect.TypeOf(data))
		logger.Logger.Info("%s sendFlow resp json:%s, type:%+v", util.RunFuncName(), string(respJsonBys), reflect.TypeOf(data))
	default:
		logger.Logger.Print("%s send flow info err type:%+v", util.RunFuncName(), reflect.TypeOf(data))
		logger.Logger.Info("%s send flow info err type:%+v", util.RunFuncName(), reflect.TypeOf(data))
	}
}
