---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1868
is_stripe_api_change: true
released_in_version: 78.9.0
---

* Add support for new value `verification_requires_additional_proof_of_registration` on enums `AccountFutureRequirementsErrorsCode`, `AccountRequirementsErrorsCode`, `BankAccountFutureRequirementsErrorsCode`, and `BankAccountRequirementsErrorsCode`
* Add support for `DefaultValue` on `CheckoutSessionCustomFieldsDropdownParams`, `CheckoutSessionCustomFieldsDropdown`, `CheckoutSessionCustomFieldsNumericParams`, `CheckoutSessionCustomFieldsNumeric`, `CheckoutSessionCustomFieldsTextParams`, and `CheckoutSessionCustomFieldsText`
* Add support for `GeneratedFrom` on `ConfirmationTokenPaymentMethodPreviewCard` and `PaymentMethodCard`
* Add support for new values `issuing_personalization_design.activated`, `issuing_personalization_design.deactivated`, `issuing_personalization_design.rejected`, and `issuing_personalization_design.updated` on enum `EventType`
