package response

var (
	VStatusBadRequestMsg = "服务器请求失败,请尝试重新操作"
)


const (
	TokenExpiredStr = "token过期"
	TokenNotValidYetStr = "token未激活"
	TokenMalformedStr = "token不合法"
	TokenInvalidStr = "token未知"
	ValidationErrorUnverifiableStr = "签名信息错误，无法验证token"
	ValidationErrorSignatureInvalidStr = "签名验证失败"
)

const (
	AuthTokenLost     = "请求未携带token，权限不足"
	AuthTokenResignin = "请重新登录"
)