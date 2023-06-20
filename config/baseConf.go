package config

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/gin-gonic/gin"
	"net/url"
	"os"
	"strings"

	"github.com/bridgewwater/gin-api-swagger-temple/util/sys"
	"github.com/spf13/viper"
)

var baseConf BaseConf

type BaseConf struct {
	BaseURL   string
	SSLEnable bool
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
			slog.Debugf("gin mode initBaseConf by yaml: runMode %v", _ginRunMode)
		} else {
			_ginRunMode = viper.GetString("runmode")
			slog.Debugf("gin mode initBaseConf by env: %s=%s", gin.EnvGinMode, _ginRunMode)
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
//	ENV_WEB_HOST 127.0.0.1:8000
func initBaseConf() {
	gin.SetMode(GinRunMode())

	ssLEnable := false
	if viper.GetBool(EnvHttpsEnable) {
		ssLEnable = true
	} else {
		ssLEnable = viper.GetBool("sslEnable")
	}

	apiBase := viper.GetString("api_base")

	uri, err := url.Parse(apiBase)
	if err != nil {
		panic(err)
	}

	slog.Debugf("uri.Host %v", uri.Host)
	baseHOSTByEnv := viper.GetString(EnvHost)
	if baseHOSTByEnv != "" {
		uri.Host = baseHOSTByEnv
		apiBase = uri.String()
	} else {
		isAutoHost := viper.GetBool(EnvAutoGetHost)
		slog.Debugf("isAutoHost %v", isAutoHost)
		if isAutoHost {
			ipv4, err := sys.NetworkLocalIP()
			if err == nil {
				addrStr := viper.GetString("addr")
				var proc string
				if ssLEnable {
					proc = "https"
				} else {
					proc = "http"
				}
				apiBase = fmt.Sprintf("%v://%v%v", proc, ipv4, addrStr)
			}
		}
	}

	if ssLEnable {
		apiBase = strings.Replace(apiBase, "http://", "https://", 1)
	}

	slog.Debugf("apiBase %v", apiBase)
	baseConf = BaseConf{
		BaseURL:   apiBase,
		SSLEnable: ssLEnable,
	}
}
