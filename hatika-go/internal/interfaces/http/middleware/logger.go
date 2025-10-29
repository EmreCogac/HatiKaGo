package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs HTTP requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		
		// Process request
		c.Next()
		
		// Calculate latency
		latency := time.Since(startTime)
		
		// Get status code
		statusCode := c.Writer.Status()
		
		// Log format: [timestamp] method path status latency
		log.Printf("[%s] %s %s %d %v",
			startTime.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			statusCode,
			latency,
		)
	}
}
