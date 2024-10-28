package zap

import (
	"samsamoohooh-go-api/internal/infra/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewCustomZapLogger(config *config.Config) (*zap.Logger, error) {

	// Define custom encoder configuration
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // Capitalize the log level names
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC timestamp format
		EncodeDuration: zapcore.SecondsDurationEncoder, // Duration in seconds
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Short caller (file and line)
	}

	// Create a lumberjack logger for file rotation
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.Logger.Filename,   // 로그 파일 경로
		MaxSize:    config.Logger.MaxSize,    // 최대 사이즈 (MB)
		MaxBackups: config.Logger.MaxBackups, // 최대 백업 파일 수
		MaxAge:     config.Logger.MaxAge,     // 최대 파일 유지 기간 (일)
		Compress:   config.Logger.Compress,   // 압축 여부
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		zapcore.Level(config.Logger.Level), // 로그 레벨
	)

	logger := zap.New(core)
	return logger, nil
}
