package middleware

import (
	"bytes"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_id_snowflake"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zlog_access"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// LoggerContentMiddleware
//
//	logger content middleware
//	this middleware will log request and response content, and copy request body to gin.Context
func LoggerContentMiddleware() gin.HandlerFunc {
	return contentLogger
}

// responseWriter
//
//	custom type gin.ResponseWriter interface
type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	// write another copy of data to a bytes buffer
	w.b.Write(b)
	// Complete the original function of gin.Context.Writer.Write()
	return w.ResponseWriter.Write(b)
}

func contentLogger(c *gin.Context) {
	writer := responseWriter{
		c.Writer,
		bytes.NewBuffer([]byte{}),
	}
	c.Writer = writer

	requestBodyStr := ""
	if c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPost {
		requestData, err := io.ReadAll(c.Request.Body)
		if err == nil {
			requestBodyStr = string(requestData)
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestData))
	}

	c.Next()

	reqUri := c.Request.RequestURI

	if zlog_access.CheckPathIsSkip(reqUri) {
		return
	}

	if zlog_access.CheckPrefixIsSkip(reqUri) {
		return
	}

	if zlog_access.CheckPrefixIsApi(reqUri) {
		// request IP
		clientIP := c.ClientIP()

		// Method
		reqMethod := c.Request.Method

		// request id
		reqId := gin_correlation_id_snowflake.GetCorrelationID(c)

		// status code
		statusCode := c.Writer.Status()
		if statusCode < 400 {
			if writer.b != nil {
				responseStr := writer.b.String()
				if len(responseStr) > 0 {
					if len(requestBodyStr) > 0 {
						zlog_access.I().Infof(
							"=> %15s | %s < %3d rid:%s -> %s\n-> req : %s \n-> resp: %s",
							clientIP,
							reqMethod,
							statusCode,
							reqId,
							reqUri,
							requestBodyStr,
							responseStr,
						)
					} else {
						zlog_access.I().Infof(
							"=> %15s | %s < %3d rid:%s -> %s , request no body\n-> resp: %s",
							clientIP,
							reqMethod,
							statusCode,
							reqId,
							reqUri,
							responseStr,
						)
					}

				} else {
					if len(requestBodyStr) > 0 {
						zlog_access.I().Infof(
							"=> %15s | %s < %3d rid:%s -> %s\n-> req : %s",
							clientIP,
							reqMethod,
							statusCode,
							reqId,
							reqUri,
							requestBodyStr,
						)
					} else {
						zlog_access.I().Infof(
							"=> %15s | %s < %3d rid:%s -> %s , request no body",
							clientIP,
							reqMethod,
							statusCode,
							reqId,
							reqUri,
						)

					}
				}
			}
		} else if statusCode < 500 {
			responseStr := writer.b.String()
			if len(responseStr) > 0 {
				if len(requestBodyStr) > 0 {
					zlog_access.I().Warnf(
						"=> %15s | %s < %3d rid:%s -> %s\n-> req : %s \n-> resp: %s",
						clientIP,
						reqMethod,
						statusCode,
						reqId,
						reqUri,
						requestBodyStr,
						responseStr,
					)
				} else {
					zlog_access.I().Warnf(
						"=> %15s | %s < %3d rid:%s -> %s , request no body\n-> resp: %s",
						clientIP,
						reqMethod,
						statusCode,
						reqId,
						reqUri,
						responseStr,
					)
				}

			} else {
				if len(requestBodyStr) > 0 {
					zlog_access.I().Warnf(
						"=> %15s | %s < %3d rid:%s -> %s\n-> req : %s",
						clientIP,
						reqMethod,
						statusCode,
						reqId,
						reqUri,
						requestBodyStr,
					)
				} else {
					zlog_access.I().Warnf(
						"=> %15s | %s < %3d rid:%s -> %s , request no body",
						clientIP,
						reqMethod,
						statusCode,
						reqId,
						reqUri,
					)

				}
			}
		} else {
			if len(requestBodyStr) > 0 {
				zlog_access.I().Errorf(
					"=> %15s | %s < %3d rid:%s -> %s\n-> req : %s",
					clientIP,
					reqMethod,
					statusCode,
					reqId,
					reqUri,
					requestBodyStr,
				)
			} else {
				zlog_access.I().Errorf(
					"=> %15s | %s < %3d rid:%s -> %s , request no body",
					clientIP,
					reqMethod,
					statusCode,
					reqId,
					reqUri,
				)

			}
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
