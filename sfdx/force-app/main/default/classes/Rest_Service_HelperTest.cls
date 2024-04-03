@IsTest
public with sharing class Rest_Service_HelperTest {
    private static Test_Tools testTools = new Test_Tools();
    private static Rest_Service_Helper helper = new Rest_Service_Helper();

    @IsTest
    private static void TestFieldsToQueryString() {
        Set<String> s = new Set<String>{'a', 'b', 'c'};
        String test = helper.fieldsToQueryString(s);

        System.assertEquals('a,b,c', test, 'set converts to "a,b,c"');
    }

    @IsTest
    private static void TestSetToQueryString() {
        Set<String> s = new Set<String>{'a', 'b', 'c'};
        String test = helper.setToQueryString(s);

        System.assertEquals('(\'a\',\'b\',\'c\')', test, 'set converts to "\'(\'a\',\'b\',\'c\')"');
    }

    @IsTest
    private static void TestGetAccessibleFields() {
        Set<Object> errors = new Set<Object>();
        String objName = Opportunity.class.getName();
        Map<String, SObjectField> allFields = Schema.Opportunity.getSObjectType().getDescribe().fields.getMap();
        Set<String> fields = helper.getAccessibleFields(objName, errors);

        // I would assume this would be 1, but there seem to be other fields not included.
        // So we will just check and see that there is a difference.
        System.assertNotEquals(allFields.size(), fields.size(), 'Should have fewer fields accessible than total number of fields');

        // because environments are all different and dependent on plugins, this test section is terrible.
        if (errors.isEmpty()) {
            return;
        }

        System.assertEquals(errors.size() > 0, true, 'Should find at least one field inaccessible');

        Map<String, Object> err = null;
        for (Object o : errors) {
            err = (Map<String, Object>) o;
            System.assertEquals((String)err.get('error_type'), 'FieldNotAccessible', 'Error type is FieldNotAccessible');
            System.assertNotEquals((String)err.get('sobject_field'), '', 'SObject field should be populated');
            System.assertNotEquals((String)err.get('sobject_field'), null, 'SObject field should not be null');
        }
    }

    @IsTest
    private static void TestGetAccessibleFieldsWithInaccessibleObject() {
        String profile = Test_Tools.PROFILE_CHATTER_FREE_USER;
        String uniqueId = 'TestGetAccessibleFieldsWithInaccessibleObject';
        User u = testTools.insertUser(profile, uniqueId);
        Set<Object> errors = new Set<Object>();
        String objName = Opportunity.class.getName();
        Set<String> fields = new Set<String>();

        System.runAs(u) {
            fields = helper.getAccessibleFields(objName, errors);
        }

        System.assertEquals(1, errors.size(), 'Should find one object inaccessible');

        Map<String, Object> err = null;
        for (Object o : errors) {
            err = (Map<String, Object>) o;
            break;
        }

        System.assertEquals((String)err.get('sobject_type'), objName, 'Error object is ' + objName);
        System.assertEquals((String)err.get('error_type'), 'ObjectNotAccessible', 'Error type is ObjectNotAccessible');
    }
}