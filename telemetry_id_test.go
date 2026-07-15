package stripe

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGetTelemetryIDReturns32CharHex(t *testing.T) {
	resetTelemetryID()
	t.Cleanup(resetTelemetryID)

	configDirOverride = t.TempDir()

	id := getTelemetryID()
	assert.Regexp(t, `^[0-9a-f]{32}$`, id)
}

func TestGetTelemetryIDIsCached(t *testing.T) {
	resetTelemetryID()
	t.Cleanup(resetTelemetryID)

	configDirOverride = t.TempDir()

	first := getTelemetryID()
	second := getTelemetryID()
	assert.Equal(t, first, second)
}

func TestGetTelemetryIDReadsExistingFile(t *testing.T) {
	resetTelemetryID()
	t.Cleanup(resetTelemetryID)

	dir := t.TempDir()
	configDirOverride = dir

	existingID := "aabbccddeeff00112233445566778899"
	err := os.WriteFile(filepath.Join(dir, "telemetry_id"), []byte(existingID), 0644)
	assert.NoError(t, err)

	id := getTelemetryID()
	assert.Equal(t, existingID, id)
}

func TestGetTelemetryIDGeneratesAndPersistsNewID(t *testing.T) {
	resetTelemetryID()
	t.Cleanup(resetTelemetryID)

	dir := t.TempDir()
	configDirOverride = dir

	id := getTelemetryID()
	assert.Regexp(t, `^[0-9a-f]{32}$`, id)

	data, err := os.ReadFile(filepath.Join(dir, "telemetry_id"))
	assert.NoError(t, err)
	assert.Equal(t, id, string(data))
}

func TestGetTelemetryIDReturnsEmptyWhenConfigDirEmpty(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod doesn't prevent writes on Windows")
	}
	resetTelemetryID()
	t.Cleanup(resetTelemetryID)

	// Use a read-only parent directory so MkdirAll fails on the child path.
	// getTelemetryID returns "" when it cannot persist the generated ID.
	roDir := t.TempDir()
	assert.NoError(t, os.Chmod(roDir, 0555))
	t.Cleanup(func() { _ = os.Chmod(roDir, 0755) })
	configDirOverride = filepath.Join(roDir, "stripe")

	id := getTelemetryID()
	assert.Equal(t, "", id)
}

func TestGetConfigDirRespectsXDGConfigHome(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("XDG is not used on Windows")
	}
	t.Setenv("XDG_CONFIG_HOME", "/custom/xdg")

	dir := getConfigDir()
	assert.Equal(t, "/custom/xdg/stripe", dir)
}

func TestGetConfigDirFallsBackToHomeConfig(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("XDG is not used on Windows")
	}
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("HOME", "/testhome")

	dir := getConfigDir()
	assert.Equal(t, "/testhome/.config/stripe", dir)
}

func TestGetTelemetryIDCreatesParentDirectories(t *testing.T) {
	resetTelemetryID()
	t.Cleanup(resetTelemetryID)

	// Point configDirOverride at a nested path that does not yet exist.
	base := t.TempDir()
	configDirOverride = filepath.Join(base, "nested", "stripe")

	id := getTelemetryID()
	assert.Regexp(t, `^[0-9a-f]{32}$`, id)

	// The directory should have been created.
	_, err := os.Stat(configDirOverride)
	assert.NoError(t, err)

	// And the file should exist inside it.
	data, err := os.ReadFile(filepath.Join(configDirOverride, "telemetry_id"))
	assert.NoError(t, err)
	assert.Equal(t, id, string(data))
}
