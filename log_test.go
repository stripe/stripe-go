package stripe

import (
	"bytes"
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
// Private functions
//

func clearBuffers(buffers ...*bytes.Buffer) {
	for _, b := range buffers {
		b.Truncate(0)
	}
}
