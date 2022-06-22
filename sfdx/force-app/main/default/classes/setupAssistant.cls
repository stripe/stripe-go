public with sharing class setupAssistant { 

     //gets steps completed from setup data to manage/presist states
    @AuraEnabled
    public static String getSetupData() {
        responseData rd = new responseData();
        Boolean isConnected = true;
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            if(stripeConnectRec.Id == null || !stripeConnectRec.Salesforce_Connected__c || !stripeConnectRec.Stripe_Connected__c) {
                isConnected = false;
            }
            Set<String> fields = new Set<String> { 
                'Id',
                'Steps_Completed__c',
                'isSetupComplete__c'
            };

            List<Setup_Data__c> setupDataQuery = [SELECT Id, Steps_Completed__c, isSetupComplete__c
                                        FROM Setup_Data__c
                                        WHERE Name = :constants.SETUP_DATA_RECORD_NAME
                                        WITH SECURITY_ENFORCED];
            Setup_Data__c setupData;

            if (!setupDataQuery.isEmpty()) {
                setupData = setupDataQuery[0];
            } else {
                setupData = new Setup_Data__c();
                setupData.Name = constants.SETUP_DATA_RECORD_NAME;
            }

            if (String.isEmpty(setupData.Steps_Completed__c)) {
                setupData.Steps_Completed__c = '{}';
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

    //validates connection status 
    @AuraEnabled
    public static String validateConnectionStatus(Boolean isConnectedCallback, String systemConnected) {
        responseData rd = new responseData();
        try {
            //gets custom setting record to determin if we are in a sandbox or production org
            rd.put('isSandbox', constants.IS_SANDBOX);
            rd.put('salesforceNamespace', constants.NAMESPACE);

            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            // checks if we have a package-level API key generated
            if(stripeConnectRec.Id != null) {

                //make callout to ruby services configuration endpoint to see if the org is connected to Stripe and Salesforce
                String route = constants.RUBY_SERVICE_BASE_URI + '/v1/configuration';
                HttpResponse response =  utilities.makeCallout(route, 'GET');

                Map<String, Object> errorBody; 
                Map<String, Object> responseBody;

                /*if we get a 200 response code from the ruby service we will check the connection status object
                in the payload to determine weather the user successfully connected to stripe or not*/
                if(response.getStatusCode() == 200) {
                    responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
                    
                    //get connection status values from response
                    Map<String, Object> connectionStatus = (Map<String,Object>)responseBody.get('connection_status');
                    stripeConnectRec.Stripe_Connected__c = (Boolean)connectionStatus.get('stripe');
                    stripeConnectRec.Salesforce_Connected__c = (Boolean)connectionStatus.get('salesforce');
                    
                    if(isConnectedCallback) {
                        if(!(Boolean)connectionStatus.get('salesforce')) {
                            rd.put('isSalesforceConnected','salesforceDisconnected');
                        } else {
                            rd.put('isSalesforceConnected','salesforceConnected');
                        }

                        if(!(Boolean)connectionStatus.get('stripe')) {
                            rd.put('isStripeConnected','stripeDisconnected');
                        } else {
                            rd.put('isStripeConnected','stripeConnected');
                        }
                        upsert stripeConnectRec;
                        return rd.getJsonString();

                    } else {
                        /*if the response shows they are connected we will udpate the Setup Connection custom setting record to signify they are connected
                        We return a connection status of freshConnection to determine if we should show a toast message of success or not*/
                        if(systemConnected == 'stripe') {
                            if (!(Boolean)connectionStatus.get('stripe')) {
                                rd.put('isStripeConnected','stripeDisconnected');               
                            } else {
                                rd.put('isStripeConnected','freshConnection');
                            }
                            upsert stripeConnectRec;
                            return rd.getJsonString();
                        }
                        
                        if(systemConnected == 'salesforce') {
                            if (!(Boolean)connectionStatus.get('salesforce')) {
                                rd.put('isSalesforceConnected','salesforceDisconnected');
                            } else {
                                rd.put('isSalesforceConnected','freshConnection');
                            }
                            upsert stripeConnectRec;
                            return rd.getJsonString();
                        }
                    }

                } else {
                    /*if we do not get a 200 status code back we create an Error Log record to show we did not get a success
                    response from the ruby service and send a status of failed to signify to show an error toast*/
                    rd.put('isStripeConnected','failed');
                    rd.put('isSalesforceConnected','failed');
                    errorLogger.create('validateConnectionStatus', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get connection status from Ruby service.');
                    return rd.getJsonString();
                }

            } else {
                //if there is no org default record we will generate the package key and create the record
                rd.put('isStripeConnected','stripeDisconnected');
                rd.put('isSalesforceConnected','salesforceDisconnected');  
                utilities.generatePackageKey(stripeConnectRec);
                return rd.getJsonString();
            }
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
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            if(stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {  
                isConnected = true;
                if(isConnectedCallback == true) {
                    //builds a map of fields lists by object    
                    Map<String, Set<Object>> fieldListByObjectMap = new Map<String, Set<Object>>{
                        'OrderItem' => utilities.getListOfFieldsByObject('OrderItem'),
                        'Order' => utilities.getListOfFieldsByObject('Order'),
                        'Account' => utilities.getListOfFieldsByObject('Account'),
                        'Product2' => utilities.getListOfFieldsByObject('Product2'),
                        'PricebookEntry' => utilities.getListOfFieldsByObject('PricebookEntry')
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
            HttpResponse response =  utilities.makeCallout(route,'GET', headers);
            
            Map<String, Object> responseBody;
            if(response.getStatusCode() == 200) {
                responseBody = (Map<String, Object>)JSON.deserializeUntyped(response.getBody());
                rd.put('formattedStripeCustomerFields', (List<Object>)responseBody.get('formattedStripeCustomerFields'));
                rd.put('formattedStripeProductItemFields', (List<Object>)responseBody.get('formattedStripeProductItemFields'));
                rd.put('formattedStripeSubscriptionFields', (List<Object>)responseBody.get('formattedStripeSubscriptionFields'));
                rd.put('formattedStripeSubscriptionItemFields', (List<Object>)responseBody.get('formattedStripeSubscriptionItemFields'));
                rd.put('formattedStripePriceFields', (List<Object>)responseBody.get('formattedStripePriceFields'));
            } else {
                errorLogger.create('getFormattedStripeObjectFields', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get mapping configuration from ruby service.');
            }
            
        } catch (Exception e) {     
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //gets all saved mappings from ruby service
    @AuraEnabled
    public static String getMappingConfigurations() {
        responseData rd = new responseData();
        Boolean isConnected = false;
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            if(stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {
                //used to store configuration maps from response
                Map<String, Object> allMappingConfigurations = new Map<String, Object>();

                //construct call out to rubys configuration endpoint
                isConnected = true;           
                String route = constants.RUBY_SERVICE_BASE_URI + '/v1/configuration';
                HttpResponse response =  utilities.makeCallout(route,'GET');

                //if the call out is successful pull out all mapping objects and add to list 
                Map<String, Object> responseBody;
                if(response.getStatusCode() == 200) {
                    responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());

                    Map<String, Object> fieldMappingsMap = (Map<String,Object>)responseBody.get('field_mappings');
                    if (fieldMappingsMap != null && !fieldMappingsMap.isEmpty()) { 
                        allMappingConfigurations.put('field_mappings', fieldMappingsMap);
                    }

                    Map<String, Object> defaultMappingsMap = (Map<String,Object>)responseBody.get('default_mappings');
                    if (defaultMappingsMap != null && !defaultMappingsMap.isEmpty()) { 
                        allMappingConfigurations.put('default_mappings', defaultMappingsMap);
                    }

                    Map<String, Object> requiredMappingsMap = (Map<String,Object>)responseBody.get('required_mappings');
                    if (requiredMappingsMap != null && !requiredMappingsMap.isEmpty()) { 
                        allMappingConfigurations.put('required_mappings', requiredMappingsMap);
                    }

                    Map<String, Object> fieldDefaultsMappingsMap = (Map<String,Object>)responseBody.get('field_defaults');
                    if (fieldDefaultsMappingsMap != null  && !fieldDefaultsMappingsMap.isEmpty()) {
                        allMappingConfigurations.put('field_defaults', fieldDefaultsMappingsMap);
                    }
                    rd.put('allMappingConfigurations', allMappingConfigurations); 
                    
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
        Boolean isConfigSaved = false;
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();  
            if(stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {     
                String route = constants.RUBY_SERVICE_BASE_URI + '/v1/configuration';
                HttpResponse response =  utilities.makeCallout(route,'PUT',jsonMappingConfigurationsObject);
                if(response.getStatusCode() == 200) {
                    isConfigSaved = true;
                } else {
                    errorLogger.create('saveMappingConfigurations', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Configuration that failed: '+jsonMappingConfigurationsObject);
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isConfigSaved', isConfigSaved);
        return rd.getJsonString();
    }

     //gets all saved sync prefrences data
     @AuraEnabled
     public static String getSyncPreferences() {
         responseData rd = new responseData();
         try {
            rd.put('isSandbox', constants.IS_SANDBOX);
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            if(stripeConnectRec.Id == null || !stripeConnectRec.Salesforce_Connected__c || !stripeConnectRec.Stripe_Connected__c) { 
                rd.put('isConnected', false); 
                return rd.getJsonString();
            }

            //make callout to ruby services configuration endpoint to get stored stripe setting values
            String route = constants.RUBY_SERVICE_BASE_URI + '/v1/configuration';
            HttpResponse response =  utilities.makeCallout(route, 'GET');

            Map<String, Object> errorBody; 
            Map<String, Object> responseBody;

            /*if we get a 200 response code from the ruby service we will check the connection status object
            in the payload to determine weather the user successfully connected to stripe or not*/
            if(response.getStatusCode() == 200) {
                responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
                
                //get connection status values from response
                Map<String, Object> connectionStatus = (Map<String,Object>)responseBody.get('connection_status');

                //ensure we are still connected
                if(!(Boolean)connectionStatus.get('stripe') || !(Boolean)connectionStatus.get('salesforce')) {
                    rd.put('isConnected', false); 
                    stripeConnectRec.Stripe_Connected__c = (Boolean)connectionStatus.get('stripe');
                    stripeConnectRec.Salesforce_Connected__c = (Boolean)connectionStatus.get('salesforce');
                    upsert stripeConnectRec;
                    return rd.getJsonString();
                }

                Map<String, Object> settings = (Map<String,Object>)responseBody.get('settings');

                List<String> settingsFields = new List<String> {
                    'default_currency',
                    'sync_record_retention',
                    'sync_start_date',
                    'api_percentage_limit',
                    'cpq_term_unit'
                };

                List<String> connectionReadOnlyFields = new List<String> {
                    'stripe_account_id',
                    'last_synced'
                };

                //gets non read only settings fields from response for frontend
                for(String settingsField : settingsFields) {
                    rd.put(settingsField, String.valueOf(settings.get(settingsField)));
                    if (settingsField == 'sync_record_retention') {
                        rd.put(settingsField, String.valueOf(settings.get(settingsField)));
                        stripeConnectRec.Sync_Record_Retention__c = String.valueOf(settings.get(settingsField));
                        upsert stripeConnectRec;
                    }
                }

                //gets read only connection fields from response for frontend
                for(String readOnlyField : connectionReadOnlyFields) {
                    rd.put(readOnlyField, String.valueOf(connectionStatus.get(readOnlyField)));
                }
            } 
            rd.put('isConnected', true); 
            rd.put('isCpqInstalled', utilities.isCpqEnabled());
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
    public static String saveSyncPreferences(String defaultCurrency, String syncRecordRetention, String syncStartDate, String apiPercentageLimit, String cpqTermUnit) {
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults(); 
            if(stripeConnectRec.Id != null) {
                Map<String, Object> resquestBody = new Map<String, Object> {
                    'settings' => new Map<String, Object> {
                        'api_percentage_limit' => apiPercentageLimit,
                        'sync_start_date' => syncStartDate,
                        'sync_record_retention' => syncRecordRetention,
                        'default_currency' => defaultCurrency,
                        'multicurrency_enabled' => UserInfo.isMultiCurrencyOrganization(),
                        'cpq_term_unit' => cpqTermUnit
                    }
                };
                String route = constants.RUBY_SERVICE_BASE_URI + '/v1/configuration';
                HttpResponse response =  utilities.makeCallout(route, 'PUT', JSON.serialize(resquestBody));

                if(response.getStatusCode() == 200) {
                    stripeConnectRec.Sync_Record_Retention__c = syncRecordRetention;
                    upsert stripeConnectRec;
                } else {
                    errorLogger.create('saveSyncPreferences', String.valueOf(response.getStatusCode()), (String)response.getStatus(), '');
                }
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
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults(); 
            
            if(stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {  
                //get sync record information for ruby service 
                List<Sync_Record__c> syncRecordList = [SELECT Id, Primary_Record_ID__c, Secondary_Record_ID__c, Resolution_Status__c, Secondary_Object_Type__c, Primary_Object_Type__c
                                                        FROM Sync_Record__c
                                                        WHERE Resolution_Status__c != 'Resolved'
                                                        AND Id = :recordId
                                                        WITH SECURITY_ENFORCED
                                                        LIMIT 1];   

                if(!syncRecordList.isEmpty()) {
                    Sync_Record__c recordToSync = syncRecordList[0];
                    String route = constants.RUBY_SERVICE_BASE_URI + '/v1/translate';
            
                    Map<String, Object> body = new Map<String, Object>{
                        'object_type' => (String)recordToSync.Primary_Object_Type__c,
                        'object_ids' => new List<String> {
                            (String)recordToSync.Primary_Record_ID__c
                        }
                    };

                    HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body));

                    if(response.getStatusCode() == 200) {
                        isSyncRecordDispactched = true;           
                    } else {
                        rd.put('isSyncRecordDispactched', false);
                        errorLogger.create('manualRetry', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Sync_Record ID that failed: '+(String)recordId);
                    }
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

            if (stripeConnectRec.Id == null || !stripeConnectRec.Salesforce_Connected__c || !stripeConnectRec.Stripe_Connected__c) {  
                rd.put('isConnected', false);
                return rd.getJsonString();
            }


            //make callout to ruby services configuration endpoint to get stored stripe setting values
            String route = constants.RUBY_SERVICE_BASE_URI + '/v1/configuration';
            HttpResponse response =  utilities.makeCallout(route, 'GET');
            
            if(response.getStatusCode() != 200) { 
                errorLogger.create('getFilterSettings', String.valueOf(response.getStatusCode()), (String)response.getStatus(), '');
                rd.put('isConnected', false);
                return rd.getJsonString();
            }
            Map<String, Object> responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
            
            //get connection status values from response
            Map<String, Object> connectionStatus = (Map<String,Object>)responseBody.get('connection_status');

            //ensure we are still connected
            if(!(Boolean)connectionStatus.get('stripe') || !(Boolean)connectionStatus.get('salesforce')) {
                rd.put('isConnected', false);
                stripeConnectRec.Stripe_Connected__c = (Boolean)connectionStatus.get('stripe');
                stripeConnectRec.Salesforce_Connected__c = (Boolean)connectionStatus.get('salesforce');
                upsert stripeConnectRec;
                return rd.getJsonString();
            }

            Map<String, Object> settings = (Map<String,Object>)responseBody.get('settings');

            Map<String, Object> filters = (Map<String,Object>)settings.get('filters');

            List<String> filterFields = new List<String> {
                'Product2',
                'Order',
                'Account'
            };

            //gets filter fields and values for front end
            for(String filterField : filterFields) {
                rd.put(filterField, String.valueOf(filters.get(filterField)));
            }

        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('isConnected', true);
        return rd.getJsonString();
    }

    //sends updated sync preference data to the ruby service and saves the values 
    @AuraEnabled
    public static String saveFilterSettings(String productFilter, String orderFilter, String accountFiter) {
        responseData rd = new responseData();
        try {
            Map<String, String> queryFilterMap = new Map<String, String> {
                'Product2' => productFilter,
                'Order' => orderFilter,
                'Account' => accountFiter
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
                    /*Using limit zero in queries so there are no records are returned therefore it 
                    will not count against the total number of rows retrieved salesforce limitation*/
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
            if (stripeConnectRec.Id == null || !stripeConnectRec.Salesforce_Connected__c || !stripeConnectRec.Stripe_Connected__c) {  
                rd.put('isFiltersSaved', false);
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            Map<String, Object> resquestBody = new Map<String, Object> {
                'settings' => new Map<String, Object> {
                    'filters' => queryFilterMap
                }
            };
            String route = constants.RUBY_SERVICE_BASE_URI + '/v1/configuration';
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
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults(); 
            
            if (stripeConnectRec.Id == null || !stripeConnectRec.Salesforce_Connected__c || !stripeConnectRec.Stripe_Connected__c) {  
                rd.put('isConnected', false);
                return rd.getJsonString();
            }

            String route = constants.RUBY_SERVICE_BASE_URI + '/v1/translate_all';
    
            Map<String, Object> body = new Map<String, Object>{
                'object_type' => objectType
            };

            HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body));

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
