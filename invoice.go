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

type InvoiceAutomaticTaxStatus string

// List of values that InvoiceAutomaticTaxStatus can take
const (
	InvoiceAutomaticTaxStatusComplete               InvoiceAutomaticTaxStatus = "complete"
	InvoiceAutomaticTaxStatusFailed                 InvoiceAutomaticTaxStatus = "failed"
	InvoiceAutomaticTaxStatusRequiresLocationInputs InvoiceAutomaticTaxStatus = "requires_location_inputs"
)

// InvoiceBillingReason is the reason why a given invoice was created
type InvoiceBillingReason string

// List of values that InvoiceBillingReason can take.
const (
	InvoiceBillingReasonAutomaticPendingInvoiceItemInvoice InvoiceBillingReason = "automatic_pending_invoice_item_invoice"
	InvoiceBillingReasonManual                             InvoiceBillingReason = "manual"
	InvoiceBillingReasonQuoteAccept                        InvoiceBillingReason = "quote_accept"
	InvoiceBillingReasonSubscription                       InvoiceBillingReason = "subscription"
	InvoiceBillingReasonSubscriptionCreate                 InvoiceBillingReason = "subscription_create"
	InvoiceBillingReasonSubscriptionCycle                  InvoiceBillingReason = "subscription_cycle"
	InvoiceBillingReasonSubscriptionThreshold              InvoiceBillingReason = "subscription_threshold"
	InvoiceBillingReasonSubscriptionUpdate                 InvoiceBillingReason = "subscription_update"
	InvoiceBillingReasonUpcoming                           InvoiceBillingReason = "upcoming"
)

// Transaction type of the mandate.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Bank account verification method.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodInstant       InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure represents
// the options for requesting 3D Secure.
type InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure can take.
const (
	InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAny       InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "any"
	InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecureAutomatic InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
)

// InvoicePaymentSettingsPaymentMethodType represents the payment method type to provide to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodType string

// List of values that InvoicePaymentSettingsPaymentMethodType can take.
const (
	InvoicePaymentSettingsPaymentMethodTypeAchCreditTransfer  InvoicePaymentSettingsPaymentMethodType = "ach_credit_transfer"
	InvoicePaymentSettingsPaymentMethodTypeAchDebit           InvoicePaymentSettingsPaymentMethodType = "ach_debit"
	InvoicePaymentSettingsPaymentMethodTypeACSSDebit          InvoicePaymentSettingsPaymentMethodType = "acss_debit"
	InvoicePaymentSettingsPaymentMethodTypeAUBECSDebit        InvoicePaymentSettingsPaymentMethodType = "au_becs_debit"
	InvoicePaymentSettingsPaymentMethodTypeBACSDebit          InvoicePaymentSettingsPaymentMethodType = "bacs_debit"
	InvoicePaymentSettingsPaymentMethodTypeBancontact         InvoicePaymentSettingsPaymentMethodType = "bancontact"
	InvoicePaymentSettingsPaymentMethodTypeBoleto             InvoicePaymentSettingsPaymentMethodType = "boleto"
	InvoicePaymentSettingsPaymentMethodTypeCard               InvoicePaymentSettingsPaymentMethodType = "card"
	InvoicePaymentSettingsPaymentMethodTypeFPX                InvoicePaymentSettingsPaymentMethodType = "fpx"
	InvoicePaymentSettingsPaymentMethodTypeGiropay            InvoicePaymentSettingsPaymentMethodType = "giropay"
	InvoicePaymentSettingsPaymentMethodTypeIdeal              InvoicePaymentSettingsPaymentMethodType = "ideal"
	InvoicePaymentSettingsPaymentMethodTypeSepaCreditTransfer InvoicePaymentSettingsPaymentMethodType = "sepa_credit_transfer"
	InvoicePaymentSettingsPaymentMethodTypeSepaDebit          InvoicePaymentSettingsPaymentMethodType = "sepa_debit"
	InvoicePaymentSettingsPaymentMethodTypeSofort             InvoicePaymentSettingsPaymentMethodType = "sofort"
	InvoicePaymentSettingsPaymentMethodTypeWechatPay          InvoicePaymentSettingsPaymentMethodType = "wechat_pay"
)

// InvoiceStatus is the reason why a given invoice was created
type InvoiceStatus string

// List of values that InvoiceStatus can take.
const (
	InvoiceStatusDeleted       InvoiceStatus = "deleted"
	InvoiceStatusDraft         InvoiceStatus = "draft"
	InvoiceStatusOpen          InvoiceStatus = "open"
	InvoiceStatusPaid          InvoiceStatus = "paid"
	InvoiceStatusUncollectible InvoiceStatus = "uncollectible"
	InvoiceStatusVoid          InvoiceStatus = "void"
)

// InvoiceCollectionMethod is the type of collection method for this invoice.
type InvoiceCollectionMethod string

// List of values that InvoiceCollectionMethod can take.
const (
	InvoiceCollectionMethodChargeAutomatically InvoiceCollectionMethod = "charge_automatically"
	InvoiceCollectionMethodSendInvoice         InvoiceCollectionMethod = "send_invoice"
)

type InvoiceUpcomingAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// The customer's shipping information. Appears on invoices emailed to this customer.
type InvoiceUpcomingCustomerDetailsShippingParams struct {
	Address *AddressParams `form:"address"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// Tax details about the customer.
type InvoiceUpcomingCustomerDetailsTaxParams struct {
	IPAddress *string `form:"ip_address"`
}

// The customer's tax IDs.
type InvoiceUpcomingCustomerDetailsTaxIDParams struct {
	Type  *string `form:"type"`
	Value *string `form:"value"`
}

// Details about the customer you want to invoice or overrides for an existing customer.
type InvoiceUpcomingCustomerDetailsParams struct {
	Address   *AddressParams                                `form:"address"`
	Shipping  *InvoiceUpcomingCustomerDetailsShippingParams `form:"shipping"`
	Tax       *InvoiceUpcomingCustomerDetailsTaxParams      `form:"tax"`
	TaxExempt *string                                       `form:"tax_exempt"`
	TaxIDs    []*InvoiceUpcomingCustomerDetailsTaxIDParams  `form:"tax_ids"`
}

// InvoiceUpcomingInvoiceItemPeriodParams represents the period associated with that invoice item
type InvoiceUpcomingInvoiceItemPeriodParams struct {
	End   *int64 `form:"end"`
	Start *int64 `form:"start"`
}

// InvoiceUpcomingInvoiceItemParams is the set of parameters that can be used when adding or modifying
// invoice items on an upcoming invoice.
// For more details see https://stripe.com/docs/api#upcoming_invoice-invoice_items.
type InvoiceUpcomingInvoiceItemParams struct {
	Amount            *int64                                  `form:"amount"`
	Currency          *string                                 `form:"currency"`
	Description       *string                                 `form:"description"`
	Discountable      *bool                                   `form:"discountable"`
	Discounts         []*InvoiceItemDiscountParams            `form:"discounts"`
	InvoiceItem       *string                                 `form:"invoiceitem"`
	Metadata          map[string]string                       `form:"metadata"`
	Period            *InvoiceUpcomingInvoiceItemPeriodParams `form:"period"`
	Price             *string                                 `form:"price"`
	PriceData         *InvoiceItemPriceDataParams             `form:"price_data"`
	Quantity          *int64                                  `form:"quantity"`
	Schedule          *string                                 `form:"schedule"`
	TaxRates          []*string                               `form:"tax_rates"`
	UnitAmount        *int64                                  `form:"unit_amount"`
	UnitAmountDecimal *float64                                `form:"unit_amount_decimal,high_precision"`
}

// Settings for automatic tax lookup for this invoice.
type InvoiceAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// InvoiceCustomFieldParams represents the parameters associated with one custom field on an invoice.
type InvoiceCustomFieldParams struct {
	Name  *string `form:"name"`
	Value *string `form:"value"`
}

// InvoiceDiscountParams represents the parameters associated with the discounts to apply to an invoice.
type InvoiceDiscountParams struct {
	Coupon   *string `form:"coupon"`
	Discount *string `form:"discount"`
}

// Additional fields for Mandate creation
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	TransactionType *string `form:"transaction_type"`
}

// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitParams struct {
	MandateOptions     *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                                  `form:"verification_method"`
}

// InvoicePaymentSettingsPaymentMethodOptionsBancontactParams is the set of parameters allowed for
// bancontact on payment_method_options on payment_settings on an invoice.
type InvoicePaymentSettingsPaymentMethodOptionsBancontactParams struct {
	PreferredLanguage *string `form:"preferred_language"`
}

// InvoicePaymentSettingsPaymentMethodOptionsCardParams is the set of parameters allowed for
// payment method options when using the card payment method.
type InvoicePaymentSettingsPaymentMethodOptionsCardParams struct {
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// InvoicePaymentSettingsPaymentMethodOptionsParams is the set of parameters allowed for
// specific payment methods on an invoice's payment settings.
type InvoicePaymentSettingsPaymentMethodOptionsParams struct {
	ACSSDebit  *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitParams  `form:"acss_debit"`
	Bancontact *InvoicePaymentSettingsPaymentMethodOptionsBancontactParams `form:"bancontact"`
	Card       *InvoicePaymentSettingsPaymentMethodOptionsCardParams       `form:"card"`
}

// InvoicePaymentSettingsParams is the set of parameters allowed for the payment_settings on an invoice.
type InvoicePaymentSettingsParams struct {
	PaymentMethodOptions *InvoicePaymentSettingsPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes   []*string                                         `form:"payment_method_types"`
}

// InvoiceTransferDataParams is the set of parameters allowed for the transfer_data hash.
type InvoiceTransferDataParams struct {
	Amount      *int64  `form:"amount"`
	Destination *string `form:"destination"`
}

// InvoiceParams is the set of parameters that can be used when creating or updating an invoice.
// For more details see https://stripe.com/docs/api#create_invoice, https://stripe.com/docs/api#update_invoice.
type InvoiceParams struct {
	Params               `form:"*"`
	AccountTaxIDs        []*string                     `form:"account_tax_ids"`
	ApplicationFeeAmount *int64                        `form:"application_fee_amount"`
	AutoAdvance          *bool                         `form:"auto_advance"`
	AutomaticTax         *InvoiceAutomaticTaxParams    `form:"automatic_tax"`
	CollectionMethod     *string                       `form:"collection_method"`
	Customer             *string                       `form:"customer"`
	CustomFields         []*InvoiceCustomFieldParams   `form:"custom_fields"`
	DaysUntilDue         *int64                        `form:"days_until_due"`
	DefaultPaymentMethod *string                       `form:"default_payment_method"`
	DefaultSource        *string                       `form:"default_source"`
	DefaultTaxRates      []*string                     `form:"default_tax_rates"`
	Description          *string                       `form:"description"`
	Discounts            []*InvoiceDiscountParams      `form:"discounts"`
	DueDate              *int64                        `form:"due_date"`
	Footer               *string                       `form:"footer"`
	OnBehalfOf           *string                       `form:"on_behalf_of"`
	Paid                 *bool                         `form:"paid"`
	PaymentSettings      *InvoicePaymentSettingsParams `form:"payment_settings"`
	Schedule             *string                       `form:"schedule"`
	StatementDescriptor  *string                       `form:"statement_descriptor"`
	Subscription         *string                       `form:"subscription"`
	TransferData         *InvoiceTransferDataParams    `form:"transfer_data"`
	// These are all for exclusive use by GetNext.
	Coupon                                  *string                               `form:"coupon"`
	CustomerDetails                         *InvoiceUpcomingCustomerDetailsParams `form:"customer_details"`
	InvoiceItems                            []*InvoiceUpcomingInvoiceItemParams   `form:"invoice_items"`
	SubscriptionBillingCycleAnchor          *int64                                `form:"subscription_billing_cycle_anchor"`
	SubscriptionBillingCycleAnchorNow       *bool                                 `form:"-"` // See custom AppendTo
	SubscriptionBillingCycleAnchorUnchanged *bool                                 `form:"-"` // See custom AppendTo
	SubscriptionCancelAt                    *int64                                `form:"subscription_cancel_at"`
	SubscriptionCancelAtPeriodEnd           *bool                                 `form:"subscription_cancel_at_period_end"`
	SubscriptionCancelNow                   *bool                                 `form:"subscription_cancel_now"`
	SubscriptionDefaultTaxRates             []*string                             `form:"subscription_default_tax_rates"`
	SubscriptionItems                       []*SubscriptionItemsParams            `form:"subscription_items"`
	SubscriptionPlan                        *string                               `form:"subscription_plan"`
	SubscriptionProrationBehavior           *string                               `form:"subscription_proration_behavior"`
	SubscriptionProrationDate               *int64                                `form:"subscription_proration_date"`
	SubscriptionQuantity                    *int64                                `form:"subscription_quantity"`
	SubscriptionStartDate                   *int64                                `form:"subscription_start_date"`
	SubscriptionTrialEnd                    *int64                                `form:"subscription_trial_end"`
	SubscriptionTrialEndNow                 *bool                                 `form:"-"` // See custom AppendTo
	SubscriptionTrialFromPlan               *bool                                 `form:"subscription_trial_from_plan"`
}

// AppendTo implements custom encoding logic for InvoiceParams.
func (i *InvoiceParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(i.SubscriptionBillingCycleAnchorNow) {
		body.Add(form.FormatKey(append(keyParts, "subscription_billing_cycle_anchor")), "now")
	}
	if BoolValue(i.SubscriptionBillingCycleAnchorUnchanged) {
		body.Add(form.FormatKey(append(keyParts, "subscription_billing_cycle_anchor")), "unchanged")
	}
	if BoolValue(i.SubscriptionTrialEndNow) {
		body.Add(form.FormatKey(append(keyParts, "subscription_trial_end")), "now")
	}
}

// InvoiceListParams is the set of parameters that can be used when listing invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
type InvoiceListParams struct {
	ListParams       `form:"*"`
	CollectionMethod *string           `form:"collection_method"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Customer         *string           `form:"customer"`
	DueDate          *int64            `form:"due_date"`
	DueDateRange     *RangeQueryParams `form:"due_date"`
	Status           *string           `form:"status"`
	Subscription     *string           `form:"subscription"`
}

// InvoiceLineListParams is the set of parameters that can be used when listing invoice line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
type InvoiceLineListParams struct {
	ListParams `form:"*"`
	// ID is the invoice ID to list invoice lines for.
	ID           *string `form:"-"` // Included in URL
	Customer     *string `form:"customer"`
	Subscription *string `form:"subscription"`
}

// InvoiceFinalizeParams is the set of parameters that can be used when finalizing invoices.
type InvoiceFinalizeParams struct {
	Params      `form:"*"`
	AutoAdvance *bool `form:"auto_advance"`
}

// InvoiceMarkUncollectibleParams is the set of parameters that can be used when marking
// invoices as uncollectible.
type InvoiceMarkUncollectibleParams struct {
	Params `form:"*"`
}

// InvoicePayParams is the set of parameters that can be used when
// paying invoices. For more details, see:
// https://stripe.com/docs/api#pay_invoice.
type InvoicePayParams struct {
	Params        `form:"*"`
	Forgive       *bool   `form:"forgive"`
	OffSession    *bool   `form:"off_session"`
	PaidOutOfBand *bool   `form:"paid_out_of_band"`
	PaymentMethod *string `form:"payment_method"`
	Source        *string `form:"source"`
}

// InvoiceSendParams is the set of parameters that can be used when sending invoices.
type InvoiceSendParams struct {
	Params `form:"*"`
}

// InvoiceVoidParams is the set of parameters that can be used when voiding invoices.
type InvoiceVoidParams struct {
	Params `form:"*"`
}

// Invoice is the resource representing a Stripe invoice.
// For more details see https://stripe.com/docs/api#invoice_object.
type Invoice struct {
	APIResource
	AccountCountry               string                   `json:"account_country"`
	AccountName                  string                   `json:"account_name"`
	AccountTaxIDs                []*TaxID                 `json:"account_tax_ids"`
	AmountDue                    int64                    `json:"amount_due"`
	AmountPaid                   int64                    `json:"amount_paid"`
	AmountRemaining              int64                    `json:"amount_remaining"`
	ApplicationFeeAmount         int64                    `json:"application_fee_amount"`
	AttemptCount                 int64                    `json:"attempt_count"`
	Attempted                    bool                     `json:"attempted"`
	AutoAdvance                  bool                     `json:"auto_advance"`
	AutomaticTax                 *InvoiceAutomaticTax     `json:"automatic_tax"`
	BillingReason                InvoiceBillingReason     `json:"billing_reason"`
	Charge                       *Charge                  `json:"charge"`
	CollectionMethod             *InvoiceCollectionMethod `json:"collection_method"`
	Created                      int64                    `json:"created"`
	Currency                     Currency                 `json:"currency"`
	Customer                     *Customer                `json:"customer"`
	CustomerAddress              *Address                 `json:"customer_address"`
	CustomerEmail                string                   `json:"customer_email"`
	CustomerName                 *string                  `json:"customer_name"`
	CustomerPhone                *string                  `json:"customer_phone"`
	CustomerShipping             *CustomerShippingDetails `json:"customer_shipping"`
	CustomerTaxExempt            CustomerTaxExempt        `json:"customer_tax_exempt"`
	CustomerTaxIDs               []*InvoiceCustomerTaxID  `json:"customer_tax_ids"`
	CustomFields                 []*InvoiceCustomField    `json:"custom_fields"`
	DefaultPaymentMethod         *PaymentMethod           `json:"default_payment_method"`
	DefaultSource                *PaymentSource           `json:"default_source"`
	DefaultTaxRates              []*TaxRate               `json:"default_tax_rates"`
	Deleted                      bool                     `json:"deleted"`
	Description                  string                   `json:"description"`
	Discount                     *Discount                `json:"discount"`
	Discounts                    []*Discount              `json:"discounts"`
	DueDate                      int64                    `json:"due_date"`
	EndingBalance                int64                    `json:"ending_balance"`
	Footer                       string                   `json:"footer"`
	HostedInvoiceURL             string                   `json:"hosted_invoice_url"`
	ID                           string                   `json:"id"`
	InvoicePDF                   string                   `json:"invoice_pdf"`
	LastFinalizationError        *Error                   `json:"last_finalization_error"`
	Lines                        *InvoiceLineList         `json:"lines"`
	Livemode                     bool                     `json:"livemode"`
	Metadata                     map[string]string        `json:"metadata"`
	NextPaymentAttempt           int64                    `json:"next_payment_attempt"`
	Number                       string                   `json:"number"`
	Object                       string                   `json:"object"`
	OnBehalfOf                   *Account                 `json:"on_behalf_of"`
	Paid                         bool                     `json:"paid"`
	PaidOutOfBand                bool                     `json:"paid_out_of_band"`
	PaymentIntent                *PaymentIntent           `json:"payment_intent"`
	PaymentSettings              *InvoicePaymentSettings  `json:"payment_settings"`
	PeriodEnd                    int64                    `json:"period_end"`
	PeriodStart                  int64                    `json:"period_start"`
	PostPaymentCreditNotesAmount int64                    `json:"post_payment_credit_notes_amount"`
	PrePaymentCreditNotesAmount  int64                    `json:"pre_payment_credit_notes_amount"`
	Quote                        *Quote                   `json:"quote"`
	ReceiptNumber                string                   `json:"receipt_number"`
	StartingBalance              int64                    `json:"starting_balance"`
	StatementDescriptor          string                   `json:"statement_descriptor"`
	Status                       InvoiceStatus            `json:"status"`
	StatusTransitions            InvoiceStatusTransitions `json:"status_transitions"`
	Subscription                 *Subscription            `json:"subscription"`
	SubscriptionProrationDate    int64                    `json:"subscription_proration_date"`
	Subtotal                     int64                    `json:"subtotal"`
	Tax                          int64                    `json:"tax"`
	ThreasholdReason             *InvoiceThresholdReason  `json:"threshold_reason"`
	Total                        int64                    `json:"total"`
	TotalDiscountAmounts         []*InvoiceDiscountAmount `json:"total_discount_amounts"`
	TotalTaxAmounts              []*InvoiceTaxAmount      `json:"total_tax_amounts"`
	TransferData                 *InvoiceTransferData     `json:"transfer_data"`
	WebhooksDeliveredAt          int64                    `json:"webhooks_delivered_at"`
}
type InvoiceAutomaticTax struct {
	Enabled bool                      `json:"enabled"`
	Status  InvoiceAutomaticTaxStatus `json:"status"`
}

// InvoiceCustomField is a structure representing a custom field on an invoice.
type InvoiceCustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// InvoiceCustomerTaxID is a structure representing a customer tax id on an invoice.
type InvoiceCustomerTaxID struct {
	Type  TaxIDType `json:"type"`
	Value string    `json:"value"`
}

// InvoiceDiscountAmount represents the aggregate amounts calculated per discount for all line items.
type InvoiceDiscountAmount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// InvoiceTaxAmount is a structure representing one of the tax amounts on an invoice.
type InvoiceTaxAmount struct {
	Amount    int64    `json:"amount"`
	Inclusive bool     `json:"inclusive"`
	TaxRate   *TaxRate `json:"tax_rate"`
}

// InvoiceThresholdReason is a structure representing a reason for a billing threshold.
type InvoiceThresholdReason struct {
	AmountGTE   int64                               `json:"amount_gte"`
	ItemReasons []*InvoiceThresholdReasonItemReason `json:"item_reasons"`
}

// InvoiceThresholdReasonItemReason is a structure representing the line items that
// triggered an invoice.
type InvoiceThresholdReasonItemReason struct {
	LineItemIDs []string `json:"line_item_ids"`
	UsageGTE    int64    `json:"usage_gte"`
}

// InvoiceTransferData represents the information for the transfer_data associated with an invoice.
type InvoiceTransferData struct {
	Amount      int64    `json:"amount"`
	Destination *Account `json:"destination"`
}
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions struct {
	TransactionType InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// If paying by `acss_debit`, this sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsACSSDebit struct {
	MandateOptions     *InvoicePaymentSettingsPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod InvoicePaymentSettingsPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// InvoicePaymentSettingsPaymentMethodOptionsBancontact contains details about the Bancontact
// payment method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsBancontact struct {
	PreferredLanguage string `json:"preferred_language"`
}

// InvoicePaymentSettingsPaymentMethodOptionsCard contains details about the Card payment
// method options to pass to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptionsCard struct {
	RequestThreeDSecure InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}

// InvoicePaymentSettingsPaymentMethodOptions represents payment-method-specific configuration to provide to the invoice's PaymentIntent.
type InvoicePaymentSettingsPaymentMethodOptions struct {
	ACSSDebit  *InvoicePaymentSettingsPaymentMethodOptionsACSSDebit  `json:"acss_debit"`
	Bancontact *InvoicePaymentSettingsPaymentMethodOptionsBancontact `json:"bancontact"`
	Card       *InvoicePaymentSettingsPaymentMethodOptionsCard       `json:"card"`
}

// InvoicePaymentSettings represents configuration settings to provide to the invoice's PaymentIntent.
type InvoicePaymentSettings struct {
	PaymentMethodOptions *InvoicePaymentSettingsPaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes   []InvoicePaymentSettingsPaymentMethodType   `json:"payment_method_types"`
}

// InvoiceStatusTransitions are the timestamps at which the invoice status was updated.
type InvoiceStatusTransitions struct {
	FinalizedAt           int64 `json:"finalized_at"`
	MarkedUncollectibleAt int64 `json:"marked_uncollectible_at"`
	PaidAt                int64 `json:"paid_at"`
	VoidedAt              int64 `json:"voided_at"`
}

// InvoiceList is a list of invoices as retrieved from a list endpoint.
type InvoiceList struct {
	APIResource
	ListMeta
	Data []*Invoice `json:"data"`
}

// UnmarshalJSON handles deserialization of an Invoice.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *Invoice) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type invoice Invoice
	var v invoice
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = Invoice(v)
	return nil
}
