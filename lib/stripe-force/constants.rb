# typed: strict
# frozen_string_literal: true

module StripeForce
  module Constants
    SALESFORCE_ACCOUNT_ID_HEADER = 'Salesforce-Account-Id'
    SALESFORCE_KEY_HEADER = 'Salesforce-Key'
    SALESFORCE_PACKAGE_NAMESPACE_HEADER = "Salesforce-Package-Namespace"

    SF_ORDER = 'Order'
    SF_PRODUCT = 'Product2'
    SF_ACCOUNT = 'Account'
    SF_OPPORTUNITY = 'Opportunity'
    SF_PRICEBOOK = 'Pricebook2'
    SF_PRICEBOOK_ENTRY = 'PricebookEntry'

    CPQ_QUOTE = 'SBQQ__Quote__c'

    CPQ_QUOTE_PRIMARY_CONTACT = 'SBQQ__PrimaryContact__c'
    CPQ_QUOTE_OPPORTUNITY = 'SBQQ__Opportunity2__c'
    CPQ_QUOTE_ORDERED = 'SBQQ__Ordered__c'
    CPQ_QUOTE_PRIMARY = 'SBQQ__Primary__c'
    CPQ_QUOTE_SUBSCRIPTION_START_DATE = 'SBQQ__StartDate__c'
    CPQ_QUOTE_SUBSCRIPTION_TERM = 'SBQQ__SubscriptionTerm__c'

    CPQ_PRODUCT_SUBSCRIPTION_TYPE = 'SBQQ__SubscriptionType__c'
    class CPQProductSubscriptionTypeOptions < T::Enum
      enums do
        RENEWABLE = new("Renewable")
      end
    end

    CPQ_PRODUCT_SUBSCRIPTION_TERM = 'SBQQ__SubscriptionTerm__c'
    CPQ_PRODUCT_CHARGE_TYPE = 'SBQQ__ChargeType__c'

    # https://help.salesforce.com/s/articleView?id=sf.cpq_subscription_fields.htm&type=5
    CPQ_PRODUCT_BILLING_TYPE = 'SBQQ__BillingType__c'
    class CPQProductBillingTypeOptions < T::Enum
      enums do
        # licensed
        ADVANCE = new("Advance")
        # metered
        ARREARS = new("Arrears")
      end
    end

    CPQ_QUOTE_LINE = 'SBQQ__QuoteLine__c'
    CPQ_QUOTE_LINE_PRODUCT = 'SBQQ__Product__c'
    CPQ_QUOTE_LINE_PRICEBOOK_ENTRY = 'SBQQ__PricebookEntryId__c'

    # custom fields added by our package
    GENERIC_STRIPE_ID = 'Stripe_ID__c'
    ORDER_INVOICE_PAYMENT_LINK = 'Stripe_Invoice_Payment_Link__c'
    ORDER_SUBSCRIPTION_PAYMENT_LINK = 'Stripe_Subscription_Payment_Link__c'

    # https://developer.salesforce.com/docs/atlas.en-us.packagingGuide.meta/packagingGuide/coa_order_type.htm
    class OrderTypeOptions < T::Enum
      enums do
        NEW = new("New")
      end
    end

    class OrderStatusOptions < T::Enum
      enums do
        ACTIVATED = new("Activated")
      end
    end
  end
end
