package emq_cmd

import (
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

type PortMapSetCmd struct {
	VehicleId   string
	CmdId int
	TaskType int

	DestPort string
	SrcPort  string
	Switch   bool
	Protocol int
	DestIp   string
}

func (setCmd *PortMapSetCmd) CreatePortMapSetCmdTopicMsg() interface{}{
	publishItem := &protobuf.Command{}

	//ItemType
	publishItem.ItemType = protobuf.Command_TaskType(setCmd.TaskType)
	//param
	portMapSetParams := &protobuf.PortRedirectSetParam{}
	portMapSetParams.DestIp = setCmd.DestIp
	portMapSetParams.DestPort = setCmd.DestPort
	portMapSetParams.SrcPort = setCmd.SrcPort
	portMapSetParams.Switch = setCmd.Switch
	portMapSetParams.Proto = protobuf.PortRedirectSetParam_Protocol(setCmd.Protocol)
	publishItem.Param, _ = proto.Marshal(portMapSetParams)
	//CmdID
	var resultcmdItemKey string
	taskTypeName := protobuf.Command_TaskType_name[int32(setCmd.TaskType)]
	cmdRandom := util.RandomString(16)
	resultcmdItemKey = createCmdId(taskTypeName, cmdRandom)
	publishItem.CmdID = resultcmdItemKey
	resultcmdItemsBys, _ := proto.Marshal(publishItem)

	logger.Logger.Info("%s createPortMapSetCmdTopicMsg publishItem:%+v",util.RunFuncName(),publishItem)
	logger.Logger.Print("%s createPortMapSetCmdTopicMsg publishItem:%+v",util.RunFuncName(),publishItem)

	return resultcmdItemsBys
}

