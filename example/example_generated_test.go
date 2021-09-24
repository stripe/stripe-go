package example

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v72"
	account "github.com/stripe/stripe-go/v72/account"
	accountlink "github.com/stripe/stripe-go/v72/accountlink"
	balancetransaction "github.com/stripe/stripe-go/v72/balancetransaction"
	configuration "github.com/stripe/stripe-go/v72/billingportal/configuration"
	billingportal_session "github.com/stripe/stripe-go/v72/billingportal/session"
	capability "github.com/stripe/stripe-go/v72/capability"
	charge "github.com/stripe/stripe-go/v72/charge"
	session "github.com/stripe/stripe-go/v72/checkout/session"
	countryspec "github.com/stripe/stripe-go/v72/countryspec"
	coupon "github.com/stripe/stripe-go/v72/coupon"
	customer "github.com/stripe/stripe-go/v72/customer"
	customerbalancetransaction "github.com/stripe/stripe-go/v72/customerbalancetransaction"
	dispute "github.com/stripe/stripe-go/v72/dispute"
	event "github.com/stripe/stripe-go/v72/event"
	invoice "github.com/stripe/stripe-go/v72/invoice"
	invoiceitem "github.com/stripe/stripe-go/v72/invoiceitem"
	authorization "github.com/stripe/stripe-go/v72/issuing/authorization"
	card "github.com/stripe/stripe-go/v72/issuing/card"
	cardholder "github.com/stripe/stripe-go/v72/issuing/cardholder"
	issuing_dispute "github.com/stripe/stripe-go/v72/issuing/dispute"
	transaction "github.com/stripe/stripe-go/v72/issuing/transaction"
	mandate "github.com/stripe/stripe-go/v72/mandate"
	order "github.com/stripe/stripe-go/v72/order"
	orderreturn "github.com/stripe/stripe-go/v72/orderreturn"
	paymentintent "github.com/stripe/stripe-go/v72/paymentintent"
	paymentmethod "github.com/stripe/stripe-go/v72/paymentmethod"
	payout "github.com/stripe/stripe-go/v72/payout"
	person "github.com/stripe/stripe-go/v72/person"
	plan "github.com/stripe/stripe-go/v72/plan"
	price "github.com/stripe/stripe-go/v72/price"
	product "github.com/stripe/stripe-go/v72/product"
	promotioncode "github.com/stripe/stripe-go/v72/promotioncode"
	earlyfraudwarning "github.com/stripe/stripe-go/v72/radar/earlyfraudwarning"
	refund "github.com/stripe/stripe-go/v72/refund"
	reversal "github.com/stripe/stripe-go/v72/reversal"
	review "github.com/stripe/stripe-go/v72/review"
	setupattempt "github.com/stripe/stripe-go/v72/setupattempt"
	setupintent "github.com/stripe/stripe-go/v72/setupintent"
	scheduledqueryrun "github.com/stripe/stripe-go/v72/sigma/scheduledqueryrun"
	sku "github.com/stripe/stripe-go/v72/sku"
	source "github.com/stripe/stripe-go/v72/source"
	taxid "github.com/stripe/stripe-go/v72/taxid"
	taxrate "github.com/stripe/stripe-go/v72/taxrate"
	location "github.com/stripe/stripe-go/v72/terminal/location"
	reader "github.com/stripe/stripe-go/v72/terminal/reader"
	_ "github.com/stripe/stripe-go/v72/testing"
	topup "github.com/stripe/stripe-go/v72/topup"
	transfer "github.com/stripe/stripe-go/v72/transfer"
	usagerecord "github.com/stripe/stripe-go/v72/usagerecord"
	webhookendpoint "github.com/stripe/stripe-go/v72/webhookendpoint"
)

func TestCustomerList(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := customer.List(params)
	assert.NotNil(t, result)
}

func TestBalanceTransactionRetrieve(t *testing.T) {
	result, _ := balancetransaction.Get("txn_xxxxxxxxxxxxx", nil)
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
	result, _ := charge.Get("ch_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestChargeUpdate(t *testing.T) {
	result, _ := charge.Update("ch_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestChargeCapture(t *testing.T) {
	result, _ := charge.Capture("ch_xxxxxxxxxxxxx", nil)
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
	result, _ := customer.Get("cus_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCustomerUpdate(t *testing.T) {
	result, _ := customer.Update("cus_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCustomerDelete(t *testing.T) {
	result, _ := customer.Del("cus_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCustomerList2(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := customer.List(params)
	assert.NotNil(t, result)
}

func TestDisputeRetrieve(t *testing.T) {
	result, _ := dispute.Get("dp_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestDisputeUpdate(t *testing.T) {
	result, _ := dispute.Update("dp_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestDisputeClose(t *testing.T) {
	result, _ := dispute.Close("dp_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestDisputeList(t *testing.T) {
	params := &stripe.DisputeListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := dispute.List(params)
	assert.NotNil(t, result)
}

func TestEventRetrieve(t *testing.T) {
	result, _ := event.Get("evt_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestEventList(t *testing.T) {
	params := &stripe.EventListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := event.List(params)
	assert.NotNil(t, result)
}

func TestMandateRetrieve(t *testing.T) {
	result, _ := mandate.Get("mandate_xxxxxxxxxxxxx", nil)
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
	result, _ := paymentintent.Get("pi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentIntentUpdate(t *testing.T) {
	result, _ := paymentintent.Update("pi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentIntentConfirm(t *testing.T) {
	result, _ := paymentintent.Confirm("pi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentIntentCapture(t *testing.T) {
	result, _ := paymentintent.Capture("pi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentIntentCancel(t *testing.T) {
	result, _ := paymentintent.Cancel("pi_xxxxxxxxxxxxx", nil)
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
	result, _ := setupintent.Get("seti_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSetupIntentUpdate(t *testing.T) {
	result, _ := setupintent.Update("seti_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSetupIntentConfirm(t *testing.T) {
	result, _ := setupintent.Confirm("seti_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSetupIntentCancel(t *testing.T) {
	result, _ := setupintent.Cancel("seti_xxxxxxxxxxxxx", nil)
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
	result, _ := payout.Get("po_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPayoutUpdate(t *testing.T) {
	result, _ := payout.Update("po_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPayoutList(t *testing.T) {
	params := &stripe.PayoutListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := payout.List(params)
	assert.NotNil(t, result)
}

func TestPayoutCancel(t *testing.T) {
	result, _ := payout.Cancel("po_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPayoutReverse(t *testing.T) {
	result, _ := payout.Reverse("po_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestProductCreate(t *testing.T) {
	params := &stripe.ProductParams{Name: stripe.String("Gold Special")}
	result, _ := product.New(params)
	assert.NotNil(t, result)
}

func TestProductRetrieve(t *testing.T) {
	result, _ := product.Get("prod_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestProductUpdate(t *testing.T) {
	result, _ := product.Update("prod_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestProductList(t *testing.T) {
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := product.List(params)
	assert.NotNil(t, result)
}

func TestProductDelete(t *testing.T) {
	result, _ := product.Del("prod_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPriceCreate(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Recurring:  &stripe.PriceRecurringParams{Interval: stripe.String("month")},
		Product:    stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, _ := price.New(params)
	assert.NotNil(t, result)
}

func TestPriceRetrieve(t *testing.T) {
	result, _ := price.Get("price_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPriceUpdate(t *testing.T) {
	result, _ := price.Update("price_xxxxxxxxxxxxx", nil)
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
	result, _ := refund.Get("re_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestRefundUpdate(t *testing.T) {
	result, _ := refund.Update("re_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestRefundList(t *testing.T) {
	params := &stripe.RefundListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := refund.List(params)
	assert.NotNil(t, result)
}

func TestPaymentMethodRetrieve(t *testing.T) {
	result, _ := paymentmethod.Get("pm_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentMethodUpdate(t *testing.T) {
	result, _ := paymentmethod.Update("pm_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentMethodList(t *testing.T) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String("card"),
	}
	result := paymentmethod.List(params)
	assert.NotNil(t, result)
}

func TestPaymentMethodAttach(t *testing.T) {
	result, _ := paymentmethod.Attach("pm_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentMethodDetach(t *testing.T) {
	result, _ := paymentmethod.Detach("pm_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSourceRetrieve(t *testing.T) {
	result, _ := source.Get("src_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSourceUpdate(t *testing.T) {
	result, _ := source.Update("src_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCheckoutSessionRetrieve(t *testing.T) {
	result, _ := session.Get("cs_test_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCheckoutSessionList(t *testing.T) {
	params := &stripe.CheckoutSessionListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := session.List(params)
	assert.NotNil(t, result)
}

func TestCouponCreate(t *testing.T) {
	params := &stripe.CouponParams{
		PercentOff:       stripe.Float64(25),
		Duration:         stripe.String("repeating"),
		DurationInMonths: stripe.Int64(3),
	}
	result, _ := coupon.New(params)
	assert.NotNil(t, result)
}

func TestCouponRetrieve(t *testing.T) {
	result, _ := coupon.Get("25_5OFF", nil)
	assert.NotNil(t, result)
}

func TestCouponUpdate(t *testing.T) {
	result, _ := coupon.Update("co_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCouponDelete(t *testing.T) {
	result, _ := coupon.Del("co_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCouponList(t *testing.T) {
	params := &stripe.CouponListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := coupon.List(params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionCreate(t *testing.T) {
	result, _ := customerbalancetransaction.New("cus_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionUpdate(t *testing.T) {
	result, _ := customerbalancetransaction.Update("cbtxn_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionList(t *testing.T) {
	result := customerbalancetransaction.List("cus_xxxxxxxxxxxxx", nil)
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
					stripe.String("email"),
					stripe.String("tax_id"),
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
	result, _ := configuration.New(params)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationUpdate(t *testing.T) {
	result, _ := configuration.Update("bpc_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationRetrieve(t *testing.T) {
	result, _ := configuration.Get("bpc_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestBillingPortalConfigurationList(t *testing.T) {
	params := &stripe.BillingPortalConfigurationListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := configuration.List(params)
	assert.NotNil(t, result)
}

func TestTaxIDCreate(t *testing.T) {
	result, _ := taxid.New("cus_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTaxIDList(t *testing.T) {
	result := taxid.List("cus_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceCreate(t *testing.T) {
	params := &stripe.InvoiceParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, _ := invoice.New(params)
	assert.NotNil(t, result)
}

func TestInvoiceRetrieve(t *testing.T) {
	result, _ := invoice.Get("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceUpdate(t *testing.T) {
	result, _ := invoice.Update("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceDelete(t *testing.T) {
	result, _ := invoice.Del("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceFinalizeInvoice(t *testing.T) {
	result, _ := invoice.FinalizeInvoice("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoicePay(t *testing.T) {
	result, _ := invoice.Pay("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceSendInvoice(t *testing.T) {
	result, _ := invoice.SendInvoice("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceVoidInvoice(t *testing.T) {
	result, _ := invoice.VoidInvoice("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceMarkUncollectible(t *testing.T) {
	result, _ := invoice.MarkUncollectible("in_xxxxxxxxxxxxx", nil)
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
	result, _ := invoiceitem.Get("ii_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceItemUpdate(t *testing.T) {
	result, _ := invoiceitem.Update("ii_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceItemDelete(t *testing.T) {
	result, _ := invoiceitem.Del("ii_xxxxxxxxxxxxx", nil)
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
		Interval: stripe.String("month"),
		Product:  &stripe.PlanProductParams{ID: stripe.String("prod_xxxxxxxxxxxxx")},
	}
	result, _ := plan.New(params)
	assert.NotNil(t, result)
}

func TestPlanRetrieve(t *testing.T) {
	result, _ := plan.Get("price_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPlanUpdate(t *testing.T) {
	result, _ := plan.Update("price_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPlanDelete(t *testing.T) {
	result, _ := plan.Del("price_xxxxxxxxxxxxx", nil)
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
	result, _ := promotioncode.Update("promo_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPromotionCodeRetrieve(t *testing.T) {
	result, _ := promotioncode.Get("promo_xxxxxxxxxxxxx", nil)
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
	result, _ := taxrate.Get("txr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTaxRateUpdate(t *testing.T) {
	result, _ := taxrate.Update("txr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTaxRateList(t *testing.T) {
	params := &stripe.TaxRateListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := taxrate.List(params)
	assert.NotNil(t, result)
}

func TestUsageRecordCreate(t *testing.T) {
	result, _ := usagerecord.New("si_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestAccountCreate(t *testing.T) {
	params := &stripe.AccountParams{
		Type:    stripe.String("custom"),
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
	result, _ := account.GetByID("acct_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestAccountUpdate(t *testing.T) {
	result, _ := account.Update("acct_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestAccountDelete(t *testing.T) {
	result, _ := account.Del("acct_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestAccountReject(t *testing.T) {
	result, _ := account.Reject("acct_xxxxxxxxxxxxx", nil)
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

func TestCapabilityUpdate(t *testing.T) {
	result, _ := capability.Update("card_payments", nil)
	assert.NotNil(t, result)
}

func TestCountrySpecList(t *testing.T) {
	params := &stripe.CountrySpecListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := countryspec.List(params)
	assert.NotNil(t, result)
}

func TestCountrySpecRetrieve(t *testing.T) {
	result, _ := countryspec.Get("US", nil)
	assert.NotNil(t, result)
}

func TestPersonCreate(t *testing.T) {
	result, _ := person.New("acct_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPersonUpdate(t *testing.T) {
	result, _ := person.Update("person_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPersonList(t *testing.T) {
	result := person.List("acct_xxxxxxxxxxxxx", nil)
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
	result, _ := topup.Get("tu_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTopupUpdate(t *testing.T) {
	result, _ := topup.Update("tu_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTopupList(t *testing.T) {
	params := &stripe.TopupListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := topup.List(params)
	assert.NotNil(t, result)
}

func TestTopupCancel(t *testing.T) {
	result, _ := topup.Cancel("tu_xxxxxxxxxxxxx", nil)
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
	result, _ := transfer.Get("tr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTransferUpdate(t *testing.T) {
	result, _ := transfer.Update("tr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTransferList(t *testing.T) {
	params := &stripe.TransferListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := transfer.List(params)
	assert.NotNil(t, result)
}

func TestTransferReversalCreate(t *testing.T) {
	result, _ := reversal.New("tr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTransferReversalUpdate(t *testing.T) {
	result, _ := reversal.Update("trr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTransferReversalList(t *testing.T) {
	result := reversal.List("tr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestRadarEarlyFraudWarningRetrieve(t *testing.T) {
	result, _ := earlyfraudwarning.Get("issfr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestRadarEarlyFraudWarningList(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := earlyfraudwarning.List(params)
	assert.NotNil(t, result)
}

func TestReviewApprove(t *testing.T) {
	result, _ := review.Approve("prv_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestReviewRetrieve(t *testing.T) {
	result, _ := review.Get("prv_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestReviewList(t *testing.T) {
	params := &stripe.ReviewListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := review.List(params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationRetrieve(t *testing.T) {
	result, _ := authorization.Get("iauth_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationUpdate(t *testing.T) {
	result, _ := authorization.Update("iauth_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationApprove(t *testing.T) {
	result, _ := authorization.Approve("iauth_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationDecline(t *testing.T) {
	result, _ := authorization.Decline("iauth_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationList(t *testing.T) {
	params := &stripe.IssuingAuthorizationListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := authorization.List(params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderCreate(t *testing.T) {
	params := &stripe.IssuingCardholderParams{
		Type:        stripe.String("individual"),
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
	result, _ := cardholder.New(params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderRetrieve(t *testing.T) {
	result, _ := cardholder.Get("ich_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingCardholderUpdate(t *testing.T) {
	result, _ := cardholder.Update("ich_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingCardholderList(t *testing.T) {
	params := &stripe.IssuingCardholderListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := cardholder.List(params)
	assert.NotNil(t, result)
}

func TestIssuingCardCreate(t *testing.T) {
	params := &stripe.IssuingCardParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Type:       stripe.String("virtual"),
	}
	result, _ := card.New(params)
	assert.NotNil(t, result)
}

func TestIssuingCardRetrieve(t *testing.T) {
	result, _ := card.Get("ic_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingCardUpdate(t *testing.T) {
	result, _ := card.Update("ic_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingCardList(t *testing.T) {
	params := &stripe.IssuingCardListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := card.List(params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeCreate(t *testing.T) {
	params := &stripe.IssuingDisputeParams{
		Transaction: stripe.String("ipi_xxxxxxxxxxxxx"),
		Evidence: &stripe.IssuingDisputeEvidenceParams{
			Reason: stripe.String("fraudulent"),
			Fraudulent: &stripe.IssuingDisputeEvidenceFraudulentParams{
				Explanation: stripe.String("Purchase was unrecognized."),
			},
		},
	}
	result, _ := issuing_dispute.New(params)
	assert.NotNil(t, result)
}

func TestIssuingDisputeSubmit(t *testing.T) {
	result, _ := issuing_dispute.Submit("idp_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingDisputeRetrieve(t *testing.T) {
	result, _ := issuing_dispute.Get("idp_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingDisputeUpdate(t *testing.T) {
	result, _ := issuing_dispute.Update("idp_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingDisputeList(t *testing.T) {
	params := &stripe.IssuingDisputeListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := issuing_dispute.List(params)
	assert.NotNil(t, result)
}

func TestIssuingTransactionRetrieve(t *testing.T) {
	result, _ := transaction.Get("ipi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingTransactionUpdate(t *testing.T) {
	result, _ := transaction.Update("ipi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestIssuingTransactionList(t *testing.T) {
	params := &stripe.IssuingTransactionListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := transaction.List(params)
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
	result, _ := location.New(params)
	assert.NotNil(t, result)
}

func TestTerminalLocationRetrieve(t *testing.T) {
	result, _ := location.Get("tml_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTerminalLocationUpdate(t *testing.T) {
	result, _ := location.Update("tml_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTerminalLocationDelete(t *testing.T) {
	result, _ := location.Del("tml_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTerminalLocationList(t *testing.T) {
	params := &stripe.TerminalLocationListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := location.List(params)
	assert.NotNil(t, result)
}

func TestTerminalReaderCreate(t *testing.T) {
	params := &stripe.TerminalReaderParams{
		RegistrationCode: stripe.String("puppies-plug-could"),
		Label:            stripe.String("Blue Rabbit"),
		Location:         stripe.String("tml_1234"),
	}
	result, _ := reader.New(params)
	assert.NotNil(t, result)
}

func TestTerminalReaderRetrieve(t *testing.T) {
	result, _ := reader.Get("tmr_P400-123-456-789", nil)
	assert.NotNil(t, result)
}

func TestTerminalReaderUpdate(t *testing.T) {
	result, _ := reader.Update("tmr_P400-123-456-789", nil)
	assert.NotNil(t, result)
}

func TestTerminalReaderDelete(t *testing.T) {
	result, _ := reader.Del("tmr_P400-123-456-789", nil)
	assert.NotNil(t, result)
}

func TestTerminalReaderList(t *testing.T) {
	params := &stripe.TerminalReaderListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := reader.List(params)
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
	result, _ := order.Get("or_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestOrderUpdate(t *testing.T) {
	result, _ := order.Update("or_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestOrderPay(t *testing.T) {
	result, _ := order.Pay("or_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestOrderList(t *testing.T) {
	params := &stripe.OrderListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := order.List(params)
	assert.NotNil(t, result)
}

func TestOrderReturnRetrieve(t *testing.T) {
	result, _ := orderreturn.Get("orret_xxxxxxxxxxxxx", nil)
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
	result, _ := sku.Get("sku_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSKUUpdate(t *testing.T) {
	result, _ := sku.Update("sku_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSKUList(t *testing.T) {
	params := &stripe.SKUListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := sku.List(params)
	assert.NotNil(t, result)
}

func TestSKUDelete(t *testing.T) {
	result, _ := sku.Del("sku_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSigmaScheduledQueryRunRetrieve(t *testing.T) {
	result, _ := scheduledqueryrun.Get("sqr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSigmaScheduledQueryRunList(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := scheduledqueryrun.List(params)
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
	result, _ := webhookendpoint.Get("we_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestWebhookEndpointUpdate(t *testing.T) {
	result, _ := webhookendpoint.Update("we_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestWebhookEndpointList(t *testing.T) {
	params := &stripe.WebhookEndpointListParams{}
	params.Filters.AddFilter("Limit", "", "3")
	result := webhookendpoint.List(params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointDelete(t *testing.T) {
	result, _ := webhookendpoint.Del("we_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}
