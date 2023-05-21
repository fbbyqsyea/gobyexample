package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 定义双向通道c1和c2
	c1 := make(chan string)
	c2 := make(chan string)

	// 开启goroutine 每个200ms向c1输入一个值 执行5次
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond * 200)
			c1 <- "result " + strconv.Itoa(i)
		}
	}()

	// 开启goroutine 每个300ms向c2输入一个值 执行5次
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond * 300)
			c2 <- "result " + strconv.Itoa(i)
		}
	}()

	// 定义一个独立的定时器
	timeout := time.After(time.Second * 1)

	// for循环可以一直启用select监听c1和c2通道
	for {
		select {
		case res1 := <-c1: // 输出c1通道接收到的值
			fmt.Println("received from chan c1: ", res1)
		case res2 := <-c2: // 输出c2通道接收到的值
			fmt.Println("received from chan c2: ", res2)
		case <-timeout: // 使用time.After在1s后让通道选择器执行超时退出操作
			fmt.Println("timeout after 1s")
			return
		}
	}

	// 输出
	// received from chan c1:  result 1
	// received from chan c2:  result 1
	// received from chan c1:  result 2
	// received from chan c1:  result 3
	// received from chan c2:  result 2
	// received from chan c1:  result 4
	// received from chan c2:  result 3
	// timeout after 1s
}
