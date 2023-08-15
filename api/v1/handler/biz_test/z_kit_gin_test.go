package biz_test

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
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
