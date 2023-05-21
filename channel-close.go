package main

import "fmt"

func main() {
	// 定义一个工作数据接收通道
	jobs := make(chan int, 5)
	// 定义一个多goroutine之间状态同步的通道
	done := make(chan bool, 0)

	// 开启goroutine异步监听jobs通道
	// 如果有数据输出数据
	// 如果通道关闭向状态同步通道同步关闭状态，然后退出当前goroutine
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 主goroutine向jobs写入数据
	for i := 0; i < 3; i++ {
		jobs <- i
	}
	// 关闭通道
	close(jobs)
	fmt.Println("sent all jobs")

	// 阻塞主goroutine 等待同步的状态 然后退出
	<-done
}
