package emq_cmd

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"strings"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

type StrategySetCmd struct {
	VehicleId string
	CmdId     int
	TaskType  int

	StrategyId string
	Type       int
	HandleMode int
	Enable     bool
	GroupId    string
}

func (setCmd *StrategySetCmd) CreateStrategyTopicMsg() interface{} {
	publishItem := &protobuf.Command{}

	//ItemType
	publishItem.ItemType = protobuf.Command_TaskType(setCmd.TaskType)
	//param
	strategySetParams := &protobuf.StrategyAddParam{}
	strategySetParams.StrategyId = setCmd.StrategyId
	strategySetParams.HandleMode = protobuf.StrategyAddParam_HandleMode(setCmd.HandleMode)
	strategySetParams.DefenseType = protobuf.StrategyAddParam_Type(setCmd.Type)
	strategySetParams.Enable = setCmd.Enable

	dipList, urlList := FetchDipUrlList(setCmd)
	strategySetParams.DIPList = dipList
	strategySetParams.URLList = urlList
	publishItem.Param, _ = proto.Marshal(strategySetParams)
	//CmdID
	var resultcmdItemKey string
	taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
	cmdRandom := util.RandomString(16)
	resultcmdItemKey = createCmdId(taskTypeName, cmdRandom)

	publishItem.CmdID = resultcmdItemKey
	resultcmdItemsBys, _ := proto.Marshal(publishItem)

	logger.Logger.Info("%s createAssetTopicMsg publishItem:%+v", util.RunFuncName(), publishItem)
	logger.Logger.Print("%s createAssetTopicMsg publishItem:%+v", util.RunFuncName(), publishItem)

	return resultcmdItemsBys
}

/*
获取策略的url,ip列表，策略合并
 */
func FetchDipUrlList(setCmd *StrategySetCmd) ([]string, []string) {
	dipList := []string{}
	urlList := []string{}

	//获取learning_result_ids
	strategyVehicleLearningResultJoin, _ := model.GetStrategyVehicleLearningResults("strategy_vehicles.vehicle_id = ?", []interface{}{setCmd.VehicleId}...)

	automatedLearningResultIds := []string{}
	for _, v := range strategyVehicleLearningResultJoin {
		automatedLearningResultIds = append(automatedLearningResultIds, v.LearningResultId)
	}

	//去重learning_result_ids
	fAutomatedLearningResultIds := util.RemoveRepeatedForSlice(automatedLearningResultIds)

	//获取automatedLearningResultModelList
	automatedLearningResultModelList := []*model.AutomatedLearningResult{}

	_ = model_base.ModelBaseImpl(&model.AutomatedLearningResult{}).
		GetModelListByCondition(&automatedLearningResultModelList,
			"learning_result_id in (?)", []interface{}{fAutomatedLearningResultIds}...)

	dipMap := map[string]interface{}{}
	urlMap := map[string]interface{}{}

	for _, learningResult := range automatedLearningResultModelList {
		switch learningResult.OriginType {
		case 1: //sample
			sampleId := learningResult.OriginId
			collectSampleItems := []*model.SampleItem{}
			_ = model_base.ModelBaseImpl(&model.SampleItem{}).GetModelListByCondition(&collectSampleItems, "sample_id = ?", []interface{}{sampleId}...)
			for _, sampleItems := range collectSampleItems {
				dip := sampleItems.DstIp
				url := sampleItems.Url
				fmt.Println("dip",dip,",url:",url)

				if strings.Trim(dip, " ") != "" {
					dipMap[dip] = dip
				}
				if strings.Trim(url, " ") != "" {
					urlMap[url] = url
				}
			}

		case 2:

		case 3:

		}
	}

	for _, dip := range dipMap {
		dipList = append(dipList, dip.(string))
	}

	for _, url := range urlMap {
		urlList = append(urlList, url.(string))
	}
	logger.Logger.Info("%s urlList:%+v,dipList:%+v", util.RunFuncName(), urlList,dipList)
	logger.Logger.Print("%s urlList:%+v,dipList:%+v", util.RunFuncName(), urlList,dipList)

	return dipList, urlList
}
