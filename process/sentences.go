package process

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goland/config"
	"goland/mq"
	"log"
	"strings"
)

// TransferData 转换数据的结构体
type TransferData struct {
	LogId          int64       `json:"log_id"`
	WordsResultNum int64       `json:"words_result_num"`
	WordsResult    []wordsList `json:"words_result"`
}

type wordsList struct {
	Words string `json:"words"`
}

// Sentences 图片转文字之后进一步结构化处理
func Sentences(text string) string {
	//fmt.Println("整理前：", text)
	pubData := TransferData{}
	err := json.Unmarshal([]byte(text), &pubData)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	//var buffer string
	var buffer bytes.Buffer
	for _, p := range pubData.WordsResult {
		//buffer += p.Words
		buffer.WriteString(p.Words)
	}

	sentences := buffer.String()
	// 去除特殊字符空格
	sentences = strings.Join(strings.Fields(sentences), "")
	//fmt.Println("整理后: ", sentences)

	// 测试: 通过 RabbitMQ 分段播放语音
	//size := len(sentences)
	//cnt := 0
	//for size > config.SplitSize {
	//	// 写入异步执行任务队列
	//	var data mq.TransferWords
	//	if config.SplitSize+cnt*config.SplitSize < len(sentences)-1 {
	//		data = mq.TransferWords{
	//			Words: sentences[cnt*config.SplitSize : config.SplitSize+cnt*config.SplitSize],
	//		}
	//	} else {
	//		data = mq.TransferWords{
	//			Words: sentences[cnt*config.SplitSize:],
	//		}
	//	}
	//	pubWords, _ := json.Marshal(data)
	//	pubSuc := mq.Publish(
	//		config.TransExchangeName,
	//		config.TransOSSRoutingKey,
	//		pubWords)
	//	if !pubSuc {
	//		// TODO: 当前发送转移信息失败, 稍后重试
	//		fmt.Println("RabbitMQ转移失败")
	//	} else {
	//		fmt.Printf("%s放入消息队列", pubWords)
	//	}
	//	cnt++
	//	size -= config.SplitSize
	//	fmt.Println("还剩余size:", size)
	//}

	// 通过 RabbitMQ 异步处理TTS
	data := mq.TransferWords{
		Words: sentences[10:], // 取出前面可能的特殊字符
	}

	pubWords, _ := json.Marshal(data)
	pubSuc := mq.Publish(
		config.TransExchangeName,
		config.TransOSSRoutingKey,
		pubWords)
	if !pubSuc {
		// TODO: 当前发送转移信息失败, 稍后重试
		fmt.Println("RabbitMQ转移失败")
	} else {
		fmt.Printf("放入消息队列成功：%s\n\n", pubWords)
	}
	return sentences
}
