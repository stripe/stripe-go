package example

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v72"
	account "github.com/stripe/stripe-go/v72/account"
	accountlink "github.com/stripe/stripe-go/v72/accountlink"
	balancetransaction "github.com/stripe/stripe-go/v72/balancetransaction"
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
	reader "github.com/stripe/stripe-go/v72/terminal/reader"
	_ "github.com/stripe/stripe-go/v72/testing"
	topup "github.com/stripe/stripe-go/v72/topup"
	transfer "github.com/stripe/stripe-go/v72/transfer"
	usagerecord "github.com/stripe/stripe-go/v72/usagerecord"
	webhookendpoint "github.com/stripe/stripe-go/v72/webhookendpoint"
)

func TestCustomerList1(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := customer.List(params)
	assert.NotNil(t, result)
}

func TestBalanceTransactionRetrieve1(t *testing.T) {
	result, _ := balancetransaction.Get("txn_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestBalanceTransactionList1(t *testing.T) {
	params := &stripe.BalanceTransactionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := balancetransaction.List(params)
	assert.NotNil(t, result)
}

func TestChargeCreate1(t *testing.T) {
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(2000),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Source:      &stripe.SourceParams{Token: stripe.String("tok_xxxx")},
		Description: stripe.String("My First Test Charge (created for API docs)"),
	}
	result, _ := charge.New(params)
	assert.NotNil(t, result)
}

func TestChargeRetrieve1(t *testing.T) {
	result, _ := charge.Get("ch_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestChargeUpdate1(t *testing.T) {
	params := &stripe.ChargeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := charge.Update("ch_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestChargeCapture1(t *testing.T) {
	result, _ := charge.Capture("ch_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestChargeList1(t *testing.T) {
	params := &stripe.ChargeListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := charge.List(params)
	assert.NotNil(t, result)
}

func TestCustomerCreate1(t *testing.T) {
	params := &stripe.CustomerParams{
		Description: stripe.String("My First Test Customer (created for API docs)"),
	}
	result, _ := customer.New(params)
	assert.NotNil(t, result)
}

func TestCustomerRetrieve1(t *testing.T) {
	result, _ := customer.Get("cus_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCustomerUpdate1(t *testing.T) {
	params := &stripe.CustomerParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := customer.Update("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerDelete1(t *testing.T) {
	result, _ := customer.Del("cus_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCustomerList2(t *testing.T) {
	params := &stripe.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := customer.List(params)
	assert.NotNil(t, result)
}

func TestDisputeRetrieve1(t *testing.T) {
	result, _ := dispute.Get("dp_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestDisputeUpdate1(t *testing.T) {
	params := &stripe.DisputeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := dispute.Update("dp_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestDisputeClose1(t *testing.T) {
	result, _ := dispute.Close("dp_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestDisputeList1(t *testing.T) {
	params := &stripe.DisputeListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := dispute.List(params)
	assert.NotNil(t, result)
}

func TestEventRetrieve1(t *testing.T) {
	result, _ := event.Get("evt_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestEventList1(t *testing.T) {
	params := &stripe.EventListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := event.List(params)
	assert.NotNil(t, result)
}

// TODO File test cases skipped

// TODO File test cases skipped

// TODO File test cases skipped

// TODO File test cases skipped

// TODO File test cases skipped

// TODO File test cases skipped

// TODO File test cases skipped

func TestMandateRetrieve1(t *testing.T) {
	result, _ := mandate.Get("mandate_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentIntentCreate1(t *testing.T) {
	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(2000),
		Currency:           stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, _ := paymentintent.New(params)
	assert.NotNil(t, result)
}

func TestPaymentIntentRetrieve1(t *testing.T) {
	result, _ := paymentintent.Get("pi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentIntentUpdate1(t *testing.T) {
	params := &stripe.PaymentIntentParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := paymentintent.Update("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentConfirm1(t *testing.T) {
	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, _ := paymentintent.Confirm("pi_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentIntentCapture1(t *testing.T) {
	result, _ := paymentintent.Capture("pi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentIntentCancel1(t *testing.T) {
	result, _ := paymentintent.Cancel("pi_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentIntentList1(t *testing.T) {
	params := &stripe.PaymentIntentListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := paymentintent.List(params)
	assert.NotNil(t, result)
}

func TestSetupIntentCreate1(t *testing.T) {
	params := &stripe.SetupIntentParams{
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, _ := setupintent.New(params)
	assert.NotNil(t, result)
}

func TestSetupIntentRetrieve1(t *testing.T) {
	result, _ := setupintent.Get("seti_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSetupIntentUpdate1(t *testing.T) {
	params := &stripe.SetupIntentParams{}
	params.AddMetadata("user_id", "3435453")
	result, _ := setupintent.Update("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentConfirm1(t *testing.T) {
	params := &stripe.SetupIntentConfirmParams{
		PaymentMethod: stripe.String("pm_card_visa"),
	}
	result, _ := setupintent.Confirm("seti_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSetupIntentCancel1(t *testing.T) {
	result, _ := setupintent.Cancel("seti_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSetupIntentList1(t *testing.T) {
	params := &stripe.SetupIntentListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := setupintent.List(params)
	assert.NotNil(t, result)
}

func TestSetupAttemptList1(t *testing.T) {
	params := &stripe.SetupAttemptListParams{
		SetupIntent: stripe.String("seti_xxxxxxxxxxxxx"),
	}
	params.Filters.AddFilter("limit", "", "3")
	result := setupattempt.List(params)
	assert.NotNil(t, result)
}

func TestPayoutCreate1(t *testing.T) {
	params := &stripe.PayoutParams{
		Amount:   stripe.Int64(1100),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	result, _ := payout.New(params)
	assert.NotNil(t, result)
}

func TestPayoutRetrieve1(t *testing.T) {
	result, _ := payout.Get("po_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPayoutUpdate1(t *testing.T) {
	params := &stripe.PayoutParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := payout.Update("po_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPayoutList1(t *testing.T) {
	params := &stripe.PayoutListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := payout.List(params)
	assert.NotNil(t, result)
}

func TestPayoutCancel1(t *testing.T) {
	result, _ := payout.Cancel("po_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPayoutReverse1(t *testing.T) {
	result, _ := payout.Reverse("po_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestProductCreate1(t *testing.T) {
	params := &stripe.ProductParams{Name: stripe.String("Gold Special")}
	result, _ := product.New(params)
	assert.NotNil(t, result)
}

func TestProductRetrieve1(t *testing.T) {
	result, _ := product.Get("prod_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestProductUpdate1(t *testing.T) {
	params := &stripe.ProductParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := product.Update("prod_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestProductList1(t *testing.T) {
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := product.List(params)
	assert.NotNil(t, result)
}

func TestProductDelete1(t *testing.T) {
	result, _ := product.Del("prod_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPriceCreate1(t *testing.T) {
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(2000),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Recurring:  &stripe.PriceRecurringParams{Interval: stripe.String("month")},
		Product:    stripe.String("prod_xxxxxxxxxxxxx"),
	}
	result, _ := price.New(params)
	assert.NotNil(t, result)
}

func TestPriceRetrieve1(t *testing.T) {
	result, _ := price.Get("price_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPriceUpdate1(t *testing.T) {
	params := &stripe.PriceParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := price.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPriceList1(t *testing.T) {
	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := price.List(params)
	assert.NotNil(t, result)
}

func TestRefundCreate1(t *testing.T) {
	params := &stripe.RefundParams{Charge: stripe.String("ch_xxxxxxxxxxxxx")}
	result, _ := refund.New(params)
	assert.NotNil(t, result)
}

func TestRefundRetrieve1(t *testing.T) {
	result, _ := refund.Get("re_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestRefundUpdate1(t *testing.T) {
	params := &stripe.RefundParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := refund.Update("re_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestRefundList1(t *testing.T) {
	params := &stripe.RefundListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := refund.List(params)
	assert.NotNil(t, result)
}

// TODO Token test cases skipped because not all Token methods are supported by GO

// TODO Token test cases skipped because not all Token methods are supported by GO

// TODO Token test cases skipped because not all Token methods are supported by GO

// TODO Token test cases skipped because not all Token methods are supported by GO

// TODO Token test cases skipped because not all Token methods are supported by GO

// TODO Token test cases skipped because not all Token methods are supported by GO

// TODO Token test cases skipped because not all Token methods are supported by GO

// TODO Invalid CC details data from mock

func TestPaymentMethodRetrieve1(t *testing.T) {
	result, _ := paymentmethod.Get("pm_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPaymentMethodUpdate1(t *testing.T) {
	params := &stripe.PaymentMethodParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := paymentmethod.Update("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodList1(t *testing.T) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Type:     stripe.String("card"),
	}
	result := paymentmethod.List(params)
	assert.NotNil(t, result)
}

func TestPaymentMethodAttach1(t *testing.T) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
	}
	result, _ := paymentmethod.Attach("pm_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPaymentMethodDetach1(t *testing.T) {
	result, _ := paymentmethod.Detach("pm_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

// TODO Payment source test cases skipped

// TODO Payment source test cases skipped

// TODO Test cases skipped

// TODO Test cases skipped

// TODO Test cases skipped

// TODO Payment source test cases skipped

// TODO Payment source test cases skipped

// TODO Payment source test cases skipped

// TODO Test cases skipped

// TODO Test cases skipped

// TODO Payment source test cases skipped

func TestSourceRetrieve1(t *testing.T) {
	result, _ := source.Get("src_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSourceUpdate1(t *testing.T) {
	params := &stripe.SourceObjectParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := source.Update("src_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

// TODO Checkout session create test cases skipped

func TestCheckoutSessionRetrieve1(t *testing.T) {
	result, _ := session.Get("sessions", nil)
	assert.NotNil(t, result)
}

func TestCheckoutSessionList1(t *testing.T) {
	params := &stripe.CheckoutSessionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := session.List(params)
	assert.NotNil(t, result)
}

func TestCouponCreate1(t *testing.T) {
	params := &stripe.CouponParams{
		PercentOff:       stripe.Float64(25),
		Duration:         stripe.String("repeating"),
		DurationInMonths: stripe.Int64(3),
	}
	result, _ := coupon.New(params)
	assert.NotNil(t, result)
}

func TestCouponRetrieve1(t *testing.T) {
	result, _ := coupon.Get("25_5OFF", nil)
	assert.NotNil(t, result)
}

func TestCouponUpdate1(t *testing.T) {
	params := &stripe.CouponParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := coupon.Update("co_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCouponDelete1(t *testing.T) {
	result, _ := coupon.Del("co_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestCouponList1(t *testing.T) {
	params := &stripe.CouponListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := coupon.List(params)
	assert.NotNil(t, result)
}

// TODO CreditNote test cases skipped

// TODO CreditNote test cases skipped

// TODO CreditNote test cases skipped

// TODO CreditNote test cases skipped

// TODO CreditNote test cases skipped

// TODO CreditNote test cases skipped

// TODO CreditNote test cases skipped

// TODO CreditNote test cases skipped

func TestCustomerBalanceTransactionCreate1(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{
		Amount:   stripe.Int64(-500),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	result, _ := customerbalancetransaction.New(params)
	assert.NotNil(t, result)
}

// TODO Test cases skipped

func TestCustomerBalanceTransactionUpdate1(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := customerbalancetransaction.Update("cus_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestCustomerBalanceTransactionList1(t *testing.T) {
	params := &stripe.CustomerBalanceTransactionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := customerbalancetransaction.List(params)
	assert.NotNil(t, result)
}

// TODO Billing portal test cases skipped

// TODO Billing portal test cases skipped

// TODO Billing portal test cases skipped

// TODO Billing portal test cases skipped

// TODO Billing portal test cases skipped

func TestTaxIDCreate1(t *testing.T) {
	params := &stripe.TaxIDParams{
		Type:  stripe.String("eu_vat"),
		Value: stripe.String("DE123456789"),
	}
	result, _ := taxid.New(params)
	assert.NotNil(t, result)
}

// TODO Test cases skipped

// TODO Test cases skipped

func TestTaxIDList1(t *testing.T) {
	params := &stripe.TaxIDListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := taxid.List(params)
	assert.NotNil(t, result)
}

func TestInvoiceCreate1(t *testing.T) {
	params := &stripe.InvoiceParams{Customer: stripe.String("cus_xxxxxxxxxxxxx")}
	result, _ := invoice.New(params)
	assert.NotNil(t, result)
}

func TestInvoiceRetrieve1(t *testing.T) {
	result, _ := invoice.Get("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceUpdate1(t *testing.T) {
	params := &stripe.InvoiceParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := invoice.Update("in_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceDelete1(t *testing.T) {
	result, _ := invoice.Del("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceFinalizeInvoice1(t *testing.T) {
	result, _ := invoice.FinalizeInvoice("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoicePay1(t *testing.T) {
	result, _ := invoice.Pay("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceSendInvoice1(t *testing.T) {
	result, _ := invoice.SendInvoice("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceVoidInvoice1(t *testing.T) {
	result, _ := invoice.VoidInvoice("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceMarkUncollectible1(t *testing.T) {
	result, _ := invoice.MarkUncollectible("in_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceList1(t *testing.T) {
	params := &stripe.InvoiceListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := invoice.List(params)
	assert.NotNil(t, result)
}

func TestInvoiceItemCreate1(t *testing.T) {
	params := &stripe.InvoiceItemParams{
		Customer: stripe.String("cus_xxxxxxxxxxxxx"),
		Price:    stripe.String("price_xxxxxxxxxxxxx"),
	}
	result, _ := invoiceitem.New(params)
	assert.NotNil(t, result)
}

func TestInvoiceItemRetrieve1(t *testing.T) {
	result, _ := invoiceitem.Get("ii_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceItemUpdate1(t *testing.T) {
	params := &stripe.InvoiceItemParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := invoiceitem.Update("ii_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestInvoiceItemDelete1(t *testing.T) {
	result, _ := invoiceitem.Del("ii_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestInvoiceItemList1(t *testing.T) {
	params := &stripe.InvoiceItemListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := invoiceitem.List(params)
	assert.NotNil(t, result)
}

func TestPlanCreate1(t *testing.T) {
	params := &stripe.PlanParams{
		Amount:   stripe.Int64(2000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Interval: stripe.String("month"),
		Product:  &stripe.PlanProductParams{ID: stripe.String("prod_xxxxxxxxxxxxx")},
	}
	result, _ := plan.New(params)
	assert.NotNil(t, result)
}

func TestPlanRetrieve1(t *testing.T) {
	result, _ := plan.Get("price_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPlanUpdate1(t *testing.T) {
	params := &stripe.PlanParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := plan.Update("price_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPlanDelete1(t *testing.T) {
	result, _ := plan.Del("price_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPlanList1(t *testing.T) {
	params := &stripe.PlanListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := plan.List(params)
	assert.NotNil(t, result)
}

func TestPromotionCodeCreate1(t *testing.T) {
	params := &stripe.PromotionCodeParams{Coupon: stripe.String("25_5OFF")}
	result, _ := promotioncode.New(params)
	assert.NotNil(t, result)
}

func TestPromotionCodeUpdate1(t *testing.T) {
	params := &stripe.PromotionCodeParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := promotioncode.Update("promo_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestPromotionCodeRetrieve1(t *testing.T) {
	result, _ := promotioncode.Get("promo_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestPromotionCodeList1(t *testing.T) {
	params := &stripe.PromotionCodeListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := promotioncode.List(params)
	assert.NotNil(t, result)
}

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

// TODO Subscription test cases skipped

func TestTaxRateCreate1(t *testing.T) {
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

func TestTaxRateRetrieve1(t *testing.T) {
	result, _ := taxrate.Get("txr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTaxRateUpdate1(t *testing.T) {
	params := &stripe.TaxRateParams{Active: stripe.Bool(false)}
	result, _ := taxrate.Update("txr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTaxRateList1(t *testing.T) {
	params := &stripe.TaxRateListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := taxrate.List(params)
	assert.NotNil(t, result)
}

func TestUsageRecordCreate1(t *testing.T) {
	params := &stripe.UsageRecordParams{
		Quantity:  stripe.Int64(100),
		Timestamp: stripe.Int64(1571252444),
	}
	result, _ := usagerecord.New(params)
	assert.NotNil(t, result)
}

// TODO Subscription test cases skipped

func TestAccountCreate1(t *testing.T) {
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

func TestAccountRetrieve1(t *testing.T) {
	result, _ := account.GetByID("acct_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestAccountUpdate1(t *testing.T) {
	params := &stripe.AccountParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := account.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountDelete1(t *testing.T) {
	result, _ := account.Del("acct_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestAccountReject1(t *testing.T) {
	params := &stripe.AccountRejectParams{Reason: stripe.String("fraud")}
	result, _ := account.Reject("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestAccountList1(t *testing.T) {
	params := &stripe.AccountListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := account.List(params)
	assert.NotNil(t, result)
}

// TODO Login link test cases skipped

func TestAccountLinkCreate1(t *testing.T) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String("acct_xxxxxxxxxxxxx"),
		RefreshURL: stripe.String("https://example.com/reauth"),
		ReturnURL:  stripe.String("https://example.com/return"),
		Type:       stripe.String("account_onboarding"),
	}
	result, _ := accountlink.New(params)
	assert.NotNil(t, result)
}

// TODO Application fee test cases skipped

// TODO Application fee test cases skipped

// TODO Fee refund test cases skipped

// TODO Fee refund test cases skipped

// TODO Fee refund test cases skipped

// TODO Fee refund test cases skipped

// TODO Test cases skipped

func TestCapabilityUpdate1(t *testing.T) {
	params := &stripe.CapabilityParams{Requested: stripe.Bool(true)}
	result, _ := capability.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

// TODO Account capabilities retrieve test case skipped

func TestCountrySpecList1(t *testing.T) {
	params := &stripe.CountrySpecListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := countryspec.List(params)
	assert.NotNil(t, result)
}

func TestCountrySpecRetrieve1(t *testing.T) {
	result, _ := countryspec.Get("US", nil)
	assert.NotNil(t, result)
}

// TODO ExternalAccount test cases skipped

// TODO ExternalAccount test cases skipped

// TODO Test cases skipped

// TODO Test cases skipped

// TODO ExternalAccount test cases skipped

// TODO ExternalAccount test cases skipped

// TODO ExternalAccount test cases skipped

// TODO Test cases skipped

// TODO Test cases skipped

// TODO ExternalAccount test cases skipped

func TestPersonCreate1(t *testing.T) {
	params := &stripe.PersonParams{
		FirstName: stripe.String("Jane"),
		LastName:  stripe.String("Diaz"),
	}
	result, _ := person.New(params)
	assert.NotNil(t, result)
}

// TODO Test cases skipped

func TestPersonUpdate1(t *testing.T) {
	params := &stripe.PersonParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := person.Update("acct_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

// TODO Test cases skipped

func TestPersonList1(t *testing.T) {
	params := &stripe.PersonListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := person.List(params)
	assert.NotNil(t, result)
}

func TestTopupCreate1(t *testing.T) {
	params := &stripe.TopupParams{
		Amount:              stripe.Int64(2000),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		Description:         stripe.String("Top-up for Jenny Rosen"),
		StatementDescriptor: stripe.String("Top-up"),
	}
	result, _ := topup.New(params)
	assert.NotNil(t, result)
}

func TestTopupRetrieve1(t *testing.T) {
	result, _ := topup.Get("tu_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTopupUpdate1(t *testing.T) {
	params := &stripe.TopupParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := topup.Update("tu_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTopupList1(t *testing.T) {
	params := &stripe.TopupListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := topup.List(params)
	assert.NotNil(t, result)
}

func TestTopupCancel1(t *testing.T) {
	result, _ := topup.Cancel("tu_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTransferCreate1(t *testing.T) {
	params := &stripe.TransferParams{
		Amount:        stripe.Int64(400),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Destination:   stripe.String("acct_xxxxxxxxxxxxx"),
		TransferGroup: stripe.String("ORDER_95"),
	}
	result, _ := transfer.New(params)
	assert.NotNil(t, result)
}

func TestTransferRetrieve1(t *testing.T) {
	result, _ := transfer.Get("tr_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestTransferUpdate1(t *testing.T) {
	params := &stripe.TransferParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := transfer.Update("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTransferList1(t *testing.T) {
	params := &stripe.TransferListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := transfer.List(params)
	assert.NotNil(t, result)
}

func TestTransferReversalCreate1(t *testing.T) {
	params := &stripe.ReversalParams{Amount: stripe.Int64(100)}
	result, _ := reversal.New(params)
	assert.NotNil(t, result)
}

// TODO Test cases skipped

func TestTransferReversalUpdate1(t *testing.T) {
	params := &stripe.ReversalParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := reversal.Update("tr_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestTransferReversalList1(t *testing.T) {
	params := &stripe.ReversalListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := reversal.List(params)
	assert.NotNil(t, result)
}

func TestRadarEarlyFraudWarningRetrieve1(t *testing.T) {
	result, _ := earlyfraudwarning.Get("early_fraud_warnings", nil)
	assert.NotNil(t, result)
}

func TestRadarEarlyFraudWarningList1(t *testing.T) {
	params := &stripe.RadarEarlyFraudWarningListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := earlyfraudwarning.List(params)
	assert.NotNil(t, result)
}

func TestReviewApprove1(t *testing.T) {
	result, _ := review.Approve("prv_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestReviewRetrieve1(t *testing.T) {
	result, _ := review.Get("prv_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestReviewList1(t *testing.T) {
	params := &stripe.ReviewListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := review.List(params)
	assert.NotNil(t, result)
}

// TODO Value list test cases skipped

// TODO Value list test cases skipped

// TODO Value list test cases skipped

// TODO Value list test cases skipped

// TODO Value list test cases skipped

// TODO Value list test cases skipped

// TODO Value list test cases skipped

// TODO Value list test cases skipped

// TODO Value list test cases skipped

func TestIssuingAuthorizationRetrieve1(t *testing.T) {
	result, _ := authorization.Get("authorizations", nil)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationUpdate1(t *testing.T) {
	params := &stripe.IssuingAuthorizationParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := authorization.Update("authorizations", params)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationApprove1(t *testing.T) {
	result, _ := authorization.Approve("authorizations", nil)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationDecline1(t *testing.T) {
	result, _ := authorization.Decline("authorizations", nil)
	assert.NotNil(t, result)
}

func TestIssuingAuthorizationList1(t *testing.T) {
	params := &stripe.IssuingAuthorizationListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := authorization.List(params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderCreate1(t *testing.T) {
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

func TestIssuingCardholderRetrieve1(t *testing.T) {
	result, _ := cardholder.Get("cardholders", nil)
	assert.NotNil(t, result)
}

func TestIssuingCardholderUpdate1(t *testing.T) {
	params := &stripe.IssuingCardholderParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := cardholder.Update("cardholders", params)
	assert.NotNil(t, result)
}

func TestIssuingCardholderList1(t *testing.T) {
	params := &stripe.IssuingCardholderListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := cardholder.List(params)
	assert.NotNil(t, result)
}

func TestIssuingCardCreate1(t *testing.T) {
	params := &stripe.IssuingCardParams{
		Cardholder: stripe.String("ich_xxxxxxxxxxxxx"),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		Type:       stripe.String("virtual"),
	}
	result, _ := card.New(params)
	assert.NotNil(t, result)
}

func TestIssuingCardRetrieve1(t *testing.T) {
	result, _ := card.Get("cards", nil)
	assert.NotNil(t, result)
}

func TestIssuingCardUpdate1(t *testing.T) {
	params := &stripe.IssuingCardParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := card.Update("cards", params)
	assert.NotNil(t, result)
}

func TestIssuingCardList1(t *testing.T) {
	params := &stripe.IssuingCardListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := card.List(params)
	assert.NotNil(t, result)
}

// TODO Issuing dispute test cases skipped

// TODO Issuing dispute test cases skipped

// TODO Issuing dispute test cases skipped

// TODO Issuing dispute test cases skipped

// TODO Issuing dispute test cases skipped

func TestIssuingTransactionRetrieve1(t *testing.T) {
	result, _ := transaction.Get("transactions", nil)
	assert.NotNil(t, result)
}

func TestIssuingTransactionUpdate1(t *testing.T) {
	params := &stripe.IssuingTransactionParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := transaction.Update("transactions", params)
	assert.NotNil(t, result)
}

func TestIssuingTransactionList1(t *testing.T) {
	params := &stripe.IssuingTransactionListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := transaction.List(params)
	assert.NotNil(t, result)
}

// TODO Terminal connection token test cases skipped

// TODO Terminal location test cases skipped

// TODO Terminal location test cases skipped

// TODO Terminal location test cases skipped

// TODO Terminal location test cases skipped

// TODO Terminal location test cases skipped

func TestTerminalReaderCreate1(t *testing.T) {
	params := &stripe.TerminalReaderParams{
		RegistrationCode: stripe.String("puppies-plug-could"),
		Label:            stripe.String("Blue Rabbit"),
		Location:         stripe.String("tml_1234"),
	}
	result, _ := reader.New(params)
	assert.NotNil(t, result)
}

func TestTerminalReaderRetrieve1(t *testing.T) {
	result, _ := reader.Get("readers", nil)
	assert.NotNil(t, result)
}

func TestTerminalReaderUpdate1(t *testing.T) {
	params := &stripe.TerminalReaderParams{Label: stripe.String("Blue Rabbit")}
	result, _ := reader.Update("readers", params)
	assert.NotNil(t, result)
}

func TestTerminalReaderDelete1(t *testing.T) {
	result, _ := reader.Del("readers", nil)
	assert.NotNil(t, result)
}

func TestTerminalReaderList1(t *testing.T) {
	params := &stripe.TerminalReaderListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := reader.List(params)
	assert.NotNil(t, result)
}

func TestOrderCreate1(t *testing.T) {
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

func TestOrderRetrieve1(t *testing.T) {
	result, _ := order.Get("or_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestOrderUpdate1(t *testing.T) {
	params := &stripe.OrderUpdateParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := order.Update("or_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestOrderPay1(t *testing.T) {
	params := &stripe.OrderPayParams{
		Source: &stripe.SourceParams{Token: stripe.String("tok_xxxx")},
	}
	result, _ := order.Pay("or_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestOrderList1(t *testing.T) {
	params := &stripe.OrderListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := order.List(params)
	assert.NotNil(t, result)
}

func TestOrderReturnRetrieve1(t *testing.T) {
	result, _ := orderreturn.Get("orret_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestOrderReturnList1(t *testing.T) {
	params := &stripe.OrderReturnListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := orderreturn.List(params)
	assert.NotNil(t, result)
}

func TestSKUCreate1(t *testing.T) {
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

func TestSKURetrieve1(t *testing.T) {
	result, _ := sku.Get("sku_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSKUUpdate1(t *testing.T) {
	params := &stripe.SKUParams{}
	params.AddMetadata("order_id", "6735")
	result, _ := sku.Update("sku_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestSKUList1(t *testing.T) {
	params := &stripe.SKUListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := sku.List(params)
	assert.NotNil(t, result)
}

func TestSKUDelete1(t *testing.T) {
	result, _ := sku.Del("sku_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestSigmaScheduledQueryRunRetrieve1(t *testing.T) {
	result, _ := scheduledqueryrun.Get("scheduled_query_runs", nil)
	assert.NotNil(t, result)
}

func TestSigmaScheduledQueryRunList1(t *testing.T) {
	params := &stripe.SigmaScheduledQueryRunListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := scheduledqueryrun.List(params)
	assert.NotNil(t, result)
}

// TODO Report run test cases skipped

// TODO Report run test cases skipped

// TODO Report run test cases skipped

// TODO Report type test cases skipped

// TODO Report type test cases skipped

func TestWebhookEndpointCreate1(t *testing.T) {
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

func TestWebhookEndpointRetrieve1(t *testing.T) {
	result, _ := webhookendpoint.Get("we_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}

func TestWebhookEndpointUpdate1(t *testing.T) {
	params := &stripe.WebhookEndpointParams{
		URL: stripe.String("https://example.com/new_endpoint"),
	}
	result, _ := webhookendpoint.Update("we_xxxxxxxxxxxxx", params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointList1(t *testing.T) {
	params := &stripe.WebhookEndpointListParams{}
	params.Filters.AddFilter("limit", "", "3")
	result := webhookendpoint.List(params)
	assert.NotNil(t, result)
}

func TestWebhookEndpointDelete1(t *testing.T) {
	result, _ := webhookendpoint.Del("we_xxxxxxxxxxxxx", nil)
	assert.NotNil(t, result)
}
