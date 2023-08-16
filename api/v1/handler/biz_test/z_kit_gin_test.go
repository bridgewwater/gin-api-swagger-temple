package biz_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	JSON = "json"
	FORM = "form"
)

var (
	ErrMethodNotSupported = errors.New("method is not supported")
	ErrMIMENotSupported   = errors.New("mime is not supported")
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
func MockJsonPost(t *testing.T, router *gin.Engine, url string, header map[string]string, param interface{}) (*httptest.ResponseRecorder, *http.Request) {
	newRequest, err := makeRequest(http.MethodPost, JSON, url, param)
	if err != nil {
		t.Fatalf("mock makeRequest name %s method [ %s ] url %v error %v", t.Name(), http.MethodPost, url, err)
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

func MockFormPost(t *testing.T, router *gin.Engine, url string, header map[string]string, param interface{}) (*httptest.ResponseRecorder, *http.Request) {
	newRequest, err := makeRequest(http.MethodPost, FORM, url, param)
	if err != nil {
		t.Fatalf("mock makeRequest name %s method [ %s ] url %v error %v", t.Name(), http.MethodPost, url, err)
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

// make request
func makeRequest(method, mime, api string, param interface{}) (request *http.Request, err error) {
	method = strings.ToUpper(method)
	mime = strings.ToLower(mime)

	switch mime {
	case JSON:
		var (
			contentBuffer *bytes.Buffer
			jsonBytes     []byte
		)
		jsonBytes, err = json.Marshal(param)
		if err != nil {
			return
		}
		contentBuffer = bytes.NewBuffer(jsonBytes)
		request, err = http.NewRequest(string(method), api, contentBuffer)
		if err != nil {
			return
		}
		request.Header.Set("Content-Type", "application/json;charset=utf-8")
	case FORM:
		queryStr := MockQueryStrFrom(param)
		var buffer io.Reader

		if (method == http.MethodDelete || method == http.MethodGet) && queryStr != "" {
			api += "?" + queryStr
		} else {
			buffer = bytes.NewReader([]byte(queryStr))
		}

		request, err = http.NewRequest(string(method), api, buffer)
		if err != nil {
			return
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	default:
		err = ErrMIMENotSupported
		return
	}
	return
}

// make request which contains uploading file
func MockFileRequest(method, api, fileName, fieldName string, param interface{}) (request *http.Request, err error) {
	method = strings.ToUpper(method)
	if method != http.MethodPost && method != http.MethodPut {
		err = ErrMethodNotSupported
		return
	}

	// create form file
	buf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buf)
	fileWriter, err := bodyWriter.CreateFormFile(fieldName, fileName)
	if err != nil {
		return
	}

	// read the file
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	// read the file to the fileWriter
	length, err := fileWriter.Write(fileBytes)
	if err != nil {
		return
	}

	bodyWriter.Close()

	// make request
	queryStr := MockQueryStrFrom(param)
	if queryStr != "" {
		api += "?" + queryStr
	}
	request, err = http.NewRequest(string(method), api, buf)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	err = request.ParseMultipartForm(int64(length))
	return
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
