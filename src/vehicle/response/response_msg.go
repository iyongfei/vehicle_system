package response

const (
	VStatusBadRequestMsg = "服务器请求失败,请尝试重新操作"
	ReqArgsIllegalMsg    = "服务器请求参数错误"
)

const (
	TrueFlag = "true"
	FalseFlag = "false"
)

const (
	OriginTypeSelf = 1
	OriginTypeSample = 2
	OriginTypeRule = 3
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

//threat
const (
	ReqGetThreatInfoFailMsg    = "威胁信息获取失败"
	ReqGetThreatInfoSuccessMsg = "威胁信息获取成功"
	ReqGetThreatInfoUnExistMsg = "威胁信息不存在"
	ReqDeleThreatInfoFailMsg   = "威胁信息删除失败"
)

//vehicle
const (
	ReqGetVehicleFailMsg    = "车载信息获取失败"
	ReqGetVehicleSuccessMsg = "车载信息获取成功"

	ReqGetVehicleUnExistMsg = "车载信息不存在"
	ReqGetVehicleExistMsg = "车载信息已存在"

	ReqDeleVehicleFailMsg      = "车载信息删除失败"
	ReqDeleVehicleSuccessMsg      = "车载信息删除成功"

	ReqUpdateVehicleFailMsg    = "车载信息更新失败"
	ReqUpdateVehicleSuccessMsg = "车载信息更新成功"

	ReqAddVehicleFailMsg    = "车载信息添加失败"
	ReqAddVehicleSuccessMsg = "车载信息添加成功"

	ReqGetVehiclesFailMsg    = "车载列表获取失败"
	ReqGetVehiclesSuccessMsg = "车载列表获取成功"
)
//depolyer
const (
	ReqGetVehicleBindLeaderUnExistMsg = "车载没有绑定管理人员信息"
	ReqGetVehicleBindLeaderFailMsg = "车载绑定管理人员信息获取失败"
	ReqUpdateVehicleBindLeaderSuccessMsg = "车载绑定管理人员信息更新成功"
)

//portMap
const (
	ReqGetPortMapUnExistMsg = "端口映射信息不存在"
	ReqGetPortMapFailMsg = "端口映射信息获取失败"
	ReqGetPortMapSuccessMsg = "端口映射信息获取成功"
	ReqUpdatePortMapSuccessMsg = "端口映射信息更新成功"
	ReqUpdatePortMapFailMsg = "端口映射信息更新失败"
)

//asset
const (
	ReqGetAssetFailMsg    = "设备信息获取失败"
	ReqGetAssetSuccessMsg = "设备信息获取成功"

	ReqGetAssetUnExistMsg = "设备信息不存在"
	ReqGetAssetExistMsg = "设备信息已存在"

	ReqDeleAssetFailMsg      = "设备信息删除失败"
	ReqDeleAssetSuccessMsg      = "设备信息删除成功"

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
	ReqGetStrategyExistMsg    = "策略已存在"
	ReqGetStrtegyUnExistMsg = "策略信息不存在"

	ReqGetStrategyListFailMsg    = "策略列表获取失败"
	ReqGetStrategyListSuccessMsg = "策略列表获取成功"

	ReqGetStrategyFailMsg    		= "策略信息获取失败"
	ReqGetStrategySuccessMsg		 = "策略信息获取成功"

	ReqDeleStrategyFailMsg          = "策略信息删除失败"
	ReqDeleStrategySuccessMsg       = "策略信息删除成功"

	ReqUpdateStrategyFailMsg    = "策略信息更新失败"
	ReqUpdateStrategySuccessMsg = "策略信息更新成功"



	ReqGetStrategyVehicleFailMsg    		= "策略-车载信息获取失败"
	ReqGetStrategyVehicleUnExistMsg    		= "策略-车载信息不存在"
	ReqGetStrategyVehicleSuccessMsg		 = "策略-车载信息获取成功"

	ReqGetStrategyVehicleListFailMsg    = "策略-车载列表获取失败"
	ReqGetStrategyVehicleListSuccessMsg = "策略-车载列表获取成功"

	ReqGetStrategyVehicleResultListFailMsg    = "车载-学习结果列表获取失败"
	ReqGetStrategyVehicleResultListSuccessMsg = "车载-学习结果列表获取成功"
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
	ReqGetFStrategyFailMsg    		= "会话策略信息获取失败"
	ReqGetFStrategySuccessMsg		 = "会话策略信息获取成功"
	//
	ReqDeleFStrategyFailMsg          = "会话策略信息删除失败"
	ReqDeleFStrategySuccessMsg       = "会话策略信息删除成功"
	//
	//ReqUpdateStrategyFailMsg    = "策略信息更新失败"
	//ReqUpdateStrategySuccessMsg = "策略信息更新成功"
	//
	//
	//
	//ReqGetStrategyVehicleFailMsg    		= "策略-车载信息获取失败"
	//ReqGetStrategyVehicleUnExistMsg    		= "策略-车载信息不存在"
	//ReqGetStrategyVehicleSuccessMsg		 = "策略-车载信息获取成功"
	//
	//ReqGetStrategyVehicleListFailMsg    = "策略-车载列表获取失败"
	//ReqGetStrategyVehicleListSuccessMsg = "策略-车载列表获取成功"
	//
	//ReqGetStrategyVehicleResultListFailMsg    = "车载-学习结果列表获取失败"
	//ReqGetStrategyVehicleResultListSuccessMsg = "车载-学习结果列表获取成功"
)