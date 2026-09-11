---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2262
is_stripe_api_change: true
released_in_version: 84.4.0-alpha.2
---

* Add support for new resource `V2CoreConnectionSession`
* Add support for `Get` and `New` methods on resource `V2CoreConnectionSession`
* Add support for `List` method on resources `V2PaymentsSettlementAllocationIntentSplit` and `V2PaymentsSettlementAllocationIntent`
* Add support for `AgenticCommerceSettings` on `AccountSessionComponentsParams`
* Add support for `TerminalHardwareOrders` and `TerminalHardwareShop` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `NetworkCostPassthroughReport` on `AccountSessionComponents`
* Add support for new values `ae_bank_account`, `ag_bank_account`, `bh_bank_account`, `gm_bank_account`, `hk_bank_account`, `kh_bank_account`, `lc_bank_account`, `mc_bank_account`, `mg_bank_account`, `my_bank_account`, `qa_bank_account`, `rw_bank_account`, `th_bank_account`, `tt_bank_account`, and `vn_bank_account` on enums `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type` and `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for `CadenceData` on `V2BillingIntentParams` and `V2BillingIntent`
* Add support for `CancellationDetails` on `V2BillingIntentActionDeactivateParams`, `V2BillingIntentActionDeactivate`, and `V2BillingPricingPlanSubscription`
* Add support for `ContactPhone` on `V2CoreAccountParams`, `V2CoreAccountTokenParams`, and `V2CoreAccount`
* Add support for `RegistrationDate` on `V2CoreAccountIdentityBusinessDetailsParams`, `V2CoreAccountIdentityBusinessDetails`, and `V2CoreAccountTokenIdentityBusinessDetailsParams`
* Add support for new value `gb_vat` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
* Add support for `Reference` on `V2MoneyManagementAdjustment`
* Add support for `AccruedFees` on `V2MoneyManagementFinancialAccount`
* Add support for `StartingBalance` on `V2MoneyManagementFinancialAccountPayments`
* Add support for new value `accrued_fees` on enum `V2MoneyManagementFinancialAccount.Type`
* Add support for `AccountHolderAddress` and `AccountHolderName` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
* Add support for `Fingerprint` on `V2MoneyManagementPayoutMethodCard`
* Add support for `CardSpend` on `V2MoneyManagementReceivedCredit` and `V2MoneyManagementReceivedDebit`
* Add support for new value `card_spend` on enum `V2MoneyManagementReceivedCredit.Type`
* Add support for new value `card_spend` on enum `V2MoneyManagementReceivedDebit.Type`
* Add support for new values `advance`, `anticipation_repayment`, `balance_transfer`, `charge_failure`, `charge`, `climate_order_purchase`, `climate_order_refund`, `connect_collection_transfer`, `connect_reserved_funds`, `contribution`, `dispute_reversal`, `financing_paydown_reversal`, `financing_paydown`, `inbound_transfer_reversal`, `issuing_dispute_fraud_liability_debit`, `issuing_dispute_provisional_credit_reversal`, `issuing_dispute_provisional_credit`, `issuing_dispute`, `minimum_balance_hold`, `network_cost`, `obligation`, `outbound_payment_reversal`, `outbound_transfer_reversal`, `partial_capture_reversal`, `payment_network_reserved_funds`, `platform_earning_refund`, `platform_earning`, `platform_fee`, `received_credit_reversal`, `received_debit_reversal`, `refund_failure`, `risk_reserved_funds`, `stripe_balance_payment_debit_reversal`, `stripe_balance_payment_debit`, `stripe_fee_tax`, `transfer_reversal`, and `unreconciled_customer_funds` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Add support for `ApplicationFeeRefund`, `ApplicationFee`, `Charge`, `Dispute`, `Payout`, `Refund`, `ReserveHold`, `ReserveRelease`, `Topup`, `TransferReversal`, and `Transfer` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
* Add support for new values `application_fee_refund`, `application_fee`, `charge`, `dispute`, `payout`, `refund`, `reserve_hold`, `reserve_release`, `topup`, `transfer_reversal`, and `transfer` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* Add support for error codes `blocked_payout_method` and `unsupported_payout_method` on `BlockedByStripeError`
* Add support for error code `invalid_payout_method_data` on `InvalidPayoutMethodError`
* Add support for error code `limit_payout_method` on `QuotaExceededError`
