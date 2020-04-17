package emq_cmd

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
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

	return nil, nil
}
