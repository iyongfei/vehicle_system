package emq_cmd

import (
	"bytes"
	"strings"
)

const (
	TaskTypeDefault = "default"
	//网关设置
	TaskTypeGwSet_Protect = "gwset_protect"
	TaskTypeGwSet_Reset   = "gwset_reset"
	//资产设置
	TaskTypeDeviceSet_Protect  = "deviceset_protect"
	TaskTypeDeviceSet_Internet = "deviceset_internet"
	TaskTypeDeviceSet_Access   = "deviceset_access"
	TaskTypeDeviceSet_Lanvisit = "deviceset_lanvisit"
	//策略添加
	TaskTypeStrategyAdd = "strategyadd"
	//策略修改
	TaskTypeStrategySet = "strategyset"
	//样本采集
	TaskTypeSampleSet = "sampleset"
	//端口映射
	TaskTypePortSet = "portset"
	//状态更新
	TaskTypeStatusUpgrade_GwInfo    = "statusupdate_gwinfo"
	TaskTypeStatusUpgrade_Device    = "statusupdate_device"
	TaskTypeStatusUpgrade_GwProtect = "statusupdate_gwprotect"
	TaskTypeGwSetGwInfo_Strategy    = "statusupdate_strategy"
	TaskTypeGwSetGwInfo_Portmap     = "statusupdate_portmap"
	TaskTypeGwSetGwInfo_Deployer    = "statusupdate_deployer"
	//更新个人信息
	TaskTypeDeployerSet = "deployerset"
	//更新版本
	TaskTypeFirmwareUpgrade = "firmwareupgrade"
	//会话策略添加
	TaskTypeFlowStrategyAdd = "flowstrategyadd"
	//会话策略设置
	TaskTypeFlowStrategySet = "flowstrategyset"
)


func createCmdId(args ...string) string {
	var buffer bytes.Buffer
	for _,arg:=range args{
		if strings.Trim(arg, " ") != ""{
			buffer.WriteString(arg)
		}
	}
	return buffer.String()
}

func ParseCmdPayload()  {



}