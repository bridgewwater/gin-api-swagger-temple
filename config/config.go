package config

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"git.sinlov.cn/bridgewwater/temp-gin-api-self/util/sys"
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

const (
	// env prefix is web
	defaultEnvPrefix string = "ENV_WEB"
	// env ENV_WEB_HTTPS_ENABLE default false
	defaultEnvHttpsEnable string = "HTTPS_ENABLE"
	// env ENV_WEB_HOST default ""
	defaultEnvHost string = "HOST"
	// env ENV_AUTO_HOST default true
	defaultEnvAutoGetHost string = "AUTO_HOST"
)

var mustConfigString = []string{
	"runmode",
	"addr",
	"name",
	"base_path",
	// project set
}

type Config struct {
	Name string
}

var baseConf BaseConf

type BaseConf struct {
	BaseURL   string
	SSLEnable bool
}

// read default config by conf/config.yaml
// can change by CLI by `-c`
// this config can config by ENV
//	ENV_WEB_HTTPS_ENABLE=false

//	ENV_WEB_HOST 127.0.0.1:8000
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// initialize configuration file
	if err := c.initConfig(); err != nil {
		return err
	}

	// initialization log package
	if err := c.initLog(); err != nil {
		return err
	}

	// init BaseConf
	initBaseConf()

	// TODO other config

	// monitor configuration changes and hot loaders
	c.watchConfig()

	return nil
}

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

func BaseURL() string {
	return baseConf.BaseURL
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath(filepath.Join("conf")) // 如果没有指定配置文件，则解析默认的配置文件 conf/config.go
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")          // 设置配置文件格式为YAML
	viper.AutomaticEnv()                 // 读取匹配的环境变量
	viper.SetEnvPrefix(defaultEnvPrefix) // 读取环境变量的前缀为 defaultEnvPrefix

	// 设置默认环境变量
	_ = os.Setenv(defaultEnvHost, "")
	_ = os.Setenv(defaultEnvHttpsEnable, "false")
	_ = os.Setenv(defaultEnvAutoGetHost, "true")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	if err := checkMustHasString(); err != nil {
		return err
	}

	return nil
}

// check config.yaml must has string key
//	config.mustConfigString
func checkMustHasString() error {
	for _, config := range mustConfigString {
		if "" == viper.GetString(config) {
			return fmt.Errorf("not has must string key [ %v ]", config)
		}
	}
	return nil
}

// Monitor configuration changes and hot loaders
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}

// Initialization log
func (c *Config) initLog() error {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	err := log.InitWithConfig(&passLagerCfg)
	return err
}
