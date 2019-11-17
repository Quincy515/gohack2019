package redis

import (
	"fmt"
	"testing"
)

func TestRedisPool(t *testing.T) {
	// 1. 获得redis的一个连接
	rConn := RedisPool().Get()
	defer rConn.Close()
	fmt.Println(rConn)
}
