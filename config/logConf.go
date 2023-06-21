package config

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/spf13/viper"
)

// Initialization log
func (c *Config) initLog() error {

	envLogLevel := viper.GetString(EnvLogLevel)
	if envLogLevel == "" {
		envLogLevel = viper.GetString("log.logger_level")
		fmt.Printf("-> app log level initLog by yaml: log.logger_level %s\n", envLogLevel)
	} else {
		fmt.Printf("-> app log level initLog by env: %s_%s=%s\n", defaultEnvPrefix, EnvLogLevel, envLogLevel)
	}

	passLagerCfg := slog.PassLagerCfg{
		LoggerLevel:    envLogLevel,
		Writers:        viper.GetString("log.writers"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		LogHideLineno:  viper.GetBool("log.log_hide_lineno"),
		RollingPolicy:  viper.GetString("log.rolling_policy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	err := slog.InitWithConfig(&passLagerCfg)
	return err
}
