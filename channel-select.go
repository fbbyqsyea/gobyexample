package main

import (
	"fmt"
	"time"
)

/// 通道选择器可以同时等待多个通道操作，go通道、协程和选择器的结合是go的一个强大的特性。

func main() {

	// 在我们的例子中，我们将从两个通道中选择。
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 各个通道将在若干时间后接收一个值，这个用来模拟例如并行的 Go 协程中阻塞的 RPC 操作
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 4)
		ch2 <- "two"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("received:", msg2)
	}

	// 我们使用 select 关键字来同时等待这两个值，并打印各自接收到的值。
	// 我们首先接收到值 "one"，然后就是预料中的 "two"了。
	// 注意从第一次和第二次 Sleeps 并发执行，总共仅运行了两秒左右。
	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case msg1 := <-ch1:
	// 		fmt.Println("received:", msg1)
	// 	case msg2 := <-ch2:
	// 		fmt.Println("received:", msg2)
	// 	}
	// }

}
