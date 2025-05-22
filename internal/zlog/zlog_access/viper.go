package zlog_access

import (
	"errors"

	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/common"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zap_encoder"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zip_rotate"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultAccessFileName = "logs/access.log"
	defaultApiFileName    = "logs/api.log"
)

func InitByViper() error {
	atomicLevel := zap.NewAtomicLevelAt(
		zap_encoder.FilterZapAtomicLevelByViper(viper.GetInt("zap.AtomicLevel")),
	) // log Level

	encoderConfig := zap_encoder.NewEncoderConfigByViper()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	confEncoding := viper.GetString("zap.Encoding")
	if confEncoding == "" {
		return errors.New("config zap.Encoding is empty")
	}

	encoder := zap_encoder.FilterZapEncoder(confEncoding, *encoderConfig)

	// format skipPath
	skipPath = common.RemoveStringDuplicateNotCopy(skipPath)

	accessFileName := viper.GetString("zap.rotate.AccessFilename")
	if accessFileName == "" {
		accessFileName = defaultAccessFileName
	}

	apiFileName := viper.GetString("zap.rotate.ApiFilename")
	if apiFileName == "" {
		apiFileName = defaultApiFileName
	}

	if accessFileName == apiFileName {
		return errors.New(
			"config [ zap.rotate.AccessFilename ] and [ zap.rotate.ApiFilename ] is same",
		)
	}

	initAccessByViper(encoder, atomicLevel, accessFileName)
	initApiByViper(encoder, apiFileName)

	return nil
}

func initAccessByViper(encoder zapcore.Encoder, atomicLevel zap.AtomicLevel, fileName string) {
	rotateLogger := zip_rotate.NewLoggerByViper()
	rotateLogger.Filename = fileName

	core := zapcore.NewCore(
		encoder, // Encoder configuration
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(rotateLogger),
		), // Print to console and file
		atomicLevel, // Log level
	)

	logZap := zap.New(core)
	newAccessAsZap(logZap, logZap.Sugar())
}

func initApiByViper(encoder zapcore.Encoder, fileName string) {
	atomicLevel := zap.NewAtomicLevelAt(
		zap_encoder.FilterZapAtomicLevelByViper(viper.GetInt("zap.Api.AtomicLevel")),
	)
	rotateLogger := zip_rotate.NewLoggerByViper()
	rotateLogger.Filename = fileName

	configApiPaths := viper.GetStringSlice("zap.Api.PrefixPaths")
	if len(configApiPaths) > 0 {
		AppendApiPrefix(configApiPaths...)
	}

	core := zapcore.NewCore(
		encoder, // Encoder configuration
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(rotateLogger),
		), // Print to console and file
		atomicLevel, // Log level
	)

	logZap := zap.New(core)
	newApiAsZap(logZap, logZap.Sugar())
}

// MockInit
//
//	for unit test init
func MockInit() {
	// format skipPath
	skipPath = common.RemoveStringDuplicateNotCopy(skipPath)

	logger, _ := zap.NewProduction()
	newAccessAsZap(logger, logger.Sugar())
	newApiAsZap(logger, logger.Sugar())
}
