package response

const (
	VStatusBadRequestMsg = "服务器请求失败,请尝试重新操作"
	ReqArgsIllegalMsg    = "服务器请求参数错误"
)

const (
	TrueFlag  = "true"
	FalseFlag = "false"
)

const (
	OriginTypeSelf      = 1
	OriginTypeSample    = 2
	OriginTypeRule      = 3
	OriginTypeWhiteList = 4
)

const (
	TokenExpiredStr                    = "token过期"
	TokenNotValidYetStr                = "token未激活"
	TokenMalformedStr                  = "token不合法"
	TokenInvalidStr                    = "token未知"
	ValidationErrorUnverifiableStr     = "签名信息错误，无法验证token"
	ValidationErrorSignatureInvalidStr = "签名验证失败"
)

//token
const (
	AuthTokenLost     = "请求未携带token，权限不足"
	AuthTokenResignin = "请重新登录"
)

//解析
const (
	UnmarshalErr = "数据解析失败"
)

//group
const (
	UnGroupName           = "未分组"
	ReqAddGroupExistMsg   = "分组已存在"
	ReqAddGroupSuccessMsg = "分组添加成功"
	ReqAddGroupFailMsg    = "分组添加失败"
)

//regist
const (
	PasswordSecret      = "vgw-1214-pwd-key"
	ReqRegistFailMsg    = "用户注册失败"
	ReqRegistSuccessMsg = "用户注册成功"
	ReqRegistExistMsg   = "该用户已注册"
	ReqRegistUnAuthMsg  = "该用户未注册"

	ReqRegistAuthFailMsg    = "该用户授权失败"
	ReqRegistAuthSuccessMsg = "该用户授权成功"
)

//white_list
const (
	ReqAddWhiteListFailMsg       = "白名单添加失败"
	ReqAddWhiteListSuccessMsg    = "白名单添加成功"
	ReqGetWhiteListSuccessMsg    = "白名单获取成功"
	ReqGetWhiteListFailMsg       = "白名单获取失败"
	ReqGetWhiteListUnExistMsg    = "白名单不存在"
	ReqDeleWhiteListFailMsg      = "白名单删除失败"
	ReqDeleWhiteListSuccessMsg   = "白名单删除成功"
	ReqUpdateWhiteListSuccessMsg = "白名单更新成功"
	ReqUpdateWhiteListFailMsg    = "白名单更新失败"
	//ReqRegistSuccessMsg = "用户注册成功"
	//ReqRegistExistMsg = "该用户已注册"
	//ReqRegistUnAuthMsg = "该用户未注册"
	//ReqRegistAuthFailMsg = "该用户授权失败"
	//ReqRegistAuthSuccessMsg = "该用户授权成功"
)

//
const (
	ReqGetFlowFailMsg     = "会话消息获取失败"
	ReqGetFlowSuccessMsg  = "会话消息获取成功"
	ReqGetFlowUnExistMsg  = "会话消息不存在"
	ReqGetFlowExistMsg    = "会话消息已存在"
	ReqAddFlowFailMsg     = "会话消息添加失败"
	ReqAddFlowSuccessMsg  = "会话消息添加成功"
	ReqEditFlowSuccessMsg = "会话消息更新成功"
	ReqEditFlowFailMsg    = "会话消息更新失败"
	ReqDeleFlowFailMsg    = "会话消息删除失败"
)

//monitor
const (
	ReqGetMonitorsFailMsg    = "监控信息获取失败"
	ReqGetMonitorsSuccessMsg = "监控信息获取成功"

	//ReqGetAssetUnExistMsg = "设备信息不存在"
	//ReqGetAssetExistMsg   = "设备信息已存在"
	//
	//ReqDeleAssetFailMsg    = "设备信息删除失败"
	//ReqDeleAssetSuccessMsg = "设备信息删除成功"
	//
	//ReqUpdateAssetFailMsg    = "设备信息更新失败"
	//ReqUpdateAssetSuccessMsg = "设备信息更新成功"
	//
	//ReqAddAssetFailMsg    = "设备信息添加失败"
	//ReqAddAssetSuccessMsg = "设备信息添加成功"
	//
	//ReqGetAssetListFailMsg    = "设备列表获取失败"
	//ReqGetAssetListSuccessMsg = "设备列表获取成功"
)

//threat
const (
	ReqGetThreatInfoFailMsg    = "威胁信息获取失败"
	ReqGetThreatInfoSuccessMsg = "威胁信息获取成功"
	ReqGetThreatInfoUnExistMsg = "威胁信息不存在"
	ReqDeleThreatInfoFailMsg   = "威胁信息删除失败"
)

//vehicle
const (
	ReqGetVehicleFailMsg    = "终端信息获取失败"
	ReqGetVehicleSuccessMsg = "终端信息获取成功"

	ReqGetVehicleUnExistMsg = "终端信息不存在"
	ReqGetVehicleExistMsg   = "终端信息已存在"

	ReqDeleVehicleFailMsg    = "终端信息删除失败"
	ReqDeleVehicleSuccessMsg = "终端信息删除成功"

	ReqUpdateVehicleFailMsg    = "终端信息更新失败"
	ReqUpdateVehicleSuccessMsg = "终端信息更新成功"

	ReqAddVehicleFailMsg    = "终端信息添加失败"
	ReqAddVehicleSuccessMsg = "终端信息添加成功"

	ReqGetVehiclesFailMsg    = "终端列表获取失败"
	ReqGetVehiclesSuccessMsg = "终端列表获取成功"
)

//depolyer
const (
	ReqGetVehicleBindLeaderUnExistMsg    = "终端没有绑定管理人员信息"
	ReqGetVehicleBindLeaderFailMsg       = "终端绑定管理人员信息获取失败"
	ReqUpdateVehicleBindLeaderSuccessMsg = "终端绑定管理人员信息更新成功"
)

//portMap
const (
	ReqGetPortMapUnExistMsg    = "端口映射信息不存在"
	ReqGetPortMapFailMsg       = "端口映射信息获取失败"
	ReqGetPortMapSuccessMsg    = "端口映射信息获取成功"
	ReqUpdatePortMapSuccessMsg = "端口映射信息更新成功"
	ReqUpdatePortMapFailMsg    = "端口映射信息更新失败"
)

//csv
const (
	ReqGetFstrategyCsvUnExistMsg = "会话策略csv不存在"
	ReqFstrategyCsvFailMsg       = "会话策略csv获取失败"
	ReqFstrategyCsvSuccessMsg    = "会话策略csv获取成功"
)

//asset
const (
	ReqGetAssetFailMsg    = "设备信息获取失败"
	ReqGetAssetSuccessMsg = "设备信息获取成功"

	ReqGetAssetUnExistMsg = "设备信息不存在"
	ReqGetAssetExistMsg   = "设备信息已存在"

	ReqDeleAssetFailMsg    = "设备信息删除失败"
	ReqDeleAssetSuccessMsg = "设备信息删除成功"

	ReqUpdateAssetFailMsg    = "设备信息更新失败"
	ReqUpdateAssetSuccessMsg = "设备信息更新成功"

	ReqAddAssetFailMsg    = "设备信息添加失败"
	ReqAddAssetSuccessMsg = "设备信息添加成功"

	ReqGetAssetListFailMsg    = "设备列表获取失败"
	ReqGetAssetListSuccessMsg = "设备列表获取成功"
)

//strategy
const (
	ReqAddStrategyFailMsg    = "策略添加失败"
	ReqAddStrategySuccessMsg = "策略添加成功"
	ReqGetStrategyExistMsg   = "策略已存在"
	ReqGetStrtegyUnExistMsg  = "策略信息不存在"

	ReqGetStrategyListFailMsg    = "策略列表获取失败"
	ReqGetStrategyListSuccessMsg = "策略列表获取成功"

	ReqGetStrategyFailMsg    = "策略信息获取失败"
	ReqGetStrategySuccessMsg = "策略信息获取成功"

	ReqDeleStrategyFailMsg    = "策略信息删除失败"
	ReqDeleStrategySuccessMsg = "策略信息删除成功"

	ReqUpdateStrategyFailMsg    = "策略信息更新失败"
	ReqUpdateStrategySuccessMsg = "策略信息更新成功"

	ReqGetStrategyVehicleFailMsg    = "策略-终端信息获取失败"
	ReqGetStrategyVehicleUnExistMsg = "策略-终端信息不存在"
	ReqGetStrategyVehicleSuccessMsg = "策略-终端信息获取成功"

	ReqGetStrategyVehicleListFailMsg    = "策略-终端列表获取失败"
	ReqGetStrategyVehicleListSuccessMsg = "策略-终端列表获取成功"

	ReqGetStrategyVehicleResultListFailMsg    = "终端-学习结果列表获取失败"
	ReqGetStrategyVehicleResultListSuccessMsg = "终端-学习结果列表获取成功"
)

//flow_statistic
const (
	ReqGetVehicleStatisticListFailMsg    = "网络状态列表获取失败"
	ReqGetVehicleStatisticListSuccessMsg = "网络状态列表获取成功"
)

//fstrategy
const (
	//ReqAddStrategyFailMsg    = "策略添加失败"
	ReqAddFStrategySuccessMsg = "会话策略添加成功"
	//ReqGetStrategyExistMsg    = "策略已存在"
	ReqGetFStrtegyUnExistMsg = "会话策略信息不存在"
	//
	ReqGetFStrategyListFailMsg    = "会话策略列表获取失败"
	ReqGetFStrategyListSuccessMsg = "会话策略列表获取成功"
	//
	ReqGetFStrategyFailMsg    = "会话策略信息获取失败"
	ReqGetFStrategySuccessMsg = "会话策略信息获取成功"
	//
	ReqDeleFStrategyFailMsg    = "会话策略信息删除失败"
	ReqDeleFStrategySuccessMsg = "会话策略信息删除成功"
	//
	//ReqUpdateStrategyFailMsg    = "策略信息更新失败"
	//ReqUpdateStrategySuccessMsg = "策略信息更新成功"
	//
	//
	//
	//ReqGetStrategyVehicleFailMsg    		= "策略-终端信息获取失败"
	//ReqGetStrategyVehicleUnExistMsg    		= "策略-终端信息不存在"
	//ReqGetStrategyVehicleSuccessMsg		 = "策略-终端信息获取成功"
	//
	//ReqGetStrategyVehicleListFailMsg    = "策略-终端列表获取失败"
	//ReqGetStrategyVehicleListSuccessMsg = "策略-终端列表获取成功"
	//
	//ReqGetStrategyVehicleResultListFailMsg    = "终端-学习结果列表获取失败"
	//ReqGetStrategyVehicleResultListSuccessMsg = "终端-学习结果列表获取成功"
)

//category
const (
	ReqAddCategoryFailMsg    = "添加指纹类别失败"
	ReqAddCategorySuccessMsg = "添加指纹类别成功"
	ReqCategoryExistMsg      = "该指纹类别已存在"
	ReqCategoryNotExistMsg   = "该指纹类别不存在"
	ReqCategoryFailMsg       = "获取指纹类别失败"
	ReqCategorySuccessMsg    = "获取指纹类别成功"

	ReqUpdateCategoryFailMsg    = "更新指纹类别失败"
	ReqUpdateCategorySuccessMsg = "更新指纹类别成功"

	ReqCategoryListFailMsg    = "获取指纹类别列表失败"
	ReqCategoryListSuccessMsg = "获取指纹类别列表成功"
)

//assetprint
const (
	ReqGetAssetFprintsUnExistMsg = "资产指纹信息不存在"
	ReqGetAssetFprintsFailMsg    = "资产指纹信息获取失败"
	ReqGetAssetFprintsSuccessMsg = "资产指纹信息获取成功"
)
