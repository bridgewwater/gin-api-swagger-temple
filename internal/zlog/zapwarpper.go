package zlog

import (
	"fmt"
	"go.uber.org/zap"
)

// S
// for zap Sugar
func S() *zap.SugaredLogger {
	if zapLog == nil {
		panic(fmt.Errorf("please use zlog.ZapLoggerInitByViper() or zlog.MockZapLoggerInit() for unit test"))
	}
	return zapLog.Sugar
}

var zapLog *zapLogger

type zapLogger struct {
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
}

// newSLogAsZap
// init zap by sugar
func newSLogAsZap(log *zap.Logger, sugar *zap.SugaredLogger) {
	if zapLog == nil {
		zapLog = &zapLogger{
			Log:   log,
			Sugar: sugar,
		}
		sugar.Info("zap log init success as sugar")
	}
}
