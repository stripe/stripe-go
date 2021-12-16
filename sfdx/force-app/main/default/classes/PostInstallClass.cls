global without sharing class PostInstallClass implements InstallHandler {
    global void onInstall(InstallContext context)  {
        // grab packaged setup connection record with global API key used to communicate with the integration service
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Salesforce_Connected__c, Stripe_Connected__c, Global_Key__c, API_Key__c
                                                            FROM Setup_Connection_Data__mdt
                                                            WITH SECURITY_ENFORCED
                                                            LIMIT 1];

        /*get setup record and make sure the package level GUID doesnt already exist 
        this key will already exist only in the case of upgrades */
        if(!setupConfigList.isEmpty() && String.isBlank(setupConfigList[0].API_Key__c)) {
            Setup_Connection_Data__mdt setupData = setupConfigList[0];

            // generate unique key for this organization used to authenticate against the integration service
            Blob blb = Crypto.GenerateAESKey(128);
            String packageGuidKey = EncodingUtil.ConvertTohex(blb);

            String route = constants.RUBY_SERVICE_BASE_URI+'post-install';
            Map<String, Object> body = new Map<String, Object>{
                'key' => packageGuidKey
            };

            //send request
            HttpResponse response =  utilities.makeCallout(route,'POST', JSON.serialize(body));

            //parse response after confirming we got a 200 response code  
            Map<String, Object> errorBody; 
            if(response.getStatusCode() == 200) {                
                Map<String, Object> setupConfigMap = new Map<String, Object> {
                    'API_Key__c' => packageGuidKey
                };

                //save package GUID in protected custom metadata
                metadataCRUD.saveSetupConfigData(setupConfigMap, 'Setup_Connection_Data__mdt.Default');
            } else {

                //create error log record if we did not get a 200 response from the ruby service 
                errorLogger.create('postInstallScript', string.valueOf(response.getStatusCode()), (String)response.getStatus(), JSON.serialize(response));
            }
        }
    }
}
