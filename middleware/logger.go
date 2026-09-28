package middleware

import (
	"bytes"
	"io"
	"log/slog"
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

		// 读取请求 Body
		var requestBody string

		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = string(bodyBytes)

			// 放回去，保证 Handler 后面还能继续读取
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 包装 ResponseWriter，用来记录响应 Body
		writer := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}

		c.Writer = writer

		// 执行真正的 Handler
		c.Next()

		// Handler 执行完成后记录日志
		slog.Info(
			"http request",
			"method", c.Request.Method,
			"url", c.Request.URL.String(),
			"requestBody", requestBody,
			"status", c.Writer.Status(),
			"responseBody", writer.body.String(),
			"cost", time.Since(start),
		)
	}
}
