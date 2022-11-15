import { LightningElement, api, track } from 'lwc';
import getPicklistValuesForMapper from '@salesforce/apex/setupAssistant.getPicklistValuesForMapper';
import getFormattedStripeObjectFields from '@salesforce/apex/setupAssistant.getFormattedStripeObjectFields';
import getMappingConfigurations from '@salesforce/apex/setupAssistant.getMappingConfigurations';
import saveMappingConfigurations from '@salesforce/apex/setupAssistant.saveMappingConfigurations';
import { getErrorMessage } from 'c/utils'

export default class DataMappingStep extends LightningElement {
    /*
     `@track` decorator is needed here to tell the lwc framework to observe 
      changes to the properties of an object and to the elements of an array
     */

    friendlyStripeObjectName = 'Customer';
    activeObject = 'customer';
    activeObjectDescription;
    staticValue;
    stripeObjectField;
    defaultSfObject;
    isConnected = false;
    isMappingsUpdated = false;
    contentLoading = new CustomEvent('contentloading');
    contentLoadingComplete = new CustomEvent('contentloadingcomplete');

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
            description: 'Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.'
        },
        "priceOrderItem": {
            objectName: "PricebookEntry",
            friendlyName: 'Price Order Item',
            description: 'Price order items define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of a single product.'
        },
        "coupon": {
            objectName: "Stripe_Coupon_Beta__c",
            friendlyName: 'Stripe Coupon',
            description: 'Coupons contain information about a percent-off or amount-off discount you might want to apply to a subscription or subscription item.'
        }
    }
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
    @track activeStripeObjectMappings;
    @track activeStripeObjectMetadataFields;
    
    // allMappingList is used to retrieve and send all user mappings `saveMappingConfigurations` 
    @track allMappingList = {
        // these are static values defined by the user in the mapper UI
        field_defaults: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_phase: {},
            subscription_item: {},
            price: {},
            price_order_item: {},
            coupon: {}
        },
        // these are user defined mappings from Salesforce fields to Stripe fields
        field_mappings: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_phase: {},
            subscription_item: {},
            price: {},
            price_order_item: {},
            coupon: {}
        },
        // these are default mappings sent from Ruby (can be overridden in map)
        default_mappings: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_phase: {},
            subscription_item: {},
            price: {},
            price_order_item: {},
            coupon: {}
        },
        // these are required mappings sent from Ruby 
        required_mappings: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_phase: {},
            subscription_item: {},
            price: {},
            price_order_item: {},
            coupon: {}
        }
    };
    @track customerMetadataFields = {metadataMapping: {
        label: '',
        name: '',
        value: '',
        description: '',
        fields: []
    }};
    @track productMetadataFields = {metadataMapping: {
        label: '',
        name: '',
        value: '',
        description: '',
        fields: []
    }};;
    @track subscriptionScheduleMetadataFields = {metadataMapping: {
        label: '',
        name: '',
        value: '',
        description: '',
        fields: []
    }};;
    @track subscriptionPhaseMetadataFields = {metadataMapping: {
        label: '',
        name: '',
        value: '',
        description: '',
        fields: []
    }};;
    @track subscriptionItemMetadataFields = {metadataMapping: {
        label: '',
        name: '',
        value: '',
        description: '',
        fields: []
    }};;
    @track priceMetadataFields = {metadataMapping: {
        label: '',
        name: '',
        value: '',
        description: '',
        fields: []
    }};
    @track priceOrderItemMetadataFields = {metadataMapping: {
        label: '',
        name: '',
        value: '',
        description: '',
        fields: []
    }};
    @track couponMetadataFields = {metadataMapping: {
        label: '',
        name: '',
        value: '',
        description: '',
        fields: []
    }};

    @track activeStripeObjectMappings = this.customerMappings;
    @track activeStripeObjectMetadataFields = this.customerMetadataFields;
    @track activeStripeObjectSections;

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

    get listOfStripeMappingObjects() {
        return [
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
        ];
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
                const metadataKey = 'metadata.' + fieldData.value;

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
        if (targetFieldIndex ) {
            const updatedSelection = {
                staticValue: true,
                sfValue: this.staticValue,
                sfValueType: 'staticMetadata'
            };
            Object.assign(this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)] , updatedSelection);
            this.valueChange();
        }
    }

    updateStaticValueChoice(event) {
        this[event.target.name] = event.target.value;
        this.staticValue = this[event.target.name];
        const targetSectionIndex = event.currentTarget.closest('lightning-accordion-section').dataset.index;
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (targetSectionIndex && targetFieldIndex ) {
            let updatedSelection = {
                hasSfValue: true,
                staticValue: true,
                sfValue: this.staticValue,
                sfValueType: 'static'
            };
            Object.assign(this.activeObjectFields[targetSectionIndex].fields[parseInt(targetFieldIndex)] , updatedSelection)
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
        this.activeStripeObjectMappings = this[this.activeObject + 'Mappings'];
        this.activeStripeObjectMetadataFields = this[this.activeObject + METADATAFIELDS];
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
            
            this.getPicklistValuesForMapper(true, '', false);
        } catch (error) {
            const errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error');
        } finally {
            this.activeObject = 'customer';
            this.activeObjectDescription = this.ACTIVE_OBJECT_INFO[this.activeObject]['description'];
            this.activeStripeObjectMappings = this.customerMappings;
            this.activeStripeObjectMetadataFields = this.customerMetadataFields;
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
            this.allMappingConfigurations = mappingConfigurationResponseData.results.allMappingConfigurations;
           
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
            sfValue: sfValue,
            sfValueType: sfValueType
        };
        return fieldObject;
    }

    @api async saveCongfiguredMappings() {
        this.loading = true;
        let saveSuccess = false;

        for (const mappingContainer of this.listOfStripeMappingObjects) {
            this.allMappingList = this.saveObjectMappings(mappingContainer.mappingsObject, this.allMappingList, mappingContainer.metadataMappingsObject, mappingContainer.object);
        }

        try {
            const saveMappingData = await saveMappingConfigurations({
                jsonMappingConfigurationsObject: JSON.stringify(this.allMappingList)
            });

            const responseData = JSON.parse(saveMappingData);
            if(responseData.error) {
                this.showToast(responseData.error, 'error', 'sticky');
                return;
            }

            const isConfigSaved = responseData.results.isConfigSaved;
            if(!isConfigSaved) {
                this.showToast('There was a problem saving data mapping', 'error', 'sticky');
                return;  
            } 

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
                    saveSuccess: saveSuccess
                }
            }));
        }
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