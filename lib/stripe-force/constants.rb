# typed: strict
# frozen_string_literal: true

module StripeForce
  module Constants
    SALESFORCE_ACCOUNT_ID_HEADER = 'Salesforce-Account-Id'
    SALESFORCE_KEY_HEADER = 'Salesforce-Key'
    SALESFORCE_PACKAGE_NAMESPACE_HEADER = "Salesforce-Package-Namespace"
    SALESFORCE_INSTANCE_TYPE_HEADER = 'Salesforce-Type'

    class SFInstanceTypes < T::Enum
      enums do
        SANDBOX = new("SANDBOX")
        SCRATCH_ORG = new("SCRATCH_ORG")
        TRIAL = new("TRIAL")
        PRODUCTION = new("PRODUCTION")
      end
    end

    SF_ORDER = 'Order'
    SF_ORDER_ITEM = 'OrderItem'
    SF_PRODUCT = 'Product2'
    SF_ACCOUNT = 'Account'
    SF_CONTACT = 'Contact'
    SF_OPPORTUNITY = 'Opportunity'
    SF_PRICEBOOK = 'Pricebook2'
    SF_PRICEBOOK_ENTRY = 'PricebookEntry'

    SF_ID = 'Id'

    CPQ_QUOTE = 'SBQQ__Quote__c'

    CPQ_QUOTE_PRIMARY_CONTACT = 'SBQQ__PrimaryContact__c'
    CPQ_QUOTE_PRICEBOOK = "SBQQ__PricebookId__c"
    CPQ_QUOTE_OPPORTUNITY = 'SBQQ__Opportunity2__c'
    CPQ_QUOTE_ORDERED = 'SBQQ__Ordered__c'
    CPQ_QUOTE_PRIMARY = 'SBQQ__Primary__c'
    CPQ_QUOTE_SUBSCRIPTION_START_DATE = 'SBQQ__StartDate__c'
    CPQ_QUOTE_SUBSCRIPTION_TERM = 'SBQQ__SubscriptionTerm__c'
    CPQ_QUOTE_SUBSCRIPTION_PRICING = 'SBQQ__SubscriptionPricing__c'

    CPQ_QUOTE_BILLING_FREQUENCY = 'SBQQ__BillingFrequency__c'
    class CPQBillingFrequencyOptions < T::Enum
      enums do
        MONTHLY = new("Monthly")
        QUARTERLY = new("Quarterly")
        SEMIANNUAL = new("Semiannual")
        ANNUAL = new("Annual")
        INVOICE_PLAN = new("Invoice Plan")
      end
    end

    CPQ_PRODUCT_SUBSCRIPTION_TYPE = 'SBQQ__SubscriptionType__c'
    class CPQProductSubscriptionTypeOptions < T::Enum
      enums do
        RENEWABLE = new("Renewable")
        ONE_TIME = new("")
      end
    end

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

    SYNC_RECORD = 'Sync_Record__c'
    class SyncRecordFields < T::Enum
      enums do
        COMPOUND_ID = new('Compound_ID__c')

        PRIMARY_RECORD_ID = new("Primary_Record_ID__c")
        PRIMARY_OBJECT_TYPE = new("Primary_Object_Type__c")

        SECONDARY_RECORD_ID = new("Secondary_Record_ID__c")
        SECONDARY_OBJECT_TYPE = new("Secondary_Object_Type__c")

        RESOLUTION_MESSAGE = new("Resolution_Message__c")
        RESOLUTION_STATUS = new("Resolution_Status__c")
      end
    end

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

    # non-cpq constants
    CONNECTOR_SETTING_SALESFORCE_NAMESPACE = "salesforce_namespace"
    CONNECTOR_SETTING_SALESFORCE_INSTANCE_TYPE = 'salesforce_instance_type'
    CONNECTOR_SETTING_CPQ_TERM_UNIT = 'cpq_term_unit'

    class SalesforceNamespaceOptions < T::Enum
      enums do
        QA = new("QaStripeConnect")
        PRODUCTION = new("stripeConnector")
        NONE = new("c")
      end
    end
  end
end
