package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	// env prefix is web
	defaultEnvPrefix string = "ENV_WEB"
	// EnvLogLevel
	//	env ENV_WEB_LOG_LEVEL default ""
	EnvLogLevel string = "LOG_LEVEL"
	// EnvHttpsEnable
	//	env ENV_WEB_HTTPS_ENABLE default false
	EnvHttpsEnable string = "HTTPS_ENABLE"
	// EnvHostName
	//	env ENV_WEB_HOSTNAME default ""
	EnvHostName string = "HOSTNAME"
	// EnvHostPort
	//	env ENV_WEB_HOST_PORT default ""
	EnvHostPort string = "HOST_PORT"
	// EnvAutoGetHost
	//	env ENV_AUTO_HOST default true, will use local ipv4
	EnvAutoGetHost string = "AUTO_HOST"
)

var mustConfigString = []string{
	"runmode",
	"port",
	"name",
	"api_base",
	"base_path",
	// project set
}

type Config struct {
	Name string
}

// Init
// read default config by conf/config.yaml
// can change by CLI by `-c`
// this config can config by ENV
// ENV_WEB_HTTPS_ENABLE=false
// ENV_AUTO_HOST=true
// ENV_WEB_HOST 127.0.0.1:8000
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

	// 设置 默认环境变量
	_ = os.Setenv(EnvHostName, "")
	_ = os.Setenv(EnvHostPort, "34565")
	_ = os.Setenv(EnvHttpsEnable, "false")
	_ = os.Setenv(EnvAutoGetHost, "false")

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

// checkMustHasString
// check config.yaml must have string key
//
// config.mustConfigString
func checkMustHasString() error {
	for _, config := range mustConfigString {
		if "" == viper.GetString(config) {
			return fmt.Errorf("not has must string key [ %v ]", config)
		}
	}
	return nil
}
