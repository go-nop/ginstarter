package logrus

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log *logrus.Logger
}

func NewLogrusLogger() *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return &Logger{log: logger}
}

func (l *Logger) Debug(ctx context.Context, message string, fields ...map[string]any) {
	l.log.WithFields(convertFields(fields)).Debug(message)
}

func (l *Logger) Info(ctx context.Context, message string, fields ...map[string]any) {
	l.log.WithFields(convertFields(fields)).Info(message)
}

func (l *Logger) Warn(ctx context.Context, message string, fields ...map[string]any) {
	l.log.WithFields(convertFields(fields)).Warn(message)
}

func (l *Logger) Error(ctx context.Context, message string, fields ...map[string]any) {
	l.log.WithFields(convertFields(fields)).Error(message)
}

func (l *Logger) Fatal(ctx context.Context, message string, fields ...map[string]any) {
	l.log.WithFields(convertFields(fields)).Fatal(message)
}

func (l *Logger) SetLevel(level string) {
	switch level {
	case "debug":
		l.log.SetLevel(logrus.DebugLevel)
	case "info":
		l.log.SetLevel(logrus.InfoLevel)
	case "warn":
		l.log.SetLevel(logrus.WarnLevel)
	case "error":
		l.log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		l.log.SetLevel(logrus.FatalLevel)
	default:
		l.log.SetLevel(logrus.InfoLevel)
	}
}

func convertFields(fields []map[string]any) logrus.Fields {
	logrusFields := logrus.Fields{}
	for _, field := range fields {
		for key, value := range field {
			logrusFields[key] = value
		}
	}
	return logrusFields
}
