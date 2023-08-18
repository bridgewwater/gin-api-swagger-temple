package zlog_access

import (
	"fmt"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zap_encoder"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zip_rotate"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const defaultAccessFileName = "logs/access.log"

func InitByViper() error {

	atomicLevel := zap.NewAtomicLevelAt(zap_encoder.FilterZapAtomicLevelByViper(viper.GetInt("zap.AtomicLevel"))) // log Level

	encoderConfig := zap_encoder.NewEncoderConfigByViper()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	confEncoding := viper.GetString("zap.Encoding")
	if confEncoding == "" {
		return fmt.Errorf("config zap.Encoding is empty")
	}
	encoder := zap_encoder.FilterZapEncoder(confEncoding, *encoderConfig)

	rotateLogger := zip_rotate.NewLoggerByViper()
	accessFileName := viper.GetString("zap.rotate.AccessFilename")
	if accessFileName == "" {
		accessFileName = defaultAccessFileName
	}
	rotateLogger.Filename = accessFileName

	core := zapcore.NewCore(
		encoder, // Encoder configuration
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(rotateLogger),
		), // Print to console and file
		atomicLevel, // Log level
	)

	//var logZap *zap.Logger
	//if viper.GetBool("zap.Development") {
	//	logZap = zap.New(core, zap.AddCaller(), zap.Development())
	//} else {
	//	logZap = zap.New(core)
	//}
	logZap := zap.New(core)
	newAccessAsZap(logZap, logZap.Sugar())

	return nil
}
