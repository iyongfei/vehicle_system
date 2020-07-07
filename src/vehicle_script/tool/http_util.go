package tool

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

func UploadFile(url string, token string, params map[string]string, nameField, fileName string, file io.Reader) (map[string]interface{}, error) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("token", token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var p map[string]interface{}
	err = json.Unmarshal(content, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func UploadEditFile(url string, token string, params map[string]string, nameField, fileName string, file io.Reader) (map[string]interface{}, error) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("token", token)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var p map[string]interface{}
	err = json.Unmarshal(content, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetFile(reqUrl string, queryParams map[string]interface{}, token string) (*http.Response, error) {
	urlReq, _ := url.Parse(reqUrl)

	params := url.Values{}
	for k, v := range queryParams {
		params.Set(k, v.(string))
	}

	urlReq.RawQuery = params.Encode()

	client := http.Client{}

	reqest, err := http.NewRequest("GET", urlReq.String(), nil)
	reqest.Header.Add("token", token)

	if err != nil {
		return nil, err
	}
	rsp, _ := client.Do(reqest)

	return rsp, nil
}
func Get(reqUrl string, queryParams map[string]interface{}, token string) (map[string]interface{}, error) {
	urlReq, _ := url.Parse(reqUrl)

	params := url.Values{}
	for k, v := range queryParams {
		params.Set(k, v.(string))
	}

	urlReq.RawQuery = params.Encode()

	client := http.Client{}

	reqest, err := http.NewRequest("GET", urlReq.String(), nil)
	reqest.Header.Add("token", token)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return nil, err
	}
	rsp, _ := client.Do(reqest)

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	var p map[string]interface{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
func Delete(reqUrl string, queryParams map[string]interface{}, token string) (map[string]interface{}, error) {
	urlReq, _ := url.Parse(reqUrl)

	params := url.Values{}
	for k, v := range queryParams {
		params.Set(k, v.(string))
	}

	urlReq.RawQuery = params.Encode()

	client := http.Client{}

	reqest, err := http.NewRequest("DELETE", urlReq.String(), nil)
	reqest.Header.Add("token", token)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return nil, err
	}
	rsp, _ := client.Do(reqest)

	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	var p map[string]interface{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

/**
模拟表单请求
*/
func PostForm(urlParam string, bodyParms map[string]interface{}, token string) (map[string]interface{}, error) {
	u, err := url.Parse(urlParam)

	if err != nil {
		return nil, err
	}
	q := u.Query()
	for k, v := range bodyParms {

		q.Set(k, v.(string))
	}

	client := http.Client{}

	reqest, err := http.NewRequest("POST", urlParam, strings.NewReader(q.Encode()))
	reqest.Header.Add("token", token)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rsp, _ := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	//defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	var p map[string]interface{}
	err = json.Unmarshal(buf, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func PutForm(urlParam string, bodyParms map[string]interface{}, token string) (map[string]interface{}, error) {
	u, err := url.Parse(urlParam)

	if err != nil {
		return nil, err
	}
	q := u.Query()
	for k, v := range bodyParms {
		q.Set(k, v.(string))
	}

	client := http.Client{}

	reqest, err := http.NewRequest("PUT", urlParam, strings.NewReader(q.Encode()))
	reqest.Header.Add("token", token)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rsp, _ := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return nil, err
	}

	var p map[string]interface{}
	err = json.Unmarshal(buf, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func PostJson(url string, body interface{}, token string) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	reqest, err := http.NewRequest("POST", url, bytes.NewReader(jsonBytes))
	reqest.Header.Add("token", token)
	reqest.Header.Set("Content-Type", "application/json")

	rsp, _ := client.Do(reqest)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	var p map[string]interface{}
	err = json.Unmarshal(buf, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
