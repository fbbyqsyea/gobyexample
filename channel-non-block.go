package main

import (
	"fmt"
	"strconv"
	"time"
)

// 常规的通过通道发送和接收数据是阻塞的。然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。

func main() {
	// 定义一个messages通道
	messages := make(chan string)

	// 使用goroutine异步，每个500ms向messages写入一条消息
	go func() {
		for i := 1; i < 5; i++ {
			time.Sleep(time.Millisecond * 500)
			messages <- "msg " + strconv.Itoa(i)
		}
	}()

	// 定义一个超时处理器
	timeout := time.After(time.Second * 3)

	for {
		// 如果在 messages 中存在，然后 select 将这个值带入 <-messages case中。如果不是，就直接到 default 分支中。
		select {
		case msg := <-messages:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("timeout 1")
			return
		default:
			time.Sleep(time.Millisecond * 300)
			fmt.Println("no message received")
		}
	}
	// 输出
	// no message received
	// no message received
	// msg 1
	// no message received
	// no message received
	// msg 2
	// no message received
	// no message received
	// msg 3
	// no message received
	// no message received
	// msg 4
	// no message received
	// no message received
	// timeout 1
}
