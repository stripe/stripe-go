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

func TestRedact(t *testing.T) {
	var stdout, stderr bytes.Buffer
	logger := &LeveledLogger{stdoutOverride: &stdout, stderrOverride: &stderr}
	for _, level := range []Level{LevelDebug, LevelInfo, LevelWarn, LevelError} {
		assertRedact := func(expectedMsg string, msg string, v ...interface{}) {
			clearBuffers(&stdout, &stderr)
			logger.Level = level

			logger.Errorf(msg, v...)
			assert.Equal(t, "", stdout.String())
			assert.Equal(t, "[ERROR] "+expectedMsg+"\n", stderr.String())
		}

		// Redacts secret and restricted keys in test and live mode
		assertRedact("sk_live_01****6789", "sk_live_0123456789")
		assertRedact("sk_test_01****6789", "sk_test_0123456789")
		assertRedact("rk_live_01****6789", "rk_live_0123456789")
		assertRedact("rk_test_01****6789", "rk_test_0123456789")

		// Redacts key created by string interpolated
		assertRedact("sk_live_01****6789", "sk_live_%s", "0123456789")

		// Redacts key within a larger message
		assertRedact("foo sk_live_01****6789 foo", "foo sk_live_0123456789 foo")

		// Does not redact too-short keys
		assertRedact("sk_live_01", "sk_live_01")
		assertRedact("sk_live_01", "sk_live_%s", "01")
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
