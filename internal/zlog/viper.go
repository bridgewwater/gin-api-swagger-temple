package zlog

import (
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zap_encoder"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zip_rotate"
	"github.com/bridgewwater/gin-api-swagger-temple/internal/zlog/zlog_access"
	"github.com/mattn/go-colorable"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// MockZapLoggerInit
// for unit test
func MockZapLoggerInit() {
	logger, _ := zap.NewProduction()
	newSLogAsZap(logger, logger.Sugar())
	zlog_access.MockInit()
}

// ZapLoggerInitByViper
// init zap logger
//
//	 viper config example:
//
//		# zap config
//		zap:
//		  AtomicLevel: -1 # DebugLevel:-1 InfoLevel:0 WarnLevel:1 ErrorLevel:2
//		  Api:
//		    PrefixPaths: "/api/v1/" # api path prefix list
//		    AtomicLevel: 0 # DebugLevel:-1 InfoLevel:0 WarnLevel:1 ErrorLevel:2 default 0
//		  FieldsAuto: false # is use auto Fields key set
//		  Fields:
//		    Key: key
//		    Val: val
//		  Development: true #  is open file and line number
//		  Encoding: console # output format, only use console or json, default is console
//		  rotate:
//		    Filename: logs/temp-gin-web.log # Log file path
//		    # AccessFilename: logs/access.log # Access log file path
//		    # ApiFilename: logs/api.log # api log file path
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

	atomicLevel := zap.NewAtomicLevelAt(zap_encoder.FilterZapAtomicLevelByViper(viper.GetInt("zap.AtomicLevel"))) // log Level

	encoderConfig := zap_encoder.NewEncoderConfigByViper()
	confEncoding := viper.GetString("zap.Encoding")
	if confEncoding == "" {
		confEncoding = "console"
	}
	encoder := zap_encoder.FilterZapEncoder(confEncoding, *encoderConfig)

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

	newSLogAsZap(logZap, logZap.Sugar())

	err := zlog_access.InitByViper()
	if err != nil {
		return err
	}

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
