package zap

import (
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

func (l *Logger) Debug(message string, fields ...map[string]any) {
	l.log.Debug(message, convertFields(fields)...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.log.Sugar().Debugf(format, args...)
}

func (l *Logger) Info(message string, fields ...map[string]any) {
	l.log.Info(message, convertFields(fields)...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.log.Sugar().Infof(format, args...)
}

func (l *Logger) Warn(message string, fields ...map[string]any) {
	l.log.Warn(message, convertFields(fields)...)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.log.Sugar().Warnf(format, args...)
}

func (l *Logger) Error(message string, fields ...map[string]any) {
	l.log.Error(message, convertFields(fields)...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.log.Sugar().Errorf(format, args...)
}

func (l *Logger) Fatal(message string, fields ...map[string]any) {
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
