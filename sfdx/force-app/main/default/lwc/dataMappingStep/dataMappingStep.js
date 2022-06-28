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

    // TODO should be able to remove @track on all non-objects/arrays
    friendlyStripeObjectName = 'Customer';
    activeObject = 'customer';
    activeObjectDescription;
    @track sfFieldOptions = [];
    @track defaultCustomerSections = [
        'standard'
    ];
    @track defaultProductSections = [
        'standard'
    ];
    @track defaultSubscriptionScheduleSections = [
        'standard'
    ];
    @track defaultSubscriptionItemSections = [
        'standard'
    ];
    @track defaultPriceSections = [
        'standard'
    ];
    @track body;
    @track staticValue;
    @track picklistIndex;
    @track stripeObjectField
    @track stripeCustomerMappings;
    @track stripeProductMappings;
    @track stripeSubscriptionMappings;
    @track stripeSubscriptionItemMappings;
    @track stripePriceMappings;
    @track fieldListByObjectMap;
    // allMappingConfigurations holds all mappings retrieved from ruby, value is set in `getMappingConfigurations`
    @track allMappingConfigurations;
    @track defaultSfObject;
    @track isConnected = false;
    @track isMappingsUpdated = false;
    @track activeStripeObjectMappings;
    @track activeStripeObjectMetadataFields;
    @track activeStripeObjectSections;
    
    // allMappingList is used to retrieve and send all user mappings `saveMappingConfigurations` 
    @track allMappingList = {
        // these are static values defined by the user in the mapper UI
        field_defaults: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_item: {},
            price: {}
        },
        // these are user defined mappings from Salesforce fields to Stripe fields
        field_mappings: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_item: {},
            price: {}
        },
        // these are default mappings sent from Ruby (can be overeidden in mapp)
        default_mappings: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_item: {},
            price: {}
        },
        // these are required mappings sent from Ruby 
        required_mappings: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_item: {},
            price: {}
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
    @track subscriptionMetadataFields = {metadataMapping: {
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

    @track activeStripeObjectMappings = this.stripeCustomerMappings;
    @track activeStripeObjectMetadataFields = this.customerMetadataFields;
    @track activeStripeObjectSections = this.defaultCustomerSections;

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

    get priceObjectActive() {
        return this.activeObject == 'price';
    }

    get activeObjectFields() {
        let activeObjectName = this.activeObject;

        // TODO this special-casing should be removed in the future
        if (this.activeObject === 'subscription-item') {
            activeObjectName = 'subscriptionItem'
        }

        activeObjectName = activeObjectName.charAt(0).toUpperCase() + activeObjectName.slice(1);
        return this['stripe' + activeObjectName + 'Mappings'];
    }

    set activeObjectFields(value) {
        let parsedVal = JSON.parse(JSON.stringify(value));
        let activeObjectName = this.activeObject;

        // TODO this special-casing should be removed in the future
        if (this.activeObject === 'subscription-item') {
            activeObjectName = 'subscriptionItem'
        }

        // TODO we should try up this transformation and self-document what it actually is
        activeObjectName = activeObjectName.charAt(0).toUpperCase() + activeObjectName.slice(1);

        this['stripe' + activeObjectName + 'Mappings'] = parsedVal;
    }

    get activeMetadataObjectFields() {
        let activeObjectName = this.activeObject;

        // TODO this special-casing should be removed in the future
        if (this.activeObject === 'subscription-item') {
            activeObjectName = 'subscriptionItem';
        }

        return this[activeObjectName + 'MetadataFields'];
    }

    set activeMetadataObjectFields(value) {
        const parsedVal = JSON.parse(JSON.stringify(value));
        let activeObjectName = this.activeObject;

        // TODO this special-casing should be removed in the future
        if (this.activeObject === 'subscription-item') {
            activeObjectName = 'subscriptionItem'
        }

        this[activeObjectName + 'MetadataFields'] = parsedVal;
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

                if(!fieldData.sfValue) {
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
        }
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
        var fieldType = this.stripeObjectField.type.toLowerCase()
        //has to be a copy to force a rerender
        let modifiedFieldOptions = JSON.parse(JSON.stringify(this.sfFieldOptions))
        switch(fieldType) {
            case 'integer' || 'decimal' || 'number':
                this.sfFieldOptions = modifiedFieldOptions.filter(fieldOptions => fieldOptions.type === 'double' ||fieldOptions.type === 'reference');
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

       if(!this.fieldListByObjectMap) {
           return;
       }

        this.loading = true;

        // future: we need to clean this up. Code duplication like this destroys velocity later on. We should be able to something like changeObject(this.activeObject) and have a single assignment statement for all of the fields.
        switch(this.activeObject) {
            case 'customer':
                this.friendlyStripeObjectName = "Account";
                this.defaultSfObject = 'Account';
                this.fieldListByObjectMap.Account.sort(this.mapperFieldSorter);
                this.sfFieldOptions = this.fieldListByObjectMap.Account;
                this.activeStripeObjectMappings = this.stripeCustomerMappings;
                this.activeStripeObjectMetadataFields = this.customerMetadataFields;
                this.activeStripeObjectSections = this.defaultCustomerSections;
               break;

            // TODO should change identifier to `subscription_schedule` instead
            case 'subscription':
                this.friendlyStripeObjectName = "Subscription Schedule"
                this.defaultSfObject = 'Order';
                this.fieldListByObjectMap.Order.sort(this.mapperFieldSorter);
                this.sfFieldOptions = this.fieldListByObjectMap.Order;
                this.activeStripeObjectMappings = this.stripeSubscriptionMappings;
                this.activeStripeObjectMetadataFields = this.subscriptionMetadataFields;
                this.activeStripeObjectSections = this.defaultSubscriptionScheduleSections;
                break;

            case 'subscription-item':
                this.friendlyStripeObjectName = "Subscription Schedule Phase Item";
                this.defaultSfObject = 'OrderItem';
                this.fieldListByObjectMap.OrderItem.sort(this.mapperFieldSorter);
                this.sfFieldOptions = this.fieldListByObjectMap.OrderItem;
                this.activeStripeObjectMappings = this.stripeSubscriptionItemMappings;
                this.activeStripeObjectMetadataFields = this.subscriptionItemMetadataFields;
                this.activeStripeObjectSections = this.defaultSubscriptionItemSections;
                break;
    
            case 'product':
                this.friendlyStripeObjectName = "Product";
                this.defaultSfObject = 'Product2';
                this.fieldListByObjectMap.Product2.sort(this.mapperFieldSorter)
                this.sfFieldOptions = this.fieldListByObjectMap.Product2;
                this.activeStripeObjectMappings = this.stripeProductMappings;
                this.activeStripeObjectMetadataFields = this.productMetadataFields;
                this.activeStripeObjectSections = this.defaultProductSections;
                break;


            case 'price':
                this.friendlyStripeObjectName = "Price";
                this.defaultSfObject = 'PricebookEntry';
                this.fieldListByObjectMap.PricebookEntry.sort(this.mapperFieldSorter)
                this.sfFieldOptions = this.fieldListByObjectMap.PricebookEntry;
                this.activeStripeObjectMappings = this.stripePriceMappings;
                this.activeStripeObjectMetadataFields = this.priceMetadataFields;
                this.activeStripeObjectSections = this.defaultPriceSections;
                break;

            default:
                 // TODO should send to sentry when that is possible
                console.log("uncaught sidebar click");
                break;
        }

        this.openActiveSection();
        this.setDescription();
        this.loading = false;
        return;

    }

    // Open 'Metadata' mapping section if metadata mappings exist
    openActiveSection() {
        this.activeStripeObjectSections = ['standard']
        let activeStripeObjectName = this.activeObject; 
        // TODO this should not be special cased but changing this would require changes in the html and other places in this file and ensuring things do not break due to the change.
        if ( this.activeObject === 'subscription-item') {
            activeStripeObjectName = 'subscriptionItem';
        }
        if (this[activeStripeObjectName + 'MetadataFields'].metadataMapping.fields.length <= 0) {
            return
        }
        this.activeStripeObjectSections = ['standard', 'metadata'];
        return;
    }

    setDescription() {
        
        switch(this.activeObject) {
            case 'customer':
              this.activeObjectDescription = 'Customer objects allow you to perform recurring charges, and to track multiple charges, that are associated with the same customer. The API allows you to create, delete, and update your customers. You can retrieve individual customers as well as a list of all your customers.';
              break;
            case 'product':
                this.activeObjectDescription = 'Products describe the specific goods or services you offer to your customers. For example, you might offer a Standard and Premium version of your goods or service; each version would be a separate Product. They can be used in conjunction with Prices to configure pricing in Checkout and Subscriptions.';
                break;
            case 'subscription':
                this.activeObjectDescription = 'A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.';
                break;
            case 'subscription-item':
                this.activeObjectDescription = 'Subscription items allow you to create customer subscriptions with more than one plan, making it easy to represent complex billing relationships.';
                break;
            case 'price':
                this.activeObjectDescription = 'Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.';
                break;
            default:
                return;
        }
        
        return;

    }

    openStripeApi(event) {
        const stripeObject = event.currentTarget.dataset.object;
        let apiUrl = "https://stripe.com/docs/api/";

        // subscription_items are unique and really represent subscription schedule phase items
        if(stripeObject == 'subscription-item') {
            apiUrl += "subscription_schedules/create#create_subscription_schedule-phases-items"
        } else if(stripeObject == 'subscription') {
            // TODO right now the wrong identifier is being used for subscriptions... ugh. Fix this when `subscription` => `subscription_schedule`
            apiUrl += "subscription_schedules/create"
        } else {
            // our API calls are always #create, so let's link to that specific area in the docs
            apiUrl += `${stripeObject}s/create`;
        }

        window.open(apiUrl, '_blank');
    }

    @api async connectedCallback() {
        this.setDescription();
        this.dispatchEvent(new CustomEvent('contentloading'));
        try {
            const getFormattedStripeObjects = await getFormattedStripeObjectFields();
            const responseData = JSON.parse(getFormattedStripeObjects);
            if(responseData.error) {
                this.showToast(responseData.error, 'error', 'sticky');
                return;
            }
            this.stripeCustomerMappings = responseData.results.formattedStripeCustomerFields;
            this.stripeProductMappings = responseData.results.formattedStripeProductItemFields;
            this.stripeSubscriptionMappings = responseData.results.formattedStripeSubscriptionFields;
            this.stripeSubscriptionItemMappings = responseData.results.formattedStripeSubscriptionItemFields;
            this.stripePriceMappings = responseData.results.formattedStripePriceFields;

            this.getPicklistValuesForMapper(true, '', false);
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error');
        } finally {
            this.activeObject = 'customer';
            this.activeStripeObjectMappings = this.stripeCustomerMappings;
            this.activeStripeObjectMetadataFields = this.customerMetadataFields;
            this.activeStripeObjectSections = this.defaultCustomerSections;
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
            const listOfStripeMappingObjects = [
                {
                    object: 'customer',
                    mappingsObject: this.stripeCustomerMappings,
                    metadataMappingsObject: this.customerMetadataFields
                },
                {
                    object: 'product',
                    mappingsObject: this.stripeProductMappings,
                    metadataMappingsObject: this.productMetadataFields
                },
                {
                    object: 'subscription_schedule',
                    mappingsObject: this.stripeSubscriptionMappings,
                    metadataMappingsObject: this.subscriptionMetadataFields
                },
                {
                    object: 'subscription_item',
                    mappingsObject: this.stripeSubscriptionItemMappings,
                    metadataMappingsObject: this.subscriptionItemMetadataFields
                },
                {
                    object: 'price',
                    mappingsObject: this.stripePriceMappings,
                    metadataMappingsObject: this.priceMetadataFields
                }
            ];
            
            for (const mappingContainer of listOfStripeMappingObjects) {
                this.setFieldMappings(mappingContainer.object, mappingContainer.mappingsObject, mappingContainer.metadataMappingsObject.metadataMapping.fields);
            }
            
            this.allMappingList['default_mappings'] = this.allMappingConfigurations['default_mappings'];
            this.allMappingList['required_mappings'] = this.allMappingConfigurations['required_mappings'];

        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error', 'sticky');
        } finally {
            this.dispatchEvent(new CustomEvent('contentloadingcomplete'));
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
        this.openActiveSection();
        
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

        const listOfStripeMappingObjects = [
            {
                object: 'customer',
                mappingsObject: this.stripeCustomerMappings,
                metadataMappingsObject: this.customerMetadataFields
            },
            {
                object: 'product',
                mappingsObject: this.stripeProductMappings,
                metadataMappingsObject: this.productMetadataFields
            },
            {
                object: 'subscription_schedule',
                mappingsObject: this.stripeSubscriptionMappings,
                metadataMappingsObject: this.subscriptionMetadataFields
            },
            {
                object: 'subscription_item',
                mappingsObject: this.stripeSubscriptionItemMappings,
                metadataMappingsObject: this.subscriptionItemMetadataFields
            },
            {
                object: 'price',
                mappingsObject: this.stripePriceMappings,
                metadataMappingsObject: this.priceMetadataFields
            }
        ];

        for (const mappingContainer of listOfStripeMappingObjects) {
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