package logger

import (
	"samsamoohooh-go-api/internal/infra/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type AppLogger struct {
	*zap.Logger
}

func NewAppLogger(config *config.Config) (*AppLogger, error) {
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
		Filename:   config.Logger.App.Filename,   // 로그 파일 경로
		MaxSize:    config.Logger.App.MaxSize,    // 최대 사이즈 (MB)
		MaxBackups: config.Logger.App.MaxBackups, // 최대 백업 파일 수
		MaxAge:     config.Logger.App.MaxAge,     // 최대 파일 유지 기간 (일)
		Compress:   config.Logger.App.Compress,   // 압축 여부
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		zapcore.Level(config.Logger.App.Level), // 로그 레벨
	)

	logger := zap.New(core)
	return &AppLogger{Logger: logger}, nil
}
