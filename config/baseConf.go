package config

import (
	"fmt"
	"git.sinlov.cn/bridgewwater/temp-gin-api-self/util/sys"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/url"
)

var baseConf BaseConf

type BaseConf struct {
	BaseURL   string
	SSLEnable bool
}

func BaseURL() string {
	return baseConf.BaseURL
}

// read default config by conf/config.yaml
// can change by CLI by `-c`
// this config can config by ENV
//	ENV_WEB_HTTPS_ENABLE=false
//	ENV_AUTO_HOST=true
//	ENV_WEB_HOST 127.0.0.1:8000
func initBaseConf() {
	ssLEnable := false
	if viper.GetString(defaultEnvHttpsEnable) == "true" {
		ssLEnable = true
	} else {
		ssLEnable = viper.GetBool("sslEnable")
	}
	runMode := viper.GetString("runmode")
	var apiBase string
	if "debug" == runMode {
		apiBase = viper.GetString("dev_url")
	} else if "test" == runMode {
		apiBase = viper.GetString("test_url")
	} else {
		apiBase = viper.GetString("prod_url")
	}

	uri, err := url.Parse(apiBase)
	if err != nil {
		panic(err)
	}
	log.Debugf("uri.Host %v", uri.Host)

	baseHOSTByEnv := viper.GetString(defaultEnvHost)
	if baseHOSTByEnv != "" {
		uri.Host = baseHOSTByEnv
		apiBase = uri.String()
	} else {
		isAutoHost := viper.GetBool(defaultEnvAutoGetHost)
		if isAutoHost {
			ipv4, err := sys.NetworkLocalIP()
			if err == nil {
				var proc string
				if ssLEnable {
					proc = "https"
				} else {
					proc = "http"
				}
				addrStr := viper.GetString("addr")
				apiBase = fmt.Sprintf("%v://%v%v", proc, ipv4, addrStr)
			}
		}
	}
	log.Debugf("apiBase %v", apiBase)
	baseConf = BaseConf{
		BaseURL:   apiBase,
		SSLEnable: ssLEnable,
	}
}
