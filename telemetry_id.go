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
	telemetryIdOnce  sync.Once
	telemetryIdValue string
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

func getTelemetryId() string {
	telemetryIdOnce.Do(func() {
		configDir := getConfigDir()
		if configDir == "" {
			return
		}

		filePath := filepath.Join(configDir, "telemetry_id")

		data, err := os.ReadFile(filePath)
		if err == nil {
			content := strings.TrimSpace(string(data))
			if content != "" {
				telemetryIdValue = content
				return
			}
		}

		b := make([]byte, 16)
		if _, err := rand.Read(b); err != nil {
			return
		}
		newId := hex.EncodeToString(b)

		if err := os.MkdirAll(configDir, 0700); err != nil {
			return
		}
		if err := os.WriteFile(filePath, []byte(newId), 0600); err != nil {
			return
		}

		telemetryIdValue = newId
	})
	return telemetryIdValue
}

// resetTelemetryId resets the telemetry ID singleton for testing purposes.
func resetTelemetryId() {
	telemetryIdOnce = sync.Once{}
	telemetryIdValue = ""
}
