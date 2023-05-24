package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/stream", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		// 创建一个只写的响应流
		responseWriter := c.Writer
		flusher, _ := responseWriter.(http.Flusher)

		// 启动 goroutine 处理流式响应
		go func() {
			defer c.Done()

			// 模拟每秒发送一个事件
			for i := 0; i < 5; i++ {
				// 构造事件数据
				event := fmt.Sprintf("Event %d", i+1)

				// 写入事件数据到响应流
				fmt.Fprintf(responseWriter, "data: %s\n\n", event)

				// 刷新响应流
				flusher.Flush()

				// 模拟延迟 1 秒
				time.Sleep(time.Second)
			}
		}()

		// 等待上下文结束
		<-c.Done()
	})

	router.Run(":8080")
}
