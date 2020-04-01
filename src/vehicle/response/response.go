package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

/**
返回对象
 */
func StructResponseObj(code int,msg string,data interface{}) interface{} {
	return Response{
		Code:code,
		Msg:msg,
		Data:data}
}

func StructResponseMap(code int,msg string,data interface{}) interface{} {
	return gin.H{
		"code":code,
		"msg":msg,
		"data":data,
	}
}



/**
返回json字符串
 */
func StructResponseJson(code int,msg string,data interface{})  interface{}{
	obj:=StructResponseObj(code,msg,data)
	ret,err:=json.Marshal(obj)
	if err!=nil{
		return nil
	}
	return string(ret)
}
