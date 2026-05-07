package middlewear

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger TODO: understand cNext and return function
func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()
		latency := time.Since(start)

		requestID, _ := c.Get(RequestIDKey)

		log.Info("request completed",
			"request_id", requestID,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency_ms", latency.Milliseconds(),
			"client_ip", c.ClientIP(),
		)
	}
}
