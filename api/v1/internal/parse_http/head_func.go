package parse_http

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/gin_kit/gin_head_check"
	"github.com/gin-gonic/gin"
	"time"
)

func HttpHeader(c *gin.Context) (HeaderContent, error) {
	return HttpHeaderItem(c, "")
}

func HttpHeaderItem(c *gin.Context, headKey string) (HeaderContent, error) {
	var headerVal string
	if headKey != "" {
		headerVal = c.GetHeader(headKey)
	}

	requestHeader := gin_head_check.HeaderCheckAsMap(c)

	r := c.Request
	return HeaderContent{
		ParseTime:  time.Now().String(),
		URL:        r.URL.String(),
		Method:     r.Method,
		RemoteAddr: r.RemoteAddr,
		Host:       r.Host,
		Proto:      r.Proto,
		ProtoMajor: r.ProtoMajor,
		ProtoMinor: r.ProtoMinor,
		Referer:    r.Referer(),
		Headers:    requestHeader,
		Header:     headerVal,
	}, nil
}
