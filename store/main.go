package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v6"
	"goland/config"
	"goland/ocr"
	"goland/process"
	"log"
	"time"
)

func main() {
	endpoint := "49.233.186.203:9000"
	accessKeyID := "S9J7W4WCMUMJ3RS0RVP0"
	secretAccessKey := "6xnkmYaYYbeBavLrb+k6L0Xxg2hkcEkHWGpCmz+J"
	useSSL := false

	// Initialize minio client object.
	// 1. 建立连接
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		fmt.Println("minioClient连接失败", err)
	}

	fmt.Printf("%#v\n", minioClient) // minioClient is now setup

	// 2. 列出bucket
	bucketName := "gohack2019"

	var pngList []string
	// 3. 获取该bucket里所有的object
	for {
		// Create a done channel to control 'ListObjects' go routine.
		doneCh := make(chan struct{})

		// Indicate to our routine to exit cleanly upon return.
		defer close(doneCh)

		// List all objects from a bucket-name with a matching prefix.
		for object := range minioClient.ListObjectsV2(bucketName, "2019", true, doneCh) {
			if object.Err != nil {
				fmt.Println(object.Err)
				return
			}
			fmt.Println("获取该bucket里所有存储对象", object.Key)
			pngList = append(pngList, object.Key)
		}
		fmt.Println("pngList: ", pngList)

		// 4. 下载图片并删除云存储
		for _, png := range pngList {
			// 下载
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
			defer cancel()

			if err := minioClient.FGetObjectWithContext(ctx, bucketName, png, png, minio.GetObjectOptions{}); err != nil {
				fmt.Println("获取图片文件失败: ", err)
				return
			}
			fmt.Println("Successfully saved ", png)

			// 删除
			err = minioClient.RemoveObject(bucketName, png)
			if err != nil {
				fmt.Println("删除图片云文件失败: ", err)
				return
			}
			log.Println("Success delete")
		}
		time.Sleep(time.Second*10)

		for len(pngList)>0 {
			png := pngList[0]
			pngList = pngList[1:]
			// 1. 读取本地图片进行文字识别
			result, err := ocr.GeneralBasic(png)
			if err != nil {
				fmt.Println("图片识别文字失败", err)
				return
			}
			if string(result) == "" {
				fmt.Println("未识别到文字", err)
				return
			}
			fmt.Println("图片识别成功")

			// 2. 文字处理
			config.Text = process.Sentences(string(result))
			//config.Text = "2019年够骇客马拉松，大家真棒！"
			//mq.StartConsume(config.TransOSSQueueName, "transfer_oss", tts.ProcessTransfer)
		}
	}
}
