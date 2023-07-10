/**
 * Created by jmather-c on 4/27/23.
 */

public with sharing class ConfigExportVFController {
    ConfigPayloadHelper payloadHelper = new ConfigPayloadHelper();

    public String getConfig() {
        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
        Boolean isConnected = utilities.isConnected(stripeConnectRec);
        responseData rd = new responseData();

        if (isConnected == false) {
            rd.put('error', 'Unable to connect to remote system.');
            return rd.getResultsJsonString();
        }

        try {
            //construct call out to rubys configuration endpoint
            String route = constants.RUBY_SERVICE_BASE_URI + '/v1/configuration';
            HttpResponse response =  utilities.makeCallout(route,'GET');

            //if the call out is successful pull out all mapping objects and add to list
            Map<String, Object> responseBody;
            if(response.getStatusCode() == 200) {
                responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
                payloadHelper.extractSettingsData(responseBody, rd, false);
                payloadHelper.extractFieldMappingData(responseBody, rd, false);
                payloadHelper.extractFilterSettingsData(responseBody, rd);
            } else {
                errorLogger.create('getConfiguration', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get mapping configuration from ruby service.');
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getResultsJsonString();
    }
}