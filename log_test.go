package stripe

import (
	"bytes"
	"log"
	"testing"

	assert "github.com/stretchr/testify/require"
)

//
// Tests
//

func TestDefaultLeveledLogger(t *testing.T) {
	// We don't set DefaultLeveledLogger by default for backwards compatibility
	// reasons. If we did, then it would override Logger, which many people
	// have been setting over the years.
	assert.Nil(t, DefaultLeveledLogger)

	// Logger continues to be set by default so that we have a non-nil object
	// to log against.
	_, ok := Logger.(*log.Logger)
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
// leveledLoggerPrintferShim
//

func TestLeveledLoggerPrintferShimDebugf(t *testing.T) {
	var stdout bytes.Buffer
	logger := &leveledLoggerPrintferShim{logger: log.New(&stdout, "", 0)}

	{
		clearBuffers(&stdout)
		logger.level = printferLevelDebug

		logger.Debugf("test")
		assert.Equal(t, "test\n", stdout.String())
	}

	// Expect no logging
	for _, level := range []printferLevel{printferLevelInfo, printferLevelError} {
		clearBuffers(&stdout)
		logger.level = level

		logger.Debugf("test")
		assert.Equal(t, "", stdout.String())
	}
}

func TestLeveledLoggerPrintferShimInfof(t *testing.T) {
	var stdout bytes.Buffer
	logger := &leveledLoggerPrintferShim{logger: log.New(&stdout, "", 0)}

	for _, level := range []printferLevel{printferLevelDebug, printferLevelInfo} {
		clearBuffers(&stdout)
		logger.level = level

		logger.Infof("test")
		assert.Equal(t, "test\n", stdout.String())
	}

	// Expect no logging
	for _, level := range []printferLevel{printferLevelError} {
		clearBuffers(&stdout)
		logger.level = level

		logger.Infof("test")
		assert.Equal(t, "", stdout.String())
	}
}

// Note: behaves identically to Errorf because historically there was no
// warning level.
func TestLeveledLoggerPrintferShimWarnf(t *testing.T) {
	var stdout bytes.Buffer
	logger := &leveledLoggerPrintferShim{logger: log.New(&stdout, "", 0)}

	for _, level := range []printferLevel{printferLevelDebug, printferLevelInfo, printferLevelError} {
		clearBuffers(&stdout)
		logger.level = level

		logger.Warnf("test")
		assert.Equal(t, "test\n", stdout.String())
	}
}

func TestLeveledLoggerPrintferShimErrorf(t *testing.T) {
	var stdout bytes.Buffer
	logger := &leveledLoggerPrintferShim{logger: log.New(&stdout, "", 0)}

	for _, level := range []printferLevel{printferLevelDebug, printferLevelInfo, printferLevelError} {
		clearBuffers(&stdout)
		logger.level = level

		logger.Errorf("test")
		assert.Equal(t, "test\n", stdout.String())
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
