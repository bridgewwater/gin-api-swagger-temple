package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sinlov-go/go-http-mock/gin_mock"
)

// MockInitSample
// for unit test
func MockInitSample(host string, baseUrl string, sslEnable bool) *gin.Engine {

	e := gin_mock.MockEngine(func(engine *gin.Engine) error {
		baseConf = &BaseConf{
			Addr:      host,
			BaseURL:   baseUrl,
			SSLEnable: sslEnable,
		}
		return nil
	})

	return e
}
