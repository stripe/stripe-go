public with sharing class utilities {
    public class SetupIsNotCompleteException extends Exception {}

    @AuraEnabled
    public static String getPackageVersion() {
        responseData rd = new responseData();

        if(String.isEmpty(constants.NAMESPACE_API)) {
            rd.put('major', '0');
            rd.put('minor', '0');
            return rd.getJsonString();
        }

        try {
            rd.put('major', System.requestVersion().major());
            rd.put('minor', System.requestVersion().minor());
        } catch(Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    public static sObject getSobjectByMap(Map<String, Object> sourceMap, String sObjectName) {
        return getSobjectByMap(sourceMap, sObjectName, constants.NAMESPACE_API);
    }

    /*
        Takes in a map and converts it into the specified sObject type.
        This function will strip away any keys that don't match a field on the sObject
        so you don't have to worry about it being dirty.
    */
    public static sObject getSobjectByMap(Map<String, Object> sourceMap, String sObjectName, String namespace) {
        if (!Schema.getGlobalDescribe().containsKey(sObjectName)) {
            if (Schema.getGlobalDescribe().containsKey(namespace + sObjectName)) {
                sObjectName = namespace + sObjectName;
            } else {
                return null;
            }
        }

        Map<String, Schema.SObjectField> fieldMap = Schema.getGlobalDescribe().get(sObjectName).getDescribe().fields.getMap();
        Map<String, Object> objMap = new Map<String, Object>();

        for (String field : sourceMap.keySet()) {
            String value = String.valueOf(sourceMap.get(field));

            //only retain the value if it exists on our object
            if (!fieldMap.containsKey(field)) {
                if (fieldMap.containsKey(namespace + field)) {
                    field = namespace + field;
                } else {
                    continue;
                }
            }
            objMap.put(field, String.isEmpty(value) ? null : value);
        }
        return (sObject)JSON.deserialize(JSON.serialize(objMap), Type.forName(sObjectName));
    }

    /*gets us a map of all the fields and only gets lookup relationships that we care about*/
    public static Set<Object> getListOfFieldsByObject(String objectApiName) {
        Set<Object> objectFieldMapList = new Set<Object>();
        try {
            //check if we have access to the object
            if (Schema.describeSObjects(new List<String>{ objectApiName })[0].isAccessible()) {
                //get all fields on object
                Map<String, Schema.SObjectType> objectSchemaMap = Schema.getGlobalDescribe();
                Schema.SObjectType objectSchema = objectSchemaMap.get(objectApiName);
                Map<String, Schema.SObjectField> objectApiFieldMap = objectSchema.getDescribe().fields.getMap();
                Boolean isObjectCustom = (Boolean)objectSchema.getDescribe().isCustom();
                for (String objectFieldName : objectApiFieldMap.keySet()) {
                    Map<String, Object> objectFieldMap = new Map<String,Object>();
                    //check we have access to the field
                    if(objectApiFieldMap.get(objectFieldName).getDescribe().isAccessible()) {
                        //get field type
                        String typeOfField = String.valueOf(objectApiFieldMap.get(objectFieldName).getDescribe().getType()).toLowercase();
                        //check if field is a custom field or not
                        Boolean isFieldCustom = (Boolean)objectApiFieldMap.get(objectFieldName).getDescribe().isCustom();
                        //check if its a lookup field
                        if(typeOfField == 'reference') {
                            //get list of objects that the field looks up to
                            List<Object> objectsFieldReferences = objectApiFieldMap.get(objectFieldName).getDescribe().getReferenceTo();
                            //iterate over objects the field looks up to and only add to list if its a field we care about
                            for(Object referencedObject : objectsFieldReferences) {
                                //make sure the field only looks up to 1 object before adding it to map
                                if(objectsFieldReferences.size() == 1) {
                                    String stringReferncedObjecect = String.valueOf(referencedObject);
                                    String relationshipName = objectApiFieldMap.get(objectFieldName).getDescribe().getRelationshipName();

                                    //use this list to target lookup fields based on the object the look up to
                                    List<String> listOfImportantObjectsApiNames = new List<String> {
                                        'Product2',
                                        'Contract',
                                        'Order',
                                        'OrderItem',
                                        'PricebookEntry',
                                        'SBQQ__QuoteLine__c',
                                        'Account',
                                        'Contact',
                                        'SBQQ__Subscription__c',
                                        'SBQQ__Quote__c',
                                        'Opportunity',
                                        'Pricebook2',
                                        'ContractLineItem',
                                        'OpportunityLineItem',
                                        'ProductItem'
                                    };
                                    if(listOfImportantObjectsApiNames.contains(stringReferncedObjecect)) {
                                        objectFieldMap = addFieldsToObjectFieldMap(objectFieldMap, objectApiFieldMap, objectFieldName, isFieldCustom, typeOfField, isObjectCustom);
                                        objectFieldMap.put('object',stringReferncedObjecect);
                                        objectFieldMap.put('relationshipName',relationshipName);
                                    }
                                } else {
                                    //in this case we now the field looks up to more than one object and we will not support traversals
                                    objectFieldMap = addFieldsToObjectFieldMap(objectFieldMap, objectApiFieldMap, objectFieldName, isFieldCustom, 'polymorphic', isObjectCustom);
                                }
                            }
                        } else {
                            //in this case we know the field is not a look up so we add to field map
                            objectFieldMap = addFieldsToObjectFieldMap(objectFieldMap, objectApiFieldMap, objectFieldName, isFieldCustom, typeOfField, isObjectCustom);
                        }
                    } else {
                        //in the case the field is not accessable we will move on to the next field
                        continue;
                    }
                    //add field map to list to be returned
                    if(objectFieldMap != null && !objectFieldMap.isEmpty()) {
                        objectFieldMapList.add(objectFieldMap);
                    }
                }
            }
            objectFieldMapList.remove(null);
        } catch (Exception e) {
            errorLogger.create('getListOfFieldsByObject', e);
        }
        return objectFieldMapList;
    }

    private static Map<String, Object> addFieldsToObjectFieldMap(Map<String, Object> objectFieldMap, Map<String, Schema.SObjectField> objectApiFieldMap, String objectFieldName, Boolean isFieldCustom, String typeOfField, Boolean isObjectCustom) {
        objectFieldMap.put('label',objectApiFieldMap.get(objectFieldName).getDescribe().getLabel());
        //REST API is case sensitive and requires CamelCase field names
        objectFieldMap.put('value',objectApiFieldMap.get(objectFieldName).getDescribe().getName());
        objectFieldMap.put('type',typeOfField);
        objectFieldMap.put('isFieldCustom',isFieldCustom);
        objectFieldMap.put('isObjectCustom',isObjectCustom);
        return objectFieldMap;
    }

    public static Boolean isCpqEnabled() {
        //checks if the namespace of the CPQ package exists (all CPQ package namespaces start with 'SBQQ'
        //if a non-admin installs this package, this check would fail
        //we are using a wildcard because there are variants of the CPQ package which have a unique namespace (SBQQW)
        Integer sbqqPackageCount = [SELECT COUNT() FROM PackageLicense WHERE NamespacePrefix LIKE 'SBQQ%'];
        return sbqqPackageCount > 0;
    }

    /* gets the organization type (sandbox or production) we are using a custom setting object so this record
    will not be copied over to newly generated sandboxes instead a new record will be created for each instance*/
    @AuraEnabled
    public static String setOrgType() {
        responseData rd = new responseData();
        try {
            Organization_Type__c orgTypeSetting = Organization_Type__c.getOrgDefaults();
            //only set value if the record doesn't exist yet
            if (orgTypeSetting.Id == null) {
                orgTypeSetting.isSandbox__c = (Boolean)constants.IS_SANDBOX;
                upsert orgTypeSetting;
            }
            /*sets custom setting record that controls weather sending errors to sentry is live or not*/
            Sentry_Active_Config__c activateSentrySetting = Sentry_Active_Config__c.getOrgDefaults();
            if (activateSentrySetting.Id == null) {
                /* this would be useful if you plan to log in to subscriber orgs via the PBO/LMA for debugging. This flag sets the 'sentry log' statements to be live so you can see exception details clearly in the debug logs. The details on what each of those custom setting fields control is in the readMe file inside the source code repo btw (https://github.com/jmather/SentryForSalesforce). */
                activateSentrySetting.IsDebug__c = true;
                activateSentrySetting.Environment_Name__c = (String)constants.ORG_ID;
                activateSentrySetting.IsIssueCreationDisabled__c = false;
                activateSentrySetting.Sentry_Config__c = 'Default';
                upsert activateSentrySetting;
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    public static List<Setup_Connection_Data__mdt> getStripeConnectionKey() {
        return [SELECT Global_Key__c
                FROM Setup_Connection_Data__mdt
                LIMIT 1];
    }

    // generates package level key and sends to ruby. Named `salesforce_organization_key` on the ruby side
    public static void generatePackageKey(Stripe_Connection__c stripeConnectRec) {

        List<Setup_Connection_Data__mdt> setupConfigList = getStripeConnectionKey();
        if(setupConfigList.isEmpty()) {
            /* This error should only be thrown in a scratch org when unmangaed key is deployed to
            an org and the custom metadata record containing the global key was not created */
            Sentry_Log.logSentry('missing metadata record with global key');
            throw new SetupIsNotCompleteException('Org is not setup correctly, missing custom metadata record with global key');
        }

        Setup_Connection_Data__mdt setupData = setupConfigList[0];

        Map<String,String> headers = new Map<String,String> {
            'Salesforce-Key' => (String)setupData.Global_Key__c
        };

        // generate unique key for this organization used to authenticate against the integration service
        Blob blb = Crypto.GenerateAESKey(128);
        String packageKey  = EncodingUtil.ConvertTohex(blb);

        String route = constants.RUBY_SERVICE_BASE_URI + '/v1/post-install';

        Map<String, Object> body = new Map<String, Object> {
            'key' => packageKey
        };

        HttpResponse response = utilities.makeCallout(route, 'POST', JSON.serialize(body), headers);

        Map<String, Object> errorBody;
        if(response.getStatusCode() == 200) {
            stripeConnectRec.API_Key__c = packageKey;
            upsert stripeConnectRec;

            // schedules the batch job that will clean up sync records
            scheduleSyncRecordDeletion.scheduleJob();
        } else {
            errorLogger.create('generatePackageKey', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get package level API key from ruby service.');
        }
    }

    public static HttpResponse makeCallout(String endPoint, String method){
        return makeCallout(endPoint, method, null, null);
    }

    public static HttpResponse makeCallout(String endPoint, String method, String body){
        return makeCallout(endPoint, method, body, null);
    }

    public static HttpResponse makeCallout(String endPoint, String method, Map<String, String> headers){
        return makeCallout(endPoint, method, null, headers);
    }

    public static HttpResponse makeCallout(String endPoint, String method, String body, Map<String, String> headers){
        Http httpMethod = new Http();
        HttpRequest request = new HttpRequest();

        Integer requestTimeoutLength = 500;
        request.setEndpoint(endPoint);
        request.setMethod(method);
        request.setTimeout(requestTimeoutLength);

        if(body != null) {
            //these headers are required for the API to respond properly
            request.setHeader('Content-Type', 'application/json');
            request.setHeader('Accept', 'application/json');
            request.setBody(body);
        }

        /* these standard headers will be add to the request if none are provided
        the ruby service expects these headers for all requests. */
        request.setHeader('Salesforce-Account-Id', (String)constants.ORG_ID);
        request.setHeader('Salesforce-Type', String.valueOf(Sentry_Environment.getInstanceType()));
        request.setHeader('Salesforce-Package-Namespace', constants.NAMESPACE);

        if(!String.isEmpty(constants.NAMESPACE_API)) {
            request.setHeader('Package-Id', (String)getPackageVersion());
        }

        /*
         this only happens in the case the key is passed in after generating
         the package level api key, otherwise we will use the organization-specific api key
         */
        if(headers != null) {
            request.setHeader('Salesforce-Key', headers.get('Salesforce-Key'));
            HttpResponse response = httpMethod.send(request);
            return response;
        }

        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

        if(stripeConnectRec.Id == null || stripeConnectRec.API_Key__c == null) {
            /*
            This error should only be thrown in a scratch org when unmangaed key is deployed to
            an org and the custom metadata record containing the global key was not created
            */
            Sentry_Log.logSentry('missing Stripe Connection org default record with api key');
            throw new SetupIsNotCompleteException('Connector is not setup correctly, missing metadata.');
        }

        request.setHeader('Salesforce-Key', (String)stripeConnectRec.API_Key__c);

        HttpResponse response = httpMethod.send(request);
        return response;
    }

    //Creates a test metadata record in test context
    @testVisible static List<Setup_Connection_Data__mdt> setupConfigMetadata {
        get {
            if (setupConfigMetadata == null)
                setupConfigMetadata = getStripeConnectionKey();
            return setupConfigMetadata;
        } set; }

}
