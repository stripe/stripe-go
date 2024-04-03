@IsTest
public with sharing class ConfigPayloadHelper_Test {
    @IsTest
    static void testStripHiddenMapperData() {
        String jsonData = '{"foo": {"bar.baz": "biz"}, "customer": {"address.state": "MI", "address.country": "USA"}, "coupon": {}}';
        Map<String, Object> data = (Map<String, Object>) JSON.deserializeUntyped(jsonData);
        Map<String, Object> customer = (Map<String, Object>) data.get('customer');
        Map<String, Object> foo = (Map<String, Object>) data.get('foo');
        List<String> hiddenMapperFields = new List<String> { 'coupon', 'foo.bar', 'customer.address.country' };
        ConfigPayloadHelper.HiddenMapperFieldAnalysis analysis =
                new ConfigPayloadHelper.HiddenMapperFieldAnalysis(hiddenMapperFields);

        ConfigPayloadHelper helper = new ConfigPayloadHelper();

        System.assert(data.containsKey('coupon'), 'coupon key exists');
        System.assert(data.containsKey('customer'), 'customer key exists');
        System.assert(customer.containsKey('address.country'), 'customer.address.country key exists');

        helper.stripHiddenMapperFields(analysis, data);

        System.assert(false == data.containsKey('coupon'), 'coupon key does not exists');
        System.assert(data.containsKey('customer'), 'customer key exists');
        System.assert(false == customer.containsKey('address.country'), 'customer.address.country key does not exists');
        System.assert(data.containsKey('foo'), 'foo key exists');
        System.assert(false == foo.containsKey('bar.baz'), 'foo.bar.baz key does not exists');

        // also making sure this does not error...
        helper.stripHiddenMapperFields(analysis, null);
    }

    @IsTest
    static void testStripPartialMapFields() {
        Map<String, String> data = new Map<String, String> {
                'a.a' => 'a',
                'a.b' => 'b',
                'ab.a' => 'c'
        };
        List<String> partialFields = new List<String> { 'a' };
        ConfigPayloadHelper helper = new ConfigPayloadHelper();

        helper.debugger = new Debug_Helper(true);

        helper.stripPartialMapFields(data, null);
        helper.stripPartialMapFields(data, partialFields);
        helper.stripPartialMapFields(null, partialFields);

        System.assert(false == data.containsKey('a.a'), 'a.a key does not exists');
        System.assert(false == data.containsKey('a.b'), 'a.b key does not exists');
        System.assert(data.containsKey('ab.a'), 'b.a key does not exists');
    }

    @IsTest
    static void testStripMapFields() {
        Map<String, String> data = new Map<String, String> {
                'b.a' => 'c',
                'b.b' => 'd',
                'c.a' => 'e'
        };
        List<String> fullFields = new List<String> { 'b.a' };
        ConfigPayloadHelper helper = new ConfigPayloadHelper();

        helper.stripMapFields(data, null);
        helper.stripMapFields(data, fullFields);
        helper.stripMapFields(null, fullFields);

        System.assert(false == data.containsKey('b.a'), 'b.a key does not exists');
        System.assert(data.containsKey('b.b'), 'b.b key exists');
        System.assert(data.containsKey('c.a'), 'c.a key exists');
    }

    @IsTest
    static void testGetListFieldFromPayload() {
        ConfigPayloadHelper helper = new ConfigPayloadHelper();
        Map<String,Object> responseBody = new Map<String, Object> {
                'empty' => new List<Object>(),
                'some' => new List<Object> { 'A', 'B' },
                'broken' => 'foo'
        };

        List<String> nullRes = helper.getListFieldFromPayload('null', responseBody);
        List<String> nullResRaw = helper.getListFieldFromPayload('null', responseBody, false);
        List<String> someRes = helper.getListFieldFromPayload('some', responseBody);
        List<String> emptyRes = helper.getListFieldFromPayload('empty', responseBody);
        List<String> brokenRes = helper.getListFieldFromPayload('broken', responseBody, false);

        System.assertNotEquals(null, nullRes, 'null ref with ensured list is not null');
        System.assertEquals(0, nullRes.size(), 'null ref with ensure list is empty');

        System.assertEquals(null, nullResRaw, 'null ref without ensure list is null');

        System.assertNotEquals(null, emptyRes, 'empty ref is not null');
        System.assertEquals(0, emptyRes.size(), 'empty ref is empty');

        System.assertNotEquals(null, someRes, 'some ref is not null');
        System.assertEquals(2, someRes.size(), 'some ref is not empty');

        System.assertEquals(null, brokenRes, 'broken ref is null');
    }

    @IsTest
    static void testExtractValueV2() {
        Setup_Connection_Data__mdt testSetupData = test_setupAssistant.getTestStripeConnectionKey();
        testSetupData.Middleware_Infrastructure_Version__c = 'v2';
        test_setupAssistant.setTestGlobalKey(testSetupData);
        test_setupAssistant.insertTestConnectedRecord();

        Boolean isV2 = utilities.isUsingV2Middleware();
        System.assert(isV2, true);

        ConfigPayloadHelper.ErrorTracker errorTracker = new ConfigPayloadHelper.ErrorTracker();
        ConfigPayloadHelper configPayloadHelper = new ConfigPayloadHelper();
        Map<String, Object> requestPayload = new Map<String,Object>{
            'foo' => 'bar',
            'cpq_prorate_precision' => 'month'

        };
        Map<String, Object> saveTo = new Map<String,Object>{};

        Boolean result = configPayloadHelper.extractValue('cpq_prorate_precision', requestPayload, saveTo, errorTracker);
        System.assert(result == false, 'cpq_prorate_precision should be extracted');
        System.assert((String)saveTo.get('cpq_prorate_precision') == 'month');
        
        result = configPayloadHelper.extractValue('foo', requestPayload, saveTo, errorTracker);
        System.assert(result == true, 'field food should be missing');
        System.assert(saveTo.get('foo') == null, 'field food should be missing');
        
        result = configPayloadHelper.extractValue('missing', requestPayload, saveTo, errorTracker);
        System.assert(result == true, 'field missing should be missing');
        System.assert(saveTo.get('foo') == null, 'field missing should be missing');
    }

     @IsTest
     static void testHiddenSyncPrefFields() {
        ConfigPayloadHelper helper = new ConfigPayloadHelper();
        List<String> actual = helper.getHiddenSyncPrefFields();
        List<String> expected = new List<String> { 'api_percentage_limit', 'cpq_term_unit' };
        System.assert(actual.size() == expected.size(), 'hidden sync pref fields should be the same size');
    }

    @IsTest
    static void testGetHiddenMapperFields() {
        ConfigPayloadHelper configPayloadHelper = new ConfigPayloadHelper();

        Map<String, Object> testConfigPayload = new Map<String, Object> {
            'a.a' => 'a',
            'a.b' => 'b',
            'ab.a' => 'c'
        };
        List<String> actual = configPayloadHelper.getHiddenMapperFields(testConfigPayload);
        System.assert(actual.size() == 0, 'hidden mapper fields should be empty');
    }
}
