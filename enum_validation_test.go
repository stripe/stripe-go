package stripe

import "testing"

// TestSubscriptionProrationBehaviorIsValid tests the IsValid method for SubscriptionProrationBehavior
func TestSubscriptionProrationBehaviorIsValid(t *testing.T) {
	testCases := []struct {
		name  string
		value SubscriptionProrationBehavior
		valid bool
	}{
		{"AlwaysInvoice", SubscriptionProrationBehaviorAlwaysInvoice, true},
		{"CreateProrations", SubscriptionProrationBehaviorCreateProrations, true},
		{"None", SubscriptionProrationBehaviorNone, true},
		{"Invalid", SubscriptionProrationBehavior("invalid"), false},
		{"Empty", SubscriptionProrationBehavior(""), false},
		{"Random", SubscriptionProrationBehavior("random_value"), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.value.IsValid(); got != tc.valid {
				t.Errorf("IsValid() = %v, want %v for value %q", got, tc.valid, tc.value)
			}
		})
	}
}

// TestSubscriptionPaymentBehaviorIsValid tests the IsValid method for SubscriptionPaymentBehavior
func TestSubscriptionPaymentBehaviorIsValid(t *testing.T) {
	testCases := []struct {
		name  string
		value SubscriptionPaymentBehavior
		valid bool
	}{
		{"AllowIncomplete", SubscriptionPaymentBehaviorAllowIncomplete, true},
		{"DefaultIncomplete", SubscriptionPaymentBehaviorDefaultIncomplete, true},
		{"ErrorIfIncomplete", SubscriptionPaymentBehaviorErrorIfIncomplete, true},
		{"PendingIfIncomplete", SubscriptionPaymentBehaviorPendingIfIncomplete, true},
		{"Invalid", SubscriptionPaymentBehavior("invalid"), false},
		{"Empty", SubscriptionPaymentBehavior(""), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.value.IsValid(); got != tc.valid {
				t.Errorf("IsValid() = %v, want %v for value %q", got, tc.valid, tc.value)
			}
		})
	}
}

// TestSubscriptionCollectionMethodIsValid tests the IsValid method for SubscriptionCollectionMethod
func TestSubscriptionCollectionMethodIsValid(t *testing.T) {
	testCases := []struct {
		name  string
		value SubscriptionCollectionMethod
		valid bool
	}{
		{"ChargeAutomatically", SubscriptionCollectionMethodChargeAutomatically, true},
		{"SendInvoice", SubscriptionCollectionMethodSendInvoice, true},
		{"Invalid", SubscriptionCollectionMethod("invalid"), false},
		{"Empty", SubscriptionCollectionMethod(""), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.value.IsValid(); got != tc.valid {
				t.Errorf("IsValid() = %v, want %v for value %q", got, tc.valid, tc.value)
			}
		})
	}
}

// TestInvoiceCollectionMethodIsValid tests the IsValid method for InvoiceCollectionMethod
func TestInvoiceCollectionMethodIsValid(t *testing.T) {
	testCases := []struct {
		name  string
		value InvoiceCollectionMethod
		valid bool
	}{
		{"ChargeAutomatically", InvoiceCollectionMethodChargeAutomatically, true},
		{"SendInvoice", InvoiceCollectionMethodSendInvoice, true},
		{"Invalid", InvoiceCollectionMethod("invalid"), false},
		{"Empty", InvoiceCollectionMethod(""), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.value.IsValid(); got != tc.valid {
				t.Errorf("IsValid() = %v, want %v for value %q", got, tc.valid, tc.value)
			}
		})
	}
}

// TestQuoteCollectionMethodIsValid tests the IsValid method for QuoteCollectionMethod
func TestQuoteCollectionMethodIsValid(t *testing.T) {
	testCases := []struct {
		name  string
		value QuoteCollectionMethod
		valid bool
	}{
		{"ChargeAutomatically", QuoteCollectionMethodChargeAutomatically, true},
		{"SendInvoice", QuoteCollectionMethodSendInvoice, true},
		{"Invalid", QuoteCollectionMethod("invalid"), false},
		{"Empty", QuoteCollectionMethod(""), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.value.IsValid(); got != tc.valid {
				t.Errorf("IsValid() = %v, want %v for value %q", got, tc.valid, tc.value)
			}
		})
	}
}

// TestSubscriptionSchedulePhaseProrationBehaviorIsValid tests the IsValid method for SubscriptionSchedulePhaseProrationBehavior
func TestSubscriptionSchedulePhaseProrationBehaviorIsValid(t *testing.T) {
	testCases := []struct {
		name  string
		value SubscriptionSchedulePhaseProrationBehavior
		valid bool
	}{
		{"AlwaysInvoice", SubscriptionSchedulePhaseProrationBehaviorAlwaysInvoice, true},
		{"CreateProrations", SubscriptionSchedulePhaseProrationBehaviorCreateProrations, true},
		{"None", SubscriptionSchedulePhaseProrationBehaviorNone, true},
		{"Invalid", SubscriptionSchedulePhaseProrationBehavior("invalid"), false},
		{"Empty", SubscriptionSchedulePhaseProrationBehavior(""), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.value.IsValid(); got != tc.valid {
				t.Errorf("IsValid() = %v, want %v for value %q", got, tc.valid, tc.value)
			}
		})
	}
}

// TestBillingPortalConfigurationProrationBehaviorIsValid tests validation for billing portal proration behavior types
func TestBillingPortalConfigurationProrationBehaviorIsValid(t *testing.T) {
	t.Run("Cancel", func(t *testing.T) {
		testCases := []struct {
			name  string
			value BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior
			valid bool
		}{
			{"AlwaysInvoice", BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorAlwaysInvoice, true},
			{"CreateProrations", BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorCreateProrations, true},
			{"None", BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehaviorNone, true},
			{"Invalid", BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior("invalid"), false},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if got := tc.value.IsValid(); got != tc.valid {
					t.Errorf("IsValid() = %v, want %v for value %q", got, tc.valid, tc.value)
				}
			})
		}
	})

	t.Run("Update", func(t *testing.T) {
		testCases := []struct {
			name  string
			value BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior
			valid bool
		}{
			{"AlwaysInvoice", BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorAlwaysInvoice, true},
			{"CreateProrations", BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorCreateProrations, true},
			{"None", BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehaviorNone, true},
			{"Invalid", BillingPortalConfigurationFeaturesSubscriptionUpdateProrationBehavior("invalid"), false},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				if got := tc.value.IsValid(); got != tc.valid {
					t.Errorf("IsValid() = %v, want %v for value %q", got, tc.valid, tc.value)
				}
			})
		}
	})
}

// TestStringFunctionReturnType tests that the String() function now returns typed pointers
func TestStringFunctionReturnType(t *testing.T) {
	t.Run("SubscriptionProrationBehavior", func(t *testing.T) {
		value := String(SubscriptionProrationBehaviorNone)
		if value == nil {
			t.Fatal("String() returned nil")
		}
		if *value != SubscriptionProrationBehaviorNone {
			t.Errorf("Expected %v, got %v", SubscriptionProrationBehaviorNone, *value)
		}
		// Verify it's the correct type (this will fail to compile if type is wrong)
		var _ *SubscriptionProrationBehavior = value
	})

	t.Run("SubscriptionPaymentBehavior", func(t *testing.T) {
		value := String(SubscriptionPaymentBehaviorAllowIncomplete)
		if value == nil {
			t.Fatal("String() returned nil")
		}
		if *value != SubscriptionPaymentBehaviorAllowIncomplete {
			t.Errorf("Expected %v, got %v", SubscriptionPaymentBehaviorAllowIncomplete, *value)
		}
		var _ *SubscriptionPaymentBehavior = value
	})

	t.Run("InvoiceCollectionMethod", func(t *testing.T) {
		value := String(InvoiceCollectionMethodChargeAutomatically)
		if value == nil {
			t.Fatal("String() returned nil")
		}
		if *value != InvoiceCollectionMethodChargeAutomatically {
			t.Errorf("Expected %v, got %v", InvoiceCollectionMethodChargeAutomatically, *value)
		}
		var _ *InvoiceCollectionMethod = value
	})
}
