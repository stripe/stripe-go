package stripe

import (
	"bytes"
	"context"
	"testing"

	assert "github.com/stretchr/testify/require"
)

//
// Tests
//

func TestDefaultLeveledLogger(t *testing.T) {
	_, ok := DefaultLeveledLogger.(*LeveledLogger)
	assert.True(t, ok)
}

//
// LeveledLogger
//

func TestLeveledLoggerDebugf(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &LeveledLogger{stdoutOverride: &stdout, stderrOverride: &stderr}

	{
		clearBuffers(&stdout, &stderr)
		logger.Level = LevelDebug

		logger.Debugf("test")
		assert.Equal(t, "[DEBUG] test\n", stdout.String())
		assert.Equal(t, "", stderr.String())
	}

	// Expect no logging
	for _, level := range []Level{LevelInfo, LevelWarn, LevelError} {
		clearBuffers(&stdout, &stderr)
		logger.Level = level

		logger.Debugf("test")
		assert.Equal(t, "", stdout.String())
		assert.Equal(t, "", stderr.String())
	}
}

func TestLeveledLoggerInfof(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &LeveledLogger{stdoutOverride: &stdout, stderrOverride: &stderr}

	for _, level := range []Level{LevelDebug, LevelInfo} {
		clearBuffers(&stdout, &stderr)
		logger.Level = level

		logger.Infof("test")
		assert.Equal(t, "[INFO] test\n", stdout.String())
		assert.Equal(t, "", stderr.String())
	}

	// Expect no logging
	for _, level := range []Level{LevelWarn, LevelError} {
		clearBuffers(&stdout, &stderr)
		logger.Level = level

		logger.Infof("test")
		assert.Equal(t, "", stdout.String())
		assert.Equal(t, "", stderr.String())
	}
}

func TestLeveledLoggerWarnf(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &LeveledLogger{stdoutOverride: &stdout, stderrOverride: &stderr}

	for _, level := range []Level{LevelDebug, LevelInfo, LevelWarn} {
		clearBuffers(&stdout, &stderr)
		logger.Level = level

		logger.Warnf("test")
		assert.Equal(t, "", stdout.String())
		assert.Equal(t, "[WARN] test\n", stderr.String())
	}

	// Expect no logging
	{
		clearBuffers(&stdout, &stderr)
		logger.Level = LevelError

		logger.Warnf("test")
		assert.Equal(t, "", stdout.String())
		assert.Equal(t, "", stderr.String())
	}
}

func TestLeveledLoggerErrorf(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &LeveledLogger{stdoutOverride: &stdout, stderrOverride: &stderr}

	for _, level := range []Level{LevelDebug, LevelInfo, LevelWarn, LevelError} {
		clearBuffers(&stdout, &stderr)
		logger.Level = level

		logger.Errorf("test")
		assert.Equal(t, "", stdout.String())
		assert.Equal(t, "[ERROR] test\n", stderr.String())
	}
}

//
// ContextLeveledLogger
//

// testContextLogger is a test implementation of ContextLeveledLoggerInterface
type testContextLogger struct {
	*LeveledLogger
	lastContext context.Context
}

func (l *testContextLogger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.lastContext = ctx
	l.LeveledLogger.Debugf(format, v...)
}

func (l *testContextLogger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.lastContext = ctx
	l.LeveledLogger.Infof(format, v...)
}

func (l *testContextLogger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l.lastContext = ctx
	l.LeveledLogger.Warnf(format, v...)
}

func (l *testContextLogger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.lastContext = ctx
	l.LeveledLogger.Errorf(format, v...)
}

func TestContextLeveledLoggerDebugf(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelDebug,
		},
	}

	ctx := context.WithValue(context.Background(), contextKey("trace_id"), "test-123")
	logger.Debugf(ctx, "test message")

	assert.Equal(t, "[DEBUG] test message\n", stdout.String())
	assert.Equal(t, "", stderr.String())
	assert.Equal(t, ctx, logger.lastContext, "Context should be passed to logger")
}

func TestContextLeveledLoggerInfof(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelInfo,
		},
	}

	ctx := context.WithValue(context.Background(), contextKey("trace_id"), "test-456")
	logger.Infof(ctx, "test message")

	assert.Equal(t, "[INFO] test message\n", stdout.String())
	assert.Equal(t, "", stderr.String())
	assert.Equal(t, ctx, logger.lastContext, "Context should be passed to logger")
}

func TestContextLeveledLoggerWarnf(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelWarn,
		},
	}

	ctx := context.WithValue(context.Background(), contextKey("trace_id"), "test-789")
	logger.Warnf(ctx, "test message")

	assert.Equal(t, "", stdout.String())
	assert.Equal(t, "[WARN] test message\n", stderr.String())
	assert.Equal(t, ctx, logger.lastContext, "Context should be passed to logger")
}

func TestContextLeveledLoggerErrorf(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelError,
		},
	}

	ctx := context.WithValue(context.Background(), contextKey("trace_id"), "test-abc")
	logger.Errorf(ctx, "test message")

	assert.Equal(t, "", stdout.String())
	assert.Equal(t, "[ERROR] test message\n", stderr.String())
	assert.Equal(t, ctx, logger.lastContext, "Context should be passed to logger")
}

func TestContextLeveledLoggerNilContext(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &testContextLogger{
		LeveledLogger: &LeveledLogger{
			stdoutOverride: &stdout,
			stderrOverride: &stderr,
			Level:          LevelInfo,
		},
	}

	// Should handle nil context gracefully
	//lint:ignore SA1012 Intentionally testing nil context handling
	logger.Infof(nil, "test message")

	assert.Equal(t, "[INFO] test message\n", stdout.String())
	assert.Equal(t, "", stderr.String())
	assert.Nil(t, logger.lastContext, "Nil context should be passed as-is")
}

//
// Private functions
//

func clearBuffers(buffers ...*bytes.Buffer) {
	for _, b := range buffers {
		b.Truncate(0)
	}
}
