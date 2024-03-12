package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetCorrelationUUIDMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		CorrelationID := uuid.NewString()

		if c.GetHeader(headerXRequestID) != "" {
			CorrelationID = c.GetHeader(headerXRequestID)
		}
		ctxCloned := c.Request.Clone(c)
		ctxCloned.Header.Set(correlationIDKey, CorrelationID)

		c.Next()
	})
}
