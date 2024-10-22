package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next()
	log.Printf(
		"%s %s %d %s in %v",
		c.Request.Method,
		c.Request.URL.Path,
		c.Writer.Status(),
		c.Request.RemoteAddr,
		time.Since(start),
	)
}
