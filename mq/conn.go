package mq

import (
	"github.com/streadway/amqp"
	"goland/config"
	"log"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

// 如果异常关闭，会接收通知
var notifyClose chan *amqp.Error

// 申请队列，保证队列存在，消息能发送到队列中
func init() {
	// 是否开启异步转移功能，开启时才初始化RabbitMQ连接
	if !config.AsyncTransferEnable {
		return
	}
	if initChannel() {
		channel.NotifyClose(notifyClose)
	}

	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channel = nil
				log.Printf("onNotifyChannelClosed: %+v\n", msg)
				initChannel()
			}
		}
	}()
}

// initChannel: 初始化Channel连接
func initChannel() bool {
	// 1. 判断channel是已经创建过
	if channel != nil {
		return true
	}

	// 2. 获得rabbitmq的一个连接
	RabbitURL := config.GetRabbitURL()
	conn, err := amqp.Dial(RabbitURL) // 完成connection的创建及赋值
	if err != nil {
		log.Println(err.Error())
		return false
	}
	//failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	// 3. 打开一个channel, 用于消息的发布与接收等
	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}
