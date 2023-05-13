package main

import (
	"fmt"
	"time"
)

func main() {

	// 使用channel来同步goroutine之间的执行状态
	done := make(chan bool, 1)
	// 运行一个 worker Go协程，并给予用于通知的通道。
	go worker(done)
	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done
	// 如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker还没开始运行时就结束了。
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 发送一个值来通知我们已经完工啦。
	done <- true
}
