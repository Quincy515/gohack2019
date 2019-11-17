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
	result, err := ocr.GeneralBasic("1.png")
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
	mq.StartConsume(config.TransOSSQueueName, "transfer_oss", tts.ProcessTransfer)
}
