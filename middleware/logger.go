package middleware

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		fmt.Println("请求方法:", c.Request.Method)
		fmt.Println("请求URL:", c.Request.URL.String())
		fmt.Println("请求Header:", c.Request.Header)

		// 请求 Body
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)

			fmt.Println("请求Body:", string(bodyBytes))

			// 放回去，保证 Handler 还能继续读取
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 先包装 ResponseWriter
		writer := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}

		c.Writer = writer

		c.Next()

		// 再执行 Handler
		c.Next()

		// Handler 执行完后，此时 writer.body 已经有返回内容
		fmt.Println("响应状态码:", c.Writer.Status())
		fmt.Println("响应Body:", writer.body.String())
		fmt.Println("请求耗时:", time.Since(start))
	}
}
