package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


/**
模拟表单请求
 */
func PostForm(urlParam string,bodyParms map[string]interface{}) (map[string]interface{}, error) {
	u,err:=url.Parse(urlParam)

	if err!=nil{
		return nil,err
	}
	q := u.Query()
	for k,v:= range bodyParms{
		q.Set(k,v.(string))
	}

	resp, err := http.PostForm(u.String(), q)
	if err!=nil{
		return nil,err
	}

	buf, err := ioutil.ReadAll(resp.Body)

	if err!=nil{
		return nil,err
	}

	var p map[string]interface{}
	err = json.Unmarshal(buf, &p)
	if err!=nil{
		return nil,err
	}
	return p,nil
}

func Get(url string,token string)(map[string]interface{}, error) {
	client := http.Client{}

	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("token", token)

	if err != nil {
		return nil,err
	}
	rsp, _ := client.Do(reqest)

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil,err
	}
	var p map[string]interface{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}


func Post(url string, body map[string]interface{},token string)(map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	reqest, err := http.NewRequest("POST", url, bytes.NewReader(jsonBytes))
	reqest.Header.Add("token", token)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rsp, _ := client.Do(reqest)
	if err != nil {
		return nil,err
	}
	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil,err
	}
	var p map[string]interface{}
	err = json.Unmarshal(buf, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}


func put(url string, reqBody string) {
	fmt.Println("PUT REQ...")
	fmt.Println("REQ:", reqBody)
	req, err := http.NewRequest("PUT", url, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RSP:", string(body))
}

func delete(url string, reqBody string) {
	fmt.Println("DELETE REQ...")
	fmt.Println("REQ:", reqBody)
	req, err := http.NewRequest("DELETE", url, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RSP:", string(body))

}

