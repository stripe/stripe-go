import { LightningElement } from 'lwc';

export default class SubscriptionFilterStep extends LightningElement {
    actionOptions = [
        {
            label: 'All Conditions Are Met',
            value: 'allConditions'
        },
        {
            label: 'Any Conditions Are Met',
            value: 'anyConditions'
        },
        {
            label: 'Custom Logic Is Met',
            value: 'custom'
        },
        {
            label: 'Always (No Cirteria)',
            value: 'always'
        }
    ];
    operatorOptions = [
        {
            label: 'Equals',
            value: 'equals'
        },
        {
            label: 'Does Not Equal',
            value: 'notEqual'
        },
        {
            label: 'Greater Than',
            value: 'greater'
        },
        {
            label: 'Less Than',
            value: 'less'
        }
    ];
    resourceOptions = [
        {
            label: 'First Name',
            value: 'firstName'
        },
        {
            label: 'Last Name',
            value: 'lastName'
        },
        {
            label: 'Date Created',
            value: 'dateCreated'
        },
        {
            label: 'Date Modified',
            value: 'dateModified'
        },
        {
            label: 'Record Owner',
            value: 'owner'
        },
        {
            label: 'Assigned To',
            value: 'assignedTo'
        },
        {
            label: 'Account',
            value: 'account'
        },
        {
            label: 'Address',
            value: 'address'
        },
        {
            label: 'City',
            value: 'city'
        },
        {
            label: 'Zip Code',
            value: 'zip'
        },
        {
            label: 'State',
            value: 'state'
        },
        {
            label: 'Country',
            value: 'country'
        }
    ];

    conditions = [
        {
            id: '1',
            selectedField: 'country',
            operator: 'notEqual',
            value: 'United States'
        },
        {
            id: '2',
            selectedField: 'owner',
            operator: 'equals',
            value: 'John Doe'
        },
        {
            id: '3',
            selectedField: '',
            operator: 'equals',
            value: ''
        },
    ];
}