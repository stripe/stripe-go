const HTTPS = require('https')
const showdown = require('showdown');
const markdownConverter = new showdown.Converter();
const excludedReadOnlyFields = {
    all: [
        'created',
        'object',
        'livemode',
    ],

    customer: [
        'delinquent',
        'discount',
        'tax.ip_address'
    ],

    product: [],

    subscription: [
        'id',
        'current_period_end',
        'current_period_start',
        'ended_at',
        'status'
    ],

    subscription_item: [],

    price: ['id'],
}

const OPTIONS = {
    hostname: 'raw.githubusercontent.com',
    port: 443,
    path: 'stripe/openapi/master/openapi/spec3.json',
    method: 'GET'
}
var responseJson = '';
const HTTPREQUEST = HTTPS.request(OPTIONS, HttpResponse => {
    HttpResponse.on('data', responseDataChunk => {
        responseJson += responseDataChunk.toString();
    })

    HttpResponse.on('end', () => {
        responseJson = JSON.parse(responseJson);
        var formattedStripeObjectsForMapper = {
            formattedStripeCustomerFields: formatStripeObjectsForMapper(responseJson['components']['schemas']['customer']['properties'], excludedReadOnlyFields.customer),
            formattedStripeProductItemFields: formatStripeObjectsForMapper(responseJson['components']['schemas']['product']['properties'], excludedReadOnlyFields.product),
            formattedStripeSubscriptionFields: formatStripeObjectsForMapper(responseJson['components']['schemas']['subscription']['properties'], excludedReadOnlyFields.subscription),
            formattedStripeSubscriptionItemFields: formatStripeObjectsForMapper(responseJson['components']['schemas']['subscription_item']['properties'], excludedReadOnlyFields.subscription_item),
            formattedStripePriceFields: formatStripeObjectsForMapper(responseJson['components']['schemas']['price']['properties'], excludedReadOnlyFields.price)
        };
        formattedStripeObjectsForMapper = JSON.stringify(formattedStripeObjectsForMapper);
        console.log(formattedStripeObjectsForMapper);
        return formattedStripeObjectsForMapper;
    })
})

HTTPREQUEST.on('error', error => {
    console.log(error);
})

HTTPREQUEST.end();

function formatStripeObjectsForMapper(stripeObjectToFormat, objectExcludedReadOnlyFields) {
    var stripeObjectMappings = [{
        label: 'Standard Mappings',
        name: 'standard',
        description: '',
        fields: []
    }];
    for (const field in stripeObjectToFormat) {
        if (!stripeObjectToFormat[field]['$ref'] && !stripeObjectToFormat[field]['anyOf'] && !excludedReadOnlyFields.all.includes(field) && !objectExcludedReadOnlyFields.includes(field)) {
            if (stripeObjectToFormat[field]['type'] && stripeObjectToFormat[field]['type'] !== 'object' && stripeObjectToFormat[field]['type'] !== 'array') {
                var fieldMap = {
                    name: field.replace(/_+/g, ' '),
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
                if (stripeObjectToFormat[field]['description']) {
                    fieldMap['description'] = markdownConverter.makeHtml(stripeObjectToFormat[field]['description']);
                }
                fieldMap['type'] = stripeObjectToFormat[field]['type'];
                stripeObjectMappings[0].fields.push(fieldMap);
                stripeObjectMappings[0].fields.sort(function(a, b) {
                    return a.name.localeCompare(b.name);
                });
            }
        } else if (stripeObjectToFormat[field]['$ref']) {
            var expandableSchemaFieldName = stripeObjectToFormat[field]['$ref'].split('/').pop();
            var expandableSchemaFieldMap = responseJson['components']['schemas'][expandableSchemaFieldName]['properties'];
            var newSection = {
                label: field.charAt(0).toUpperCase() + field.slice(1).replace(/_+/g, ' '),
                name: field,
                description: '',
                fields: []
            };
            stripeObjectMappings.push(newSection);

            for (const expandableField in expandableSchemaFieldMap) {
                if (expandableSchemaFieldMap[expandableField]['type'] && expandableSchemaFieldMap[expandableField]['type'] !== 'object'
                && expandableSchemaFieldMap[expandableField]['type'] !== 'array' && !excludedReadOnlyFields.all.includes(expandableField) && !objectExcludedReadOnlyFields.includes(expandableField)) {
                    var hashFieldName = expandableField.replace(/_+/g, ' ');
                    var hashFieldValue = field + '.' + expandableField;
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
                    if (expandableSchemaFieldMap[expandableField]['description']) {
                        fieldExpandableMap['description'] = markdownConverter.makeHtml(expandableSchemaFieldMap[expandableField]['description']);
                    }
                    fieldExpandableMap['type'] = expandableSchemaFieldMap[expandableField]['type'];
                    if (expandableSchemaFieldMap[expandableField]['properties']) {
                        stripeObjectMappings = this.checkForNestedProperties(expandableField, expandableSchemaFieldMap, stripeObjectMappings);
                    } else {
                        var index = stripeObjectMappings.findIndex(objectSection => {
                            return objectSection.name === field;
                        });
                        if (index) {
                            stripeObjectMappings[index].fields.push(fieldExpandableMap);
                        } else {
                            stripeObjectMappings[stripeObjectMappings.length - 1].fields.push(fieldExpandableMap);
                        }
                    }
                }
            }
        } else if (stripeObjectToFormat[field]['anyOf']) {
            if (stripeObjectToFormat[field]['anyOf'].length && stripeObjectToFormat[field]['anyOf'][0]['$ref']) {
                var nestedExpandableField = stripeObjectToFormat[field]['anyOf'][0]['$ref'].split('/').pop();
                var nestedExpandableFieldMap = responseJson['components']['schemas'][nestedExpandableField]['properties'];
                if (nestedExpandableFieldMap) {
                    stripeObjectMappings = checkforNestedFields(field, stripeObjectToFormat, stripeObjectMappings, nestedExpandableFieldMap, objectExcludedReadOnlyFields);
                }
            }
        }
    }
    stripeObjectMappings = stripeObjectMappings.filter(function(section) {
        return section.fields.length;
    });
    return stripeObjectMappings;
}

function checkforNestedFields(field, stripeObjectToFormat, stripeObjectMappings, nestedExpandableFieldMap, objectExcludedReadOnlyFields) {
    const NEWSECTION = {
        label: field.charAt(0).toUpperCase() + field.slice(1).replace(/_+/g, ' ').replace(/\./g, ' '),
        name: field,
        description: '',
        fields: []
    };
    stripeObjectMappings.push(NEWSECTION);
    for (const expandableField in nestedExpandableFieldMap) {
        if (nestedExpandableFieldMap[expandableField]['description'] && !excludedReadOnlyFields.all.includes(expandableField) && !objectExcludedReadOnlyFields.includes(expandableField)) {
            if ((nestedExpandableFieldMap[expandableField]['type'] && nestedExpandableFieldMap[expandableField]['type'] !== 'object'
            && nestedExpandableFieldMap[expandableField]['type'] !== 'array')) {
                var hashFieldName = expandableField.replace(/_+/g, ' ');
                var hashFieldValue = field + '.' + expandableField;
                var fieldExpandableMap = {
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
                fieldExpandableMap['description'] = markdownConverter.makeHtml(nestedExpandableFieldMap[expandableField]['description']);
                fieldExpandableMap['type'] = nestedExpandableFieldMap[expandableField]['type'];
                stripeObjectMappings[stripeObjectMappings.length - 1].fields.sort(function(a, b) {
                    return a.name.localeCompare(b.name);
                });
                var index = stripeObjectMappings.findIndex(objectSection => {
                    return objectSection.name === field;
                });

                if (index) {
                    stripeObjectMappings[index].fields.push(fieldExpandableMap);
                } else {
                    stripeObjectMappings[stripeObjectMappings.length - 1].fields.push(fieldExpandableMap);
                }
            }
        } else {
            if (responseJson['components']['schemas'][expandableField] && responseJson['components']['schemas'][expandableField]['properties']
            && expandableField !== 'coupon'&& !excludedReadOnlyFields.all.includes(expandableField) && !objectExcludedReadOnlyFields.includes(expandableField)) {
                var newNestedExpandableFieldMap = responseJson['components']['schemas'][expandableField]['properties'];
                var newNestedFieldName = field + '.' + expandableField.charAt(0).toUpperCase() + expandableField.slice(1).replace(/_+/g, ' ');
                stripeObjectMappings = checkforNestedFields(newNestedFieldName, stripeObjectToFormat, stripeObjectMappings, newNestedExpandableFieldMap, objectExcludedReadOnlyFields);
            }
        }
    }
    return stripeObjectMappings;
}