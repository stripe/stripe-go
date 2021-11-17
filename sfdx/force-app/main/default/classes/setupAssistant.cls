public with sharing class setupAssistant {
    public String namespace { get; set; }

    public setupAssistant() {
        namespace = constants.NAMESPACE;
    }

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

    @AuraEnabled
    public static String saveData(Map<String, Object> setupData) {
        responseData rd = new responseData();
        try {
             if (setupData != null) {
                List<Setup_Configuration_Data__mdt> data = [SELECT Id, Steps_Completed__c, label FROM Setup_Configuration_Data__mdt WHERE label = 'Default' WITH SECURITY_ENFORCED];
                if(!data.isEmpty()){
                    Setup_Configuration_Data__mdt dataObj = (Setup_Configuration_Data__mdt)utilities.getSobjectByMap(setupData, constants.NAMESPACE_API + 'Setup_Configuration_Data__mdt');
                    if (!String.isEmpty(dataObj.Steps_Completed__c) && !String.isEmpty(data[0].Steps_Completed__c)) {
                        Map<String, Object> stepsCompleted = (Map<String, Object>)JSON.deserializeUntyped(data[0].Steps_Completed__c);
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

    @AuraEnabled
    public static String isCpqEnabled() {
        responseData rd = new responseData();
        try {
            Boolean isCpqInstalled = UserInfo.isCurrentUserLicensed('SBQQDS');
            rd.put('isCpqInstalled',isCpqInstalled);
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    @testVisible static List<Setup_Connection_Data__mdt> setupConfigMetadata { 
        get {
            if (setupConfigMetadata == null)
                setupConfigMetadata = [SELECT Salesforce_Connected__c, Stripe_Connected__c, Global_Key__c, API_Key__c
                                       FROM Setup_Connection_Data__mdt
                                       WITH SECURITY_ENFORCED]; 
            return setupConfigMetadata;
        } set; }

    @AuraEnabled
    public static String validateConnectionStatus(Boolean isConnectedCallback) {
        responseData rd = new responseData();
        try {
            List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Salesforce_Connected__c, Stripe_Connected__c, Global_Key__c, API_Key__c
                                                                FROM Setup_Connection_Data__mdt
                                                                WITH SECURITY_ENFORCED];   
                                                                
            if(Test.isRunningTest()){  
                setupConfigList = setupConfigMetadata;
            }

            if(!setupConfigList.isEmpty()) {
                Setup_Connection_Data__mdt setupData = setupConfigList[0];
                if((setupData.Salesforce_Connected__c == true && setupData.Stripe_Connected__c == true) && isConnectedCallback == true) {
                    rd.put('isConnected',true);
                    return rd.getJsonString();
                } else if ((setupData.Salesforce_Connected__c != true && setupData.Stripe_Connected__c != true) && isConnectedCallback == true) {
                    rd.put('isConnected',false);
                    return rd.getJsonString();
                }
                String route = constants.RUBY_SERVICE_BASE_URI+'configuration';
                HttpResponse response =  utilities.makeCallout(route,'GET');

                Map<String, Object> errorBody; 
                Map<String, Object> responseBody;
                if(response.getStatusCode() == 200) {
                    responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
                    Map<String, Object> connectionStatus = (Map<String,Object>)responseBody.get('connection_status');
                    if((Boolean)connectionStatus.get('salesforce') == true)setupData.Salesforce_Connected__c = true; 
                    if((Boolean)connectionStatus.get('stripe') == true)setupData.Stripe_Connected__c = true; 
                    if(((Boolean)connectionStatus.get('stripe') == true) && ((Boolean)connectionStatus.get('salesforce') == true)) {
                        rd.put('isConnected','fresh');
                    } else {
                        rd.put('isConnected',false);
                    }
                    Map<String, Object> setupConfigMap = new Map<String, Object> {
                        'Salesforce_Connected__c' => setupData.Salesforce_Connected__c,
                        'Stripe_Connected__c' => setupData.Stripe_Connected__c 
                    };
                    metadataCRUD.saveSetupConfigData(setupConfigMap, 'Setup_Connection_Data__mdt.Default');
                } else {
                    rd.put('isConnected',false);
                    errorBody = (Map<String, Object>)JSON.deserializeUntyped(response.getBody());
                    errorLogger.create('validateConnectionStatus', (String)errorBody.get('message'), (String)response.getStatus(), JSON.serialize(errorBody));
                }
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }
}
