package zap_encoder

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// FilterZapTimeEncoder
// default ISO8601TimeEncoder
func FilterZapTimeEncoder(timeEncoder string) func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	switch timeEncoder {
	default:
		return zapcore.ISO8601TimeEncoder
	case "ISO8601TimeEncoder":
		return zapcore.ISO8601TimeEncoder
	case "EpochMillisTimeEncoder":
		return zapcore.EpochMillisTimeEncoder
	case "EpochNanosTimeEncoder":
		return zapcore.EpochNanosTimeEncoder
	case "EpochTimeEncoders":
		return zapcore.EpochTimeEncoder
	}
}

// FilterZapDurationEncoder
// default SecondsDurationEncoder
func FilterZapDurationEncoder(encodeDuration string) func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	switch encodeDuration {
	default:
		return zapcore.SecondsDurationEncoder
	case "SecondsDurationEncoder":
		return zapcore.SecondsDurationEncoder
	case "NanosDurationEncoder":
		return zapcore.NanosDurationEncoder
	case "StringDurationEncoder":
		return zapcore.StringDurationEncoder
	}
}

// FilterZapCallerEncoder
// default FullCallerEncoder
func FilterZapCallerEncoder(encodeCaller string) func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	switch encodeCaller {
	default:
		return zapcore.FullCallerEncoder
	case "FullCallerEncoder":
		return zapcore.FullCallerEncoder
	case "ShortCallerEncoder":
		return zapcore.ShortCallerEncoder
	}

}

// FilterZapEncodeLevel
// default CapitalLevelEncoder
func FilterZapEncodeLevel(level string) func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	default:
		return zapcore.CapitalLevelEncoder
	case "CapitalLevelEncoder":
		return zapcore.CapitalLevelEncoder
	case "CapitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder
	case "LowercaseLevelEncoder":
		return zapcore.LowercaseLevelEncoder
	case "LowercaseColorLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder
	}
}

// FilterZapEncoder
// default zapcore.NewConsoleEncoder(encoderConfig)
func FilterZapEncoder(zapEncoding string, encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	var encoder zapcore.Encoder
	switch zapEncoding {
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return encoder
}

// FilterZapAtomicLevelByViper
// default zap.InfoLevel
func FilterZapAtomicLevelByViper(atomicLevel int) zapcore.Level {
	var atomViper zapcore.Level
	switch atomicLevel {
	default:
		atomViper = zap.InfoLevel
	case -1:
		atomViper = zap.DebugLevel
	case 0:
		atomViper = zap.InfoLevel
	case 1:
		atomViper = zap.WarnLevel
	case 2:
		atomViper = zap.ErrorLevel
	}
	return atomViper
}

// NewEncoderConfigByViper
// new zapcore.EncoderConfig
func NewEncoderConfigByViper() *zapcore.EncoderConfig {
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
		EncodeLevel:    FilterZapEncodeLevel(viper.GetString("zap.EncoderConfig.EncodeLevel")),
		EncodeTime:     FilterZapTimeEncoder(viper.GetString("zap.EncoderConfig.TimeEncoder")), // ISO8601TimeEncoder ISO8601 UTC time
		EncodeDuration: FilterZapDurationEncoder(viper.GetString("zap.EncoderConfig.EncodeDuration")),
		EncodeCaller:   FilterZapCallerEncoder(viper.GetString("zap.EncoderConfig.EncodeCaller")),
	}
}
