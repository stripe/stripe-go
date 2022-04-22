const HTTPS = require('https')
const showdown = require('showdown');
const markdownConverter = new showdown.Converter();
const excludedReadOnlyFields = {
    all: [
        'created',
        'object',
        'livemode',
        'metadata'
    ],

    customer: [
        'delinquent',
        'discount',
    ],

    product: [
        'price'
    ],

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
    path: 'stripe/openapi/d41298323ca953bf0b475b685a63515aed1e0c73/openapi/spec3.json',
    method: 'GET'
}

let openApiSpec = '';

function extractStripeObject(openApiSpec, stripeObjectType) {
    return openApiSpec['paths'][`/v1/${stripeObjectType}s`]['post']['requestBody']['content']['application/x-www-form-urlencoded']['schema']['properties']
}

// TODO we should use async/await here instead
const HTTPREQUEST = HTTPS.request(OPTIONS, HttpResponse => {
    HttpResponse.on('data', responseDataChunk => {
        openApiSpec += responseDataChunk.toString();
    })

    HttpResponse.on('end', () => {
        openApiSpec = JSON.parse(openApiSpec);

        // TODO reduce duplication here and just have a single array of mappings we support
        var formattedStripeObjectsForMapper = {
            formattedStripeCustomerFields: formatStripeObjectsForMapper(extractStripeObject(openApiSpec, 'customer'), excludedReadOnlyFields.customer),
            formattedStripeProductItemFields: formatStripeObjectsForMapper(extractStripeObject(openApiSpec, 'product'), excludedReadOnlyFields.product),
            formattedStripeSubscriptionFields: formatStripeObjectsForMapper(extractStripeObject(openApiSpec, 'subscription_schedule'), excludedReadOnlyFields.subscription),
            formattedStripeSubscriptionItemFields: formatStripeObjectsForMapper(extractStripeObject(openApiSpec, 'subscription_item'), excludedReadOnlyFields.subscription_item),
            formattedStripePriceFields: formatStripeObjectsForMapper(extractStripeObject(openApiSpec, 'price'), excludedReadOnlyFields.price)
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
        } else if (stripeObjectToFormat[field]['$ref'] && !excludedReadOnlyFields.all.includes(field) && !objectExcludedReadOnlyFields.includes(field)) {
            var expandableSchemaFieldName = stripeObjectToFormat[field]['$ref'].split('/').pop();
            var expandableSchemaFieldMap = openApiSpec['components']['schemas'][expandableSchemaFieldName]['properties'];
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
        } else if (stripeObjectToFormat[field]['anyOf'] && openApiSpec['components']['schemas'][field] && !excludedReadOnlyFields.all.includes(field) && !objectExcludedReadOnlyFields.includes(field)) {
            if (stripeObjectToFormat[field]['anyOf'].length) {
                var nestedExpandableFieldMap = openApiSpec['components']['schemas'][field]['properties'];
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
            if (openApiSpec['components']['schemas'][expandableField] && openApiSpec['components']['schemas'][expandableField]['properties']
            && expandableField !== 'coupon'&& !excludedReadOnlyFields.all.includes(expandableField) && !objectExcludedReadOnlyFields.includes(expandableField)) {
                var newNestedExpandableFieldMap = openApiSpec['components']['schemas'][expandableField]['properties'];
                var newNestedFieldName = field + '.' + expandableField.charAt(0) + expandableField.slice(1).replace(/_+/g, ' ');
                stripeObjectMappings = checkforNestedFields(newNestedFieldName, stripeObjectToFormat, stripeObjectMappings, newNestedExpandableFieldMap, objectExcludedReadOnlyFields);
            }
        }
    }
    return stripeObjectMappings;
}
