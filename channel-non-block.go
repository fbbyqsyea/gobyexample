package main

import "fmt"

// 常规的通过通道发送和接收数据是阻塞的。然而，我们可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select。

func main() {
	messages := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("no message received")
	}
}
