package response

var (
	VStatusBadRequestMsg  = "服务器请求失败,请尝试重新操作"
	ReqArgsIllegalMsg     = "服务器请求参数错误"
)


const (
	TokenExpiredStr							= "token过期"
	TokenNotValidYetStr                     = "token未激活"
	TokenMalformedStr                       = "token不合法"
	TokenInvalidStr 					    = "token未知"
	ValidationErrorUnverifiableStr          = "签名信息错误，无法验证token"
	ValidationErrorSignatureInvalidStr      = "签名验证失败"
)

const (
	AuthTokenLost     = "请求未携带token，权限不足"
	AuthTokenResignin = "请重新登录"
)

//分组
const (
	UnGroupName = "未分组"
	ReqAddGroupExistMsg = "分组已存在"
	ReqAddGroupSuccessMsg = "分组添加成功"
	ReqAddGroupFailMsg = "分组添加失败"
)
//注册
const (

	PasswordSecret = "vgw-1214-pwd-key"
	ReqRegistFailMsg = "用户注册失败"
	ReqRegistSuccessMsg = "用户注册成功"
	ReqRegistExistMsg = "该用户已注册"
	ReqRegistUnAuthMsg = "该用户未注册"

	ReqRegistAuthFailMsg = "该用户授权失败"
	ReqRegistAuthSuccessMsg = "该用户授权成功"
)

//白名单
const (

	ReqAddWhiteListFailMsg = "白名单添加失败"
	ReqAddWhiteListSuccessMsg = "白名单添加成功"
	ReqGetWhiteListSuccessMsg = "白名单获取成功"
	ReqGetWhiteListFailMsg = "白名单获取失败"
	ReqGetWhiteListUnExistMsg = "白名单不存在"
	ReqDeleWhiteListFailMsg = "白名单删除失败"
	ReqDeleWhiteListSuccessMsg = "白名单删除成功"
	ReqUpdateWhiteListSuccessMsg = "白名单更新成功"
	ReqUpdateWhiteListFailMsg = "白名单更新失败"
	//ReqRegistSuccessMsg = "用户注册成功"
	//ReqRegistExistMsg = "该用户已注册"
	//ReqRegistUnAuthMsg = "该用户未注册"
	//ReqRegistAuthFailMsg = "该用户授权失败"
	//ReqRegistAuthSuccessMsg = "该用户授权成功"
)
//
const (
	ReqGetFlowFailMsg = "会话消息获取失败"
	ReqGetFlowSuccessMsg = "会话消息获取成功"
	ReqGetFlowUnExistMsg = "会话消息不存在"
	ReqGetFlowExistMsg = "会话消息已存在"
	ReqAddFlowFailMsg = "会话消息添加失败"
	ReqAddFlowSuccessMsg = "会话消息添加成功"
	ReqEditFlowSuccessMsg = "会话消息更新成功"
	ReqEditFlowFailMsg = "会话消息更新失败"
	ReqDeleFlowFailMsg = "会话消息删除失败"
)

//threat
const (
	ReqGetThreatInfoFailMsg = "威胁信息获取失败"
	ReqGetThreatInfoSuccessMsg = "威胁信息获取成功"
	ReqGetThreatInfoUnExistMsg = "威胁信息不存在"
	ReqDeleThreatInfoFailMsg = "威胁信息删除失败"
)

//
const (
	ReqGetVehicleFailMsg = "车载信息获取失败"
	ReqGetVehicleSuccessMsg = "车载信息获取成功"
	ReqGetVehicleUnExistMsg = "车载信息不存在"
	ReqDeleVehicleFailMsg = "车载信息删除失败"
	ReqUpdateVehicleFailMsg = "车载信息更新失败"
	ReqUpdateVehicleSuccessMsg = "车载信息更新成功"
)