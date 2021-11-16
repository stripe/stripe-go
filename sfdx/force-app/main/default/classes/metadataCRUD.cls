public with sharing class metadataCRUD implements Metadata.DeployCallback {
    public void handleResult(Metadata.DeployResult result, Metadata.DeployCallbackContext context) {
        if (result.status == Metadata.DeployStatus.Succeeded) {
            System.debug(' success : '+ result);
        } else {
            System.debug(' fail : '+ result);
        }
    }

    public static void saveSetupConfigData(Map<String, Object> setupData, String MdtObjectName) {
        Metadata.CustomMetadata setupConfiguration = new Metadata.CustomMetadata();
        setupConfiguration.fullName = Constants.NAMESPACE_API + MdtObjectName;
        setupConfiguration.label = 'Default';
        for(String key : setupData.keySet()) {
            Metadata.CustomMetadataValue customField = new Metadata.CustomMetadataValue();
            customField.field = Constants.NAMESPACE_API + key;
            customField.value = setupData.get(key);
            setupConfiguration.values.add(customField);
        }
        Metadata.DeployContainer mdContainer = new Metadata.DeployContainer();
        mdContainer.addMetadata(setupConfiguration);
        metadataCRUD callback = new metadataCRUD();
        if(!Test.isRunningTest()) {
            Id jobId = Metadata.Operations.enqueueDeployment(mdContainer, callback);
        }
    }
}
