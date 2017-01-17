package source

import (
	"math/rand"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestSource(t *testing.T) {

	sourceParams := &stripe.SourceObjectParams{
		Type:     "bitcoin",
		Amount:   1000,
		Currency: currency.USD,
		Owner: &stripe.SourceOwnerParams{
			Email: "jenny.rosen@example.com",
		},
	}

	s, err := New(sourceParams)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	if s.ID == "" {
		t.Errorf("ID is not set %v", s.ID)
	}
	if s.Created == 0 {
		t.Errorf("Created date is not set")
	}

	if s.Currency != "usd" {
		t.Errorf("Currency is invalid: %v", s.Currency)
	}
	if s.Amount != 1000 {
		t.Errorf("Amount is invalid: %v", s.Amount)
	}

	if s.Owner.Email != "jenny.rosen@example.com" {
		t.Errorf("Email is invalid: %v", s.Owner.Email)
	}

	if s.Status != stripe.SourceStatusPending {
		t.Errorf("Source status is invalid: %v", s.Status)
	}
	if s.Flow != stripe.FlowReceiver {
		t.Errorf("Source flow is invalid: %v", s.Flow)
	}

	if s.Receiver.RefundAttributesStatus != stripe.RefundAttributesMissing {
		t.Errorf("Source receiver refund attributes status is invalid: %v", s.Receiver.RefundAttributesStatus)
	}
	if s.Receiver.RefundAttributesMethod != stripe.RefundAttributesEmail {
		t.Errorf("Source receiver refund attributes method is invalid: %v", s.Receiver.RefundAttributesMethod)
	}
	if s.Receiver.Address == "" {
		t.Errorf("Source receiver address is invalid: %v", s.Receiver.Address)
	}

	if addr, ok := s.TypeData["address"]; !ok || addr == "" {
		t.Errorf("Source typedata address is invalid: %v", addr)
	}
	if amount, ok := s.TypeData["amount_received"]; !ok || amount != float64(0) {
		t.Errorf("Source typedata address is invalid: %v", amount)
	}
}

func TestSourceUpdate(t *testing.T) {
	params := &stripe.SourceObjectParams{
		Type:     "bitcoin",
		Amount:   1000,
		Currency: currency.USD,
		Owner: &stripe.SourceOwnerParams{
			Email: "jenny.rosen@example.com",
		},
	}

	source, _ := New(params)

	params = &stripe.SourceObjectParams{}
	params.AddMeta("foo", "bar")

	source, err := Update(source.ID, params)
	if err != nil {
		t.Error(err)
	}

	if source.Meta["foo"] != params.Meta["foo"] {
		t.Errorf("Meta %v does not match expected Meta %v\n", source.Meta, params.Meta)
	}
}
