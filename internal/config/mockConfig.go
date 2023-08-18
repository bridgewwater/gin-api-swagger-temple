package config

import (
	"github.com/gin-gonic/gin"
	"os"
)

// MockInitSample
// for unit test
func MockInitSample(host string, baseUrl string, sslEnable bool) *gin.Engine {
	e := gin.Default()
	gin.SetMode(fetchGinRunMode())

	baseConf = BaseConf{
		Addr:      host,
		BaseURL:   baseUrl,
		SSLEnable: sslEnable,
	}

	return e
}

func fetchGinRunMode() string {
	ginMode := os.Getenv(gin.EnvGinMode)
	if ginMode == "" {
		ginMode = gin.TestMode
	}
	return ginMode
}
