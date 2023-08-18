package zlog

import (
	"fmt"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zip_rotate"
	"github.com/mattn/go-colorable"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// ZapLoggerInitByViper
// init zap logger
//
//	 viper config example:
//
//		# zap config
//		zap:
//		  AtomicLevel: -1 # DebugLevel:-1 InfoLevel:0 WarnLevel:1 ErrorLevel:2
//		  FieldsAuto: false # is use auto Fields key set
//		  Fields:
//		    Key: key
//		    Val: val
//		  Development: true #  is open file and line number
//		  Encoding: console # output format, only use console or json, default is console
//		  rotate:
//		    Filename: logs/temp-gin-web.log # Log file path
//		    MaxSize: 16 # Maximum size of each zlog file, Unit: M
//		    MaxBackups: 10 # How many backups are saved in the zlog file
//		    MaxAge: 7 # How many days can the file be keep, Unit: day
//		    Compress: true # need compress
//		  EncoderConfig:
//		    TimeKey: time
//		    LevelKey: level
//		    NameKey: logger
//		    CallerKey: caller
//		    MessageKey: msg
//		    StacktraceKey: stacktrace
//		    TimeEncoder: ISO8601TimeEncoder # ISO8601TimeEncoder EpochMillisTimeEncoder EpochNanosTimeEncoder EpochTimeEncoder default is ISO8601TimeEncoder
//		    EncodeDuration: SecondsDurationEncoder # NanosDurationEncoder SecondsDurationEncoder StringDurationEncoder default is SecondsDurationEncoder
//		    EncodeLevel: CapitalColorLevelEncoder # CapitalLevelEncoder CapitalColorLevelEncoder LowercaseColorLevelEncoder LowercaseLevelEncoder default is CapitalLevelEncoder
//		    EncodeCaller: ShortCallerEncoder # ShortCallerEncoder FullCallerEncoder default is FullCallerEncoder
//
// return error
func ZapLoggerInitByViper() error {

	atomicLevel := zap.NewAtomicLevelAt(filterZapAtomicLevelByViper(viper.GetInt("zap.AtomicLevel"))) // log Level

	encoderConfig := newEncoderConfigByViper()
	confEncoding := viper.GetString("zap.Encoding")
	if confEncoding == "" {
		return fmt.Errorf("config zap.Encoding is empty")
	}
	encoder := filterZapEncoder(confEncoding, *encoderConfig)

	rotateLogger := zip_rotate.NewLoggerByViper()

	var core zapcore.Core
	if isAddColorableOut() {
		core = zapcore.NewCore(
			encoder, // Encoder configuration
			zapcore.NewMultiWriteSyncer(
				zapcore.AddSync(colorable.NewColorableStdout()),
				zapcore.AddSync(rotateLogger),
			), // Print to console and file
			atomicLevel, // Log level
		)
	} else {
		core = zapcore.NewCore(
			encoder, // Encoder configuration
			zapcore.NewMultiWriteSyncer(
				zapcore.AddSync(os.Stdout),
				zapcore.AddSync(rotateLogger),
			), // Print to console and file
			atomicLevel, // Log level
		)
	}

	var filed zap.Option
	if viper.GetBool("zap.FieldsAuto") {
		filed = zap.Fields( //the initialization field
			zap.String(viper.GetString("zap.Fields.Key"), viper.GetString("zap.Fields.Val")),
		)
	}

	var logZap *zap.Logger
	if viper.GetBool("zap.Development") {
		if filed != nil {
			logZap = zap.New(core, zap.AddCaller(), zap.Development(), filed)
		} else {
			logZap = zap.New(core, zap.AddCaller(), zap.Development())
		}
	} else {
		if filed != nil {
			logZap = zap.New(core, filed)
		} else {
			logZap = zap.New(core)
		}
	}

	newZapLog(logZap, logZap.Sugar())

	return nil
}

func isAddColorableOut() bool {
	zapEncodeLevel := viper.GetString("zap.EncoderConfig.EncodeLevel")
	switch zapEncodeLevel {
	default:
		return false
	case "CapitalColorLevelEncoder":
		return true
	case "LowercaseColorLevelEncoder":
		return true
	}
}

func newEncoderConfigByViper() *zapcore.EncoderConfig {
	timeKey := viper.GetString("zap.EncoderConfig.TimeKey")
	levelKey := viper.GetString("zap.EncoderConfig.LevelKey")
	nameKey := viper.GetString("zap.EncoderConfig.NameKey")
	callerKey := viper.GetString("zap.EncoderConfig.CallerKey")
	messageKey := viper.GetString("zap.EncoderConfig.MessageKey")
	stacktraceKey := viper.GetString("zap.EncoderConfig.StacktraceKey")
	return &zapcore.EncoderConfig{
		TimeKey:        timeKey,
		LevelKey:       levelKey,
		NameKey:        nameKey,
		CallerKey:      callerKey,
		MessageKey:     messageKey,
		StacktraceKey:  stacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    filterZapEncodeLevel(viper.GetString("zap.EncoderConfig.EncodeLevel")),
		EncodeTime:     filterZapTimeEncoder(viper.GetString("zap.EncoderConfig.TimeEncoder")), // ISO8601TimeEncoder ISO8601 UTC time
		EncodeDuration: filterZapDurationEncoder(viper.GetString("zap.EncoderConfig.EncodeDuration")),
		EncodeCaller:   filterZapCallerEncoder(viper.GetString("zap.EncoderConfig.EncodeCaller")),
	}
}
