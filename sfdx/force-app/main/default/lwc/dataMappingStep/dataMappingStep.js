import { LightningElement, api, track } from 'lwc';

export default class DataMappingStep extends LightningElement {
    activeObject = 'customer';
    sfObjectOptions = [
        {
            label: '-- Select Salesforce Object --',
            value: ''
        },
        {
            label: 'Account',
            value: 'account'
        },
        {
            label: 'Contact',
            value: 'contact'
        },
        {
            label: 'Lead',
            value: 'lead'
        },
        {
            label: 'Opportunity',
            value: 'opportunity'
        },
        {
            label: 'Order',
            value: 'order'
        },
        {
            label: 'Product',
            value: 'product'
        },
    ];
    sfRecordTypeOptions = [
        {
            label: '-- Select Record Type --',
            value: ''
        },
        {
            label: 'Default',
            value: 'default'
        },
        {
            label: 'Client',
            value: 'client'
        },
    ];
    sfFieldOptions = [
        {
            label: '-- Select Salesforce Field --',
            value: ''
        },
        {
            label: 'Name',
            value: 'name'
        },
        {
            label: 'First Name',
            value: 'firstName'
        },
        {
            label: 'Last Name',
            value: 'lastName'
        },
        {
            label: 'Email',
            value: 'email'
        },
        {
            label: 'Phone',
            value: 'phone'
        },
        {
            label: 'Shipping Address',
            value: 'shippingAddress'
        },
        {
            label: 'Shipping City',
            value: 'shippingCity'
        },
        {
            label: 'Shipping State',
            value: 'shippingState'
        },
        {
            label: 'Shipping Country',
            value: 'shippingCountry'
        },
        {
            label: 'Shipping Zip Code',
            value: 'shippingZip'
        },
    ];

    defaultCustomerSections = [
        'standard',
        'address'
    ];

    defaultProductSections = [
        'standard'
    ];

    // DEMO DATA //
    @track stripeCustomerMappings = [
        {
            label: 'Standard Mappings',
            name: 'standard',
            description: '',
            fields: [
                {
                    name: 'Name',
                    description: 'The customer’s full name or business name.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Email',
                    description: 'The customer’s email address.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: 'email'
                },
                {
                    name: 'Phone',
                    description: 'The customer’s phone number.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: 'phone'
                },
                {
                    name: 'Description',
                    description: 'An arbitrary string attached to the object. Often useful for displaying to users.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        }, 
        {
            label: 'Address Group',
            name: 'address',
            description: '',
            fields: [
                {
                    name: 'Line 1',
                    description: 'Address line 1 (e.g., street, PO Box, or company name).',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Line 2',
                    description: 'Address line 2 (e.g., apartment, suite, unit, or building).',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'City',
                    description: 'City, district, suburb, town, or village.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'State',
                    description: 'State, county, province, or region.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Postal Code',
                    description: 'ZIP or postal code.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Country',
                    description: 'Two-letter country code (ISO 3166-1 alpha-2).',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        }, 
        {
            label: 'Shipping Address Group',
            name: 'shipping',
            description: '',
            fields: [
                {
                    name: 'Line 1',
                    description: 'Address line 1 (e.g., street, PO Box, or company name).',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Line 2',
                    description: 'Address line 2 (e.g., apartment, suite, unit, or building).',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'City',
                    description: 'City, district, suburb, town, or village.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'State',
                    description: 'State, county, province, or region.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Postal Code',
                    description: 'ZIP or postal code.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Country',
                    description: 'Two-letter country code (ISO 3166-1 alpha-2).',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Name',
                    description: 'Customer name.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Phone',
                    description: 'Customer phone (including extension).',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                
            ]
        }, 
        {
            label: 'Additional Information',
            name: 'additional',
            description: '',
            fields: [
                {
                    name: 'Currency',
                    description: 'Three-letter ISO code for the currency the customer can be charged in for recurring billing purposes.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Tax Exempt',
                    description: 'Describes the customer’s tax exemption status. One of none, exempt, or reverse. When set to reverse, invoice and receipt PDFs include the text “Reverse charge”.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                }
            ]
        }
    ];

    @track stripeProductMappings = [
        {
            label: 'Standard Mappings',
            name: 'standard',
            description: '',
            fields: [
                {
                    name: 'Name',
                    description: 'The product’s name, meant to be displayable to the customer. Whenever this product is sold via a subscription, name will show up on associated invoice line item descriptions.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Active',
                    description: 'Whether the product is currently available for purchase.',
                    type: 'Boolean',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: 'email'
                },
                {
                    name: 'Description',
                    description: 'The product’s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        }, 
        {
            label: 'Additional Information',
            name: 'additional',
            description: '',
            fields: [
                {
                    name: 'Images',
                    description: 'A list of up to 8 URLs of images for this product, meant to be displayable to the customer.',
                    type: 'Array',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Shippable',
                    description: 'Whether this product is shipped (i.e., physical goods).',
                    type: 'Boolean',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Statement Descriptor',
                    description: 'Extra information about a product which will appear on your customer’s credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Tax Code',
                    description: 'A tax code ID.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Unit Label',
                    description: 'A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'URL',
                    description: 'A URL of a publicly-accessible webpage for this product.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        },
        {
            label: 'Package Dimensions',
            name: 'dimensions',
            description: '',
            fields: [
                {
                    name: 'Package Dimensions Height',
                    description: 'Height, in inches.',
                    type: 'Decimal',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Package Dimensions Length',
                    description: 'Length, in inches.',
                    type: 'Decimal',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Package Dimensions Weight',
                    description: 'Weight, in ounces.',
                    type: 'Decimal',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Package Dimensions Width',
                    description: 'Width, in inches.',
                    type: 'Decimal',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        }
    ];

    @track stripeSubscriptionMappings = [
        {
            label: 'Standard Mappings',
            name: 'standard',
            description: '',
            fields: [
                {
                    name: 'Current Period End',
                    description: 'End of the current period that the subscription has been invoiced for. At the end of this period, a new invoice will be created.',
                    type: 'Timestamp',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Current Period Start',
                    description: 'Start of the current period that the subscription has been invoiced for.',
                    type: 'Timestamp',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: 'email'
                },
                {
                    name: 'Default Payment Method',
                    description: 'ID of the default payment method for the subscription. It must belong to the customer associated with the subscription. This takes precedence over default_source. If neither are set, invoices will use the customer’s invoice_settings.default_payment_method or default_source.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Customer',
                    description: 'ID of the customer who owns the subscription.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        }, 
        {
            label: 'Settings',
            name: 'settings',
            description: '',
            fields: [
                {
                    name: 'Cancel at Period End',
                    description: 'If the subscription has been canceled with the at_period_end flag set to true, cancel_at_period_end on the subscription will be true. You can use this attribute to determine whether a subscription that has a status of active is scheduled to be canceled at the end of the current period.',
                    type: 'Boolean',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Application Fee Percent',
                    description: 'A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner’s Stripe account.',
                    type: 'Decimal',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Billing Cycle Anchor',
                    description: 'Determines the date of the first full invoice, and, for plans with month or year intervals, the day of the month for subsequent invoices.',
                    type: 'Timestamp',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Cancel At',
                    description: 'A date in the future at which the subscription will automatically get canceled',
                    type: 'Timestamp',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Collection Method',
                    description: 'Either charge_automatically, or send_invoice. When charging automatically, Stripe will attempt to pay this subscription at the end of the cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Trial End',
                    description: 'If the subscription has a trial, the end of that trial.',
                    type: 'Timestamp',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Trial Start',
                    description: 'If the subscription has a trial, the beginning of that trial.',
                    type: 'Timestamp',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        },
        {
            label: 'Child Objects',
            name: 'child-objects',
            description: '',
            fields: [
                {
                    name: 'Items',
                    description: 'List of subscription items, each with an attached price.',
                    type: 'List',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Billing Thresholds',
                    description: 'Define thresholds at which an invoice will be sent, and the subscription advanced to a new billing period',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        },
        {
            label: 'Subhashes',
            name: 'subhashes',
            description: '',
            fields: [
                {
                    name: 'Automatic Tax',
                    description: 'Automatic tax settings for this subscription.',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Default Payment Method',
                    description: 'ID of the default payment method for the subscription. It must belong to the customer associated with the subscription. This takes precedence over default_source. If neither are set, invoices will use the customer’s invoice_settings.default_payment_method or default_source.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Default Source',
                    description: 'ID of the default payment source for the subscription. It must belong to the customer associated with the subscription and be in a chargeable state. If default_payment_method is also set, default_payment_method will take precedence. If neither are set, invoices will use the customer’s invoice_settings.default_payment_method or default_source.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Discount',
                    description: 'Describes the current discount applied to this subscription, if there is one. When billing, a discount applied to a subscription overrides a discount applied on a customer-wide basis.',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Pause Collection',
                    description: 'If specified, payment collection for this subscription will be paused.',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Payment Settings',
                    description: 'Payment settings passed on to invoices created by the subscription.',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Pending Invoice Item Interval',
                    description: 'Specifies an interval for how often to bill for any pending invoice items. It is analogous to calling Create an invoice for the given subscription at the specified interval.',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Schedule',
                    description: 'The schedule attached to the subscription.',
                    type: 'String',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Transfer Data',
                    description: 'The account (if any) the subscription’s payments will be attributed to for tax reporting, and where funds from each payment will be transferred to for each of the subscription’s invoices.',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        },
    ];

    @track stripeSubscriptionItemMappings = [
        {
            label: 'Standard Mappings',
            name: 'standard',
            description: '',
            fields: [
                {
                    name: 'Quantity',
                    description: 'The quantity of the plan to which the customer should be subscribed.',
                    type: 'Integer',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        }, 
        {
            label: 'Parent Objects',
            name: 'parent-objects',
            description: '',
            fields: [
                {
                    name: 'Price',
                    description: 'The price the customer is subscribed to.',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Subscription',
                    description: 'The subscription this subscription_item belongs to.',
                    type: 'Boolean',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
                {
                    name: 'Billing Thresholds',
                    description: 'Define thresholds at which an invoice will be sent, and the related subscription advanced to a new billing period.',
                    type: 'Hash',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        },
        {
            label: 'Child Objects',
            name: 'child-objects',
            description: '',
            fields: [
                {
                    name: 'Tax Rates',
                    description: 'The tax rates which apply to this subscription_item. When set, the default_tax_rates on the subscription do not apply to this subscription_item.',
                    type: 'Array',
                    defaultValue: 'Object > Field > Object > Field',
                    hasOverride: false,
                    staticValue: false,
                    sfValue: ''
                },
            ]
        }
    ];

    @track customerMetadataFields = [];
    @track productMetadataFields = [];
    @track subscriptionMetadataFields = [];
    @track subscriptionItemMetadataFields = [];

    toggleStaticValue(event) {
        let targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        let targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetSectionIndex && targetFieldIndex ) {
            let targetSection = this.stripeCustomerMappings[parseInt(targetSectionIndex)];
            targetSection.fields[parseInt(targetFieldIndex)].staticValue = !(targetSection.fields[parseInt(targetFieldIndex)].staticValue);
            targetSection.fields[parseInt(targetFieldIndex)].hasOverride = true;
            targetSection.fields[parseInt(targetFieldIndex)].sfValue = '';
        }
    }

    resetToDefault(event) {
        let targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        let targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetSectionIndex && targetFieldIndex) {
            let targetSection = this.stripeCustomerMappings[targetSectionIndex];
            targetSection.fields[parseInt(targetFieldIndex)].staticValue = false;
            targetSection.fields[parseInt(targetFieldIndex)].hasOverride = false;
            targetSection.fields[parseInt(targetFieldIndex)].sfValue = '';
        }
    }

    updatePicklist(event) {
        let targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        let targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetSectionIndex && targetFieldIndex) {
            let targetSection = this.stripeCustomerMappings[targetSectionIndex];
            event.currentTarget.value = event.detail.value;
            targetSection.fields[parseInt(targetFieldIndex)].hasOverride = true;
            targetSection.fields[parseInt(targetFieldIndex)].sfValue = event.detail.value;
        }
    }

    addMetadataRow() {
        let emptyMapping = {
            name: '',
            staticValue: false,
            sfValue: ''
        };

        switch(this.activeObject) {
            case 'customer':
                this.customerMetadataFields.unshift(emptyMapping);
                break;
            case 'product':
                this.productMetadataFields.unshift(emptyMapping);
                break; 
            case 'subscription':
                this.subscriptionMetadataFields.unshift(emptyMapping);
                break; 
            case 'subscription-item':
                this.subscriptionItemMetadataFields.unshift(emptyMapping);
                break; 
            default: 
                break;
        }
    }

    changeActiveObject(event) {
        this.activeObject = event.detail.name;
    }

    get customerObjectActive() {
        return this.activeObject == 'customer';
    }

    get productObjectActive() {
        return this.activeObject == 'product';
    }

    get subscriptionObjectActive() {
        return this.activeObject == 'subscription';
    }

    get subscriptionItemObjectActive() {
        return this.activeObject == 'subscription-item';
    }

}