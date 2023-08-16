package biz_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func MockRequestPost(t *testing.T, router *gin.Engine, url string, header map[string]string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	return MockRequest(t, router, http.MethodPost, url, header, body)
}

func MockRequestGet(t *testing.T, router *gin.Engine, url string, header map[string]string) (*httptest.ResponseRecorder, *http.Request) {
	return MockRequest(t, router, http.MethodGet, url, header, nil)
}

func MockRequest(t *testing.T, router *gin.Engine, method, url string, header map[string]string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	newRequest, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("mock request name %s method [ %s ] url %v error %v", t.Name(), method, url, err)
	}
	if len(header) > 0 {
		for k, v := range header {
			newRequest.Header.Add(k, v)
		}
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, newRequest)
	return recorder, newRequest
}

func FetchGinRunMode() string {
	ginMode := os.Getenv(gin.EnvGinMode)
	if ginMode == "" {
		ginMode = gin.TestMode
	}
	return ginMode
}

// MockQueryStrFrom
//
//	make query string from params
func MockQueryStrFrom(params interface{}) (result string) {
	if params == nil {
		return
	}
	value := reflect.ValueOf(params)

	switch value.Kind() {
	case reflect.Struct:
		var formName string
		for i := 0; i < value.NumField(); i++ {
			if formName = value.Type().Field(i).Tag.Get("form"); formName == "" {
				// don't tag the form name, use camel name
				formName = GetCamelNameFrom(value.Type().Field(i).Name)
			}
			result += "&" + formName + "=" + fmt.Sprintf("%v", value.Field(i).Interface())
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			result += "&" + fmt.Sprintf("%v", key.Interface()) + "=" + fmt.Sprintf("%v", value.MapIndex(key).Interface())
		}
	default:
		return
	}

	if result != "" {
		result = result[1:]
	}
	return
}

// GetCamelNameFrom
//
//	get the Camel name of the original name
func GetCamelNameFrom(name string) string {
	result := ""
	i := 0
	j := 0
	r := []rune(name)
	for m, v := range r {
		// if the char is the capital
		if v >= 'A' && v < 'a' {
			// if the prior is the lower-case || if the prior is the capital and the latter is the lower-case
			if (m != 0 && r[m-1] >= 'a') || ((m != 0 && r[m-1] >= 'A' && r[m-1] < 'a') && (m != len(r)-1 && r[m+1] >= 'a')) {
				i = j
				j = m
				result += name[i:j] + "_"
			}
		}
	}

	result += name[j:]
	return strings.ToLower(result)
}
