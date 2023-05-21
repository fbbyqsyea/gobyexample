package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/long_async_operation", func(c *gin.Context) {
		// 启动goroutine
		go func() {
			for i := 0; i < 10; i++ {
				// 进行耗时操作
				time.Sleep(1 * time.Second)
				// 结束时打印日志
				log.Println("long_async_operation done")
			}
		}()
	})

	r.Run(":8081")
}
