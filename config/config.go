package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"path/filepath"
)

var mustConfigString = []string{
	"runmode",
	"addr",
	"name",
	"name",
	"base_path",
	// project set
}

type Config struct {
	Name string
}

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

	// monitor configuration changes and hot loaders
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath(filepath.Join("learn02", "conf")) // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
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
			return errors.New(fmt.Sprintf("not has must string key [ %v ]", config))
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
