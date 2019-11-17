# 介绍
2019 Go Hack 马拉松

所见即所得

整个过程模拟人眼看到的内容，拍照成图片，图片转文字，文字转语音，语音播放

# 模块划分
- Minio云存储图片          使用Minio云存储
- 图片转文字 文字识别服务    使用百度AI ocr服务
- 处理文字发送给MQ          使用阿里云AMQP服务
- 文字合成语音服务          使用百度AI tts服务
- 语音文件通过Redis发布     使用公有Redis服务
- 通过Redis订阅播放语音     独立部署

# 使用说明

1. 使用的开源第三方库，感谢以下的开源支持

```bash
	go get github.com/garyburd/redigo 
	go get github.com/hajimehoshi/oto 
	go get github.com/streadway/amqp 
	go get github.com/tosone/minimp3
	go get github.com/minio/minio-go/v6 
```

2. 项目启动 

开启消息队列服务文字转语言
```bash
go build main.go // 在项目根目录中
./main // 启动main

```

开启从云储存获取图片服务
```bash
cd store
go build main.go 
./main
```

开启语音播放服务
```bash
cd audio
go build main.go 
./main
```

# 项目目录
```
├── audio
│   └── main.go         通过Redis发布订阅播放音频
├── config              配置中心
│   ├── global.go       全局配置
│   ├── mq.go           消息队列配置
│   ├── ocr.go          百度 OCR API 配置
│   └── tts.go          百度 TTS API 配置
├── main.go             主程序 处理图片转文字，通过MQ文字转语音
├── mq                  消息队列
│   ├── conn.go         连接
│   ├── consumer.go     消费者
│   ├── define.go       结构体定义
│   └── producer.go     发布者
├── ocr                 百度图片文字识别
│   ├── main.go         处理图片转文字
│   └── ocr_test.go
├── process             文字处理，去除空格，发送给MQ
│   └── sentences.go
├── redis               Redis
│   ├── conn.go         连接
│   └── conn_test.go
├── tts                 百度文字语音合成
│   ├── main.go         处理文字合成语音功能
│   └── tts_test.go
└── util                工具
    ├── fetchSha1.go    给生成的音频文件进行Sha1
    ├── fetchToken.go   获取百度AI token
    └── fetchUser.go    获取Aliyun AMQP用户
```

# 进度说明
- [x] Minio云储存图片          使用Minio云存储服务✔️
- [x] 图片转文字 文字识别服务    使用百度AI ocr服务 ✔️
- [x] 处理文字发送给MQ          使用阿里云AMQP服务 ✔️
- [x] 文字合成语音服务          使用百度AI tts服务 ✔️
- [x] 语音文件通过Redis发布     使用公有Redis服务  ✔️
- [x] 通过Redis订阅播放语音     独立部署          ✔️
- [ ] 生成的语音文件存放公有云    跨机器播放语音
- [ ] 自然语言处理，过滤广告等不必要文字            
- [ ] 语音交互，如语音唤醒、语音识别、语音控制       
- [ ] 更多功能，更多场景，敬请期待...

# 共同学习
> 该项目是2019 Go Hack 马拉松项目，因为能力有限，希望有小伙伴可以一起参与共同学习进步
1. Fork 本仓库
2. 新建分支
3. 提交代码
4. 新建 Pull Request

> 期待与您一起学习进步，下面是我的微信二维码：

![微信二维码](./wechat.png)