package logger

import (
	"samsamoohooh-go-api/internal/infra/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var instance *zap.Logger

func Initialize(cfg *config.Config) error {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 환경에 따라 로그 레벨 설정
	var logLevel zapcore.Level
	if cfg.HTTP.Development {
		logLevel = zapcore.DebugLevel
	} else {
		logLevel = zapcore.InfoLevel
	}

	atomicLevel := zap.NewAtomicLevelAt(logLevel)
	zapConfig := zap.Config{
		Level:            atomicLevel,
		Development:      cfg.HTTP.Development,
		Encoding:         cfg.Logger.Encoding,
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return err
	}

	instance = logger
	return nil
}

func Sync() error {
	return instance.Sync()
}
func Get() *zap.Logger {
	return instance
}
