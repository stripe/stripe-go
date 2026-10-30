//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"github.com/shopspring/decimal"
	"time"
)

type V2BillingContractTermQuantityChange struct {
	APIResource
	// The ID of the parent contract.
	Contract string `json:"contract"`
	// When this quantity change was created.
	Created time.Time `json:"created"`
	// When the quantity change takes effect.
	EffectiveAt time.Time `json:"effective_at"`
	// The ID of the quantity change.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The quantity at the effective time.
	Quantity decimal.Decimal `json:"quantity"`
	// The ID of the recurring-price term.
	Term string `json:"term"`
}
