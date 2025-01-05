package log

import (
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
	Debug(message string, fields ...map[string]any)
	// Debugf logs a formatted message at the debug level.
	Debugf(format string, args ...any)
	// Info logs a message at the info level.
	Info(message string, fields ...map[string]any)
	// Infof logs a formatted message at the info level.
	Infof(format string, args ...any)
	// Warn logs a message at the warn level.
	Warn(message string, fields ...map[string]any)
	// Warnf logs a formatted message at the warn level.
	Warnf(format string, args ...any)
	// Error logs a message at the error level.
	Error(message string, fields ...map[string]any)
	// Errorf logs a formatted message at the error level.
	Errorf(format string, args ...any)
	// Fatal logs a message at the fatal level.
	Fatal(message string, fields ...map[string]any)
}

// Driver is the type for the logger driver.
type Driver string

const (
	// ZapDriver is the zap driver.
	ZapDriver Driver = "zap"
	// LogrusDriver is the logrus driver.
	LogrusDriver Driver = "logrus"
)

// NewLogger returns a new Logger instance.
func NewLogger(driver Driver) Logger {
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
func Debug(message string, fields ...map[string]any) {
	log.Debug(message, fields...)
}

// Debugf logs a formatted message at the debug level.
func Debugf(format string, args ...any) {
	log.Debugf(format, args...)
}

// Info logs a message at the info level.
func Info(message string, fields ...map[string]any) {
	log.Info(message, fields...)
}

// Infof logs a formatted message at the info level.
func Infof(format string, args ...any) {
	log.Infof(format, args...)
}

// Warn logs a message at the warn level.
func Warn(message string, fields ...map[string]any) {
	log.Warn(message, fields...)
}

// Warnf logs a formatted message at the warn level.
func Warnf(format string, args ...any) {
	log.Warnf(format, args...)
}

// Error logs a message at the error level.
func Error(message string, fields ...map[string]any) {
	log.Error(message, fields...)
}

// Errorf logs a formatted message at the error level.
func Errorf(format string, args ...any) {
	log.Errorf(format, args...)
}

// Fatal logs a message at the fatal level.
func Fatal(message string, fields ...map[string]any) {
	log.Fatal(message, fields...)
}
