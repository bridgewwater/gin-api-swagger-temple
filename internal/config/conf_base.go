package config

import (
	"fmt"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
	"strings"

	"github.com/bridgewwater/gin-api-swagger-temple/internal/sys"
	"github.com/spf13/viper"
)

var baseConf *BaseConf

type BaseConf struct {
	buildId string

	Addr      string
	BaseURL   string
	SSLEnable bool
}

// FetchBuildId
//
//	get build id
func FetchBuildId() string {
	return baseConf.buildId
}

// Addr
//
//	get Addr
func Addr() string {
	return baseConf.Addr
}

// BaseURL
//
//	get base url
func BaseURL() string {
	return baseConf.BaseURL
}

var _ginRunMode string

// GinRunMode
//
//	get gin run mode
func GinRunMode() string {
	if _ginRunMode == "" {
		ginRunMode := os.Getenv(gin.EnvGinMode)
		if ginRunMode != "" {
			_ginRunMode = ginRunMode
			zlog.S().Debugf("gin mode initBaseConf by yaml: runMode %v", _ginRunMode)
		} else {
			_ginRunMode = viper.GetString("runmode")
			zlog.S().Debugf("gin mode initBaseConf by env: %s=%s", gin.EnvGinMode, _ginRunMode)
		}
	}
	return _ginRunMode
}

// initBaseConf
//
//	read default config by conf/config.yaml
//	can change by CLI by `-c`
//	this config can config by ENV
//
//	ENV_WEB_HTTPS_ENABLE=false
//	ENV_AUTO_HOST=true
//	ENV_WEB_HOST_PORT 34565
//	ENV_WEB_HOSTNAME  0.0.0.0
//
// this function will check base config
func initBaseConf(buildId string) {
	gin.SetMode(GinRunMode())

	sslEnable := false
	if viper.GetBool(EnvHttpsEnable) {
		sslEnable = true
	} else {
		sslEnable = viper.GetBool("sslEnable")
	}

	apiBase := viper.GetString("api_base")

	apiBaseUrl, err := url.Parse(apiBase)
	if err != nil {
		panic(err)
	}

	//zlog.S().Debugf("api_base.Hostname %v", apiBaseUrl.Hostname())
	//zlog.S().Debugf("api_base.Port %v", apiBaseUrl.Port())

	runPort := viper.GetString("port")
	if viper.GetString(EnvHostPort) != "" {
		runPort = viper.GetString(EnvHostPort)
		zlog.S().Debugf("port change by env as: %s", runPort)
	}
	baseHostNameByEnv := viper.GetString(EnvHostName)

	if baseHostNameByEnv != "" {
		apiBaseUrl.Host = fmt.Sprintf("%s:%s", baseHostNameByEnv, runPort)
		apiBase = apiBaseUrl.String()
	} else {
		isAutoHost := viper.GetBool(EnvAutoGetHost)
		zlog.S().Debugf("isAutoHost %v", isAutoHost)
		if isAutoHost {
			ipv4, errLocalIp := sys.NetworkLocalIP()
			if errLocalIp == nil {
				var proc string
				if sslEnable {
					proc = "https"
				} else {
					proc = "http"
				}
				apiBase = fmt.Sprintf("%s://%s:%s", proc, ipv4, runPort)
				apiBaseUrl.Host = fmt.Sprintf("%s:%s", ipv4, runPort)
			}
		}
	}

	if sslEnable {
		apiBase = strings.Replace(apiBase, "http://", "https://", 1)
	}

	zlog.S().Debugf("run as apiBase: %s", apiBase)
	baseConf = &BaseConf{
		buildId: buildId,

		Addr:      apiBaseUrl.Host,
		BaseURL:   apiBase,
		SSLEnable: sslEnable,
	}
}
