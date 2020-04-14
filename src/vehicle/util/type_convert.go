package util

import (
	"encoding/json"
	"reflect"
)

/**
结构体====>map
 */

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

/**
结构体指针====>map
 */

func StructPtr2Map(obj interface{}) map[string]interface{} {
	obj_v := reflect.ValueOf(obj)
	v := obj_v.Elem()
	typeOfType := v.Type()
	var data = make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		data[typeOfType.Field(i).Name] = field.Interface()
	}
	return data
}


/**
转结体，指针转====>json
 */
func Struct2Json(obj interface{}) string{
	jsonBys,err := json.Marshal(obj)
	if err!=nil{
		return ""
	}
	return string(jsonBys)
}


/**
json====>结构体
 */

func Json2Struct(jsonStr interface{},obj interface{}){
	err := json.Unmarshal([]byte(jsonStr.(string)),obj)
	if err!=nil{
		obj = nil
	}
}

/**
json====>map
 */
func Json2Map(jsonStr interface{}) interface{}{

	var tempMap map[string]interface{}

	err:=json.Unmarshal([]byte(jsonStr.(string)),&tempMap)

	if err!=nil{
		return nil
	}
	return tempMap
}

/**
map====>json
 */


func Map2Json(mapData interface{}) interface{}{
	jsonBys,err := json.Marshal(mapData)
	if err!=nil{
		return ""
	}
	return string(jsonBys)
}
/**
map转结构体
 */

func Map2Struct(mapData interface{}) interface{}{
	return nil
}