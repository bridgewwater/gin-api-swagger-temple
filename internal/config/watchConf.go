package config

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Monitor configuration changes and hot loaders
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		zlog.S().Debugf("Config file changed: %s", e.Name)
	})
}
