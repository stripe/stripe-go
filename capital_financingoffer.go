//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of financing being offered.
type CapitalFinancingOfferFinancingType string

// List of values that CapitalFinancingOfferFinancingType can take
const (
	CapitalFinancingOfferFinancingTypeCashAdvance CapitalFinancingOfferFinancingType = "cash_advance"
	CapitalFinancingOfferFinancingTypeFlexLoan    CapitalFinancingOfferFinancingType = "flex_loan"
)

// Describes the type of user the offer is being extended to.
type CapitalFinancingOfferOfferedTermsCampaignType string

// List of values that CapitalFinancingOfferOfferedTermsCampaignType can take
const (
	CapitalFinancingOfferOfferedTermsCampaignTypeNewlyEligibleUser      CapitalFinancingOfferOfferedTermsCampaignType = "newly_eligible_user"
	CapitalFinancingOfferOfferedTermsCampaignTypePreviouslyEligibleUser CapitalFinancingOfferOfferedTermsCampaignType = "previously_eligible_user"
	CapitalFinancingOfferOfferedTermsCampaignTypeRepeatUser             CapitalFinancingOfferOfferedTermsCampaignType = "repeat_user"
)

// Financing product identifier.
type CapitalFinancingOfferProductType string

// List of values that CapitalFinancingOfferProductType can take
const (
	CapitalFinancingOfferProductTypeRefill   CapitalFinancingOfferProductType = "refill"
	CapitalFinancingOfferProductTypeStandard CapitalFinancingOfferProductType = "standard"
)

// The current status of the offer.
type CapitalFinancingOfferStatus string

// List of values that CapitalFinancingOfferStatus can take
const (
	CapitalFinancingOfferStatusAccepted    CapitalFinancingOfferStatus = "accepted"
	CapitalFinancingOfferStatusCanceled    CapitalFinancingOfferStatus = "canceled"
	CapitalFinancingOfferStatusCompleted   CapitalFinancingOfferStatus = "completed"
	CapitalFinancingOfferStatusDelivered   CapitalFinancingOfferStatus = "delivered"
	CapitalFinancingOfferStatusExpired     CapitalFinancingOfferStatus = "expired"
	CapitalFinancingOfferStatusFullyRepaid CapitalFinancingOfferStatus = "fully_repaid"
	CapitalFinancingOfferStatusPaidOut     CapitalFinancingOfferStatus = "paid_out"
	CapitalFinancingOfferStatusRejected    CapitalFinancingOfferStatus = "rejected"
	CapitalFinancingOfferStatusReplaced    CapitalFinancingOfferStatus = "replaced"
	CapitalFinancingOfferStatusUndelivered CapitalFinancingOfferStatus = "undelivered"
)

// See [financing_type](https://stripe.com/docs/api/capital/connect_financing_object#financing_offer_object-financing_type).
type CapitalFinancingOfferType string

// List of values that CapitalFinancingOfferType can take
const (
	CapitalFinancingOfferTypeCashAdvance CapitalFinancingOfferType = "cash_advance"
	CapitalFinancingOfferTypeFlexLoan    CapitalFinancingOfferType = "flex_loan"
)

// Retrieves the financing offers available for Connected accounts that belong to your platform.
type CapitalFinancingOfferListParams struct {
	ListParams `form:"*"`
	// limit list to offers belonging to given connected account
	ConnectedAccount *string           `form:"connected_account"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// limit list to offers with given status
	Status *string `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *CapitalFinancingOfferListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Get the details of the financing offer
type CapitalFinancingOfferParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CapitalFinancingOfferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Acknowledges that platform has received and delivered the financing_offer to
// the intended merchant recipient.
type CapitalFinancingOfferMarkDeliveredParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CapitalFinancingOfferMarkDeliveredParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// This is an object representing the terms of an offer of financing from
// Stripe Capital to a Connected account. This resource represents
// the terms accepted by the Connected account, which may differ from those
// offered.
type CapitalFinancingOfferAcceptedTerms struct {
	// Amount of financing offered, in minor units.
	AdvanceAmount int64 `json:"advance_amount"`
	// Currency that the financing offer is transacted in. For example, `usd`.
	Currency Currency `json:"currency"`
	// Fixed fee amount, in minor units.
	FeeAmount int64 `json:"fee_amount"`
	// Populated when the `product_type` of the `financingoffer` is `refill`.
	// Represents the discount amount on remaining premium for the existing loan at payout time.
	PreviousFinancingFeeDiscountAmount int64 `json:"previous_financing_fee_discount_amount"`
	// Per-transaction rate at which Stripe will withhold funds to repay the financing.
	WithholdRate float64 `json:"withhold_rate"`
}

// This is an object representing the terms of an offer of financing from
// Stripe Capital to a Connected account. This resource represents
// both the terms offered to the Connected account.
type CapitalFinancingOfferOfferedTerms struct {
	// Amount of financing offered, in minor units.
	AdvanceAmount int64 `json:"advance_amount"`
	// Describes the type of user the offer is being extended to.
	CampaignType CapitalFinancingOfferOfferedTermsCampaignType `json:"campaign_type"`
	// Currency that the financing offer is transacted in. For example, `usd`.
	Currency Currency `json:"currency"`
	// Fixed fee amount, in minor units.
	FeeAmount int64 `json:"fee_amount"`
	// Populated when the `product_type` of the `financingoffer` is `refill`.
	// Represents the discount rate percentage on remaining fee on the existing loan. When the `financing_offer`
	// is paid out, the `previous_financing_fee_discount_amount` will be computed as the multiple of this rate
	// and the remaining fee.
	PreviousFinancingFeeDiscountRate float64 `json:"previous_financing_fee_discount_rate"`
	// Per-transaction rate at which Stripe will withhold funds to repay the financing.
	WithholdRate float64 `json:"withhold_rate"`
}

// This is an object representing an offer of financing from
// Stripe Capital to a Connect subaccount.
type CapitalFinancingOffer struct {
	APIResource
	// This is an object representing the terms of an offer of financing from
	// Stripe Capital to a Connected account. This resource represents
	// the terms accepted by the Connected account, which may differ from those
	// offered.
	AcceptedTerms *CapitalFinancingOfferAcceptedTerms `json:"accepted_terms"`
	// The ID of the merchant associated with this financing object.
	Account string `json:"account"`
	// Time at which the offer was created. Given in seconds since unix epoch.
	Created int64 `json:"created"`
	// Time at which the offer expires. Given in seconds since unix epoch.
	ExpiresAfter float64 `json:"expires_after"`
	// The type of financing being offered.
	FinancingType CapitalFinancingOfferFinancingType `json:"financing_type"`
	// A unique identifier for the financing object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The object type: financing_offer.
	Object string `json:"object"`
	// This is an object representing the terms of an offer of financing from
	// Stripe Capital to a Connected account. This resource represents
	// both the terms offered to the Connected account.
	OfferedTerms *CapitalFinancingOfferOfferedTerms `json:"offered_terms"`
	// Financing product identifier.
	ProductType CapitalFinancingOfferProductType `json:"product_type"`
	// The ID of the financing offer that replaced this offer.
	Replacement string `json:"replacement"`
	// The ID of the financing offer that this offer is a replacement for.
	ReplacementFor string `json:"replacement_for"`
	// The current status of the offer.
	Status CapitalFinancingOfferStatus `json:"status"`
	// See [financing_type](https://stripe.com/docs/api/capital/connect_financing_object#financing_offer_object-financing_type).
	Type CapitalFinancingOfferType `json:"type"`
}

// CapitalFinancingOfferList is a list of FinancingOffers as retrieved from a list endpoint.
type CapitalFinancingOfferList struct {
	APIResource
	ListMeta
	Data []*CapitalFinancingOffer `json:"data"`
}
