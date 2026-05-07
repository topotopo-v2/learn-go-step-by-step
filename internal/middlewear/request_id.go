package middlewear

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.NewString()

		c.Set(RequestIDKey, requestId)
		c.Writer.Header().Set(RequestIDKey, requestId)

		c.Next()
	}
}
