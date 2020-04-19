package emq_cmd

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

type FStrategySetCmd struct {
	VehicleId string
	CmdId     int
	TaskType  int

	FstrategyId string
	Type       int
	HandleMode int
	Enable     bool
	GroupId    string
}

func (setCmd *FStrategySetCmd) CreateFStrategyTopicMsg() interface{} {
	publishItem := &protobuf.Command{}

	//ItemType
	publishItem.ItemType = protobuf.Command_TaskType(setCmd.TaskType)
	//param
	fstrategySetParams := &protobuf.FlowStrategyAddParam{}
	fstrategySetParams.FlowStrategyId = setCmd.FstrategyId
	fstrategySetParams.HandleMode = protobuf.FlowStrategyAddParam_HandleMode(setCmd.HandleMode)
	fstrategySetParams.DefenseType = protobuf.FlowStrategyAddParam_Type(setCmd.Type)
	fstrategySetParams.Enable = setCmd.Enable

	dipPortList:= FetchDipPortList(setCmd)
	fstrategySetParams.DIPList = dipPortList

	publishItem.Param, _ = proto.Marshal(fstrategySetParams)
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

func FetchDipPortList(setCmd *FStrategySetCmd) []*protobuf.FlowStrategyAddParam_FlowStrategyItem {
		//获取learning_result_ids
	strategyVehicleLearningResultJoin, _ := model.GetFlowStrategyVehicleItems(
		"fstrategy_vehicles.vehicle_id = ?", []interface{}{setCmd.VehicleId}...)

	fstrategyVehicleItemIds := []string{}
	for _, v := range strategyVehicleLearningResultJoin {
		fstrategyVehicleItemIds = append(fstrategyVehicleItemIds, v.FstrategyItemId)
	}

	//去重learning_result_ids
	fFstrategyVehicleItems := util.RemoveRepeatedForSlice(fstrategyVehicleItemIds)

	fstrategyVehicleItems := []*model.FstrategyItem{}

	_ = model_base.ModelBaseImpl(&model.AutomatedLearningResult{}).
		GetModelListByCondition(&fstrategyVehicleItems,
			"fstrategy_item_id in (?)", []interface{}{fFstrategyVehicleItems}...)


	fProtobufStrategyVehicleItems := []*protobuf.FlowStrategyAddParam_FlowStrategyItem{}
	mapper := map[string][]uint32{}
	for _, fItem := range fstrategyVehicleItems {
		//去重
		dip:=fItem.DstIp
		dport:=fItem.DstPort

		if len(mapper[dip]) == 0{

			mapper[dip] = []uint32{dport}
		}else {
			if !util.IsExistInSlice(dport,mapper[dip]){
				mapper[dip] = append(mapper[dip],dport)
			}
		}
	}

	for dip,ports:=range mapper{
		for _,port:=range ports{
			fProtobufStrategyVehicleItem := &protobuf.FlowStrategyAddParam_FlowStrategyItem{}
			fProtobufStrategyVehicleItem.DstIp = uint32(util.InetAton(dip))
			fProtobufStrategyVehicleItem.DstPort = port

			fProtobufStrategyVehicleItems = append(fProtobufStrategyVehicleItems,fProtobufStrategyVehicleItem)
		}
	}

	return fProtobufStrategyVehicleItems
}

