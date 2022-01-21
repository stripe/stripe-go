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

    //saves steps completed as a setup data metadata record to manage/presist states
    @AuraEnabled
    public static String saveData(Map<String, Object> setupData) {
        responseData rd = new responseData();
        try {
             if (setupData != null) {
                List<Setup_Configuration_Data__mdt> data = [SELECT Id, Steps_Completed__c, label FROM Setup_Configuration_Data__mdt WHERE label = 'Default' WITH SECURITY_ENFORCED];
                if(!data.isEmpty()) {
                    Setup_Configuration_Data__mdt setupConfigurationRecord = data[0];
                    Setup_Configuration_Data__mdt dataObj = (Setup_Configuration_Data__mdt)utilities.getSobjectByMap(setupData, constants.NAMESPACE_API + 'Setup_Configuration_Data__mdt');
                    if (!String.isEmpty(dataObj.Steps_Completed__c) && !String.isEmpty(setupConfigurationRecord.Steps_Completed__c)) {
                        Map<String, Object> stepsCompleted = (Map<String, Object>)JSON.deserializeUntyped(setupConfigurationRecord.Steps_Completed__c);
                        Map<String, Object> newStepsCompleted = (Map<String, Object>)JSON.deserializeUntyped(dataObj.Steps_Completed__c);
                        newStepsCompleted.putAll(stepsCompleted);
                        setupData.put('Steps_Completed__c',json.serialize(newStepsCompleted));
                    }          
                }
                metadataCRUD.saveSetupConfigData(setupData, 'Setup_Configuration_Data.Default');
            }
            return getSetupData();
        } catch (Exception e) {
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
            String query = 'SELECT ' + String.join(new List<String>(fields), ',') + ' FROM ' + constants.NAMESPACE_API + 'Setup_Configuration_Data__mdt WHERE label = \'Default\' WITH SECURITY_ENFORCED';
            List<Setup_Configuration_Data__mdt> data = Database.query(query);
            Setup_Configuration_Data__mdt setupData = new Setup_Configuration_Data__mdt();
            if (!data.isEmpty()) {
                setupData = data[0];
            }
            setupData.label = 'Default';
            if (String.isEmpty(setupData.Steps_Completed__c)) {
                setupData.Steps_Completed__c = '{}';
            }
            rd.put('setupData', setupData);
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
                    Map<String, Object> connectionStatus = (Map<String,Object>)responseBody.get('connection_status');
                    if((Boolean)connectionStatus.get('salesforce') == true)stripeConnectRec.Salesforce_Connected__c = true; 
                    if((Boolean)connectionStatus.get('stripe') == true)stripeConnectRec.Stripe_Connected__c = true; 

                    /*if the response shows they are connected we will udpate the Setup Connection metadata record to signify they are connected. 
                    We return a connection status of fresh to determine if we should show a toast message of success or not*/
                    if(((Boolean)connectionStatus.get('stripe') == true) && ((Boolean)connectionStatus.get('salesforce') == true)) {
                        rd.put('isConnected','fresh');
                    } else {
                        rd.put('isConnected',false);
                    }
                    upsert stripeConnectRec;

                } else {
                    /*if we do not get a 200 status code back we create an Error Log record to show we did not get a success
                    response from the ruby service and send a status of failed to signify to show an error toast*/
                    rd.put('isConnected','failed');
                    errorLogger.create('validateConnectionStatus', string.valueOf(response.getStatusCode()), (String)response.getStatus(), JSON.serialize(response));
                }
            } else {
                //if there is no record saved we are not connected
                rd.put('isConnected',false);
                //if there is no record we need to go get a new package level key 
                List<Setup_Connection_Data__mdt> setupConfigList = utilities.getStripeConnectionKey(); 
                Setup_Connection_Data__mdt setupData = setupConfigList[0];

                // generate unique key for this organization used to authenticate against the integration service
                Blob blb = Crypto.GenerateAESKey(128);
                String packageGuidKey = EncodingUtil.ConvertTohex(blb);
                
                String route = constants.RUBY_SERVICE_BASE_URI+'post-install';
                
                Map<String, Object> body = new Map<String, Object>{
                'key' => packageGuidKey
                };

                String globalKey = setupData.Global_Key__c;
                Map<String, String> headers = new Map<String, String>{
                    'Salesforce-Account-Id' => (String)constants.ORG_ID,
                    'Salesforce-Key' => globalKey
                };
                HttpResponse response =  utilities.makeCallout(route, 'POST', JSON.serialize(body), headers);
    
                Map<String, Object> errorBody; 
                if(response.getStatusCode() == 200) {     
                    stripeConnectRec.API_Key__c = packageGuidKey;
                    upsert stripeConnectRec;
                } 
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    /*gets object fields to populate picklist in mapper
    NOTE: This can be sped up to reduce load time potentially by making this method async (by annotating it with this @future) 
    and only called on page load*/
    @AuraEnabled
    public static String getPicklistValuesForMapper(Boolean isConnectedCallback, String ObjectApiName) {
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();   

            if(stripeConnectRec.Id != null) {
                if(stripeConnectRec.Salesforce_Connected__c == true && stripeConnectRec.Stripe_Connected__c == true) {  
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

            if(stripeConnectRec.Id != null) {
                //used to store configuration maps from response
                Map<String, Object> allMappingConfigurations = new Map<String, Object>();
                //construct call out to rubys configuration endpoint
                if(stripeConnectRec.Salesforce_Connected__c == true && stripeConnectRec.Stripe_Connected__c == true) {  
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
                        rd.put('allMappingConfigurations',allMappingConfigurations); 
                        
                    } else {
                        errorLogger.create('getMappingConfigurations', string.valueOf(response.getStatusCode()), (String)response.getStatus(), JSON.serialize(response));
                    }
                } else {
                    rd.put('isConnected',false); 
                }
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
            if(stripeConnectRec.Id != null) {
                if(stripeConnectRec.Salesforce_Connected__c == true && stripeConnectRec.Stripe_Connected__c == true) {                
                    String route = constants.RUBY_SERVICE_BASE_URI+'configuration';
                    HttpResponse response =  utilities.makeCallout(route,'PUT',jsonMappingConfigurationsObject);
                    if(response.getStatusCode() == 200) {
                        rd.put('isConfigSaved',true);
                    } else {
                        rd.put('isConfigSaved',false);
                        errorLogger.create('saveMappingConfigurations', string.valueOf(response.getStatusCode()), (String)response.getStatus(), JSON.serialize(response));
                    }
                }
            }
        
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }
    
    @AuraEnabled
    public static String manualRetry(String jsonMapObj) {
        responseData rd = new responseData();
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults(); 
            
            if(stripeConnectRec.Id != null) {
                if(stripeConnectRec.Salesforce_Connected__c == true && stripeConnectRec.Stripe_Connected__c == true) {             
                    String route = constants.RUBY_SERVICE_BASE_URI+'translate';
                    HttpResponse response =  utilities.makeCallout(route,'POST',jsonMapObj);
                    Map<String, Object> responseBody;
                    if(response.getStatusCode() == 200) {

                    } else {
                        errorLogger.create('manualRetry', string.valueOf(response.getStatusCode()), (String)response.getStatus(), JSON.serialize(response));
                    }
                }
            }
        
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }
}
