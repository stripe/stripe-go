global without sharing class PostInstallClass implements InstallHandler {
    global void onInstall(InstallContext context)  {
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Salesforce_Connected__c, Stripe_Connected__c, Global_Key__c, API_Key__c
                                                            FROM Setup_Connection_Data__mdt
                                                            WITH SECURITY_ENFORCED];
        if(!setupConfigList.isEmpty()) {
            Setup_Connection_Data__mdt setupData = setupConfigList[0];
            Blob blb = Crypto.GenerateAESKey(128);
            String hex = EncodingUtil.ConvertTohex(blb);
            String packageGuidKey = hex.SubString(0,8)+ '-' + hex.SubString(8,12) + '-' + hex.SubString(12,16);
            String orgId = constants.ORG_ID;
            //String globalKey = setupData.Global_Key__c; //Arnold comment out when in package 
            String globalKey = 'ABC123';
            String route = constants.RUBY_SERVICE_BASE_URI+'post-install';

            Map<String, String> headers = new Map<String,String>{
                'Salesforce-Account-Id' => orgId,
                'Salesforce-Key' => globalKey,
                'key' => packageGuidKey
            };
            HttpResponse response =  utilities.makeCallout(route,'POST', headers);

            Map<String, Object> errorBody; 
            if(response.getStatusCode() == 200) {
                setupData.API_Key__c = packageGuidKey;
                Map<String, Object> setupConfig = (Map<String, Object>) JSON.deserializeUntyped(JSON.serialize(setupData));
                metadataCRUD.saveSetupConfigData(setupConfig, 'Setup_Connection_Data__mdt.Default');
            } else {
                errorBody = (Map<String, Object>)JSON.deserializeUntyped(response.getBody());
                errorLogger.create('postInstallScript', (String)errorBody.get('message'), (String)response.getStatus(), JSON.serialize(errorBody));
            }
        }
    }
}
