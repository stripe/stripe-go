import { LightningElement, api, track } from 'lwc';
import getPicklistValuesForMapper from '@salesforce/apex/setupAssistant.getPicklistValuesForMapper';
import getFormattedStripeObjectFields from '@salesforce/apex/setupAssistant.getFormattedStripeObjectFields';
import getMappingConfigurations from '@salesforce/apex/setupAssistant.getMappingConfigurations';
import saveMappingConfigurations from '@salesforce/apex/setupAssistant.saveMappingConfigurations';
import { getErrorMessage } from 'c/utils';
import Alert from 'c/alert';

const TERMINATION_METADATA_PREFIX = 'sbc_termination.';

const blankMetadataMapping = () => {
    return {
        metadataMapping: {
            label: '',
            name: '',
            value: '',
            description: '',
            alerts: [],
            fields: []
        }
    };
}

const mappingListDefaults = () => {
    return {
        customer: {},
        product: {},
        subscription_schedule: {},
        subscription_phase: {},
        subscription_item: {},
        price: {},
        price_order_item: {},
        coupon: {}
    };
};

export default class DataMappingStep extends LightningElement {
    /*
     `@track` decorator is needed here to tell the lwc framework to observe 
      changes to the properties of an object and to the elements of an array
     */

    friendlyStripeObjectName = 'Customer';
    @track activeObject = 'customer';
    activeObjectDescription;
    activeObjectAlerts;
    staticValue;
    stripeObjectField;
    defaultSfObject;
    isConnected = false;
    isMappingsUpdated = false;
    configurationHash;
    contentLoading = new CustomEvent('contentloading');
    contentLoadingComplete = new CustomEvent('contentloadingcomplete');

    MENU = [
        { object: "customer", name: "customer", label: "Customer", hidden: false },
        { object: "subscription_schedule", name: "subscriptionSchedule", label: "Subscription Schedule", hidden: false },
        { object: "subscription_phase", name: "subscriptionPhase", label: "Subscription Phase", hidden: false },
        { object: "subscription_item", name: "subscriptionItem", label: "Subscription Item", hidden: false },
        { object: "product", name: "product", label: "Product", hidden: false },
        { object: "price", name: "price", label: "Price", hidden: false },
        { object: "price_order_item", name: "priceOrderItem", label: "Price (Order Item)", hidden: false },
        { object: "coupon", name: "coupon", label: "Coupon", hidden: false },
    ]

    ACTIVE_OBJECT_INFO = {
        "customer": {
            objectName: "Account",
            friendlyName: 'Account',
            description: 'Customer objects allow you to perform recurring charges, and to track multiple charges, that are associated with the same customer. The API allows you to create, delete, and update your customers. You can retrieve individual customers as well as a list of all your customers.'
        },
        "subscriptionSchedule": {
            objectName: "Order",
            friendlyName: 'Subscription Schedule',
            description: 'A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.'
        },
        "subscriptionPhase": {
            objectName: "Order",
            friendlyName: 'Subscription Phase',
            description: 'Subscription Phases allow you to create subscription schedules with multiple billing phases.'
        },
        "subscriptionItem": {
            objectName: "OrderItem",
            friendlyName: 'Subscription Schedule Phase Item',
            description: 'Subscription items allow you to create customer subscriptions with more than one plan, making it easy to represent complex billing relationships.'
        },
        "product": {
            objectName: "Product2",
            friendlyName: 'Product',
            description: 'Products describe the specific goods or services you offer to your customers. For example, you might offer a Standard and Premium version of your goods or service; each version would be a separate Product. They can be used in conjunction with Prices to configure pricing in Checkout and Subscriptions.'
        },
        "price": {
            objectName: "PricebookEntry",
            friendlyName: 'Price',
            description: 'Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.',
        },
        "priceOrderItem": {
            objectName: "OrderItem",
            friendlyName: 'Price (Order Item)',
            description: 'Price order items define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of a single product.',
            alerts: [
                {
                    type: Alert.Types.warning,
                    message: 'In some cases we create Stripe Prices out of Order Items, for example when the UnitPrice on an Order Item differs from the associated Pricebook Entry\'s UnitPrice. You can set the mappings for this case here.',
                },
            ],
        },
        "coupon": {
            objectName: "Order_Stripe_Coupon__c",
            friendlyName: 'Stripe Coupon on Order',
            description: 'Coupons contain information about a percent-off or amount-off discount you might want to apply to a subscription or subscription item.'
        },
    }
    @track hiddenMapperFields = [];
    @track sfFieldOptions = [];
    @track customerMappings;
    @track productMappings;
    @track subscriptionScheduleMappings;
    @track subscriptionPhaseMappings;
    @track subscriptionItemMappings;
    @track priceMappings;
    @track priceOrderItemMappings;
    @track couponMappings;
    @track fieldListByObjectMap;
    // allMappingConfigurations holds all mappings retrieved from ruby, value is set in `getMappingConfigurations`
    @track allMappingConfigurations;

    // allMappingList is used to retrieve and send all user mappings `saveMappingConfigurations` 
    @track allMappingList = {
        // these are static values defined by the user in the mapper UI
        field_defaults: mappingListDefaults(),
        // these are user defined mappings from Salesforce fields to Stripe fields
        field_mappings: mappingListDefaults(),
        // these are default mappings sent from Ruby (can be overridden in map)
        default_mappings: mappingListDefaults(),
        // these are required mappings sent from Ruby 
        required_mappings: mappingListDefaults()
    };
    @track customerMetadataFields = blankMetadataMapping();
    @track productMetadataFields = blankMetadataMapping();
    @track subscriptionScheduleMetadataFields = blankMetadataMapping();
    @track subscriptionPhaseMetadataFields = blankMetadataMapping();
    @track subscriptionItemMetadataFields = blankMetadataMapping();
    @track priceMetadataFields = blankMetadataMapping();
    @track priceOrderItemMetadataFields = blankMetadataMapping();
    @track couponMetadataFields = blankMetadataMapping();

    @track activeStripeObjectSections;

    get hasActiveAlerts() {
        return this.activeObjectAlerts !== undefined && this.activeObjectAlerts instanceof Array;
    }

    get priceOrderItemObjectActive() {
        return this.activeObject == 'priceOrderItem';
    }

    get activeObjectFields() {
        return this[this.activeObject + 'Mappings'];
    }

    set activeObjectFields(value) {
        let parsedVal = JSON.parse(JSON.stringify(value));
        this[this.activeObject + 'Mappings'] = parsedVal;
    }

    get activeMetadataObjectFields() {
        let activeObjectName = this.activeObject;
        return this[activeObjectName + 'MetadataFields'];
    }

    set activeMetadataObjectFields(value) {
        const parsedVal = JSON.parse(JSON.stringify(value));
        let activeObjectName = this.activeObject;
        this[activeObjectName + 'MetadataFields'] = parsedVal;
    }

    get menuItems() {
        return this.applyHiddenMapperFieldDataToTopLevel(this.MENU);
    }

    get listOfStripeMappingObjects() {
        return this.applyHiddenMapperFieldDataToTopLevel([
            {
                object: 'customer',
                mappingsObject: this.customerMappings,
                metadataMappingsObject: this.customerMetadataFields
            },
            {
                object: 'product',
                mappingsObject: this.productMappings,
                metadataMappingsObject: this.productMetadataFields
            },
            {
                object: 'subscription_schedule',
                mappingsObject: this.subscriptionScheduleMappings,
                metadataMappingsObject: this.subscriptionScheduleMetadataFields
            },
            {
                object: 'subscription_phase',
                mappingsObject: this.subscriptionPhaseMappings,
                metadataMappingsObject: this.subscriptionPhaseMetadataFields
            },
            {
                object: 'subscription_item',
                mappingsObject: this.subscriptionItemMappings,
                metadataMappingsObject: this.subscriptionItemMetadataFields
            },
            {
                object: 'price',
                mappingsObject: this.priceMappings,
                metadataMappingsObject: this.priceMetadataFields
            },
            {
                object: 'price_order_item',
                mappingsObject: this.priceOrderItemMappings,
                metadataMappingsObject: this.priceOrderItemMetadataFields
            },
            {
                object: 'coupon',
                mappingsObject: this.couponMappings, 
                metadataMappingsObject: this.couponMetadataFields
            }
        ]);
    } 

    valueChange() {
        if(!this.isMappingsUpdated) {
            this.isMappingsUpdated = true;
            this.dispatchEvent(new CustomEvent('valuechange'));
        }
        this.isMappingsUpdated = false;
    }

    debounce(targetInput) {
        setTimeout(() => {
            targetInput.dropdownLoading = false;
        }, 1);
    }

    saveObjectMappings(stripeObjectMappings, listOfAllMappings, listOfMetadataFields, stripeObjectName) {
        for(let i = 0; i < stripeObjectMappings.length; i++) {
            for(let j = 0; j < stripeObjectMappings[i].fields.length; j++) {
                
                const fieldData = stripeObjectMappings[i].fields[j];

                if(!fieldData.value || !fieldData.sfValue) {
                    continue
                }
                
                if(fieldData.staticValue === true) {
                    listOfAllMappings.field_defaults[stripeObjectName][fieldData.value] = fieldData.sfValue;
                } else {
                    listOfAllMappings.field_mappings[stripeObjectName][fieldData.value] = fieldData.sfValue;
                }
            }
        }
        listOfAllMappings = this.saveMetadataMappings(listOfMetadataFields.metadataMapping.fields, listOfAllMappings, stripeObjectName);

        return listOfAllMappings;
    }
    
    saveMetadataMappings(metadataFieldList, listOfAllMappings, stripeObjectName) {
        if(metadataFieldList.length) {
            for(let i = 0; i < metadataFieldList.length; i++) {
                const fieldData = metadataFieldList[i];
                // console.log({ fieldData: JSON.parse(JSON.stringify(fieldData)) });
                const keyPrefix = fieldData.onTermination ? TERMINATION_METADATA_PREFIX : '';
                const metadataKey = keyPrefix + 'metadata.' + fieldData.value;

                if(!fieldData.sfValue) {
                    continue
                }

                if(fieldData.staticValue === true) {
                    listOfAllMappings.field_defaults[stripeObjectName][metadataKey] = fieldData.sfValue;
                } else {
                    listOfAllMappings.field_mappings[stripeObjectName][metadataKey] = fieldData.sfValue;
                }
            }
        }
        return listOfAllMappings;
    }

    toggleMetaStaticValue(event) {
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;

        if (targetFieldIndex ) {
            const intFieldIndex = parseInt(targetFieldIndex);
            const toggledStaticValue = !this.activeMetadataObjectFields.metadataMapping.fields[intFieldIndex].staticValue;
            this.activeMetadataObjectFields.metadataMapping.fields[intFieldIndex].staticValue = toggledStaticValue

            if(toggledStaticValue === true) {
                this.activeMetadataObjectFields.metadataMapping.fields[intFieldIndex].sfValue = '';
            }

            this.valueChange();
        }
    }

    debounce(targetInput) {
        setTimeout(() => {
            targetInput.dropdownLoading = false;
        }, 1);
    }

    updateFieldList(event) {
        this.sfFieldOptions = [];
        const  selectedObject = event.detail.object;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        const targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        this.stripeObjectField = this.activeObjectFields[targetSectionIndex].fields[targetFieldIndex];
        let targetInput = event.currentTarget;
        if(this.fieldListByObjectMap[selectedObject]) {
            this.sfFieldOptions = this.fieldListByObjectMap[selectedObject];
            this.filterSfOptionsByStripeFieldType();
            // Push updating the child loading attribute to the next render cycle
            this.debounce(targetInput);
        } else {
            this.getPicklistValuesForMapper(false, selectedObject, false, targetInput);
        }
    }

    updateMetadataFieldList(event) {
        this.sfFieldOptions = [];
        let selectedObject = event.detail.object;
        let targetInput = event.currentTarget;
        if(this.fieldListByObjectMap[selectedObject]) {
            this.sfFieldOptions = this.fieldListByObjectMap[selectedObject];
            // Push updating the child loading attribute to the next render cycle
            this.debounce(targetInput);
        } else {
            this.getPicklistValuesForMapper(false, selectedObject, true, targetInput);
        }
    }

    toggleStaticValue(event) {
        const targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetSectionIndex && targetFieldIndex ) {
            const updatedSelection = {
                hasSfValue: true,
                hasOverride: true,
                staticValue: true,
                sfValue: '',
                sfValueType: 'static'
            };
            updatedSelection.hasOverride = this.activeObjectFields[parseInt(targetSectionIndex)].fields[parseInt(targetFieldIndex)].defaultValue ? true : false;
            updatedSelection.staticValue = !this.activeObjectFields[parseInt(targetSectionIndex)].fields[parseInt(targetFieldIndex)].staticValue;
            Object.assign(this.activeObjectFields[targetSectionIndex].fields[parseInt(targetFieldIndex)] , updatedSelection);
            this.valueChange();
        }
    }

    removeMetaRow(event) {
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetFieldIndex ) {
            const value = this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)];
            this.activeMetadataObjectFields.metadataMapping.fields.splice(this.activeMetadataObjectFields.metadataMapping.fields.findIndex(metadataField => metadataField.name === value.name),1);
        }
        this.valueChange();
    }

    updateStripeMetadataName(event) {
        this.staticValue = event.target.value;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetFieldIndex ) {
            const updatedSelection = {
                name: this.staticValue,
                value: this.staticValue
            };
            Object.assign(this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)] , updatedSelection);
            this.valueChange();
        }
    }

    updateMetaStaticValueChoice(event) {
        this.staticValue = event.target.value;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetFieldIndex) {
            const updatedSelection = {
                staticValue: true,
                sfValue: this.staticValue,
                sfValueType: 'staticMetadata'
            };
            Object.assign(this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)] , updatedSelection);
            this.valueChange();
        }
    }

    toggleMetaOnTermination(event) {
        const checked = event.target.value;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetFieldIndex) {
            const currentVal = this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)].onTermination;
            // console.log({ checked, currentVal, nextVal: !currentVal })
            const updatedSelection = {
                onTermination: !currentVal,
            };
            Object.assign(this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)] , updatedSelection);
            this.valueChange();
        }
    }

    updateStaticValueChoice(event) {
        const targetSectionIndex = event.detail.sectionIndex;
        const targetFieldIndex = event.detail.fieldIndex;
        if (this.activeObjectFields[targetSectionIndex].fields[targetFieldIndex]) {
            const updatedSelection = {
                hasSfValue: true,
                staticValue: true,
                sfValue: event.detail.value,
                sfValueType: 'static'
            };
            Object.assign(this.activeObjectFields[targetSectionIndex].fields[targetFieldIndex] , updatedSelection)
            this.valueChange();
        }
    }

    resetToDefault(event) {
        const targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetSectionIndex && targetFieldIndex) {
            const updatedSelection = {
                hasSfValue: false,
                staticValue: false,
                hasOverride: false,
                sfValue: '',
                sfValueType: 'default'
            };
            Object.assign(this.activeObjectFields[targetSectionIndex].fields[parseInt(targetFieldIndex)], updatedSelection);
            this.valueChange();
        }
    }

    toggleOverride(event) {
        const targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetSectionIndex && targetFieldIndex) {
            const updatedSelection = {
                hasSfValue: true,
                hasOverride: false,
                hasRequiredValue: false,
                staticValue: false
            };
            updatedSelection.hasOverride = this.activeObjectFields[targetSectionIndex].fields[parseInt(targetFieldIndex)].defaultValue ? true : false;
            updatedSelection.hasRequiredValue = this.activeObjectFields[targetSectionIndex].fields[parseInt(targetFieldIndex)].requiredValue ? true : false;
            Object.assign(this.activeObjectFields[targetSectionIndex].fields[parseInt(targetFieldIndex)], updatedSelection);

            this.openPicklist(event);
        }
    }

    openPicklist(event) {
        const targetContainer = event.currentTarget.closest('td'); // Get the parent container of the targeted input 
        const staticInput = targetContainer.querySelector('lightning-input');
        if (staticInput) {
            staticInput.focus();
        } 
        this.updatePicklistChoices(event);
        
        //setting timeout to force next render cycle and allow picker to to be rendered
        setTimeout(() => {
            const fieldPicker = targetContainer.querySelector('c-field-picker');
            if (fieldPicker) {
                fieldPicker.focus();
            }
        }, 100);
    }

    updatePicklistChoices(event) {
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        const targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        this.stripeObjectField = this.activeObjectFields[targetSectionIndex].fields[targetFieldIndex];
        this.filterSfOptionsByStripeFieldType()
    }

    filterSfOptionsByStripeFieldType() {
       if (!this.stripeObjectField.type) {
            return;
        } 

        const stripeFieldType = this.stripeObjectField.type.toLowerCase()

        // has to be a copy to force a rerender
        const modifiedFieldOptions = JSON.parse(JSON.stringify(this.sfFieldOptions))

        // TODO this whole switch statement is a mess and is confusing, need to clean this up
        switch(stripeFieldType) {
            case 'integer' || 'decimal' || 'number':
                // NOTE `string` was specifically added as an acceptable value because of the payment terms mapping
                this.sfFieldOptions = modifiedFieldOptions.filter(fieldOptions => fieldOptions.type === 'double' || fieldOptions.type === 'reference' || fieldOptions.type == 'picklist');
                return;
            case 'timestamp':
                this.sfFieldOptions = modifiedFieldOptions.filter(fieldOptions => fieldOptions.type.includes('date') || fieldOptions.type === 'reference' )
                return;
            case 'boolean':
                this.sfFieldOptions = modifiedFieldOptions.filter(fieldOptions => fieldOptions.type === 'boolean' || fieldOptions.type === 'reference' )
                return;
            default:
                this.sfFieldOptions = modifiedFieldOptions
                return;
        }
    }

    updateMetaPicklist(event) {
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetFieldIndex ) {
            const updatedSelection = {
                hasSfValue: true,
                sfValue: event.detail.value,
                sfValueType: this.sfFieldOptions[event.detail.index].type
            };
            Object.assign(this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)] , updatedSelection);
            this.valueChange();
        }
    }

    updatePicklist(event) {
        const targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if(targetSectionIndex && targetFieldIndex) {
            const updatedSelection = {
                hasSfValue: true,
                sfValue: event.detail.value,
                sfValueType: this.sfFieldOptions[event.detail.index].type
            };
            Object.assign(this.activeObjectFields[targetSectionIndex].fields[parseInt(targetFieldIndex)] , updatedSelection);
            this.valueChange();
        }
    }

    addMetadataRow() {
        const metadataFieldObj = {
            name: '',
            value: '',
            type: 'metadata',
            defaultValue: '',
            hasOverride: false,
            staticValue: false,
            onTermination: false,
            sfValue: '',
            sfValueType: 'metadata'
        };
        this.activeMetadataObjectFields.metadataMapping.fields.push(metadataFieldObj);
    }

    mapperFieldSorter(a, b) {
        return a.label.localeCompare(b.label);
    }

    // triggered on each sidebar click
    changeActiveObject(event) {
        this.activeObject = event.detail.name;

        if(!this.fieldListByObjectMap || !this.activeObject) {
           return;
        }
    
        this.dispatchEvent(this.contentLoading);
        this.activeObjectDescription = this.ACTIVE_OBJECT_INFO[this.activeObject]['description'];
        this.activeObjectAlerts = this.ACTIVE_OBJECT_INFO[this.activeObject]['alerts'];
        this.updateObjectFieldsAndSetMappings(this.ACTIVE_OBJECT_INFO[this.activeObject]['objectName'], this.ACTIVE_OBJECT_INFO[this.activeObject]['friendlyName']);
        this.openActiveSection();
        this.dispatchEvent(this.contentLoadingComplete);
        return;
    }

    updateObjectFieldsAndSetMappings(objectName, friendlyName) {
        const METADATAFIELDS = 'MetadataFields';
        this.friendlyStripeObjectName = friendlyName;
        this.defaultSfObject = objectName;
        this.fieldListByObjectMap[objectName].sort(this.mapperFieldSorter);
        this.sfFieldOptions = this.fieldListByObjectMap[objectName];
    }

    // Open 'Metadata' mapping section if metadata mappings exist
    openActiveSection() {
        this.activeStripeObjectSections = ['standard'];

        if (this[this.activeObject + 'MetadataFields'].metadataMapping.fields.length > 0) {
            this.activeStripeObjectSections = ['standard', 'metadata']
        }
        
        return this[this.activeObject + 'Sections'] = this.activeStripeObjectSections
    }

    openStripeApi(event) {
        const stripeObject = event.currentTarget.dataset.object;
        let apiUrl = "https://stripe.com/docs/api/";

        // subscription_items are unique and really represent subscription schedule phase items
        if(stripeObject == 'subscriptionItem') {
            apiUrl += "subscription_schedules/create#create_subscription_schedule-phases-items"
        } else if(stripeObject == 'subscriptionSchedule' || stripeObject == 'subscriptionPhase') {
            apiUrl += "subscription_schedules/create"
        }else if(stripeObject == 'priceOrderItem') {
            apiUrl += "prices/create"
        } else {
            // our API calls are always #create, so let's link to that specific area in the docs
            apiUrl += `${stripeObject}s/create`;
        }

        window.open(apiUrl, '_blank');
    }

    @api async connectedCallback() {
        this.dispatchEvent(this.contentLoading);
        try {
            const getFormattedStripeObjects = await getFormattedStripeObjectFields();
            const responseData = JSON.parse(getFormattedStripeObjects);
            if(responseData.error) {
                this.showToast(responseData.error, 'error', 'sticky');
                return;
            }

            this.customerMappings = responseData.results.formattedStripeCustomerFields;
            this.productMappings = responseData.results.formattedStripeProductItemFields;
            this.subscriptionScheduleMappings = responseData.results.formattedStripeSubscriptionFields;
            this.subscriptionPhaseMappings = responseData.results.formattedStripeSubscriptionSchedulePhaseFields;
            this.subscriptionItemMappings = responseData.results.formattedStripeSubscriptionItemFields;
            this.priceMappings = responseData.results.formattedStripePriceFields;
            this.priceOrderItemMappings = responseData.results.formattedStripePriceOrderItemFields;
            this.couponMappings = responseData.results.formattedStripeCouponFields;

            await this.getPicklistValuesForMapper(true, '', false);
        } catch (error) {
            const errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error');
        } finally {
            this.activeObject = 'customer';
            this.activeObjectDescription = this.ACTIVE_OBJECT_INFO[this.activeObject]['description'];
            this.activeObjectAlerts = this.ACTIVE_OBJECT_INFO[this.activeObject]['alerts'];
            // console.log('activeMappings', JSON.parse(JSON.stringify(this.activeStripeObjectMappings)));
        }
    }

    @api async getPicklistValuesForMapper(isConnectedCallback, ObjectName, isMetadataRow, targetElement) {
        this.loading = true;
        try {
            const getPicklistValues = await getPicklistValuesForMapper({
                isConnectedCallback: isConnectedCallback,
                ObjectApiName: ObjectName
            });

            const picklistValueResponseData =  JSON.parse(getPicklistValues);
            if(picklistValueResponseData.error) {
                this.showToast(picklistValueResponseData.error, 'error', 'sticky');
                return;
            }

            if(!picklistValueResponseData.results.isConnected) {
                return;
            }

            this.isConnected = true;
            if(isConnectedCallback === true) {
                this.fieldListByObjectMap = picklistValueResponseData.results.fieldListByObjectMap;
                this.sfFieldOptions = this.fieldListByObjectMap.Account
                this.defaultSfObject = 'Account';
                this.sfFieldOptions.sort(function(a, b) {
                    return a.label.localeCompare(b.label);
                });
            } else {
                this.fieldListByObjectMap[picklistValueResponseData.results.ObjectApiName] = picklistValueResponseData.results.listOfObjectFields;

                this.sfFieldOptions = this.fieldListByObjectMap[picklistValueResponseData.results.ObjectApiName]
                this.sfFieldOptions.sort(function(a, b) {
                    return a.label.localeCompare(b.label);
                });
                if(!isMetadataRow) {
                    this.filterSfOptionsByStripeFieldType();
                }
            }

            const getMappingConfigs = await getMappingConfigurations();
            const mappingConfigurationResponseData =  JSON.parse(getMappingConfigs);
            if(mappingConfigurationResponseData.error) {
                this.showToast(mappingConfigurationResponseData.error, 'error', 'sticky');
                return;
            }
            if(!mappingConfigurationResponseData.results.isConnected) {
                return;
            }

            this.configurationHash = mappingConfigurationResponseData.results.configurationHash;
            this.hiddenMapperFields = mappingConfigurationResponseData.results.hiddenMapperFields;
            // console.log('hiddenMapperFields', JSON.parse(JSON.stringify(this.hiddenMapperFields)));
            this.allMappingConfigurations = mappingConfigurationResponseData.results.allMappingConfigurations;
            // console.log('listOfStripeMappingObjects', JSON.parse(JSON.stringify(this.listOfStripeMappingObjects)));
            for (const mappingContainer of this.listOfStripeMappingObjects) {
                this.setFieldMappings(mappingContainer.object, mappingContainer.mappingsObject, mappingContainer.metadataMappingsObject.metadataMapping.fields);
            }
            
            this.allMappingList['default_mappings'] = this.allMappingConfigurations['default_mappings'];
            this.allMappingList['required_mappings'] = this.allMappingConfigurations['required_mappings'];

        } catch (error) {
            const errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error', 'sticky');
        } finally {
            this.openActiveSection();
            this.dispatchEvent(this.contentLoadingComplete);
            if(targetElement) targetElement.dropdownLoading = false;
        }
    }

    get activeObjectSnakeCase() {
        return this.activeObject.replace(/[A-Z]/g, letter => `_${letter.toLowerCase()}`);
    }

    get activeStripeObjectMetadataFields() {
        return this[this.activeObject + 'MetadataFields'];
    }

    get activeStripeObjectMappings() {
        return this.applyHiddenMapperFieldDataToCategories(this.activeObjectSnakeCase, this[this.activeObject + 'Mappings']);
    }

    applyHiddenMapperFieldDataToCategories(objectName, mappingObjs) {
        const segments = this.hiddenMapperFieldSegments;
        // console.log('segments.paths', segments.paths);
        if (segments.paths.length === 0) {
            return mappingObjs;
        }

        const paths = this.getHiddenMapperFieldPathsForObject(objectName, segments.paths);
        // console.log('paths', paths);

        if (paths.length === 0) {
            return mappingObjs;
        }

        for (let i = 0; i < mappingObjs.length; i++) {
            mappingObjs[i].hidden = paths.indexOf(mappingObjs[i].name) !== -1;
            if (mappingObjs[i].hidden) {
                continue;
            }

            this.applyHiddenMapperFieldDataToFieldLevel(mappingObjs[i], paths);
        }

        return mappingObjs;
    }

    applyHiddenMapperFieldDataToFieldLevel(section, paths) {
        for (let i = 0; i < section.fields.length; i++) {
            const field = section.fields[i];
            // handles standard object.section.field format
            const check1 = section.name + '.' + field.value;
            // console.log('check1', check1);
            field.hidden = paths.indexOf(check1) !== -1
            const fieldSegment = field.value.indexOf('.');
            if (field.hidden === false && fieldSegment !== -1) {
                const sectionEmbed = field.value.indexOf(section.name + '.');
                if (sectionEmbed !== -1) {
                    // handles subscription_schedule billing_thresholds default_settings.billing_thresholds.amount_gte
                    //   referenced as subscription_schedule.billing_thresholds.amount_gte
                    // handles subscription_schedule prebilling prebilling.iterations
                    //   referenced as subscription_schedule.prebilling.iterations
                    const check2 = section.name + field.value.substring(sectionEmbed + section.name.length);
                    // console.log('check2', check2);
                    field.hidden = paths.indexOf(check2) !== -1;
                }
            }
        }
    }

    applyHiddenMapperFieldDataToTopLevel(mappingObjs) {
        const segments = this.hiddenMapperFieldSegments;

        if (Object.keys(segments.objects).length === 0) {
            return mappingObjs;
        }

        for (let i = 0; i < mappingObjs.length; i++) {
            const mappingObj = mappingObjs[i];
            mappingObj.hidden = segments.objects[mappingObj.object];
        }
        return mappingObjs;
    }

    /**
     *
     * @param {string} objectName
     * @param {Array<string>} paths
     * @return {Array<string>}
     */
    getHiddenMapperFieldPathsForObject(objectName, paths) {
        const found = [];
        const match = objectName + '.';
        for (let i = 0; i < paths.length; i++) {
            // console.log('check', match, paths[0]);
            if (paths[i].indexOf(match) === 0) {
                found.push(paths[i].substring(match.length));
            }
        }
        return found;
    }

    /**
     *
     * @returns {{objects: Object<string, boolean>, paths: Array<string>}}
     */
    get hiddenMapperFieldSegments() {
        const ret = { objects: {}, paths: [] };
        for (let i = 0; i < this.hiddenMapperFields.length; i++) {
            if (this.hiddenMapperFields[i].indexOf('.') === -1) {
                ret.objects[this.hiddenMapperFields[i]] = true;
            } else {
                ret.paths.push(this.hiddenMapperFields[i]);
            }
        }

        return ret;
    }

    setFieldMappings(stripeObject, stripeObjectMap, metadataFieldList) {
        // TODO way too many nested for loops, need to simplify this
        for(let i = 0; i < stripeObjectMap.length; i++) {
            for(let j = 0; j < stripeObjectMap[i].fields.length; j++) {
                if(this.allMappingConfigurations.default_mappings &&
                   this.allMappingConfigurations.default_mappings[stripeObject] &&
                   this.allMappingConfigurations.default_mappings[stripeObject].length !== 0) {

                    for(const value in this.allMappingConfigurations.default_mappings[stripeObject]) {
                        const fieldData = stripeObjectMap[i].fields[j];
                        if(fieldData.value === value && this.allMappingConfigurations.default_mappings[stripeObject][value]) {
                            fieldData.sfValue = '';
                            fieldData.hasOverride = false;
                            fieldData.hasSfValue = false;
                            fieldData.staticValue = false;
                            fieldData.defaultValue = this.allMappingConfigurations.default_mappings[stripeObject][value];
                        }
                    }
                }
                if(this.allMappingConfigurations.field_defaults && this.allMappingConfigurations.field_defaults[stripeObject]) {
                    for(const value in this.allMappingConfigurations.field_defaults[stripeObject]) {
                        const fieldData = stripeObjectMap[i].fields[j];
                        if(fieldData.value === value) {
                            fieldData.staticValue = true
                            fieldData.sfValue = this.allMappingConfigurations.field_defaults[stripeObject][value];
                            if(fieldData.sfValue) {
                                fieldData.hasSfValue = true
                            }
                            if(fieldData.defaultValue) {
                                fieldData.hasOverride = true
                            }
                        } else if (value.startsWith('metadata.')) {
                            let realfieldMapValue = value.replace('metadata.','')
                            
                            let metadataFieldObj = this.getFieldObject(realfieldMapValue, realfieldMapValue, 'metadata', '', false, true, this.allMappingConfigurations.field_defaults[stripeObject][value], 'metadata');

                            if (metadataFieldList.length && metadataFieldList.filter(field =>  field.name === realfieldMapValue ).length > 0) {
                                //do nothing
                            } else {
                                metadataFieldList.push(metadataFieldObj);
                            }
                        } else if (value.startsWith(TERMINATION_METADATA_PREFIX + 'metadata.')) {
                            let realfieldMapValue = value.replace(TERMINATION_METADATA_PREFIX + 'metadata.','')

                            let metadataFieldObj = this.getFieldObject(realfieldMapValue, realfieldMapValue, 'metadata', '', false, true, this.allMappingConfigurations.field_defaults[stripeObject][value], 'metadata');
                            metadataFieldObj.onTermination = true;

                            if (metadataFieldList.length && metadataFieldList.filter(field =>  field.name === realfieldMapValue ).length > 0) {
                                //do nothing
                            } else {
                                metadataFieldList.push(metadataFieldObj);
                            }
                        }
                    }
                }
                if(this.allMappingConfigurations.field_mappings && this.allMappingConfigurations.field_mappings[stripeObject]) {
                    for(const value in this.allMappingConfigurations.field_mappings[stripeObject]) {
                        const fieldData = stripeObjectMap[i].fields[j];
                        if(fieldData.value === value && !value.startsWith('metadata.')) {
                            fieldData.staticValue = false
                            fieldData.sfValue = this.allMappingConfigurations.field_mappings[stripeObject][value];
                            if(fieldData.sfValue)fieldData.hasSfValue = true
                            if(fieldData.defaultValue)fieldData.hasOverride = true
                        } else if (value.startsWith('metadata.') ) {
                            let realValue = value.replace('metadata.','')
                            let metadataFieldObj = this.getFieldObject(realValue, realValue, 'metadata', '', false, false, this.allMappingConfigurations.field_mappings[stripeObject][value], 'metadata');
                            if (metadataFieldList.length && metadataFieldList.filter(field => field.name === realValue).length > 0) {
                                //do nothing
                            } else {
                                metadataFieldList.push(metadataFieldObj);
                            }
                        } else if (value.startsWith(TERMINATION_METADATA_PREFIX + 'metadata.') ) {
                            let realValue = value.replace(TERMINATION_METADATA_PREFIX + 'metadata.','')
                            let metadataFieldObj = this.getFieldObject(realValue, realValue, 'metadata', '', false, false, this.allMappingConfigurations.field_mappings[stripeObject][value], 'metadata');
                            metadataFieldObj.onTermination = true;
                            if (metadataFieldList.length && metadataFieldList.filter(field => field.name === realValue).length > 0) {
                                //do nothing
                            } else {
                                metadataFieldList.push(metadataFieldObj);
                            }
                        }
                    }
                }

                if(this.allMappingConfigurations.required_mappings && this.allMappingConfigurations.required_mappings[stripeObject]) {
                    for(const value in this.allMappingConfigurations.required_mappings[stripeObject]) {
                        const fieldData = stripeObjectMap[i].fields[j];
                        if(fieldData.value === value && !value.startsWith('metadata.')) {
                            fieldData.requiredValue = this.allMappingConfigurations.required_mappings[stripeObject][value];
                            fieldData.hasRequiredValue = true;
                        }
                    }
                }
            }
        }
    }

    getFieldObject(name, value, type, defaultValue, hasOverride, staticValue, sfValue, sfValueType) {
        let fieldObject = {
            name: name,
            value: value,
            type: type,
            defaultValue: defaultValue,
            hasOverride: hasOverride,
            staticValue: staticValue,
            onTermination: false,
            sfValue: sfValue,
            sfValueType: sfValueType
        };
        return fieldObject;
    }

    @api async saveDataMappings() {
        this.loading = true;
        let saveSuccess = false;

        let mappingData = JSON.parse(JSON.stringify(this.allMappingList));

        for (const mappingContainer of this.listOfStripeMappingObjects) {
            mappingData = this.saveObjectMappings(mappingContainer.mappingsObject, mappingData, mappingContainer.metadataMappingsObject, mappingContainer.object);
        }

        const payload = mappingData;
        payload.configuration_hash = this.configurationHash;
        delete payload.default_mappings;
        delete payload.required_mappings;

        try {
            const saveMappingData = await saveMappingConfigurations({
                jsonMappingConfigurationsObject: JSON.stringify(payload)
            });

            const responseData = JSON.parse(saveMappingData);
            if(responseData.error) {
                this.showToast(responseData.error, 'error', 'sticky');
                return;
            }

            const isConfigSaved = responseData.results.isConfigSaved;
            if(!isConfigSaved) {
                this.showToast('There was a problem saving the data mapping', 'error', 'sticky');
                return;  
            }

            this.configurationHash = responseData.results.configurationHash;

            this.showToast('Data mapping was successfully saved', 'success');
            saveSuccess = true;
            this.isMappingsUpdated = false;
        } catch (error) {
            const errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error', 'sticky');
        } finally {
            this.loading = false;

            // triggers the `oncompletesave` bound function, which is tied to `completeSave` in the setup component
            this.dispatchEvent(new CustomEvent('savecomplete', {
                detail: {
                    saveSuccess: saveSuccess,
                    configurationHash: this.configurationHash,
                }
            }));
        }
    }

    get allowTerminationMetadata() {
        return this.defaultSfObject === 'Order' || this.defaultSfObject === 'OrderItem';
    }

    @api updateConfigHash(configHash) {
        this.configurationHash = configHash;
    }

    showToast(message, variant, mode) {
        this.dispatchEvent(new CustomEvent('showtoast', {
            bubbles: true,
            composed: true,
            detail: {
                toast: {
                    message: message,
                    variant: variant ? variant : 'info',
                    mode: mode ? mode : 'dismissible'
                }
            }
        }));
    }
}