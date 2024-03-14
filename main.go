package main

import (
	"fmt"
	"rateLimiting/api"
	"time"
)

func main() {
	// 漏桶算法测试
	lb := api.NewLeakyBucket(5, 1) // 容量为5，速率为1（每秒1个令牌）
	for i := 0; i < 10; i++ {
		fmt.Printf("lb:%+v\n", lb)
		fmt.Println("Leaky Bucket:", lb.AllowRequest())
		time.Sleep(time.Second / 5)
	}

	// 令牌桶算法测试
	tb := api.NewTokenBucket(5, 1) // 容量为5，速率为1（每秒1个令牌）
	for i := 0; i < 10; i++ {
		fmt.Println("Token Bucket:", tb.AllowRequest())
		time.Sleep(time.Second / 2)
	}

	// 滑动窗口算法测试
	sw := api.NewSlidingWindow(3, time.Second*5) // 5秒内最多允许3个请求
	for i := 0; i < 10; i++ {
		fmt.Println("Sliding Window:", sw.AllowRequest())
		time.Sleep(time.Second / 2)
	}
}
