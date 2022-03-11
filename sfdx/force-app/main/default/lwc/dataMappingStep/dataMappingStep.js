import { LightningElement, api, track } from 'lwc';
import { loadScript } from 'lightning/platformResourceLoader'
import getPicklistValuesForMapper from '@salesforce/apex/setupAssistant.getPicklistValuesForMapper';
import getMappingConfigurations from '@salesforce/apex/setupAssistant.getMappingConfigurations';
import saveMappingConfigurations from '@salesforce/apex/setupAssistant.saveMappingConfigurations';
import saveData from '@salesforce/apex/setupAssistant.saveData';
import showdownJs from '@salesforce/resourceUrl/showdownJS';

export default class DataMappingStep extends LightningElement {
    @track activeObject = '';
    @track sfFieldOptions = [];
    @track defaultCustomerSections = [
        'standard'
    ];
    @track defaultProductSections = [
        'standard'
    ];
    @track defaultSubscriptionSections = [
        'standard'
    ];
    @track defaultSubscriptionItemSections = [
        'standard'
    ];
    @track defaultPriceSections = [
        'standard'
    ];

    @track body
    @track staticValue
    @track picklistIndex;
    @track stripeObjectField
    @track stripeCustomerMappings;
    @track stripeProductMappings;
    @track stripeSubscriptionMappings;
    @track stripeSubscriptionItemMappings;
    @track stripePriceMappings;
    @track defaultSfObject;
    @track fieldListByObjectMap;
    @track allMappingConfigurations
    @track isConnected = false;
    markdownConverter;  // Stores instance of Showdown Markdown converter; set on connectedCallback 

    @track allMappingList = {
        field_defaults: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_item: {},
            price: {}
        },
        field_mappings: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_item: {},
            price: {}
        },
        default_mappings: {
            customer: {},
            product: {},
            subscription_schedule: {},
            subscription_item: {},
            price: {}
        },
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
        if (this.activeObject === 'subscription-item')activeObjectName = 'subscriptionItem'
        activeObjectName = activeObjectName.charAt(0).toUpperCase() + activeObjectName.slice(1);
        return this['stripe' + activeObjectName + 'Mappings'];
    }

    set activeObjectFields(value) {
        let parsedVal = JSON.parse(JSON.stringify(value));
        let activeObjectName = this.activeObject;
        if (this.activeObject === 'subscription-item')activeObjectName = 'subscriptionItem'
        activeObjectName = activeObjectName.charAt(0).toUpperCase() + activeObjectName.slice(1);
        this['stripe' + activeObjectName + 'Mappings'] = parsedVal;
    } 

    get activeMetadataObjectFields() {
        let activeObjectName = this.activeObject;
        if (this.activeObject === 'subscription-item')activeObjectName = 'subscriptionItem'
        return this[activeObjectName + 'MetadataFields'];
    }

    set activeMetadataObjectFields(value) {
        let parsedVal = JSON.parse(JSON.stringify(value));
        let activeObjectName = this.activeObject;
        if (this.activeObject === 'subscription-item')activeObjectName = 'subscriptionItem'
        this[activeObjectName + 'MetadataFields'] = parsedVal;
    } 

    saveObjectMappings(stripeObjectMappings, listOfAllMappings, listOfMetadataFields, stripeObjectName) {
        for(let i = 0; i < stripeObjectMappings.length; i++) {
            for(let j = 0; j < stripeObjectMappings[i].fields.length; j++) {
                if(stripeObjectMappings[i].fields[j].sfValue) {
                    if(stripeObjectMappings[i].fields[j].staticValue === true) {
                        listOfAllMappings.field_defaults[stripeObjectName][stripeObjectMappings[i].fields[j].value] = stripeObjectMappings[i].fields[j].sfValue;
                    } else {
                        listOfAllMappings.field_mappings[stripeObjectName][stripeObjectMappings[i].fields[j].value] = stripeObjectMappings[i].fields[j].sfValue;
                    }
                    if(stripeObjectMappings[i].fields[j].defaultValue)listOfAllMappings.default_mappings[stripeObjectName][stripeObjectMappings[i].fields[j].value] = stripeObjectMappings[i].fields[j].defaultValue;
                }
            }
        }

        if(listOfMetadataFields.metadataMapping.fields.length) {
           for(let i = 0; i < listOfMetadataFields.metadataMapping.fields.length; i++) {
                if(listOfMetadataFields.metadataMapping.fields[i].sfValue) {
                    if(listOfMetadataFields.metadataMapping.fields[i].staticValue === true) {
                        listOfAllMappings.field_defaults[stripeObjectName]['metadata.'+listOfMetadataFields.metadataMapping.fields[i].value] = listOfMetadataFields.metadataMapping.fields[i].sfValue;
                    } else {
                        listOfAllMappings.field_mappings[stripeObjectName]['metadata.'+listOfMetadataFields.metadataMapping.fields[i].value] = listOfMetadataFields.metadataMapping.fields[i].sfValue;
                    }
                    if(listOfMetadataFields.metadataMapping.fields[i].defaultValue)listOfAllMappings.default_mappings[stripeObjectName]['metadata.'+listOfMetadataFields.metadataMapping.fields[i].value] = listOfMetadataFields.metadataMapping.fields[i].defaultValue;
                }
            } 
        }
    }


    toggleMetaStaticValue(event) {
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index; 
        if (targetFieldIndex ) {
            this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)].staticValue = !this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)].staticValue;
            if(this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)].staticValue === true)this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)].sfValue = ''; 
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
        }
    }

    removeMetaRow(event) {
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index; 
        if (targetFieldIndex ) {
            const value = this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)];
            this.activeMetadataObjectFields.metadataMapping.fields.splice(this.activeMetadataObjectFields.metadataMapping.fields.findIndex(metadataField => metadataField.name === value.name),1);
        }
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
                hasOverride: true,
                staticValue: false
            };
            updatedSelection.hasOverride = this.activeObjectFields[targetSectionIndex].fields[parseInt(targetFieldIndex)].defaultValue ? true : false;
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
        var fieldType = this.stripeObjectField.type.toLowerCase()
        //has to be a copy to force a rerender
        let modifiedFieldOptions = JSON.parse(JSON.stringify(this.sfFieldOptions))
        if(fieldType === 'integer' || fieldType === 'decimal' || fieldType === 'number') {
            this.sfFieldOptions = modifiedFieldOptions.filter(fieldOptions => fieldOptions.type === 'double' ||fieldOptions.type === 'reference') 
        } else if (fieldType === 'timestamp') {
            this.sfFieldOptions = modifiedFieldOptions.filter(fieldOptions => fieldOptions.type.includes('date') || fieldOptions.type === 'reference' ) 
        } else if (fieldType === 'boolean') {
            this.sfFieldOptions = modifiedFieldOptions.filter(fieldOptions => fieldOptions.type === 'boolean' || fieldOptions.type === 'reference' ) 
        } 
    }

    updateMetaPicklist(event) {
        const targetFieldIndex = event.currentTarget.closest('tr').dataset.index;
        if (this.targetFieldIndex ) {
            this.activeMetadataObjectFields.metadataMapping.fields[parseInt(targetFieldIndex)].sfValue = event.detail.value;
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

    changeActiveObject(event) {
        this.activeObject = event.detail.name;
          if(this.activeObject === 'customer' && this.fieldListByObjectMap){
            this.defaultSfObject = 'Account';
            this.fieldListByObjectMap.Account.sort(function(a, b) {
                    return a.label.localeCompare(b.label);
            });
            this.sfFieldOptions = this.fieldListByObjectMap.Account
         } else if(this.activeObject  === 'product' && this.fieldListByObjectMap){
            this.defaultSfObject = 'Product2';
            this.fieldListByObjectMap.Product2.sort(function(a, b) {
                    return a.label.localeCompare(b.label);
            });
            this.sfFieldOptions = this.fieldListByObjectMap.Product2
            if(this.fieldListByObjectMap)this.setFieldMappings('product', this.stripeProductMappings, this.productMetadataFields.metadataMapping.fields);
         } else if(this.activeObject === 'subscription' && this.fieldListByObjectMap) {
            this.defaultSfObject = 'Order';
            this.fieldListByObjectMap.Order.sort(function(a, b) {
                    return a.label.localeCompare(b.label);
            });
            this.sfFieldOptions = this.fieldListByObjectMap.Order
            if(this.fieldListByObjectMap) {
                //Arnold delete this when mike fixed nomenclature in ruby service
                this.setFieldMappings('subscription', this.stripeSubscriptionMappings, this.subscriptionMetadataFields.metadataMapping.fields);
                this.setFieldMappings('subscription_schedule', this.stripeSubscriptionMappings, this.subscriptionMetadataFields.metadataMapping.fields);
                 
            }
         } else if(this.activeObject === 'subscription-item' && this.fieldListByObjectMap) {
            this.defaultSfObject = 'Order Item';
            this.fieldListByObjectMap.OrderItem.sort(function(a, b) {
                    return a.label.localeCompare(b.label);
            });
            this.sfFieldOptions = this.fieldListByObjectMap.OrderItem
            if(this.fieldListByObjectMap)this.setFieldMappings('subscription_item', this.stripeSubscriptionItemMappings, this.subscriptionItemMetadataFields.metadataMapping.fields);
         } else if(this.activeObject === 'price' && this.fieldListByObjectMap) {
            this.defaultSfObject = 'Pricebook Entry';
            this.fieldListByObjectMap.PricebookEntry.sort(function(a, b) {
                   return a.label.localeCompare(b.label);
           });
           this.sfFieldOptions = this.fieldListByObjectMap.PricebookEntry
           if(this.fieldListByObjectMap)this.setFieldMappings('price', this.stripePriceMappings, this.priceMetadataFields.metadataMapping.fields);
        }
    }

    async connectedCallback() {
        try {
            loadScript(this, showdownJs + '/showdown-1.9.1/dist/showdown.js').then(()=> {
                this.markdownConverter = new showdown.Converter();
            });

            const stripeOpenSpecEndPoint = "https://raw.githubusercontent.com/stripe/openapi/master/openapi/spec3.json";
            const response = await fetch(stripeOpenSpecEndPoint);
            const repos = await response.json();
            this.body = repos;
            this.stripeCustomerMappings = this.formatStripeObjectsForMapper( this.body['components']['schemas']['customer']['properties']);
            this.stripeProductMappings = this.formatStripeObjectsForMapper(this.body['components']['schemas']['product']['properties']);
            this.stripeSubscriptionMappings = this.formatStripeObjectsForMapper(this.body['components']['schemas']['subscription']['properties']);
            this.stripeSubscriptionItemMappings = this.formatStripeObjectsForMapper(this.body['components']['schemas']['subscription_item']['properties']);
            this.stripePriceMappings = this.formatStripeObjectsForMapper(this.body['components']['schemas']['price']['properties']);
            this.getPicklistValuesForMapper(true, '', false);
        } catch (error) {
            this.showToast(error.message, 'error');
        } finally {
            this.isLoaded = true;
            this.activeObject = 'customer';
        }
    }

    formatStripeObjectsForMapper(objectJsonConfigMap) {
        var stripeObjectMappings = [
            {
                label: 'Standard Mappings',
                name: 'standard',
                description: '',
                fields: []
            }
        ]
    
        for (const field in objectJsonConfigMap) { 
            if(!objectJsonConfigMap[field]['$ref'] && !objectJsonConfigMap[field]['anyOf']) {
                if(objectJsonConfigMap[field]['type'] && objectJsonConfigMap[field]['type'] !== 'object' && objectJsonConfigMap[field]['type'] !== 'array') {
                    let fieldMap = {
                        name: field.replaceAll('_',' '),
                        value: field,
                        description: '',
                        type: '',
                        defaultValue: '',
                        hasOverride: false,
                        staticValue: false,
                        hasSfValue: false,
                        sfValue: '',
                        sfValueType: ''
                    };
                    if(objectJsonConfigMap[field]['description'])fieldMap['description'] = this.markdownConverter.makeHtml(objectJsonConfigMap[field]['description']);
                    fieldMap['type'] = objectJsonConfigMap[field]['type']    
                    stripeObjectMappings[0].fields.push(fieldMap)
                    stripeObjectMappings[0].fields.sort(function(a, b) {
                        return a.name.localeCompare(b.name);
                    });

                }
            } else if(objectJsonConfigMap[field]['$ref']) {         
                var expandableSchemaFieldName = objectJsonConfigMap[field]['$ref'].split('/').pop();
                var  expandableSchemaFieldMap = this.body['components']['schemas'][expandableSchemaFieldName]['properties']

                let newSection = {
                    label: field.charAt(0).toUpperCase() + field.slice(1).replaceAll('_',' '),
                    name: field,
                    description: '',
                    fields: []
                }
                stripeObjectMappings.push(newSection)
                for (const expandableField in expandableSchemaFieldMap) {
                    if(expandableSchemaFieldMap[expandableField]['type'] && expandableSchemaFieldMap[expandableField]['type'] !== 'object' && expandableSchemaFieldMap[expandableField]['type'] !== 'array') {
                        var hashFieldName = expandableField.replaceAll('_',' ')
                        var hashFieldValue = field+'.'+expandableField;
                        let fieldExpandableMap = {
                            name: hashFieldName,
                            value: hashFieldValue,
                            description: '',
                            type: '',
                            defaultValue: '',
                            hasOverride: false,
                            staticValue: false,
                            hasSfValue: false,
                            sfValue: '',
                            sfValueType: ''
                        };
                        if(expandableSchemaFieldMap[expandableField]['description'])fieldExpandableMap['description'] = this.markdownConverter.makeHtml(expandableSchemaFieldMap[expandableField]['description']);
                        fieldExpandableMap['type'] = expandableSchemaFieldMap[expandableField]['type'] 
                        if(expandableSchemaFieldMap[expandableField]['properties']) {  
                            stripeObjectMappings = this.checkForNestedProperties(expandableField,expandableSchemaFieldMap,stripeObjectMappings);
                        } else {
                             var index = stripeObjectMappings.findIndex(objectSection => {
                            return objectSection.name === field;
                            });
                            if(index) {
                                stripeObjectMappings[index].fields.push(fieldExpandableMap);
                            } else {
                                stripeObjectMappings[stripeObjectMappings.length-1].fields.push(fieldExpandableMap);
                            }
                        }
                    }
                }
            } else if(objectJsonConfigMap[field]['anyOf']) {
                if(objectJsonConfigMap[field]['anyOf'].length && objectJsonConfigMap[field]['anyOf'][0]['$ref']) {         
                    var  nestedExpandableField = objectJsonConfigMap[field]['anyOf'][0]['$ref'].split('/').pop()
                    var nestedExpandableFieldMap = this.body['components']['schemas'][nestedExpandableField]['properties']
                    if(nestedExpandableFieldMap)stripeObjectMappings = this.checkforNestedFields(field,objectJsonConfigMap,stripeObjectMappings, nestedExpandableFieldMap)
                }

            }        
        }
          stripeObjectMappings = stripeObjectMappings.filter(function( section ) {
            return section.fields.length;
        });
          
        return stripeObjectMappings
    }

    checkforNestedFields(field, objectJsonConfigMap, stripeObjectMappings, nestedExpandableFieldMap) {
            let newSection = {
                label: field.charAt(0).toUpperCase() + field.slice(1).replaceAll('_',' ').replaceAll('.',' '),
                name: field,
                description: '',
                fields: []
            }
            stripeObjectMappings.push(newSection)
            for (const expandableField in nestedExpandableFieldMap) {
                if(nestedExpandableFieldMap[expandableField]['description']) {
                    if( (nestedExpandableFieldMap[expandableField]['type'] && nestedExpandableFieldMap[expandableField]['type'] !== 'object' && nestedExpandableFieldMap[expandableField]['type'] !== 'array')) {
                        var hashFieldName = expandableField.replaceAll('_',' ')
                        var hashFieldValue = field+'.'+expandableField;
                        let fieldExpandableMap = {
                            name: hashFieldName,
                            value: hashFieldValue,
                            description: '',
                            type: '',
                            defaultValue: '',
                            hasOverride: false,
                            staticValue: false,
                            hasSfValue: false,
                            sfValue: '',
                            sfValueType: ''
                        };


                        fieldExpandableMap['description'] = this.markdownConverter.makeHtml(nestedExpandableFieldMap[expandableField]['description']);
                        fieldExpandableMap['type'] = nestedExpandableFieldMap[expandableField]['type']  
                        
                        stripeObjectMappings[stripeObjectMappings.length-1].fields.sort(function(a, b) {
                            return a.name.localeCompare(b.name);
                        });

                        var index = stripeObjectMappings.findIndex(objectSection => {
                            return objectSection.name === field;
                        });
                        if(index) {
                            stripeObjectMappings[index].fields.push(fieldExpandableMap);
                        } else {
                            stripeObjectMappings[stripeObjectMappings.length-1].fields.push(fieldExpandableMap);
                        }
                    }
                } else {
                    //HERE WE EXPLICTLY EXCLUDE NESTED (I.E Shipping.Adresss.Countty)
                    if(this.body['components']['schemas'][expandableField]['properties'] && expandableField !== 'coupon') {
                        var newNestedExpandableFieldMap = this.body['components']['schemas'][expandableField]['properties'];
                        var newNestedFieldName = field+'.'+expandableField.charAt(0).toUpperCase() + expandableField.slice(1).replaceAll('_',' ');
                        stripeObjectMappings = this.checkforNestedFields(newNestedFieldName,objectJsonConfigMap,stripeObjectMappings, newNestedExpandableFieldMap);
                    }
                }
            }
        
        return stripeObjectMappings; 
    }
    
    @api async getPicklistValuesForMapper(isConnectedCallback, ObjectName, isMetadataRow, targetElement) {
        this.loading = true;
        try {
            const getPicklistValues = await getPicklistValuesForMapper({ 
                isConnectedCallback: isConnectedCallback,
                ObjectApiName: ObjectName
            });
            this.data =  JSON.parse(getPicklistValues);
            if(this.data.isSuccess) {
                if(this.data.results.isConnected) {
                    this.isConnected = true;
                    if(isConnectedCallback === true) {
                        this.fieldListByObjectMap = this.data.results.fieldListByObjectMap;
                        this.sfFieldOptions = this.fieldListByObjectMap.Account
                        this.defaultSfObject = 'Account';
                        this.sfFieldOptions.sort(function(a, b) {
                            return a.label.localeCompare(b.label);
                        });
                    } else {
                        this.fieldListByObjectMap[this.data.results.ObjectApiName] = this.data.results.listOfObjectFields;
                        
                        this.sfFieldOptions = this.fieldListByObjectMap[this.data.results.ObjectApiName]
                        this.sfFieldOptions.sort(function(a, b) {
                            return a.label.localeCompare(b.label);
                        });
                        if(!isMetadataRow)this.filterSfOptionsByStripeFieldType();
                    }
                }
            } else {
                this.showToast(this.data.error, 'error');
            }

            const getMappingConfigs = await getMappingConfigurations();
            this.data =  JSON.parse(getMappingConfigs);
            if(this.data.isSuccess) {
                if(this.data.results.isConnected) {
                    this.allMappingConfigurations = this.data.results.allMappingConfigurations;
                    this.setFieldMappings('customer', this.stripeCustomerMappings, this.customerMetadataFields.metadataMapping.fields)
                }
            } else {
                this.showToast(this.data.error, 'error');
            }
        } catch (error) {
            this.showToast(error.message, 'error');
        } finally {
            if(targetElement) targetElement.dropdownLoading = false;
        }
    }
    
    setFieldMappings(stripeObject, stripeObjectMap, metadataFieldList) {
        for(let i = 0; i < stripeObjectMap.length; i++) {
            for(let j = 0; j < stripeObjectMap[i].fields.length; j++) {
                if(this.allMappingConfigurations.default_mapping && 
                   this.allMappingConfigurations.default_mapping[stripeObject] && 
                   this.allMappingConfigurations.default_mapping[stripeObject].length !== 0) {

                    for(const value in this.allMappingConfigurations.default_mapping[stripeObject]) {
                        if(stripeObjectMap[i].fields[j].value === value && this.allMappingConfigurations.default_mapping[stripeObject][value]) {
                            stripeObjectMap[i].fields[j].sfValue = '';
                            stripeObjectMap[i].fields[j].hasOverride = false; 
                            stripeObjectMap[i].fields[j].hasSfValue = false; 
                            stripeObjectMap[i].fields[j].staticValue = false; 
                            stripeObjectMap[i].fields[j].defaultValue = this.allMappingConfigurations.default_mapping[stripeObject][value];
                        } 
                    }
                }
                if(this.allMappingConfigurations.field_defaults && this.allMappingConfigurations.field_defaults[stripeObject]) {
                    for(const value in this.allMappingConfigurations.field_defaults[stripeObject]) {
                        if(stripeObjectMap[i].fields[j].value === value) {
                            stripeObjectMap[i].fields[j].staticValue = true
                            stripeObjectMap[i].fields[j].sfValue = this.allMappingConfigurations.field_defaults[stripeObject][value];
                            if(stripeObjectMap[i].fields[j].sfValue) {
                                stripeObjectMap[i].fields[j].hasSfValue = true
                            }
                            if(stripeObjectMap[i].fields[j].defaultValue) {
                                stripeObjectMap[i].fields[j].hasOverride = true
                            }
                        } else if (value.startsWith('metadata.')) {
                            let realfieldMapValue = value.replace('metadata.','')
                            let metadataFieldObj = {
                                name: realfieldMapValue,
                                value: realfieldMapValue,
                                type: 'metadata',
                                defaultValue: '',
                                hasOverride: false,
                                staticValue: true,
                                sfValue: this.allMappingConfigurations.field_defaults[stripeObject][value],
                                sfValueType: 'metadata'
                            };
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
                        if(stripeObjectMap[i].fields[j].value === value && !value.startsWith('metadata.')) {
                            stripeObjectMap[i].fields[j].staticValue = false
                            stripeObjectMap[i].fields[j].sfValue = this.allMappingConfigurations.field_mappings[stripeObject][value];
                            if(stripeObjectMap[i].fields[j].sfValue)stripeObjectMap[i].fields[j].hasSfValue = true
                            if(stripeObjectMap[i].fields[j].defaultValue)stripeObjectMap[i].fields[j].hasOverride = true
                        } else if (value.startsWith('metadata.') ) {
                            let realValue = value.replace('metadata.','')
                            let metadataFieldObj = {
                                name: realValue,
                                value: realValue,
                                type: 'metadata',
                                defaultValue: '',
                                hasOverride: false,
                                staticValue: false,
                                sfValue: this.allMappingConfigurations.field_mappings[stripeObject][value],
                                sfValueType: 'metadata'
                            };
                            if (metadataFieldList.length && metadataFieldList.filter(field => field.name === realValue).length > 0) {
                                //do nothing
                            } else {
                                metadataFieldList.push(metadataFieldObj);
                            }
                        } 
                    }
                }
            }
        }
    }

    async next(event) {
        event.stopPropagation();
        this.loading = true;
        this.saveObjectMappings(this.stripeCustomerMappings,this.allMappingList,this.customerMetadataFields,'customer');
        this.saveObjectMappings(this.stripeProductMappings,this.allMappingList,this.productMetadataFields,'product');
        this.saveObjectMappings(this.stripeSubscriptionMappings,this.allMappingList,this.subscriptionMetadataFields,'subscription_schedule');
        this.saveObjectMappings(this.stripeSubscriptionItemMappings,this.allMappingList,this.subscriptionItemMetadataFields,'subscription_item');
        this.saveObjectMappings(this.stripePriceMappings,this.allMappingList,this.priceMetadataFields,'price');
        try {
            const saveMappingData = await saveMappingConfigurations({ 
                jsonMappingConfigurationsObject: JSON.stringify(this.allMappingList)
            });
            this.data =  JSON.parse(saveMappingData);
            if(this.data.isSuccess) {
                let isConfigSaved = this.data.results.isConfigSaved;
                if(isConfigSaved === true) {
                    this.showToast('Saved successfully', 'success')
                } else {
                    this.showToast('The mappings were not saved successfully', 'error')
                }
            } else {
                this.showToast(this.data.error, 'error');
            }

            const saveSetupData = await saveData({
                newSetupDataRec: {
                    Steps_Completed__c: JSON.stringify({
                        'C-DATA-MAPPING-STEP': 2
                    })
                }
            });
            this.data =  JSON.parse(saveSetupData);
            if(this.data.isSuccess) {
                this.setupData = this.data.setupData;
                this.dispatchEvent(new CustomEvent('next', {
                    bubbles: true,
                    composed: true
                }));
            } else {
                this.showToast(this.data.error, 'error');
            }
        } catch (error) {
            this.showToast(error.message, 'error');
        } finally {
            this.loading = false;
        }
    }

    back() {
        this.dispatchEvent(new CustomEvent('exit', {
            bubbles: true,
            composed: true
        }));
    }

    showToast(message, variant, mode) {
        this.dispatchEvent(new CustomEvent('showtoast', {
            bubbles: true,
            composed: true,
            detail: { 
                toast: {
                    message: message,
                    variant: variant,
                    mode: mode
                }
            }
        })); 
    }
}