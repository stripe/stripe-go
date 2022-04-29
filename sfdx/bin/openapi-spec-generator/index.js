const HTTPS = require('https')
const showdown = require('showdown');
const markdownConverter = new showdown.Converter();

const excludedFields = {
    all: [
        'object',
        // metadata is special cased in the data mapper
        'metadata',
        'expand'
    ],

    customer: [],

    product: [
        'price'
    ],

    subscription_schedule: [
        // the customer relationship is managed directly by the integration, omit in the UI
        'customer'
    ],

    subscription_item: [],

    price: [
        // product<>price relationship is managed directly by the integration, we should ignore
        'product'
    ],
}

const ListOfStripeObjects = [
    'customer',
    'product',
    'subscription_schedule',
    'subscription_item',
    'price',
]

const OPTIONS = {
    hostname: 'raw.githubusercontent.com',
    port: 443,
    path: 'stripe/openapi/d41298323ca953bf0b475b685a63515aed1e0c73/openapi/spec3.json',
    method: 'GET'
}

let openApiSpec = '';

function extractStripeObject(openApiSpec, stripeObjectType) {
    return openApiSpec['paths'][`/v1/${stripeObjectType}s`]['post']['requestBody']['content']['application/x-www-form-urlencoded']['schema']['properties'];
}

// TODO we should use async/await here instead
const HTTPREQUEST = HTTPS.request(OPTIONS, HttpResponse => {
    HttpResponse.on('data', responseDataChunk => {
        openApiSpec += responseDataChunk.toString();
    })

    HttpResponse.on('end', () => {
        openApiSpec = JSON.parse(openApiSpec);


        var formattedStripeObjectsForMapper = {};
        for (const stripeObject of ListOfStripeObjects) {
            var convertObjectName = stripeObject.charAt(0).toUpperCase() + stripeObject.slice(1);

            // TODO change mapper response to match here so there is now need for all this formatting https://github.com/stripe/stripe-salesforce/issues/364
            // NOTE when this naming change is made to the mapper it will require all to upgrade package or the mapper will not work
            switch (stripeObject) {
                case 'subscription_schedule':
                    convertObjectName = 'Subscription'
                    break;
                case 'subscription_item':
                    convertObjectName = 'SubscriptionItem'
                    break;
                case 'product':
                    convertObjectName = 'ProductItem'
                    break;
                default:
                    break;
            }

            formattedStripeObjectsForMapper['formattedStripe' + convertObjectName + 'Fields'] = formatStripeObjectsForMapper(extractStripeObject(openApiSpec, stripeObject), excludedFields[stripeObject]);

        }
        formattedStripeObjectsForMapper = JSON.stringify(formattedStripeObjectsForMapper);
        console.log(formattedStripeObjectsForMapper);
        return formattedStripeObjectsForMapper;
    })
})

HTTPREQUEST.on('error', error => {
    console.log(error);
})

HTTPREQUEST.end();

function arrayEquals(a, b) {
    return Array.isArray(a) &&
        Array.isArray(b) &&
        a.length === b.length &&
        a.every((val, index) => val === b[index]);
}


function formatStripeObjectsForMapper(stripeObjectToFormat, objectExcludedReadOnlyFields) {
    let stripeObjectMappings = [{
        label: 'Standard Mappings',
        name: 'standard',
        description: '',
        fields: []
    }];

    for (const field in stripeObjectToFormat) {
        if (excludedFields.all.includes(field) || objectExcludedReadOnlyFields.includes(field)) {
            continue;
        }

        const fieldData = stripeObjectToFormat[field];

        // if `type` does not exist, `anyOf` is often provided which provides a list of options
        // in many cases, a enum with an empty string is an option. I'm not sure why, but I believe we can safely filter these out and treat
        // them as a type with anyOf containing the other options
        const acceptableTypes = (fieldData['anyOf'] || []).filter(t => !arrayEquals(t.enum, [''])).map(t => t.type)

        // mapper does not support array mapping
        if(fieldData['type'] === 'array' || arrayEquals(acceptableTypes, ['array'])) {
            continue;
        }

        // if we don't have an object reference, then this is a standard field
        if (
            !fieldData['$ref'] &&
            !acceptableTypes.includes('object') &&

            (fieldData['type'] || '') !== 'object'
        ) {
            const fieldMap = getNewFieldObject(field.replace(/_+/g, ' '), field);
            fieldMap['type'] = fieldData['type'];

            if (fieldData['description']) {
                fieldMap['description'] = markdownConverter.makeHtml(fieldData['description']);
            }

            // standard field section is always at the top of the array
            stripeObjectMappings[0].fields.push(fieldMap);
            continue;
        }

        // TODO we do not have the logic below adjusted for the new understanding that `acceptableTypes` could contain *just* an object reference

        if (fieldData['anyOf'] && fieldData['anyOf'].length &&
            openApiSpec['components']['schemas'][field] ) {
            var nestedExpandableFieldMap = openApiSpec['components']['schemas'][field]['properties'];
            if (nestedExpandableFieldMap) {
                stripeObjectMappings = checkforNestedFields(field, stripeObjectToFormat, stripeObjectMappings, nestedExpandableFieldMap, objectExcludedReadOnlyFields);
            }
        }

        var expandableSchemaFieldName;
        var expandableSchemaFieldMap;

        /*In this case we are getting the field name assoicated with this object and going to
        '['components']['schemas']' tree path associated with the current field to get the
        related subfields, descriptions and types to add to them mapper*/
        // TODO in what cases does this occur?
        if (fieldData['$ref']) {
            expandableSchemaFieldName = fieldData['$ref'].split('/').pop();
            expandableSchemaFieldMap = openApiSpec['components']['schemas'][expandableSchemaFieldName]['properties'];
        }

        //In this case we are getting all the descriptions, subfields and types from the object directly
        if (fieldData['type'] === 'object' || fieldData['properties']) {
            expandableSchemaFieldName = field;
            expandableSchemaFieldMap = fieldData['properties'];
        }

        // TODO this should be separated out into it's own method
        // TODO the description of the fields below is NOT included, it needs to be looked up elsewhere in the schema https://github.com/stripe/stripe-salesforce/issues/376

        for (const expandableField in expandableSchemaFieldMap) {
            if (excludedFields.all.includes(expandableField) || objectExcludedReadOnlyFields.includes(expandableField)) {
                continue
            }

            if (expandableSchemaFieldMap[expandableField] && expandableSchemaFieldMap[expandableField]['type'] && expandableSchemaFieldMap[expandableField]['type'] !== 'object') {
                var newSection = {
                    label: field.charAt(0).toUpperCase() + field.slice(1).replace(/_+/g, ' '),
                    name: field,
                    description: '',
                    fields: []
                };
                stripeObjectMappings.push(newSection);
                const hashFieldName = expandableField.replace(/_+/g, ' ');
                const hashFieldValue = field + '.' + expandableField;
                stripeObjectMappings = addNewFieldToSection(stripeObjectMappings, hashFieldName, hashFieldValue, expandableSchemaFieldMap, expandableField, field);

            } else if (fieldData['properties'] && fieldData['properties'][expandableField]['properties']) {
                expandableSchemaFieldMap = fieldData['properties'][expandableField]['properties'];

                for (const subfield in expandableSchemaFieldMap) {
                    var newSection = {
                        label: expandableField.charAt(0).toUpperCase() + expandableField.slice(1).replace(/_+/g, ' ')+ ' ' +subfield.charAt(0).toUpperCase() + subfield.slice(1).replace(/_+/g, ' '),
                        name: [field, expandableField, subfield].join("."),
                        description: '',
                        fields: []
                    };
                    stripeObjectMappings.push(newSection);
                    var nestedHashFieldName = subfield.replace(/_+/g, ' ');
                    var nestedHashFieldValue = field + '.' + expandableField + '.' + subfield;
                    stripeObjectMappings = addNewFieldToSection(stripeObjectMappings, nestedHashFieldName, nestedHashFieldValue, expandableSchemaFieldMap, subfield, field);

                }
            }
        }
    }

    stripeObjectMappings = stripeObjectMappings.filter(function(section) {
        return section.fields.length > 0;
    });

    // sort all fields alphabetically
    stripeObjectMappings[0].fields.sort((a, b) => a.name.localeCompare(b.name));

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

        if (excludedFields.all.includes(expandableField) || objectExcludedReadOnlyFields.includes(expandableField)) {
            continue;
        }

        if ((nestedExpandableFieldMap[expandableField]['type'] && nestedExpandableFieldMap[expandableField]['type'] !== 'object'
        && nestedExpandableFieldMap[expandableField]['type'] !== 'array' && nestedExpandableFieldMap[expandableField]['description'])) {
            var hashFieldName = expandableField.replace(/_+/g, ' ');
            var hashFieldValue = field + '.' + expandableField;
            let fieldExpandableMap = getNewFieldObject(hashFieldName, hashFieldValue);
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
        } else {
            if (openApiSpec['components']['schemas'][expandableField] && openApiSpec['components']['schemas'][expandableField]['properties']) {
                var newNestedExpandableFieldMap = openApiSpec['components']['schemas'][expandableField]['properties'];
                var newNestedFieldName = field + '.' + expandableField.charAt(0) + expandableField.slice(1).replace(/_+/g, ' ');
                stripeObjectMappings = checkforNestedFields(newNestedFieldName, stripeObjectToFormat, stripeObjectMappings, newNestedExpandableFieldMap, objectExcludedReadOnlyFields);
            }
        }
    }
    return stripeObjectMappings;
}

function addNewFieldToSection(stripeObjectMappings, hashFieldName, hashFieldValue, expandableSchemaFieldMap, expandableField, field) {
    let fieldExpandableMap = getNewFieldObject(hashFieldName, hashFieldValue);
    fieldExpandableMap['description'] = markdownConverter.makeHtml(expandableSchemaFieldMap[expandableField]['description']);
    fieldExpandableMap['type'] = expandableSchemaFieldMap[expandableField]['type'];
    var index = stripeObjectMappings.findIndex(objectSection => {
        return objectSection.name === field;
    });
    if (index) {
        stripeObjectMappings[index].fields.push(fieldExpandableMap);
    } else {
        stripeObjectMappings[stripeObjectMappings.length - 1].fields.push(fieldExpandableMap);
    }
    return stripeObjectMappings;
}

function getNewFieldObject(fieldName, fieldValue) {
    const fieldMap = {
        name: fieldName,
        value: fieldValue,
        description: '',
        type: '',
        defaultValue: '',
        hasOverride: false,
        staticValue: false,
        hasSfValue: false,
        sfValue: '',
        sfValueType: ''
    };
    return fieldMap;
}