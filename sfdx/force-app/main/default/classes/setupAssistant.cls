public with sharing class setupAssistant {
    @TestVisible
    private static ConfigPayloadHelper payloadHelper = new ConfigPayloadHelper();

    @TestVisible
    private static Debug_Helper debugger = new Debug_Helper();

    static Boolean isConnected() {
        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
        return utilities.isConnected(stripeConnectRec);
    }

    // downloads sync record and config data from quick action button on sync record component
    @AuraEnabled(Cacheable=true)
    public static String getDownloadForSupportUrl(Id recordId) {
        Pagereference ref = Page.DownloadForSupportConfig;
        ref.getParameters().put('recordId', recordId);
        return ref.getUrl();
    }

    @AuraEnabled(Cacheable=true)
    public static String getExportableConfigDownloadUrl() {
        return Page.ConfigExportDownload.getUrl();
    }

    // gets steps completed from setup data to manage/presist states
    @AuraEnabled
    public static String getSetupData() {
        responseData rd = new responseData();
        Boolean isConnected = isConnected();
        try {
            List<Setup_Data__c> setupDataQuery = [
                SELECT Id, Steps_Completed__c, isSetupComplete__c
                FROM Setup_Data__c
                WHERE Name = :constants.SETUP_DATA_RECORD_NAME
                WITH SECURITY_ENFORCED
            ];
            Setup_Data__c setupData;

            if (!setupDataQuery.isEmpty()) {
                setupData = setupDataQuery[0];
            } else {
                // default values must be set explicitly since this record is NOT upserted and does not inherit database defaults
                setupData = new Setup_Data__c(
                    Steps_Completed__c = '{}',
                    Name = constants.SETUP_DATA_RECORD_NAME,
                    isSetupComplete__c = false
                );
            }

            rd.put('setupData', setupData);
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isConnected', isConnected);
        return rd.getJsonString();
    }

    //saves steps completed as a setup data record to manage/presist states
    @AuraEnabled
    public static String saveData(Map<String, Object> newSetupDataRec, Boolean isSetupComplete) {
        responseData rd = new responseData();
        try {
            if (newSetupDataRec != null) {
                Setup_Data__c setupData = (Setup_Data__c)utilities.getSobjectByMap(newSetupDataRec, constants.NAMESPACE_API + 'Setup_Data__c');

                List<Setup_Data__c> data = [SELECT Id, Steps_Completed__c
                                            FROM Setup_Data__c
                                            WHERE Name = :constants.SETUP_DATA_RECORD_NAME
                                            WITH SECURITY_ENFORCED];

                if (!data.isEmpty()) {
                    // override id
                    setupData.Id = data[0].Id;

                    // merge completed steps
                    if (!String.isEmpty(setupData.Steps_Completed__c) && !String.isEmpty(data[0].Steps_Completed__c)) {
                        Map<String, Object> stepsCompleted = (Map<String, Object>)JSON.deserializeUntyped(data[0].Steps_Completed__c);
                        Map<String, Object> newStepsCompleted = (Map<String, Object>)JSON.deserializeUntyped(setupData.Steps_Completed__c);
                        newStepsCompleted.putAll(stepsCompleted);
                        setupData.Steps_Completed__c = JSON.serialize(newStepsCompleted);
                    }

                    //update setup complete flag if true
                    if(isSetupComplete) {
                        setupData.isSetupComplete__c = true;
                    }
                }

                // override name
                setupData.Name = constants.SETUP_DATA_RECORD_NAME;

                upsert setupData;
            }
            return getSetupData();
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    // validates connection status
    @AuraEnabled
    public static String validateConnectionStatus(Boolean isConnectedCallback, String systemConnected) {
        responseData rd = new responseData();

        try {
            // gets custom setting record to determin if we are in a sandbox or production org
            rd.put('isSandbox', constants.IS_SANDBOX);
            rd.put('salesforceNamespace', constants.NAMESPACE);

            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            // checks if we have a org-level API key generated, if we don't we should generate one
            if(stripeConnectRec.Id == null) {
                // NOTE this is a VERY important line! `generatePackageKey` makes another callout and persists the organization-specific API key
                // on the ruby service. This key is also persisted locally within the salesforce account.
                utilities.generatePackageKey(stripeConnectRec);

                rd.put('isStripeConnected','stripeDisconnected');
                rd.put('isSalesforceConnected','salesforceDisconnected');
                return rd.getJsonString();
            }

            // make callout to ruby services configuration endpoint to see if the org is connected to Stripe and Salesforce
            String route = utilities.getPlatformConfigEndpoint();
            HttpResponse response = utilities.makeCallout(route, 'GET');

            Map<String, Object> errorBody;
            Map<String, Object> responseBody;

            /*
            if we do not get a 200 status code back we create an Error Log record to show we did not get a success
            response from the ruby service and send a status of failed to signify to show an error toast
            */
            if(response.getStatusCode() != 200) {
                rd.put('isStripeConnected','failed');
                rd.put('isSalesforceConnected','failed');
                errorLogger.create('validateConnectionStatus', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get connection status from Ruby service.');
                return rd.getJsonString();
            }

            /*
            if we get a 200 response code from the ruby service we will check the connection status object
            in the payload to determine weather the user successfully connected to stripe or not
            */
            responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

            payloadHelper.extractConnectionStatus(responseBody, stripeConnectRec);

            // TODO why do we care if it's a connected callback? Why can't we just return the same data regardless?
            if(isConnectedCallback) {
                // why do we need distinct strings here for the status? Can't we just pass a boolean?
                // because of the 'freshConnection' value in the condition below.
                String status = stripeConnectRec.Salesforce_Connected__c ? 'salesforceConnected' : 'salesforceDisconnected';
                rd.put('isSalesforceConnected', status);

                status = stripeConnectRec.Stripe_Connected__c ? 'stripeConnected' : 'stripeDisconnected';
                rd.put('isStripeConnected', status);
            } else {
                /*if the response shows they are connected we will udpate the Setup Connection custom setting record to signify they are connected
                We return a connection status of freshConnection to determine if we should show a toast message of success or not*/
                if(systemConnected == constants.SYSTEM_STRIPE) {
                    String status = stripeConnectRec.Stripe_Connected__c ? 'freshConnection' : 'stripeDisconnected';
                    rd.put('isStripeConnected', status);
                } else if(systemConnected == constants.SYSTEM_SALESFORCE) {
                    String status = stripeConnectRec.Stripe_Connected__c ? 'freshConnection' : 'salesforceDisconnected';
                    rd.put('isSalesforceConnected', status);
                }
            }

            upsert stripeConnectRec;
            return rd.getJsonString();
        } catch (Exception e) {
            rd.addError(e);
        }

        return rd.getJsonString();
    }

    //gets object fields to populate picklist in mapper
    @AuraEnabled
    public static String getPicklistValuesForMapper(Boolean isConnectedCallback, String ObjectApiName) {
        responseData rd = new responseData();
        Boolean isConnected = false;
        try {
            isConnected = isConnected();

            if (isConnected) {
                if (isConnectedCallback == true) {
                    //builds a map of fields lists by object
                    Map<String, Set<Object>> fieldListByObjectMap = new Map<String, Set<Object>>{
                        'OrderItem' => utilities.getListOfFieldsByObject('OrderItem'),
                        'Order' => utilities.getListOfFieldsByObject('Order'),
                        'Account' => utilities.getListOfFieldsByObject('Account'),
                        'Product2' => utilities.getListOfFieldsByObject('Product2'),
                        'PricebookEntry' => utilities.getListOfFieldsByObject('PricebookEntry'),
                        'Order_Stripe_Coupon__c' => utilities.getListOfFieldsByObject(constants.NAMESPACE_API + 'Order_Stripe_Coupon__c')
                    };
                    rd.put('fieldListByObjectMap', fieldListByObjectMap);
                } else {
                    rd.put('ObjectApiName', ObjectApiName);
                    rd.put('listOfObjectFields', utilities.getListOfFieldsByObject(ObjectApiName));
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isConnected', isConnected);
        return rd.getJsonString();
    }


    //gets formatted stripe object fields
    @AuraEnabled
    public static String getFormattedStripeObjectFields() {
        responseData rd = new responseData();
        try {
            //using a blank string here to avoid error for no `Salesforce-Key` in make callout function
            Map<String,String> headers = new Map<String,String> {
                'Salesforce-Key' => ''
            };

            String route = constants.RUBY_SERVICE_BASE_URI + '/openapi.json';
            Map<String, Object> responseBody = null;

            Setup_Connection_Data__mdt connectionData = utilities.getConnectionData();
            if (connectionData.Use_Local_OpenAPI_Spec__c) {
                String NS = constants.NAMESPACE == 'c' ? null : constants.NAMESPACE;
                List<StaticResource> resources = [SELECT Body FROM StaticResource WHERE Name = 'OpenApiOverride' AND NamespacePrefix = :NS];
                if (resources.size() > 0) {
                    responseBody = (Map<String, Object>)JSON.deserializeUntyped(resources[0].Body.toString());
                }
            }

            HttpResponse response;
            if (responseBody == null) {
                response = utilities.makeCallout(route,'GET', headers);
                if (response.getStatusCode() == 200) {
                    responseBody = (Map<String, Object>)JSON.deserializeUntyped(response.getBody());
                }
            }

            if(responseBody != null) {
                rd.put('formattedStripeCustomerFields', (List<Object>)responseBody.get('formattedStripeCustomerFields'));
                rd.put('formattedStripeProductItemFields', (List<Object>)responseBody.get('formattedStripeProductItemFields'));
                rd.put('formattedStripeSubscriptionFields', (List<Object>)responseBody.get('formattedStripeSubscriptionFields'));
                rd.put('formattedStripeSubscriptionSchedulePhaseFields', (List<Object>)responseBody.get('formattedStripeSubscriptionSchedulePhaseFields'));
                rd.put('formattedStripeSubscriptionItemFields', (List<Object>)responseBody.get('formattedStripeSubscriptionItemFields'));
                rd.put('formattedStripePriceFields', (List<Object>)responseBody.get('formattedStripePriceFields'));
                rd.put('formattedStripePriceOrderItemFields', (List<Object>)responseBody.get('formattedStripePriceFields'));
                rd.put('formattedStripeCouponFields', (List<Object>)responseBody.get('formattedStripeCouponFields'));
            } else {
                if (response != null) {
                    errorLogger.create('getFormattedStripeObjectFields', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get mapping configuration from ruby service.');
                } else {
                    errorLogger.create('getFormattedStripeObjectFields', 'An unknown error occurred', null, null);
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //gets all saved mappings from ruby service
    @AuraEnabled(Cacheable = false)
    public static String getMappingConfigurations() {
        responseData rd = new responseData();
        Boolean isConnected = false;
        try {
            isConnected = isConnected();

            if (isConnected) {
                //used to store configuration maps from response
                Map<String, Object> allMappingConfigurations = new Map<String, Object>();

                //construct call out to rubys configuration endpoint
                String route = utilities.getPlatformConfigEndpoint();
                HttpResponse response =  utilities.makeCallout(route,'GET');

                //if the call out is successful pull out all mapping objects and add to list
                Map<String, Object> responseBody;
                if (response.getStatusCode() == 200) {
                    responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

                    payloadHelper.extractConfigHash(responseBody, rd);

                    payloadHelper.extractFieldMappingData(responseBody, rd, true);
                } else {
                    errorLogger.create('getMappingConfigurations', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get mapping configuration from ruby service.');
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isConnected',isConnected);
        return rd.getJsonString();
    }

    //saves mappings and sends them to the ruby service to store
    @AuraEnabled
    public static String saveMappingConfigurations(String jsonMappingConfigurationsObject) {
        responseData rd = new responseData();
        try {
            if(isConnected()) {
                String route = utilities.getPlatformConfigEndpoint();
                HttpResponse response =  utilities.makeCallout(route, 'PUT', jsonMappingConfigurationsObject);
                rd = utilities.validateCalloutResponse(response, 'saveMappingConfigurations');
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

     //gets all saved sync prefrences data
     @AuraEnabled(Cacheable = false)
     public static String getSyncPreferences() {
         responseData rd = new responseData();
         try {
            rd.put('isSandbox', constants.IS_SANDBOX);
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            if(utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            //make callout to ruby services configuration endpoint to get stored stripe setting values
            String route = utilities.getPlatformConfigEndpoint();
            HttpResponse response =  utilities.makeCallout(route, 'GET');

            Map<String, Object> errorBody;
            Map<String, Object> responseBody;

            /*if we get a 200 response code from the ruby service we will check the connection status object
            in the payload to determine weather the user successfully connected to stripe or not*/
            if (response.getStatusCode() == 200) {
                responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

                payloadHelper.extractConnectionStatus(responseBody, stripeConnectRec);

                //ensure we are still connected
                Boolean isConnected = stripeConnectRec.Stripe_Connected__c && stripeConnectRec.Salesforce_Connected__c;
                if (isConnected == false) {
                    rd.put('isConnected', false);
                    upsert stripeConnectRec;
                    return rd.getJsonString();
                }

                payloadHelper.extractConfigHash(responseBody, rd);
                payloadHelper.extractSettingsData(responseBody, rd, true);

                String recRetention = rd.getString(ConfigPayloadHelper.SYNC_RECORD_RETENTION);
                if (stripeConnectRec.Sync_Record_Retention__c != recRetention) {
                    stripeConnectRec.Sync_Record_Retention__c = recRetention;
                    upsert stripeConnectRec;
                }
            }
            rd.put('isConnected', true);
        } catch (Exception e) {
             rd.addError(e);
        }

        return rd.getJsonString();
    }

    //gets multi currency options if it is enablec in the org
    @AuraEnabled
    public static String getMulticurrencySelectionOptions() {
        responseData rd = new responseData();
        try {
            Boolean multiCurrencyEnabled = UserInfo.isMultiCurrencyOrganization();
            Map<String,System.OrgLimit> limitsMap = OrgLimits.getMap();
            System.OrgLimit apiRequestsLimit = limitsMap.get('DailyApiRequests');
            if (multiCurrencyEnabled) {
                List<Object> supportedISOCodes = new List<Object>{
                    new Map<String, Object> {
                        'label' => 'Select default currency...',
                        'value' => ''
                    }
                };

                /* using a dynamic query here because the currencyType object can only be referenced directly
                when the multicurrency feature is enabled in the org otherwise we get a type error*/
                String supportedISOCodesQuery = 'SELECT ISOCode FROM CurrencyType WHERE IsActive = true';
                List<sObject> supportedISOCodesQueryResults = Database.query(supportedISOCodesQuery);

                //Iterate through the supported currencies to build list of supported ISO code
                for(sObject supportedISOCode : supportedISOCodesQueryResults) {
                    supportedISOCodes.add(new Map<String, Object> {
                        'label' => supportedISOCode.get('ISOCode'),
                        'value' => supportedISOCode.get('ISOCode')
                    });
                }
                rd.put('supportedISOCodes', supportedISOCodes);
            }
            rd.put('isMultiCurrencyEnabled', multiCurrencyEnabled);
            rd.put('orgMaxApiLimit', apiRequestsLimit.getLimit());
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //sends updated sync preference data to the ruby service and saves the values
    @AuraEnabled
    public static String saveSyncPreferences(String defaultCurrency, String syncRecordRetention, String syncStartDate, String apiPercentageLimit, String cpqTermUnit, String cpqProratePrecision, Boolean pollingEnabled, String configurationHash) {
        responseData rd = new responseData();
        try {
            if (isConnected()) {
                Map<String, Object> settings = new Map<String, Object> {
                    'polling_enabled' => pollingEnabled,
                    'sync_start_date' => syncStartDate,
                    'sync_record_retention' => syncRecordRetention,
                    'default_currency' => UserInfo.getDefaultCurrency(),
                    'cpq_term_unit' => cpqTermUnit,
                    'cpq_prorate_precision' => cpqProratePrecision
                };

                if (apiPercentageLimit != null) {
                    settings.put('api_percentage_limit', apiPercentageLimit);
                }

                if (cpqTermUnit != null) {
                    settings.put('cpq_term_unit', cpqTermUnit);
                }

                Map<String, Object> resquestBody = new Map<String, Object> {
                    'configuration_hash' => configurationHash,
                    'settings' => settings
                };

                String route = utilities.getPlatformConfigEndpoint();
                HttpResponse response =  utilities.makeCallout(route, 'PUT', JSON.serialize(resquestBody));
                rd = utilities.validateCalloutResponse(response, 'saveSyncPreferences');

            }
        } catch (Exception e) {
           rd.addError(e);
        }
        return rd.getJsonString();
    }

    //activates polling in the ruby service
    @AuraEnabled
    public static String activatePolling(String syncStartDate, Boolean isConfigEnabled, String configurationHash) {
        responseData rd = new responseData();
        try {
            if (isConnected()) {
                Map<String, Object> resquestBody = new Map<String, Object> {
                        'enabled' => isConfigEnabled,
                        'configuration_hash' => configurationHash,
                        'settings' => new Map<String, Object> {
                                'polling_enabled' => true,
                                'sync_start_date' => syncStartDate
                        }
                };

                String route = utilities.getPlatformConfigEndpoint();
                HttpResponse response =  utilities.makeCallout(route, 'PUT', JSON.serialize(resquestBody));
                rd = utilities.validateCalloutResponse(response, 'activatePolling');
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //sends records to be synced to ruby service called from quick action button on record detail page
    @AuraEnabled
    public static String manualRetry(Id recordId) {
        responseData rd = new responseData();
        Boolean isSyncRecordDispactched = false;
        try {
            if (isConnected()) {
                //get sync record information for ruby service
                List<Sync_Record__c> syncRecordList = [SELECT Id, Primary_Record_ID__c, Secondary_Record_ID__c, Resolution_Status__c, Secondary_Object_Type__c, Primary_Object_Type__c
                                                        FROM Sync_Record__c
                                                        WHERE Resolution_Status__c != 'Resolved'
                                                        AND Id = :recordId
                                                        WITH SECURITY_ENFORCED
                                                        LIMIT 1];

                if (syncRecordList.isEmpty()) {
                    rd.put('isSyncRecordDispactched', isSyncRecordDispactched);
                    return rd.getJsonString();
                }

                Sync_Record__c recordToSync = syncRecordList[0];
                
                String route = utilities.getPlatformTranslateEndpoint();

                Map<String, Object> body = new Map<String, Object>{
                    'object_type' => (String)recordToSync.Primary_Object_Type__c,
                    'object_ids' => new List<String> {
                        (String)recordToSync.Primary_Record_ID__c
                    }
                };

                HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body));

                if (response.getStatusCode() == 200) {
                    isSyncRecordDispactched = true;
                } else {
                    rd.put('isSyncRecordDispactched', false);
                    errorLogger.create('manualRetry', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Sync_Record ID that failed: '+(String)recordId);
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isSyncRecordDispactched', isSyncRecordDispactched);
        return rd.getJsonString();
    }

     //gets all saved filter settings
     @AuraEnabled
    public static String getFilterSettings() {
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            if (utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            //make callout to ruby services configuration endpoint to get stored stripe setting values
            String route = utilities.getPlatformConfigEndpoint();
            HttpResponse response =  utilities.makeCallout(route, 'GET');

            if (response.getStatusCode() != 200) {
                errorLogger.create('getFilterSettings', String.valueOf(response.getStatusCode()), (String)response.getStatus(), '');
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            Map<String, Object> responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

            //get connection status values from response
            Boolean isConnected = payloadHelper.extractConnectionStatus(responseBody, stripeConnectRec);

            //ensure we are still connected
            if (isConnected == false) {
                rd.put('isConnected', false);
                upsert stripeConnectRec;
                return rd.getJsonString();
            }

            payloadHelper.extractFilterSettingsData(responseBody, rd);
        } catch (Exception e) {
            rd.addError(e);
        }

        rd.put('isConnected', true);
        return rd.getJsonString();
    }

    //sends updated sync preference data to the ruby service and saves the values
    @AuraEnabled
    public static String saveFilterSettings(String productFilter, String orderFilter, String accountFiter, String priceBookFilter) {
        responseData rd = new responseData();
        try {
            Map<String, String> queryFilterMap = new Map<String, String> {
                'Product2' => productFilter,
                'Order' => orderFilter,
                'Account' => accountFiter,
                'PricebookEntry' => priceBookFilter
            };

            List<Map<String, Object>> validationErrorMapList = new List<Map<String, Object>> ();
            //Iterate through filters and run query to ensure they are valid
            for (String objectName : queryFilterMap.keySet()) {

                if (String.isBlank(queryFilterMap.get(objectName)) || queryFilterMap.get(objectName) == null) {
                    queryFilterMap.put(objectName, null);
                    continue;
                }

                String whereClause = queryFilterMap.get(objectName);
                try {
                    /*
                    Using limit zero in queries so there are no records are returned therefore it
                    will not count against the total number of rows retrieved salesforce limitation
                    */
                    List<sObject> sobjList = Database.query('SELECT Id FROM ' + objectName + ' WHERE ' + whereClause + ' WITH SECURITY_ENFORCED LIMIT 0');
                } catch (QueryException e) {
                    validationErrorMapList.add(new Map<String, Object> {
                        'Object' => objectName,
                        'Error' => String.valueOf(e)
                    });
                }
            }

            if (!validationErrorMapList.isEmpty()) {
                rd.put('isValidationError', true);
                rd.put('isFiltersSaved', false);
                rd.put('ValidationErrors', validationErrorMapList);
                return rd.getJsonString();
            }

            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
            if (utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isFiltersSaved', false);
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            Map<String, Object> resquestBody = new Map<String, Object> {
                'settings' => new Map<String, Object> {
                    'filters' => queryFilterMap
                }
            };
            String route = utilities.getPlatformConfigEndpoint();
            HttpResponse response =  utilities.makeCallout(route, 'PUT', JSON.serialize(resquestBody));

            if(response.getStatusCode() != 200) {
                rd.put('isValidationError', false);
                rd.put('isFiltersSaved', false);
                errorLogger.create('saveFilterSettings', String.valueOf(response.getStatusCode()), (String)response.getStatus(), '');
            }

        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isFiltersSaved', true);
        rd.put('isValidationError', false);
        return rd.getJsonString();
    }

    //sends object name for all records syncing to the ruby service
    @AuraEnabled
    public static String syncAllRecords(String objectType) {
        debugger.debug('setupAssistant.syncAllRecords', 'objectType: ' + objectType);
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            if (utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            debugger.debug('setupAssistant.syncAllRecords', 'is connected');

            String route = utilities.getPlatformTranslateAllEndpoint();

            Map<String, Object> body = new Map<String, Object>{
                'object_type' => objectType
            };

            debugger.debug('setupAssistant.syncAllRecords', 'route: ' + route);
            debugger.debug('setupAssistant.syncAllRecords', 'body: ' + JSON.serialize(body));
            HttpResponse response = utilities.makeCallout(route, 'POST', JSON.serialize(body));

            debugger.debug('setupAssistant.syncAllRecords', 'Got response status code: ' + response.getStatusCode());

            if (response.getStatusCode() != 200) {
                rd.put('syncAllRecordsDispatched', false);
                errorLogger.create('syncAllRecords', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Object type that failed: ' + objectType);
                return rd.getJsonString();
            }

            rd.put('syncAllRecordsDispatched', true);
            rd.put('isConnected', true);

        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //sends object name for all records syncing to the ruby service
    @AuraEnabled
    public static String manualTranslation(String translationRecordID) {
        responseData rd = new responseData();
        try {
            List<String> listOfAcceptableObjects = new List<String> {
                'Account',
                'Order',
                'Pricebook2',
                'Product2'
            };
            rd.put('isRecordValid', false);

            Id recordId = Id.valueOf(translationRecordID);
            String objectName = string.valueOf(recordId.getsobjecttype());
            if(!listOfAcceptableObjects.contains(objectName)) {
                rd.put('errorMessage', 'this is an id for a record on the ' + objectName + ' object which is not suppoted, full list of supported objects: ' + listOfAcceptableObjects);
                return rd.getJsonString();
            }

            String validationQuery = 'SELECT Id FROM ' + objectName + ' WHERE Id = :recordId WITH SECURITY_ENFORCED LIMIT 1';
            List<SObject> listOfFoundRecords = Database.query(validationQuery);
            if (listOfFoundRecords.isEmpty()) {
                rd.put('errorMessage', 'No ' + objectName + ' objects found for this Salesforce record Id');
                return rd.getJsonString();
            }

            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();

            if (utilities.isConnected(stripeConnectRec) == false) {
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            //here is where we need to validate if its record or not
            String route = utilities.getPlatformTranslateEndpoint();

            Map<String, Object> body = new Map<String, Object> {
                'object_type' => objectName,
                'object_ids' => new List<String> {
                    translationRecordID
                }
            };

            HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body));
            if(response.getStatusCode() == 200) {
                rd.put('isRecordValid', true);
            } else {
                errorLogger.create('manualTranslation', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Translation Record ID That Failed: '+(String)recordId);
                throw new CalloutException('There was an error in queueing this record for translation, check sync logs for more detail');
            }

        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //performs a permission check on user before granting access to the application
    @AuraEnabled
    public static String checkUserPermissions() {
        responseData rd = new responseData();
        try {
            //Map for front end state setting
            Map<String, Object> permissionIssueMap = new Map<String, Object> {
                'isPermSetAssigned' => false,
                'isSystemPermissionMissing' => false,
                'isObjectPermissionMissing' => false,
                'missingPermissionSets' => new List<String>(),
                'missingObjectPermissionList' => new List<String>(),
                'missingSystemPermissionList' => new List<String>()
            };

            //List of Objects to check access on
            List<String> listOfObjectsToCheckAccessOn = new List<String> {
                'Account',
                'Product2',
                'Order',
                'OrderItem',
                'PricebookEntry'
            };

            if ((Boolean)utilities.isCpqEnabled()) {
                listOfObjectsToCheckAccessOn.add('SBQQ__Quote__c');
                listOfObjectsToCheckAccessOn.add('SBQQ__OrderItemConsumptionSchedule__c');
                listOfObjectsToCheckAccessOn.add('SBQQ__OrderItemConsumptionRate__c');
                listOfObjectsToCheckAccessOn.add('ConsumptionSchedule');
                listOfObjectsToCheckAccessOn.add('ConsumptionRate');
            }

            List<Map<String,Object>> missingObjectPermissionList = validateObjectLevelPermission(listOfObjectsToCheckAccessOn);


            //add list of missing permissions if any were found
            if (!missingObjectPermissionList.isEmpty()) {
                permissionIssueMap.put('isObjectPermissionMissing', true);
                permissionIssueMap.put('missingObjectPermissionList', missingObjectPermissionList);
            }

            /* this list holds system permission fields which hold boolean denoting if the user has access or not.
            This is incase there are other system permissions need to be checked in the future. Full list of Profile fields
            here https://developer.salesforce.com/docs/atlas.en-us.sfFieldRef.meta/sfFieldRef/salesforce_field_reference_Profile.htm.*/
            List<String> permissionList = new List<String> {
                'Name',
                'PermissionsEditActivatedOrders'
            };

            //Used to get label names of permissions
            Map<String, Schema.SObjectType> schemaMap = Schema.getGlobalDescribe();
            Schema.SObjectType profileSchema = schemaMap.get('Profile');
            Map<String, Schema.SObjectField> fieldMap = profileSchema.getDescribe().fields.getMap();

            //Map of system permissions by label
            Map<String, Object> activePermissionMap = new Map<String, Object>();
            for (String permisson : permissionList) {
                if (permisson == 'Name') {
                    continue;
                }
                activePermissionMap.put(fieldMap.get(permisson).getDescribe().getLabel(), false);
            }

            List<String> userFieldsList = new List<String> {
                'Id',
                'Name'
            };

            for (String permission : permissionList) {
                userFieldsList.add('Profile.' + permission);
            }

            List<String> permissionSetAssignmentFieldsList = new List<String>();
            for (String permission : permissionList) {
                permissionSetAssignmentFieldsList.add('PermissionSet.' + permission);
            }

            List<SObject> userPermissions = Database.Query (
                String.join (
                    new List<String> {
                        'SELECT',
                        String.join(userFieldsList, ', '),
                        ',',
                        '(SELECT Id,',
                        String.join(permissionSetAssignmentFieldsList, ', '),
                        'FROM PermissionSetAssignments) FROM User WHERE Id = \'' +
                        UserInfo.getUserId() + '\' WITH SECURITY_ENFORCED LIMIT 1'
                    },
                    ' '
                )
            );

            if (userPermissions.isEmpty()) {
                rd.put('permissionIssueMap', permissionIssueMap);
                return rd.getJsonString();
            }
            SObject userProfileSystemPermissions = userPermissions[0].getSobject('Profile');

            //check profile level permissions
            for (String permisson : permissionList) {
                String permissionLabel = fieldMap.get(permisson).getDescribe().getLabel();
                if (permisson != 'Name') {
                    activePermissionMap.put(permissionLabel, (Boolean)userProfileSystemPermissions.get(permisson));
                }
            }

            Map<String, Object> validateUserPermissionsMap = validateUserPermissions(userPermissions, activePermissionMap, permissionIssueMap, fieldMap, permissionList);
            activePermissionMap = (Map<String, Object>)validateUserPermissionsMap.get('activePermissionMap');
            permissionIssueMap = (Map<String, Object>)validateUserPermissionsMap.get('permissionIssueMap');


            List<String> missingSystemPermissionList = new List<String>();
            for (String permission : activePermissionMap.keySet()) {
                if (!(Boolean)activePermissionMap.get(permission)) {
                    permissionIssueMap.put('isSystemPermissionMissing', true);
                    missingSystemPermissionList.add(permission);
                }
            }

            if (!missingSystemPermissionList.isEmpty()) {
                permissionIssueMap.put('missingSystemPermissionList', missingSystemPermissionList);
            }

            rd.put('permissionIssueMap', permissionIssueMap);

        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    // builds map based on object permissioning
    private static List<Map<String,Object>> validateObjectLevelPermission(List<String> listOfObjectsToCheckAccessOn) {
        List<Map<String,Object>> missingObjectPermissionList = new List<Map<String,Object>>();
        for (String objectApiName : listOfObjectsToCheckAccessOn) {
            Boolean hasObjectAccess = true;
            DescribeSObjectResult objectDescribe = dmlManager.getSObjectDescribeByObjectName(objectApiName);

            Map<String,Object> missingPermissionsMap = new Map<String,Object> {
                'Object' => objectDescribe.getLabel(),
                'Access' => true,
                'Create' => true,
                'Edit' => true
            };

            if (!(Boolean)objectDescribe.isAccessible()) {
                hasObjectAccess = false;
                missingPermissionsMap.put('Access', false);
            }

            if (!(Boolean)objectDescribe.isCreateable()) {
                hasObjectAccess = false;
                missingPermissionsMap.put('Create', false);
            }

            if (!(Boolean)objectDescribe.isUpdateable()) {
                hasObjectAccess = false;
                missingPermissionsMap.put('Edit', false);
            }

            if (!hasObjectAccess) {
                missingObjectPermissionList.add(missingPermissionsMap);
            }
        }
        return missingObjectPermissionList;
    }

    // check each permset assigned to the user
    private static Map<String, Object> validateUserPermissions(List<SObject> userPermissions, Map<String, Object> activePermissionMap, Map<String, Object> permissionIssueMap, Map<String, Schema.SObjectField> fieldMap, List<String> permissionList) {
        List<SObject> userPermissionsAssignments = userPermissions[0].getSObjects('PermissionSetAssignments');
        Boolean isPermSetMissing = true;
        if (!userPermissionsAssignments.isEmpty()) {
            List<String> listOfAssignedPermissionSets = new List<String> {
                constants.PACKAGED_PERMISSION_SET_NAME
            };
            for (SObject userPermissionsAssignment : userPermissionsAssignments) {
                SObject userPermSetSystemPermissions = userPermissionsAssignment.getSobject('PermissionSet');

                if(userPermSetSystemPermissions.get('Name') == constants.PACKAGED_PERMISSION_SET_NAME.replace(' ','_')) {
                    isPermSetMissing = false;
                }

                for (String permisson : permissionList) {
                    String permissionLabel = fieldMap.get(permisson).getDescribe().getLabel();
                    if(permisson != 'Name') {
                        Boolean systemPermissionValue = (Boolean)userPermSetSystemPermissions.get(permisson);
                        //only update the permission in the map if it is true in the permset and false in the map
                        if (!(Boolean)activePermissionMap.get(permissionLabel) && systemPermissionValue) {
                            activePermissionMap.put(permissionLabel, systemPermissionValue);
                        }
                    }
                }
            }
        }

        //set missing permission set state
        if (isPermSetMissing) {
            permissionIssueMap.put('missingPermissionSets', new List<String> {
                constants.PACKAGED_PERMISSION_SET_NAME
            });
        } else {
            permissionIssueMap.put('isPermSetAssigned', true);
        }

        return new Map<String, Object> {
           'activePermissionMap' => activePermissionMap,
           'permissionIssueMap' => permissionIssueMap
        };
    }


    //Creates a test metadata record in test context
    @testVisible static List<Setup_Connection_Data__mdt> setupConfigMetadata {
        get {
            if (setupConfigMetadata == null)
                setupConfigMetadata = utilities.getStripeConnectionKey();
            return setupConfigMetadata;
        } set; }
}
