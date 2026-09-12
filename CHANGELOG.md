<!--
THIS IS A GENERATED FILE. Any changes you make to it directly will be blown away.
Instead, edit a corresponding `.change.md` file and run `hark build`.
-->

# Changelog

> This changelog only covers the **public preview** releases. Each release builds on the most recent GA release; see those notes in [the GA changelog](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md).

## 86.5.0-beta.1 - 2026-08-26
This release changes the pinned API version to `2026-08-26.preview`.

* [#2408](https://github.com/stripe/stripe-go/pull/2408) Add non-verified methods to managed handlers
* ⚠️ [#2413](https://github.com/stripe/stripe-go/pull/2413) Update generated code for beta
  * Add support for new resources `V2CoreApprovalRequest`, `V2SignalsAccountActivity`, `V2SignalsAccountEvaluation`, and `V2SignalsAccountSignal`
  * Add support for `Get` and `List` methods on resource `V2SignalsAccountSignal`
  * Add support for `Get` and `New` methods on resource `V2SignalsAccountEvaluation`
  * Add support for `Del`, `Get`, and `New` methods on resource `V2SignalsAccountActivity`
  * Add support for `Cancel`, `Get`, `List`, and `Update` methods on resource `V2CoreApprovalRequest`
  * Add support for `Disable` method on resource `V2MoneyManagementPayoutMethod`
  * Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsPaymentMethodSettingsFeaturesParams`
  * ⚠️ Remove support for `PaymentMethodTypes` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, and `SetupIntentParams`
  * ⚠️ Change type of `ProductCatalogTrialOffer.Price` and `ProductCatalogTrialOfferEndBehaviorTransition.Price` from `$Price` to `deletable($Price)`
  * Add support for `Billie` on `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`
  * Add support for new value `billie` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
  * Add support for `PayoutMethods` on `V2CoreAccountDefaultsParams` and `V2CoreAccountDefaults`
  * Add support for `Restricted` on `V2CoreVaultGbBankAccount` and `V2CoreVaultUsBankAccount`
  * Add support for `EnabledDeliverySchemes` on `V2MoneyManagementPayoutMethodBankAccount`
  * ⚠️ Remove support for `EnabledDeliveryOptions` on `V2MoneyManagementPayoutMethodBankAccount`
  * Add support for new value `disabled` on enum `V2MoneyManagementPayoutMethodUsageStatus.Payments`
  * Add support for new value `disabled` on enum `V2MoneyManagementPayoutMethodUsageStatus.Transfers`
  * Add support for event notifications `V2CoreApprovalRequestApprovedEvent`, `V2CoreApprovalRequestCanceledEvent`, `V2CoreApprovalRequestCreatedEvent`, `V2CoreApprovalRequestExpiredEvent`, `V2CoreApprovalRequestFailedEvent`, `V2CoreApprovalRequestRejectedEvent`, and `V2CoreApprovalRequestSucceededEvent` with related object `V2CoreApprovalRequest`
  * Add support for event notification `V2SignalsAccountEvaluationCompleteEvent` with related object `V2SignalsAccountEvaluation`
  * Add support for error codes `authentication_failure`, `capability_not_active`, `expired_payment_method`, `incorrect_postal_code`, `invalid_canceled_subscription_fields`, and `payment_method_restricted` on `QuotePreviewInvoiceLastFinalizationError`
  * Add support for error code `default_payout_method_cannot_be_disabled` on `CannotProceedError`

## 86.3.0-beta.1 - 2026-07-29
This release changes the pinned API version to `2026-07-29.preview`.

* ⚠️ [#2380](https://github.com/stripe/stripe-go/pull/2380) Update generated code for beta
  * Add support for `Get` and `List` methods on resource `ProductCatalogTrialOffer`
  * Add support for `TaxItems` on `ChargeCapturePaymentDetailsCarRentalDataTotalTaxParams`, `ChargeCapturePaymentDetailsFlightDataTotalTaxParams`, `ChargeCapturePaymentDetailsLodgingDataTotalTaxParams`, `ChargePaymentDetailsCarRentalDataTotalTaxParams`, `ChargePaymentDetailsFlightDataTotalTaxParams`, `ChargePaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentPaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentPaymentDetailsCarRentalDatumTotalTax`, `PaymentIntentPaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentPaymentDetailsFlightDatumTotalTax`, `PaymentIntentPaymentDetailsLodgingDataTotalTaxParams`, and `PaymentIntentPaymentDetailsLodgingDatumTotalTax`
  * ⚠️ Remove support for `Taxes` on `ChargeCapturePaymentDetailsCarRentalDataTotalTaxParams`, `ChargeCapturePaymentDetailsFlightDataTotalTaxParams`, `ChargeCapturePaymentDetailsLodgingDataTotalTaxParams`, `ChargePaymentDetailsCarRentalDataTotalTaxParams`, `ChargePaymentDetailsFlightDataTotalTaxParams`, `ChargePaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentPaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentPaymentDetailsCarRentalDatumTotalTax`, `PaymentIntentPaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentPaymentDetailsFlightDatumTotalTax`, `PaymentIntentPaymentDetailsLodgingDataTotalTaxParams`, and `PaymentIntentPaymentDetailsLodgingDatumTotalTax`
  * Add support for `TaxID` on `CheckoutSessionCollectedInformation`
  * ⚠️ Remove support for `TaxIDs` on `CheckoutSessionCollectedInformation`
  * Add support for `Mode` on `FinancialConnectionsSessionManualEntry`
  * Add support for `Name` on `IssuingCardholderParams`
  * Add support for new value `ic_nif` on enums `OrderTaxDetailsTaxId.Type` and `QuotePreviewInvoiceCustomerTaxIds.Type`
  * Add support for new values `alipay` and `mb_way` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
  * Add support for `CustomFields`, `Description`, and `Footer` on `QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettings` and `QuotePreviewSubscriptionSchedulePhaseInvoiceSettings`
  * Add support for `Trial` on `QuotePreviewSubscriptionSchedulePhase`
  * ⚠️ Remove support for `ACSSDebit`, `AUBECSDebit`, `AfterpayClearpay`, `Alipay`, `Alma`, `AmazonPay`, `BACSDebit`, `BLIK`, `Bancontact`, `Billie`, `Bizum`, `Boleto`, `CardPresent`, `CashApp`, `Crypto`, `CustomerBalance`, `EPS`, `FPX`, `Giropay`, `Gopay`, `Grabpay`, `IDBankTransfer`, `IDEAL`, `InteracPresent`, `KakaoPay`, `Konbini`, `KrCard`, `MbWay`, `Mobilepay`, `Multibanco`, `NaverPay`, `NzBankAccount`, `OXXO`, `P24`, `PayByBank`, `PayNow`, `Payco`, `Paypal`, `Paypay`, `Payto`, `Pix`, `PromptPay`, `Qris`, `Rechnung`, `RevolutPay`, `SEPADebit`, `SamsungPay`, `Satispay`, `Scalapay`, `Shopeepay`, `Sofort`, `StripeBalance`, `Sunbit`, `Swish`, `TWINT`, `USBankAccount`, `Upi`, `WeChatPay`, and `Zip` on `SharedPaymentGrantedTokenPaymentMethodDetails`
  * ⚠️ Remove support for values `acss_debit`, `afterpay_clearpay`, `alipay`, `alma`, `amazon_pay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `billie`, `bizum`, `blik`, `boleto`, `card_present`, `cashapp`, `crypto`, `custom`, `customer_balance`, `eps`, `fpx`, `giropay`, `gopay`, `grabpay`, `id_bank_transfer`, `ideal`, `interac_present`, `kakao_pay`, `konbini`, `kr_card`, `mb_way`, `mobilepay`, `multibanco`, `naver_pay`, `nz_bank_account`, `oxxo`, `p24`, `pay_by_bank`, `payco`, `paynow`, `paypal`, `paypay`, `payto`, `pix`, `promptpay`, `qris`, `rechnung`, `revolut_pay`, `samsung_pay`, `satispay`, `scalapay`, `sepa_debit`, `shopeepay`, `sofort`, `stripe_balance`, `sunbit`, `swish`, `twint`, `upi`, `us_bank_account`, `wechat_pay`, and `zip` from enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
  * Add support for `UseStripeSDK` on `SharedPaymentIssuedTokenParams` and `SharedPaymentIssuedToken`
  * Add support for `RedirectToURL` on `SharedPaymentIssuedTokenNextAction`
  * ⚠️ Change type of `SharedPaymentIssuedTokenNextAction.Type` from `literal('use_stripe_sdk')` to `enum('redirect_to_url'|'use_stripe_sdk')`
  * Add support for `Livemode` on `TaxLocation`
  * Add support for `Source` on `V2IamActivityLogDetailsUserRoles`
  * Add support for `Payout` on `V2MoneyManagementReceivedCreditBalanceTransfer`
  * ⚠️ Remove support for `PayoutV1` on `V2MoneyManagementReceivedCreditBalanceTransfer`
  * Add support for new value `payout` on enum `V2MoneyManagementReceivedCreditBalanceTransfer.Type`
  * Add support for error codes `us_bank_account_microdeposits_cannot_be_confirmed` and `us_bank_account_microdeposits_cannot_be_sent` on `ControlledByAlternateResourceError`

## 86.2.0-beta.1 - 2026-06-24
This release changes the pinned API version to `2026-06-24.preview`.

* ⚠️ [#2363](https://github.com/stripe/stripe-go/pull/2363) Update generated code for beta
  * Add support for `Redaction` on `Card`, `Charge`, `CheckoutSession`, `Customer`, `IssuingAuthorization`, `IssuingCard`, `IssuingCardholder`, `IssuingDispute`, `IssuingTransaction`, `PaymentIntent`, `PaymentMethod`, `SetupIntent`, `Source`, and `Token`
  * Add support for `DisclaimerVariant` on `CapitalFinancingOffer` and `CapitalFinancingSummaryDetails`
  * Add support for `Active` on `FinancialConnectionsAccountStatusDetails` and `FinancialConnectionsAuthorizationStatusDetails`
  * Add support for new value `institution_requirement` on enum `FinancialConnectionsAccountStatusDetailsInactive.Cause`
  * Change type of `FinancialConnectionsSessionLimitsParams.Accounts` from `longInteger` to `emptyable(longInteger)`
  * Add support for `Pause` on `InvoiceCreatePreviewSubscriptionDetailsParams`
  * Add support for new value `satispay` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
  * Add support for `ReleaseDetails` on `ReserveHold`
  * Add support for `BuyerID` on `SharedPaymentGrantedTokenPaymentMethodDetailsBizum` and `SharedPaymentGrantedTokenPaymentMethodDetailsBlik`
  * Add support for `Fingerprint` on `SharedPaymentGrantedTokenPaymentMethodDetailsPix`
  * Add support for new value `money_manager` on enums `EventsV2CoreAccountLinkReturnedEvent.Configurations`, `V2CoreAccountLinkUseCaseAccountOnboarding.Configurations`, and `V2CoreAccountLinkUseCaseAccountUpdate.Configurations`
  * ⚠️ Add support for new value `money_manager` on enum `V2CoreAccount.AppliedConfigurations`
  * ⚠️ Remove support for value `storer` from enum `V2CoreAccount.AppliedConfigurations`
  * Add support for `MoneyManager` on `V2CoreAccountConfigurationParams`, `V2CoreAccountConfiguration`, `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, `V2CoreAccountIdentityAttestationsTermsOfService`, and `V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams`
  * ⚠️ Remove support for `Storer` on `V2CoreAccountConfigurationParams`, `V2CoreAccountConfiguration`, `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, `V2CoreAccountIdentityAttestationsTermsOfService`, and `V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams`
  * Add support for new values `business_storage.inbound.eur`, `business_storage.inbound.gbp`, `business_storage.inbound.usd`, `business_storage.outbound.eur`, `business_storage.outbound.gbp`, `business_storage.outbound.usd`, `received_credits.bank_accounts`, and `received_debits.bank_accounts` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for new value `money_manager` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Configuration` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Configuration`
  * ⚠️ Remove support for `MaximumRps` on `V2CoreBatchJobParams` and `V2CoreBatchJob`
  * Add support for `BIC` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
  * ⚠️ Remove support for `SwiftCode` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
  * Add support for `Processing` on `V2MoneyManagementOutboundPaymentStatusDetails` and `V2MoneyManagementOutboundTransferStatusDetails`
  * Add support for new values `fx_rate_drift_exceeded_after_review`, `payout_method_amount_limit_exceeded`, and `review_rejected` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
  * Add support for new values `fx_rate_drift_exceeded_after_review` and `review_rejected` on enum `V2MoneyManagementOutboundTransferStatusDetailsFailed.Reason`
  * Add support for `AccountHolderName` on `V2MoneyManagementReceivedCreditBankTransferUsBankAccount`
  * Add support for new value `capability_inactive` on enum `V2MoneyManagementReceivedDebitStatusDetailsFailed.Reason`
  * Add support for `Statuses` on `V2MoneyManagementFinancialAccountListParams`
  * ⚠️ Remove support for `Status` on `V2MoneyManagementFinancialAccountListParams`
  * Add support for event notifications `V2CoreAccountIncludingConfigurationMoneyManagerCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationMoneyManagerUpdatedEvent` with related object `V2CoreAccount`
  * Add support for event notification `V2MoneyManagementOutboundPaymentUnderReviewEvent` with related object `V2MoneyManagementOutboundPayment`
  * Add support for event notification `V2MoneyManagementOutboundTransferUnderReviewEvent` with related object `V2MoneyManagementOutboundTransfer`
  * ⚠️ Remove support for event notifications `V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationStorerUpdatedEvent` with related object `V2CoreAccount`
  * Add support for error codes `anomalous_money_movement_request`, `failed_tax_calculation`, `financial_account_balance_does_not_support_currency`, `financial_account_capability_not_enabled`, and `financial_account_capability_restricted` on `QuotePreviewInvoiceLastFinalizationError`
  * Add support for error code `default_us_bank_account_cannot_be_archived` on `CannotProceedError`

## 85.3.0-beta.1 - 2026-05-27
This release changes the pinned API version to `2026-05-27.preview`.

* ⚠️ [#2350](https://github.com/stripe/stripe-go/pull/2350) Update generated code for beta
  * Add support for `Pause` method on resource `Subscription`
  * Add support for `Get` method on resource `V2IamActivityLog`
  * Add support for new value `mastercard` on enum `IssuingSettlement.Network`
  * ⚠️ Change type of `ProductCatalogTrialOfferEndBehaviorTransition.Price` from `string` to `expandable($Price)`
  * Add support for `AmountPaidOffStripe` on `QuotePreviewInvoice`
  * Add support for new value `twint` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
  * Add support for `Discountable` on `QuotePreviewSubscriptionSchedulePhaseAddInvoiceItem`
  * Add support for `Bizum` and `Scalapay` on `SharedPaymentGrantedTokenPaymentMethodDetails`
  * Add support for new values `bizum` and `scalapay` on enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
  * Add support for `PaymentBehavior` on `SubscriptionResumeParams`
  * Add support for `StatusDetails` on `Subscription`
  * Add support for new values `ao_bank_account`, `az_bank_account`, `bd_bank_account`, `bo_bank_account`, `br_bank_account`, `cl_bank_account`, `ga_bank_account`, `gh_bank_account`, `gi_bank_account`, `hn_bank_account`, `kr_bank_account`, `kz_bank_account`, `la_bank_account`, `ne_bank_account`, `ng_bank_account`, `ni_bank_account`, `py_bank_account`, `sa_bank_account`, `sm_bank_account`, and `uy_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * ⚠️ Change type of `V2MoneyManagementReceivedCreditBankTransferGbBankAccount.Network` from `literal('fps')` to `enum('chaps'|'fps')`
  * Add support for error codes `payment_method_microdeposit_processing_error` and `siret_invalid` on `QuotePreviewInvoiceLastFinalizationError`

## 85.2.0-beta.2 - 2026-04-24
* ⚠️ [#2347](https://github.com/stripe/stripe-go/pull/2347) Update generated code for beta
  * Add support for new resources `V2CommerceProductCatalogImport`, `V2DataReportingQueryRun`, `V2ExtendWorkflowRun`, `V2ExtendWorkflow`, `V2IamActivityLog`, `V2NetworkBusinessProfile`, and `V2OrchestratedCommerceAgreement`
  * Add support for `Confirm`, `Get`, `List`, `New`, and `Terminate` methods on resource `V2OrchestratedCommerceAgreement`
  * Add support for `Get` and `Me` methods on resource `V2NetworkBusinessProfile`
  * Add support for `List` method on resource `V2IamActivityLog`
  * Add support for `Get` and `List` methods on resource `V2ExtendWorkflowRun`
  * Add support for `Get`, `Invoke`, and `List` methods on resource `V2ExtendWorkflow`
  * Add support for `Get` and `New` methods on resources `V2CommerceProductCatalogImport` and `V2DataReportingQueryRun`
  * ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptions.Konbini`, `V2BillingCollectionSettingPaymentMethodOptions.Konbini`, `V2BillingCollectionSettingPaymentMethodOptionsParams.Konbini`, and `V2BillingCollectionSettingVersionPaymentMethodOptions.Konbini` from `map(string: dynamic)` to `an object`
  * ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptions.SEPADebit`, `V2BillingCollectionSettingPaymentMethodOptions.SEPADebit`, `V2BillingCollectionSettingPaymentMethodOptionsParams.SEPADebit`, and `V2BillingCollectionSettingVersionPaymentMethodOptions.SEPADebit` from `map(string: dynamic)` to `an object`
  * Add support for new values `cn_bank_account` and `jp_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for new values `futsu` and `toza` on enums `V2CoreVaultGbBankAccount.BankAccountType` and `V2MoneyManagementPayoutMethodBankAccount.BankAccountType`
  * ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitProcessing` from `map(string: dynamic)` to `an object`
  * ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitQueued` from `map(string: dynamic)` to `an object`
  * ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitSucceeded` from `map(string: dynamic)` to `an object`
  * Add support for new value `payout_method_amount_limit_exceeded` on enum `V2MoneyManagementOutboundTransferStatusDetailsFailed.Reason`
  * ⚠️ Add support for new values `inbound_transfer_reversal`, `outbound_payment_reversal`, `outbound_transfer_reversal`, `received_credit_reversal`, `received_debit_reversal`, and `stripe_fee_tax` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * ⚠️ Remove support for value `return` from enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * Change type of `V2CoreBatchJobEndpointParams.HTTPMethod` from `literal('post')` to `enum('delete'|'post')`
  * Add support for new value `meter_event_value_too_many_digits` on enums `EventsV1BillingMeterErrorReportTriggeredEventReasonErrorType.Code` and `EventsV1BillingMeterNoMeterFoundEventReasonErrorType.Code`
  * Add support for `TreasuryTransaction` on `EventsV2MoneyManagementTransactionCreatedEvent`
  * Add support for event notifications `V2CommerceProductCatalogImportsFailedEvent`, `V2CommerceProductCatalogImportsProcessingEvent`, `V2CommerceProductCatalogImportsSucceededEvent`, and `V2CommerceProductCatalogImportsSucceededWithErrorsEvent` with related object `V2CommerceProductCatalogImport`
  * Add support for event notifications `V2DataReportingQueryRunCreatedEvent`, `V2DataReportingQueryRunFailedEvent`, `V2DataReportingQueryRunSucceededEvent`, and `V2DataReportingQueryRunUpdatedEvent` with related object `V2DataReportingQueryRun`
  * Add support for event notifications `V2ExtendWorkflowRunFailedEvent`, `V2ExtendWorkflowRunStartedEvent`, and `V2ExtendWorkflowRunSucceededEvent` with related object `V2ExtendWorkflowRun`
  * Add support for event notifications `V2OrchestratedCommerceAgreementConfirmedEvent`, `V2OrchestratedCommerceAgreementCreatedEvent`, `V2OrchestratedCommerceAgreementPartiallyConfirmedEvent`, and `V2OrchestratedCommerceAgreementTerminatedEvent` with related object `V2OrchestratedCommerceAgreement`
  * Add support for error type `CannotProceedError`

## 85.2.0-beta.1 - 2026-04-23
This release changes the pinned API version to `2026-04-22.preview`.

* ⚠️ [#2336](https://github.com/stripe/stripe-go/pull/2336) Update generated code for beta
  * Add support for new resources `SharedPaymentGrantedToken` and `SharedPaymentIssuedToken`
  * Add support for `Get` method on resource `SharedPaymentGrantedToken`
  * Add support for `New` and `Revoke` test helper methods on resource `SharedPaymentGrantedToken`
  * Add support for `Get`, `New`, and `Revoke` methods on resource `SharedPaymentIssuedToken`
  * Add support for `BLIK` on `CheckoutSessionPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
  * Add support for new values `fo_vat`, `gi_tin`, `it_cf`, and `py_ruc` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
  * Add support for `SharedPaymentGrantedToken` on `ConfirmationTokenPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `ValidationErrors` on `PrivacyRedactionJob`
  * Add support for `TaxDetails` on `Product`
  * Add support for new value `blik` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
  * ⚠️ Change type of `QuotePreviewInvoiceTotalTaxesTaxRateDetails.TaxRate` from `string` to `expandable($TaxRate)`
  * Add support for `AdmissionsTax`, `AttendanceTax`, `EntertainmentTax`, `GrossReceiptsTax`, `HospitalityTax`, `LuxuryTax`, `ResortTax`, and `TourismTax` on `TaxRegistrationCountryOptionsUsParams`
  * Add support for `Purpose` on `TreasuryOutboundPaymentParams` and `TreasuryOutboundPayment`
  * Add support for error codes `action_blocked` and `approval_required` on `QuotePreviewInvoiceLastFinalizationError`

## 85.1.0-beta.1 - 2026-03-25
This release changes the pinned API version to `2026-03-25.preview`.

It is built on top of SDK version 85.0.0 which contains breaking changes. Please review the [changelog for 85.0.0](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md#8500---2026-03-25) if upgrading from older SDK versions.

* ⚠️ [#2314](https://github.com/stripe/stripe-go/pull/2314) Update generated code for beta
  * Add support for new resources `ProductCatalogTrialOffer`, `TaxLocation`, and `V2CoreBatchJob`
  * Add support for `New` method on resource `ProductCatalogTrialOffer`
  * Add support for `Get`, `List`, and `New` methods on resource `TaxLocation`
  * Add support for `Cancel`, `Get`, and `New` methods on resource `V2CoreBatchJob`
  * Add support for `PerformanceLocation` on `TaxCalculationLineItemParams` and `TaxCalculationLineItem`
  * Add support for new value `performance` on enums `TaxCalculationLineItemTaxBreakdown.Sourcing`, `TaxCalculationShippingCostTaxBreakdown.Sourcing`, and `TaxTransactionShippingCostTaxBreakdown.Sourcing`
  * Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `resort_tax`, and `tourism_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
  * Add support for `TrialOffer` on `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionAddParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionSetParams`, `InvoiceCreatePreviewScheduleDetailsPhaseItemParams`, `QuoteLineActionAddItemParams`, `QuoteLineActionAddItem`, `QuoteLineActionSetItemParams`, `QuoteLineActionSetItems`, `QuotePreviewSubscriptionSchedulePhaseItem`, `SubscriptionScheduleAmendAmendmentItemActionAddParams`, `SubscriptionScheduleAmendAmendmentItemActionSetParams`, `SubscriptionSchedulePhaseItemParams`, and `SubscriptionSchedulePhaseItem`
  * Add support for `RiskReserved` on `Balance`
  * ⚠️ Remove support for `SourceType` on `ChargePaymentMethodDetailsStripeBalance`, `ConfirmationTokenPaymentMethodDataStripeBalanceParams`, `ConfirmationTokenPaymentMethodPreviewStripeBalance`, `PaymentAttemptRecordPaymentMethodDetailsStripeBalance`, `PaymentIntentConfirmPaymentMethodDataStripeBalanceParams`, `PaymentIntentPaymentMethodDataStripeBalanceParams`, `PaymentMethodStripeBalanceParams`, `PaymentMethodStripeBalance`, `PaymentRecordPaymentMethodDetailsStripeBalance`, `SetupIntentConfirmPaymentMethodDataStripeBalanceParams`, and `SetupIntentPaymentMethodDataStripeBalanceParams`
  * Add support for `TaxDetails` on `CheckoutSessionLineItemPriceDataProductDataParams`, `InvoiceAddLinesLinePriceDataProductDataParams`, `InvoiceLineItemPriceDataProductDataParams`, `InvoiceUpdateLinesLinePriceDataProductDataParams`, `PaymentLinkLineItemPriceDataProductDataParams`, `PlanProductParams`, `PriceProductDataParams`, and `ProductParams`
  * Add support for `PendingInvoiceItemInterval` on `CheckoutSessionSubscriptionDataParams`
  * Add support for `Hosted` and `UIMode` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
  * Add support for `URL` on `FinancialConnectionsSession`
  * Add support for `ExpiresAfterSeconds` on `InvoicePaymentSettingsPaymentMethodOptionsPixParams`, `InvoicePaymentSettingsPaymentMethodOptionsPix`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsPix`, `SubscriptionPaymentSettingsPaymentMethodOptionsPixParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsPix`
  * Add support for `CurrentTrial` on `InvoiceCreatePreviewSubscriptionDetailsItemParams`, `SubscriptionItemParams`, and `SubscriptionItem`
  * Add support for `Surcharge` on `PaymentIntentAmountDetailsParams`, `PaymentIntentAmountDetails`, `PaymentIntentCaptureAmountDetailsParams`, `PaymentIntentConfirmAmountDetailsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsParams`
  * Add support for `AmountDetails` and `PaymentDetails` on `PaymentIntentDecrementAuthorizationParams`
  * Add support for `MandateOptions` on `PaymentIntentPaymentMethodOptionsStripeBalance`
  * Add support for `ManagedPayments` on `PaymentLinkParams` and `PaymentLink`
  * Add support for `StripeBalance` on `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
  * Add support for `BillingCycleAnchor` on `SubscriptionTrialSettingsEndBehaviorParams` and `SubscriptionTrialSettingsEndBehavior`
  * Add support for `AdmissionsTax`, `AttendanceTax`, `EntertainmentTax`, `GrossReceiptsTax`, `HospitalityTax`, `LuxuryTax`, `ResortTax`, and `TourismTax` on `TaxRegistrationCountryOptionsUs`
  * Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `resort_tax`, and `tourism_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
  * Add support for `Requirements` on `TaxCode`
  * ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptionsCardMandateOptions.Amount`, `V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptions.Amount`, `V2BillingCollectionSettingPaymentMethodOptionsCardMandateOptionsParams.Amount`, and `V2BillingCollectionSettingVersionPaymentMethodOptionsCardMandateOptions.Amount` from `longInteger` to `int64_string`
  * Add support for new values `ar_bank_account`, `co_bank_account`, and `eg_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for `Timezone` on `V2CoreAccountDefaultsParams` and `V2CoreAccountDefaults`
  * Add support for `AzureEventGrid` on `V2CoreEventDestinationParams` and `V2CoreEventDestination`
  * Add support for new value `no_azure_partner_topic_exists` on enum `V2CoreEventDestinationStatusDetailsDisabled.Reason`
  * Add support for new value `azure_event_grid` on enum `V2CoreEventDestination.Type`
  * Add support for `SupportedCurrencies` on `V2CoreVaultGbBankAccount`, `V2CoreVaultUsBankAccount`, and `V2MoneyManagementPayoutMethodCard`
  * Add support for `Restricted` on `V2MoneyManagementPayoutMethod`
  * Add support for `Currencies` on `V2MoneyManagementPayoutMethodsBankAccountSpecCountriesField`
  * Add support for `Counterparty` and `Description` on `V2MoneyManagementTransaction`
  * ⚠️ Add support for `Currency` on `V2CoreVaultGbBankAccountParams`, `V2CoreVaultUsBankAccountParams`, `V2MoneyManagementOutboundSetupIntentPayoutMethodDataBankAccountParams`, and `V2MoneyManagementOutboundSetupIntentPayoutMethodDataCardParams`
  * Add support for `IBAN` on `V2CoreVaultGbBankAccountParams`
  * Add support for new value `currency` on enum `InvalidPaymentMethodError.InvalidParam`
  * Add support for event notifications `V2CoreBatchJobBatchFailedEvent`, `V2CoreBatchJobCanceledEvent`, `V2CoreBatchJobCompletedEvent`, `V2CoreBatchJobCreatedEvent`, `V2CoreBatchJobReadyForUploadEvent`, `V2CoreBatchJobTimeoutEvent`, `V2CoreBatchJobUpdatedEvent`, `V2CoreBatchJobUploadTimeoutEvent`, `V2CoreBatchJobValidatingEvent`, and `V2CoreBatchJobValidationFailedEvent` with related object `V2CoreBatchJob`
  * Add support for error code `service_period_coupon_with_metered_tiered_item_unsupported` on `QuotePreviewInvoiceLastFinalizationError`
* [#2325](https://github.com/stripe/stripe-go/pull/2325) Update generated code for beta
  * Release specs are identical.
* [#2331](https://github.com/stripe/stripe-go/pull/2331) Update generated code for beta

## 84.5.0-beta.1 - 2026-02-25
This release changes the pinned API version to `2026-02-25.preview`.

* [#2276](https://github.com/stripe/stripe-go/pull/2276) Update generated code for beta
  * Add support for `SmartDisputes` on `AccountSettingsParams`, `AccountSettings`, `V2CoreAccountConfigurationMerchantParams`, and `V2CoreAccountConfigurationMerchant`
  * Add support for `EmailCustomersOnSuccessfulPayment` on `AccountSettingsPaymentsParams` and `AccountSettingsPayments`
  * Add support for `ManagedPayments` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentIntent`, `SetupIntent`, and `Subscription`
  * Add support for new value `lk_vat` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
  * Add support for new value `pay_by_bank` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
  * Add support for new values `bt_bank_account`, `cr_bank_account`, `do_bank_account`, `gt_bank_account`, `md_bank_account`, `mk_bank_account`, `mo_bank_account`, `mz_bank_account`, `pe_bank_account`, `pk_bank_account`, `tw_bank_account`, and `uz_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for `Purpose` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundPayment`
  * Add support for `BranchNumber` and `SwiftCode` on `V2MoneyManagementPayoutMethodBankAccount`
  * Add support for error codes `storer_capability_missing` and `storer_capability_not_active` on `QuotePreviewInvoiceLastFinalizationError`

## 84.4.0-beta.1 - 2026-01-28
This release changes the pinned API version to `2026-01-28.preview`.

* ⚠️ [#2238](https://github.com/stripe/stripe-go/pull/2238) Fix passing context to EventNotificationHandler callbacks & update example
  - Fixes a bug where the first argument to event registration functions (e.g. `stripe.EventNotificationHandler.OnV1BillingMeterErrorReportTriggered`) didn't take a `context.Context` argument. To fix, we added the `ctx` argument:
      - before: `func (h *EventNotificationHandler) OnV1BillingMeterErrorReportTriggered(callback func(notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error`
      - after: `func (h *EventNotificationHandler) OnV1BillingMeterErrorReportTriggered(callback func(ctx context.Context, notif *V1BillingMeterErrorReportTriggeredEventNotification, client *Client) error) error`
  - this is a breaking change if you're already using the new `EventNotificationHandler`. You'll need to update the function you're registering.
* [#2249](https://github.com/stripe/stripe-go/pull/2249) Update generated code for beta
  * Add support for new resource `FinancialConnectionsAuthorization`
  * Add support for `Get` method on resource `FinancialConnectionsAuthorization`
  * Add support for `DetachPayment` method on resource `Invoice`
  * Remove support for `Cancel`, `ListLineItems`, and `Reopen` methods on resource `Order`
  * Remove support for `AttachCadence` method on resource `Subscription`
  * Add support for `AdditionalFiles` and `Site` on `AccountSettingsPaypayPaymentsParams` and `AccountSettingsPaypayPayments`
  * Remove support for `Capital` on `AccountSettings`
  * Add support for new value `pl_nip` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
  * Add support for new value `capital.financing_summary.line_of_credit_update` on enum `Event.Type`
  * Add support for `Authorization` and `StatusDetails` on `FinancialConnectionsAccount`
  * Add support for `RelinkOptions` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
  * Add support for `RelinkResult` on `FinancialConnectionsSession`
  * Remove support for `BillingCadence` on `InvoiceCreatePreviewParams`, `SubscriptionParams`, and `Subscription`
  * Remove support for `BillingCadenceDetails` on `InvoiceParent` and `QuotePreviewInvoiceParent`
  * Remove support for value `billing_cadence_details` from enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
  * Add support for `CarRentalData`, `FlightData`, and `LodgingData` on `PaymentIntentPaymentDetails`
  * Add support for new values `ae_bank_account`, `ag_bank_account`, `bh_bank_account`, `gm_bank_account`, `hk_bank_account`, `kh_bank_account`, `lc_bank_account`, `mc_bank_account`, `mg_bank_account`, `my_bank_account`, `qa_bank_account`, `rw_bank_account`, `th_bank_account`, `tt_bank_account`, and `vn_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for `AlternativeReference` on `V2CoreVaultGbBankAccount`, `V2CoreVaultUsBankAccount`, and `V2MoneyManagementPayoutMethod`
  * Add support for `AccountHolderAddress` and `AccountHolderName` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
  * Add support for `Fingerprint` on `V2MoneyManagementPayoutMethodCard`
  * Add support for snapshot event `EventTypeInvoicePaymentDetached` with resource `InvoicePayment`
  * Add support for error code `request_blocked` on `QuotePreviewInvoiceLastFinalizationError`
  * Add support for error codes `blocked_payout_method` and `unsupported_payout_method` on `BlockedByStripeError`
  * Add support for error code `invalid_payout_method_data` on `InvalidPayoutMethodError`
  * Add support for error code `limit_payout_method` on `QuotaExceededError`

## 84.2.0-beta.1 - 2025-12-16
This release changes the pinned API version to `2025-12-15.preview`.

* [#2209](https://github.com/stripe/stripe-go/pull/2209) Add EventNotificationHandler
  * This is a new, simplified way to handle event notifications (AKA thin event webhooks). Learn more in the docs: https://docs.stripe.com/webhooks/event-notification-handlers?lang=go
* [#2222](https://github.com/stripe/stripe-go/pull/2222) Update generated code for beta
  * Add support for new resources `ReserveHold`, `ReservePlan`, and `ReserveRelease`
  * Add support for `Get` and `List` methods on resources `ReserveHold` and `ReserveRelease`
  * Add support for `Get` method on resource `ReservePlan`
  * Change type of `V2FinancialAddressGeneratedMicrodeposits.Amounts` from `amount` to `an object`
  * Change type of `CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams.Amount`, `PaymentIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.Amount`, `PaymentIntentPaymentMethodOptionsPaytoMandateOptionsParams.Amount`, `SetupIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.Amount`, and `SetupIntentPaymentMethodOptionsPaytoMandateOptionsParams.Amount` from `longInteger` to `emptyable(longInteger)`
  * Change type of `CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams.AmountType`, `PaymentIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.AmountType`, `PaymentIntentPaymentMethodOptionsPaytoMandateOptionsParams.AmountType`, `SetupIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.AmountType`, and `SetupIntentPaymentMethodOptionsPaytoMandateOptionsParams.AmountType` from `enum('fixed'|'maximum')` to `emptyable(enum('fixed'|'maximum'))`
  * Change type of `CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams.EndDate`, `PaymentIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.EndDate`, `PaymentIntentPaymentMethodOptionsPaytoMandateOptionsParams.EndDate`, `SetupIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.EndDate`, and `SetupIntentPaymentMethodOptionsPaytoMandateOptionsParams.EndDate` from `string` to `emptyable(string)`
  * Change type of `CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams.PaymentSchedule`, `PaymentIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.PaymentSchedule`, `PaymentIntentPaymentMethodOptionsPaytoMandateOptionsParams.PaymentSchedule`, `SetupIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.PaymentSchedule`, and `SetupIntentPaymentMethodOptionsPaytoMandateOptionsParams.PaymentSchedule` from `enum` to `emptyable(enum)`
  * Change type of `CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams.PaymentsPerPeriod`, `PaymentIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.PaymentsPerPeriod`, `PaymentIntentPaymentMethodOptionsPaytoMandateOptionsParams.PaymentsPerPeriod`, `SetupIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.PaymentsPerPeriod`, and `SetupIntentPaymentMethodOptionsPaytoMandateOptionsParams.PaymentsPerPeriod` from `longInteger` to `emptyable(longInteger)`
  * Change type of `CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams.Purpose`, `PaymentIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.Purpose`, `PaymentIntentPaymentMethodOptionsPaytoMandateOptionsParams.Purpose`, `SetupIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.Purpose`, and `SetupIntentPaymentMethodOptionsPaytoMandateOptionsParams.Purpose` from `enum` to `emptyable(enum)`
  * Change type of `CheckoutSessionPaymentMethodOptionsPaytoMandateOptionsParams.StartDate`, `SetupIntentConfirmPaymentMethodOptionsPaytoMandateOptionsParams.StartDate`, and `SetupIntentPaymentMethodOptionsPaytoMandateOptionsParams.StartDate` from `string` to `emptyable(string)`
  * Add support for `AsyncWorkflows` on `PaymentIntent`
  * Add support for `Payto` on `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`
  * Add support for new value `payto` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
  * Remove support for `Requested` on `V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTax`, `V2CoreAccountConfigurationMerchantCapabilitiesAchDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAcssDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAffirmPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAlmaPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAuBecsDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesBacsDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesBancontactPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesBlikPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesBoletoPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesCardPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesCashappPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesEpsPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesFpxPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesGbBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesIdealPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesJcbPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesJpBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesKrCardPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesLinkPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesMxBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesOxxoPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesP24Payments`, `V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesPaycoPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesPaynowPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesPromptpayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesSepaBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesSepaDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesStripeBalancePayouts`, `V2CoreAccountConfigurationMerchantCapabilitiesSwishPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesTwintPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesUsBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesZipPayments`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocal`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWire`, `V2CoreAccountConfigurationRecipientCapabilitiesCards`, `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalancePayouts`, `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfers`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesBankAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesEur`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesGbp`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsd`, `V2CoreAccountConfigurationStorerCapabilitiesInboundTransfersBankAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsBankAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCards`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersBankAccounts`, and `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersFinancialAccounts`
  * Add support for new values `al_bank_account`, `am_bank_account`, `bn_bank_account`, `bw_bank_account`, `dz_bank_account`, `gy_bank_account`, `jm_bank_account`, `jo_bank_account`, `kw_bank_account`, `lk_bank_account`, `ma_bank_account`, `om_bank_account`, and `tz_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Change type of `V2CoreAccountIdentityBusinessDetailsAnnualRevenue.Amount`, `V2CoreAccountIdentityBusinessDetailsAnnualRevenueParams.Amount`, `V2CoreAccountIdentityBusinessDetailsMonthlyEstimatedRevenue.Amount`, `V2CoreAccountIdentityBusinessDetailsMonthlyEstimatedRevenueParams.Amount`, `V2CoreAccountTokenIdentityBusinessDetailsAnnualRevenueParams.Amount`, `V2CoreAccountTokenIdentityBusinessDetailsMonthlyEstimatedRevenueParams.Amount`, `V2FinancialAddressCreditSimulationCreditParams.Amount`, `V2MoneyManagementAdjustment.Amount`, `V2MoneyManagementInboundTransfer.Amount`, `V2MoneyManagementInboundTransferParams.Amount`, `V2MoneyManagementOutboundPayment.Amount`, `V2MoneyManagementOutboundPaymentParams.Amount`, `V2MoneyManagementOutboundPaymentQuote.Amount`, `V2MoneyManagementOutboundPaymentQuoteEstimatedFee.Amount`, `V2MoneyManagementOutboundPaymentQuoteParams.Amount`, `V2MoneyManagementOutboundTransfer.Amount`, `V2MoneyManagementOutboundTransferParams.Amount`, `V2MoneyManagementReceivedCredit.Amount`, `V2MoneyManagementReceivedDebit.Amount`, and `V2MoneyManagementTransaction.Amount` from `amount` to `an object`
  * Add support for new values `at_stn`, `at_vat`, `be_vat`, `bg_vat`, `ca_gst_hst`, `cy_he`, `cy_vat`, `cz_vat`, `de_stn`, `dk_vat`, `ee_vat`, `es_vat`, `fi_vat`, `fr_rna`, `gr_afm`, `gr_vat`, `hr_mbs`, `hr_oib`, `hr_vat`, `hu_tin`, `hu_vat`, `ie_trn`, `ie_vat`, `lt_vat`, `lu_nif`, `lu_vat`, `lv_vat`, `mt_tin`, `mt_vat`, `my_itn`, `nl_rsin`, `nl_vat`, `nz_ird`, `pl_nip`, `pl_vat`, `ro_orc`, `ro_vat`, `se_vat`, `si_tin`, `si_vat`, `sk_dic`, and `sk_vat` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
  * Remove support for value `hk_mbs` from enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
  * Add support for new values `ar_cuil`, `at_stn`, `be_nrn`, `bg_ucn`, `bn_nric`, `ca_sin`, `ch_oasi`, `cl_rut`, `cn_pp`, `co_nuip`, `cr_ci`, `cy_tic`, `cz_rc`, `dk_cpr`, `do_cie`, `ec_ci`, `ee_ik`, `es_nif`, `fi_hetu`, `fr_nir`, `gb_nino`, `gr_afm`, `hr_oib`, `hu_ad`, `id_nik`, `ie_ppsn`, `is_kt`, `it_cf`, `jp_inc`, `ke_pin`, `li_peid`, `lt_ak`, `lu_nif`, `lv_pk`, `ng_nin`, `no_nin`, `nz_ird`, `pl_pesel`, `pt_nif`, `ro_cnp`, `se_pin`, `sk_dic`, `tr_tin`, `uy_dni`, and `za_id` on enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
  * Add support for `FinancialConnectionsAccount` on `V2CoreVaultUsBankAccount` and `V2MoneyManagementPayoutMethodBankAccount`
  * Change type of `V2MoneyManagementFinancialAccountBalance.Available`, `V2MoneyManagementTransactionBalanceImpact.Available`, and `V2MoneyManagementTransactionEntryBalanceImpact.Available` from `amount` to `an object`
  * Change type of `V2MoneyManagementFinancialAccountBalance.InboundPending`, `V2MoneyManagementTransactionBalanceImpact.InboundPending`, and `V2MoneyManagementTransactionEntryBalanceImpact.InboundPending` from `amount` to `an object`
  * Change type of `V2MoneyManagementFinancialAccountBalance.OutboundPending`, `V2MoneyManagementTransactionBalanceImpact.OutboundPending`, and `V2MoneyManagementTransactionEntryBalanceImpact.OutboundPending` from `amount` to `an object`
  * Change type of `V2MoneyManagementInboundTransferFrom.Debited`, `V2MoneyManagementOutboundPaymentFrom.Debited`, `V2MoneyManagementOutboundPaymentQuoteFrom.Debited`, and `V2MoneyManagementOutboundTransferFrom.Debited` from `amount` to `an object`
  * Change type of `V2MoneyManagementInboundTransferTo.Credited`, `V2MoneyManagementOutboundPaymentQuoteTo.Credited`, `V2MoneyManagementOutboundPaymentTo.Credited`, and `V2MoneyManagementOutboundTransferTo.Credited` from `amount` to `an object`
  * Add support for `Transfer` on `V2MoneyManagementReceivedCreditBalanceTransfer`
  * Add support for new value `transfer` on enum `V2MoneyManagementReceivedCreditBalanceTransfer.Type`
  * Add support for event notification `V2MoneyManagementPayoutMethodCreatedEvent` with related object `V2MoneyManagementPayoutMethod`
  * Add support for error type `ControlledByAlternateResourceError`
  * Remove support for error type `RateLimitError`
  * Add support for error code `account_token_required_for_v2_account` on `QuotePreviewInvoiceLastFinalizationError`

## 84.1.0-beta.1 - 2025-11-18
This release changes the pinned API version to `2025-11-17.preview`.

* [#2199](https://github.com/stripe/stripe-go/pull/2199) Update generated code for beta
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

## 83.2.0-beta.1 - 2025-10-29
This release changes the pinned API version to `2025-10-29.preview`.

* [#2156](https://github.com/stripe/stripe-go/pull/2156) Update generated code for beta
  * Add support for `Update` method on resource `V2MoneyManagementFinancialAccount`
  * Add support for `ConfirmMicrodeposits`, `List`, and `SendMicrodeposits` methods on resource `V2CoreVaultUSBankAccount`
  * Add support for `List` method on resource `V2CoreVaultGBBankAccount`
  * Add support for new value `verification_data_not_found` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
  * Add support for `PaymentPortalURL` on `ChargePaymentMethodDetailsRechnung`, `PaymentAttemptRecordPaymentMethodDetailsRechnung`, and `PaymentRecordPaymentMethodDetailsRechnung`
  * Add support for `TaxIDElement` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
  * Add support for `StartingAfter` on `PaymentAttemptRecordListParams`
  * Add support for new value `solana` on enums `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network` and `PaymentRecordPaymentMethodDetailsCrypto.Network`
  * Add support for `Reference` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`
  * Add support for `SubscriptionReference` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna`
  * Add support for `Closed` on `V2CoreAccountListParams` and `V2CoreAccount`
  * Add support for new value `payment_method` on enum `V2CoreAccountConfigurationCustomerAutomaticIndirectTax.LocationSource`
  * Add support for `USD` on `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams` and `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrencies`
  * Add support for new values `application_custom` and `application_express` on enum `V2CoreAccountDefaultsResponsibilities.FeesCollector`
  * Add support for `RepresentativeDeclaration` on `V2CoreAccountIdentityAttestationsParams` and `V2CoreAccountIdentityAttestations`
  * Add support for new value `holds_currencies.usd` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for `Verification` on `V2CoreVaultUsBankAccount`
  * Add support for `V1ID` on `EventsV2MoneyManagementTransactionCreatedEvent`
  * Remove support for thin event `V2BillingBillSettingUpdatedEvent` with related object `V2BillingBillSetting`
  * Add support for error code `payment_intent_rate_limit_exceeded` on `QuotePreviewInvoiceLastFinalizationError`
  * Add support for error codes `blocked_payout_method_crypto_wallet` and `unsupported_payout_method_crypto_wallet` on `BlockedByStripeError`
  * Add support for error code `outbound_flow_from_closed_financial_account_unsupported` on `FeatureNotEnabledError`
  * Add support for error code `limit_payout_method_crypto_wallet` on `QuotaExceededError`
* [#2191](https://github.com/stripe/stripe-go/pull/2191) Update generated code for beta
  * Add support for `CryptoStorer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams`

## 83.1.0-beta.3 - 2025-10-21
* [#2183](https://github.com/stripe/stripe-go/pull/2183) Fix URL serialization for array query parameters that affected V2 GET APIs

## 83.1.0-beta.2 - 2025-10-08
* Contains bug fixes and improvements from [v83.0.1](https://github.com/stripe/stripe-go/blob/v83.0.1/CHANGELOG.md#8301---2025-10-08).

## 83.1.0-beta.1 - 2025-09-30
This release changes the pinned API version to `2025-09-30.preview`.

It is built on top of SDK version 83.0.0 which contains breaking changes. Please review the [changelog for 83.0.0](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md#8300---2025-09-30) if upgrading from older SDK versions.

* [#2123](https://github.com/stripe/stripe-go/pull/2123) Update generated code for beta
  * Add support for new resources `BillingAnalyticsMeterUsageRow`, `BillingAnalyticsMeterUsage`, `V2BillingBillSettingVersion`, `V2BillingBillSetting`, `V2BillingCadence`, `V2BillingCollectionSettingVersion`, `V2BillingCollectionSetting`, and `V2BillingProfile`
  * Remove support for resources `BillingMeterUsageRow` and `BillingMeterUsage`
  * Add support for `Get` method on resource `BillingAnalyticsMeterUsage`
  * Remove support for `Get` method on resource `BillingMeterUsage`
  * Add support for `ReportPaymentAttemptInformational` method on resource `PaymentRecord`
  * Add support for `Get`, `List`, `New`, and `Update` methods on resources `V2BillingBillSetting`, `V2BillingCollectionSetting`, and `V2BillingProfile`
  * Add support for `Get` and `List` methods on resources `V2BillingBillSettingVersion` and `V2BillingCollectionSettingVersion`
  * Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2BillingCadence`
  * Add support for new value `crypto_wallet` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for `Profile` on `V2CoreAccountDefaultsParams` and `V2CoreAccountDefaults`
  * Add support for `IP` on `V2CoreAccountIdentityAttestationsDirectorshipDeclaration`, `V2CoreAccountIdentityAttestationsOwnershipDeclaration`, `V2CoreAccountIdentityAttestationsTermsOfServiceAccountParams`, `V2CoreAccountIdentityAttestationsTermsOfServiceAccount`, `V2CoreAccountIdentityAttestationsTermsOfServiceStorerParams`, `V2CoreAccountIdentityAttestationsTermsOfServiceStorer`, `V2CoreAccountIdentityIndividualAdditionalTermsOfServiceAccount`, `V2CorePersonAdditionalTermsOfServiceAccountParams`, and `V2CorePersonAdditionalTermsOfServiceAccount`
  * Remove support for `IP` on `V2CoreAccountIdentityAttestationsDirectorshipDeclaration`, `V2CoreAccountIdentityAttestationsOwnershipDeclaration`, `V2CoreAccountIdentityAttestationsTermsOfServiceAccountParams`, `V2CoreAccountIdentityAttestationsTermsOfServiceAccount`, `V2CoreAccountIdentityAttestationsTermsOfServiceStorerParams`, `V2CoreAccountIdentityAttestationsTermsOfServiceStorer`, `V2CoreAccountIdentityIndividualAdditionalTermsOfServiceAccount`, `V2CorePersonAdditionalTermsOfServiceAccountParams`, and `V2CorePersonAdditionalTermsOfServiceAccount`
  * Remove support for `DoingBusinessAs`, `ProductDescription`, and `URL` on `V2CoreAccountIdentityBusinessDetailsParams` and `V2CoreAccountIdentityBusinessDetails`
  * Add support for `SettlementCurrency` on `V2MoneyManagementFinancialAddress`
  * Add support for `SEPABankAccount` on `V2MoneyManagementFinancialAddressCredentials` and `V2MoneyManagementReceivedCreditBankTransfer`
  * Add support for new value `sepa_bank_account` on enum `V2MoneyManagementFinancialAddressCredentials.Type`
  * Add support for `AmountDetails` and `PaymentsOrchestration` on `V2PaymentsOffSessionPaymentParams` and `V2PaymentsOffSessionPayment`
  * Add support for new value `authorization_expired` on enum `V2PaymentsOffSessionPayment.FailureReason`
  * Add support for `RetryPolicy` on `V2PaymentsOffSessionPaymentRetryDetailsParams` and `V2PaymentsOffSessionPaymentRetryDetails`
  * Add support for new values `heuristic` and `scheduled` on enum `V2PaymentsOffSessionPaymentRetryDetails.RetryStrategy`
  * Change type of `V2MoneyManagementOutboundPaymentQuoteFxQuote.LockDuration` from `literal('five_minutes')` to `enum('five_minutes'|'none')`
  * Add support for new value `none` on enum `V2MoneyManagementOutboundPaymentQuoteFxQuote.LockStatus`
  * Add support for new value `crypto_wallet` on enum `V2MoneyManagementPayoutMethod.Type`
  * Add support for `OriginType` on `V2MoneyManagementReceivedCreditBankTransfer`
  * Remove support for `PaymentMethodType` on `V2MoneyManagementReceivedCreditBankTransfer`
  * Add support for `MinimumBalanceByCurrency` on `BalanceSettingsPaymentsPayoutsParams` and `BalanceSettingsPaymentsPayouts`
  * Change type of `BalanceSettingsPaymentsSettlementTimingParams.DelayDaysOverride` from `longInteger` to `emptyable(longInteger)`
  * Remove support for values `saturday` and `sunday` from enum `BalanceSettingsPaymentsPayoutsSchedule.WeeklyPayoutDays`
  * Add support for `DelayDaysOverride` on `BalanceSettingsPaymentsSettlementTiming`
  * Add support for `AutomaticTax` and `InvoiceCreation` on `CheckoutSessionParams`
  * Add support for `UnitLabel` on `CheckoutSessionLineItemPriceDataProductDataParams`
  * Add support for `InvoiceSettings` on `CheckoutSessionSubscriptionDataParams`
  * Add support for `IntendedSubmissionMethod` on `DisputeParams` and `Dispute`
  * Change type of `DisputeSmartDisputes.RecommendedEvidence` from `string` to `array(string)`
  * Add support for new value `prevented` on enum `Dispute.Status`
  * Add support for `Pix` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
  * Add support for `BillingSchedules` on `InvoiceCreatePreviewSubscriptionDetailsParams`, `SubscriptionParams`, and `Subscription`
  * Add support for new value `pix` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for `Paypay` on `PaymentAttemptRecordPaymentMethodDetails` and `PaymentRecordPaymentMethodDetails`
  * Add support for `Wallet` on `PaymentAttemptRecordPaymentMethodDetailsCard` and `PaymentRecordPaymentMethodDetailsCard`
  * Add support for `Flexible` on `QuotePreviewSubscriptionScheduleBillingMode`
  * Add support for `BilledUntil` on `SubscriptionItem`
  * Add support for `MandateData` and `PaymentMethodOptions` on `V2PaymentsOffSessionPaymentParams`
  * Add support for `Type` on `V2MoneyManagementFinancialAddressParams`
  * Remove support for `Currency` on `V2MoneyManagementFinancialAddressParams`
  * Add support for new values `financial_addressses.crypto_wallets`, `holds_currencies.usdc`, `outbound_payments.crypto_wallets`, and `outbound_transfers.crypto_wallets` on enum `EventsV2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent.UpdatedCapability`
  * Add support for thin event `V2BillingBillSettingUpdatedEvent` with related object `V2BillingBillSetting`
  * Add support for error type `RateLimitError`
  * Add support for error codes `financial_connections_account_pending_account_numbers` and `financial_connections_account_unavailable_account_numbers` on `QuotePreviewInvoiceLastFinalizationError`
  * Add support for error code `invalid_payout_method_crypto_wallet` on `InvalidPayoutMethodError`
* [#2132](https://github.com/stripe/stripe-go/pull/2132) Update generated code for beta
  * Add support for new value `billing_cadence_details` on enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
  * Add support for `AttachCadence` method on resource `Subscription`
  * Add support for `BillingCadenceDetails` on `InvoiceParent` and `QuotePreviewInvoiceParent`
  * Add support for `BillingCadence` on `InvoiceCreatePreviewParams`, `SubscriptionParams`, and `Subscription`

## 82.6.0-beta.1 - 2025-08-27
This release changes the pinned API version to `2025-08-27.preview`.

* [#2096](https://github.com/stripe/stripe-go/pull/2096) Update generated code for beta
  * Add support for `Get` and `List` methods on resource `InvoicePayment`
  * Add support for `List` method on resource `Mandate`
  * Add support for `Applied` on `V2CoreAccountConfigurationCustomerParams`, `V2CoreAccountConfigurationCustomer`, `V2CoreAccountConfigurationMerchantParams`, `V2CoreAccountConfigurationMerchant`, `V2CoreAccountConfigurationRecipientParams`, `V2CoreAccountConfigurationRecipient`, `V2CoreAccountConfigurationStorerParams`, and `V2CoreAccountConfigurationStorer`
  * Add support for new values `ao_nif`, `az_tin`, `bd_etin`, `cr_cpj`, `cr_nite`, `do_rcn`, `gt_nit`, `kz_bin`, `mz_nuit`, `pe_ruc`, `pk_ntn`, `sa_crn`, and `sa_tin` on enum `V2CoreAccountIdentityBusinessDetailsIdNumbers.Type`
  * Add support for new values `ao_nif`, `az_tin`, `bd_brc`, `bd_etin`, `bd_nid`, `cr_cpf`, `cr_dimex`, `cr_nite`, `do_rcn`, `gt_nit`, `kz_iin`, `mz_nuit`, `pe_dni`, `pk_cnic`, `pk_snic`, and `sa_tin` on enums `V2CoreAccountIdentityIndividualIdNumbers.Type` and `V2CorePersonIdNumbers.Type`
  * Change type of `BillingAlertTriggered.Value` from `longInteger` to `decimal_string`
  * Add support for `DisplayName` on `V2MoneyManagementFinancialAccountParams` and `V2MoneyManagementFinancialAccount`
  * Add support for new value `currency_conversion` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * Add support for `CurrencyConversion` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
  * Add support for new value `currency_conversion` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
  * Add support for `Payments` on `BalanceSettingsParams` and `BalanceSettings`
  * Remove support for `DebitNegativeBalances`, `Payouts`, and `SettlementTiming` on `BalanceSettingsParams` and `BalanceSettings`
  * Add support for `Mandate` on `ChargePaymentMethodDetailsPix`, `PaymentAttemptRecordPaymentMethodDetailsPix`, and `PaymentRecordPaymentMethodDetailsPix`
  * Add support for `CouponData` on `CheckoutSessionDiscountParams`
  * Add support for `MandateOptions` on `CheckoutSessionPaymentMethodOptionsPixParams`, `CheckoutSessionPaymentMethodOptionsPix`, `PaymentIntentConfirmPaymentMethodOptionsPixParams`, `PaymentIntentPaymentMethodOptionsPixParams`, and `PaymentIntentPaymentMethodOptionsPix`
  * Change type of `CheckoutSessionPaymentMethodOptionsPix.SetupFutureUsage`, `CheckoutSessionPaymentMethodOptionsPixParams.SetupFutureUsage`, `PaymentIntentConfirmPaymentMethodOptionsPixParams.SetupFutureUsage`, `PaymentIntentPaymentMethodOptionsPix.SetupFutureUsage`, and `PaymentIntentPaymentMethodOptionsPixParams.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
  * Add support for `Amount` on `MandateMultiUse`, `PaymentAttemptRecord`, and `PaymentRecord`
  * Add support for `Currency` on `MandateMultiUse`
  * Add support for `Pix` on `MandatePaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
  * Add support for `Limit` on `PaymentAttemptRecordListParams`
  * Add support for `AmountAuthorized`, `AmountRefunded`, and `Application` on `PaymentAttemptRecord` and `PaymentRecord`
  * Add support for `ProcessorDetails` on `PaymentAttemptRecord`, `PaymentRecordReportPaymentParams`, and `PaymentRecord`
  * Remove support for `PaymentReference` on `PaymentAttemptRecord`, `PaymentRecordReportPaymentParams`, and `PaymentRecord`
  * Add support for `Installments` on `PaymentAttemptRecordPaymentMethodDetailsAlma` and `PaymentRecordPaymentMethodDetailsAlma`
  * Add support for `TransactionID` on `PaymentAttemptRecordPaymentMethodDetailsAlma`, `PaymentAttemptRecordPaymentMethodDetailsAmazonPay`, `PaymentAttemptRecordPaymentMethodDetailsBillie`, `PaymentAttemptRecordPaymentMethodDetailsKakaoPay`, `PaymentAttemptRecordPaymentMethodDetailsKrCard`, `PaymentAttemptRecordPaymentMethodDetailsNaverPay`, `PaymentAttemptRecordPaymentMethodDetailsPayco`, `PaymentAttemptRecordPaymentMethodDetailsRevolutPay`, `PaymentAttemptRecordPaymentMethodDetailsSamsungPay`, `PaymentAttemptRecordPaymentMethodDetailsSatispay`, `PaymentRecordPaymentMethodDetailsAlma`, `PaymentRecordPaymentMethodDetailsAmazonPay`, `PaymentRecordPaymentMethodDetailsBillie`, `PaymentRecordPaymentMethodDetailsKakaoPay`, `PaymentRecordPaymentMethodDetailsKrCard`, `PaymentRecordPaymentMethodDetailsNaverPay`, `PaymentRecordPaymentMethodDetailsPayco`, `PaymentRecordPaymentMethodDetailsRevolutPay`, `PaymentRecordPaymentMethodDetailsSamsungPay`, and `PaymentRecordPaymentMethodDetailsSatispay`
  * Add support for `Location` and `Reader` on `PaymentAttemptRecordPaymentMethodDetailsPaynow` and `PaymentRecordPaymentMethodDetailsPaynow`
  * Add support for `LatestActiveMandate` on `PaymentMethod`
  * Add support for `Metadata` and `Period` on `QuotePreviewSubscriptionSchedulePhaseAddInvoiceItem`
  * Add support for `PixDisplayQRCode` on `SetupIntentNextAction`
  * Add support for `ReaderSecurity` on `TerminalConfigurationParams` and `TerminalConfiguration`
  * Add support for error codes `customer_session_expired` and `india_recurring_payment_mandate_canceled` on `QuotePreviewInvoiceLastFinalizationError`

## 82.5.0-beta.2 - 2025-08-08
* [#2103](https://github.com/stripe/stripe-go/pull/2103) Bring back invoice payments APIs that were missing in the public preview SDKs
  * Add support for new resource `InvoicePayment`
  * Add support for `Get` and `List` methods on resource `InvoicePayment`

## 82.5.0-beta.1 - 2025-07-30
This release changes the pinned API version to `2025-07-30.preview`.

* [#2083](https://github.com/stripe/stripe-go/pull/2083) Update generated code for beta
  * Add support for new resources `BillingMeterUsageRow`, `BillingMeterUsage`, and `TerminalOnboardingLink`
  * Add support for `Get` method on resource `BillingMeterUsage`
  * Add support for `New` method on resource `TerminalOnboardingLink`
  * Add support for `MonthlyPayoutDays` and `WeeklyPayoutDays` on `BalanceSettingsPayoutsScheduleParams` and `BalanceSettingsPayoutsSchedule`
  * Remove support for `MonthlyAnchor` and `WeeklyAnchor` on `BalanceSettingsPayoutsScheduleParams` and `BalanceSettingsPayoutsSchedule`
  * Add support for `DelayDaysOverride` on `BalanceSettingsSettlementTimingParams`
  * Remove support for `DelayDays` on `BalanceSettingsSettlementTimingParams`
  * Add support for `UpdateDiscounts` on `CheckoutSessionPermissionsParams`
  * Add support for `Discounts` and `SubscriptionData` on `CheckoutSessionParams`
  * Add support for `SmartDisputes` on `Dispute`
  * Add support for `Upi` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
  * Add support for new value `upi` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for `TransactionID` on `PaymentAttemptRecordPaymentMethodDetailsCashapp` and `PaymentRecordPaymentMethodDetailsCashapp`
  * Add support for `AmountDetails` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentIncrementAuthorizationParams`, and `PaymentIntentParams`
  * Add support for `PaymentDetails` on `PaymentIntentIncrementAuthorizationParams`
  * Add support for `Storer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
  * Add support for `CollectionOptions` on `V2CoreAccountLinkUseCaseAccountOnboardingParams`, `V2CoreAccountLinkUseCaseAccountOnboarding`, `V2CoreAccountLinkUseCaseAccountUpdateParams`, and `V2CoreAccountLinkUseCaseAccountUpdate`
  * Change type of `V2CoreAccountLinkUseCaseAccountOnboarding.Configurations`, `V2CoreAccountLinkUseCaseAccountOnboardingParams.Configurations`, `V2CoreAccountLinkUseCaseAccountUpdate.Configurations`, and `V2CoreAccountLinkUseCaseAccountUpdateParams.Configurations` from `literal('recipient')` to `enum('customer'|'merchant'|'recipient'|'storer')`
  * Add support for `BankAccountType` on `V2MoneyManagementPayoutMethodBankAccount`
  * Add support for thin event `V2CoreAccountLinkReturnedEvent`
  * Add support for thin event `V2MoneyManagementPayoutMethodUpdatedEvent` with related object `V2MoneyManagementPayoutMethod`
  * Remove support for thin event `V2CoreAccountLinkCompletedEvent`
  * Remove support for thin event `V2OffSessionPaymentRequiresCaptureEvent` with related object `V2PaymentsOffSessionPayment`

## 82.4.0-beta.2 - 2025-07-09
* [#2084](https://github.com/stripe/stripe-go/pull/2084) Pull in V2 FinancialAccount changes for June release
  * Add support for `Close` and `New` methods on resource `V2MoneyManagementFinancialAccount`
  * Add support for new value `storer` on enum `V2CoreAccount.AppliedConfigurations`
  * Add support for `Storer` on `V2CoreAccountConfigurationParams` and `V2CoreAccountConfiguration`
  * Add support for new values `financial_addresses.bank_accounts`, `holds_currencies.gbp`, `inbound_transfers.financial_accounts`, `outbound_payments.bank_accounts`, `outbound_payments.cards`, `outbound_payments.financial_accounts`, `outbound_transfers.bank_accounts`, and `outbound_transfers.financial_accounts` on enum `V2CoreAccountRequirementsEntriesImpactRestrictsCapabilities.Capability`
  * Add support for new value `storer` on enum `V2CoreAccountRequirementsEntriesImpactRestrictsCapabilities.Configuration`
  * Add support for `StatusDetails` on `V2MoneyManagementFinancialAccount`
  * Add support for `Status` on `V2MoneyManagementFinancialAccountListParams`
  * Add support for thin events `V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationStorerUpdatedEvent` with related object `V2CoreAccount`
  * Add support for error types `AlreadyExistsError` and `NonZeroBalanceError`

## 82.4.0-beta.1 - 2025-07-01
This release changes the pinned API version to `2025-06-30.preview`.

* [#2069](https://github.com/stripe/stripe-go/pull/2069) Update generated code for beta
  * Change type of `CheckoutSessionSubscriptionDataParams.BillingMode`, `InvoiceCreatePreviewScheduleDetailsParams.BillingMode`, `InvoiceCreatePreviewSubscriptionDetailsParams.BillingMode`, `QuoteSubscriptionData.BillingMode`, `QuoteSubscriptionDataParams.BillingMode`, `SubscriptionParams.BillingMode`, and `SubscriptionScheduleParams.BillingMode` from `enum('classic'|'flexible')` to `billing_mode`
  * Add support for `SubmissionMethod` on `DisputeEvidenceDetails`
  * Add support for `OnDemand` and `Subscriptions` on `OrderPaymentSettingsPaymentMethodOptionsKlarnaParams`
  * Change type of `OrderPaymentSettingsPaymentMethodOptionsKlarna.SetupFutureUsage` and `OrderPaymentSettingsPaymentMethodOptionsKlarnaParams.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session'|'on_session')`
  * Add support for `Crypto` on `PaymentAttemptRecordPaymentMethodDetails` and `PaymentRecordPaymentMethodDetails`
  * Add support for new value `buut` on enums `PaymentAttemptRecordPaymentMethodDetailsIdeal.Bank` and `PaymentRecordPaymentMethodDetailsIdeal.Bank`
  * Add support for new value `BUUTNL2A` on enums `PaymentAttemptRecordPaymentMethodDetailsIdeal.BIC` and `PaymentRecordPaymentMethodDetailsIdeal.BIC`
  * Change type of `PaymentIntentConfirmPaymentMethodOptionsGopayParams.SetupFutureUsage`, `PaymentIntentPaymentMethodOptionsGopay.SetupFutureUsage`, and `PaymentIntentPaymentMethodOptionsGopayParams.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
  * Add support for new value `crypto` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
  * Change type of `QuotePreviewSubscriptionSchedule.BillingMode`, `Subscription.BillingMode`, and `SubscriptionSchedule.BillingMode` from `enum('classic'|'flexible')` to `SubscriptionsResourceBillingMode`
  * Change type of `SubscriptionMigrateParams.BillingMode` from `literal('flexible')` to `billing_mode_migrate`
  * Remove support for `BillingModeDetails` on `Subscription`
  * Add support for new value `xx` on enums `V2CoreAccountConfigurationCustomerAutomaticIndirectTaxLocation.Country`, `V2CoreAccountConfigurationCustomerShippingAddress.Country`, `V2CoreAccountConfigurationMerchantSupportAddress.Country`, `V2CoreAccountIdentity.Country`, `V2CoreAccountIdentityBusinessDetailsAddress.Country`, `V2CoreAccountIdentityBusinessDetailsScriptAddressesKana.Country`, `V2CoreAccountIdentityBusinessDetailsScriptAddressesKanji.Country`, `V2CoreAccountIdentityIndividualAdditionalAddresses.Country`, `V2CoreAccountIdentityIndividualAddress.Country`, `V2CoreAccountIdentityIndividualScriptAddressesKana.Country`, `V2CoreAccountIdentityIndividualScriptAddressesKanji.Country`, `V2CorePersonAdditionalAddresses.Country`, `V2CorePersonAddress.Country`, `V2CorePersonScriptAddressesKana.Country`, `V2CorePersonScriptAddressesKanji.Country`, and `V2MoneyManagementFinancialAccount.Country`
  * Add support for new value `unsupported_entity_type` on enums `V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesAchDebitPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesAcssDebitPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesAuBecsDebitPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesBacsDebitPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesBlikPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesCashappPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesEpsPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesFpxPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesGbBankTransferPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesIdealPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesJcbPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesJpBankTransferPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesMxBankTransferPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesOxxoPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesPaynowPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesPromptpayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesSepaBankTransferPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesSepaDebitPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesStripeBalancePayoutsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesTwintPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesUsBankTransferPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsStatusDetail.Code`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalStatusDetail.Code`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireStatusDetail.Code`, `V2CoreAccountConfigurationRecipientCapabilitiesCardsStatusDetail.Code`, `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalancePayoutsStatusDetail.Code`, and `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersStatusDetail.Code`
  * Add support for `ProofOfAddress` on `V2CoreAccountIdentityBusinessDetailsDocumentsParams` and `V2CoreAccountIdentityBusinessDetailsDocuments`
  * Add support for new value `xx` on enums `V2CoreAccountIdentityIndividual.Nationalities` and `V2CorePerson.Nationalities`
  * Add support for `Metadata` on `V2MoneyManagementFinancialAccount`
  * Remove support for `Description` on `V2MoneyManagementFinancialAccount`
  * Add support for new value `pending` on enum `V2MoneyManagementFinancialAccount.Status`
  * Remove support for `Attempts` on `V2PaymentsOffSessionPayment`
  * Add support for `FromAccount`, `OutboundPayment`, and `OutboundTransfer` on `V2MoneyManagementReceivedCreditBalanceTransfer`
  * Change type of `V2MoneyManagementReceivedCreditBalanceTransfer.Type` from `literal('payout_v1')` to `enum('outbound_payment'|'outbound_transfer'|'payout_v1')`
  * Add support for error codes `recipient_feature_not_active`, `storer_capability_missing`, and `storer_capability_not_active` on `FeatureNotEnabledError`
  * Remove support for error code `outbound_payment_recipient_feature_not_active` on `FeatureNotEnabledError`
  * Add support for error code `insufficient_funds` on `InsufficientFundsError`
  * Remove support for error codes `outbound_payment_insufficient_funds` and `outbound_transfer_insufficient_funds` on `InsufficientFundsError`
  * Add support for error codes `recipient_amount_limit_exceeded` and `recipient_count_limit_exceeded` on `QuotaExceededError`
  * Remove support for error codes `outbound_payment_recipient_amount_limit_exceeded` and `outbound_payment_recipient_count_limit_exceeded` on `QuotaExceededError`
  * Add support for error code `recipient_email_does_not_exist` on `RecipientNotNotifiableError`
  * Remove support for error code `outbound_payment_recipient_email_does_not_exist` on `RecipientNotNotifiableError`

## 82.3.0-beta.2 - 2025-06-26
* [#2081](https://github.com/stripe/stripe-go/pull/2081) Pull in OffSessionPayment changes for the May release

## 82.3.0-beta.1 - 2025-05-29
This release changes the pinned API version to `2025-05-28.preview`.

* [#2060](https://github.com/stripe/stripe-go/pull/2060) Update generated code for beta
  ### Breaking changes
  * Remove support for deprecated previews
    * Remove support for resources `BillingMeterErrorReport`, `GiftCardsCard`, `GiftCardsTransaction`, and `PrivacyRedactionJobRootObjects`
    * Remove support for `Get`, `List`, `New`, `Update`, and `Validate` methods on resource `GiftCardsCard`
    * Remove support for `Cancel`, `Confirm`, `Get`, `List`, `New`, and `Update` methods on resource `GiftCardsTransaction`
    * Remove support for `Provisioning` on `ProductParams` and `Product`
    * Remove support for snapshot event `EventTypeBillingMeterErrorReportTriggered` with resource `BillingMeterErrorReport`
    * Remove support for error codes `gift_card_balance_insufficient`, `gift_card_code_exists`, and `gift_card_inactive` on `Error` and `QuotePreviewInvoiceLastFinalizationError`
  * Remove support for `Credits` on `OrderParams` and `Order`
  * Remove support for `AmountRemaining` on `Order`
  * Remove support for `AmountCredit` on `OrderTotalDetails`
  * Remove support for `AsyncWorkflows` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentDecrementAuthorizationParams`, `PaymentIntentIncrementAuthorizationParams`, `PaymentIntentParams`, and `PaymentIntent`
  * Remove support for values `credits_attributed_to_debits` and `legacy_prorations` from enums `QuotePreviewSubscriptionSchedule.BillingMode`, `QuoteSubscriptionData.BillingMode`, `Subscription.BillingMode`, and `SubscriptionSchedule.BillingMode`
  * Remove support for `StatusDetails` and `Status` on `TaxAssociation`
  * Change type of `InvoiceCreatePreviewSubscriptionDetailsParams.CancelAt` and `SubscriptionParams.CancelAt` from `DateTime` to `DateTime | enum('max_period_end'|'min_period_end')`
  * Change type of `CheckoutSessionLineItemParams.Quantity` from `emptyable(longInteger)` to `longInteger`
  * Change type of `PrivacyRedactionJob.Objects` from `$Privacy.RedactionJobRootObjects` to `RedactionResourceRootObjects`
  * Change type of `PrivacyRedactionJob.Status` from `string` to `enum`
  * Change type of `PrivacyRedactionJob.ValidationBehavior` from `string` to `enum('error'|'fix')`
  * Change type of `PrivacyRedactionJobValidationError.Code` from `string` to `enum`
  * Change type of `PrivacyRedactionJobValidationError.ErroringObject` from `map(string: string)` to `RedactionResourceErroringObject`

  ### Other changes
  * Add support for `Migrate` method on resource `Subscription`
  * Add support for `Distance`, `PickupLocationName`, `ReturnLocationName`, and `VehicleIdentificationNumber` on `ChargeCapturePaymentDetailsCarRentalParams`, `ChargePaymentDetailsCarRentalParams`, `PaymentIntentCapturePaymentDetailsCarRentalParams`, `PaymentIntentConfirmPaymentDetailsCarRentalParams`, `PaymentIntentPaymentDetailsCarRentalParams`, and `PaymentIntentPaymentDetailsCarRental`
  * Add support for `DriverIdentificationNumber` and `DriverTaxNumber` on `ChargeCapturePaymentDetailsCarRentalDriverParams`, `ChargePaymentDetailsCarRentalDriverParams`, `PaymentIntentCapturePaymentDetailsCarRentalDriverParams`, `PaymentIntentConfirmPaymentDetailsCarRentalDriverParams`, `PaymentIntentPaymentDetailsCarRentalDriverParams`, and `PaymentIntentPaymentDetailsCarRentalDriver`
  * Add support for `Institution` on `FinancialConnectionsAccount`
  * Add support for `Countries` on `FinancialConnectionsInstitution`
  * Add support for `Location` and `Reader` on `PaymentAttemptRecordPaymentMethodDetailsAffirm`, `PaymentAttemptRecordPaymentMethodDetailsWechatPay`, `PaymentRecordPaymentMethodDetailsAffirm`, and `PaymentRecordPaymentMethodDetailsWechatPay`
  * Add support for `Hooks` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentDecrementAuthorizationParams`, `PaymentIntentIncrementAuthorizationParams`, `PaymentIntentParams`, and `PaymentIntent`
  * Add support for `CardPresent` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptions`
  * Change type of `PaymentRecordReportPaymentAttemptCanceledParams.Metadata`, `PaymentRecordReportPaymentAttemptFailedParams.Metadata`, `PaymentRecordReportPaymentAttemptGuaranteedParams.Metadata`, `PaymentRecordReportPaymentAttemptParams.Metadata`, and `PaymentRecordReportPaymentParams.Metadata` from `map(string: string)` to `emptyable(map(string: string))`
  * Add support for `Livemode` on `PrivacyRedactionJob`
  * Add support for new values `classic` and `flexible` on enums `QuotePreviewSubscriptionSchedule.BillingMode`, `QuoteSubscriptionData.BillingMode`, `Subscription.BillingMode`, and `SubscriptionSchedule.BillingMode`
  * Add support for `BillingThresholds` on `QuotePreviewSubscriptionScheduleDefaultSettings`, `QuotePreviewSubscriptionSchedulePhaseItem`, and `QuotePreviewSubscriptionSchedulePhase`
  * Add support for `BillingModeDetails` on `Subscription`
  * Add support for `TaxTransactionAttempts` on `TaxAssociation`
  * Add support for `ConfirmConfig` on `TerminalReaderActionConfirmPaymentIntent` and `TerminalReaderConfirmPaymentIntentParams`
  * Add support for error code `forwarding_api_upstream_error` on `QuotePreviewInvoiceLastFinalizationError`

## 82.2.0-beta.2 - 2025-04-30
* [#2059](https://github.com/stripe/stripe-go/pull/2059) Update generated code for beta
  Release specs are identical.

## 82.2.0-beta.1 - 2025-04-30
This release changes the pinned API version to `2025-04-30.preview`.

* ⚠️ [#2054](https://github.com/stripe/stripe-go/pull/2054) Unexport GetBaseEvent and V2ErrorCode
  * ⚠️ Unexported `GetBaseEvent` --> `getBaseEvent`. This function should not be called. Instead, type-cast the `V2Event` to a concrete Event struct.
  * ⚠️ Removed the `V2ErrorCode` type. Error codes should not be handled programmatically for V2 errors.
* [#2034](https://github.com/stripe/stripe-go/pull/2034) Update generated code for beta
  This release changes the pinned API version to `2025-04-30.preview`.

  * Add support for `BillingMode` on `CheckoutSessionSubscriptionDataParams`, `InvoiceCreatePreviewScheduleDetailsParams`, `InvoiceCreatePreviewSubscriptionDetailsParams`, `QuotePreviewSubscriptionSchedule`, `QuoteSubscriptionDataParams`, `QuoteSubscriptionData`, `SubscriptionParams`, `SubscriptionScheduleParams`, `SubscriptionSchedule`, and `Subscription`
  * Add support for new values `aw_tin`, `az_tin`, `bd_bin`, `bf_ifu`, `bj_ifu`, `cm_niu`, `cv_nif`, `et_tin`, `kg_tin`, and `la_tin` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
  * Add support for `AccountNumber` on `ConfirmationTokenPaymentMethodPreviewAcssDebit` and `PaymentMethodAcssDebit`
  * Add support for new value `balance_settings.updated` on enum `Event.Type`

## 82.1.0-beta.3 - 2025-04-17
* [#2023](https://github.com/stripe/stripe-go/pull/2023) Update generated code for beta
  * Add support for new resources `FxQuote` and `PaymentIntentAmountDetailsLineItem`
  * Add support for new services `fxquote.Client` (accessed by `client.API.FxQuotes`) and `paymentintentamountdetailslineitem.Client` (accessed by `client.API.PaymentIntentAmountDetailsLineItems`)
  * Remove support for service `invoicepayment.Client` (accessed by `client.API.InvoicePayments`)
  * Add support for `Get`, `List`, and `New` methods on resource `FxQuote`
  * Remove support for `AttachPaymentIntent` method on resource `Invoice`
  * Remove support for `Get` and `List` methods on resource `InvoicePayment`
  * Add support for `List` method on resource `PaymentIntentAmountDetailsLineItem`
  * Add support for `List` method on service `paymentintentamountdetailslineitem.Client`
  * Add support for `Get`, `List`, and `New` methods on service `fxquote.Client`
  * Remove support for `Get` and `List` methods on service `invoicepayment.Client`
  * Remove support for `AttachPaymentIntent` method on service `invoice.Client`
  * Add support for `RegistrationDate` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
  * Add support for `USCfpbData` on `AccountParams`, `PersonParams`, `Person`, and `TokenPersonParams`
  * Add support for `CustomerReference` and `OrderReference` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
  * Add support for `TaxID` on `ChargeBillingDetails`, `ConfirmationTokenPaymentMethodDataBillingDetailsParams`, `ConfirmationTokenPaymentMethodPreviewBillingDetails`, `PaymentIntentConfirmPaymentMethodDataBillingDetailsParams`, `PaymentIntentPaymentMethodDataBillingDetailsParams`, `PaymentMethodBillingDetailsParams`, `PaymentMethodBillingDetails`, `SetupIntentConfirmPaymentMethodDataBillingDetailsParams`, `SetupIntentPaymentMethodDataBillingDetailsParams`, `TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParams`, and `TreasuryOutboundPaymentDestinationPaymentMethodDataBillingDetailsParams`
  * Add support for `PriceData` on `CheckoutSessionLineItemParams`
  * Add support for `Script` on `CouponParams` and `Coupon`
  * Add support for `Type` on `Coupon`
  * Add support for new value `fx_quote.expired` on enum `Event.Type`
  * Add support for new value `affirm` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for `FxQuote` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, `PaymentIntent`, `TransferParams`, and `Transfer`
  * Add support for `DiscountAmount`, `LineItems`, `Shipping`, and `Tax` on `PaymentIntentAmountDetails`
  * Add support for `Pix` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
  * Add support for `PendingReason` on `Refund`
  * Add support for `Aw`, `Az`, `Bd`, `Bj`, `ET`, `Kg`, `La`, and `Ph` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
  * Add support for snapshot event `EventTypeFxQuoteExpired` with resource `FxQuote`

## 82.1.0-beta.2 - 2025-04-10
* [#2014](https://github.com/stripe/stripe-go/pull/2014) Handle top-level External Account service
  - Add support for top-level External Account CRUDL calls by adding separate `GetCard`/`GetBankAccount`, `NewCard`/`NewBankAccount`, etc. calls.
* [#2019](https://github.com/stripe/stripe-go/pull/2019) Update generated code for beta
  ### Breaking changes
  * Change type of `V2MoneyManagementReceivedDebit.StatusTransitions` from `an object` to `nullable(an object)`
  * Remove support for values `bank_accounts.local_uk`, `bank_accounts.wire_uk`, `cards_uk`, and `crypto_wallets_v2` from enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`

  ### Additions
  * Add support for new resources `Privacy.RedactionJobRootObjects`, `Privacy.RedactionJobValidationError`, and `Privacy.RedactionJob`
  * Add support for `Cancel`, `Get`, `List`, `New`, `Run`, `Update`, and `Validate` methods on resource `RedactionJob`
  * Add support for `Get` and `List` methods on resource `RedactionJobValidationError`
  * Add support for `MinorityOwnedBusinessDesignation` on `AccountBusinessProfileParams` and `AccountBusinessProfile`
  * Add support for `ExportTaxTransactions` and `PaymentDisputes` on `AccountSessionComponentsParams`
  * Add support for new value `tax_id_prohibited` on enums `InvoiceLastFinalizationError.Code`, `PaymentIntentLastPaymentError.Code`, `QuotePreviewInvoiceLastFinalizationError.Code`, `SetupAttemptSetupError.Code`, `SetupIntentLastSetupError.Code`, and `StripeError.Code`
  * Add support for new value `verification_legal_entity_structure_mismatch` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
  * Add support for new value `fixed_term_loan` on enum `CapitalFinancingOffer.Type`
  * Add support for `WalletOptions` on `CheckoutSessionParams` and `CheckoutSession`
  * Add support for new values `privacy.redaction_job.canceled`, `privacy.redaction_job.created`, `privacy.redaction_job.ready`, `privacy.redaction_job.succeeded`, and `privacy.redaction_job.validation_error` on enum `Event.Type`
  * Add support for `Klarna` on `PaymentMethodDomain`
  * Add support for `In` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`

## 82.1.0-beta.1 - 2025-04-02
This release changes the pinned API version to `2025-03-31.preview`.

#### New APIs for Accounts v2 in private preview

See [SaaS platform payments with subscription billing using Accounts v2](https://docs.stripe.com/connect/accounts-v2/saas-platform-payments-billing)

* [#2009](https://github.com/stripe/stripe-go/pull/2009) Update generated code for beta
  This release changes the pinned API version to `2025-03-31.preview`

### Breaking Changes
* ⚠️ Change type of `InvoiceCreatePreviewSubscriptionDetailsParams.CancelAt` and `SubscriptionParams.CancelAt` from `DateTime | literal('min_period_end')` to `DateTime`
* ⚠️ Change type of `PaymentAttemptRecord.payment_method_details.type` and `PaymentRecord.payment_method_details.type` from `literal('custom')` to `string`
* ⚠️ Change type of `PaymentAttemptRecord.payment_record` from `string` to `nullable(string)`
* ⚠️ Change type of `PaymentRecord.latest_payment_attempt_record` from `string` to `nullable(string)`
* ⚠️ Remove support for `AmountOverpaid` on `InvoicePayment`
* ⚠️ Remove support for `ApplicationFeeAmount`, `PaidOutOfBand`, and `Paid` on `QuotePreviewInvoice`
* ⚠️ Remove support for `BillingThresholds` on `QuotePreviewSubscriptionScheduleDefaultSettings`, `QuotePreviewSubscriptionSchedulePhasesItems`, and `QuotePreviewSubscriptionSchedulePhases`
* ⚠️ Remove support for `RateCardSubscriptionDetails` on `InvoiceItemParent`
* ⚠️ Remove support for `Value` on `TerminalReaderActionCollectInputsInputsSelectionChoices`, `TerminalReaderActionCollectInputsInputsSelection`, and `TerminalReaderCollectInputsInputsSelectionChoicesParams`
* ⚠️ Remove support for values `out_of_band_payment` and `payment_record` from enum `InvoicePaymentPayment.Type`

### Additions
* Add support for `Billie` on `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`
* Add support for `CancelAtPeriodEnd` on `InvoiceCreatePreviewSubscriptionDetailsParams`, `SubscriptionParams`, and `Subscription`
* Add support for `ConfirmationSecret` and `Parent` on `QuotePreviewInvoice`
* Add support for `Context` on `Event`
* Add support for `CustomerAccount` on `BillingCreditBalanceSummaryParams`, `BillingCreditBalanceSummary`, `BillingCreditBalanceTransactionListParams`, `BillingCreditGrantListParams`, `BillingCreditGrantParams`, `BillingCreditGrant`, `BillingPortalSessionParams`, `BillingPortalSession`, `CashBalance`, `CheckoutSessionListParams`, `CheckoutSessionParams`, `CheckoutSession`, `ConfirmationTokenPaymentMethodPreview`, `CreditNoteListParams`, `CreditNote`, `CustomerBalanceTransaction`, `CustomerCashBalanceTransaction`, `CustomerSessionParams`, `CustomerSession`, `Customer`, `Discount`, `FinancialConnectionsAccountAccountHolder`, `FinancialConnectionsAccountListAccountHolderParams`, `FinancialConnectionsSessionAccountHolderParams`, `FinancialConnectionsSessionAccountHolder`, `InvoiceCreatePreviewParams`, `InvoiceItemListParams`, `InvoiceItemParams`, `InvoiceItem`, `InvoiceListParams`, `InvoiceParams`, `Invoice`, `PaymentIntentListParams`, `PaymentIntentParams`, `PaymentIntent`, `PaymentMethodAttachParams`, `PaymentMethod`, `PromotionCodeListParams`, `PromotionCodeParams`, `PromotionCode`, `QuoteListParams`, `QuoteParams`, `QuotePreviewInvoice`, `QuotePreviewSubscriptionSchedule`, `Quote`, `SetupAttempt`, `SetupIntentListParams`, `SetupIntentParams`, `SetupIntent`, `SubscriptionListParams`, `SubscriptionParams`, `SubscriptionScheduleListParams`, `SubscriptionScheduleParams`, `SubscriptionSchedule`, `Subscription`, `TaxIdOwner`, and `TaxId`
* Add support for `Del`, `Get`, `List`, `New`, and `Update` methods on resource `TlExternalAccount`
* Add support for `DurationInMonths` on `CouponParams`
* Add support for `Get` and `Update` methods on resource `BalanceSettings`
* Add support for `ID` and `Text` on `TerminalReaderActionCollectInputsInputsSelectionChoices`, `TerminalReaderActionCollectInputsInputsSelection`, and `TerminalReaderCollectInputsInputsSelectionChoicesParams`
* Add support for `Installments` on `ConfirmationTokenPaymentMethodOptionsCard`
* Add support for new resources `BalanceSettings` and `TlExternalAccount`
* Add support for new value `repeating` on enum `Coupon.Duration`
* Add support for new values `forwarding_api_retryable_upstream_error`, `v2_account_disconnection_unsupported`, and `v2_account_missing_configuration` on enum `QuotePreviewInvoiceLastFinalizationError.Code`
* Add support for new values `nz_bank_account` and `stripe_balance` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
* Add support for `NzBankAccount` on `PaymentAttemptRecordPaymentMethodDetails` and `PaymentRecordPaymentMethodDetails`
* Add support for `PaymentMethodOptions` on `ConfirmationTokenParams`
* Add support for `PayoutMethod` on `PayoutParams` and `Payout`
* Add support for `Provider` on `CheckoutSessionAutomaticTax`, `InvoiceAutomaticTax`, `QuoteAutomaticTax`, and `QuotePreviewInvoiceAutomaticTax`
* Add support for `RelatedCustomerAccount` on `IdentityVerificationSessionListParams`, `IdentityVerificationSessionParams`, and `IdentityVerificationSession`
* Add support for `StripeBalancePayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `StripeBalance` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `TaxCalculationReference` on `CreditNoteLineItem`, `InvoiceLineItem`, and `LineItem`
* Add support for `UpdateLineItems` on `CheckoutSessionPermissionsParams` and `CheckoutSessionPermissions`

### Changes
* [#2002](https://github.com/stripe/stripe-go/pull/2002) Update behavior of AddBetaHeader
  - AddBetaVersion will use the highest version number used for a beta feature instead of return an `error` on a conflict as it had done previously.
* Change type of `InvoiceItemParent.Type` from `enum('rate_card_subscription_details'|'subscription_details')` to `literal('subscription_details')`
* Change type of `InvoicePayment.is_default` from `nullable(boolean)` to `boolean`
* Change type of `PaymentAttemptRecord.payment_method_details.custom` and `PaymentRecord.payment_method_details.custom` from `nullable(PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCustomDetails)` to `PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCustomDetails`

### New APIs for Accounts v2 in private preview
* Add support for `AcknowledgeConfirmationOfPayee`, `Archive`, `Get`, `InitiateConfirmationOfPayee`, and `New` methods on resource `V2.Core.Vault.GbBankAccount`
* Add support for `Archive`, `Get`, `New`, and `Update` methods on resource `V2.Core.Vault.UsBankAccount`
* Add support for `Close`, `Get`, `List`, `New`, and `Update` methods on resource `V2.Core.Account`
* Add support for new resources `V2.Core.AccountLink`, `V2.Core.Account`, `V2.Core.Person`, `V2.Core.Vault.GbBankAccount`, `V2.Core.Vault.UsBankAccount`
* Add support for `New` method on resource `V2.Core.AccountLink`
* Add support for new thin events `V2CoreAccountIncludingConfigurationCustomerCapabilityStatusUpdatedEvent`, `V2CoreAccountIncludingConfigurationCustomerUpdatedEvent`, `V2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent`, `V2CoreAccountIncludingConfigurationMerchantUpdatedEvent`, `V2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent`, `V2CoreAccountIncludingConfigurationRecipientUpdatedEvent`, `V2CoreAccountIncludingIdentityUpdatedEvent`, and `V2CoreAccountIncludingRequirementsUpdatedEvent`
* Add support for new thin event `V2CoreAccountLinkCompletedEvent` with related object `V2.Core.AccountLink`
* Add support for new thin events `V2CoreAccountPersonCreatedEvent`, `V2CoreAccountPersonDeletedEvent`, and `V2CoreAccountPersonUpdatedEvent` with related object `V2.Core.Person`

### New APIs for Money CardManagement
* Add support for `Archive`, `Get`, `List`, and `Unarchive` methods on resource `V2.MoneyManagement.PayoutMethod`
* Add support for `Cancel`, `Get`, `List`, and `New` methods on resources `V2.MoneyManagement.OutboundPayment` and `V2.MoneyManagement.OutboundTransfer`
* Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2.MoneyManagement.OutboundSetupIntent`
* Add support for `Get` and `List` methods on resources `V2.MoneyManagement.Adjustment`, `V2.MoneyManagement.FinancialAccount`, `V2.MoneyManagement.ReceivedCredit`, `V2.MoneyManagement.ReceivedDebit`, `V2.MoneyManagement.TransactionEntry`, and `V2.MoneyManagement.Transaction`
* Add support for `Get`, `List`, and `New` methods on resources `V2.MoneyManagement.FinancialAddress` and `V2.MoneyManagement.InboundTransfer`
* Add support for `Get` method on resource `V2.MoneyManagement.PayoutMethodsBankAccountSpec`
* Add support for new thin events `V2MoneyManagementInboundTransferAvailableEvent`, `V2MoneyManagementInboundTransferBankDebitFailedEvent`, `V2MoneyManagementInboundTransferBankDebitProcessingEvent`, `V2MoneyManagementInboundTransferBankDebitQueuedEvent`, `V2MoneyManagementInboundTransferBankDebitReturnedEvent`, and `V2MoneyManagementInboundTransferBankDebitSucceededEvent` with related object `V2.MoneyManagement.InboundTransfer`
* Add support for new thin events `V2MoneyManagementOutboundPaymentCanceledEvent`, `V2MoneyManagementOutboundPaymentCreatedEvent`, `V2MoneyManagementOutboundPaymentFailedEvent`, `V2MoneyManagementOutboundPaymentPostedEvent`, and `V2MoneyManagementOutboundPaymentReturnedEvent` with related object `V2.MoneyManagement.OutboundPayment`
* Add support for new thin events `V2MoneyManagementOutboundTransferCanceledEvent`, `V2MoneyManagementOutboundTransferCreatedEvent`, `V2MoneyManagementOutboundTransferFailedEvent`, `V2MoneyManagementOutboundTransferPostedEvent`, and `V2MoneyManagementOutboundTransferReturnedEvent` with related object `V2.MoneyManagement.OutboundTransfer`
* Add support for new thin events `V2MoneyManagementReceivedCreditAvailableEvent`, `V2MoneyManagementReceivedCreditFailedEvent`, `V2MoneyManagementReceivedCreditReturnedEvent`, and `V2MoneyManagementReceivedCreditSucceededEvent` with related object `V2.MoneyManagement.ReceivedCredit`
* Add support for new thin events `V2MoneyManagementReceivedDebitCanceledEvent`, `V2MoneyManagementReceivedDebitFailedEvent`, `V2MoneyManagementReceivedDebitPendingEvent`, `V2MoneyManagementReceivedDebitSucceededEvent`, and `V2MoneyManagementReceivedDebitUpdatedEvent` with related object `V2.MoneyManagement.ReceivedDebit`
* Add support for new error types `AlreadyCanceledError`, `BlockedByStripeError`, `ControlledByDashboardError`, `FeatureNotEnabledError`, `FinancialAccountNotOpenError`, `InsufficientFundsError`, `InvalidPayoutMethodError`, `NotCancelableError`, and `RecipientNotNotifiableError`
* Add support for new resources `V2.FinancialAddressCreditSimulation`, `V2.FinancialAddressGeneratedMicrodeposits`, `V2.MoneyManagement.Adjustment`, `V2.MoneyManagement.FinancialAccount`, `V2.MoneyManagement.FinancialAddress`, `V2.MoneyManagement.InboundTransfer`, `V2.MoneyManagement.OutboundPaymentQuote`, `V2.MoneyManagement.OutboundPayment`, `V2.MoneyManagement.OutboundSetupIntent`, `V2.MoneyManagement.OutboundTransfer`, `V2.MoneyManagement.PayoutMethod`, `V2.MoneyManagement.PayoutMethodsBankAccountSpec`, `V2.MoneyManagement.ReceivedCredit`, `V2.MoneyManagement.ReceivedDebit`, `V2.MoneyManagement.TransactionEntry`, and `V2.MoneyManagement.Transaction`
* Add support for `New` method on resource `V2.MoneyManagement.OutboundPaymentQuote`
* Add support for new values `account_number`, `fedwire_routing_number`, and `routing_number` on enum `InvalidPaymentMethod.InvalidParam`
* Add support for new thin event `V2MoneyManagementFinancialAccountCreatedEvent` with related object `V2.MoneyManagement.FinancialAccount`
* Add support for new thin events `V2MoneyManagementFinancialAddressActivatedEvent` and `V2MoneyManagementFinancialAddressFailedEvent` with related object `V2.MoneyManagement.FinancialAddress`

## 81.5.0-beta.1 - 2025-03-18
This release changes the pinned API version to `2025-02-24.acacia`.

* [#1997](https://github.com/stripe/stripe-go/pull/1997) Beta SDK updates between Open API versions 1473 and 1505
  * Add support for `SucceedInputCollection` and `TimeoutInputCollection` test helper methods on resource `Terminal.Reader`
  * Add support for `TargetDate` on `OrderPaymentSettingsPaymentMethodOptionsAcssDebitParams`, `OrderPaymentSettingsPaymentMethodOptionsAcssDebit`, `OrderPaymentSettingsPaymentMethodOptionsSepaDebitParams`, and `OrderPaymentSettingsPaymentMethodOptionsSepaDebit`
* [#1987](https://github.com/stripe/stripe-go/pull/1987) Update generated code for beta
  * Add support for `SucceedInputCollection` and `TimeoutInputCollection` test helper methods on resource `Terminal.Reader`
  * Add support for new value `setup_intent_mobile_wallet_unsupported` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
  * Remove support for `Carrier`, `Phone`, and `TrackingNumber` on `CheckoutSessionCollectedInformationShippingDetails`
  * Add support for `InterchangeFeesAmount`, `NetTotalAmount`, `NetworkFeesAmount`, `OtherFeesAmount`, `OtherFeesCount`, and `TransactionAmount` on `IssuingSettlement`
  * Remove support for `InterchangeFees`, `NetTotal`, `NetworkFees`, and `TransactionVolume` on `IssuingSettlement`
  * Add support for `TargetDate` on `OrderPaymentSettingsPaymentMethodOptionsAcssDebitParams`, `OrderPaymentSettingsPaymentMethodOptionsAcssDebit`, `OrderPaymentSettingsPaymentMethodOptionsSepaDebitParams`, and `OrderPaymentSettingsPaymentMethodOptionsSepaDebit`
  * Add support for `ACHCreditTransfer`, `ACHDebit`, `ACSSDebit`, `AUBECSDebit`, `Affirm`, `AfterpayClearpay`, `Alipay`, `Alma`, `AmazonPay`, `BACSDebit`, `BLIK`, `Bancontact`, `Boleto`, `CardPresent`, `Card`, `CashApp`, `CustomerBalance`, `EPS`, `FPX`, `Giropay`, `Gopay`, `Grabpay`, `IDBankTransfer`, `IDEAL`, `InteracPresent`, `KakaoPay`, `Klarna`, `Konbini`, `KrCard`, `Link`, `MbWay`, `Mobilepay`, `Multibanco`, `NaverPay`, `OXXO`, `P24`, `PayByBank`, `PayNow`, `Payco`, `Paypal`, `Payto`, `Pix`, `PromptPay`, `Qris`, `Rechnung`, `RevolutPay`, `SEPACreditTransfer`, `SEPADebit`, `SamsungPay`, `Shopeepay`, `Sofort`, `StripeAccount`, `Swish`, `TWINT`, `USBankAccount`, `WeChatPay`, `WeChat`, and `Zip` on `PaymentAttemptRecordPaymentMethodDetails` and `PaymentRecordPaymentMethodDetails`
  * Change type of `PaymentAttemptRecordPaymentMethodDetailsCustom` and `PaymentRecordPaymentMethodDetailsCustom` from `nullable(PaymentsPrimitivesPaymentRecordsResourcePaymentMethodDetailsResourceCustomDetails)` to `PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCustomDetails`
  * Change type of `PaymentAttemptRecordPaymentMethodDetailsType` and `PaymentRecordPaymentMethodDetailsType` from `literal('custom')` to `string`
  * Add support for `Wifi` on `TerminalConfigurationParams` and `TerminalConfiguration`

## 81.4.0-beta.1 - 2025-02-07
* [#1972](https://github.com/stripe/stripe-go/pull/1972) Support time.Time instead of int64 for date fields
* [#1981](https://github.com/stripe/stripe-go/pull/1981) Revert "Support time.Time instead of int64 for date fields"
* [#1973](https://github.com/stripe/stripe-go/pull/1973) Update generated code for beta
  * Add support for `RejectedReason` on `AccountRiskControls`
  * Add support for `ProductTaxCodeSelector` on `AccountSessionComponentsParams`
  * Add support for `Prices` on `BillingCreditBalanceSummaryFilterApplicabilityScopeParams`, `BillingCreditGrantApplicabilityConfigScopeParams`, and `BillingCreditGrantApplicabilityConfigScope`
  * Add support for `BrandProduct` on `ChargePaymentMethodDetailsAmazonPayFundingCard` and `ChargePaymentMethodDetailsRevolutPayFundingCard`
  * Add support for `Restrictions` on `CheckoutSessionPaymentMethodOptionsCardParams` and `CheckoutSessionPaymentMethodOptionsCard`

## 81.3.0-beta.3 - 2025-01-23
* [#1971](https://github.com/stripe/stripe-go/pull/1971) Update generated code for beta
  * Remove support for `StripeAccount` on `TerminalReaderActionCollectPaymentMethod`, `TerminalReaderActionConfirmPaymentIntent`, `TerminalReaderActionProcessPaymentIntent`, and `TerminalReaderActionRefundPayment`

## 81.3.0-beta.2 - 2025-01-17
This release changes the pinned API version to `2025-01-27.acacia`.

* [#1963](https://github.com/stripe/stripe-go/pull/1963) Update generated code for beta
  * Add support for `PayByBankPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for `DirectorshipDeclaration` on `AccountCompanyParams` and `TokenAccountCompanyParams`
  * Add support for `ProofOfUltimateBeneficialOwnership` on `AccountDocumentsParams`
  * Add support for `TaxThresholdMonitoring` on `AccountSessionComponentsParams`
  * Add support for `FinancialAccountTransactions`, `FinancialAccount`, `IssuingCard`, and `IssuingCardsList` on `AccountSessionComponents`
  * Add support for new value `always_invoice` on enum `BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior`
  * Add support for `PayByBank` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `Discounts` on `CheckoutSession`
  * Add support for new value `SD` on enums `CheckoutSessionShippingAddressCollectionAllowedCountries` and `PaymentLinkShippingAddressCollectionAllowedCountries`
  * Add support for new value `pay_by_bank` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
  * Add support for `PhoneNumberCollection` on `PaymentLinkParams`
  * Add support for new value `pay_by_bank` on enum `PaymentLinkPaymentMethodTypes`
  * Add support for `Jpy` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`

## 81.3.0-beta.1 - 2025-01-09
* [#1958](https://github.com/stripe/stripe-go/pull/1958) Update generated code for beta
  * Add support for `Close` method on resource `Treasury.FinancialAccount`
  * Add support for `OwnershipExemptionReason` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
  * Add support for `DirectorshipDeclaration` on `AccountCompany`
  * Add support for `AdviceCode` on `ChargeOutcome`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
  * Remove support for value `always_invoice` from enum `BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior`
  * Add support for `BrandProduct` on `Card`, `SourceCardPresent`, `SourceCard`, and `SourceThreeDSecure`
  * Add support for `Country` on `ChargePaymentMethodDetailsPaypal`, `ConfirmationTokenPaymentMethodPreviewPaypal`, and `PaymentMethodPaypal`
  * Add support for new value `al_tin` on enums `CheckoutSessionCollectedInformationTaxIdsType` and `OrderTaxDetailsTaxIdsType`
  * Add support for `Nickname` on `TreasuryFinancialAccountParams` and `TreasuryFinancialAccount`
  * Add support for `ForwardingSettings` on `TreasuryFinancialAccountParams`
  * Add support for `IsDefault` on `TreasuryFinancialAccount`
  * Add support for `DestinationPaymentMethodData` on `TreasuryOutboundTransferParams`
  * Add support for `FinancialAccount` on `TreasuryOutboundTransferDestinationPaymentMethodDetails`
  * Change type of `TreasuryOutboundTransferDestinationPaymentMethodDetailsType` from `literal('us_bank_account')` to `enum('financial_account'|'us_bank_account')`
  * Add support for `OutboundTransfer` on `TreasuryReceivedCreditLinkedFlowsSourceFlowDetails`
  * Add support for new value `outbound_transfer` on enum `TreasuryReceivedCreditLinkedFlowsSourceFlowDetailsType`

## 81.2.0-beta.3 - 2024-12-12
This release changes the pinned API version to `2024-12-18.acacia`.

* [#1956](https://github.com/stripe/stripe-go/pull/1956) Update generated code for beta
  * Add support for `AllowRedisplay` on `Card` and `Source`
  * Add support for new values `am_tin`, `ao_tin`, `ba_tin`, `bb_tin`, `bs_tin`, `cd_nif`, `gn_nif`, `kh_tin`, `me_pib`, `mk_vat`, `mr_nif`, `np_pan`, `sn_ninea`, `sr_fin`, `tj_tin`, `ug_tin`, `zm_tin`, and `zw_tin` on enums `CheckoutSessionCollectedInformationTaxIdsType` and `OrderTaxDetailsTaxIdsType`
  * Add support for new value `network_fallback` on enum `IssuingAuthorizationRequestHistoryReason`
  * Remove support for `AmountRefunded` on `PaymentRecord`
  * Add support for `Account` on `TerminalReaderActionCollectPaymentMethod`, `TerminalReaderActionConfirmPaymentIntent`, `TerminalReaderActionProcessPaymentIntent`, and `TerminalReaderActionRefundPayment`

## 81.2.0-beta.2 - 2024-12-05
* [#1953](https://github.com/stripe/stripe-go/pull/1953) Update generated code for beta
  * Add support for `AutomaticIndirectTax` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for new values `payout_minimum_balance_hold` and `payout_minimum_balance_release` on enum `BalanceTransactionType`
  * Add support for `ReferencePrefix` on `CheckoutSessionPaymentMethodOptionsBacsDebitMandateOptionsParams`, `CheckoutSessionPaymentMethodOptionsBacsDebitMandateOptions`, `CheckoutSessionPaymentMethodOptionsSepaDebitMandateOptionsParams`, `CheckoutSessionPaymentMethodOptionsSepaDebitMandateOptions`, `OrderPaymentSettingsPaymentMethodOptionsSepaDebitMandateOptionsParams`, `OrderPaymentSettingsPaymentMethodOptionsSepaDebitMandateOptions`, `PaymentIntentConfirmPaymentMethodOptionsBacsDebitMandateOptionsParams`, `PaymentIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsBacsDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsBacsDebitMandateOptions`, `PaymentIntentPaymentMethodOptionsSepaDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsSepaDebitMandateOptions`, `SetupIntentConfirmPaymentMethodOptionsBacsDebitMandateOptionsParams`, `SetupIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams`, `SetupIntentPaymentMethodOptionsBacsDebitMandateOptionsParams`, `SetupIntentPaymentMethodOptionsBacsDebitMandateOptions`, `SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams`, and `SetupIntentPaymentMethodOptionsSepaDebitMandateOptions`
  * Add support for `DisabledReason` on `InvoiceAutomaticTax`, `SubscriptionAutomaticTax`, `SubscriptionScheduleDefaultSettingsAutomaticTax`, and `SubscriptionSchedulePhasesAutomaticTax`
  * Add support for `TrialPeriodDays` on `PaymentLinkSubscriptionDataParams`

## 81.2.0-beta.1 - 2024-11-21
* [#1952](https://github.com/stripe/stripe-go/pull/1952) Update generated code for beta
  * Add support for `NetworkAdviceCode` and `NetworkDeclineCode` on `ChargeOutcome`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
  * Add support for `Funding` on `ChargePaymentMethodDetailsAmazonPay` and `ChargePaymentMethodDetailsRevolutPay`
  * Add support for `AmountRequested` and `PartialAuthorization` on `ChargePaymentMethodDetailsCard`
  * Add support for `Metadata` on `CheckoutSessionLineItemsParams` and `LineItem`
  * Add support for `LineItems` on `CheckoutSessionParams`, `CheckoutSessionPermissionsUpdateParams`, and `CheckoutSessionPermissionsUpdate`
  * Add support for new value `invoice.overpaid` on enum `EventType`
  * Add support for `AdjustableQuantity` and `Display` on `LineItem`
  * Add support for `RequestPartialAuthorization` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCard`
  * Add support for `PaymentMethodOptions` on `PaymentIntentIncrementAuthorizationParams`

## 81.1.0-beta.3 - 2024-11-14
This release changes the pinned API version to `2024-11-20.acacia`.

* [#1950](https://github.com/stripe/stripe-go/pull/1950) Update generated code for beta
  * Add support for `AccountHolderAddress` and `BankAddress` on `FundingInstructionsBankTransferFinancialAddressesIban`, `FundingInstructionsBankTransferFinancialAddressesSortCode`, `FundingInstructionsBankTransferFinancialAddressesSpei`, `FundingInstructionsBankTransferFinancialAddressesZengin`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesIban`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSortCode`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSpei`, and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesZengin`
  * Add support for `AccountHolderName` on `FundingInstructionsBankTransferFinancialAddressesSpei` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSpei`

## 81.1.0-beta.2 - 2024-11-07
* [#1949](https://github.com/stripe/stripe-go/pull/1949) Update generated code for beta
  * Add support for new resources `Issuing.FraudLiabilityDebit`, `PaymentAttemptRecord`, and `PaymentRecord`
  * Add support for `Get` and `List` methods on resources `FraudLiabilityDebit` and `PaymentAttemptRecord`
  * Add support for `Get`, `ReportPaymentAttemptCanceled`, `ReportPaymentAttemptFailed`, `ReportPaymentAttemptGuaranteed`, `ReportPaymentAttempt`, and `ReportPayment` methods on resource `PaymentRecord`
  * Change type of `AccountFutureRequirementsDisabledReason` and `AccountRequirementsDisabledReason` from `string` to `enum`
  * Remove support for `MoneyMovement` on `AccountSessionComponentsFinancialAccountFeaturesParams`
  * Add support for `CardManagement`, `CardSpendDisputeManagement`, `CardholderManagement`, and `SpendControlManagement` on `AccountSessionComponentsIssuingCardFeaturesParams`
  * Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsIssuingCardsListFeaturesParams`
  * Add support for `AdaptivePricing` on `CheckoutSessionParams` and `CheckoutSession`
  * Add support for `MandateOptions` on `CheckoutSessionPaymentMethodOptionsBacsDebitParams`, `CheckoutSessionPaymentMethodOptionsBacsDebit`, `CheckoutSessionPaymentMethodOptionsSepaDebitParams`, and `CheckoutSessionPaymentMethodOptionsSepaDebit`
  * Add support for `RequestDecrementalAuthorization`, `RequestExtendedAuthorization`, `RequestIncrementalAuthorization`, `RequestMulticapture`, and `RequestOvercapture` on `CheckoutSessionPaymentMethodOptionsCardParams` and `CheckoutSessionPaymentMethodOptionsCard`
  * Add support for `CaptureMethod` on `CheckoutSessionPaymentMethodOptionsKakaoPayParams`, `CheckoutSessionPaymentMethodOptionsKrCardParams`, `CheckoutSessionPaymentMethodOptionsNaverPayParams`, `CheckoutSessionPaymentMethodOptionsPaycoParams`, and `CheckoutSessionPaymentMethodOptionsSamsungPayParams`
  * Add support for new value `li_vat` on enums `CheckoutSessionCollectedInformationTaxIdsType`, `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `OrderTaxDetailsTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
  * Add support for new values `invoice.payment_attempt_required` and `issuing_fraud_liability_debit.created` on enum `EventType`
  * Add support for `AccountHolderAddress`, `AccountHolderName`, `AccountType`, and `BankAddress` on `FundingInstructionsBankTransferFinancialAddressesAba`, `FundingInstructionsBankTransferFinancialAddressesSwift`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesAba`, and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSwift`
  * Add support for `PaymentRecordData` and `PaymentRecord` on `InvoiceAttachPaymentParams`
  * Remove support for `OutOfBandPayment` on `InvoiceAttachPaymentParams`
  * Add support for `AmountOverpaid` on `Invoice`
  * Add support for new value `custom` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
  * Add support for `MerchantAmount` and `MerchantCurrency` on `IssuingAuthorizationParams`
  * Add support for new value `link` on enums `PaymentIntentPaymentMethodOptionsCardNetwork`, `SetupIntentPaymentMethodOptionsCardNetwork`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCardNetwork`
  * Add support for `SubmitType` on `PaymentLinkParams`
  * Add support for new value `service_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationTaxBreakdownTaxRateDetailsTaxType`, `TaxRateTaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetailsTaxType`

## 81.1.0-beta.1 - 2024-10-29
This release changes the pinned API version to `2024-10-28.acacia`.

* [#1941](https://github.com/stripe/stripe-go/pull/1941) Do not allow setting API Version directly
  * `stripe.APIVersion` is no longer settable. If you were using this to set the beta headers, use the helper method `stripe.AddBetaVersion()` instead.
* [#1945](https://github.com/stripe/stripe-go/pull/1945) Update generated code for beta
  * Add support for `TriggerAction` method on resource `PaymentIntent`
  * Add support for `IDBankTransferPaymentsBca` and `IDBankTransferPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for `BankBcaOnboarding` on `AccountSettingsParams` and `AccountSettings`
  * Add support for `SendMoney` on `AccountSessionComponentsRecipientsFeaturesParams`
  * Remove support for value `payout_statement_descriptor_profanity` from enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
  * Add support for `IDBankTransfer` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `RefundDestinationDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentPaymentMethodDataParams`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
  * Add support for `Gopay`, `Qris`, and `Shopeepay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`

## 80.3.0-beta.1 - 2024-10-18
* [#1934](https://github.com/stripe/stripe-go/pull/1934) Update generated code for beta
  * Add support for `AlmaPayments`, `GopayPayments`, `KakaoPayPayments`, `KrCardPayments`, `NaverPayPayments`, `PaycoPayments`, `QrisPayments`, `SamsungPayPayments`, `ShopeepayPayments`, `TreasuryEvolve`, `TreasuryFifthThird`, and `TreasuryGoldmanSachs` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for `ScheduleAtPeriodEnd` on `BillingPortalConfigurationFeaturesSubscriptionUpdateParams` and `BillingPortalConfigurationFeaturesSubscriptionUpdate`
  * Add support for `Alma` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `RefundDestinationDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `Gopay`, `Qris`, and `Shopeepay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for new values `alma`, `gopay`, `qris`, and `shopeepay` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
  * Add support for `Metadata` on `ForwardingRequestParams`
  * Add support for new values `jp_credit_transfer`, `kakao_pay`, `kr_card`, `naver_pay`, and `payco` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
  * Remove support for value `expired` from enum `IssuingAuthorizationStatus`
  * Add support for new values `alma`, `gopay`, `qris`, and `shopeepay` on enum `PaymentLinkPaymentMethodTypes`
  * Add support for `AmazonPay` on `PaymentMethodDomain`
  * Add support for `ExternalReference` on `TaxFormListPayeeParams` and `TaxFormPayee`
  * Change type of `TaxFormListPayeeTypeParams` and `TaxFormPayeeType` from `literal('account')` to `enum('account'|'external_reference')`
  * Add support for `AuSerr`, `CaMrdp`, `EUDac7`, `GBMrdp`, and `NzMrdp` on `TaxForm`
  * Add support for new values `au_serr`, `ca_mrdp`, `eu_dac7`, `gb_mrdp`, and `nz_mrdp` on enum `TaxFormType`
  * Add support for `Pln` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
  * Add support for `Bank` on `TreasuryFinancialAccountFeaturesFinancialAddressesAbaParams`, `TreasuryFinancialAccountFeaturesFinancialAddressesAba`, and `TreasuryFinancialAccountUpdateFeaturesFinancialAddressesAbaParams`

## 80.2.0-beta.2 - 2024-10-08
* [#1932](https://github.com/stripe/stripe-go/pull/1932) Update generated code for beta
  * Add support for `SubmitCard` test helper method on resource `Issuing.Card`
  * Add support for `Groups` on `AccountParams` and `Account`
  * Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsAccountManagementFeaturesParams`, `AccountSessionComponentsAccountManagementFeatures`, `AccountSessionComponentsAccountOnboardingFeaturesParams`, `AccountSessionComponentsAccountOnboardingFeatures`, `AccountSessionComponentsBalancesFeaturesParams`, `AccountSessionComponentsBalancesFeatures`, `AccountSessionComponentsFinancialAccountFeaturesParams`, `AccountSessionComponentsNotificationBannerFeaturesParams`, `AccountSessionComponentsNotificationBannerFeatures`, `AccountSessionComponentsPayoutsFeaturesParams`, and `AccountSessionComponentsPayoutsFeatures`
  * Add support for `CardSpendDisputeManagement` and `SpendControlManagement` on `AccountSessionComponentsIssuingCardsListFeaturesParams`
  * Add support for new value `payout_statement_descriptor_profanity` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
  * Add support for `KakaoPay` and `KrCard` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `NaverPay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `Payco` and `SamsungPay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for new values `by_tin`, `ma_vat`, `md_vat`, `tz_vat`, `uz_tin`, and `uz_vat` on enums `CheckoutSessionCollectedInformationTaxIdsType`, `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `OrderTaxDetailsTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
  * Add support for new values `kakao_pay`, `kr_card`, `naver_pay`, `payco`, and `samsung_pay` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
  * Add support for new value `refund.failed` on enum `EventType`
  * Add support for `Metadata` on `ForwardingRequest`
  * Add support for new value `expired` on enum `IssuingAuthorizationStatus`
  * Add support for `LineItems` on `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypal`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, `PaymentIntentPaymentMethodOptionsPaypalParams`, and `PaymentIntentPaymentMethodOptionsPaypal`
  * Add support for new value `retail_delivery_fee` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationTaxBreakdownTaxRateDetailsTaxType`, `TaxRateTaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetailsTaxType`
  * Add support for `FlatAmount` and `RateType` on `TaxCalculationTaxBreakdownTaxRateDetails` and `TaxRate`
  * Add support for `By`, `Cr`, `Ec`, `Ma`, `Md`, `RU`, `Rs`, `Tz`, and `Uz` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
  * Add support for new value `state_retail_delivery_fee` on enum `TaxRegistrationCountryOptionsUsType`

## 80.2.0-beta.1 - 2024-10-03
This release changes the pinned API version to `2024-09-30.acacia`.

* [#1930](https://github.com/stripe/stripe-go/pull/1930) Updates beta branch with changes in master
  * Add support for `ReportingChart` on `AccountSessionComponentsParams`
  * Remove support for `FromSchedule` on `QuoteSubscriptionData`
  * Add support for `AllowRedisplay` on `TerminalReaderCollectPaymentMethodCollectConfigParams`
  * Moved raw request functionality from `preview` package to `rawrequest` package, and removed `preview`

## 79.13.0-beta.1 - 2024-09-18
* [#1920](https://github.com/stripe/stripe-go/pull/1920) Update generated code for beta
  * Remove support for resource `QuotePhase`
  * Remove support for `Get` and `ListLineItems` methods on resource `QuotePhase`
  * Add support for `SendMoney` and `TransferBalance` on `AccountSessionComponentsFinancialAccountFeaturesParams`
  * Add support for new value `rechnung` on enum `PaymentLinkPaymentMethodTypes`

## 79.12.0-beta.1 - 2024-09-13
* [#1915](https://github.com/stripe/stripe-go/pull/1915) Use pinned version of tools in Makefile and ci.yml (beta)
* [#1911](https://github.com/stripe/stripe-go/pull/1911) Update generated code for beta
  * Add support for new resources `Issuing.DisputeSettlementDetail` and `Issuing.Settlement`
  * Add support for `Get` and `List` methods on resource `DisputeSettlementDetail`
  * Remove support for `List` method on resource `QuotePhase`
  * Add support for new values `issuing_dispute_settlement_detail.created`, `issuing_dispute_settlement_detail.updated`, `issuing_settlement.created`, and `issuing_settlement.updated` on enum `EventType`
  * Add support for `Settlement` on `IssuingTransactionListParams` and `IssuingTransaction`

## 79.11.0-beta.1 - 2024-09-05
* [#1908](https://github.com/stripe/stripe-go/pull/1908) Update generated code for beta
  * Add support for new resources `Billing.MeterErrorReport` and `Terminal.ReaderCollectedData`
  * Add support for `Get` method on resource `ReaderCollectedData`
  * Add support for `Recipients` on `AccountSessionComponentsParams`
  * Add support for new value `terminal_reader_collected_data_invalid` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
  * Add support for `BusinessName`, `Email`, `Phone`, and `TaxIDs` on `CheckoutSessionCollectedInformation`
  * Add support for new value `billing.meter_error_report.triggered` on enum `EventType`
  * Add support for `RegulatoryReportingFile` on `IssuingCreditUnderwritingRecordCorrectParams`, `IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams`, `IssuingCreditUnderwritingRecordReportDecisionParams`, and `IssuingCreditUnderwritingRecord`
  * Add support for new value `mb_way` on enum `PaymentLinkPaymentMethodTypes`
  * Remove support for `Rechnung` on `PaymentMethodParams`

## 79.9.0-beta.2 - 2024-08-22
* [#1907](https://github.com/stripe/stripe-go/pull/1907) Update generated code for beta
  * Add support for `MbWayPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for `MbWay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for new value `mb_way` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
  * Remove support for values `accepted`, `partner_rejected`, and `submitted` from enum `DisputeEvidenceDetailsEnhancedEligibilityVisaCompellingEvidence3Status`
  * Add support for new value `hr_oib` on enum `OrderTaxDetailsTaxIdsType`
  * Remove support for `Phases` on `QuoteParams`
  * Remove support for `FromSchedule` on `QuoteSubscriptionDataParams`

## 79.9.0-beta.1 - 2024-08-15
* [#1905](https://github.com/stripe/stripe-go/pull/1905) Update generated code for beta
  * Add support for `CapitalFinancingApplication` and `CapitalFinancing` on `AccountSessionComponentsParams`
  * Add support for `Permissions` on `CheckoutSessionParams` and `CheckoutSession`
  * Add support for `CollectedInformation` on `CheckoutSessionParams` and `CheckoutSession`
  * Add support for `ShippingOptions` on `CheckoutSessionParams`

## 79.8.0-beta.1 - 2024-08-12
* ⚠️ [#1900](https://github.com/stripe/stripe-go/pull/1900) Update generated code for beta
  * Add support for `CapitalFinancingApplication` and `CapitalFinancing` on `AccountSessionComponents`
  * Add support for `Payto` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
  * Add support for new value `custom` on enum `CheckoutSessionUiMode`
  * ⚠️  Remove support for `RiskCorrelationID` on `PaymentIntentConfirmPaymentMethodOptionsRechnungParams`, `PaymentIntentPaymentMethodOptionsRechnungParams`, and `PaymentIntentPaymentMethodOptionsRechnung`
  * Add support for new value `payto` on enum `PaymentLinkPaymentMethodTypes`

## 79.7.0-beta.1 - 2024-08-01
* [#1898](https://github.com/stripe/stripe-go/pull/1898) Update generated code for beta
  * Add support for `AttachPayment` method on resource `Invoice`
  * Add support for `AppInstall` and `AppViewport` on `AccountSessionComponentsParams`
  * Remove support for `PartnerRejectedDetails` on `DisputeEvidenceDetailsEnhancedEligibilityVisaCompellingEvidence3`
  * Add support for `LinesInvalid` on `QuoteStatusDetailsStaleLastReason`
  * Add support for new value `lines_invalid` on enum `QuoteStatusDetailsStaleLastReasonType`
  * Add support for `LastPriceMigrationError` on `SubscriptionSchedule` and `Subscription`

## 79.6.1-beta.1 - 2024-07-25
* [#1893](https://github.com/stripe/stripe-go/pull/1893) Update generated code for beta
  * Add support for new resources `Billing.AlertTriggered` and `Billing.Alert`
  * Add support for `Activate`, `Archive`, `Deactivate`, `Get`, `List`, and `New` methods on resource `Alert`
  * Add support for new values `issuing.account_closed_for_not_providing_business_model_clarification`, `issuing.account_closed_for_not_providing_url_clarification`, and `issuing.account_closed_for_not_providing_use_case_clarification` on enum `AccountNoticeReason`
  * Add support for `DisplayName` on `TreasuryFinancialAccountParams` and `TreasuryFinancialAccount`

## 79.6.0-beta.1 - 2024-07-25
* [#1889](https://github.com/stripe/stripe-go/pull/1889) Update generated code for beta
  * Add support for new resource `Tax.Association`
  * Add support for `Find` method on resource `Association`
  * Add support for `Capital` on `AccountSettingsParams` and `AccountSettings`
  * Add support for `AsyncWorkflows` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentDecrementAuthorizationParams`, `PaymentIntentIncrementAuthorizationParams`, `PaymentIntentParams`, and `PaymentIntent`
  * Add support for `Payto` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`

## 79.4.0-beta.1 - 2024-07-11
* [#1885](https://github.com/stripe/stripe-go/pull/1885) Update generated code for beta
  * Add support for new value `not_qualified` on enum `DisputeEvidenceDetailsEnhancedEligibilityVisaCompellingEvidence3Status`

## 79.3.0-beta.1 - 2024-07-05
* ⚠️ [#1882](https://github.com/stripe/stripe-go/pull/1882) Update generated code for beta
  * ⚠️ Remove support for `PaymentMethodUpdate` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`. Users are expected to completely migrate from using `payment_method_update`.
  * Add support for new resource `FinancialConnections.Institution`
  * Add support for `Get` and `List` methods on resource `Institution`
  * Add support for `Institution` on `CheckoutSessionPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`, `FinancialConnectionsSessionFiltersParams`, `FinancialConnectionsSessionFilters`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`, `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`, `SetupIntentConfirmPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`, `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`
  * Add support for `PaymentMethodAllowRedisplayFilters`, `PaymentMethodRedisplayLimit`, `PaymentMethodRedisplay`, and `PaymentMethodSaveUsage` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`
  * Add support for new value `balance` on enum `FinancialConnectionsAccountSubscriptions`

## 79.2.0-beta.1 - 2024-06-27
This release changes the pinned API version to `2024-06-20`.

* [#1877](https://github.com/stripe/stripe-go/pull/1877) Update generated code for beta
  * Remove support for `PaymentMethodSetAsDefault` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`
  * Add support for new value `ch_uid` on enum `OrderTaxDetailsTaxIdsType`

## 78.12.0-beta.1 - 2024-06-13
* [#1872](https://github.com/stripe/stripe-go/pull/1872) Update generated code for beta
  * Add support for new value `de_stn` on enum `OrderTaxDetailsTaxIdsType`

## 78.11.0-beta.1 - 2024-06-06
* [#1869](https://github.com/stripe/stripe-go/pull/1869) Update generated code for beta
  * Add support for `TWINT` on `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`

## 78.10.0-beta.1 - 2024-05-30
* [#1867](https://github.com/stripe/stripe-go/pull/1867) Update generated code for beta
  * Keeping up with the changes from version 78.9.0

## 78.9.0-beta.1 - 2024-05-23
* [#1865](https://github.com/stripe/stripe-go/pull/1865) Update generated code for beta

## 78.8.0-beta.1 - 2024-05-16
* [#1861](https://github.com/stripe/stripe-go/pull/1861) Update generated code for beta

## 78.7.0-beta.1 - 2024-05-09
* [#1857](https://github.com/stripe/stripe-go/pull/1857) Update generated code for beta
  * No new beta features. Merging changes from the main branch.

## 78.6.0-beta.1 - 2024-05-02
* [#1854](https://github.com/stripe/stripe-go/pull/1854) Update generated code for beta
  * Add support for `RechnungPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for `Rechnung` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `Multibanco` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
  * Add support for new value `rechnung` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`

## 78.5.0-beta.1 - 2024-04-25
* [#1850](https://github.com/stripe/stripe-go/pull/1850) Update generated code for beta
  * Add support for `PaymentMethodSettings` on `AccountSessionComponentsParams`
  * Add support for `CancelSubscriptionSchedule` on `QuoteLine` and `QuoteLinesParams`

## 78.4.0-beta.1 - 2024-04-18
* [#1848](https://github.com/stripe/stripe-go/pull/1848) Update generated code for beta
  * Add support for `CapitalOverview`, `TaxRegistrations`, and `TaxSettings` on `AccountSessionComponentsParams`
  * Add support for `ExternalAccountCollection` on `AccountSessionComponentsFinancialAccountFeaturesParams`
  * Add support for `SubscriptionTrialFromPlan` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`

## 78.2.0-beta.1 - 2024-04-11
This release changes the pinned API version to `2024-04-10`.

* [#1843](https://github.com/stripe/stripe-go/pull/1843) Update generated code for beta
  * Add support for `Get` method on resources `Entitlements.ActiveEntitlement` and `Entitlements.Feature`
  * Add support for `Fees`, `Losses`, `RequirementCollection`, and `StripeDashboard` on `AccountControllerParams`
  * Add support for new values `bh_vat`, `kz_bin`, `ng_tin`, and `om_vat` on enum `OrderTaxDetailsTaxIdsType`
  * Add support for `HostedVoucherURL` on `PaymentIntentNextActionMultibancoDisplayDetails`
  * Add support for `Toggles` on `TerminalReaderActionCollectInputsInputs` and `TerminalReaderCollectInputsInputsParams`
  * Add support for `Email`, `Numeric`, `Phone`, and `Text` on `TerminalReaderActionCollectInputsInputs`

## 76.25.0-beta.1 - 2024-04-04
* [#1839](https://github.com/stripe/stripe-go/pull/1839) Update generated code for beta
  * Add support for `Update` method on resource `Entitlements.Feature`
  * Add support for `RiskControls` on `AccountParams` and `Account`
  * Add support for `PromotionCode` on `InvoiceAddLinesLinesDiscountsParams`, `InvoiceUpdateLinesLinesDiscountsParams`, `QuoteLineItemsDiscountsParams`, `QuoteLinesActionsAddDiscountParams`, and `QuotePhasesLineItemsDiscountsParams`

## 76.24.0-beta.1 - 2024-03-28
* [#1831](https://github.com/stripe/stripe-go/pull/1831) Update generated code for beta
  * Add support for `FinancialAccountTransactions`, `FinancialAccount`, `IssuingCard`, and `IssuingCardsList` on `AccountSessionComponentsParams`
  * Remove support for `SubscriptionBillingCycleAnchor`, `SubscriptionCancelAtPeriodEnd`, `SubscriptionCancelAt`, `SubscriptionCancelNow`, `SubscriptionDefaultTaxRates`, `SubscriptionItems`, `SubscriptionPrebilling`, `SubscriptionProrationBehavior`, `SubscriptionProrationDate`, `SubscriptionResumeAt`, `SubscriptionStartDate`, and `SubscriptionTrialEnd` on `InvoiceCreatePreviewParams`

## 76.23.0-beta.1 - 2024-03-21
* [#1827](https://github.com/stripe/stripe-go/pull/1827) Update generated code for beta
  * Add support for new resources `Entitlements.ActiveEntitlementSummary` and `Entitlements.ActiveEntitlement`
  * Add support for `List` method on resource `ActiveEntitlement`
  * Add support for `Mobilepay` on `ConfirmationTokenPaymentMethodDataParams` and `ConfirmationTokenPaymentMethodPreview`
  * Add support for `UseStripeSDK` on `ConfirmationToken`
  * Remove support for `PaymentMethod` on `ConfirmationToken`
  * Change type of `ConfirmationTokenMandateData` from `ConfirmationTokensResourceMandateData` to `nullable(ConfirmationTokensResourceMandateData)`
  * Add support for new value `mobilepay` on enum `ConfirmationTokenPaymentMethodPreviewType`
  * Add support for `Metadata` on `EntitlementsFeatureParams` and `EntitlementsFeature`
  * Add support for `Active` on `EntitlementsFeature`
  * Add support for new value `entitlements.active_entitlement_summary.updated` on enum `EventType`
  * Remove support for value `customer.entitlement_summary.updated` from enum `EventType`

## 76.22.0-beta.1 - 2024-03-14
* [#1825](https://github.com/stripe/stripe-go/pull/1825) Update generated code for beta
  * Add support for new resources `Billing.MeterEventAdjustment`, `Billing.MeterEvent`, and `Billing.Meter`
  * Add support for `Deactivate`, `Get`, `List`, `New`, `Reactivate`, and `Update` methods on resource `Meter`
  * Add support for `New` method on resources `MeterEventAdjustment` and `MeterEvent`
  * Add support for `New` test helper method on resource `ConfirmationToken`
  * Add support for `AddLines`, `RemoveLines`, and `UpdateLines` methods on resource `Invoice`
  * Add support for `Multibanco` on `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `RefundDestinationDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for new value `multibanco` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
  * Add support for `SecondLine` on `IssuingPhysicalBundleFeatures`
  * Add support for `MultibancoDisplayDetails` on `PaymentIntentNextAction`
  * Add support for `Meter` on `PlanParams`, `Plan`, `PriceListRecurringParams`, `PriceRecurringParams`, and `PriceRecurring`

## 76.21.0-beta.1 - 2024-03-07
* [#1822](https://github.com/stripe/stripe-go/pull/1822) Update generated code for beta
  * Add support for new value `billing_period_end` on enum `QuoteLineEndsAtType`

## 76.20.0-beta.1 - 2024-02-29
* [#1819](https://github.com/stripe/stripe-go/pull/1819) Add helper to set beta version
* [#1816](https://github.com/stripe/stripe-go/pull/1816) Update generated code for beta
  * Remove support for resource `Entitlements.Event`
  * Change type of `ConfirmationTokenMandateData` from `nullable(ConfirmationTokensResourceMandateData)` to `ConfirmationTokensResourceMandateData`
  * Remove support for `Quantity` and `Type` on `EntitlementsFeatureParams` and `EntitlementsFeature`
  * Add support for `Livemode` on `IssuingPersonalizationDesign`
  * Add support for `ApplicationFeeAmount`, `Description`, `Metadata`, and `TransferData` on `PaymentIntentDecrementAuthorizationParams`
  * Add support for `EnableCustomerCancellation` on `TerminalReaderActionCollectPaymentMethodCollectConfig` and `TerminalReaderCollectPaymentMethodCollectConfigParams`

## 76.19.0-beta.1 - 2024-02-22
* [#1815](https://github.com/stripe/stripe-go/pull/1815) Update generated code for beta

## 76.18.0-beta.1 - 2024-02-16
* [#1813](https://github.com/stripe/stripe-go/pull/1813) Update generated code for beta
  * Add support for `DecrementAuthorization` method on resource `PaymentIntent`
  * Add support for `PaytoPayments` and `TWINTPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for `Payto` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
  * Add support for `TWINT` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `DecrementalAuthorization` on `ChargePaymentMethodDetailsCard`
  * Add support for `DisplayBrand` on `ConfirmationTokenPaymentMethodPreviewCard`
  * Add support for new values `payto` and `twint` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
  * Add support for new value `no_voec` on enum `OrderTaxDetailsTaxIdsType`
  * Add support for `RequestDecrementalAuthorization` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCard`

## 76.17.0-beta.1 - 2024-02-08
* [#1810](https://github.com/stripe/stripe-go/pull/1810) Update generated code for beta
  * Add support for `PaymentMethodOptions` on `ConfirmationToken`
  * Add support for new value `velobank` on enum `ConfirmationTokenPaymentMethodPreviewP24Bank`

## 76.16.0-beta.1 - 2024-02-01
* [#1806](https://github.com/stripe/stripe-go/pull/1806) Update generated code for beta
  * Add support for new resources `Entitlements.Event` and `Entitlements.Feature`
  * Add support for `New` method on resource `Event`
  * Add support for `List` and `New` methods on resource `Feature`
  * Add support for `Swish` on `ConfirmationTokenPaymentMethodPreview`
  * Add support for new value `swish` on enum `ConfirmationTokenPaymentMethodPreviewType`
  * Add support for new value `customer.entitlement_summary.updated` on enum `EventType`
  * Add support for `AccountTaxIDs` on `InvoiceCreatePreviewScheduleDetailsPhasesInvoiceSettingsParams`, `InvoiceUpcomingLinesScheduleDetailsPhasesInvoiceSettingsParams`, and `InvoiceUpcomingScheduleDetailsPhasesInvoiceSettingsParams`
  * Add support for `Feature` on `ProductFeaturesParams` and `ProductFeatures`

## 76.15.0-beta.1 - 2024-01-25
* [#1801](https://github.com/stripe/stripe-go/pull/1801) Update generated code for beta
  Release specs are identical.
* [#1799](https://github.com/stripe/stripe-go/pull/1799) Update generated code for beta
  * Add support for new value `nn` on enum `ConfirmationTokenPaymentMethodPreviewIdealBank`
  * Add support for new value `NNBANL2G` on enum `ConfirmationTokenPaymentMethodPreviewIdealBic`
* [#1802](https://github.com/stripe/stripe-go/pull/1802) Update generated code for beta
  * Add support for `CreatePreview` method on resource `Invoice`
  * Add support for `ChargedOffAt` on `CapitalFinancingOffer`
  * Add support for new values `disabled` and `enabled` on enums `CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemove`, `CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSave`, `CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefault`, and `CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdate`
  * Remove support for values `auto` and `never` from enums `CustomerSessionComponentsPaymentElementFeaturesPaymentMethodRemove`, `CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSave`, `CustomerSessionComponentsPaymentElementFeaturesPaymentMethodSetAsDefault`, and `CustomerSessionComponentsPaymentElementFeaturesPaymentMethodUpdate`
  * Add support for `EnhancedEvidence` on `DisputeEvidenceParams` and `DisputeEvidence`
  * Add support for `EnhancedEligibilityTypes` on `Dispute`
  * Add support for `EnhancedEligibility` on `DisputeEvidenceDetails`
  * Add support for `PromotionCode` on `InvoiceUpcomingLinesScheduleDetailsAmendmentsDiscountActionsAddParams`, `InvoiceUpcomingLinesScheduleDetailsAmendmentsDiscountActionsRemoveParams`, `InvoiceUpcomingLinesScheduleDetailsAmendmentsDiscountActionsSetParams`, `InvoiceUpcomingLinesScheduleDetailsAmendmentsItemActionsAddDiscountsParams`, `InvoiceUpcomingLinesScheduleDetailsAmendmentsItemActionsSetDiscountsParams`, `InvoiceUpcomingLinesScheduleDetailsPhasesAddInvoiceItemsDiscountsParams`, `InvoiceUpcomingLinesScheduleDetailsPhasesDiscountsParams`, `InvoiceUpcomingLinesScheduleDetailsPhasesItemsDiscountsParams`, `InvoiceUpcomingLinesSubscriptionDetailsItemsDiscountsParams`, `InvoiceUpcomingLinesSubscriptionItemsDiscountsParams`, `InvoiceUpcomingScheduleDetailsAmendmentsDiscountActionsAddParams`, `InvoiceUpcomingScheduleDetailsAmendmentsDiscountActionsRemoveParams`, `InvoiceUpcomingScheduleDetailsAmendmentsDiscountActionsSetParams`, `InvoiceUpcomingScheduleDetailsAmendmentsItemActionsAddDiscountsParams`, `InvoiceUpcomingScheduleDetailsAmendmentsItemActionsSetDiscountsParams`, `InvoiceUpcomingScheduleDetailsPhasesAddInvoiceItemsDiscountsParams`, `InvoiceUpcomingScheduleDetailsPhasesDiscountsParams`, `InvoiceUpcomingScheduleDetailsPhasesItemsDiscountsParams`, `InvoiceUpcomingSubscriptionDetailsItemsDiscountsParams`, `InvoiceUpcomingSubscriptionItemsDiscountsParams`, `QuoteLineActionsAddDiscount`, `QuoteLineActionsAddItemDiscounts`, `QuoteLineActionsRemoveDiscount`, `QuoteLineActionsSetDiscounts`, `QuoteLineActionsSetItemsDiscounts`, `QuoteLinesActionsAddItemDiscountsParams`, `QuoteLinesActionsRemoveDiscountParams`, `QuoteLinesActionsSetDiscountsParams`, `QuoteLinesActionsSetItemsDiscountsParams`, `QuotePhasesDiscountsParams`, `SubscriptionAddInvoiceItemsDiscountsParams`, `SubscriptionDiscountsParams`, `SubscriptionItemDiscountsParams`, `SubscriptionItemsDiscountsParams`, `SubscriptionScheduleAmendAmendmentsDiscountActionsAddParams`, `SubscriptionScheduleAmendAmendmentsDiscountActionsRemoveParams`, `SubscriptionScheduleAmendAmendmentsDiscountActionsSetParams`, `SubscriptionScheduleAmendAmendmentsItemActionsAddDiscountsParams`, `SubscriptionScheduleAmendAmendmentsItemActionsSetDiscountsParams`, `SubscriptionSchedulePhasesAddInvoiceItemsDiscountsParams`, `SubscriptionSchedulePhasesAddInvoiceItemsDiscounts`, `SubscriptionSchedulePhasesDiscountsParams`, `SubscriptionSchedulePhasesDiscounts`, `SubscriptionSchedulePhasesItemsDiscountsParams`, and `SubscriptionSchedulePhasesItemsDiscounts`

## 76.14.0-beta.1 - 2024-01-18
* [#1795](https://github.com/stripe/stripe-go/pull/1795) Update generated code for beta
  * Add support for `Amount` on `ChargeCapturePaymentDetailsFlightSegmentsParams`, `ChargePaymentDetailsFlightSegmentsParams`, `PaymentIntentCapturePaymentDetailsFlightSegmentsParams`, `PaymentIntentConfirmPaymentDetailsFlightSegmentsParams`, and `PaymentIntentPaymentDetailsFlightSegmentsParams`
  * Add support for `NumberOfRooms` and `RoomClass` on `ChargeCapturePaymentDetailsLodgingParams`, `ChargePaymentDetailsLodgingParams`, `PaymentIntentCapturePaymentDetailsLodgingParams`, `PaymentIntentConfirmPaymentDetailsLodgingParams`, and `PaymentIntentPaymentDetailsLodgingParams`
  * Add support for `BuyButton` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
  * Add support for new values `high_risk_industry`, `insufficient_margin_ratio`, `insufficient_operating_profit`, `insufficient_reserves`, `insufficient_time_in_network`, `lacking_cash_account`, and `poor_payment_history_with_platform` on enum `IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasons`
  * Add support for new values `high_risk_industry`, `insufficient_margin_ratio`, `insufficient_operating_profit`, `insufficient_reserves`, `insufficient_time_in_network`, and `lacking_cash_account` on enums `IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasons` and `IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasons`
* [#1801](https://github.com/stripe/stripe-go/pull/1801) Update generated code for beta
  Release specs are identical.
* [#1799](https://github.com/stripe/stripe-go/pull/1799) Update generated code for beta
  * Add support for new value `nn` on enum `ConfirmationTokenPaymentMethodPreviewIdealBank`
  * Add support for new value `NNBANL2G` on enum `ConfirmationTokenPaymentMethodPreviewIdealBic`

## 76.13.0-beta.1 - 2024-01-12
* [#1793](https://github.com/stripe/stripe-go/pull/1793) Update generated code for beta
* [#1795](https://github.com/stripe/stripe-go/pull/1795) Update generated code for beta
  * Add support for `Amount` on `ChargeCapturePaymentDetailsFlightSegmentsParams`, `ChargePaymentDetailsFlightSegmentsParams`, `PaymentIntentCapturePaymentDetailsFlightSegmentsParams`, `PaymentIntentConfirmPaymentDetailsFlightSegmentsParams`, and `PaymentIntentPaymentDetailsFlightSegmentsParams`
  * Add support for `NumberOfRooms` and `RoomClass` on `ChargeCapturePaymentDetailsLodgingParams`, `ChargePaymentDetailsLodgingParams`, `PaymentIntentCapturePaymentDetailsLodgingParams`, `PaymentIntentConfirmPaymentDetailsLodgingParams`, and `PaymentIntentPaymentDetailsLodgingParams`
  * Add support for `BuyButton` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
  * Add support for new values `high_risk_industry`, `insufficient_margin_ratio`, `insufficient_operating_profit`, `insufficient_reserves`, `insufficient_time_in_network`, `lacking_cash_account`, and `poor_payment_history_with_platform` on enum `IssuingCreditUnderwritingRecordDecisionApplicationRejectedReasons`
  * Add support for new values `high_risk_industry`, `insufficient_margin_ratio`, `insufficient_operating_profit`, `insufficient_reserves`, `insufficient_time_in_network`, and `lacking_cash_account` on enums `IssuingCreditUnderwritingRecordDecisionCreditLimitDecreasedReasons` and `IssuingCreditUnderwritingRecordDecisionCreditLineClosedReasons`

## 76.12.0-beta.1 - 2024-01-04
* [#1793](https://github.com/stripe/stripe-go/pull/1793) Update generated code for beta
  * Updated stable APIs to the latest version

## 76.11.0-beta.1 - 2023-12-22
* [#1791](https://github.com/stripe/stripe-go/pull/1791) Update generated code for beta
  * Add support for `CapitalFinancingPromotion` on `AccountSessionComponentsParams` and `AccountSessionComponents`
  * Add support for new value `shipping_address_invalid` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
  * Change type of `InvoiceIssuer` and `SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer` from `nullable(ConnectAccountReference)` to `ConnectAccountReference`
  * Change type of `PaymentLinkSubscriptionDataInvoiceSettings` from `nullable(PaymentLinksResourceSubscriptionDataInvoiceSettings)` to `PaymentLinksResourceSubscriptionDataInvoiceSettings`
  * Add support for `ShipFromDetails` on `TaxCalculationParams`, `TaxCalculation`, and `TaxTransaction`

## 76.10.0-beta.1 - 2023-12-14
* [#1779](https://github.com/stripe/stripe-go/pull/1779) Track usage of `raw_request`
* [#1785](https://github.com/stripe/stripe-go/pull/1785) Update generated code for beta
  * Add support for `PreviewMode` and `SubscriptionDetails` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
  * Remove support for `SubscriptionTrialFromPlan` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
  * Add support for `BillingBehavior`, `EndBehavior`, and `ProrationBehavior` on `InvoiceUpcomingLinesScheduleDetailsParams` and `InvoiceUpcomingScheduleDetailsParams`

## 76.9.0-beta.1 - 2023-12-08
* [#1776](https://github.com/stripe/stripe-go/pull/1776) Update generated code for beta
  * Add support for `Get` method on resource `FinancialConnections.Transaction`
  * Remove support for `IssuingCard` and `IssuingCardsList` on `AccountSessionComponentsParams`
  * Add support for `PaymentMethodRemove`, `PaymentMethodSave`, and `PaymentMethodSetAsDefault` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`
  * Remove support for `PaymentMethodDetach` and `PaymentMethodSetAsCustomerDefault` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`

## 76.8.0-beta.1 - 2023-11-30
* [#1773](https://github.com/stripe/stripe-go/pull/1773) Update generated code for beta

## 76.7.0-beta.1 - 2023-11-21
* [#1770](https://github.com/stripe/stripe-go/pull/1770) Update generated code for beta
  * Rename `Receipient` to `Recipient` beneath `PaymentDetails` on `Charge` and `PaymentIntent` APIs.
  * Add support for `Components` on `CustomerSessionParams` and `CustomerSession`

## 76.6.0-beta.1 - 2023-11-16
* [#1765](https://github.com/stripe/stripe-go/pull/1765) Update generated code for beta
  * Add support for `IssuingCard` and `IssuingCardsList` on `AccountSessionComponentsParams`
  * Add support for `EventDetails` and `Subscription` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
  * Add support for `Affiliate` and `Delivery` on `ChargeCapturePaymentDetailsCarRentalParams`, `ChargeCapturePaymentDetailsFlightParams`, `ChargeCapturePaymentDetailsLodgingParams`, `ChargePaymentDetailsCarRentalParams`, `ChargePaymentDetailsFlightParams`, `ChargePaymentDetailsLodgingParams`, `PaymentIntentCapturePaymentDetailsCarRentalParams`, `PaymentIntentCapturePaymentDetailsFlightParams`, `PaymentIntentCapturePaymentDetailsLodgingParams`, `PaymentIntentConfirmPaymentDetailsCarRentalParams`, `PaymentIntentConfirmPaymentDetailsFlightParams`, `PaymentIntentConfirmPaymentDetailsLodgingParams`, `PaymentIntentPaymentDetailsCarRentalParams`, `PaymentIntentPaymentDetailsCarRental`, `PaymentIntentPaymentDetailsFlightParams`, and `PaymentIntentPaymentDetailsLodgingParams`
  * Add support for `Drivers` on `ChargeCapturePaymentDetailsCarRentalParams`, `ChargePaymentDetailsCarRentalParams`, `PaymentIntentCapturePaymentDetailsCarRentalParams`, `PaymentIntentConfirmPaymentDetailsCarRentalParams`, `PaymentIntentPaymentDetailsCarRentalParams`, and `PaymentIntentPaymentDetailsCarRental`
  * Add support for `Passengers` on `ChargeCapturePaymentDetailsFlightParams`, `ChargeCapturePaymentDetailsLodgingParams`, `ChargePaymentDetailsFlightParams`, `ChargePaymentDetailsLodgingParams`, `PaymentIntentCapturePaymentDetailsFlightParams`, `PaymentIntentCapturePaymentDetailsLodgingParams`, `PaymentIntentConfirmPaymentDetailsFlightParams`, `PaymentIntentConfirmPaymentDetailsLodgingParams`, `PaymentIntentPaymentDetailsFlightParams`, and `PaymentIntentPaymentDetailsLodgingParams`
  * Add support for `Created` on `CustomerSession`

## 76.5.0-beta.1 - 2023-11-10
* [#1763](https://github.com/stripe/stripe-go/pull/1763) Update generated code for beta
  * Add support for new value `quote.reestimate_failed` on enum `EventType`
  * Add support for `Metadata` on `QuotePhase` and `QuotePhasesParams`
  * Add support for `LastReestimationDetails` on `QuoteComputed`

## 76.4.0-beta.1 - 2023-11-02
* [#1761](https://github.com/stripe/stripe-go/pull/1761) Update generated code for beta
  * Add support for `AttachPaymentIntent` method on resource `Invoice`
  * Add support for `RevolutPay` on `ConfirmationTokenPaymentMethodPreview`
  * Add support for new value `revolut_pay` on enum `ConfirmationTokenPaymentMethodPreviewType`
  * Add support for `Refunds` on `CreditNoteParams`, `CreditNotePreviewLinesParams`, `CreditNotePreviewParams`, and `CreditNote`
  * Add support for `PostPaymentAmount` and `PrePaymentAmount` on `CreditNote`
  * Add support for new value `invoice.payment.overpaid` on enum `EventType`
  * Add support for `ScheduleDetails` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
  * Add support for `AmountsDue` on `InvoiceParams` and `Invoice`
  * Add support for `Payments` on `Invoice`
  * Add support for `Created` on `IssuingPersonalizationDesign`

## 76.3.0-beta.1 - 2023-10-26
* [#1758](https://github.com/stripe/stripe-go/pull/1758) Update generated code for beta
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

## 76.2.0-beta.1 - 2023-10-17
This release changes the pinned API version to `2023-10-16`.

* [#1757](https://github.com/stripe/stripe-go/pull/1757) Update generated code for beta
* [#1754](https://github.com/stripe/stripe-go/pull/1754) Update generated code for beta
  - Update pinned API version to `2023-10-16`

## 75.12.0-beta.1 - 2023-10-16
* [#1752](https://github.com/stripe/stripe-go/pull/1752) Update generated code for beta

## 75.11.0-beta.1 - 2023-10-11
* [#1745](https://github.com/stripe/stripe-go/pull/1745) Update generated code for beta
  * Add support for new resources `AccountNotice` and `Issuing.CreditUnderwritingRecord`
  * Add support for `Get`, `List`, and `Update` methods on resource `AccountNotice`
  * Add support for `Correct`, `CreateFromApplication`, `CreateFromProactiveReview`, `Get`, `List`, and `ReportDecision` methods on resource `CreditUnderwritingRecord`
  * Change type of `CheckoutSessionAutomaticTaxLiabilityAccount`, `CheckoutSessionInvoiceCreationInvoiceDataIssuerAccount`, `InvoiceAutomaticTaxLiabilityAccount`, `InvoiceIssuerAccount`, `QuoteAutomaticTaxLiabilityAccount`, `QuoteInvoiceSettingsIssuerAccount`, `SubscriptionAutomaticTaxLiabilityAccount`, `SubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityAccount`, `SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerAccount`, `SubscriptionSchedulePhasesAutomaticTaxLiabilityAccount`, and `SubscriptionSchedulePhasesInvoiceSettingsIssuerAccount` from `nullable(expandable($Account))` to `expandable($Account)`
  * Add support for new values `account_notice.created` and `account_notice.updated` on enum `EventType`
  * Add support for new values `local_amusement_tax` and `state_communications_tax` on enum `TaxRegistrationCountryOptionsUsType`

## 75.10.0-beta.1 - 2023-10-05
* [#1744](https://github.com/stripe/stripe-go/pull/1744) Update generated code for beta
  * Add support for `MarkDraft` and `MarkStale` methods on resource `Quote`
  * Remove support for `DraftQuote` and `MarkStaleQuote` methods on resource `Quote`
  * Add support for `Liability` on `CheckoutSessionAutomaticTaxParams` and `CheckoutSessionAutomaticTax`
  * Add support for `Issuer` on `CheckoutSessionInvoiceCreationInvoiceDataParams` and `CheckoutSessionInvoiceCreationInvoiceData`
  * Add support for `InvoiceSettings` on `CheckoutSessionSubscriptionDataParams`
  * Add support for `PersonalizationDesign` on `IssuingCardListParams`
  * Add support for `AllowBackdatedLines` on `QuoteParams` and `Quote`

## 75.9.0-beta.1 - 2023-09-28
* [#1742](https://github.com/stripe/stripe-go/pull/1742) Beta: fix incompatible combination
* [#1740](https://github.com/stripe/stripe-go/pull/1740) Update generated code for beta
  * Rename resources `Issuing.CardDesign` and `Issuing.CardBundle` to `Issuing.PersonalizationDesign` and `Issuing.PhysicalBundle`
  * Add support for `Features` on `AccountSessionComponentsAccountOnboardingParams`, `AccountSessionComponentsPaymentDetailsParams`, `AccountSessionComponentsPaymentDetails`, `AccountSessionComponentsPaymentsParams`, `AccountSessionComponentsPayments`, `AccountSessionComponentsPayoutsParams`, and `AccountSessionComponentsPayouts`
  * Add support for `Reason` on `Event`

## 75.8.0-beta.1 - 2023-09-21
* [#1737](https://github.com/stripe/stripe-go/pull/1737) Update generated code for beta
  * Remove support for `Customer` on `ConfirmationToken`
  * Add support for `Issuer` on `InvoiceParams`, `InvoiceUpcomingLinesParams`, `InvoiceUpcomingParams`, `Invoice`, `QuoteInvoiceSettingsParams`, `QuoteInvoiceSettings`, `SubscriptionScheduleDefaultSettingsInvoiceSettingsParams`, `SubscriptionScheduleDefaultSettingsInvoiceSettings`, `SubscriptionSchedulePhasesInvoiceSettingsParams`, and `SubscriptionSchedulePhasesInvoiceSettings`
  * Add support for `OnBehalfOf` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
  * Add support for `Liability` on `InvoiceAutomaticTaxParams`, `InvoiceAutomaticTax`, `InvoiceUpcomingAutomaticTaxParams`, `InvoiceUpcomingLinesAutomaticTaxParams`, `QuoteAutomaticTaxParams`, `QuoteAutomaticTax`, `SubscriptionAutomaticTaxParams`, `SubscriptionAutomaticTax`, `SubscriptionScheduleDefaultSettingsAutomaticTaxParams`, `SubscriptionScheduleDefaultSettingsAutomaticTax`, `SubscriptionSchedulePhasesAutomaticTaxParams`, and `SubscriptionSchedulePhasesAutomaticTax`
  * Add support for `InvoiceSettings` on `SubscriptionParams`

## 75.7.0-beta.1 - 2023-09-14
* [#1732](https://github.com/stripe/stripe-go/pull/1732) Update generated code for beta
  * Add support for new resource `ConfirmationToken`
  * Add support for `Get` method on resource `ConfirmationToken`
  * Add support for `New` method on resource `Issuing.CardDesign`
  * Add support for `RejectTestmode` test helper method on resource `Issuing.CardDesign`
  * Add support for new value `issuing_card_design.rejected` on enum `EventType`
  * Add support for `Features` on `IssuingCardBundle`
  * Add support for `Preferences` on `IssuingCardDesignListParams`, `IssuingCardDesignParams`, and `IssuingCardDesign`
  * Remove support for `Preference` on `IssuingCardDesignListParams`, `IssuingCardDesignParams`, and `IssuingCardDesign`
  * Add support for `CardBundle` on `IssuingCardDesignParams`
  * Add support for `CardLogo` and `CarrierText` on `IssuingCardDesignParams` and `IssuingCardDesign`
  * Change type of `IssuingCardDesignLookupKeyParams` and `IssuingCardDesignNameParams` from `string` to `emptyStringable(string)`
  * Add support for `RejectionReasons` on `IssuingCardDesign`
  * Add support for `ConfirmationToken` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, `SetupIntentConfirmParams`, and `SetupIntentParams`

## 75.5.0-beta.1 - 2023-09-07
* [#1730](https://github.com/stripe/stripe-go/pull/1730) Update generated code for beta
  Release specs are identical.
* [#1725](https://github.com/stripe/stripe-go/pull/1725) Update generated code for beta
  * Remove support for `SubmitCard` test helper method on resource `Issuing.Card`
  * Add support for `TaxForms` on `AccountSettingsParams` and `AccountSettings`
  * Add support for `CardDesign` on `IssuingCardParams`
  * Remove support for value `submitted` from enum `IssuingCardShippingStatus`
  * Add support for new value `platform_default` on enum `IssuingCardDesignPreference`

## 75.4.0-beta.1 - 2023-08-31
* [#1721](https://github.com/stripe/stripe-go/pull/1721) Update generated code for beta
  * Rename `quote.PreviewInvoices` and `quote.PreviewSubscriptionSchedules` to `quotepreviewinvoice.List` and `quotepreviewschedule.List`
  * Add support for `Components` on `AccountSessionParams` and `AccountSession`

## 75.0.0-beta.1 - 2023-08-24
This release changes the pinned API version to `2023-08-16`.

* [#1719](https://github.com/stripe/stripe-go/pull/1719) Move beta version back
* [#1715](https://github.com/stripe/stripe-go/pull/1715) Update generated code for beta
  * Add support for new resources `QuotePreviewInvoice` and `QuotePreviewSchedule`
  * Remove support for `AppliesTo` on `Invoice` and `SubscriptionSchedule`
  * Add support for `Cl`, `Co`, `ID`, `Kr`, `MX`, `My`, `Sa`, `TH`, `TR`, and `Vn` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
  * Remove support for `Hk` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`

## 74.31.0-beta.1 - 2023-08-10
* [#1701](https://github.com/stripe/stripe-go/pull/1701) Update generated code for beta
  * Add support for `Paypal` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`

## 74.30.0-beta.1 - 2023-08-03
* [#1697](https://github.com/stripe/stripe-go/pull/1697) Update generated code for beta
  * Add support for `SubmitCard` test helper method on resource `Issuing.Card`
  * Add support for `AddressValidation` on `IssuingCardShippingParams` and `IssuingCardShipping`
  * Add support for new value `submitted` on enum `IssuingCardShippingStatus`
  * Change type of `OrderDescriptionParams`, `OrderLineItemsProductDataDescriptionParams`, `OrderLineItemsProductDataTaxCodeParams`, `OrderShippingDetailsPhoneParams`, `PaymentMethodConfigurationListApplicationParams`, and `QuoteSubscriptionDataOverridesDescriptionParams` from `string` to `emptyStringable(string)`
  * Add support for `Reason` on `QuoteMarkStaleQuoteParams`
  * Add support for `MarkedStale` on `QuoteStatusDetailsStaleLastReason`

## 74.29.0-beta.1 - 2023-07-28
* [#1690](https://github.com/stripe/stripe-go/pull/1690) Update generated code for beta
  * Remove support for values `excluded_territory`, `jurisdiction_unsupported`, and `vat_exempt` from enums `OrderShippingCostTaxesTaxabilityReason`, `OrderTotalDetailsBreakdownTaxesTaxabilityReason`, and `QuotePhaseTotalDetailsBreakdownTaxesTaxabilityReason`
  * Add support for new value `ro_tin` on enum `OrderTaxDetailsTaxIdsType`
  * Add support for new values `email`, `numeric`, `phone`, and `text` on enum `TerminalReaderActionCollectInputsInputsType`
* [#1692](https://github.com/stripe/stripe-go/pull/1692) Update generated code for beta
  * Add support for new resource `Tax.Form`
  * Add support for `Get`, `List`, and `PDF` methods on resource `Form`
  * Add support for `PaymentMethodConfiguration` on `CheckoutSessionParams` and `SetupIntentParams`
  * Add support for `PaymentMethodConfigurationDetails` on `CheckoutSession` and `SetupIntent`
* [#1694](https://github.com/stripe/stripe-go/pull/1694) Update generated code for beta
  * Release specs are identical.

## 74.27.0-beta.1 - 2023-07-13
* [#1685](https://github.com/stripe/stripe-go/pull/1685) Update generated code for beta
* [#1689](https://github.com/stripe/stripe-go/pull/1689) Update generated code for beta
  Release specs are identical.
* [#1687](https://github.com/stripe/stripe-go/pull/1687) Update generated code for beta
  * Add support for new resource `PaymentMethodConfiguration`
  * Add support for `Get`, `List`, `New`, and `Update` methods on resource `PaymentMethodConfiguration`
  * Add support for `PaymentMethodConfiguration` on `PaymentIntentParams`
  * Add support for `PaymentMethodConfigurationDetails` on `PaymentIntent`

## 74.25.0-beta.1 - 2023-06-29
* [#1683](https://github.com/stripe/stripe-go/pull/1683) Update generated code for beta
  * Add support for `Metadata` on `InvoiceSubscriptionDetails`
  * Add support for new values `ad_nrt`, `ar_cuit`, `bo_tin`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `pe_ruc`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, and `vn_tin` on enum `OrderTaxDetailsTaxIdsType`

## 74.24.0-beta.1 - 2023-06-22
* [#1677](https://github.com/stripe/stripe-go/pull/1677) Update generated code for beta
  * Add support for new resource `CustomerSession`
  * Add support for `New` method on resource `CustomerSession`
  * Change type of `TaxRegistrationCountryOptionsUsTypeParams` and `TaxRegistrationCountryOptionsUsType` from `literal('state_sales_tax')` to `enum('local_lease_tax'|'state_sales_tax')`

## 74.23.0-beta.2 - 2023-06-15
* [#1666](https://github.com/stripe/stripe-go/pull/1666) Consolidate Beta SDKs section in README
* [#1665](https://github.com/stripe/stripe-go/pull/1665) Update generated code for beta
  * Add support for `SubscriptionDetails` on `Invoice`
  * Add support for new values `aba` and `swift` on enum `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypes`
  * Add support for new value `us_bank_transfer` on enum `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType`
  * Add support for `SetPauseCollection` on `QuoteLine`, `QuoteLinesParams`, and `SubscriptionScheduleAmendAmendmentsParams`
  * Add support for new value `pause_collection_start` on enums `QuoteSubscriptionDataBillOnAcceptanceBillFromType` and `QuoteSubscriptionDataOverridesBillOnAcceptanceBillFromType`
  * Add support for `PauseCollection` on `SubscriptionSchedulePhasesParams` and `SubscriptionSchedulePhases`
  * Add support for `LocalAmusementTax` on `TaxRegistrationCountryOptionsUsParams` and `TaxRegistrationCountryOptionsUs`
  * Remove support for `Locations` on `TaxSettingsParams` and `TaxSettings`
* [#1669](https://github.com/stripe/stripe-go/pull/1669) Update generated code for beta
* [#1673](https://github.com/stripe/stripe-go/pull/1673) Update generated code for beta
* [#1675](https://github.com/stripe/stripe-go/pull/1675) Update generated code for beta
  * Add support for `PaymentDetails` on `ChargeCaptureParams`, `ChargeParams`, `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentParams`, and `PaymentIntent`
  * Add support for `StatementDetails` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCard`

## 74.23.0-beta.1 - 2023-06-08
* [#1669](https://github.com/stripe/stripe-go/pull/1669) Update generated code for beta
  * Updated stable APIs to the latest version

## 74.22.0-beta.1 - 2023-06-01
* [#1663](https://github.com/stripe/stripe-go/pull/1663) Handle developer message in preview error responses
* [#1648](https://github.com/stripe/stripe-go/pull/1648) Introduce stripe.RawRequest as a canonical way to request APIs without definitions
  * Please refer to the [Custom Requests README section](https://github.com/stripe/stripe-go/tree/beta#custom-requests) for usage instructions.
* [#1659](https://github.com/stripe/stripe-go/pull/1659) Update generated code for beta
* [#1665](https://github.com/stripe/stripe-go/pull/1665) Update generated code for beta
  * Add support for `SubscriptionDetails` on `Invoice`
  * Add support for new values `aba` and `swift` on enum `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypes`
  * Add support for new value `us_bank_transfer` on enum `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType`
  * Add support for `SetPauseCollection` on `QuoteLine`, `QuoteLinesParams`, and `SubscriptionScheduleAmendAmendmentsParams`
  * Add support for new value `pause_collection_start` on enums `QuoteSubscriptionDataBillOnAcceptanceBillFromType` and `QuoteSubscriptionDataOverridesBillOnAcceptanceBillFromType`
  * Add support for `PauseCollection` on `SubscriptionSchedulePhasesParams` and `SubscriptionSchedulePhases`
  * Add support for `LocalAmusementTax` on `TaxRegistrationCountryOptionsUsParams` and `TaxRegistrationCountryOptionsUs`
  * Remove support for `Locations` on `TaxSettingsParams` and `TaxSettings`

## 74.21.0-beta.1 - 2023-05-25
* [#1663](https://github.com/stripe/stripe-go/pull/1663) Handle developer message in preview error responses
* [#1648](https://github.com/stripe/stripe-go/pull/1648) Introduce stripe.RawRequest as a canonical way to request APIs without definitions
* [#1659](https://github.com/stripe/stripe-go/pull/1659) Update generated code for beta

## 74.20.0-beta.1 - 2023-05-19
* [#1658](https://github.com/stripe/stripe-go/pull/1658) Update generated code for beta
  * Add support for `Subscribe` and `Unsubscribe` methods on resource `FinancialConnections.Account`
  * Add support for `NextRefreshAvailableAt` on `FinancialConnectionsAccountBalanceRefresh`, `FinancialConnectionsAccountInferredBalancesRefresh`, `FinancialConnectionsAccountOwnershipRefresh`, and `FinancialConnectionsAccountTransactionRefresh`
  * Add support for `StatusDetails` and `Status` on `TaxSettings`

## 74.19.0-beta.1 - 2023-05-11
* [#1654](https://github.com/stripe/stripe-go/pull/1654) Update generated code for beta
  * Add support for `PayerEmail`, `PayerName`, and `SellerProtection` on `ChargePaymentMethodDetailsPaypal`
  * Add support for `CaptureMethod`, `PreferredLocale`, `ReferenceID`, and `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsPaypalParams`
  * Add support for `Reference` on `CheckoutSessionPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypal`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, `PaymentIntentPaymentMethodOptionsPaypalParams`, and `PaymentIntentPaymentMethodOptionsPaypal`
  * Add support for `RiskCorrelationID` on `CheckoutSessionPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, and `PaymentIntentPaymentMethodOptionsPaypalParams`
  * Remove support for `BillingAgreementID` and `Currency` on `CheckoutSessionPaymentMethodOptionsPaypalParams`
  * Add support for `Fingerprint`, `PayerID`, and `VerifiedEmail` on `MandatePaymentMethodDetailsPaypal` and `PaymentMethodPaypal`
  * Add support for `TaxabilityReason` and `TaxableAmount` on `OrderShippingCostTaxes`, `OrderTotalDetailsBreakdownTaxes`, and `QuotePhaseTotalDetailsBreakdownTaxes`
  * Add support for `HeadOffice` on `TaxSettingsParams` and `TaxSettings`

## 74.18.0-beta.1 - 2023-05-04
* [#1651](https://github.com/stripe/stripe-go/pull/1651) Update generated code for beta
  * Updated stable APIs to the latest version

## 74.17.0-beta.1 - 2023-04-27
* [#1645](https://github.com/stripe/stripe-go/pull/1645) Update generated code for beta
  * Add support for `BillingCycleAnchor` and `ProrationBehavior` on `CheckoutSessionSubscriptionDataParams`
  * Add support for `TerminalID` on `IssuingAuthorizationMerchantData` and `IssuingTransactionMerchantData`
  * Add support for `Metadata` on `PaymentIntentCaptureParams`
  * Add support for `Checks` on `SetupAttemptPaymentMethodDetailsCard`
  * Add support for `TaxBreakdown` on `TaxCalculationShippingCost` and `TaxTransactionShippingCost`
  * Change type of `TaxRegistrationActiveFromParams` and `TaxRegistrationExpiresAtParams` from `longInteger` to `longInteger | literal('now')`

## 74.16.0-beta.3 - 2023-04-20
* [#1642](https://github.com/stripe/stripe-go/pull/1642) Update generated code for beta
  * Add support for `Zip` on `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`
  * Add support for `CountryOptions` on `TaxRegistrationParams` and `TaxRegistration`
  * Remove support for `State` and `Type` on `TaxRegistrationParams` and `TaxRegistration`

## 74.16.0-beta.2 - 2023-04-13
* [#1639](https://github.com/stripe/stripe-go/pull/1639) Update generated code for beta
  * Add support for `CollectPaymentMethod` and `ConfirmPaymentIntent` methods on resource `Terminal.Reader`
  * Add support for `PaypalPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for new value `REVOIE23` on enums `ChargePaymentMethodDetailsIdealBic`, `PaymentMethodIdealBic`, and `SetupAttemptPaymentMethodDetailsIdealBic`
  * Add support for `CollectPaymentMethod` and `ConfirmPaymentIntent` on `TerminalReaderAction`
  * Add support for `StripeAccount` on `TerminalReaderActionProcessPaymentIntent` and `TerminalReaderActionRefundPayment`
  * Add support for new values `collect_payment_method` and `confirm_payment_intent` on enum `TerminalReaderActionType`

## 74.16.0-beta.1 - 2023-04-06
* [#1637](https://github.com/stripe/stripe-go/pull/1637) Update generated code for beta
  * Add support for `TreasuryTransaction` on `CapitalFinancingTransactionListParams`
  * Add support for `Transaction` on `CapitalFinancingTransactionDetails`
  * Add support for new value `link` on enum `PaymentMethodCardWalletType`

## 74.15.0-beta.1 - 2023-03-30
* [#1634](https://github.com/stripe/stripe-go/pull/1634) Update generated code
  * Add support for new value `ioss` on enum `TaxRegistrationType`

## 74.14.0-beta.1 - 2023-03-23
* [#1623](https://github.com/stripe/stripe-go/pull/1623) Update generated code for beta (new)
  * Add support for new resources `Tax.CalculationLineItem` and `Tax.TransactionLineItem`
  * Add support for `CollectInputs` method on resource `Terminal.Reader`
  * Add support for `FinancingOffer` on `CapitalFinancingSummary`
  * Add support for `FxRate` on `CheckoutSessionCurrencyConversion`
  * Add support for new value `link` on enum `PaymentLinkPaymentMethodTypes`
  * Add support for `AutomaticPaymentMethods` on `SetupIntentParams` and `SetupIntent`
  * Remove support for `Preview` on `TaxCalculationParams`
  * Add support for `TaxBreakdown` on `TaxCalculation`
  * Remove support for `TaxSummary` on `TaxCalculation`
  * Change type of `TaxCalculationLineItems` from `$LineItem` to `$Tax.CalculationLineItem`
  * Change type of `TaxTransactionLineItems` from `$LineItem` to `$Tax.TransactionLineItem`
  * Add support for `CollectInputs` on `TerminalReaderAction`
  * Add support for new value `collect_inputs` on enum `TerminalReaderActionType`

## 74.13.0-beta.1 - 2023-03-16
* [#1621](https://github.com/stripe/stripe-go/pull/1621) API Updates
  * Add support for `CreateFromCalculation` method on resource `Tax.Transaction`
  * Change type of `InvoiceAppliesTo` from `nullable(QuotesResourceQuoteLinesAppliesTo)` to `QuotesResourceQuoteLinesAppliesTo`
  * Add support for `Paypal` on `MandatePaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
  * Add support for `SetupFutureUsage` on `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypal`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, `PaymentIntentPaymentMethodOptionsPaypalParams`, and `PaymentIntentPaymentMethodOptionsPaypal`
  * Add support for new value `automatic_async` on enums `OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethod` and `OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethod`
  * Remove support for `AppliesTo` on `QuotePreviewInvoiceLinesParams`
  * Add support for `ShippingCost` on `TaxCalculationParams`, `TaxCalculation`, `TaxTransactionCreateReversalParams`, and `TaxTransaction`
  * Add support for `TaxBreakdown` on `TaxCalculation`
  * Remove support for `TaxSummary` on `TaxCalculation`

## 74.12.0-beta.1 - 2023-03-09
* [#1617](https://github.com/stripe/stripe-go/pull/1617) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Remove support for `ListTransactions` method on resource `Tax.Transaction`
  * Add support for `UpdateBehavior` on `SubscriptionPrebillingParams`, `SubscriptionPrebilling`, `SubscriptionSchedulePrebillingParams`, and `SubscriptionSchedulePrebilling`
  * Add support for `Prebilling` on `SubscriptionScheduleAmendParams`
  * Change type of `SubscriptionScheduleAppliesTo` from `nullable(QuotesResourceQuoteLinesAppliesTo)` to `QuotesResourceQuoteLinesAppliesTo`
  * Add support for `TaxabilityOverride` on `TaxCalculationCustomerDetailsParams`, `TaxCalculationCustomerDetails`, and `TaxTransactionCustomerDetails`
  * Add support for `TaxSummary` on `TaxCalculation`
  * Remove support for `TaxBreakdown` on `TaxCalculation`
  * Add support for `TaxBehavior` on `TaxSettingsDefaultsParams` and `TaxSettingsDefaults`

## 74.11.0-beta.1 - 2023-03-02
* [#1615](https://github.com/stripe/stripe-go/pull/1615) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for new resources `Issuing.CardBundle` and `Issuing.CardDesign`
  * Add support for `Get` and `List` methods on resource `CardBundle`
  * Add support for `Get`, `List`, and `Update` methods on resource `CardDesign`
  * Remove support for `Controller` on `AccountParams`
  * Add support for `CardDesign` on `IssuingCardParams` and `IssuingCard`

## 74.10.0-beta.1 - 2023-02-23
* [#1610](https://github.com/stripe/stripe-go/pull/1610) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for `ManualEntry` on `CheckoutSessionPaymentMethodOptionsUsBankAccountFinancialConnections`, `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountFinancialConnectionsParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnections`, `SetupIntentConfirmPaymentMethodOptionsUsBankAccountFinancialConnectionsParams`, `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsParams`, and `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnections`
  * Add support for new value `igst` on enum `TaxCalculationTaxBreakdownTaxRateDetailsTaxType`

## 74.9.0-beta.1 - 2023-02-16
* [#1608](https://github.com/stripe/stripe-go/pull/1608) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for `CurrencyConversion` on `CheckoutSession`
  * Add support for `Limits` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
  * Remove support for `Enabled` on `FinancialConnectionsSessionManualEntryParams`
  * Change type of `QuoteStatusDetailsCanceled` from `nullable(QuotesResourceStatusDetailsCanceledStatusDetails)` to `QuotesResourceStatusDetailsCanceledStatusDetails`
  * Change type of `QuoteStatusDetailsStale` from `nullable(QuotesResourceStatusDetailsStaleStatusDetails)` to `QuotesResourceStatusDetailsStaleStatusDetails`
  * Remove support for `Reference` on `TaxCalculationParams` and `TaxCalculation`
  * Add support for `Reference` on `TaxTransactionParams`

## 74.8.0-beta.1 - 2023-02-02
* [#1601](https://github.com/stripe/stripe-go/pull/1601) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for new resource `FinancialConnections.Transaction`
  * Add support for `List` method on resource `Transaction`
  * Add support for `Prefetch` on `-PaymentMethodOptionsUsBankAccountFinancialConnectionsParams` and `-PaymentMethodOptionsUsBankAccountFinancialConnections` across several APIs.
  * * Add support for `InferredBalancesRefresh`, `Subscriptions`, and `TransactionRefresh` on `FinancialConnectionsAccount`
  * Add support for `ManualEntry` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
  * Add support for `StatusDetails` and `Status` on `FinancialConnectionsSession`
  * Add support for new value `ownership` on enums `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPermissions` and `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPermissions`
  * Add support for `AccountNumber` on `PaymentMethodUsBankAccount`
  * Remove support for `ID` on `QuoteLinesStartsAtLineEndsAtParams`

## 74.7.0-beta.2 - 2023-01-26
* [#1598](https://github.com/stripe/stripe-go/pull/1598) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for `ListTransactions` method on resource `Tax.Transaction`
  * Add support for `BillingAgreementID` on `CheckoutSessionPaymentMethodOptionsPaypalParams`
  * Change type of `QuoteSubscriptionDataOverridesParams` from `array(create_specs)` to `emptyStringable(array(update_specs))`

## 74.7.0-beta.1 - 2023-01-19
* [#1596](https://github.com/stripe/stripe-go/pull/1596) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for `Tax.Settings` resource.

## 74.6.0-beta.2 - 2023-01-12
* [#1591](https://github.com/stripe/stripe-go/pull/1591) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for new resource `Tax.Registration`
  * Add support for `List`, `New`, and `Update` methods on resource `Registration`
  * Add support for `Controller` on `AccountParams`
  * Add support for `Application` and `Dashboard` on `AccountController`
  * Remove support for `Timestamp` on `QuoteLineActionsAddDiscountDiscountEnd`
  * Change type of `QuoteLineActionsAddDiscountDiscountEndType` from `literal('timestamp')` to `literal('line_ends_at')`
  * Remove support for `Index` on `QuoteLineActionsAddItemDiscounts`, `QuoteLineActionsRemoveDiscount`, `QuoteLineActionsSetDiscounts`, `QuoteLineActionsSetItemsDiscounts`, `SubscriptionSchedulePhasesAddInvoiceItemsDiscounts`, `SubscriptionSchedulePhasesDiscounts`, and `SubscriptionSchedulePhasesItemsDiscounts`

## 74.6.0-beta.1 - 2023-01-05
* [#1589](https://github.com/stripe/stripe-go/pull/1589) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for `MarkStaleQuote` method on resource `Quote`
  * Add support for `Duration` and `LineEndsAt` on `QuoteSubscriptionDataBillOnAcceptanceBillUntilParams` and `QuoteSubscriptionDataOverridesBillOnAcceptanceBillUntilParams`
  * Remove support for `LineStartsAt` on `QuoteSubscriptionDataBillOnAcceptanceBillUntilParams` and `QuoteSubscriptionDataOverridesBillOnAcceptanceBillUntilParams`
  * Add support for `Metadata` on `TerminalReaderActionRefundPayment` and `TerminalReaderRefundPaymentParams`

## 74.5.0-beta.1 - 2022-12-22
* [#1587](https://github.com/stripe/stripe-go/pull/1587) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Move `taxcalculation` package to `tax/calculation` package, and `taxtransaction` package to `tax/transaction` package.

## 74.4.0-beta.1 - 2022-12-15
* [#1585](https://github.com/stripe/stripe-go/pull/1585) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for new resources `QuoteLine`, `TaxCalculation`, and `TaxTransaction`
  * Add support for `ListLineItems` and `New` methods on resource `TaxCalculation`
  * Add support for `CreateReversal`, `Get`, and `New` methods on resource `TaxTransaction`

## 74.3.0-beta.1 - 2022-12-08
This release changes the pinned API version to `2022-11-15`.

* [#1582](https://github.com/stripe/stripe-go/pull/1582) API Updates for beta branch
  * Updated stable APIs to the latest version
* [#1580](https://github.com/stripe/stripe-go/pull/1580) API Updates for beta branch
  * Updated stable APIs to the latest version

## 73.17.0-beta.1 - 2022-11-10
* [#1572](https://github.com/stripe/stripe-go/pull/1572) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for `DiscountEnd` on `*DiscountParams`.
  * Add support for `URL` on `IssuingAuthorizationMerchantData`.

## 73.15.0-beta.2 - 2022-11-02
* [#1564](https://github.com/stripe/stripe-go/pull/1564) API Updates for beta branch
  * Updated stable APIs to the latest version

## 73.15.0-beta.1 - 2022-10-21
* [#1561](https://github.com/stripe/stripe-go/pull/1561) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for `Paypal` on `ChargePaymentMethodDetails` and `Source`
  * Add support for `NetworkData` on `IssuingTransaction`
  * Add support for new value `paypal` on enum `SourceType`
  * Add support for `BillingCycleAnchor` on `SubscriptionScheduleAmendAmendmentsParams`

## 73.14.0-beta.1 - 2022-10-14
* [#1559](https://github.com/stripe/stripe-go/pull/1559) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for `ScheduleSettings` on `SubscriptionScheduleAmendParams`

## 73.13.0-beta.1 - 2022-10-07
* [#1552](https://github.com/stripe/stripe-go/pull/1552) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add `ReferenceID` to `OrderPaymentSettingsPaymentMethodOptionsPaypalParams` and `OrderPaymentSettingsPaymentMethodOptionsPaypal`
  * Rename `CapitalFinancingSummaries` client to `CapitalFinancingSummary`.

## 73.11.0-beta.1 - 2022-09-26
* [#1548](https://github.com/stripe/stripe-go/pull/1548) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add `FinancingOffer`, `FinancingSummary` and `FinancingTransaction` resources.

## 73.5.0-beta.1 - 2022-08-26
* [#1535](https://github.com/stripe/stripe-go/pull/1535) API Updates for beta branch
  * Updated stable APIs to the latest version
  * Add support for the beta [Gift Card API](https://stripe.com/docs/gift-cards).

## 73.4.0-beta.1 - 2022-08-23
* [#1529](https://github.com/stripe/stripe-go/pull/1529) Allow setting APIVersion
* [#1531](https://github.com/stripe/stripe-go/pull/1531) API Updates for beta branch
  - Updated stable APIs to the latest version
  - `Stripe-Version` beta headers are not pinned by-default and need to be manually specified, please refer to [beta SDKs README section](https://github.com/stripe/stripe-go/blob/master/README.md#beta-sdks)

## 73.3.0-beta.1 - 2022-08-11
* [#1525](https://github.com/stripe/stripe-go/pull/1525) API Updates for beta branch
  - Updated stable APIs to the latest version
  - Add `RefundPayment` method to Terminal resource

## 73.1.0-beta.1 - 2022-08-03
* [#1514](https://github.com/stripe/stripe-go/pull/1514) API Updates for beta branch
  - Updated stable APIs to the latest version
  - Added the `Order` resource support

## 72.121.0-beta.1 - 2022-07-22
* [#1498](https://github.com/stripe/stripe-go/pull/1498) API Updates for beta branch
  - Updated stable APIs to the latest version
  - Add `Price.MigrateTo` property
  - Add `SubscriptionSchedule.Amend` method.
  - Add `Discount.SubscriptionItem` property.
  - Add `Quote.SubscriptionData.BillingBehavior`, `BillingCycleAnchor`, `EndBehavior`, `FromSchedule`, `FromSubscription`, `Prebilling`, `ProrationBehavior` properties.
  - Add `Phases` parameter to `Quote.Create`
  - Add `Subscription.Discounts`, `Prebilling` properties.
* [#1503](https://github.com/stripe/stripe-go/pull/1503) API Updates for beta branch
  - Updated stable APIs to the latest version
  - Add `QuotePhase` resource
* [#1506](https://github.com/stripe/stripe-go/pull/1506) API Updates for beta branch
  - Updated stable APIs to the latest version

## 72.119.0-beta.1 - 2022-07-07
* [#1493](https://github.com/stripe/stripe-go/pull/1493) API Updates for beta branch
  - Include `server_side_confirmation_beta=v1` beta
  - Add `secretKeyConfirmation` to `PaymentIntent`

## 72.115.0-beta.1 - 2022-06-15
* [#1476](https://github.com/stripe/stripe-go/pull/1476) API Updates for beta branch
  Add support for NetworkDetails properties on ReceivedCredits/ReceivedDebits resource
