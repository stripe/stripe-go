/**
 * Created by jmather-c on 5/1/23.
 */

@IsTest
public with sharing class ConfigExportVFCont_Test {
    @IsTest
    static public void TestGetConfig() {
        Setup_Connection_Data__mdt testSetupData = test_setupAssistant.getTestStripeConnectionKey();
        test_setupAssistant.setTestGlobalKey(testSetupData);
        test_setupAssistant.insertTestConnectedRecord();
        Test.setMock(HttpCalloutMock.class, new test_setupAssistant.UnifiedConfigMock());
        ConfigExportVFController controller = new ConfigExportVFController();

        Test.startTest();
        String response = controller.getConfig();
        Map<String, Object> resultsMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        Map<String, Object> allMappingConfigurations = (Map<String, Object>)resultsMap.get('allMappingConfigurations');
        Map<String, Object> default_mappings = (Map<String, Object>)allMappingConfigurations.get('default_mappings');
        Map<String, Object> field_defaults = (Map<String, Object>)allMappingConfigurations.get('field_defaults');
        Map<String, Object> field_mappings = (Map<String, Object>)allMappingConfigurations.get('field_mappings');
        Map<String, Object> required_mappings = (Map<String, Object>)allMappingConfigurations.get('required_mappings');

        System.assertNotEquals(allMappingConfigurations, null);
        System.assertEquals(default_mappings, null);
        System.assertEquals(required_mappings, null);
        System.assertNotEquals(field_defaults, null);
        System.assertNotEquals(field_mappings, null);
        System.assertEquals('2 days', (String)resultsMap.get('sync_record_retention'));
        System.assertEquals('yesterday', (String)resultsMap.get('sync_start_date'));

        System.assertEquals('testOrderFilter', (String)resultsMap.get('order_filter'));
        System.assertEquals('testAccountFilter', (String)resultsMap.get('account_filter'));
        System.assertEquals('testProduct2Filter', (String)resultsMap.get('product_filter'));
        System.assertEquals('testPricebookEntryFilter', (String)resultsMap.get('pricebook_entry_filter'));
    }
}
