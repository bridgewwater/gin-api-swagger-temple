package zlog_access

import (
	"fmt"
	"go.uber.org/zap"
)

// I
//
//	for api zap Sugar
func I() *zap.SugaredLogger {
	if zapApiLog == nil {
		panic(fmt.Errorf("please use zlog_access.InitByViper()"))
	}
	return zapApiLog.Sugar
}

var zapApiLog *zapApiLogger

type zapApiLogger struct {
	Log   *zap.Logger
	Sugar *zap.SugaredLogger
}

// newApiAsZap
//
//	init zap api by sugar
func newApiAsZap(log *zap.Logger, sugar *zap.SugaredLogger) {
	if zapApiLog == nil {
		zapApiLog = &zapApiLogger{
			Log:   log,
			Sugar: sugar,
		}
		sugar.Info("zap log init api logger as sugar now")
	}
}
