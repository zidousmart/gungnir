package middleware

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zidousmart/gungnir/internal/log"
)

func SetAccessLogger(logger *log.GLogger) {
	ngxLogger = logger
}

// 访问日志中间件
func MiddleAccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 响应数据接口
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 打印请求日志
		ngxLogInfo(c, "", 0)

		c.Next()

		duration := time.Since(start).Milliseconds()

		// 打印响应日志
		response := blw.body.String()
		if c.Writer.Status() == http.StatusOK {
			ngxLogInfo(c, response, duration)
		} else {
			ngxLogError(c, response, duration)
		}
	}
}

var ngxLogger *log.GLogger

// ngx log
func ngxLogInfo(c *gin.Context, response string, duration int64) {
	if ngxLogger == nil {
		return
	}

	ngxLogger.NgxInfo(c, response, duration)
}

func ngxLogError(c *gin.Context, response string, duration int64) {
	if ngxLogger == nil {
		return
	}

	ngxLogger.NgxError(c, response, duration)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
