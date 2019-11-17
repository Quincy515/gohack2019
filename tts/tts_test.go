package tts

import (
	"fmt"
	"goland/config"
	"testing"
)

func TestText2audio(t *testing.T) {
	config.Text = "2019Gohack我们都是好孩子!"
	resp, err := Text2audio(config.Text)
	if resp != 200 || err != nil {
		fmt.Println("处理文字转语音失败!", err)
		return
	}
}
