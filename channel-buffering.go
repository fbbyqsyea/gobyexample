package main

import "fmt"

func main() {
	// 默认通道是没有缓冲的，这意味着只有在对于的接收通道准备接收数据时，才能进行发送操作。
	// 缓冲通道允许，在没有接收方的情况下，缓存一定数量的值

	// 创建一个可以缓冲两个值的通道
	messages := make(chan string, 2)

	// 向通道中发送2个值
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
