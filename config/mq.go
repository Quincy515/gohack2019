package config

import (
	"bytes"
	"goland/util"
)

const (
	AsyncTransferEnable  = true                             // 是否开启文件异步转移(默认false同步),这里设置成true为异步
	TransExchangeName    = "uploadserver.trans"             // 用于文件transfer的交换机
	TransOSSQueueName    = "uploadserver.trans.oss"         // oss转移队列名
	TransOSSErrQueueName = "uploadserver.trans.oss.err"     // oss转移失败后写入另一个队列的队列名
	TransOSSRoutingKey   = "oss"                            // routingkey
	AccessKeyID          = "LTAI4FoCTo2vbGoKvSDzd9DN"       // 阿里云访问key
	AccessKeySecret      = "hCqtRj3Pr2m9tHWowlvcQ303mjXMsD" // 阿里云访问secret
	ResourceOwnerId      = uint64(1703956476638309)         // 账户ID在安全设置
)

// 获取rabbitmq服务的入口url
// url格式amqp://账号:密码@rabbitmq服务器地址:端口号/vhost
func GetRabbitURL() string {
	var buffer bytes.Buffer
	username := util.GetUserName(AccessKeyID, ResourceOwnerId)
	password := util.GetPassword(AccessKeySecret)
	buffer.WriteString("amqp://")
	buffer.WriteString(username)
	buffer.WriteString(":")
	buffer.WriteString(password)

	// 从控制台获取接入点, 如果你使用的是杭州Region，那么Endpoint会形如1703956476638309.mq-amqp.cn-hangzhou-a.aliyuncs.com
	buffer.WriteString("@1703956476638309.mq-amqp.cn-hangzhou-a.aliyuncs.com:5672/uploadserver")
	RabbitURL := buffer.String()
	return RabbitURL
}
