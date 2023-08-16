package config

import (
	"github.com/bar-counter/slog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Monitor configuration changes and hot loaders
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		slog.Debugf("Config file changed: %s", e.Name)
	})
}