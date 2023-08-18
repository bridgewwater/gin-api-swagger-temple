package middleware

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zlog_access"
	"github.com/gin-gonic/gin"
	"time"
)

// LoggerMiddleware
// just use logger to record
// ip will try X-Forwarded-For, X-Real-Ip
// filter at status code
// less than 400 use Warn
// less than 500 use Error
// other use Warning
// use as
//
//	g.Use(middleware.LoggerMiddleware())
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// request router
		reqUri := c.Request.RequestURI

		if zlog_access.CheckPathIsSkip(reqUri) {
			return
		}

		// start time
		startTime := time.Now()
		// to next
		c.Next()
		// end time
		endTime := time.Now()

		// latency time
		latencyTime := endTime.Sub(startTime)

		// request IP
		clientIP := c.ClientIP()

		// Method
		reqMethod := c.Request.Method

		// status code
		statusCode := c.Writer.Status()
		if statusCode < 400 {
			zlog_access.A().Infof(
				"=> %15s %13v | %s < %3d -> %s",
				clientIP,
				latencyTime,
				reqMethod,
				statusCode,
				reqUri)
		} else if statusCode < 500 {
			zlog_access.A().Warnf(
				"=> %15s %13v | %s < %3d -> %s",
				clientIP,
				latencyTime,
				reqMethod,
				statusCode,
				reqUri)
		} else {
			zlog_access.A().Errorf(
				"=> %15s %13v | %s < %3d -> %s",
				clientIP,
				latencyTime,
				reqMethod,
				statusCode,
				reqUri)
		}
	}
}

// LoggerToMongo logger to mongo
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// LoggerToMQ logger to MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
