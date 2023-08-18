package parse_http

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/gin_kit/gin_head_check"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type FormContent struct {
	URL          string            `json:"url"`
	ParseTime    string            `json:"parse_time,omitempty"`
	Method       string            `json:"method,omitempty"`
	RemoteAddr   string            `json:"remote_addr,omitempty"`
	Host         string            `json:"host,omitempty"`
	Proto        string            `json:"proto,omitempty"`
	ProtoMajor   int               `json:"proto_major,omitempty"`
	ProtoMinor   int               `json:"proto_minor,omitempty"`
	Param        string            `json:"param,omitempty"`
	Referer      string            `json:"referer,omitempty"`
	Headers      map[string]string `json:"headers,omitempty"`
	PostFormData map[string]string `json:"postData,omitempty"`
	QueryStrings map[string]string `json:"queryString,omitempty"`
}

func FormPost(c *gin.Context) (FormContent, error) {
	return FormPostParam(c, "")
}

// FormPostParam
// will remove HEADER Authorization
func FormPostParam(c *gin.Context, paramKey string) (FormContent, error) {
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
		return FormContent{}, err
	}
	postFormData := make(map[string]string)
	for k, v := range r.PostForm {
		postFormData[k] = strings.Join(v, "")
	}

	return FormContent{
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
		PostFormData: postFormData,
	}, nil
}
