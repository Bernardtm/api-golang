package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func CacheMiddleware(maxAge int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age="+strconv.Itoa(maxAge))
		c.Next()
	}
}
