package stripe

import (
	"regexp"
	"testing"
)

func TestNewUUID4(t *testing.T) {
	id, err := newUUID4()
	if err != nil {
		t.Fatalf("newUUID4() returned error: %v", err)
	}

	uuidPattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	if !uuidPattern.MatchString(id) {
		t.Errorf("newUUID4() = %q, does not match UUID v4 pattern", id)
	}
}

func TestNewUUID4Uniqueness(t *testing.T) {
	id1, err := newUUID4()
	if err != nil {
		t.Fatalf("first newUUID4() returned error: %v", err)
	}
	id2, err := newUUID4()
	if err != nil {
		t.Fatalf("second newUUID4() returned error: %v", err)
	}
	if id1 == id2 {
		t.Errorf("newUUID4() returned the same value twice: %q", id1)
	}
}
