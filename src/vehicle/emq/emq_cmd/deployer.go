package emq_cmd

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

type DeployerSetCmd struct {
	VehicleId string
	CmdId     int
	TaskType  int

	Name    string
	Phone   string
	DevName string
}

func (setCmd *DeployerSetCmd) CreateDeployerTopicMsg() interface{} {
	publishItem := &protobuf.Command{}

	//ItemType
	publishItem.ItemType = protobuf.Command_TaskType(setCmd.TaskType)
	//param
	deployerSetParams := &protobuf.DeployerSetParam{}
	deployerSetParams.Name = setCmd.Name
	deployerSetParams.Phone = setCmd.Phone
	deployerSetParams.DevName = setCmd.DevName
	publishItem.Param, _ = proto.Marshal(deployerSetParams)
	//CmdID
	var resultcmdItemKey string
	taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
	cmdRandom := util.RandomString(16)
	resultcmdItemKey = createCmdId(taskTypeName, cmdRandom, util.RandomString(16))
	publishItem.CmdID = resultcmdItemKey

	resultcmdItemsBys, _ := proto.Marshal(publishItem)

	logger.Logger.Info("%s createDeployerTopicMsg publishItem:%+v", util.RunFuncName(), publishItem)
	logger.Logger.Print("%s createDeployerTopicMsg publishItem:%+v", util.RunFuncName(), publishItem)

	return resultcmdItemsBys
}
