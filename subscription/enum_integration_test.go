package subscription

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v83"
	_ "github.com/stripe/stripe-go/v83/testing"
)

// TestTypedEnumsIntegration demonstrates the end-to-end functionality of typed enums
// for request parameters. This test validates that all new typed enum types work correctly
// with stripe-mock and that the form encoding properly handles typed enum pointers.
//
// This is an integration test that validates:
// - Typed enums compile and are accepted by the type system
// - Form encoding correctly serializes typed enum pointers to string values
// - stripe-mock accepts the serialized values (validates correct string representation)
//
// This test showcases the type safety improvements from issue #1679:
// - Compile-time validation of enum values
// - IDE autocomplete support
// - Prevention of typos in enum values
// - Consistent types between request params and response fields
func TestTypedEnumsIntegration(t *testing.T) {
	// This test demonstrates all the new typed enums working together in a real API call
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_123"),

		// Typed enum: SubscriptionCollectionMethod (replaces *string)
		// Before: CollectionMethod: stripe.String("charge_automatically")
		// After: Type-safe constant with compile-time validation
		CollectionMethod: stripe.String(stripe.SubscriptionCollectionMethodChargeAutomatically),

		// Typed enum: SubscriptionPaymentBehavior (NEW - includes missing default_incomplete)
		// Before: PaymentBehavior: stripe.String("allow_incomplete")
		// After: Compiler prevents typos like "allow_imcomplete"
		PaymentBehavior: stripe.String(stripe.SubscriptionPaymentBehaviorAllowIncomplete),

		// Typed enum: SubscriptionProrationBehavior (NEW)
		// Before: ProrationBehavior: stripe.String("create_prorations")
		// After: IDE autocomplete shows all valid values
		ProrationBehavior: stripe.String(stripe.SubscriptionProrationBehaviorCreateProrations),

		// Subscription items demonstrating Currency and Interval typed enums
		Items: []*stripe.SubscriptionItemsParams{
			{
				PriceData: &stripe.SubscriptionItemPriceDataParams{
					// Typed enum: Currency
					// The generic String[T] function preserves the Currency type
					Currency: stripe.String(stripe.CurrencyUSD),
					Product:  stripe.String("prod_123"),

					// Recurring pricing with typed Interval enum
					Recurring: &stripe.SubscriptionItemPriceDataRecurringParams{
						// Typed enum: PriceRecurringInterval
						// Before: Interval: stripe.String("month")
						// After: Type-safe constant
						Interval:      stripe.String(stripe.PriceRecurringIntervalMonth),
						IntervalCount: stripe.Int64(1),
					},
					UnitAmount: stripe.Int64(1000),
				},
			},
		},
	}

	// Execute the API call against stripe-mock
	// This validates that:
	// 1. The String[T] function correctly converts typed enums to pointers
	// 2. The form encoder properly serializes typed enum pointers
	// 3. stripe-mock accepts the request (validates correct string values)
	// 4. The response can be successfully parsed
	sub, err := New(params)

	// Verify the call succeeded
	assert.Nil(t, err)
	assert.NotNil(t, sub)

	// Additional validation: verify the subscription has expected properties
	assert.NotEmpty(t, sub.ID)
	assert.Equal(t, "cus_123", sub.Customer.ID)
}

// TestTypedEnumsUpdate demonstrates typed enums work correctly in update operations
func TestTypedEnumsUpdate(t *testing.T) {
	// Test updating a subscription with typed proration behavior
	params := &stripe.SubscriptionParams{
		// Demonstrate different proration behavior value
		ProrationBehavior: stripe.String(stripe.SubscriptionProrationBehaviorNone),
	}

	sub, err := Update("sub_123", params)

	assert.Nil(t, err)
	assert.NotNil(t, sub)
}

// TestTypedEnumsAllPaymentBehaviors validates all payment behavior enum values
func TestTypedEnumsAllPaymentBehaviors(t *testing.T) {
	// This test validates that all SubscriptionPaymentBehavior values work correctly,
	// including the previously missing "default_incomplete" value from issue #1679

	testCases := []struct {
		name     string
		behavior stripe.SubscriptionPaymentBehavior
	}{
		{"AllowIncomplete", stripe.SubscriptionPaymentBehaviorAllowIncomplete},
		{"DefaultIncomplete", stripe.SubscriptionPaymentBehaviorDefaultIncomplete}, // Previously missing!
		{"ErrorIfIncomplete", stripe.SubscriptionPaymentBehaviorErrorIfIncomplete},
		{"PendingIfIncomplete", stripe.SubscriptionPaymentBehaviorPendingIfIncomplete},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := &stripe.SubscriptionParams{
				Customer:        stripe.String("cus_123"),
				PaymentBehavior: stripe.String(tc.behavior),
				Items: []*stripe.SubscriptionItemsParams{
					{Price: stripe.String("price_123")},
				},
			}

			sub, err := New(params)

			assert.Nil(t, err)
			assert.NotNil(t, sub)
		})
	}
}

// TestTypedEnumsAllProrationBehaviors validates all proration behavior enum values
func TestTypedEnumsAllProrationBehaviors(t *testing.T) {
	testCases := []struct {
		name     string
		behavior stripe.SubscriptionProrationBehavior
	}{
		{"AlwaysInvoice", stripe.SubscriptionProrationBehaviorAlwaysInvoice},
		{"CreateProrations", stripe.SubscriptionProrationBehaviorCreateProrations},
		{"None", stripe.SubscriptionProrationBehaviorNone},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := &stripe.SubscriptionParams{
				Customer:          stripe.String("cus_123"),
				ProrationBehavior: stripe.String(tc.behavior),
				Items: []*stripe.SubscriptionItemsParams{
					{Price: stripe.String("price_123")},
				},
			}

			sub, err := New(params)

			assert.Nil(t, err)
			assert.NotNil(t, sub)
		})
	}
}

// TestTypedEnumsAllCollectionMethods validates all collection method enum values
func TestTypedEnumsAllCollectionMethods(t *testing.T) {
	testCases := []struct {
		name   string
		method stripe.SubscriptionCollectionMethod
	}{
		{"ChargeAutomatically", stripe.SubscriptionCollectionMethodChargeAutomatically},
		{"SendInvoice", stripe.SubscriptionCollectionMethodSendInvoice},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := &stripe.SubscriptionParams{
				Customer:         stripe.String("cus_123"),
				CollectionMethod: stripe.String(tc.method),
				Items: []*stripe.SubscriptionItemsParams{
					{Price: stripe.String("price_123")},
				},
			}

			// Add DaysUntilDue for send_invoice collection method
			if tc.method == stripe.SubscriptionCollectionMethodSendInvoice {
				params.DaysUntilDue = stripe.Int64(30)
			}

			sub, err := New(params)

			assert.Nil(t, err)
			assert.NotNil(t, sub)
		})
	}
}
