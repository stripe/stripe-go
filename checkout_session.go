package stripe

import (
	"encoding/json"
)

// CheckoutSessionCustomerDetailsTaxIDsType is the list of allowed values for type
// on the tax_ids inside customer_details of a checkout session.
type CheckoutSessionCustomerDetailsTaxIDsType string

// List of values that CheckoutSessionCustomerDetailsTaxIDsType can take.
const (
	CheckoutSessionCustomerDetailsTaxIDsTypeAETRN   CheckoutSessionCustomerDetailsTaxIDsType = "ae_trn"
	CheckoutSessionCustomerDetailsTaxIDsTypeAUABN   CheckoutSessionCustomerDetailsTaxIDsType = "au_abn"
	CheckoutSessionCustomerDetailsTaxIDsTypeBRCNPJ  CheckoutSessionCustomerDetailsTaxIDsType = "br_cnpj"
	CheckoutSessionCustomerDetailsTaxIDsTypeBRCPF   CheckoutSessionCustomerDetailsTaxIDsType = "br_cpf"
	CheckoutSessionCustomerDetailsTaxIDsTypeCABN    CheckoutSessionCustomerDetailsTaxIDsType = "ca_bn"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAQST   CheckoutSessionCustomerDetailsTaxIDsType = "ca_qst"
	CheckoutSessionCustomerDetailsTaxIDsTypeCHVAT   CheckoutSessionCustomerDetailsTaxIDsType = "ch_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeCLTIN   CheckoutSessionCustomerDetailsTaxIDsType = "cl_tin"
	CheckoutSessionCustomerDetailsTaxIDsTypeESCIF   CheckoutSessionCustomerDetailsTaxIDsType = "es_cif"
	CheckoutSessionCustomerDetailsTaxIDsTypeEUVAT   CheckoutSessionCustomerDetailsTaxIDsType = "eu_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeGBVAT   CheckoutSessionCustomerDetailsTaxIDsType = "gb_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeHKBR    CheckoutSessionCustomerDetailsTaxIDsType = "hk_br"
	CheckoutSessionCustomerDetailsTaxIDsTypeIDNPWP  CheckoutSessionCustomerDetailsTaxIDsType = "id_npwp"
	CheckoutSessionCustomerDetailsTaxIDsTypeINGST   CheckoutSessionCustomerDetailsTaxIDsType = "in_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeJPCN    CheckoutSessionCustomerDetailsTaxIDsType = "jp_cn"
	CheckoutSessionCustomerDetailsTaxIDsTypeJPRN    CheckoutSessionCustomerDetailsTaxIDsType = "jp_rn"
	CheckoutSessionCustomerDetailsTaxIDsTypeKRBRN   CheckoutSessionCustomerDetailsTaxIDsType = "kr_brn"
	CheckoutSessionCustomerDetailsTaxIDsTypeLIUID   CheckoutSessionCustomerDetailsTaxIDsType = "li_uid"
	CheckoutSessionCustomerDetailsTaxIDsTypeMXRFC   CheckoutSessionCustomerDetailsTaxIDsType = "mx_rfc"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYFRP   CheckoutSessionCustomerDetailsTaxIDsType = "my_frp"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYITN   CheckoutSessionCustomerDetailsTaxIDsType = "my_itn"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYSST   CheckoutSessionCustomerDetailsTaxIDsType = "my_sst"
	CheckoutSessionCustomerDetailsTaxIDsTypeNOVAT   CheckoutSessionCustomerDetailsTaxIDsType = "no_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeNZGST   CheckoutSessionCustomerDetailsTaxIDsType = "nz_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeRUINN   CheckoutSessionCustomerDetailsTaxIDsType = "ru_inn"
	CheckoutSessionCustomerDetailsTaxIDsTypeRUKPP   CheckoutSessionCustomerDetailsTaxIDsType = "ru_kpp"
	CheckoutSessionCustomerDetailsTaxIDsTypeSAVAT   CheckoutSessionCustomerDetailsTaxIDsType = "sa_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeSGGST   CheckoutSessionCustomerDetailsTaxIDsType = "sg_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeSGUEN   CheckoutSessionCustomerDetailsTaxIDsType = "sg_uen"
	CheckoutSessionCustomerDetailsTaxIDsTypeTHVAT   CheckoutSessionCustomerDetailsTaxIDsType = "th_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeTWVAT   CheckoutSessionCustomerDetailsTaxIDsType = "tw_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeUnknown CheckoutSessionCustomerDetailsTaxIDsType = "unknown"
	CheckoutSessionCustomerDetailsTaxIDsTypeUSEIN   CheckoutSessionCustomerDetailsTaxIDsType = "us_ein"
	CheckoutSessionCustomerDetailsTaxIDsTypeZAVAT   CheckoutSessionCustomerDetailsTaxIDsType = "za_vat"
)

// CheckoutSessionCustomerDetailsTaxExempt is the list of allowed values for
// tax_exempt inside customer_details of a checkout session.
type CheckoutSessionCustomerDetailsTaxExempt string

// List of values that CheckoutSessionCustomerDetailsTaxExempt can take.
const (
	CheckoutSessionCustomerDetailsTaxExemptExempt  CheckoutSessionCustomerDetailsTaxExempt = "exempt"
	CheckoutSessionCustomerDetailsTaxExemptNone    CheckoutSessionCustomerDetailsTaxExempt = "none"
	CheckoutSessionCustomerDetailsTaxExemptReverse CheckoutSessionCustomerDetailsTaxExempt = "reverse"
)

// CheckoutSessionMode is the list of allowed values for the mode on a Session.
type CheckoutSessionMode string

// List of values that CheckoutSessionMode can take.
const (
	CheckoutSessionModePayment      CheckoutSessionMode = "payment"
	CheckoutSessionModeSetup        CheckoutSessionMode = "setup"
	CheckoutSessionModeSubscription CheckoutSessionMode = "subscription"
)

// CheckoutSessionPaymentStatus is the list of allowed values for the payment status on a Session.`
type CheckoutSessionPaymentStatus string

// List of values that CheckoutSessionPaymentStatus can take.
const (
	CheckoutSessionPaymentStatusNoPaymentRequired CheckoutSessionPaymentStatus = "no_payment_required"
	CheckoutSessionPaymentStatusPaid              CheckoutSessionPaymentStatus = "paid"
	CheckoutSessionPaymentStatusUnpaid            CheckoutSessionPaymentStatus = "unpaid"
)

// CheckoutSessionSubmitType is the list of allowed values for the submit type on  a Session.
type CheckoutSessionSubmitType string

// List of values that CheckoutSessionSubmitType can take.
const (
	CheckoutSessionSubmitTypeAuto   CheckoutSessionSubmitType = "auto"
	CheckoutSessionSubmitTypeBook   CheckoutSessionSubmitType = "book"
	CheckoutSessionSubmitTypeDonate CheckoutSessionSubmitType = "donate"
	CheckoutSessionSubmitTypePay    CheckoutSessionSubmitType = "pay"
)

// CheckoutSessionLineItemAdjustableQuantityParams represents the parameters for
// an adjustable quantity on a checkout session's line items.
type CheckoutSessionLineItemAdjustableQuantityParams struct {
	Enabled *bool  `form:"enabled"`
	Maximum *int64 `form:"maximum"`
	Minimum *int64 `form:"minimum"`
}

// CheckoutSessionLineItemPriceDataProductDataParams is the set of parameters that can be used when
// creating a product created inline for a line item.
type CheckoutSessionLineItemPriceDataProductDataParams struct {
	Description *string           `form:"description"`
	Images      []*string         `form:"images"`
	Metadata    map[string]string `form:"metadata"`
	Name        *string           `form:"name"`
}

// CheckoutSessionLineItemPriceDataRecurringParams is the set of parameters for the recurring
// components of a price created inline for a line item.
type CheckoutSessionLineItemPriceDataRecurringParams struct {
	AggregateUsage  *string `form:"aggregate_usage"`
	Interval        *string `form:"interval"`
	IntervalCount   *int64  `form:"interval_count"`
	TrialPeriodDays *int64  `form:"trial_period_days"`
	UsageType       *string `form:"usage_type"`
}

// CheckoutSessionLineItemPriceDataParams is a structure representing the parameters to create
// an inline price for a line item.
type CheckoutSessionLineItemPriceDataParams struct {
	Currency          *string                                            `form:"currency"`
	Product           *string                                            `form:"product"`
	ProductData       *CheckoutSessionLineItemPriceDataProductDataParams `form:"product_data"`
	Recurring         *CheckoutSessionLineItemPriceDataRecurringParams   `form:"recurring"`
	UnitAmount        *int64                                             `form:"unit_amount"`
	UnitAmountDecimal *float64                                           `form:"unit_amount_decimal,high_precision"`
}

// CheckoutSessionDiscountParams is the set of parameters allowed for discounts on
// a checkout session.
type CheckoutSessionDiscountParams struct {
	Coupon        *string `form:"coupon"`
	PromotionCode *string `form:"promotion_code"`
}

// CheckoutSessionLineItemParams is the set of parameters allowed for a line item
// on a checkout session.
type CheckoutSessionLineItemParams struct {
	AdjustableQuantity *CheckoutSessionLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	Amount             *int64                                           `form:"amount"`
	Currency           *string                                          `form:"currency"`
	Description        *string                                          `form:"description"`
	DynamicTaxRates    []*string                                        `form:"dynamic_tax_rates"`
	Images             []*string                                        `form:"images"`
	Name               *string                                          `form:"name"`
	Price              *string                                          `form:"price"`
	PriceData          *CheckoutSessionLineItemPriceDataParams          `form:"price_data"`
	Quantity           *int64                                           `form:"quantity"`
	TaxRates           []*string                                        `form:"tax_rates"`
}

// CheckoutSessionPaymentIntentDataTransferDataParams is the set of parameters allowed for the
// transfer_data hash.
type CheckoutSessionPaymentIntentDataTransferDataParams struct {
	Amount      *int64  `form:"amount"`
	Destination *string `form:"destination"`
}

// CheckoutSessionPaymentIntentDataParams is the set of parameters allowed for the
// payment intent creation on a checkout session.
type CheckoutSessionPaymentIntentDataParams struct {
	Params                    `form:"*"`
	ApplicationFeeAmount      *int64                                              `form:"application_fee_amount"`
	CaptureMethod             *string                                             `form:"capture_method"`
	Description               *string                                             `form:"description"`
	Metadata                  map[string]string                                   `form:"metadata"`
	OnBehalfOf                *string                                             `form:"on_behalf_of"`
	ReceiptEmail              *string                                             `form:"receipt_email"`
	SetupFutureUsage          *string                                             `form:"setup_future_usage"`
	Shipping                  *ShippingDetailsParams                              `form:"shipping"`
	StatementDescriptor       *string                                             `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                                             `form:"statement_descriptor_suffix"`
	TransferData              *CheckoutSessionPaymentIntentDataTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                                             `form:"transfer_group"`
}

// CheckoutSessionSetupIntentDataParams is the set of parameters allowed for the setup intent
// creation on a checkout session.
type CheckoutSessionSetupIntentDataParams struct {
	Params      `form:"*"`
	Description *string `form:"description"`
	OnBehalfOf  *string `form:"on_behalf_of"`
}

// CheckoutSessionShippingAddressCollectionParams is the set of parameters allowed for the
// shipping address collection.
type CheckoutSessionShippingAddressCollectionParams struct {
	AllowedCountries []*string `form:"allowed_countries"`
}

// CheckoutSessionSubscriptionDataItemsParams is the set of parameters allowed for one item on a
// checkout session associated with a subscription.
type CheckoutSessionSubscriptionDataItemsParams struct {
	Plan     *string   `form:"plan"`
	Quantity *int64    `form:"quantity"`
	TaxRates []*string `form:"tax_rates"`
}

// CheckoutSessionSubscriptionDataTransferDataParams is the set of parameters allowed
// for the transfer_data hash.
type CheckoutSessionSubscriptionDataTransferDataParams struct {
	AmountPercent *float64 `form:"amount_percent"`
	Destination   *string  `form:"destination"`
}

// CheckoutSessionSubscriptionDataParams is the set of parameters allowed for the subscription
// creation on a checkout session.
type CheckoutSessionSubscriptionDataParams struct {
	Params                `form:"*"`
	ApplicationFeePercent *float64                                           `form:"application_fee_percent"`
	Coupon                *string                                            `form:"coupon"`
	DefaultTaxRates       []*string                                          `form:"default_tax_rates"`
	Items                 []*CheckoutSessionSubscriptionDataItemsParams      `form:"items"`
	Metadata              map[string]string                                  `form:"metadata"`
	TransferData          *CheckoutSessionSubscriptionDataTransferDataParams `form:"transfer_data"`
	TrialEnd              *int64                                             `form:"trial_end"`
	TrialFromPlan         *bool                                              `form:"trial_from_plan"`
	TrialPeriodDays       *int64                                             `form:"trial_period_days"`
}

// CheckoutSessionParams is the set of parameters that can be used when creating
// a checkout session.
// For more details see https://stripe.com/docs/api/checkout/sessions/create
type CheckoutSessionParams struct {
	Params                    `form:"*"`
	AllowPromotionCodes       *bool                                           `form:"allow_promotion_codes"`
	BillingAddressCollection  *string                                         `form:"billing_address_collection"`
	CancelURL                 *string                                         `form:"cancel_url"`
	ClientReferenceID         *string                                         `form:"client_reference_id"`
	Customer                  *string                                         `form:"customer"`
	CustomerEmail             *string                                         `form:"customer_email"`
	Discounts                 []*CheckoutSessionDiscountParams                `form:"discounts"`
	LineItems                 []*CheckoutSessionLineItemParams                `form:"line_items"`
	Locale                    *string                                         `form:"locale"`
	Mode                      *string                                         `form:"mode"`
	PaymentIntentData         *CheckoutSessionPaymentIntentDataParams         `form:"payment_intent_data"`
	PaymentMethodTypes        []*string                                       `form:"payment_method_types"`
	SetupIntentData           *CheckoutSessionSetupIntentDataParams           `form:"setup_intent_data"`
	ShippingAddressCollection *CheckoutSessionShippingAddressCollectionParams `form:"shipping_address_collection"`
	ShippingRates             []*string                                       `form:"shipping_rates"`
	SubmitType                *string                                         `form:"submit_type"`
	SubscriptionData          *CheckoutSessionSubscriptionDataParams          `form:"subscription_data"`
	SuccessURL                *string                                         `form:"success_url"`
}

// CheckoutSessionListLineItemsParams is the set of parameters that can be
// used when listing line items on a session.
type CheckoutSessionListLineItemsParams struct {
	ListParams `form:"*"`
	Session    *string `form:"-"` // Included in URL
}

// CheckoutSessionListParams is the set of parameters that can be
// used when listing sessions.
// For more details see: https://stripe.com/docs/api/checkout/sessions/list
type CheckoutSessionListParams struct {
	ListParams    `form:"*"`
	PaymentIntent *string `form:"payment_intent"`
	Subscription  *string `form:"subscription"`
}

// CheckoutSessionCustomerDetailsTaxIDs represent customer's tax IDs at the
// time of checkout.
type CheckoutSessionCustomerDetailsTaxIDs struct {
	Type  CheckoutSessionCustomerDetailsTaxIDsType `json:"type"`
	Value string                                   `json:"value"`
}

// CheckoutSessionCustomerDetails represent the customer details including
// the tax exempt status and the customer's tax IDs.
type CheckoutSessionCustomerDetails struct {
	Email     string                                  `json:"email"`
	TaxExempt CheckoutSessionCustomerDetailsTaxExempt `json:"tax_exempt"`
	TaxIDs    []*CheckoutSessionCustomerDetailsTaxIDs `json:"tax_ids"`
}

// CheckoutSessionShippingAddressCollection is the set of parameters allowed for the
// shipping address collection.
type CheckoutSessionShippingAddressCollection struct {
	AllowedCountries []string `json:"allowed_countries"`
}

// CheckoutSessionTotalDetailsBreakdownDiscount represent the details of one discount applied to a session.
type CheckoutSessionTotalDetailsBreakdownDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// CheckoutSessionTotalDetailsBreakdownTax represent the details of tax rate applied to a session.
type CheckoutSessionTotalDetailsBreakdownTax struct {
	Amount  int64    `json:"amount"`
	TaxRate *TaxRate `json:"tax_rate"`
}

// CheckoutSessionTotalDetailsBreakdown is the set of properties detailing a breakdown of taxes and discounts applied to a session if any.
type CheckoutSessionTotalDetailsBreakdown struct {
	Discounts []*CheckoutSessionTotalDetailsBreakdownDiscount `json:"discounts"`
	Taxes     []*CheckoutSessionTotalDetailsBreakdownTax      `json:"taxes"`
}

// CheckoutSessionTotalDetails is the set of properties detailing how the amounts were calculated.
type CheckoutSessionTotalDetails struct {
	AmountDiscount int64                                 `json:"amount_discount"`
	AmountShipping int64                                 `json:"amount_shipping"`
	AmountTax      int64                                 `json:"amount_tax"`
	Breakdown      *CheckoutSessionTotalDetailsBreakdown `json:"breakdown"`
}

// CheckoutSession is the resource representing a Stripe checkout session.
// For more details see https://stripe.com/docs/api/checkout/sessions/object
type CheckoutSession struct {
	APIResource
	AllowPromotionCodes       bool                                      `json:"allow_promotion_codes"`
	CancelURL                 string                                    `json:"cancel_url"`
	CustomerDetails           *CheckoutSessionCustomerDetails           `json:"customer_details"`
	AmountSubtotal            int64                                     `json:"amount_subtotal"`
	AmountTotal               int64                                     `json:"amount_total"`
	ClientReferenceID         string                                    `json:"client_reference_id"`
	Currency                  Currency                                  `json:"currency"`
	Customer                  *Customer                                 `json:"customer"`
	CustomerEmail             string                                    `json:"customer_email"`
	Deleted                   bool                                      `json:"deleted"`
	ID                        string                                    `json:"id"`
	LineItems                 *LineItemList                             `json:"line_items"`
	Livemode                  bool                                      `json:"livemode"`
	Locale                    string                                    `json:"locale"`
	Metadata                  map[string]string                         `json:"metadata"`
	Mode                      CheckoutSessionMode                       `json:"mode"`
	Object                    string                                    `json:"object"`
	PaymentIntent             *PaymentIntent                            `json:"payment_intent"`
	PaymentMethodTypes        []string                                  `json:"payment_method_types"`
	PaymentStatus             CheckoutSessionPaymentStatus              `json:"payment_status"`
	SetupIntent               *SetupIntent                              `json:"setup_intent"`
	Shipping                  *ShippingDetails                          `json:"shipping"`
	ShippingAddressCollection *CheckoutSessionShippingAddressCollection `json:"shipping_address_collection"`
	Subscription              *Subscription                             `json:"subscription"`
	SubmitType                CheckoutSessionSubmitType                 `json:"submit_type"`
	SuccessURL                string                                    `json:"success_url"`
	TotalDetails              *CheckoutSessionTotalDetails              `json:"total_details"`
}

// CheckoutSessionList is a list of sessions as retrieved from a list endpoint.
type CheckoutSessionList struct {
	APIResource
	ListMeta
	Data []*CheckoutSession `json:"data"`
}

// UnmarshalJSON handles deserialization of a checkout session.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *CheckoutSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type session CheckoutSession
	var v session
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = CheckoutSession(v)
	return nil
}
