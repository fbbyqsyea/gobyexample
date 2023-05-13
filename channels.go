package main

import "fmt"

func main() {
	// 通道是连接多个go协程的管道，可以在一个go协程中发送数据到通道，在另外一个协程中接收数据

	// 使用 make(chan val-type)来创建一个新的通道。
	messages := make(chan string)
	// 使用channel <- 语法发送一个新值到通道，这里在协程中发送数据到messages通道
	go func() {
		messages <- "ping"
	}()
	// 使用<-channel语法从channel中接收一个值，并打印出来
	message := <-messages
	fmt.Println(message)

	// 默认发送和接收操作都是阻塞的，直到发送方和接收方都准备完毕，这个特性可以让我们在不使用其他的同步操作的情况下，在程序结尾的时候等待接收ping消息
}
