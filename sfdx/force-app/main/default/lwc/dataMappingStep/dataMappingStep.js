import { LightningElement } from 'lwc';

export default class DataMappingStep extends LightningElement {
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

    // DEMO DATA //
    stripeCustomerFields = [
        {
            name: 'Name',
            staticValue: false,
            sfValue: 'name'
        },
        {
            name: 'Description',
            staticValue: false,
            sfValue: ''
        },
        {
            name: 'Email',
            staticValue: false,
            sfValue: 'email'
        },
        {
            name: 'Phone',
            staticValue: false,
            sfValue: 'phone'
        },
        {
            name: 'Address Line 1',
            staticValue: false,
            sfValue: 'shippingAddress'
        },
        {
            name: 'Address Line 2',
            staticValue: false,
            sfValue: ''
        },
        {
            name: 'City',
            staticValue: false,
            sfValue: 'shippingCity'
        },
        {
            name: 'State',
            staticValue: false,
            sfValue: 'shippingState'
        },
        {
            name: 'Postal Code',
            staticValue: false,
            sfValue: 'shippingZip'
        },
        {
            name: 'Country',
            staticValue: false,
            sfValue: 'shippingCountry'
        },
        {
            name: 'Shipping Address Line 1',
            staticValue: false,
            sfValue: 'shippingAddress'
        },
        {
            name: 'Shipping Address Line 2',
            staticValue: false,
            sfValue: ''
        },
        {
            name: 'Shipping City',
            staticValue: false,
            sfValue: 'shippingCity'
        },
        {
            name: 'Shipping State',
            staticValue: false,
            sfValue: 'shippingState'
        },
        {
            name: 'Shipping Postal Code',
            staticValue: false,
            sfValue: 'shippingZip'
        },
        {
            name: 'Shipping Country',
            staticValue: false,
            sfValue: 'shippingCountry'
        },
    ];

    customFields = [
        {
            name: 'Source',
            staticValue: true,
            sfValue: 'Salesforce'
        }
    ];

}