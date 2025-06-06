package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestInternalSetUsage_Deduplication(t *testing.T) {
	p := &Params{}
	p.InternalSetUsage([]string{"foo", "bar"})
	assert.ElementsMatch(t, []string{"foo", "bar"}, p.usage)

	// Add duplicate and new usage
	p.InternalSetUsage([]string{"bar", "baz"})
	// Should contain all unique values
	assert.ElementsMatch(t, []string{"foo", "bar", "baz"}, p.usage)
}

func TestInternalSetUsage_Nil(t *testing.T) {
	p := &Params{}
	p.InternalSetUsage(nil)
	assert.Empty(t, p.usage)
}

func TestInternalSetUsage_Empty(t *testing.T) {
	p := &Params{}
	p.InternalSetUsage([]string{})
	assert.Empty(t, p.usage)
}

func TestInternalSetUsage_Idempotent(t *testing.T) {
	p := &Params{}
	p.InternalSetUsage([]string{"foo"})
	p.InternalSetUsage([]string{"foo"})
	assert.Equal(t, []string{"foo"}, p.usage)
}
