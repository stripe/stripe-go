public with sharing class setupAssistant {
    //gets the namespace of the package
    public String namespace { get; set; }

    public setupAssistant() {
        namespace = constants.NAMESPACE;
    }

    //gets package version
    @AuraEnabled
    public static String getPackageVersion() {
        responseData rd = new responseData();
        try {
            rd.put('major', System.requestVersion().major());
            rd.put('minor', System.requestVersion().minor());
        } catch(Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

     //gets steps completed from setup data to manage/presist states
    @AuraEnabled
    public static String getSetupData() {
        responseData rd = new responseData();
        try {
            Set<String> fields = new Set<String> { 
                'Id',
                'Steps_Completed__c'
            };
            String query = 'SELECT ' + String.join(new List<String>(fields), ',') + ' FROM ' + constants.NAMESPACE_API + 'Setup_Data__c WHERE Name = \'setupData\' WITH SECURITY_ENFORCED';
            List<Setup_Data__c> data = Database.query(query);
            Setup_Data__c setupData = new Setup_Data__c();
            if (!data.isEmpty()) {
                setupData = data[0];
            }
            setupData.Name = 'setupData';
            if (String.isEmpty(setupData.Steps_Completed__c)) {
                setupData.Steps_Completed__c = '{}';
            }
            rd.put('setupData', setupData);
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    } 

    //saves steps completed as a setup data record to manage/presist states
    @AuraEnabled
    public static String saveData(Map<String, Object> newSetupDataRec) {
        responseData rd = new responseData();
        try {
            if (newSetupDataRec != null) {
                Setup_Data__c setupData = (Setup_Data__c)utilities.getSobjectByMap(newSetupDataRec, constants.NAMESPACE_API + 'Setup_Data__c');
            
                List<Setup_Data__c> data = [SELECT Id, Steps_Completed__c FROM Setup_Data__c WHERE Name = 'setupData' WITH SECURITY_ENFORCED];

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
                }

                // override name
                setupData.Name = 'setupData';

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
    public static String validateConnectionStatus(Boolean isConnectedCallback) {
        responseData rd = new responseData();
        try {
            //gets custom setting record to determin if we are in a sandbox or production org
            rd.put('isSandbox',constants.IS_SANDBOX);

            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            if(stripeConnectRec.Id != null) {
                /*checks to see if the user is already connected this runs on render of the component 
                by passing a boolean param to denote this is being called on render*/
                if((stripeConnectRec.Salesforce_Connected__c == true && stripeConnectRec.Stripe_Connected__c == true) && isConnectedCallback == true) {
                    rd.put('isConnected',true);
                    return rd.getJsonString();
                } else if ((stripeConnectRec.Salesforce_Connected__c != true && stripeConnectRec.Stripe_Connected__c != true) && isConnectedCallback == true) {
                    rd.put('isConnected',false);                
                    return rd.getJsonString();
                }
                //make callout to ruby services configuration endpoint to see if the org is connected to Stripe and Salesforce
                String route = constants.RUBY_SERVICE_BASE_URI+'configuration';
                HttpResponse response =  utilities.makeCallout(route,'GET');

                Map<String, Object> errorBody; 
                Map<String, Object> responseBody;

                /*if we get a 200 response code from the ruby service we will check the connection status object
                in the payload to determine weather the user successfully connected to stripe or not*/
                if(response.getStatusCode() == 200) {
                    responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
                    
                    //get connection status values from response
                    Map<String, Object> connectionStatus = (Map<String,Object>)responseBody.get('connection_status');
                    if((Boolean)connectionStatus.get('salesforce') != true) rd.put('isConnected','salesforceDisconnected');
                    if((Boolean)connectionStatus.get('stripe') != true) rd.put('isConnected','stripeDisconnected');


                    /*if the response shows they are connected we will udpate the Setup Connection custom setting record to signify they are connected
                    We return a connection status of fresh to determine if we should show a toast message of success or not*/
                    if((Boolean)connectionStatus.get('stripe') || (Boolean)connectionStatus.get('salesforce')) {
                        Map<String, Object> settings = (Map<String,Object>)responseBody.get('settings');
                        List<String> connectionSettingsFields = new List<String> {
                            'stripe_account_id',
                            'last_synced',
                            'default_currency',
                            'sync_record_retention',
                            'sync_start_date',
                            'api_percentage_limit'
                        };
                        for(String settingsField : connectionSettingsFields) {
                            rd.put( settingsField, string.valueOf(connectionStatus.get(settingsField)));
                            switch on settingsField {
                                when 'stripe_account_id' {
                                    stripeConnectRec.Stripe_Account_Id__c = string.valueOf(connectionStatus.get(settingsField));
                                }
                                when 'last_synced' {
                                    stripeConnectRec.Last_Synced__c = string.valueOf(connectionStatus.get(settingsField));
                                }
                                when 'default_currency' {
                                    stripeConnectRec.Default_Currency__c = string.valueOf(connectionStatus.get(settingsField));
                                }
                                when 'sync_record_retention' {
                                    stripeConnectRec.Sync_Record_Retention__c = string.valueOf(connectionStatus.get(settingsField));
                                }
                                when 'sync_start_date' {
                                    stripeConnectRec.Sync_Start_Date__c = string.valueOf(connectionStatus.get(settingsField));
                                }
                                when 'api_percentage_limit' {
                                    stripeConnectRec.Api_Percentage_Limit__c = string.valueOf(connectionStatus.get(settingsField));
                                }
                            }
                        }
                        stripeConnectRec.Stripe_Connected__c = (Boolean)connectionStatus.get('stripe');
                        stripeConnectRec.Salesforce_Connected__c = (Boolean)connectionStatus.get('salesforce');
                        upsert stripeConnectRec;
                        rd.put('isConnected','fresh');
                    } else {
                        rd.put('isConnected',false);
                    }
                } else {
                    /*if we do not get a 200 status code back we create an Error Log record to show we did not get a success
                    response from the ruby service and send a status of failed to signify to show an error toast*/
                    rd.put('isConnected','failed');
                    errorLogger.create('validateConnectionStatus', string.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get connection status from Ruby service.');
                }
            } else {
                //if there is no record saved we are not connected
                rd.put('isConnected',false);

                // generate unique key for this organization used to authenticate against the integration service
                Blob blb = Crypto.GenerateAESKey(128);
                String packageKey  = EncodingUtil.ConvertTohex(blb);
                
                String route = constants.RUBY_SERVICE_BASE_URI+'post-install';
                
                Map<String, Object> body = new Map<String, Object>{
                'key' => packageKey 
                };

                HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body));
    
                Map<String, Object> errorBody; 
                if(response.getStatusCode() == 200) {     
                    stripeConnectRec.API_Key__c = packageKey ;
                    upsert stripeConnectRec;
                } else {
                    rd.put('isConnected','failed');
                    errorLogger.create('validateConnectionStatus', string.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get package level API key from ruby service.');
                }
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
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            if(stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {  
                rd.put('isConnected',true);  
                if(isConnectedCallback == true) {
                    //builds a map of fields lists by object    
                    Map<String, Set<Object>> fieldListByObjectMap = new Map<String, Set<Object>>{
                        'OrderItem' => utilities.getListOfFieldsByObject('OrderItem'),
                        'Order' => utilities.getListOfFieldsByObject('Order'),
                        'Account' => utilities.getListOfFieldsByObject('Account'),
                        'Product2' => utilities.getListOfFieldsByObject('Product2'),
                        'PricebookEntry' => utilities.getListOfFieldsByObject('PricebookEntry')
                    };   
                    rd.put('fieldListByObjectMap',fieldListByObjectMap);
                } else {
                    rd.put('ObjectApiName',ObjectApiName);
                    rd.put('listOfObjectFields',utilities.getListOfFieldsByObject(ObjectApiName));
                }
            } else {
                rd.put('isConnected',false);  
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
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            if(stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {
                //used to store configuration maps from response
                Map<String, Object> allMappingConfigurations = new Map<String, Object>();

                //construct call out to rubys configuration endpoint
                rd.put('isConnected',true);               
                String route = constants.RUBY_SERVICE_BASE_URI+'configuration';
                HttpResponse response =  utilities.makeCallout(route,'GET');

                //if the call out is successful pull out all mapping objects and add to list 
                Map<String, Object> responseBody;
                if(response.getStatusCode() == 200) {
                    responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
                    Map<String, Object> fieldMappingsMap = (Map<String,Object>)responseBody.get('field_mappings');
                    if (fieldMappingsMap != null && !fieldMappingsMap.isEmpty()) { 
                        allMappingConfigurations.put('field_mappings', fieldMappingsMap);
                    }
                    Map<String, Object> defaultMappingMap = (Map<String,Object>)responseBody.get('default_mapping');
                    if (defaultMappingMap != null && !defaultMappingMap.isEmpty()) { 
                        allMappingConfigurations.put('default_mapping', defaultMappingMap);
                    }

                    Map<String, Object> requiredMappingsMap = (Map<String,Object>)responseBody.get('required_mappings');
                    if (requiredMappingsMap != null && !requiredMappingsMap.isEmpty()) { 
                        allMappingConfigurations.put('required_mappings', requiredMappingsMap);
                    }

                    Map<String, Object> fieldDefaultsMappingsMap = (Map<String,Object>)responseBody.get('field_defaults');
                    if (fieldDefaultsMappingsMap != null  && !fieldDefaultsMappingsMap.isEmpty()) {
                        allMappingConfigurations.put('field_defaults', fieldDefaultsMappingsMap);
                    }
                    rd.put('allMappingConfigurations',allMappingConfigurations); 
                    
                } else {
                    errorLogger.create('getMappingConfigurations', string.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get mapping configuration from ruby service.');
                }

            } else {
                rd.put('isConnected',false); 
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    //saves mappings and sends them to the ruby service to store 
    @AuraEnabled
    public static String saveMappingConfigurations(String jsonMappingConfigurationsObject) {
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();  
            if(stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {     
                String route = constants.RUBY_SERVICE_BASE_URI+'configuration';
                HttpResponse response =  utilities.makeCallout(route,'PUT',jsonMappingConfigurationsObject);
                if(response.getStatusCode() == 200) {
                    rd.put('isConfigSaved',true);
                } else {
                    rd.put('isConfigSaved',false);
                    errorLogger.create('saveMappingConfigurations', string.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Configuration that failed: '+jsonMappingConfigurationsObject);
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
                    String route = constants.RUBY_SERVICE_BASE_URI + 'translate';
            
                    Map<String, Object> body = new Map<String, Object>{
                        'object_type' => (String)recordToSync.Primary_Object_Type__c,
                        'object_ids' => new List<String> {
                            (String)recordToSync.Primary_Record_ID__c
                        }
                    };

                    HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body));

                    if(response.getStatusCode() == 200) {
                        rd.put('isSyncRecordDispactched',true);
                    } else {
                        errorLogger.create('manualRetry', string.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Sync_Record ID that failed: '+(String)recordId);
                    }
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }
    //Creates a test metadata record in test context
    @testVisible static List<Setup_Connection_Data__mdt> setupConfigMetadata { 
        get {
            if (setupConfigMetadata == null)
                setupConfigMetadata = utilities.getStripeConnectionKey(); 
            return setupConfigMetadata;
        } set; }
}
