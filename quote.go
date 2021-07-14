//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/v72/form"
)

// The status of the most recent automated tax calculation for this quote.
type QuoteAutomaticTaxStatus string

// List of values that QuoteAutomaticTaxStatus can take
const (
	QuoteAutomaticTaxStatusComplete               QuoteAutomaticTaxStatus = "complete"
	QuoteAutomaticTaxStatusFailed                 QuoteAutomaticTaxStatus = "failed"
	QuoteAutomaticTaxStatusRequiresLocationInputs QuoteAutomaticTaxStatus = "requires_location_inputs"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay invoices at the end of the subscription cycle or on finalization using the default payment method attached to the subscription or customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions. Defaults to `charge_automatically`.
type QuoteCollectionMethod string

// List of values that QuoteCollectionMethod can take
const (
	QuoteCollectionMethodChargeAutomatically QuoteCollectionMethod = "charge_automatically"
	QuoteCollectionMethodSendInvoice         QuoteCollectionMethod = "send_invoice"
)

// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
type QuoteComputedRecurringInterval string

// List of values that QuoteComputedRecurringInterval can take
const (
	QuoteComputedRecurringIntervalDay   QuoteComputedRecurringInterval = "day"
	QuoteComputedRecurringIntervalMonth QuoteComputedRecurringInterval = "month"
	QuoteComputedRecurringIntervalWeek  QuoteComputedRecurringInterval = "week"
	QuoteComputedRecurringIntervalYear  QuoteComputedRecurringInterval = "year"
)

// The status of the quote.
type QuoteStatus string

// List of values that QuoteStatus can take
const (
	QuoteStatusAccepted QuoteStatus = "accepted"
	QuoteStatusCanceled QuoteStatus = "canceled"
	QuoteStatusDraft    QuoteStatus = "draft"
	QuoteStatusOpen     QuoteStatus = "open"
)

// Retrieves the quote with the given ID.
type QuoteParams struct {
	Params                `form:"*"`
	ApplicationFeeAmount  *int64                       `form:"application_fee_amount"`
	ApplicationFeePercent *float64                     `form:"application_fee_percent"`
	AutomaticTax          *QuoteAutomaticTaxParams     `form:"automatic_tax"`
	CollectionMethod      *string                      `form:"collection_method"`
	Customer              *string                      `form:"customer"`
	DefaultTaxRates       []*string                    `form:"default_tax_rates"`
	Description           *string                      `form:"description"`
	Discounts             []*QuoteDiscountParams       `form:"discounts"`
	ExpiresAt             *int64                       `form:"expires_at"`
	Footer                *string                      `form:"footer"`
	FromQuote             *QuoteFromQuoteParams        `form:"from_quote"`
	Header                *string                      `form:"header"`
	InvoiceSettings       *QuoteInvoiceSettingsParams  `form:"invoice_settings"`
	LineItems             []*QuoteLineItemParams       `form:"line_items"`
	OnBehalfOf            *string                      `form:"on_behalf_of"`
	SubscriptionData      *QuoteSubscriptionDataParams `form:"subscription_data"`
	TransferData          *QuoteTransferDataParams     `form:"transfer_data"`
}

// Settings for automatic tax lookup for this quote and resulting invoices and subscriptions.
type QuoteAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// The discounts applied to the quote. You can only set up to one discount.
type QuoteDiscountParams struct {
	Coupon   *string `form:"coupon"`
	Discount *string `form:"discount"`
}

// All invoices will be billed using the specified settings.
type QuoteInvoiceSettingsParams struct {
	DaysUntilDue *int64 `form:"days_until_due"`
}

// The recurring components of a price such as `interval` and `usage_type`.
type QuoteLineItemPriceDataRecurringParams struct {
	Interval      *string `form:"interval"`
	IntervalCount *int64  `form:"interval_count"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
type QuoteLineItemPriceDataParams struct {
	Currency          *string                                `form:"currency"`
	Product           *string                                `form:"product"`
	Recurring         *QuoteLineItemPriceDataRecurringParams `form:"recurring"`
	TaxBehavior       *string                                `form:"tax_behavior"`
	UnitAmount        *int64                                 `form:"unit_amount"`
	UnitAmountDecimal *float64                               `form:"unit_amount_decimal,high_precision"`
}

// A list of line items the customer is being quoted for. Each line item includes information about the product, the quantity, and the resulting cost.
type QuoteLineItemParams struct {
	ID        *string                       `form:"id"`
	Price     *string                       `form:"price"`
	PriceData *QuoteLineItemPriceDataParams `form:"price_data"`
	Quantity  *int64                        `form:"quantity"`
	TaxRates  []*string                     `form:"tax_rates"`
}

// When creating a subscription or subscription schedule, the specified configuration data will be used. There must be at least one line item with a recurring price for a subscription or subscription schedule to be created. A subscription schedule is created if `subscription_data[effective_date]` is present and in the future, otherwise a subscription is created.
type QuoteSubscriptionDataParams struct {
	EffectiveDate                 *int64 `form:"effective_date"`
	EffectiveDateCurrentPeriodEnd *bool  `form:"-"` // See custom AppendTo
	TrialPeriodDays               *int64 `form:"trial_period_days"`
}

// AppendTo implements custom encoding logic for QuoteSubscriptionDataParams.
func (q *QuoteSubscriptionDataParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(q.EffectiveDateCurrentPeriodEnd) {
		body.Add(form.FormatKey(append(keyParts, "effective_date")), "current_period_end")
	}
}

// The data with which to automatically create a Transfer for each of the invoices.
type QuoteTransferDataParams struct {
	Amount        *int64   `form:"amount"`
	AmountPercent *float64 `form:"amount_percent"`
	Destination   *string  `form:"destination"`
}

// Clone an existing quote. The new quote will be created in `status=draft`. When using this parameter, you cannot specify any other parameters except for `expires_at`.
type QuoteFromQuoteParams struct {
	IsRevision *bool   `form:"is_revision"`
	Quote      *string `form:"quote"`
}

// Returns a list of your quotes.
type QuoteListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"customer"`
	Status     *string `form:"status"`
}

// Cancels the quote.
type QuoteCancelParams struct {
	Params `form:"*"`
}

// Finalizes the quote.
type QuoteFinalizeQuoteParams struct {
	Params    `form:"*"`
	ExpiresAt *int64 `form:"expires_at"`
}

// Accepts the specified quote.
type QuoteAcceptParams struct {
	Params `form:"*"`
}

// When retrieving a quote, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type QuoteListLineItemsParams struct {
	ListParams `form:"*"`
	Quote      *string `form:"-"` // Included in URL
}

// When retrieving a quote, there is an includable upfront.line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of upfront line items.
type QuoteListComputedUpfrontLineItemsParams struct {
	ListParams `form:"*"`
	Quote      *string `form:"-"` // Included in URL
}

// Download the PDF for a finalized quote
type QuotePDFParams struct {
	Params `form:"*"`
}
type QuoteAutomaticTax struct {
	Enabled bool                    `json:"enabled"`
	Status  QuoteAutomaticTaxStatus `json:"status"`
}

// The aggregated line item discounts.
type QuoteComputedRecurringTotalDetailsBreakdownDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// The aggregated line item tax amounts by rate.
type QuoteComputedRecurringTotalDetailsBreakdownTax struct {
	Amount int64    `json:"amount"`
	Rate   *TaxRate `json:"rate"`
}
type QuoteComputedRecurringTotalDetailsBreakdown struct {
	Discounts []*QuoteComputedRecurringTotalDetailsBreakdownDiscount `json:"discounts"`
	Taxes     []*QuoteComputedRecurringTotalDetailsBreakdownTax      `json:"taxes"`
}
type QuoteComputedRecurringTotalDetails struct {
	AmountDiscount int64                                        `json:"amount_discount"`
	AmountShipping int64                                        `json:"amount_shipping"`
	AmountTax      int64                                        `json:"amount_tax"`
	Breakdown      *QuoteComputedRecurringTotalDetailsBreakdown `json:"breakdown"`
}

// The definitive totals and line items the customer will be charged on a recurring basis. Takes into account the line items with recurring prices and discounts with `duration=forever` coupons only. Defaults to `null` if no inputted line items with recurring prices.
type QuoteComputedRecurring struct {
	AmountSubtotal int64                               `json:"amount_subtotal"`
	AmountTotal    int64                               `json:"amount_total"`
	Interval       QuoteComputedRecurringInterval      `json:"interval"`
	IntervalCount  int64                               `json:"interval_count"`
	TotalDetails   *QuoteComputedRecurringTotalDetails `json:"total_details"`
}

// The aggregated line item discounts.
type QuoteComputedUpfrontTotalDetailsBreakdownDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// The aggregated line item tax amounts by rate.
type QuoteComputedUpfrontTotalDetailsBreakdownTax struct {
	Amount int64    `json:"amount"`
	Rate   *TaxRate `json:"rate"`
}
type QuoteComputedUpfrontTotalDetailsBreakdown struct {
	Discounts []*QuoteComputedUpfrontTotalDetailsBreakdownDiscount `json:"discounts"`
	Taxes     []*QuoteComputedUpfrontTotalDetailsBreakdownTax      `json:"taxes"`
}
type QuoteComputedUpfrontTotalDetails struct {
	AmountDiscount int64                                      `json:"amount_discount"`
	AmountShipping int64                                      `json:"amount_shipping"`
	AmountTax      int64                                      `json:"amount_tax"`
	Breakdown      *QuoteComputedUpfrontTotalDetailsBreakdown `json:"breakdown"`
}
type QuoteComputedUpfront struct {
	AmountSubtotal int64                             `json:"amount_subtotal"`
	AmountTotal    int64                             `json:"amount_total"`
	LineItems      *LineItemList                     `json:"line_items"`
	TotalDetails   *QuoteComputedUpfrontTotalDetails `json:"total_details"`
}
type QuoteComputed struct {
	Recurring *QuoteComputedRecurring `json:"recurring"`
	Upfront   *QuoteComputedUpfront   `json:"upfront"`
}

// Details of the quote that was cloned. See the [cloning documentation](https://stripe.com/docs/quotes/clone) for more details.
type QuoteFromQuote struct {
	IsRevision bool   `json:"is_revision"`
	Quote      *Quote `json:"quote"`
}

// All invoices will be billed using the specified settings.
type QuoteInvoiceSettings struct {
	DaysUntilDue int64 `json:"days_until_due"`
}
type QuoteStatusTransitions struct {
	AcceptedAt  int64 `json:"accepted_at"`
	CanceledAt  int64 `json:"canceled_at"`
	FinalizedAt int64 `json:"finalized_at"`
}
type QuoteSubscriptionData struct {
	EffectiveDate   int64 `json:"effective_date"`
	TrialPeriodDays int64 `json:"trial_period_days"`
}

// The aggregated line item discounts.
type QuoteTotalDetailsBreakdownDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// The aggregated line item tax amounts by rate.
type QuoteTotalDetailsBreakdownTax struct {
	Amount int64    `json:"amount"`
	Rate   *TaxRate `json:"rate"`
}
type QuoteTotalDetailsBreakdown struct {
	Discounts []*QuoteTotalDetailsBreakdownDiscount `json:"discounts"`
	Taxes     []*QuoteTotalDetailsBreakdownTax      `json:"taxes"`
}
type QuoteTotalDetails struct {
	AmountDiscount int64                       `json:"amount_discount"`
	AmountShipping int64                       `json:"amount_shipping"`
	AmountTax      int64                       `json:"amount_tax"`
	Breakdown      *QuoteTotalDetailsBreakdown `json:"breakdown"`
}

// The account (if any) the payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the invoices.
type QuoteTransferData struct {
	Amount        int64    `json:"amount"`
	AmountPercent float64  `json:"amount_percent"`
	Destination   *Account `json:"destination"`
}

// A Quote is a way to model prices that you'd like to provide to a customer.
// Once accepted, it will automatically create an invoice, subscription or subscription schedule.
type Quote struct {
	APIResource
	AmountSubtotal        int64                   `json:"amount_subtotal"`
	AmountTotal           int64                   `json:"amount_total"`
	ApplicationFeeAmount  int64                   `json:"application_fee_amount"`
	ApplicationFeePercent float64                 `json:"application_fee_percent"`
	AutomaticTax          *QuoteAutomaticTax      `json:"automatic_tax"`
	CollectionMethod      QuoteCollectionMethod   `json:"collection_method"`
	Computed              *QuoteComputed          `json:"computed"`
	Created               int64                   `json:"created"`
	Currency              Currency                `json:"currency"`
	Customer              *Customer               `json:"customer"`
	DefaultTaxRates       []*TaxRate              `json:"default_tax_rates"`
	Description           string                  `json:"description"`
	Discounts             []*Discount             `json:"discounts"`
	ExpiresAt             int64                   `json:"expires_at"`
	Footer                string                  `json:"footer"`
	FromQuote             *QuoteFromQuote         `json:"from_quote"`
	Header                string                  `json:"header"`
	ID                    string                  `json:"id"`
	Invoice               *Invoice                `json:"invoice"`
	InvoiceSettings       *QuoteInvoiceSettings   `json:"invoice_settings"`
	LineItems             *LineItemList           `json:"line_items"`
	Livemode              bool                    `json:"livemode"`
	Metadata              map[string]string       `json:"metadata"`
	Number                string                  `json:"number"`
	Object                string                  `json:"object"`
	OnBehalfOf            *Account                `json:"on_behalf_of"`
	Status                QuoteStatus             `json:"status"`
	StatusTransitions     *QuoteStatusTransitions `json:"status_transitions"`
	Subscription          *Subscription           `json:"subscription"`
	SubscriptionData      *QuoteSubscriptionData  `json:"subscription_data"`
	SubscriptionSchedule  *SubscriptionSchedule   `json:"subscription_schedule"`
	TotalDetails          *QuoteTotalDetails      `json:"total_details"`
	TransferData          *QuoteTransferData      `json:"transfer_data"`
}

// QuoteList is a list of Quotes as retrieved from a list endpoint.
type QuoteList struct {
	APIResource
	ListMeta
	Data []*Quote `json:"data"`
}

// UnmarshalJSON handles deserialization of a Quote.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (q *Quote) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		q.ID = id
		return nil
	}

	type quote Quote
	var v quote
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*q = Quote(v)
	return nil
}
