package main

import "fmt"

func main() {
	// 关闭通道意味着不能再往通道里面推送数据
	// 这个特性可以给通道的接收方传达工作完成的信号
	jobs := make(chan int, 5)
	done := make(chan bool, 0)

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

	for i := 0; i < 3; i++ {
		jobs <- i
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
