package middleware

import "github.com/gin-gonic/gin"

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement jwk validation
		// token := c.Request.FormValue("token")
		c.Next()
	}
}
