package mq

import (
	"github.com/streadway/amqp"
)

// Publish: 发布消息
func Publish(exchange, routingKey string, msg []byte) bool {
	// 1. 判断channel是否正常
	if !initChannel() {
		return false
	}

	// 2. 执行消息发布动作-发送消息到队列中
	if nil == channel.Publish(
		exchange,
		routingKey,
		false, /*如果没有对应的queue, 就会丢弃这条消息,
		如果为true,根据exchange类型和routingKey规则,
		如果无法找到符合条件的队列那么会把发送的消息返回给发送者*/
		false, /*如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者,
		则会把消息返还给发送者*/
		amqp.Publishing{ // 发送的信息
			ContentType: "text/plain",
			Body:        msg}) {
		return true
	}
	return false
}
