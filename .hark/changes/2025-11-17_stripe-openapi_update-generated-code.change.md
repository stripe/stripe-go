---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2199
is_stripe_api_change: true
released_in_version: 84.1.0-beta.1
---

* Add support for new resources `V2CoreAccountPersonToken` and `V2CoreAccountToken`
* Remove support for resource `V2PaymentsOffSessionPayment`
* Add support for `Get` and `New` methods on resources `V2CoreAccountPersonToken` and `V2CoreAccountToken`
* Remove support for `Cancel`, `Capture`, `Get`, `List`, and `New` methods on resource `V2PaymentsOffSessionPayment`
* Add support for `SpecifiedCommercialTransactionsActURL` on `AccountBusinessProfileParams` and `AccountBusinessProfile`
* Add support for `PaypayPayments` on `AccountSettingsParams` and `AccountSettings`
* Change type of `BillingAnalyticsMeterUsageMeterParams.DimensionFilters` from `string` to `array(string)`
* Change type of `BillingAnalyticsMeterUsageMeterParams.TenantFilters` from `string` to `array(string)`
* Add support for `CarRentalData`, `FlightData`, and `LodgingData` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, and `PaymentIntentPaymentDetailsParams`
* Add support for `SupplementaryPurchaseData` on `OrderPaymentSettingsPaymentMethodOptionsKlarnaParams`, `PaymentIntentConfirmPaymentMethodOptionsKlarnaParams`, and `PaymentIntentPaymentMethodOptionsKlarnaParams`
* Add support for `AllowRedisplay` and `CustomerAccount` on `PaymentMethodListParams`
* Add support for `FutureRequirements` on `V2CoreAccount`
* Add support for `KonbiniPayments` and `ScriptStatementDescriptor` on `V2CoreAccountConfigurationMerchantParams` and `V2CoreAccountConfigurationMerchant`
* Add support for `EUR` on `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams` and `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrencies`
* Add support for `RequirementsCollector` on `V2CoreAccountDefaultsResponsibilities`
* Add support for new value `ar_cuit` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
* Add support for new value `ar_dni` on enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
* Remove support for `Collector` on `V2CoreAccountRequirements`
* Add support for new value `holds_currencies.eur` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new values `payment_method` and `person` on enum `V2CoreAccountRequirementsEntryReference.Type`
* Remove support for value `resource` from enum `V2CoreAccountRequirementsEntryReference.Type`
* Remove support for value `future_requirements` from enum `V2CoreAccountRequirementsEntryRequestedReason.Code`
* Add support for `Changes` on `V2CoreEvent`
* Remove support for value `sepa_bank_account` from enum `V2MoneyManagementFinancialAddressCredentials.Type`
* Add support for `AccountToken` on `V2CoreAccountParams`
* Add support for `PersonToken` on `V2CoreAccountPersonParams`
* Add support for thin event `V2CoreHealthEventGenerationFailureResolvedEvent`
* Remove support for thin events `V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEvent`, `V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEvent`, `V2PaymentsOffSessionPaymentCanceledEvent`, `V2PaymentsOffSessionPaymentCreatedEvent`, `V2PaymentsOffSessionPaymentFailedEvent`, `V2PaymentsOffSessionPaymentRequiresCaptureEvent`, and `V2PaymentsOffSessionPaymentSucceededEvent` with related object `V2PaymentsOffSessionPayment`
