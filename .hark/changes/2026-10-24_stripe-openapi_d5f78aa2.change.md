---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2478
semver_level: minor
is_stripe_api_change: true
---

* Add support for `BillingGroup` on `BillingAlertNotification`, `BillingAlertSpendThresholdFiltersParams`, `BillingAlertSpendThresholdFilters`, `BillingAlertTriggeredGroupBy`, and `BillingBillingAlertAlertNotificationListParams`
* Add support for new values `get_fund` and `nhil` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
* Add support for `PhysicalAddress` on `AccountCompany`
* Add support for `TossPay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Mandate` on `ChargePaymentMethodDetailsPaypay`, `PaymentAttemptRecordPaymentMethodDetailsPaypay`, and `PaymentRecordPaymentMethodDetailsPaypay`
* Add support for `AutomaticPassThroughFee` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `Paypay` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for new values `bw_vat`, `ci_vat`, and `gh_tin` on enums `CheckoutSessionCollectedInformationTaxId.Type`, `CheckoutSessionCustomerDetailsTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, `QuotePreviewInvoiceCustomerTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for `AdjustableQuantity` on `CheckoutSessionItemSubscriptionItem`
* Add support for new value `toss_pay` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `MarketingAttributions` on `CustomerParams` and `Customer`
* Add support for `DeliveredAt` and `ShippedAt` on `IssuingCardShipping`
* Add support for `MerchantOfRecord` on `PaymentIntentConfirmPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentDetailsMoneyServicesAccountFunding`
* Add support for new value `toss_pay` on enums `PaymentIntent.AllowedPaymentMethodTypes` and `SetupIntent.AllowedPaymentMethodTypes`
* Add support for new value `toss_pay` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `PaymentMethodData` and `PaymentMethodOptions` on `RefundParams`
* Add support for `DebitBehavior` on `SubscriptionPaymentSettingsPaymentMethodOptionsSepaDebitParams` and `SubscriptionPaymentSettingsPaymentMethodOptionsSepaDebit`
* Add support for `TriggeringSignal` on `V2SignalsAccountActivityParams` and `V2SignalsAccountActivity`
* Add support for `MfaChallenges` on `V2SignalsAccountActivityLoginDecisionParams` and `V2SignalsAccountActivityLoginDecision`
