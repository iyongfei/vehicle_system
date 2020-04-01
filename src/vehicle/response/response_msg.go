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

//注册
const (

	ReqAddWhiteListFailMsg = "白名单添加失败"
	ReqAddWhiteListSuccessMsg = "白名单添加成功"
	ReqGetWhiteListSuccessMsg = "白名单获取成功"
	ReqGetWhiteListFailMsg = "白名单获取失败"
	//ReqRegistSuccessMsg = "用户注册成功"
	//ReqRegistExistMsg = "该用户已注册"
	//ReqRegistUnAuthMsg = "该用户未注册"
	//
	//ReqRegistAuthFailMsg = "该用户授权失败"
	//ReqRegistAuthSuccessMsg = "该用户授权成功"
)

