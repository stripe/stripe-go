---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1758
is_stripe_api_change: true
released_in_version: 76.3.0-beta.1
---

* Add support for new resource `Margin`
* Add support for `Get`, `List`, `New`, and `Update` methods on resource `Margin`
* Add support for `Subsellers` on `CheckoutSessionPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypal`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, `PaymentIntentPaymentMethodOptionsPaypalParams`, `PaymentIntentPaymentMethodOptionsPaypal`, `SetupIntentConfirmPaymentMethodOptionsPaypalParams`, `SetupIntentPaymentMethodOptionsPaypalParams`, and `SetupIntentPaymentMethodOptionsPaypal`
* Add support for `DefaultMargins` on `InvoiceParams` and `Invoice`
* Add support for `TotalMarginAmounts` on `Invoice`
* Add support for `Margins` on `InvoiceItemParams` and `InvoiceItem`
* Add support for new values `applicant_is_not_beneficial_owner`, `current_account_tier_ineligible`, `customer_requested_account_closure`, `dispute_rate_too_high`, and `invalid_business_license` on enum `IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasons`
* Remove support for values `change_in_financial_state`, `change_in_utilization_of_credit_line`, `decrease_in_income_to_expense_ratio`, `decrease_in_social_media_performance`, `exceeds_acceptable_platform_exposure`, `has_recent_credit_limit_increase`, `insufficient_credit_utilization`, `insufficient_usage_as_qualified_expenses`, and `poor_payment_history_with_platform` from enum `IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasons`
* Add support for new values `applicant_is_not_beneficial_owner`, `current_account_tier_ineligible`, `customer_requested_account_closure`, `dispute_rate_too_high`, and `invalid_business_license` on enums `IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasons` and `IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasons`
* Add support for `IsDefault` on `IssuingPersonalizationDesignListPreferencesParams`, `IssuingPersonalizationDesignPreferencesParams`, and `IssuingPersonalizationDesignPreferences`
* Add support for `IsPlatformDefault` on `IssuingPersonalizationDesignListPreferencesParams` and `IssuingPersonalizationDesignPreferences`
* Remove support for `AccountDefault` on `IssuingPersonalizationDesignListPreferencesParams`, `IssuingPersonalizationDesignPreferencesParams`, and `IssuingPersonalizationDesignPreferences`
* Remove support for `PlatformDefault` on `IssuingPersonalizationDesignListPreferencesParams` and `IssuingPersonalizationDesignPreferences`
* Add support for `Liability` on `PaymentLinkAutomaticTaxParams` and `PaymentLinkAutomaticTax`
* Add support for `Issuer` on `PaymentLinkInvoiceCreationInvoiceDataParams` and `PaymentLinkInvoiceCreationInvoiceData`
* Add support for `InvoiceSettings` on `PaymentLinkSubscriptionDataParams` and `PaymentLinkSubscriptionData`
* Add support for new value `accept_failed_validations` on enum `QuoteStatusDetailsStaleLastReasonType`
