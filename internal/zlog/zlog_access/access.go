package zlog_access

import (
	"fmt"
	"go.uber.org/zap"
)

// A
//
//	for access zap Sugar
func A() *zap.SugaredLogger {
	if zapAccessLog == nil {
		panic(fmt.Errorf("please use zlog_access.InitByViper()"))
	}
	return zapAccessLog.Sugar
}

var zapAccessLog *zapAccessLogger

type zapAccessLogger struct {
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
}

// newAccessAsZap
//
//	init zap access by sugar
func newAccessAsZap(log *zap.Logger, sugar *zap.SugaredLogger) {
	if zapAccessLog == nil {
		zapAccessLog = &zapAccessLogger{
			Log:   log,
			Sugar: sugar,
		}
		sugar.Info("zap log init access logger as sugar now")
	}
}
