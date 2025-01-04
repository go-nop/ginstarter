package log_test

import (
	"context"
	"testing"

	"github.com/go-nop/ginstarter/pkg/log"
	"github.com/go-nop/ginstarter/pkg/log/driver/zap"
	"github.com/stretchr/testify/assert"
)

func TestNewLoggerSingleton(t *testing.T) {
	t.Run("should only create one logger instance", func(t *testing.T) {
		logger1 := log.NewLogger(log.ZapDriver)
		logger2 := log.NewLogger(log.LogrusDriver)

		assert.Equal(t, logger1, logger2)
		assert.IsType(t, &zap.Logger{}, logger1)
	})
}

func TestLogging(t *testing.T) {
	_ = log.NewLogger(log.ZapDriver)

	ctx := context.Background()

	t.Run("should log debug message", func(t *testing.T) {
		assert.NotPanics(t, func() {
			log.Debug(ctx, "debug message", map[string]any{"key": "value"})
		})
	})

	t.Run("should log info message", func(t *testing.T) {
		assert.NotPanics(t, func() {
			log.Info(ctx, "info message", map[string]any{"key": "value"})
		})
	})

	t.Run("should log warn message", func(t *testing.T) {
		assert.NotPanics(t, func() {
			log.Warn(ctx, "warn message", map[string]any{"key": "value"})
		})
	})

	t.Run("should log error message", func(t *testing.T) {
		assert.NotPanics(t, func() {
			log.Error(ctx, "error message", map[string]any{"key": "value"})
		})
	})

	// t.Run("should log fatal message", func(t *testing.T) {
	// 	assert.NotPanics(t, func() {
	// 		// Note: Fatal will call os.Exit(1), so we don't actually want to call it in a test.
	// 		// Instead, we can just ensure the method exists and doesn't panic.
	// 		log.Fatal(ctx, "fatal message", map[string]any{"key": "value"})
	// 	})
	// })
}
