package core

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestB(t *testing.T) {
	zap.L().Info("hello world", zap.String("name", "zhangsan"))
}

func TestA(t *testing.T) {
	//development, _ := zap.NewDevelopment()
	production, _ := zap.NewProduction()
	//development.Info("hello world", zap.String("name", "zhangsan"))
	//production.Info("hello world", zap.String("name", "zhangsan"))

	conf := zap.Config{
		Level:         zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:   false,
		Encoding:      "json",
		OutputPaths:   []string{"stdout", "../log/app.log"},
		EncoderConfig: getE(),
	}
	production, _ = conf.Build()
	production.Info("hello world", zap.String("name", "zhangsan"))
}

func getE() zapcore.EncoderConfig {
	config := zap.NewDevelopmentEncoderConfig()
	config.TimeKey = "timestamp"
	config.LevelKey = "level"
	config.NameKey = "logger"
	config.CallerKey = "caller"
	config.MessageKey = "message"
	config.StacktraceKey = "stacktrace"

	config.LineEnding = zapcore.DefaultLineEnding

	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	//config.EncodeDuration = zapcore.SecondsDurationEncoder
	config.EncodeDuration = zapcore.StringDurationEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	return config
}

func GetEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,      // 默认换行符"\n"
			EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 日志等级序列为小写字符串，如:InfoLevel被序列化为 "info"
			EncodeTime:     zapcore.EpochTimeEncoder,       // 日志时间格式显示
			EncodeDuration: zapcore.SecondsDurationEncoder, // 时间序列化，Duration为经过的浮点秒数
			EncodeCaller:   zapcore.ShortCallerEncoder,     // 日志行号显示
		})
}
