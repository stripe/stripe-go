---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2192
is_stripe_api_change: true
released_in_version: 83.2.0-alpha.1
---

* Add support for `ReportRefund` method on resource `PaymentRecord`
* Add support for new value `verification_data_not_found` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
* Add support for `Tenants` on `BillingAnalyticsMeterUsageRow`
* Add support for `RepresentativeDeclaration` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
* Add support for `Transfer` on `ApplicationFeeFeeSource`
* Add support for new value `transfer` on enum `ApplicationFeeFeeSource.Type`
* Add support for `TransitBalancesTotal` on `Balance`
* Add support for new value `transit` on enum `BalanceTransaction.BalanceType`
* Add support for `TenantGroupByKeys` on `BillingAnalyticsMeterUsageMeterParams`
* Add support for `PaymentMethodConfiguration` on `BillingPortalConfigurationFeaturesPaymentMethodUpdateParams`
* Add support for new value `solana` on enums `ChargePaymentMethodDetailsCrypto.Network`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network`, and `PaymentRecordPaymentMethodDetailsCrypto.Network`
* Add support for `PaymentPortalURL` on `ChargePaymentMethodDetailsRechnung`, `PaymentAttemptRecordPaymentMethodDetailsRechnung`, and `PaymentRecordPaymentMethodDetailsRechnung`
* Add support for `TWINT` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for new value `custom` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `CustomerSheet`, `MobilePaymentElement`, and `TaxIDElement` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
* Add support for `Provider` on `CustomerTax`
* Remove support for `RiskDetails` on `DelegatedCheckoutRequestedSessionParams`
* Add support for `RiskDetails` on `DelegatedCheckoutRequestedSessionConfirmParams`
* Add support for new value `platform_terms_of_service` on enum `File.Purpose`
* Add support for `StartingAfter` on `PaymentAttemptRecordListParams`
* Add support for `Reference` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`
* Add support for `AllocatedFunds` on `PaymentIntent`
* Add support for `SubscriptionReference` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna`
* Add support for `NameCollection` on `PaymentLinkParams` and `PaymentLink`
* Add support for `Crypto` on `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, and `RefundDestinationDetails`
* Add support for `MbWay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Add support for `Custom` on `PaymentMethodParams` and `PaymentMethod`
* Add support for `ExcludedPaymentMethodTypes` on `SetupIntentParams` and `SetupIntent`
* Add support for `Tw` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Add support for `Gip` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
* Add support for `LastSeenAt` on `TerminalReader`
* Add support for `ApplicationFeeAmount` on `TransferParams` and `Transfer`
* Add support for `ApplicationFee` on `Transfer`
* Add support for `HighRiskActivitiesDescription`, `HighRiskActivities`, `MoneyServicesDescription`, `OperatesInProhibitedCountries`, `ParticipatesInRegulatedActivity`, `PurposeOfFundsDescription`, `PurposeOfFunds`, `RegulatedActivity`, `SourceOfFundsDescription`, and `SourceOfFunds` on `V2CoreAccountConfigurationStorerParams` and `V2CoreAccountConfigurationStorer`
* Add support for `CryptoWallets` on `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesParams`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddresses`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPayments`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersParams`, and `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfers`
* Add support for `Usdc` on `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams` and `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrencies`
* Add support for `CryptoStorer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
* Add support for `ComplianceScreeningDescription` on `V2CoreAccountIdentityBusinessDetailsParams` and `V2CoreAccountIdentityBusinessDetails`
* Add support for `ExternalAmount` on `V2MoneyManagementReceivedCredit` and `V2MoneyManagementReceivedDebit`
* Add support for error code `payment_intent_rate_limit_exceeded` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `QuotePreviewInvoiceLastFinalizationError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
