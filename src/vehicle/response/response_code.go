package response

/**
状态码
 */
const (
	//请求成功,数据返回
	VStatusOK   = 2000
	//请求成功,尚未处理,在异步操作最适合
	VStatusOKAccepted  = 2002
	//服务器成功处理了请求，但不需要返回任何实体内容
	VStatusNoContent   = 2004

	//表示请求的资源已经永久的搬到了其他位置,就是说资源已经被分配了新的URI
	VStatusMovedPermanently = 3001
	//和3001很像，只不过资源是临时移动，资源在将来可能还会改变
	VStatusFound = 3002

	//表示请求报文存在语法错误或参数错误
	VStatusBadRequest   = 4000
	//表示发送的请求需要有HTTP认证信息或者是认证失败了
	VStatusUnauthorized   = 4001
	//表示对请求资源的访问被服务器拒绝了
	VStatusForbidden = 4003
	//表示服务器找不到你请求的资源
	VStatusNotFound = 4004
	//请求方法异常
	VStatusMethodNotAllowed = 4005

	//表示服务器执行请求的时候出错了
	VStatusServerError   = 5000
	//作为网关或者代理工作的服务器尝试执行请求时，从上游服务器接收到无效的响应
	VStatusBadGateway = 5002
	//能及时从上游服务器收到响应。
	VStatusGatewayTimeout = 5004
)



