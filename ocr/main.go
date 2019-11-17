package ocr

import (
	"encoding/base64"
	"fmt"
	"goland/config"
	"goland/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ocr(baseUrl, path string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("获取文件失败:", err)
	}

	sourceString := base64.StdEncoding.EncodeToString(fileBytes)

	token, err := util.FetchToken(config.OCRApiKek, config.OCRSecretKey)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s?access_token=%s", baseUrl, token)
	//todo options参数抽出来
	params := url.Values{
		"image": {sourceString},
	}
	res, err := http.PostForm(urlStr, params)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	return body, err
}

func GeneralBasic(path string) ([]byte, error) {
	//通用文字识别
	generalBasic := "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	return ocr(generalBasic, path)
}

func main() {
	result, err := GeneralBasic("1.png")
	if err != nil {
		fmt.Println("图片识别文字失败", err)
		return
	}
	fmt.Println(string(result))
	config.Text = string(result)
}
