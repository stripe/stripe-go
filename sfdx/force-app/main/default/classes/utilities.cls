public with sharing class utilities {
    // generate with: System.debug(EncodingUtil.to(Crypto.generateAesKey(256)));
    private static String stateEncryptionKey = '98ae2216688a8b879c25f17ed793bfa7bb11201ddda97ad01b2d282a06aa191a';
    private static String stateSigningKey = '98ae2216688a8b879c25f17ed793bfa7bb11201ddda97ad01b2d282a06aa191a';

    @TestVisible
    static ConfigPayloadHelper configPayloadHelper = new ConfigPayloadHelper();

    @TestVisible
    static Debug_Helper debugger = new Debug_Helper();

    public class SetupIsNotCompleteException extends Exception {}
    public class HttpCalloutFailedException extends Exception {}

    public static Boolean isPackagedEnvironment() {
        return !String.isEmpty(constants.NAMESPACE_API);
    }

    public static Boolean isConnected(Stripe_Connection__c stripeConnectRec) {
        return stripeConnectRec != null && stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c;
    }

    public static List<Map<String, String>> getMultiCurrencyOptions() {
        List<Map<String, String>> supportedISOCodes = new List<Map<String, String>>();
        Boolean multiCurrencyEnabled = UserInfo.isMultiCurrencyOrganization();
        if (multiCurrencyEnabled) {
            supportedISOCodes.add(new Map<String, String> {
                'label' => 'Select default currency...',
                'value' => ''
            });

            /* using a dynamic query here because the currencyType object can only be referenced directly
            when the multicurrency feature is enabled in the org otherwise we get a type error*/
            String supportedISOCodesQuery = 'SELECT ISOCode FROM CurrencyType WHERE IsActive = true';
            List<SObject> supportedISOCodesQueryResults = Database.query(supportedISOCodesQuery);

            //Iterate through the supported currencies to build list of supported ISO code
            for(SObject supportedISOCode : supportedISOCodesQueryResults) {
                supportedISOCodes.add(new Map<String, String> {
                        'label' => (String) supportedISOCode.get('ISOCode'),
                        'value' => (String) supportedISOCode.get('ISOCode')
                });
            }
        }

        return supportedISOCodes;
    }

    @AuraEnabled
    public static String getPackageVersion() {
        responseData rd = new responseData();

        if(!isPackagedEnvironment()) {
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

    public static String getPackageVersionString() {

        if(!isPackagedEnvironment()) {
            return '0.0.0';
        }

        String major = String.valueOf(System.requestVersion().major());
        String minor = String.valueOf(System.requestVersion().minor());
        String patch = String.valueOf(System.requestVersion().patch());

        if (patch == null) {
            patch = '0';
        }

        return major + '.' + minor + '.' + patch;
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
                                        'ProductItem',
                                        'Quote_Stripe_Coupon__c'
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
        objectFieldMap.put('label', objectApiFieldMap.get(objectFieldName).getDescribe().getLabel());
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
        Integer sbqqPackageCount = Database.countQuery('SELECT COUNT() FROM PackageLicense WHERE NamespacePrefix LIKE \'SBQQ%\'');
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

    public static String getStripeConnectionKey() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            return setupConfigMetadata[0].Global_Key__c;
        }

        if (Test.isRunningTest() == false) {
            Sentry_Log.logSentry('missing metadata record with global key');
        }
        throw new SetupIsNotCompleteException('Org is not setup correctly, missing custom metadata record with global key');
    }

    public static Boolean shouldInitializeMiddlewareConnection() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            debugger.debug('utilities.shouldInitializeMiddlewareConnection', 'Initializing middleware connection');
            Boolean isV1 = setupConfigMetadata[0].Middleware_Infrastructure_Version__c == 'v1';
            Boolean isV2 = !isV1;
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
            Stripe_Account__c acct = getPrimaryStripeAccount();
            Boolean needsAPIKey = stripeConnectRec.API_Key__c == null || String.isBlank(stripeConnectRec.API_Key__c);
            debugger.debug('utilities.shouldInitializeMiddlewareConnection:needsAPIKey', needsAPIKey);
            debugger.debug('utilities.shouldInitializeMiddlewareConnection:acct', acct);
            debugger.debug('utilities.shouldInitializeMiddlewareConnection:isV2', isV2);
            if ((isV1 && needsAPIKey) || (isV2 && needsAPIKey && acct != null)) {
                return true;
            } else if (isV2) {
                return false;
            }
        }

        return false;
    }

    public static Boolean isUsingV1Middleware() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            return setupConfigMetadata[0].Middleware_Infrastructure_Version__c == 'v1';
        }
        return false;
    }

    @AuraEnabled(Cacheable=false)
    public static Boolean isUsingV2Middleware() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            return setupConfigMetadata[0].Middleware_Infrastructure_Version__c != 'v1';
        }
        return true;
    }

    public static Blob getOAuthStateSigningKey() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            String encKey = setupConfigMetadata[0].OAuth_State_Signing_Key__c;
            if (String.isBlank(encKey) || encKey == null) {
                encKey = stateSigningKey;
            }
            return EncodingUtil.convertFromHex(encKey);
        }

        return null;
    }

    public static Boolean isDeveloperModeEnabled() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            return setupConfigMetadata[0].Enable_JS_Debug__c || setupConfigMetadata[0].Enable_Apex_Debug__c || setupConfigMetadata[0].Local_Authorization_Target__c;
        }

        return false;
    }

    public static Boolean isApexDebuggingEnabled() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            return setupConfigMetadata[0].Enable_Apex_Debug__c;
        }

        return false;
    }

    public static Boolean isJSDebuggingEnabled() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            return setupConfigMetadata[0].Enable_JS_Debug__c;
        }

        return false;
    }

    public static Boolean isOAuthAutoCloseEnabled() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            return setupConfigMetadata[0].OAuth_AutoClose_Window__c;
        }

        return false;
    }

    public static Stripe_Connection__c getStripeConnection() {
        return Stripe_Connection__c.getOrgDefaults();
    }

    public static Integer getSyncRecordRetentionDays() {
        Stripe_Connection__c config = getStripeConnection();
        debugger.debug('utilities.getSyncRecordRetentionDays:Sync_Record_Retention__c', JSON.serialize(config.Sync_Record_Retention__c));

        if (String.isNotEmpty(config.Sync_Record_Retention__c)) {
            debugger.debug('utilities.getSyncRecordRetentionDays:value', JSON.serialize(Integer.valueOf(config.Sync_Record_Retention__c)));
            return Integer.valueOf(config.Sync_Record_Retention__c);
        }

        return 5000;
    }

    public static Setup_Connection_Data__mdt getConnectionData() {
        if (setupConfigMetadata != null && setupConfigMetadata.isEmpty() == false) {
            return setupConfigMetadata[0];
        }

        return new Setup_Connection_Data__mdt();
    }

    public static List<Setup_Connection_Data__mdt> getStripeConnectionData() {
        return setupConfigMetadata;
    }

    // generates package level key and sends to ruby. Named `salesforce_organization_key` on the ruby side
    public static void generatePackageKey(Stripe_Connection__c stripeConnectRec) {
        List<Setup_Connection_Data__mdt> setupConfigList = setupConfigMetadata;
        debugger.debug('utilities.generatePackageKey', JSON.serialize(stripeConnectRec));

        if(setupConfigList.isEmpty()) {
            /* This error should only be thrown in a scratch org when unmangaed key is deployed to
            an org and the custom metadata record containing the global key was not created */
            Sentry_Log.logSentry('missing metadata record with global key');
            throw new SetupIsNotCompleteException('Org is not setup correctly, missing custom metadata record with global key');
        }

        Setup_Connection_Data__mdt setupData = setupConfigList[0];
        String authorization = 'Bearer ' + StripeOAuthState.create().toString();

        Map<String,String> headers = new Map<String,String> {
            'Salesforce-Key' => (String)setupData.Global_Key__c,
            'Authorization' => authorization
        };

        // generate unique key for this organization used to authenticate against the integration service
        Blob blb = Crypto.generateAesKey(128);
        String packageKey  = EncodingUtil.convertToHex(blb);

        String route = UriTargets.getPostInstallEndpoint();

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

    public static HttpResponse makeCallout(String endPoint, String method, Integer timeoutLength){
        return makeCallout(endPoint, method, null, null, timeoutLength);
    }

    public static HttpResponse makeCallout(String endPoint, String method, String body){
        return makeCallout(endPoint, method, body, null);
    }

    public static HttpResponse makeCallout(String endPoint, String method, Map<String, String> headers){
        return makeCallout(endPoint, method, null, headers);
    }

    public static HttpResponse makeCallout(String endPoint, String method, String body, Map<String, String> headers) {
        return makeCallout(endPoint, method, body, headers, 5000);
    }

    public static HttpResponse makeCallout(String endPoint, String method, String body, Map<String, String> headers, Integer timeoutLength){
        Http httpMethod = new Http();
        HttpRequest request = buildRequest(endPoint, method, body, headers, timeoutLength);
        debugger.debug('utilities.makeCallout', 'Calling endpoint: ' + endPoint);
        HttpResponse response = httpMethod.send(request);
        debugger.debug('utilities.makeCallout', 'Response status: ' + response.getStatusCode());
        if (response.getStatusCode() > 299) {
            debugger.debug('utilities.makeCallout', 'Response body: ' + response.getBody());
            for (String header : response.getHeaderKeys()) {
                debugger.debug('utilities.makeCallout', 'Response header: ' + header + ' : ' + response.getHeader(header));
            }
        }

        return response;
    }
    public static HttpRequest buildRequest(String endPoint, String method, Integer timeoutLength) {
        return buildRequest(endPoint, method, null, null, timeoutLength);
    }
    public static HttpRequest buildRequest(String endPoint, String method, String body, Map<String, String> headers, Integer timeoutLength){
        debugger.debug('utilities.buildRequest', 'endpoint: ' + endPoint);
        debugger.debug('utilities.buildRequest', 'method: ' + method);
        debugger.debug('utilities.buildRequest', 'body: ' + body);
        debugger.debug('utilities.buildRequest', 'headers: ' + headers);

        HttpRequest request = new HttpRequest();

        // 5s timeout in case the ruby server is backed up
        Integer requestTimeoutLength = 5000;
        request.setEndpoint(endPoint);
        request.setMethod(method);
        request.setTimeout(requestTimeoutLength);

        if (body != null) {
            //these headers are required for the API to respond properly
            request.setHeader('Content-Type', 'application/json');
            request.setHeader('Accept', 'application/json');
            request.setBody(body);
            debugger.debug('utilities.buildRequest', 'Set body with default body headers');
        }

        /* these standard headers will be add to the request if none are provided
        the ruby service expects these headers for all requests. */
        request.setHeader('Salesforce-Account-Id', (String)constants.ORG_ID);
        debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Account-Id" to: ' + constants.ORG_ID);
        request.setHeader('Salesforce-Type', String.valueOf(Sentry_Environment.getInstanceType()));
        debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Type" to: ' + Sentry_Environment.getInstanceType());
        request.setHeader('Salesforce-Package-Namespace', constants.NAMESPACE);
        debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Package-Namespace" to: ' + constants.NAMESPACE);

        request.setHeader('X-Requested-With', 'Salesforce');
        debugger.debug('utilities.buildRequest', 'Set header "X-Requested-With" to: Salesforce');

        // TODO this does not seem to be coming over properly
        if (utilities.isPackagedEnvironment()) {
            request.setHeader('Salesforce-Package-Id', (String)getPackageVersionString());
            debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Package-Id" to: ' + getPackageVersionString());
        } else {
            // For non-package environments, lets just set it as version 0
            request.setHeader('Salesforce-Package-Id', '0.0.0.0');
            debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Package-Id" to: 0.0.0.0');
        }

        debugger.debug('utilities.buildRequest', 'Set default headers');

        Stripe_Account__c acct = getPrimaryStripeAccount();
        if (acct != null) {
            debugger.debug('utilities.buildRequest', 'Using V2');
            
            // when we are deleting or setting default account, we need to use the header passed in
            // otherwise we default to the primary acct
            String acctId = headers != null && headers.get('Stripe-Primary-Account-Id') != null ? headers.get('Stripe-Primary-Account-Id') : acct.Stripe_ID__c;
            request.setHeader('Stripe-Primary-Account-Id', acctId);
            debugger.debug('utilities.buildRequest', 'Set header "Stripe-Primary-Account-Id" to: ' + acctId);
            
            String acctLivemode = headers != null ? headers.get('Stripe-Primary-Account-Livemode') : null;
            if (acctLivemode == null) {
                acctLivemode = acct.Is_Live_Mode__c ? 'true' : 'false';
            }
            request.setHeader('Stripe-Primary-Account-Livemode', acctLivemode);
            debugger.debug('utilities.buildRequest', 'Set header "Stripe-Primary-Account-Livemode" to: ' + acctLivemode);
        }

        /*
         this only happens in the case the key is passed in after generating
         the package level api key, otherwise we will use the organization-specific api key
         */
        if (headers != null && headers.get('Salesforce-Key') != null) {
            debugger.debug('utilities.buildRequest', 'Set custom headers');
            request.setHeader('Salesforce-Key', headers.get('Salesforce-Key'));
            debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Key" to: ' + redactString(headers.get('Salesforce-Key')));
            return request;
        }

        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
        Boolean needsSetup = stripeConnectRec.Id == null || stripeConnectRec.API_Key__c == null;

        if (needsSetup && isUsingV1Middleware()) {
            debugger.debug('utilities.buildRequest', 'Missing Stripe_Connection__c.Id or API_Key__c');
            /*
            This error should only be thrown in a scratch org when unmangaed key is deployed to
            an org and the custom metadata record containing the global key was not created
            */
            Sentry_Log.logSentry('Missing Stripe Connection org default record with api key');
            throw new SetupIsNotCompleteException('Connector is not setup correctly, missing metadata.');
        }

        if (isUsingV1Middleware()) {
            debugger.debug('utilities.buildRequest', 'Setting default API key header');
            request.setHeader('Salesforce-Key', (String) stripeConnectRec.API_Key__c);
            debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Key" to: ' + redactString(stripeConnectRec.API_Key__c));
        } else if (isUsingV2Middleware()) {
            debugger.debug('utilities.buildRequest', 'Setting default API key header');
            if (stripeConnectRec.API_Key__c == null) {
                request.setHeader('Salesforce-Key', '');
                debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Key" to: [REDACTED - EMPTY STRING]');
            } else {
                request.setHeader('Salesforce-Key', stripeConnectRec.API_Key__c);
                debugger.debug('utilities.buildRequest', 'Set header "Salesforce-Key" to: ' + redactString(stripeConnectRec.API_Key__c));
            }
        }
        request.setHeader('X-Middleware-Authorization', StripeOAuthState.create().toString());
        debugger.debug('utilities.buildRequest', 'Set header "X-Middleware-Authorization" to: ' + StripeOAuthState.create().toString());
        return request;
    }

    public static String redactString(String input) {
        return input == null ? null : input.replaceAll('(?<=.{3}).(?=.{3})', '*');
    }

    public static responseData validateCalloutResponse(HttpResponse response, String actionTaken) {
        responseData rd = new responseData();
        rd.put('isConfigSaved', false);
        
        if(response.getStatusCode() >= 200 && response.getStatusCode() < 300) {
            rd.put('isConfigSaved', true);
            if (actionTaken == 'setStripeAccountAsDefault' || actionTaken == 'deleteStripeAccount' || actionTaken == 'saveSyncPreferences' || actionTaken == 'saveFilterSettings' || actionTaken == 'saveMappingConfigurations' || actionTaken == 'activatePolling' || actionTaken == 'saveConfiguration') {
                Map<String, Object> responseBody = (Map<String, Object>) JSON.deserializeUntyped(response.getBody());
                configPayloadHelper.extractConfigHash(responseBody, rd);
            }
            return rd;
        } 

        if((response.getStatusCode() < 200 || response.getStatusCode() >= 300) && (response.getBody() != null)) {
            Map<String, Object> responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
            Object errorMessage = responseBody.get('error');
            throw new CalloutException(errorMessage.toString());
        }
        
        errorLogger.create(actionTaken, String.valueOf(response.getStatusCode()), (String)response.getStatus(), '');
        return rd;
    }
    
    // keeps us from repeating this logic over and over
    // multiCurrFields is usually some reference to CurrencyIsoCode
    public static List<SObject> execMultiCurrLookupQuery(SObjectType objType, String idField, Set<Id> lookupIds, List<String> objFields, List<String> multiCurrFields) {
        List<String> fields = new List<String>(objFields);

        if (UserInfo.isMultiCurrencyOrganization()) {
            fields.addAll(multiCurrFields);
        }

        String fieldList = String.join(fields, ', ');
        String ids = idsToQueryString(lookupIds);

        String table = objType.getDescribe().getName();

        String query = 'SELECT ' + fieldList + ' FROM ' + table;
        query += ' WHERE ' + idField + ' IN ' + ids;

        // this fanciness is to ensure that the array type matches the field so you don't run into silly errors
        Type retType = getDynamicType('List', table);
        List<SObject> retVal = (List<SObject>) retType.newInstance();
        retVal.addAll(Database.query(query));

        return retVal;
    }

    public static String idsToQueryString(Set<Id> ids) {
        List<String> idList = new List<String>();
        for (Id idObj : ids) {
            idList.add(idObj);
        }

        return '(\'' + String.join(idList, '\',\'') + '\')';
    }

    public static Type getDynamicType(String augmentation, String objectType) {
        return getDynamicType(augmentation + '<' + objectType + '>');
    }

    public static Type getDynamicType(String objectType) {
        return Type.forName(objectType);
    }

    public static String getCurrencyIsoCodePreferNonDefault() {
        if (UserInfo.isMultiCurrencyOrganization() == false) {
            return UserInfo.getDefaultCurrency();
        }

        String supportedISOCodesQuery = 'SELECT ISOCode FROM CurrencyType WHERE IsActive = true';
        List<SObject> supportedISOCodesQueryResults = Database.query(supportedISOCodesQuery);
        String isoCurrency = UserInfo.getDefaultCurrency();

        // try to use a non-default isocode for better tests.
        for (SObject iso : supportedISOCodesQueryResults) {
            String isoCode = (String) iso.get('ISOCode');
            if (isoCode != isoCurrency && isoCode != UserInfo.getDefaultCurrency()) {
                isoCurrency = isoCode;
            }
        }

        return isoCurrency;
    }

    public static Stripe_Account__c getPrimaryStripeAccount() {
        List<Stripe_Account__c> accts = [
                SELECT Id, Name, Stripe_ID__c, Is_Live_Mode__c, External_ID__c, Connection_Status__c, Is_Primary__c, CSAC_ID__c
                FROM Stripe_Account__c
                WHERE Is_Primary__c = TRUE
                WITH SECURITY_ENFORCED
        ];

        if (accts.isEmpty()) {
            return null;
        }

        return accts[0];
    }

    //Creates a test metadata record in test context
    @TestVisible static List<Setup_Connection_Data__mdt> setupConfigMetadata {
        get {
            if (setupConfigMetadata == null) {
                setupConfigMetadata = [SELECT Global_Key__c,
                                                Local_Authorization_Target__c, Enable_Apex_Debug__c, OAuth_AutoClose_Window__c,
                                                Enable_JS_Debug__c, Platform_Target__c, OAuth_State_Signing_Key__c,
                                                Middleware_Infrastructure_Version__c
                                        FROM Setup_Connection_Data__mdt
                                        LIMIT 1];
            }
            return setupConfigMetadata;
        } set; }
}