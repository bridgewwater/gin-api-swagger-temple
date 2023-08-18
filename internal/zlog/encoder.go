package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// filterZapTimeEncoder
// default ISO8601TimeEncoder
func filterZapTimeEncoder(timeEncoder string) func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
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

// filterZapDurationEncoder
// default SecondsDurationEncoder
func filterZapDurationEncoder(encodeDuration string) func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
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

// filterZapCallerEncoder
// default FullCallerEncoder
func filterZapCallerEncoder(encodeCaller string) func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	switch encodeCaller {
	default:
		return zapcore.FullCallerEncoder
	case "FullCallerEncoder":
		return zapcore.FullCallerEncoder
	case "ShortCallerEncoder":
		return zapcore.ShortCallerEncoder
	}

}

// filterZapEncodeLevel
// default CapitalLevelEncoder
func filterZapEncodeLevel(level string) func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
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

// filterZapEncoder
// default zapcore.NewConsoleEncoder(encoderConfig)
func filterZapEncoder(zapEncoding string, encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
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

// filterZapAtomicLevelByViper
// default zap.InfoLevel
func filterZapAtomicLevelByViper(atomicLevel int) zapcore.Level {
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
