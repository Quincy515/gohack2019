package main

import (
	"fmt"
	"goland/config"
	"goland/mq"
	"goland/ocr"
	"goland/process"
	"goland/tts"
)

func main() {
	// 1. 图片转文字
	result, err := ocr.GeneralBasic("2.png")
	if err != nil {
		fmt.Println("图片识别文字失败", err)
		return
	}
	if string(result) == "" {
		fmt.Println("未识别到文字", err)
		return
	}

	// 2. 文字处理
	config.Text = process.Sentences(string(result))
	//config.Text = "2019年够骇客马拉松，大家真棒！"
	mq.StartConsume(config.TransOSSQueueName, "transfer_oss", tts.ProcessTransfer)
	// 3. 文字转语音播放
	//text:=strings.Join(strings.Fields(config.Text), "")
	//resp, err := tts.Text2audio(text)
	//if resp != 200 || err != nil {
	//	fmt.Println("处理文字转语音失败!", err)
	//	return
	//}
	//
	//if err := tts.PlayAudio(); err != nil {
	//	fmt.Println("播放mp3失败!", err)
	//	return
	//}
}
