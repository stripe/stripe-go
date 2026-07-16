package stripe

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var (
	telemetryIDOnce  sync.Once
	telemetryIDValue string
)

func getConfigDir() string {
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return ""
		}
		return filepath.Join(appData, "Stripe")
	}
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg != "" {
		return filepath.Join(xdg, "stripe")
	}
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return ""
	}
	return filepath.Join(home, ".config", "stripe")
}

// configDirOverride is used by tests to redirect file I/O to a temp directory.
var configDirOverride string

func getTelemetryID() string {
	telemetryIDOnce.Do(func() {
		configDir := configDirOverride
		if configDir == "" {
			configDir = getConfigDir()
		}
		if configDir == "" {
			return
		}

		filePath := filepath.Join(configDir, "telemetry_id")

		data, err := os.ReadFile(filePath)
		if err == nil {
			content := strings.TrimSpace(string(data))
			if content != "" {
				telemetryIDValue = content
				return
			}
		}

		b := make([]byte, 16)
		if _, err := rand.Read(b); err != nil {
			return
		}
		newID := hex.EncodeToString(b)

		if err := os.MkdirAll(configDir, 0755); err != nil {
			return
		}
		if err := os.WriteFile(filePath, []byte(newID), 0644); err != nil {
			return
		}

		telemetryIDValue = newID
	})
	return telemetryIDValue
}

// resetTelemetryID resets the telemetry ID singleton for testing purposes.
func resetTelemetryID() {
	telemetryIDOnce = sync.Once{}
	telemetryIDValue = ""
	configDirOverride = ""
}
