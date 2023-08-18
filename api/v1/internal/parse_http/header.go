package parse_http

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/gin_kit/gin_head_check"
	"github.com/gin-gonic/gin"
	"time"
)

type HeaderContent struct {
	ParseTime  string            `json:"parse_time,omitempty"`
	URL        string            `json:"url"`
	Method     string            `json:"method,omitempty"`
	RemoteAddr string            `json:"remote_addr,omitempty"`
	Host       string            `json:"host,omitempty"`
	Proto      string            `json:"proto,omitempty"`
	ProtoMajor int               `json:"proto_major,omitempty"`
	ProtoMinor int               `json:"proto_minor,omitempty"`
	Param      string            `json:"param,omitempty"`
	Referer    string            `json:"referer,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Header     string            `json:"header,omitempty"`
}

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
