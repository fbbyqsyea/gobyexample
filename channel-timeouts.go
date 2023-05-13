package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- "result 1"
	}()

	for i := 0; i < 2; i++ {
		select {
		case res := <-ch1:
			fmt.Println(res)
		case <-time.After(time.Second):
			fmt.Println("timeout 1")
		}
	}
}
