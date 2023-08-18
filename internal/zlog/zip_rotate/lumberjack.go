package zip_rotate

import (
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLoggerByViper() *lumberjack.Logger {
	fileName := viper.GetString("zap.rotate.Filename")
	if fileName == "" {
		fileName = "logs/app.log"
	}
	maxSize := viper.GetInt("zap.rotate.MaxSize")
	if maxSize < 0 {
		maxSize = 0
	}
	maxBackUps := viper.GetInt("zap.rotate.MaxBackups")
	if maxBackUps < 0 {
		maxBackUps = 0
	}
	maxAge := viper.GetInt("zap.rotate.MaxAge")
	if maxAge < 0 {
		maxAge = 0
	}
	return &lumberjack.Logger{
		Filename:   fileName,                             // Log file path
		MaxSize:    maxSize,                              // Maximum size of each log file Unit: M
		MaxBackups: maxBackUps,                           // How many backups are saved in the log file
		MaxAge:     maxAge,                               // How many days can the file be keep
		Compress:   viper.GetBool("zap.rotate.Compress"), // need compress
	}
}
