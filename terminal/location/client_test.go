package location

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestTerminalLocationDel(t *testing.T) {
	location, err := Del("loc_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "terminal.location", location.Object)
}

func TestTerminalLocationGet(t *testing.T) {
	location, err := Get("loc_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "terminal.location", location.Object)
}

func TestTerminalLocationList(t *testing.T) {
	i := List(&stripe.TerminalLocationListParams{})

	// Verify that we can get at least one location
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.TerminalLocation())
	assert.Equal(t, "terminal.location", i.TerminalLocation().Object)
	assert.NotNil(t, i.TerminalLocationList())
}

func TestTerminalLocationNew(t *testing.T) {
	location, err := New(&stripe.TerminalLocationParams{
		DisplayName: stripe.String("name"),
		Address: &stripe.AddressParams{
			Country:    stripe.String("US"),
			City:       stripe.String("San Francisco"),
			PostalCode: stripe.String("12345"),
			Line1:      stripe.String("line-1"),
			State:      stripe.String("CA"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "terminal.location", location.Object)
}

func TestTerminalLocationUpdate(t *testing.T) {
	location, err := Update("loc_123", &stripe.TerminalLocationParams{
		DisplayName: stripe.String("new name"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, "terminal.location", location.Object)
}
