---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/2400
is_stripe_api_change: true
released_in_version: 86.2.0
---

* Add support for new resource `FinancialConnectionsAuthorization`
* Add support for `Unreject` method on resource `Account`
* Add support for `List` method on resource `PaymentRecord`
* Add support for new values `mass_transit_parking_tax` and `parking_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, `TaxRate.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
* Add support for new value `chaps` on enums `FundingInstructionsBankTransferFinancialAddress.SupportedNetworks` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddress.SupportedNetworks`
* Add support for `SmartDisputesManagement` on `AccountSessionComponentsDisputesListFeaturesParams`, `AccountSessionComponentsDisputesListFeatures`, `AccountSessionComponentsPaymentDetailsFeaturesParams`, `AccountSessionComponentsPaymentDetailsFeatures`, `AccountSessionComponentsPaymentDisputesFeaturesParams`, `AccountSessionComponentsPaymentDisputesFeatures`, `AccountSessionComponentsPaymentsFeaturesParams`, and `AccountSessionComponentsPaymentsFeatures`
* Add support for `AdministrativeAddress` and `PrincipalPlaceOfBusiness` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
* Add support for `SEPADebitPayments` on `AccountSettingsParams`
* Remove support for `ProofOfRegistration` on `AccountDocumentsParams`.  This field was limited-use and is being deprecated.
* Add support for `PayoutsAction` on `AccountRejectParams`
* Add support for new value `data_share_only` on enums `ChargePaymentMethodDetailsCardThreeDSecure.Result`, `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure.Result`, `PaymentRecordPaymentMethodDetailsCardThreeDSecure.Result`, and `SetupAttemptPaymentMethodDetailsCardThreeDSecure.Result`
* Remove support for `DynamicTaxRates` on `CheckoutSessionLineItemParams`.  This field is limited-use and is being deprecated.
* Add support for `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsPaycoParams`, `CheckoutSessionPaymentMethodOptionsPayco`, `CheckoutSessionPaymentMethodOptionsSamsungPayParams`, `CheckoutSessionPaymentMethodOptionsSamsungPay`, `PaymentIntentConfirmPaymentMethodOptionsPaycoParams`, `PaymentIntentConfirmPaymentMethodOptionsSamsungPayParams`, `PaymentIntentPaymentMethodOptionsPaycoParams`, `PaymentIntentPaymentMethodOptionsPayco`, `PaymentIntentPaymentMethodOptionsSamsungPayParams`, `PaymentIntentPaymentMethodOptionsSamsungPay`, and `PaymentLinkPaymentIntentDataParams`
* Add support for new value `ic_nif` on enums `CheckoutSessionCustomerDetailsTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, `TaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for new values `bnp_paribas`, `citibank`, and `mbsb_bank` on enums `ConfirmationTokenPaymentMethodPreviewFpx.Bank`, `PaymentAttemptRecordPaymentMethodDetailsFpx.Bank`, and `PaymentRecordPaymentMethodDetailsFpx.Bank`
* Add support for `Network` on `DisputePaymentMethodDetailsCard`
* Add support for new values `financial_connections.account.expected_deactivation_date_updated`, `financial_connections.account.supported_payment_method_types_updated`, `financial_connections.account.upcoming_deactivation`, `financial_connections.authorization.expected_deactivation_date_updated`, and `financial_connections.authorization.upcoming_deactivation` on enum `Event.Type`
* Add support for `Limits` and `ManualEntry` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
* Add support for `RequirePaymentMethodSupport` on `FinancialConnectionsSessionFiltersParams` and `FinancialConnectionsSessionFilters`
* Add support for `BankAccountToken` on `FinancialConnectionsSession`
* Add support for `Metadata` on `InvoiceCreatePreviewSubscriptionDetailsParams`
* Add support for new values `alipay` and `mb_way` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for new value `stripe_internal_error` on enum `IssuingAuthorizationRequestHistory.Reason`
* Add support for `BusinessName` on `IssuingCardShippingParams` and `IssuingCardShipping`
* Add support for new value `correos` on enum `IssuingCardShipping.Carrier`
* Add support for `AllowedPaymentMethodTypes` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, `PaymentIntent`, `SetupIntentConfirmParams`, `SetupIntentParams`, and `SetupIntent`
* Add support for `Referrer` on `PaymentIntentConfirmRadarOptionsParams` and `PaymentIntentRadarOptionsParams`
* Add support for `ConsentCollection` and `ShippingOptions` on `PaymentLinkParams`
* Add support for `CustomFields`, `Description`, and `Footer` on `QuoteInvoiceSettingsParams`, `QuoteInvoiceSettings`, `SubscriptionScheduleDefaultSettingsInvoiceSettingsParams`, `SubscriptionScheduleDefaultSettingsInvoiceSettings`, `SubscriptionSchedulePhaseInvoiceSettingsParams`, and `SubscriptionSchedulePhaseInvoiceSettings`
* Add support for `CustomerAccount` and `Customer` on `Refund`
* Add support for `PaymentMethod` on `Refund` and `Topup`
* Add support for `Trial` on `SubscriptionSchedulePhase`
* Add support for `MassTransitParkingTax` and `ParkingTax` on `TaxRegistrationCountryOptionsUsParams` and `TaxRegistrationCountryOptionsUs`
* Add support for new values `mass_transit_parking_tax` and `parking_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
* Add support for `InitiatedBy` and `PaymentMethodOptions` on `Topup`
* Add support for `AdditionalAddresses` on `V2CoreAccountIdentityBusinessDetailsParams`, `V2CoreAccountIdentityBusinessDetails`, and `V2CoreAccountTokenIdentityBusinessDetailsParams`
* Add support for snapshot events `EventTypeFinancialConnectionsAccountExpectedDeactivationDateUpdated`, `EventTypeFinancialConnectionsAccountSupportedPaymentMethodTypesUpdated`, and `EventTypeFinancialConnectionsAccountUpcomingDeactivation` with resource `FinancialConnectionsAccount`
* Add support for snapshot events `EventTypeFinancialConnectionsAuthorizationExpectedDeactivationDateUpdated` and `EventTypeFinancialConnectionsAuthorizationUpcomingDeactivation` with resource `FinancialConnectionsAuthorization`
