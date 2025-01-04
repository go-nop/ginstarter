package zap

import (
	"context"

	configurator "github.com/go-nop/ginstarter/pkg/config"
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.Logger
}

func NewZapLogger() *Logger {
	var logger *zap.Logger
	appEnv := configurator.GetEnv("APP_ENV", "development")
	if appEnv == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	return &Logger{log: logger}
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...map[string]any) {
	l.log.Debug(message, convertFields(fields)...)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...map[string]any) {
	l.log.Info(message, convertFields(fields)...)
}

func (l *Logger) Warn(ctx context.Context, message string, fields ...map[string]any) {
	l.log.Warn(message, convertFields(fields)...)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...map[string]any) {
	l.log.Error(message, convertFields(fields)...)
}

func (l *Logger) Fatal(ctx context.Context, message string, fields ...map[string]any) {
	l.log.Fatal(message, convertFields(fields)...)
}

func convertFields(fields []map[string]any) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, field := range fields {
		for key, value := range field {
			zapFields[i] = zap.Any(key, value)
		}
	}
	return zapFields
}
