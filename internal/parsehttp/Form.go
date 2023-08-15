package parsehttp

import (
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type Form struct {
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

func FormPost(c *gin.Context) (Form, error) {
	return FormPostParam(c, "")
}

func FormPostParam(c *gin.Context, paramKey string) (Form, error) {
	var paramVal string
	if paramKey != "" {
		paramVal = c.Param(paramKey)
	}
	requestHeader := make(map[string]string)
	r := c.Request

	// remove HEADER Authorization
	for k, v := range r.Header {
		if k != "Authorization" {
			requestHeader[k] = strings.Join(v, "")
		}
	}

	queryStrings := make(map[string]string)
	for k, v := range r.URL.Query() {
		queryStrings[k] = strings.Join(v, "")
	}
	err := r.ParseForm()
	if err != nil {
		return Form{}, err
	}
	postFormData := make(map[string]string)
	for k, v := range r.PostForm {
		postFormData[k] = strings.Join(v, "")
	}

	return Form{
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
