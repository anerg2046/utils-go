package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Post 发起POST请求
func Post(url string, data interface{}, contentType string) ([]byte, error) {
	// 超时时间：15秒
	client := &http.Client{Timeout: 15 * time.Second}
	jsonStr, _ := json.Marshal(data)
	//fmt.Println(string(jsonStr))
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}
