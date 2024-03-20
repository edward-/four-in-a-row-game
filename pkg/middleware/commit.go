package middleware

import (
	"github.com/edward-/four-in-a-row-game/build"
	"github.com/gin-gonic/gin"
)

func CommitVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		version := build.Version
		hashCommint := build.HashCommit

		c.Writer.Header().Set(versionKey, version)
		c.Writer.Header().Set(buildKey, hashCommint)
		c.Next()
	}
}
