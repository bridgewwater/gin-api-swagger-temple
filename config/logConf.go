package config

import (
	"github.com/bar-counter/slog"
	"github.com/spf13/viper"
)

// Initialization log
func (c *Config) initLog() error {
	passLagerCfg := slog.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rolling_policy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	err := slog.InitWithConfig(&passLagerCfg)
	return err
}
