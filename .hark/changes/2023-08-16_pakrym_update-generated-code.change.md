---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1699
is_breaking: true
released_in_version: 75.0.0
---

* Add `Metadata` and `Expand` to individual `Params` classes.
* `Expand`, `AddExpand`, `Metadata` and `AddMetadata` on embedded `Params` struct were deprecated.
  Before:

  ```go
  params := &stripe.AccountParams{
            Params: stripe.Params{
	            Expand: []*string{stripe.String("business_profile")},
	            Metadata: map[string]string{
		            "order_id": "6735",
	            },
            },
  }
  ```

  After:
  ```go
  params := &stripe.AccountParams{
            Expand: []*string{stripe.String("business_profile")},
            Metadata: map[string]string{
                     "order_id": "6735",
            },
  }
  ```
  You don't have to change your calls to `AddMetadata` and `AddExpand`
  Before/After:
  ```go
  params.AddMetadata("order_id", "6735")
  params.AddExpand("business_profile")
  ```
- ⚠️ Removed deprecated `excluded_territory`, `jurisdiction_unsupported`, `vat_exempt` taxability reasons:
  - `CheckoutSessionShippingCostTaxTaxabilityReasonExcludedTerritory`
  - `CheckoutSessionShippingCostTaxTaxabilityReasonJurisdictionUnsupported`
  - `CheckoutSessionShippingCostTaxTaxabilityReasonVATExempt`
  - `CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonExcludedTerritory`
  - `CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonJurisdictionUnsupported`
  - `CheckoutSessionTotalDetailsBreakdownTaxTaxabilityReasonVATExempt`
  - `CreditNoteShippingCostTaxTaxabilityReasonExcludedTerritory`
  - `CreditNoteShippingCostTaxTaxabilityReasonJurisdictionUnsupported`
  - `CreditNoteShippingCostTaxTaxabilityReasonVATExempt`
  - `InvoiceShippingCostTaxTaxabilityReasonExcludedTerritory`
  - `InvoiceShippingCostTaxTaxabilityReasonJurisdictionUnsupported`
  - `InvoiceShippingCostTaxTaxabilityReasonVATExempt`
  - `LineItemTaxTaxabilityReasonExcludedTerritory`
  - `LineItemTaxTaxabilityReasonJurisdictionUnsupported`
  - `LineItemTaxTaxabilityReasonVATExempt`
  - `QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonExcludedTerritory`
  - `QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonJurisdictionUnsupported`
  - `QuoteComputedRecurringTotalDetailsBreakdownTaxTaxabilityReasonVATExempt`
  - `QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonExcludedTerritory`
  - `QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonJurisdictionUnsupported`
  - `QuoteComputedUpfrontTotalDetailsBreakdownTaxTaxabilityReasonVATExempt`
  - `QuoteTotalDetailsBreakdownTaxTaxabilityReasonExcludedTerritory`
  - `QuoteTotalDetailsBreakdownTaxTaxabilityReasonJurisdictionUnsupported`
  - `QuoteTotalDetailsBreakdownTaxTaxabilityReasonVATExempt`
- ⚠️ Removed deprecated error code constant `ErrorCodeCardDeclinedRateLimitExceeded`, prefer `ErrorCodeCardDeclineRateLimitExceeded`.
- ⚠️ Removed deprecated error code constant `ErrorCodeInvalidSwipeData`.
- ⚠️ Removed deprecated error code constant `ErrorCodeInvoicePamentIntentRequiresAction` prefer `ErrorCodeInvoicePaymentIntentRequiresAction`.
- ⚠️ Removed deprecated error code constant `ErrorCodeSepaUnsupportedAccount`, prefer `ErrorCodeSEPAUnsupportedAccount`.
- ⚠️ Removed deprecated error code constant `ErrorCodeSkuInactive`, prefer `ErrorCodeSKUInactive`.
- ⚠️ Removed deprecated error code constant `ErrorCodeinstantPayoutsLimitExceeded`, prefer `ErrorCodeInstantPayoutsLimitExceeded`.
