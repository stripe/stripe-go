package example

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v72"
	account "github.com/stripe/stripe-go/v72/account"
	accountlink "github.com/stripe/stripe-go/v72/accountlink"
	balancetransaction "github.com/stripe/stripe-go/v72/balancetransaction"
	billingportal_configuration "github.com/stripe/stripe-go/v72/billingportal/configuration"
	billingportal_session "github.com/stripe/stripe-go/v72/billingportal/session"
	capability "github.com/stripe/stripe-go/v72/capability"
	charge "github.com/stripe/stripe-go/v72/charge"
	checkout_session "github.com/stripe/stripe-go/v72/checkout/session"
	countryspec "github.com/stripe/stripe-go/v72/countryspec"
	coupon "github.com/stripe/stripe-go/v72/coupon"
	customer "github.com/stripe/stripe-go/v72/customer"
	customerbalancetransaction "github.com/stripe/stripe-go/v72/customerbalancetransaction"
	dispute "github.com/stripe/stripe-go/v72/dispute"
	event "github.com/stripe/stripe-go/v72/event"
	invoice "github.com/stripe/stripe-go/v72/invoice"
	invoiceitem "github.com/stripe/stripe-go/v72/invoiceitem"
	issuing_authorization "github.com/stripe/stripe-go/v72/issuing/authorization"
	issuing_card "github.com/stripe/stripe-go/v72/issuing/card"
	issuing_cardholder "github.com/stripe/stripe-go/v72/issuing/cardholder"
	issuing_dispute "github.com/stripe/stripe-go/v72/issuing/dispute"
	issuing_transaction "github.com/stripe/stripe-go/v72/issuing/transaction"
	mandate "github.com/stripe/stripe-go/v72/mandate"
	order "github.com/stripe/stripe-go/v72/order"
	orderreturn "github.com/stripe/stripe-go/v72/orderreturn"
	paymentintent "github.com/stripe/stripe-go/v72/paymentintent"
	paymentlink "github.com/stripe/stripe-go/v72/paymentlink"
	paymentmethod "github.com/stripe/stripe-go/v72/paymentmethod"
	payout "github.com/stripe/stripe-go/v72/payout"
	person "github.com/stripe/stripe-go/v72/person"
	plan "github.com/stripe/stripe-go/v72/plan"
	price "github.com/stripe/stripe-go/v72/price"
	product "github.com/stripe/stripe-go/v72/product"
	promotioncode "github.com/stripe/stripe-go/v72/promotioncode"
	radar_earlyfraudwarning "github.com/stripe/stripe-go/v72/radar/earlyfraudwarning"
	refund "github.com/stripe/stripe-go/v72/refund"
	reversal "github.com/stripe/stripe-go/v72/reversal"
	review "github.com/stripe/stripe-go/v72/review"
	setupattempt "github.com/stripe/stripe-go/v72/setupattempt"
	setupintent "github.com/stripe/stripe-go/v72/setupintent"
	shippingrate "github.com/stripe/stripe-go/v72/shippingrate"
	sigma_scheduledqueryrun "github.com/stripe/stripe-go/v72/sigma/scheduledqueryrun"
	sku "github.com/stripe/stripe-go/v72/sku"
	source "github.com/stripe/stripe-go/v72/source"
	taxid "github.com/stripe/stripe-go/v72/taxid"
	taxrate "github.com/stripe/stripe-go/v72/taxrate"
	terminal_configuration "github.com/stripe/stripe-go/v72/terminal/configuration"
	terminal_connectiontoken "github.com/stripe/stripe-go/v72/terminal/connectiontoken"
	terminal_location "github.com/stripe/stripe-go/v72/terminal/location"
	terminal_reader "github.com/stripe/stripe-go/v72/terminal/reader"
	testhelpers_refund "github.com/stripe/stripe-go/v72/testhelpers/refund"
	testhelpers_testclock "github.com/stripe/stripe-go/v72/testhelpers/testclock"
	_ "github.com/stripe/stripe-go/v72/testing"
	topup "github.com/stripe/stripe-go/v72/topup"
	transfer "github.com/stripe/stripe-go/v72/transfer"
	usagerecord "github.com/stripe/stripe-go/v72/usagerecord"
	usagerecordsummary "github.com/stripe/stripe-go/v72/usagerecordsummary"
	webhookendpoint "github.com/stripe/stripe-go/v72/webhookendpoint"
)

func TestCustomerList(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := customer.List(params)
	assert.NotNil(t, result)
}

func TestBalanceTransactionRetrieve(t *testing.T) {
	params := &stripe.BalanceTransactionParams{}
	result, _ := balancetransaction.Get("txn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestBalanceTransactionList(t *testing.T) {
	params := &stripe.BalanceTransactionListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := balancetransaction.List(params)
	assert.NotNil(t, result)
}

func TestChargeCreate(t *testing.T) {
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Source:      &stripe.SourceParams{Token: stripe.String("tok_xxxx")},
		Description: stripe.String("My First Test Charge (created for API docs)"),
	}
	result, _ := charge.New(params)
	assert.NotNil(t, result)
}

func TestChargeRetrieve(t *testing.T) {
	params := &stripe.ChargeParams{}
	result, _ := charge.Get("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestChargeUpdate(t *testing.T) {
	params := &stripe.ChargeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := charge.Update("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestChargeCapture(t *testing.T) {
	params := &stripe.CaptureParams{}
	result, _ := charge.Capture("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestChargeList(t *testing.T) {
	params := &stripe.ChargeListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := charge.List(params)
	assert.NotNil(t, result)
}

func TestCustomerCreate(t *testing.T) {
	params := &stripe.CustomerParams{
		Description: stripe.String("My First Test Customer (created for API docs)"),
	}
	result, _ := customer.New(params)
	assert.NotNil(t, result)
}

func TestCustomerRetrieve(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, _ := customer.Get("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerUpdate(t *testing.T) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := customer.Update("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerDelete(t *testing.T) {
	params := &stripe.CustomerParams{}
	result, _ := customer.Del("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerList2(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := customer.List(params)
	assert.NotNil(t, result)
}

func TestDisputeRetrieve(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, _ := dispute.Get("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestDisputeUpdate(t *testing.T) {
	params := &stripe.DisputeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := dispute.Update("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestDisputeClose(t *testing.T) {
	params := &stripe.DisputeParams{}
	result, _ := dispute.Close("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestDisputeList(t *testing.T) {
	params := &stripe.DisputeListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := dispute.List(params)
	assert.NotNil(t, result)
}

func TestEventRetrieve(t *testing.T) {
	params := &stripe.EventParams{}
	result, _ := event.Get("evt_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestEventList(t *testing.T) {
	params := &stripe.EventListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := event.List(params)
	assert.NotNil(t, result)
}

func TestMandateRetrieve(t *testing.T) {
	params := &stripe.MandateParams{}
	result, _ := mandate.Get("mandate_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentCreate(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(2000),
		Currency:           stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, _ := paymentintent.New(params)
	assert.NotNil(t, result)
}

func TestPaymentIntentRetrieve(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	result, _ := paymentintent.Get("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentUpdate(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := paymentintent.Update("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentConfirm(t *testing.T) {
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, _ := paymentintent.Confirm("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentCapture(t *testing.T) {
	params := &stripe.PaymentIntentCaptureParams{}
	result, _ := paymentintent.Capture("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentCancel(t *testing.T) {
	params := &stripe.PaymentIntentCancelParams{}
	result, _ := paymentintent.Cancel("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentList(t *testing.T) {
	params := &stripe.PaymentIntentListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := paymentintent.List(params)
	assert.NotNil(t, result)
}

func TestSetupIntentCreate(t *testing.T) {
	params := &stripe.SetupIntentParams{
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, _ := setupintent.New(params)
	assert.NotNil(t, result)
}

func TestSetupIntentRetrieve(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	result, _ := setupintent.Get("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentUpdate(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	params.AddMetadata("user_id", "3435453")
	result, _ := setupintent.Update("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentConfirm(t *testing.T) {
	params := &stripe.SetupIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, _ := setupintent.Confirm("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentCancel(t *testing.T) {
	params := &stripe.SetupIntentCancelParams{}
	result, _ := setupintent.Cancel("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentList(t *testing.T) {
	params := &stripe.SetupIntentListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := setupintent.List(params)
	assert.NotNil(t, result)
}

func TestSetupAttemptList(t *testing.T) {
	params := &stripe.SetupAttemptListParams{
		SetupIntent: stripe.String("seti_xxxxxxxxxxxxx"),
	}
	params.Filters.AddFilter("Limit", "", "3")
	result := setupattempt.List(params)
	assert.NotNil(t, result)
}

func TestPayoutCreate(t *testing.T) {
	params := &stripe.PayoutParams{
		Amount:   stripe.Int64(1100),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	result, _ := payout.New(params)
	assert.NotNil(t, result)
}

func TestPayoutRetrieve(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, _ := payout.Get("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPayoutUpdate(t *testing.T) {
	params := &stripe.PayoutParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := payout.Update("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPayoutList(t *testing.T) {
	params := &stripe.PayoutListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := payout.List(params)
	assert.NotNil(t, result)
}

func TestPayoutCancel(t *testing.T) {
	params := &stripe.PayoutParams{}
	result, _ := payout.Cancel("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPayoutReverse(t *testing.T) {
	params := &stripe.PayoutReverseParams{}
	result, _ := payout.Reverse("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestProductCreate(t *testing.T) {
	params := &stripe.ProductParams{Name: stripe.String("Gold Special")}
	result, _ := product.New(params)
	assert.NotNil(t, result)
}

func TestProductRetrieve(t *testing.T) {
	params := &stripe.ProductParams{}
	result, _ := product.Get("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestProductUpdate(t *testing.T) {
	params := &stripe.ProductParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := product.Update("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestProductList(t *testing.T) {
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := product.List(params)
	assert.NotNil(t, result)
}

func TestProductDelete(t *testing.T) {
	params := &stripe.ProductParams{}
	result, _ := product.Del("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPriceCreate(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, _ := price.New(params)
	assert.NotNil(t, result)
}

func TestPriceRetrieve(t *testing.T) {
	params := &stripe.PriceParams{}
	result, _ := price.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPriceUpdate(t *testing.T) {
	params := &stripe.PriceParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := price.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPriceList(t *testing.T) {
	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := price.List(params)
	assert.NotNil(t, result)
}

func TestRefundCreate(t *testing.T) {
	params := &stripe.RefundParams{Charge: stripe.String("ch_xxxxxxxxxxxxx")}
	result, _ := refund.New(params)
	assert.NotNil(t, result)
}

func TestRefundRetrieve(t *testing.T) {
	params := &stripe.RefundParams{}
	result, _ := refund.Get("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestRefundUpdate(t *testing.T) {
	params := &stripe.RefundParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := refund.Update("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestRefundList(t *testing.T) {
	params := &stripe.RefundListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := refund.List(params)
	assert.NotNil(t, result)
}

func TestPaymentMethodRetrieve(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	result, _ := paymentmethod.Get("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodUpdate(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := paymentmethod.Update("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodList(t *testing.T) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
	}
	result := paymentmethod.List(params)
	assert.NotNil(t, result)
}

func TestPaymentMethodAttach(t *testing.T) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, _ := paymentmethod.Attach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodDetach(t *testing.T) {
	params := &stripe.PaymentMethodDetachParams{}
	result, _ := paymentmethod.Detach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSourceRetrieve(t *testing.T) {
	params := &stripe.SourceObjectParams{}
	result, _ := source.Get("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSourceUpdate(t *testing.T) {
	params := &stripe.SourceObjectParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := source.Update("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCheckoutSessionRetrieve(t *testing.T) {
	params := &stripe.CheckoutSessionParams{}
	result, _ := checkout_session.Get("cs_test_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCheckoutSessionList(t *testing.T) {
	params := &stripe.CheckoutSessionListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := checkout_session.List(params)
	assert.NotNil(t, result)
}

func TestCouponCreate(t *testing.T) {
	params := &stripe.CouponParams{
		PercentOff:       stripe.Float64(25),
		Duration:         stripe.String(string(stripe.CouponDurationRepeating)),
		DurationInMonths: stripe.Int64(3),
	}
	result, _ := coupon.New(params)
	assert.NotNil(t, result)
}

func TestCouponRetrieve(t *testing.T) {
	params := &stripe.CouponParams{}
	result, _ := coupon.Get("25_5OFF", params)
	assert.NotNil(t, result)
}

func TestCouponUpdate(t *testing.T) {
	params := &stripe.CouponParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := coupon.Update("co_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCouponDelete(t *testing.T) {
	params := &stripe.CouponParams{}
	result, _ := coupon.Del("co_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCouponList(t *testing.T) {
	params := &stripe.CouponListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := coupon.List(params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionCreate(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Amount:   stripe.Int64(-500),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	result, _ := customerbalancetransaction.New(params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionRetrieve(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{}
	result, _ := customerbalancetransaction.Get("cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionUpdate(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := customerbalancetransaction.Update("cbtxn_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionList(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := customerbalancetransaction.List(params)
	assert.NotNil(t, result)
}

func TestBillingPortalSessionCreate(t *testing.T) {
	params := &stripe.BillingPortalSessionParams{
		Customer:  stripe.String("cus_xxxxxxxxxxxxx"),
		ReturnURL: stripe.String("https://example.com/account"),
	}
	result, _ := billingportal_session.New(params)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationCreate(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{
		Features: &stripe.BillingPortalConfigurationFeaturesParams{
			CustomerUpdate: &stripe.BillingPortalConfigurationFeaturesCustomerUpdateParams{
				AllowedUpdates: []*string{
					stripe.String(string(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateEmail)),
					stripe.String(string(stripe.BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdateTaxID)),
				},
				Enabled: stripe.Bool(true),
			},
			InvoiceHistory: &stripe.BillingPortalConfigurationFeaturesInvoiceHistoryParams{
				Enabled: stripe.Bool(true),
			},
		},
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, _ := billingportal_configuration.New(params)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationUpdate(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{
		BusinessProfile: &stripe.BillingPortalConfigurationBusinessProfileParams{
			PrivacyPolicyURL:  stripe.String("https://example.com/privacy"),
			TermsOfServiceURL: stripe.String("https://example.com/terms"),
		},
	}
	result, _ := billingportal_configuration.Update("bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationRetrieve(t *testing.T) {
	params := &stripe.BillingPortalConfigurationParams{}
	result, _ := billingportal_configuration.Get("bpc_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationList(t *testing.T) {
	params := &stripe.BillingPortalConfigurationListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := billingportal_configuration.List(params)
	assert.NotNil(t, result)
}

func TestTaxIDCreate(t *testing.T) {
	params := &stripe.TaxIDParams{
		Type:  stripe.String(string(stripe.TaxIDTypeEUVAT)),
		Value: stripe.String("DE123456789"),
	}
	result, _ := taxid.New(params)
	assert.NotNil(t, result)
}

func TestTaxIDRetrieve(t *testing.T) {
	params := &stripe.TaxIDParams{}
	result, _ := taxid.Get("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTaxIDDelete(t *testing.T) {
	params := &stripe.TaxIDParams{}
	result, _ := taxid.Del("txi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTaxIDList(t *testing.T) {
	params := &stripe.TaxIDListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := taxid.List(params)
	assert.NotNil(t, result)
}

func TestInvoiceCreate(t *testing.T) {
	params := &stripe.InvoiceParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, _ := invoice.New(params)
	assert.NotNil(t, result)
}

func TestInvoiceRetrieve(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, _ := invoice.Get("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceUpdate(t *testing.T) {
	params := &stripe.InvoiceParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := invoice.Update("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceDelete(t *testing.T) {
	params := &stripe.InvoiceParams{}
	result, _ := invoice.Del("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceFinalizeInvoice(t *testing.T) {
	params := &stripe.InvoiceFinalizeParams{}
	result, _ := invoice.FinalizeInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoicePay(t *testing.T) {
	params := &stripe.InvoicePayParams{}
	result, _ := invoice.Pay("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceSendInvoice(t *testing.T) {
	params := &stripe.InvoiceSendParams{}
	result, _ := invoice.SendInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceVoidInvoice(t *testing.T) {
	params := &stripe.InvoiceVoidParams{}
	result, _ := invoice.VoidInvoice("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceMarkUncollectible(t *testing.T) {
	params := &stripe.InvoiceMarkUncollectibleParams{}
	result, _ := invoice.MarkUncollectible("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceList(t *testing.T) {
	params := &stripe.InvoiceListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := invoice.List(params)
	assert.NotNil(t, result)
}

func TestInvoiceItemCreate(t *testing.T) {
	params := &stripe.InvoiceItemParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Price:    stripe.String("price_xxxxxxxxxxxxx"),
	}
	result, _ := invoiceitem.New(params)
	assert.NotNil(t, result)
}

func TestInvoiceItemRetrieve(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	result, _ := invoiceitem.Get("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceItemUpdate(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := invoiceitem.Update("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceItemDelete(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	result, _ := invoiceitem.Del("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceItemList(t *testing.T) {
	params := &stripe.InvoiceItemListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := invoiceitem.List(params)
	assert.NotNil(t, result)
}

func TestPlanCreate(t *testing.T) {
	params := &stripe.PlanParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Interval: stripe.String(string(stripe.PlanIntervalMonth)),
		Product:  &stripe.PlanProductParams{ID: stripe.String("prod_xxxxxxxxxxxxx")},
	}
	result, _ := plan.New(params)
	assert.NotNil(t, result)
}

func TestPlanRetrieve(t *testing.T) {
	params := &stripe.PlanParams{}
	result, _ := plan.Get("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPlanUpdate(t *testing.T) {
	params := &stripe.PlanParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := plan.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPlanDelete(t *testing.T) {
	params := &stripe.PlanParams{}
	result, _ := plan.Del("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPlanList(t *testing.T) {
	params := &stripe.PlanListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := plan.List(params)
	assert.NotNil(t, result)
}

func TestPromotionCodeCreate(t *testing.T) {
	params := &stripe.PromotionCodeParams{Coupon: stripe.String("25_5OFF")}
	result, _ := promotioncode.New(params)
	assert.NotNil(t, result)
}

func TestPromotionCodeUpdate(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := promotioncode.Update("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPromotionCodeRetrieve(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	result, _ := promotioncode.Get("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPromotionCodeList(t *testing.T) {
	params := &stripe.PromotionCodeListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := promotioncode.List(params)
	assert.NotNil(t, result)
}

func TestTaxRateCreate(t *testing.T) {
	params := &stripe.TaxRateParams{
		DisplayName:  stripe.String("VAT"),
		Description:  stripe.String("VAT Germany"),
		Jurisdiction: stripe.String("DE"),
		Percentage:   stripe.Float64(16),
		Inclusive:    stripe.Bool(false),
	}
	result, _ := taxrate.New(params)
	assert.NotNil(t, result)
}

func TestTaxRateRetrieve(t *testing.T) {
	params := &stripe.TaxRateParams{}
	result, _ := taxrate.Get("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTaxRateUpdate(t *testing.T) {
	params := &stripe.TaxRateParams{Active: stripe.Bool(false)}
	result, _ := taxrate.Update("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTaxRateList(t *testing.T) {
	params := &stripe.TaxRateListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := taxrate.List(params)
	assert.NotNil(t, result)
}

func TestUsageRecordCreate(t *testing.T) {
	params := &stripe.UsageRecordParams{
		Quantity:  stripe.Int64(100),
		Timestamp: stripe.Int64(1571252444),
	}
	result, _ := usagerecord.New(params)
	assert.NotNil(t, result)
}

func TestUsageRecordSummaryList(t *testing.T) {
	params := &stripe.UsageRecordSummaryListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := usagerecordsummary.List(params)
	assert.NotNil(t, result)
}

func TestAccountCreate(t *testing.T) {
	params := &stripe.AccountParams{
		Type:    stripe.String(string(stripe.AccountTypeCustom)),
		Country: stripe.String("US"),
		Email:   stripe.String("jenny.rosen@example.com"),
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
	}
	result, _ := account.New(params)
	assert.NotNil(t, result)
}

func TestAccountRetrieve(t *testing.T) {
	params := &stripe.AccountParams{}
	result, _ := account.GetByID("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountUpdate(t *testing.T) {
	params := &stripe.AccountParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := account.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountDelete(t *testing.T) {
	params := &stripe.AccountParams{}
	result, _ := account.Del("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountReject(t *testing.T) {
	params := &stripe.AccountRejectParams{Reason: stripe.String("fraud")}
	result, _ := account.Reject("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountList(t *testing.T) {
	params := &stripe.AccountListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := account.List(params)
	assert.NotNil(t, result)
}

func TestAccountLinkCreate(t *testing.T) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String("acct_xxxxxxxxxxxxx"),
		RefreshURL: stripe.String("https://example.com/reauth"),
		ReturnURL:  stripe.String("https://example.com/return"),
		Type:       stripe.String("account_onboarding"),
	}
	result, _ := accountlink.New(params)
	assert.NotNil(t, result)
}

func TestCapabilityRetrieve(t *testing.T) {
	params := &stripe.CapabilityParams{}
	result, _ := capability.Get("card_payments", params)
	assert.NotNil(t, result)
}

func TestCapabilityUpdate(t *testing.T) {
	params := &stripe.CapabilityParams{Requested: stripe.Bool(true)}
	result, _ := capability.Update("card_payments", params)
	assert.NotNil(t, result)
}

func TestCapabilityList(t *testing.T) {
	params := &stripe.CapabilityListParams{}
	result := capability.List(params)
	assert.NotNil(t, result)
}

func TestCountrySpecList(t *testing.T) {
	params := &stripe.CountrySpecListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := countryspec.List(params)
	assert.NotNil(t, result)
}

func TestCountrySpecRetrieve(t *testing.T) {
	params := &stripe.CountrySpecParams{}
	result, _ := countryspec.Get("US", params)
	assert.NotNil(t, result)
}

func TestPersonCreate(t *testing.T) {
	params := &stripe.PersonParams{
		FirstName: stripe.String("Jane"),
		LastName:  stripe.String("Diaz"),
	}
	result, _ := person.New(params)
	assert.NotNil(t, result)
}

func TestPersonUpdate(t *testing.T) {
	params := &stripe.PersonParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := person.Update("person_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPersonList(t *testing.T) {
	params := &stripe.PersonListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := person.List(params)
	assert.NotNil(t, result)
}

func TestTopupCreate(t *testing.T) {
	params := &stripe.TopupParams{
		Amount:              stripe.Int64(2000),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		Description:         stripe.String("Top-up for Jenny Rosen"),
		StatementDescriptor: stripe.String("Top-up"),
	}
	result, _ := topup.New(params)
	assert.NotNil(t, result)
}

func TestTopupRetrieve(t *testing.T) {
	params := &stripe.TopupParams{}
	result, _ := topup.Get("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTopupUpdate(t *testing.T) {
	params := &stripe.TopupParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := topup.Update("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTopupList(t *testing.T) {
	params := &stripe.TopupListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := topup.List(params)
	assert.NotNil(t, result)
}

func TestTransferCreate(t *testing.T) {
	params := &stripe.TransferParams{
		Amount:        stripe.Int64(400),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Destination:   stripe.String("acct_xxxxxxxxxxxxx"),
		TransferGroup: stripe.String("ORDER_95"),
	}
	result, _ := transfer.New(params)
	assert.NotNil(t, result)
}

func TestTransferRetrieve(t *testing.T) {
	params := &stripe.TransferParams{}
	result, _ := transfer.Get("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTransferUpdate(t *testing.T) {
	params := &stripe.TransferParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := transfer.Update("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTransferList(t *testing.T) {
	params := &stripe.TransferListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := transfer.List(params)
	assert.NotNil(t, result)
}

func TestReversalCreate(t *testing.T) {
	params := &stripe.ReversalParams{Amount: stripe.Int64(100)}
	result, _ := reversal.New(params)
	assert.NotNil(t, result)
}

func TestReversalRetrieve(t *testing.T) {
	params := &stripe.ReversalParams{}
	result, _ := reversal.Get("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestReversalUpdate(t *testing.T) {
	params := &stripe.ReversalParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := reversal.Update("trr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestReversalList(t *testing.T) {
	params := &stripe.ReversalListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := reversal.List(params)
	assert.NotNil(t, result)
}

func TestRadarEarlyFraudWarningRetrieve(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningParams{}
	result, _ := radar_earlyfraudwarning.Get("issfr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestRadarEarlyFraudWarningList(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := radar_earlyfraudwarning.List(params)
	assert.NotNil(t, result)
}

func TestReviewApprove(t *testing.T) {
	params := &stripe.ReviewApproveParams{}
	result, _ := review.Approve("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestReviewRetrieve(t *testing.T) {
	params := &stripe.ReviewParams{}
	result, _ := review.Get("prv_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestReviewList(t *testing.T) {
	params := &stripe.ReviewListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := review.List(params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationRetrieve(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	result, _ := issuing_authorization.Get("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationUpdate(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := issuing_authorization.Update("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationApprove(t *testing.T) {
	params := &stripe.IssuingAuthorizationApproveParams{}
	result, _ := issuing_authorization.Approve("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationDecline(t *testing.T) {
	params := &stripe.IssuingAuthorizationDeclineParams{}
	result, _ := issuing_authorization.Decline("iauth_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationList(t *testing.T) {
	params := &stripe.IssuingAuthorizationListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := issuing_authorization.List(params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderCreate(t *testing.T) {
	params := &stripe.IssuingCardholderParams{
		Type:        stripe.String(string(stripe.IssuingCardholderTypeIndividual)),
		Name:        stripe.String("Jenny Rosen"),
		Email:       stripe.String("jenny.rosen@example.com"),
		PhoneNumber: stripe.String("+18888675309"),
		Billing: &stripe.IssuingCardholderBillingParams{
			Address: &stripe.AddressParams{
				Line1:      stripe.String("1234 Main Street"),
				City:       stripe.String("San Francisco"),
				State:      stripe.String("CA"),
				Country:    stripe.String("US"),
				PostalCode: stripe.String("94111"),
			},
		},
	}
	result, _ := issuing_cardholder.New(params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderRetrieve(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	result, _ := issuing_cardholder.Get("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderUpdate(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := issuing_cardholder.Update("ich_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderList(t *testing.T) {
	params := &stripe.IssuingCardholderListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := issuing_cardholder.List(params)
	assert.NotNil(t, result)
}

func TestIssuingCardCreate(t *testing.T) {
	params := &stripe.IssuingCardParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Type:       stripe.String(string(stripe.IssuingCardTypeVirtual)),
	}
	result, _ := issuing_card.New(params)
	assert.NotNil(t, result)
}

func TestIssuingCardRetrieve(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	result, _ := issuing_card.Get("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingCardUpdate(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := issuing_card.Update("ic_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingCardList(t *testing.T) {
	params := &stripe.IssuingCardListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := issuing_card.List(params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeCreate(t *testing.T) {
	params := &stripe.IssuingDisputeParams{
		Transaction: stripe.String("ipi_xxxxxxxxxxxxx"),
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Reason: stripe.String(string(stripe.IssuingDisputeEvidenceReasonFraudulent)),
			Fraudulent: &stripe.IssuingDisputeEvidenceFraudulentParams{
				Explanation: stripe.String("Purchase was unrecognized."),
			},
		},
	}
	result, _ := issuing_dispute.New(params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeSubmit(t *testing.T) {
	params := &stripe.IssuingDisputeSubmitParams{}
	result, _ := issuing_dispute.Submit("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeRetrieve(t *testing.T) {
	params := &stripe.IssuingDisputeParams{}
	result, _ := issuing_dispute.Get("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeUpdate(t *testing.T) {
	params := &stripe.IssuingDisputeParams{
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Reason: stripe.String(string(stripe.IssuingDisputeEvidenceReasonNotReceived)),
			NotReceived: &stripe.IssuingDisputeEvidenceNotReceivedParams{
				ExpectedAt:         stripe.Int64(1590000000),
				Explanation:        stripe.String(""),
				ProductDescription: stripe.String("Baseball cap"),
				ProductType:        stripe.String(string(stripe.IssuingDisputeEvidenceNotReceivedProductTypeMerchandise)),
			},
		},
	}
	result, _ := issuing_dispute.Update("idp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeList(t *testing.T) {
	params := &stripe.IssuingDisputeListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := issuing_dispute.List(params)
	assert.NotNil(t, result)
}

func TestIssuingTransactionRetrieve(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	result, _ := issuing_transaction.Get("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingTransactionUpdate(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := issuing_transaction.Update("ipi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestIssuingTransactionList(t *testing.T) {
	params := &stripe.IssuingTransactionListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := issuing_transaction.List(params)
	assert.NotNil(t, result)
}

func TestTerminalConnectionTokenCreate(t *testing.T) {
	params := &stripe.TerminalConnectionTokenParams{}
	result, _ := terminal_connectiontoken.New(params)
	assert.NotNil(t, result)
}

func TestTerminalLocationCreate(t *testing.T) {
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
		Address: &stripe.AccountAddressParams{
			Line1:      stripe.String("1234 Main Street"),
			City:       stripe.String("San Francisco"),
			Country:    stripe.String("US"),
			PostalCode: stripe.String("94111"),
		},
	}
	result, _ := terminal_location.New(params)
	assert.NotNil(t, result)
}

func TestTerminalLocationRetrieve(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, _ := terminal_location.Get("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalLocationUpdate(t *testing.T) {
	params := &stripe.TerminalLocationParams{
		DisplayName: stripe.String("My First Store"),
	}
	result, _ := terminal_location.Update("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalLocationDelete(t *testing.T) {
	params := &stripe.TerminalLocationParams{}
	result, _ := terminal_location.Del("tml_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTerminalLocationList(t *testing.T) {
	params := &stripe.TerminalLocationListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := terminal_location.List(params)
	assert.NotNil(t, result)
}

func TestTerminalReaderCreate(t *testing.T) {
	params := &stripe.TerminalReaderParams{
		RegistrationCode: stripe.String("puppies-plug-could"),
		Label:            stripe.String("Blue Rabbit"),
		Location:         stripe.String("tml_1234"),
	}
	result, _ := terminal_reader.New(params)
	assert.NotNil(t, result)
}

func TestTerminalReaderUpdate(t *testing.T) {
	params := &stripe.TerminalReaderParams{Label: stripe.String("Blue Rabbit")}
	result, _ := terminal_reader.Update("tmr_P400-123-456-789", params)
	assert.NotNil(t, result)
}

func TestTerminalReaderDelete(t *testing.T) {
	params := &stripe.TerminalReaderParams{}
	result, _ := terminal_reader.Del("tmr_P400-123-456-789", params)
	assert.NotNil(t, result)
}

func TestTerminalReaderList(t *testing.T) {
	params := &stripe.TerminalReaderListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := terminal_reader.List(params)
	assert.NotNil(t, result)
}

func TestOrderCreate(t *testing.T) {
	params := &stripe.OrderParams{
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Email:    stripe.String("jenny.rosen@example.com"),
		Items: []*stripe.OrderItemParams{
			&stripe.OrderItemParams{
				Type:   stripe.String("sku"),
				Parent: stripe.String("sku_xxxxxxxxxxxxx"),
			},
		},
		Shipping: &stripe.ShippingParams{
			Name: stripe.String("Jenny Rosen"),
			Address: &stripe.AddressParams{
				Line1:      stripe.String("1234 Main Street"),
				City:       stripe.String("San Francisco"),
				State:      stripe.String("CA"),
				Country:    stripe.String("US"),
				PostalCode: stripe.String("94111"),
			},
		},
	}
	result, _ := order.New(params)
	assert.NotNil(t, result)
}

func TestOrderRetrieve(t *testing.T) {
	params := &stripe.OrderParams{}
	result, _ := order.Get("or_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestOrderUpdate(t *testing.T) {
	params := &stripe.OrderUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := order.Update("or_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestOrderPay(t *testing.T) {
	params := &stripe.OrderPayParams{
		Source: &stripe.SourceParams{Token: stripe.String("tok_xxxx")},
	}
	result, _ := order.Pay("or_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestOrderList(t *testing.T) {
	params := &stripe.OrderListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := order.List(params)
	assert.NotNil(t, result)
}

func TestOrderReturnRetrieve(t *testing.T) {
	params := &stripe.OrderReturnParams{}
	result, _ := orderreturn.Get("orret_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestOrderReturnList(t *testing.T) {
	params := &stripe.OrderReturnListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := orderreturn.List(params)
	assert.NotNil(t, result)
}

func TestSKUCreate(t *testing.T) {
	params := &stripe.SKUParams{
		Attributes: map[string]string{"size": "Medium", "gender": "Unisex"},
		Price:      stripe.Int64(1500),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Inventory: &stripe.InventoryParams{
			Type:     stripe.String(string(stripe.SKUInventoryTypeFinite)),
			Quantity: stripe.Int64(500),
		},
		Product: stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, _ := sku.New(params)
	assert.NotNil(t, result)
}

func TestSKURetrieve(t *testing.T) {
	params := &stripe.SKUParams{}
	result, _ := sku.Get("sku_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSKUUpdate(t *testing.T) {
	params := &stripe.SKUParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := sku.Update("sku_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSKUList(t *testing.T) {
	params := &stripe.SKUListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := sku.List(params)
	assert.NotNil(t, result)
}

func TestSKUDelete(t *testing.T) {
	params := &stripe.SKUParams{}
	result, _ := sku.Del("sku_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSigmaScheduledQueryRunRetrieve(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunParams{}
	result, _ := sigma_scheduledqueryrun.Get("sqr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSigmaScheduledQueryRunList(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := sigma_scheduledqueryrun.List(params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointCreate(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/my/webhook/endpoint"),
		EnabledEvents: []*string{
			stripe.String("charge.failed"),
			stripe.String("charge.succeeded"),
		},
	}
	result, _ := webhookendpoint.New(params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointRetrieve(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, _ := webhookendpoint.Get("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointUpdate(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/new_endpoint"),
	}
	result, _ := webhookendpoint.Update("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointList(t *testing.T) {
	params := &stripe.WebhookEndpointListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := webhookendpoint.List(params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointDelete(t *testing.T) {
	params := &stripe.WebhookEndpointParams{}
	result, _ := webhookendpoint.Del("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerListPaymentMethods(t *testing.T) {
	params := &stripe.CustomerListPaymentMethodsParams{
		Type: stripe.String("card"),
	}
	result := customer.ListPaymentMethods(params)
	assert.NotNil(t, result)
}

func TestCheckoutSessionExpire(t *testing.T) {
	params := &stripe.CheckoutSessionExpireParams{}
	result, _ := checkout_session.Expire("sess_xyz", params)
	assert.NotNil(t, result)
}

func TestShippingRateCreate(t *testing.T) {
	params := &stripe.ShippingRateParams{
		DisplayName: stripe.String("Sample Shipper"),
		FixedAmount: &stripe.ShippingRateFixedAmountParams{
			Currency: stripe.String(string(stripe.CurrencyUSD)),
			Amount:   stripe.Int64(400),
		},
		Type: stripe.String("fixed_amount"),
	}
	result, _ := shippingrate.New(params)
	assert.NotNil(t, result)
}

func TestShippingRateList(t *testing.T) {
	params := &stripe.ShippingRateListParams{}
	result := shippingrate.List(params)
	assert.NotNil(t, result)
}

func TestPaymentIntentCreate2(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(1099),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, _ := paymentintent.New(params)
	assert.NotNil(t, result)
}

func TestPaymentLinkCreate(t *testing.T) {
	params := &stripe.PaymentLinkParams{
		LineItems: []*stripe.PaymentLinkLineItemParams{
			&stripe.PaymentLinkLineItemParams{
				Price:    stripe.String("price_xxxxxxxxxxxxx"),
				Quantity: stripe.Int64(1),
			},
		},
	}
	result, _ := paymentlink.New(params)
	assert.NotNil(t, result)
}

func TestPaymentLinkListLineItems(t *testing.T) {
	params := &stripe.PaymentLinkListLineItemsParams{}
	result := paymentlink.ListLineItems(params)
	assert.NotNil(t, result)
}

func TestPaymentLinkRetrieve(t *testing.T) {
	params := &stripe.PaymentLinkParams{}
	result, _ := paymentlink.Get("pl_xyz", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentVerifyMicrodeposits(t *testing.T) {
	params := &stripe.PaymentIntentVerifyMicrodepositsParams{}
	result, _ := paymentintent.VerifyMicrodeposits("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentVerifyMicrodeposits(t *testing.T) {
	params := &stripe.SetupIntentVerifyMicrodepositsParams{}
	result, _ := setupintent.VerifyMicrodeposits("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockCreate(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{
		FrozenTime: stripe.Int64(123),
		Name:       stripe.String("cogsworth"),
	}
	result, _ := testhelpers_testclock.New(params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockRetrieve(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, _ := testhelpers_testclock.Get("clock_xyz", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockList(t *testing.T) {
	params := &stripe.TestHelpersTestClockListParams{}
	result := testhelpers_testclock.List(params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockDelete(t *testing.T) {
	params := &stripe.TestHelpersTestClockParams{}
	result, _ := testhelpers_testclock.Del("clock_xyz", params)
	assert.NotNil(t, result)
}

func TestTestHelpersTestClockAdvance(t *testing.T) {
	params := &stripe.TestHelpersTestClockAdvanceParams{
		FrozenTime: stripe.Int64(142),
	}
	result, _ := testhelpers_testclock.Advance("clock_xyz", params)
	assert.NotNil(t, result)
}

func TestCustomerCreateFundingInstructions(t *testing.T) {
	params := &stripe.CustomerCreateFundingInstructionsParams{
		BankTransfer: &stripe.CustomerCreateFundingInstructionsBankTransferParams{
			RequestedAddressTypes: []*string{stripe.String("zengin")},
			Type:                  stripe.String("jp_bank_transfer"),
		},
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		FundingType: stripe.String("bank_transfer"),
	}
	result, _ := customer.CreateFundingInstructions("cus_123", params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationList(t *testing.T) {
	params := &stripe.TerminalConfigurationListParams{}
	result := terminal_configuration.List(params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationRetrieve(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, _ := terminal_configuration.Get("uc_123", params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationCreate(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, _ := terminal_configuration.New(params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationUpdate(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{
		Tipping: &stripe.TerminalConfigurationTippingParams{
			USD: &stripe.TerminalConfigurationTippingUSDParams{
				FixedAmounts: []*int64{stripe.Int64(10)},
			},
		},
	}
	result, _ := terminal_configuration.Update("uc_123", params)
	assert.NotNil(t, result)
}

func TestTerminalConfigurationDelete(t *testing.T) {
	params := &stripe.TerminalConfigurationParams{}
	result, _ := terminal_configuration.Del("uc_123", params)
	assert.NotNil(t, result)
}

func TestTestHelpersRefundExpire(t *testing.T) {
	params := &stripe.TestHelpersRefundExpireParams{}
	result, _ := testhelpers_refund.Expire("re_123", params)
	assert.NotNil(t, result)
}
