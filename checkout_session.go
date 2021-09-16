//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// CheckoutSessionCustomerDetailsTaxIDsType is the list of allowed values for type
// on the tax_ids inside customer_details of a checkout session.
type CheckoutSessionCustomerDetailsTaxIDsType string

// List of values that CheckoutSessionCustomerDetailsTaxIDsType can take.
const (
	CheckoutSessionCustomerDetailsTaxIDsTypeAETRN    CheckoutSessionCustomerDetailsTaxIDsType = "ae_trn"
	CheckoutSessionCustomerDetailsTaxIDsTypeAUABN    CheckoutSessionCustomerDetailsTaxIDsType = "au_abn"
	CheckoutSessionCustomerDetailsTaxIDsTypeAUARN    CheckoutSessionCustomerDetailsTaxIDsType = "au_arn"
	CheckoutSessionCustomerDetailsTaxIDsTypeBRCNPJ   CheckoutSessionCustomerDetailsTaxIDsType = "br_cnpj"
	CheckoutSessionCustomerDetailsTaxIDsTypeBRCPF    CheckoutSessionCustomerDetailsTaxIDsType = "br_cpf"
	CheckoutSessionCustomerDetailsTaxIDsTypeCABN     CheckoutSessionCustomerDetailsTaxIDsType = "ca_bn"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAGSTHST CheckoutSessionCustomerDetailsTaxIDsType = "ca_gst_hst"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTBC  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_bc"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTMB  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_mb"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTSK  CheckoutSessionCustomerDetailsTaxIDsType = "ca_pst_sk"
	CheckoutSessionCustomerDetailsTaxIDsTypeCAQST    CheckoutSessionCustomerDetailsTaxIDsType = "ca_qst"
	CheckoutSessionCustomerDetailsTaxIDsTypeCHVAT    CheckoutSessionCustomerDetailsTaxIDsType = "ch_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeCLTIN    CheckoutSessionCustomerDetailsTaxIDsType = "cl_tin"
	CheckoutSessionCustomerDetailsTaxIDsTypeESCIF    CheckoutSessionCustomerDetailsTaxIDsType = "es_cif"
	CheckoutSessionCustomerDetailsTaxIDsTypeEUVAT    CheckoutSessionCustomerDetailsTaxIDsType = "eu_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeGBVAT    CheckoutSessionCustomerDetailsTaxIDsType = "gb_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeHKBR     CheckoutSessionCustomerDetailsTaxIDsType = "hk_br"
	CheckoutSessionCustomerDetailsTaxIDsTypeIDNPWP   CheckoutSessionCustomerDetailsTaxIDsType = "id_npwp"
	CheckoutSessionCustomerDetailsTaxIDsTypeILVAT    CheckoutSessionCustomerDetailsTaxIDsType = "il_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeINGST    CheckoutSessionCustomerDetailsTaxIDsType = "in_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeJPCN     CheckoutSessionCustomerDetailsTaxIDsType = "jp_cn"
	CheckoutSessionCustomerDetailsTaxIDsTypeJPRN     CheckoutSessionCustomerDetailsTaxIDsType = "jp_rn"
	CheckoutSessionCustomerDetailsTaxIDsTypeKRBRN    CheckoutSessionCustomerDetailsTaxIDsType = "kr_brn"
	CheckoutSessionCustomerDetailsTaxIDsTypeLIUID    CheckoutSessionCustomerDetailsTaxIDsType = "li_uid"
	CheckoutSessionCustomerDetailsTaxIDsTypeMXRFC    CheckoutSessionCustomerDetailsTaxIDsType = "mx_rfc"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYFRP    CheckoutSessionCustomerDetailsTaxIDsType = "my_frp"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYITN    CheckoutSessionCustomerDetailsTaxIDsType = "my_itn"
	CheckoutSessionCustomerDetailsTaxIDsTypeMYSST    CheckoutSessionCustomerDetailsTaxIDsType = "my_sst"
	CheckoutSessionCustomerDetailsTaxIDsTypeNOVAT    CheckoutSessionCustomerDetailsTaxIDsType = "no_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeNZGST    CheckoutSessionCustomerDetailsTaxIDsType = "nz_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeRUINN    CheckoutSessionCustomerDetailsTaxIDsType = "ru_inn"
	CheckoutSessionCustomerDetailsTaxIDsTypeRUKPP    CheckoutSessionCustomerDetailsTaxIDsType = "ru_kpp"
	CheckoutSessionCustomerDetailsTaxIDsTypeSAVAT    CheckoutSessionCustomerDetailsTaxIDsType = "sa_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeSGGST    CheckoutSessionCustomerDetailsTaxIDsType = "sg_gst"
	CheckoutSessionCustomerDetailsTaxIDsTypeSGUEN    CheckoutSessionCustomerDetailsTaxIDsType = "sg_uen"
	CheckoutSessionCustomerDetailsTaxIDsTypeTHVAT    CheckoutSessionCustomerDetailsTaxIDsType = "th_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeTWVAT    CheckoutSessionCustomerDetailsTaxIDsType = "tw_vat"
	CheckoutSessionCustomerDetailsTaxIDsTypeUnknown  CheckoutSessionCustomerDetailsTaxIDsType = "unknown"
	CheckoutSessionCustomerDetailsTaxIDsTypeUSEIN    CheckoutSessionCustomerDetailsTaxIDsType = "us_ein"
	CheckoutSessionCustomerDetailsTaxIDsTypeZAVAT    CheckoutSessionCustomerDetailsTaxIDsType = "za_vat"
)

type CheckoutSessionAutomaticTaxStatus string

// List of values that CheckoutSessionAutomaticTaxStatus can take
const (
	CheckoutSessionAutomaticTaxStatusComplete               CheckoutSessionAutomaticTaxStatus = "complete"
	CheckoutSessionAutomaticTaxStatusFailed                 CheckoutSessionAutomaticTaxStatus = "failed"
	CheckoutSessionAutomaticTaxStatusRequiresLocationInputs CheckoutSessionAutomaticTaxStatus = "requires_location_inputs"
)

// Describes whether Checkout should collect the customer's billing address.
type CheckoutSessionBillingAddressCollection string

// List of values that CheckoutSessionBillingAddressCollection can take
const (
	CheckoutSessionBillingAddressCollectionAuto     CheckoutSessionBillingAddressCollection = "auto"
	CheckoutSessionBillingAddressCollectionRequired CheckoutSessionBillingAddressCollection = "required"
)

// If `opt_in`, the customer consents to receiving promotional communications
// from the merchant about this Checkout Session.
type CheckoutSessionConsentPromotions string

// List of values that CheckoutSessionConsentPromotions can take
const (
	CheckoutSessionConsentPromotionsOptIn  CheckoutSessionConsentPromotions = "opt_in"
	CheckoutSessionConsentPromotionsOptOut CheckoutSessionConsentPromotions = "opt_out"
)

// If set to `auto`, enables the collection of customer consent for promotional communications. The Checkout
// Session will determine whether to display an option to opt into promotional communication
// from the merchant depending on the customer's locale. Only available to US merchants.
type CheckoutSessionConsentCollectionPromotions string

// List of values that CheckoutSessionConsentCollectionPromotions can take
const (
	CheckoutSessionConsentCollectionPromotionsAuto CheckoutSessionConsentCollectionPromotions = "auto"
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

// List of Stripe products where this mandate can be selected automatically. Returned when the Session is in `setup` mode.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultForInvoice      CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "invoice"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultForSubscription CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "subscription"
)

// CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule is the list of allowed values for payment_schedule
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleCombined CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "combined"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleInterval CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "interval"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleSporadic CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "sporadic"
)

// CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType is the list of allowed values for transaction_type
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod is the list of allowed values for verification_method
type CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodInstant       CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
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
	TaxCode     *string           `form:"tax_code"`
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
	TaxBehavior       *string                                            `form:"tax_behavior"`
	UnitAmount        *int64                                             `form:"unit_amount"`
	UnitAmountDecimal *float64                                           `form:"unit_amount_decimal,high_precision"`
}

// Configure a Checkout Session that can be used to recover an expired session.
type CheckoutSessionAfterExpirationRecoveryParams struct {
	AllowPromotionCodes *bool `form:"allow_promotion_codes"`
	Enabled             *bool `form:"enabled"`
}

// Configure actions after a Checkout Session has expired.
type CheckoutSessionAfterExpirationParams struct {
	Recovery *CheckoutSessionAfterExpirationRecoveryParams `form:"recovery"`
}

// Settings for automatic tax lookup for this session and resulting payments, invoices, and subscriptions.
type CheckoutSessionAutomaticTaxParams struct {
	Enabled *bool `form:"enabled"`
}

// Configure fields for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollectionParams struct {
	Promotions *string `form:"promotions"`
}

// Controls what fields on Customer can be updated by the Checkout Session. Can only be provided when `customer` is provided.
type CheckoutSessionCustomerUpdateParams struct {
	Address  *string `form:"address"`
	Name     *string `form:"name"`
	Shipping *string `form:"shipping"`
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

// CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams is the set of parameters allowed for mandate_options for acss debit.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	CustomMandateURL    *string   `form:"custom_mandate_url"`
	DefaultFor          []*string `form:"default_for"`
	IntervalDescription *string   `form:"interval_description"`
	PaymentSchedule     *string   `form:"payment_schedule"`
	TransactionType     *string   `form:"transaction_type"`
}

// CheckoutSessionPaymentMethodOptionsACSSDebitParams is the set of parameters allowed for acss_debit on payment_method_options.
type CheckoutSessionPaymentMethodOptionsACSSDebitParams struct {
	Currency           *string                                                           `form:"currency"`
	MandateOptions     *CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                           `form:"verification_method"`
}

// CheckoutSessionPaymentMethodOptionsBoletoParams is the set of parameters allowed for boleto on payment_method_options.
type CheckoutSessionPaymentMethodOptionsBoletoParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}

// CheckoutSessionPaymentMethodOptionsOXXOParams is the set of parameters allowed for oxxo on payment_method_options.
type CheckoutSessionPaymentMethodOptionsOXXOParams struct {
	ExpiresAfterDays *int64 `form:"expires_after_days"`
}

// CheckoutSessionPaymentMethodOptionsWechatPayParams is the set of parameters allowed for wechat_pay on payment_method_options.
type CheckoutSessionPaymentMethodOptionsWechatPayParams struct {
	AppID  *string `form:"app_id"`
	Client *string `form:"client"`
}

// CheckoutSessionPaymentMethodOptionsParams is the set of allowed parameters for payment_method_options on a checkout session.
type CheckoutSessionPaymentMethodOptionsParams struct {
	ACSSDebit *CheckoutSessionPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	Boleto    *CheckoutSessionPaymentMethodOptionsBoletoParams    `form:"boleto"`
	OXXO      *CheckoutSessionPaymentMethodOptionsOXXOParams      `form:"oxxo"`
	WechatPay *CheckoutSessionPaymentMethodOptionsWechatPayParams `form:"wechat_pay"`
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

// Controls tax ID collection settings for the session.
type CheckoutSessionTaxIDCollectionParams struct {
	Enabled *bool `form:"enabled"`
}

// CheckoutSessionParams is the set of parameters that can be used when creating
// a checkout session.
// For more details see https://stripe.com/docs/api/checkout/sessions/create
type CheckoutSessionParams struct {
	Params                    `form:"*"`
	AfterExpiration           *CheckoutSessionAfterExpirationParams           `form:"after_expiration"`
	AllowPromotionCodes       *bool                                           `form:"allow_promotion_codes"`
	AutomaticTax              *CheckoutSessionAutomaticTaxParams              `form:"automatic_tax"`
	BillingAddressCollection  *string                                         `form:"billing_address_collection"`
	CancelURL                 *string                                         `form:"cancel_url"`
	ClientReferenceID         *string                                         `form:"client_reference_id"`
	ConsentCollection         *CheckoutSessionConsentCollectionParams         `form:"consent_collection"`
	Customer                  *string                                         `form:"customer"`
	CustomerEmail             *string                                         `form:"customer_email"`
	CustomerUpdate            *CheckoutSessionCustomerUpdateParams            `form:"customer_update"`
	Discounts                 []*CheckoutSessionDiscountParams                `form:"discounts"`
	ExpiresAt                 *int64                                          `form:"expires_at"`
	LineItems                 []*CheckoutSessionLineItemParams                `form:"line_items"`
	Locale                    *string                                         `form:"locale"`
	Mode                      *string                                         `form:"mode"`
	PaymentIntentData         *CheckoutSessionPaymentIntentDataParams         `form:"payment_intent_data"`
	PaymentMethodOptions      *CheckoutSessionPaymentMethodOptionsParams      `form:"payment_method_options"`
	PaymentMethodTypes        []*string                                       `form:"payment_method_types"`
	SetupIntentData           *CheckoutSessionSetupIntentDataParams           `form:"setup_intent_data"`
	ShippingAddressCollection *CheckoutSessionShippingAddressCollectionParams `form:"shipping_address_collection"`
	ShippingRates             []*string                                       `form:"shipping_rates"`
	SubmitType                *string                                         `form:"submit_type"`
	SubscriptionData          *CheckoutSessionSubscriptionDataParams          `form:"subscription_data"`
	SuccessURL                *string                                         `form:"success_url"`
	TaxIDCollection           *CheckoutSessionTaxIDCollectionParams           `form:"tax_id_collection"`
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

// When set, configuration used to recover the Checkout Session on expiry.
type CheckoutSessionAfterExpirationRecovery struct {
	AllowPromotionCodes bool   `json:"allow_promotion_codes"`
	Enabled             bool   `json:"enabled"`
	ExpiresAt           int64  `json:"expires_at"`
	URL                 string `json:"url"`
}

// When set, provides configuration for actions to take if this Checkout Session expires.
type CheckoutSessionAfterExpiration struct {
	Recovery *CheckoutSessionAfterExpirationRecovery `json:"recovery"`
}
type CheckoutSessionAutomaticTax struct {
	Enabled bool                              `json:"enabled"`
	Status  CheckoutSessionAutomaticTaxStatus `json:"status"`
}

// Results of `consent_collection` for this session.
type CheckoutSessionConsent struct {
	Promotions CheckoutSessionConsentPromotions `json:"promotions"`
}

// When set, provides configuration for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollection struct {
	Promotions CheckoutSessionConsentCollectionPromotions `json:"promotions"`
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

// CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptions represent mandate options for acss debit.
type CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptions struct {
	CustomMandateURL    string                                                                    `json:"custom_mandate_url"`
	DefaultFor          []CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor    `json:"default_for"`
	IntervalDescription string                                                                    `json:"interval_description"`
	PaymentSchedule     CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	TransactionType     CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// CheckoutSessionPaymentMethodOptionsACSSDebit represent the options for acss debit on payment_method_options.
type CheckoutSessionPaymentMethodOptionsACSSDebit struct {
	Currency           string                                                         `json:"currency"`
	MandateOptions     *CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod CheckoutSessionPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// CheckoutSessionPaymentMethodOptionsBoleto represent the options for boleto on payment_method_options.
type CheckoutSessionPaymentMethodOptionsBoleto struct {
	ExpiresAfterDays int64 `json:"expires_after_days"`
}

// CheckoutSessionPaymentMethodOptionsOXXO represent the options for oxxo on payment_method_options.
type CheckoutSessionPaymentMethodOptionsOXXO struct {
	ExpiresAfterDays int64 `json:"expires_after_days"`
}

// CheckoutSessionPaymentMethodOptions represent payment-method-specific options for a checkout session.
type CheckoutSessionPaymentMethodOptions struct {
	ACSSDebit *CheckoutSessionPaymentMethodOptionsACSSDebit `json:"acss_debit"`
	Boleto    *CheckoutSessionPaymentMethodOptionsBoleto    `json:"boleto"`
	OXXO      *CheckoutSessionPaymentMethodOptionsOXXO      `json:"oxxo"`
}

// CheckoutSessionShippingAddressCollection is the set of parameters allowed for the
// shipping address collection.
type CheckoutSessionShippingAddressCollection struct {
	AllowedCountries []string `json:"allowed_countries"`
}
type CheckoutSessionTaxIDCollection struct {
	Enabled bool `json:"enabled"`
}

// CheckoutSessionTotalDetailsBreakdownDiscount represent the details of one discount applied to a session.
type CheckoutSessionTotalDetailsBreakdownDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// CheckoutSessionTotalDetailsBreakdownTax represent the details of tax rate applied to a session.
type CheckoutSessionTotalDetailsBreakdownTax struct {
	Amount  int64    `json:"amount"`
	Rate    *TaxRate `json:"rate"`
	TaxRate *TaxRate `json:"tax_rate"` // Do not use: use `Rate`
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
	AfterExpiration           *CheckoutSessionAfterExpiration           `json:"after_expiration"`
	AllowPromotionCodes       bool                                      `json:"allow_promotion_codes"`
	AmountSubtotal            int64                                     `json:"amount_subtotal"`
	AmountTotal               int64                                     `json:"amount_total"`
	AutomaticTax              *CheckoutSessionAutomaticTax              `json:"automatic_tax"`
	BillingAddressCollection  CheckoutSessionBillingAddressCollection   `json:"billing_address_collection"`
	CancelURL                 string                                    `json:"cancel_url"`
	ClientReferenceID         string                                    `json:"client_reference_id"`
	Consent                   *CheckoutSessionConsent                   `json:"consent"`
	ConsentCollection         *CheckoutSessionConsentCollection         `json:"consent_collection"`
	Currency                  Currency                                  `json:"currency"`
	Customer                  *Customer                                 `json:"customer"`
	CustomerDetails           *CheckoutSessionCustomerDetails           `json:"customer_details"`
	CustomerEmail             string                                    `json:"customer_email"`
	Deleted                   bool                                      `json:"deleted"`
	ExpiresAt                 int64                                     `json:"expires_at"`
	ID                        string                                    `json:"id"`
	LineItems                 *LineItemList                             `json:"line_items"`
	Livemode                  bool                                      `json:"livemode"`
	Locale                    string                                    `json:"locale"`
	Metadata                  map[string]string                         `json:"metadata"`
	Mode                      CheckoutSessionMode                       `json:"mode"`
	Object                    string                                    `json:"object"`
	PaymentIntent             *PaymentIntent                            `json:"payment_intent"`
	PaymentMethodOptions      *CheckoutSessionPaymentMethodOptions      `json:"payment_method_options"`
	PaymentMethodTypes        []string                                  `json:"payment_method_types"`
	PaymentStatus             CheckoutSessionPaymentStatus              `json:"payment_status"`
	RecoveredFrom             string                                    `json:"recovered_from"`
	SetupIntent               *SetupIntent                              `json:"setup_intent"`
	Shipping                  *ShippingDetails                          `json:"shipping"`
	ShippingAddressCollection *CheckoutSessionShippingAddressCollection `json:"shipping_address_collection"`
	SubmitType                CheckoutSessionSubmitType                 `json:"submit_type"`
	Subscription              *Subscription                             `json:"subscription"`
	SuccessURL                string                                    `json:"success_url"`
	TaxIDCollection           *CheckoutSessionTaxIDCollection           `json:"tax_id_collection"`
	TotalDetails              *CheckoutSessionTotalDetails              `json:"total_details"`
	URL                       string                                    `json:"url"`
}

// UnmarshalJSON handles deserialization of a CheckoutSession.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *CheckoutSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type checkoutSession CheckoutSession
	var v checkoutSession
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = CheckoutSession(v)
	return nil
}

// CheckoutSessionList is a list of sessions as retrieved from a list endpoint.
type CheckoutSessionList struct {
	APIResource
	ListMeta
	Data []*CheckoutSession `json:"data"`
}
