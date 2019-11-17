package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TTSTokenResp 百度语音识别返回的数据结构
type RespMsg struct {
	AccessToken   string `json:"access_token"`
	ExpiresIn     int64  `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	SessionSecret string `json:"session_secret"`
}

// FetchToken 获取百度语音识别token
func FetchToken(apikey, secret string) (string, error) {
	// 1. 发起Get请求
	url := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s",
		apikey, secret)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("通过API Key和Secret Key获取百度Access_token失败:", err.Error())
		return "", err
	}
	// 2. 读取请求返回的结果
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		fmt.Printf("获取返回值发生错误: %s\n", err)
		return "", err
	}
	// 3. 转换为结构体
	ubody := &RespMsg{}
	if err := json.Unmarshal(body, ubody); err != nil {
		fmt.Println("认证失败")
		return "", err
	}
	return ubody.AccessToken, nil
}
