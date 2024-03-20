package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

const timeOutRequest time.Duration = 800 * time.Millisecond

func defaultResponse(c *gin.Context) {
	c.String(http.StatusRequestTimeout, "timeout")
}

func WithTimeout() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(timeOutRequest),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(defaultResponse),
	)
}
