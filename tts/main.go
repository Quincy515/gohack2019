package tts

import (
	"encoding/json"
	"errors"
	"fmt"
	"goland/config"
	rPool "goland/redis"
	"goland/util"
	"io/ioutil"
	"net/http"
)

var (
	formats = map[int]string{3: "mp3", 4: "pcm-16k", 5: "pcm-8k", 6: "wav"}
	format  = formats[config.AUE]
	token   string
	//ctx     *oto.Context // 如果已有音频正在播放，则关闭后开始播放新音频
)

// Text2audio 文字转音频
func Text2audio(text string) (int, error) {
	token, err := util.FetchToken(config.TTSApiKek, config.TTSSecretKey)
	if err != nil {
		return 500, err
	} else if token == "" {
		return 502, nil
	}

	//fmt.Println("text:", text)
	urlT2a := "http://tsn.baidu.com/text2audio?tex=" + text + "&lan=zh&cuid=" + config.CUID + "&ctp=1&tok=" + token
	//fmt.Println("urlT2a: ", urlT2a)

	resp, err := http.Get(urlT2a)
	if err == nil {
		contentType := resp.Header.Get("Content-Type")
		if contentType == "application/json" {
			// 返回json失败处理
			type ErrResp struct {
				ErrMsg string `json:"err_msg"`
				ErrNo  int    `json:"err_no"`
			}
			var respContent ErrResp

			body, err := ioutil.ReadAll(resp.Body)
			if err == nil && resp.StatusCode == 200 {
				_ = json.Unmarshal(body, &respContent)
			}
			_ = resp.Body.Close()
		} else if contentType == "audio/mp3" {
			// 返回audio mp3
			var audio []byte
			audio, err := ioutil.ReadAll(resp.Body)
			// 1. 计算sha1
			fileHash := util.Sha1(audio)
			filename := fileHash + ".mp3"
			if err == nil {
				// write to file
				err = ioutil.WriteFile(filename, audio, 0666)
				if err != nil {
					fmt.Println("存储音频文件失败!", err)
					return 500, err
				}
				fmt.Println("存储音频文件成功！")
				// 2. 获得redis的一个连接
				rConn := rPool.RedisPool().Get()
				defer rConn.Close()
				// 3. 发布redis消息
				reply, err := rConn.Do("PUBLISH", "audio", filename)
				if err != nil {
					fmt.Println("发布redis消息失败: ", err)
				}
				fmt.Println("发送redis消息成功: ", reply)
			}
			_ = resp.Body.Close()
		} else {
			err = errors.New("Unknown Content-Type ! =>" + contentType + " | url: " + urlT2a)
		}
	}
	return 200, err
}

