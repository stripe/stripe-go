package testing

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCompareVersions(t *testing.T) {
	assert.Equal(t, 0, compareVersions("1.2.3", "1.2.3"))

	assert.Equal(t, 1, compareVersions("1.2.3", "1.2.4"))
	assert.Equal(t, -1, compareVersions("1.2.4", "1.2.3"))

	assert.Equal(t, 1, compareVersions("0.2.3", "1.2.3"))
	assert.Equal(t, -1, compareVersions("1.2.3", "0.2.3"))

	assert.Equal(t, 1, compareVersions("1.2.3", "1.22.3"))
	assert.Equal(t, -1, compareVersions("1.22.3", "1.2.3"))

	assert.Equal(t, 1, compareVersions("1.2", "1.22.3"))
	assert.Equal(t, -1, compareVersions("1.22.3", "1.2"))

	assert.Equal(t, 1, compareVersions("1", "1.22.3"))
	assert.Equal(t, -1, compareVersions("1.22.3", "1"))
}
