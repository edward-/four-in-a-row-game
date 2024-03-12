package middleware

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func CommitVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := ioutil.ReadFile("REVISION")
		// If we cant read file, just skip to the next request handler
		// This is pretty much a NOOP middlware :)
		if err != nil {
			// Make sure to log error so it could be spotted
			log.Println("revision middleware error:", err)

			return func(c *gin.Context) {
				c.Next()
			}
		}

		// Clean up the value since it could contain line breaks
		revision := strings.TrimSpace(string(data))

		// Set out header value for each response
		return func(c *gin.Context) {
			c.Writer.Header().Set("X-Revision", revision)
			c.Next()
		}

		c.Next()
	}
}
