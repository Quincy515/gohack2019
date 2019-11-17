package ocr

import (
	"fmt"
	"goland/config"
	"testing"
)

func TestOcr(t *testing.T) {
	result, err := GeneralBasic("1.png")
	if err != nil {
		fmt.Println("图片识别文字失败", err)
		return
	}
	fmt.Println(string(result))
	config.Text = string(result)
}
