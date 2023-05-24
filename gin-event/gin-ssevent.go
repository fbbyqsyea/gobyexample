package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/stream", func(c *gin.Context) {
		// 设置响应头，指定为事件流格式
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		// 启动事件流
		flusher, _ := c.Writer.(http.Flusher)
		if flusher == nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		eventStream := make(chan string)

		go func() {
			writer := bufio.NewWriter(c.Writer)

			for {
				event := <-eventStream
				fmt.Fprintf(writer, "data: %s\n\n", event)
				writer.Flush()
				flusher.Flush()
			}
		}()

		// 模拟产生事件并写入事件流
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			event := "Event " + strconv.Itoa(i+1)
			eventStream <- event
		}
	})

	router.Run(":8080")
}
