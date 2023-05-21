package main

import "fmt"

func main() {
	// for range 来遍历通道
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
	}()

	for v := range c {
		fmt.Println(v)
	}
	// 输出
	// 1
	// 2
}
