package response

import "encoding/json"

type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}



func StructResponseObj(code int,msg string,data interface{}) interface{} {
	return Response{
		Code:code,
		Msg:msg,
		Data:data}
}

func StructResponseJson(code int,msg string,data interface{})  interface{}{
	obj:=StructResponseObj(code,msg,data)
	ret,err:=json.Marshal(obj)
	if err!=nil{
		return nil
	}
	return string(ret)
}
