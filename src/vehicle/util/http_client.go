package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//import (
//	"bytes"
//	"encoding/json"
//	"github.com/pkg/errors"
//	"io/ioutil"
//	"net/http"
//	"time"
//)

func PostJSON(url string, body map[string]interface{}) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("%s json marshal err:%v",RunFuncName(),err.Error())
	}

	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Post(url, "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("%s http err:%v",RunFuncName(),err.Error())
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s read body err %v",RunFuncName(),err.Error())
	}

	var p map[string]interface{}
	err = json.Unmarshal(buf, &p)
	if err != nil {
		return nil, fmt.Errorf("%s json unmarshal err %v",RunFuncName(),err.Error())
	}

	return p, nil
}
