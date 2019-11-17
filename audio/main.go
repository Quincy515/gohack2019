package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
	"log"
	"time"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	rPool "goland/redis"
)

// 播放音频 服务
func main() {
	var err error

	var dec *minimp3.Decoder
	var data, file []byte
	var player *oto.Player

	if file, err = ioutil.ReadFile("../go.mp3"); err != nil {
		log.Fatal(err)
	}

	if dec, data, err = minimp3.DecodeFull(file); err != nil {
		log.Fatal(err)
	}

	if player, err = oto.NewPlayer(dec.SampleRate, dec.Channels, 2, 1024); err != nil {
		log.Fatal(err)
	}

	player.Write(data)

	// 2. 获得redis的一个连接
	rConn := rPool.RedisPool().Get()
	defer rConn.Close()
	psc := redis.PubSubConn{Conn: rConn}
	psc.Subscribe("audio")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			// process pushed message
			if file, err = ioutil.ReadFile("../"+string(v.Data)); err != nil {
				fmt.Println("读取来自redis的音频失败", err)
			}

			if dec, data, err = minimp3.DecodeFull(file); err != nil {
				fmt.Println("解析来自redis的音频失败", err)

			}
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println("接收redis消息失败: ", err)
		}

		_, _ = player.Write(data)
	}

	<-time.After(time.Second)
	dec.Close()
	player.Close()
}
