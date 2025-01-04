package log

import (
	"context"
	"sync"

	"github.com/go-nop/ginstarter/pkg/log/driver/logrus"
	"github.com/go-nop/ginstarter/pkg/log/driver/zap"
)

var (
	log        Logger
	loggerOnce sync.Once
)

// Logger is the interface that wraps the basic logging methods.
type Logger interface {
	// Debug logs a message at the debug level.
	Debug(ctx context.Context, message string, fields ...map[string]any)
	// Info logs a message at the info level.
	Info(ctx context.Context, message string, fields ...map[string]any)
	// Warn logs a message at the warn level.
	Warn(ctx context.Context, message string, fields ...map[string]any)
	// Error logs a message at the error level.
	Error(ctx context.Context, message string, fields ...map[string]any)
	// Fatal logs a message at the fatal level.
	Fatal(ctx context.Context, message string, fields ...map[string]any)
}

const (
	// ZapDriver is the zap driver.
	ZapDriver = "zap"
	// LogrusDriver is the logrus driver.
	LogrusDriver = "logrus"
)

// NewLogger returns a new Logger instance.
func NewLogger(driver string) Logger {
	loggerOnce.Do(func() {
		switch driver {
		case LogrusDriver:
			log = logrus.NewLogrusLogger()
		default:
			log = zap.NewZapLogger()
		}
	})
	return log
}

// Debug logs a message at the debug level.
func Debug(ctx context.Context, message string, fields ...map[string]any) {
	log.Debug(ctx, message, fields...)
}

// Info logs a message at the info level.
func Info(ctx context.Context, message string, fields ...map[string]any) {
	log.Info(ctx, message, fields...)
}

// Warn logs a message at the warn level.
func Warn(ctx context.Context, message string, fields ...map[string]any) {
	log.Warn(ctx, message, fields...)
}

// Error logs a message at the error level.
func Error(ctx context.Context, message string, fields ...map[string]any) {
	log.Error(ctx, message, fields...)
}

// Fatal logs a message at the fatal level.
func Fatal(ctx context.Context, message string, fields ...map[string]any) {
	log.Fatal(ctx, message, fields...)
}
