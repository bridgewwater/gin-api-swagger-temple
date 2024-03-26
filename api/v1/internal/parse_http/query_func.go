package parse_http

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/gin_kit/gin_head_check"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func HttpQuery(c *gin.Context) (QueryContent, error) {
	return HttpQueryParam(c, "")
}

func HttpQueryParam(c *gin.Context, paramKey string) (QueryContent, error) {
	var paramVal string
	if paramKey != "" {
		paramVal = c.Param(paramKey)
	}
	requestHeader := gin_head_check.HeaderCheckAsMap(c)

	r := c.Request

	queryStrings := make(map[string]string)
	for k, v := range r.URL.Query() {
		queryStrings[k] = strings.Join(v, "")
	}
	err := r.ParseForm()
	if err != nil {
		return QueryContent{}, err
	}

	return QueryContent{
		ParseTime:    time.Now().String(),
		URL:          r.URL.String(),
		Method:       r.Method,
		RemoteAddr:   r.RemoteAddr,
		Host:         r.Host,
		Proto:        r.Proto,
		ProtoMajor:   r.ProtoMajor,
		ProtoMinor:   r.ProtoMinor,
		Param:        paramVal,
		Referer:      r.Referer(),
		Headers:      requestHeader,
		QueryStrings: queryStrings,
	}, nil
}
