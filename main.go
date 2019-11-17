package main

import (
	"encoding/json"
	"fmt"
	"goland/config"
	"goland/mq"
	"goland/tts"
	"log"
)

// ProcessTransfer 通过 RabbitMQ 异步处理文字转语音
func ProcessTransfer(msg []byte) bool {
	// 1. 解析msg
	fmt.Println("MQ消费者解析消息:\n", string(msg))
	pubData := mq.TransferWords{}
	err := json.Unmarshal(msg, &pubData)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	fmt.Println("解析消息队列消息成功")

	// 2. 文字转语音播放
	resp, err := tts.Text2audio(pubData.Words)
	if resp != 200 || err != nil {
		fmt.Println("处理文字转语音失败!", err)
		return false
	}
	return true
}

func main() {
	fmt.Println("消息队列服务启动中, 开始监听转移任务队列...")
	mq.StartConsume(config.TransOSSQueueName, "transfer_oss", ProcessTransfer)
}