package filestore

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger interface همان است
type Logger interface {
	Info(ctx context.Context, msg string, fields map[string]interface{})
	Error(ctx context.Context, msg string, fields map[string]interface{})
}

// ZapLogger حرفه‌ای با zap
type ZapLogger struct {
	logger *zap.Logger
}

// NewZapLogger یک logger حرفه‌ای با JSON output می‌سازد
func NewZapLogger() *ZapLogger {
	// Config برای production-ready logger
	cfg := zap.Config{
		Encoding:         "json",                  // JSON structured logs
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return &ZapLogger{logger: logger}
}

// Info با context و فیلدهای دلخواه
func (l *ZapLogger) Info(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logger.With(convertFields(fields)...).Info(msg)
}

// Error با context و فیلدهای دلخواه
func (l *ZapLogger) Error(ctx context.Context, msg string, fields map[string]interface{}) {
	l.logger.With(convertFields(fields)...).Error(msg)
}

// تبدیل map[string]interface{} به zap.Field
func convertFields(fields map[string]interface{}) []zap.Field {
	zFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zFields = append(zFields, zap.Any(k, v))
	}
	return zFields
}