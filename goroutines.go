package main

import (
	"fmt"
	"time"
)

func main() {

	// 直接调用
	f("direct")
	// go run goroutines.go
	// direct : 0
	// direct : 1
	// direct : 2

	// 在一个go的协程中调用这个函数，这个新的go协程会并行执行这个函数调用
	go f("goroutine")

	// 我们也可以定义一个匿名函数
	go func(from string) {
		f(from)
	}("goroutine1")
	// 由于是异步执行，所以需要等待执行结束
	// time.Sleep(time.Second)
	// 同时也可以监听用户输入事件，回车结束
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")

	// direct : 0
	// direct : 1
	// direct : 2
	// goroutine1 : 0
	// goroutine : 0
	// goroutine : 1
	// goroutine1 : 1
	// goroutine1 : 2
	// goroutine : 2

	// done
}

// 定义一个函数，循环3次输出:"from:index"
func f(from string) {
	for i := 0; i < 3; i++ {
		// 为了能看出多个协程执行效果，每循环一次，休眠1s
		time.Sleep(time.Second)
		fmt.Println(from, ":", i)
	}
}
