package ginzerolog

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)


//Logger function is gin middleware to use zerolog
func Logger() gin.HandlerFunc {
	return LoggerWithWriter(gin.DefaultWriter)
}

func LoggerWithWriter(out io.Writer) gin.HandlerFunc {
	log := zerolog.New(out).With().Timestamp().Logger()

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		latency := time.Since(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		event := log.Info()
		if comment != "" {
			event = log.Error()
		}

		event.
			Int("statusCode", statusCode).
			Dur("latency", latency).
			Str("clientIP", clientIP).
			Str("method", method).
			Str("path", path).
			Msg(comment)
	}
}
