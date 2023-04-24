
/**
 * Created by jmather-c on 3/14/23.
 */

@IsTest
public with sharing class RelationshipMapper_Test {
    static Map<String, ConfigData> testData = new Map<String, ConfigData> {
            'price' => new ConfigData('price', new Map<String, String> {
                    'metadata.Product' => 'Product2.Name',
                    'metadata.TaxCode' => 'Product2.SBQQ__TaxCode__c',
                    'metadata.Interval' => 'Product2.SBQQ__SubscriptionTerm__c',
                    'metadata.Usage Type' => 'Product2.SBQQ__BillingType__c',
                    'metadata.Description' => 'Product2.Description',
                    'metadata.Interval Count' => 'Product2.SBQQ__BillingFrequency__c',
                    'metadata.Unit Amount Decimal' => 'UnitPrice'
            }),
            'product' => new ConfigData('product', new Map<String, String> {
                    'name' => 'Name',
                    'description' => 'Description',
                    'metadata.TaxCode' => 'SBQQ__TaxCode__c',
                    'metadata.Tax Code' => 'SBQQ__TaxCode__c',
//            'metadata.Product Base' => 'Product_Base__c',
//            'metadata.Product Line' => 'Product_Line__c',
//            'metadata.Product Type' => 'Product_Type__c',
//            'metadata.Product Group' => 'Product_Group__c',
                    'metadata.Product Family' => 'Family'
//            'metadata.Product Segment' => 'Product_Segment__c',
//            'metadata.Product Super Group' => 'Product_Super_Group__c'
            }),
            'customer' => new ConfigData('customer', new Map<String, String> {
                    'name' => 'Name',
//            'email' => 'Primary_Contact__c.Email',
                    'phone' => 'Phone',
                    'description' => 'Description',
                    'address.city' => 'BillingCity',
                    'address.line1' => 'BillingStreet',
                    'address.state' => 'BillingState',
//            'shipping.name' => 'Primary_Contact__c.Litmos__Full_Name__c',
                    'shipping.phone' => 'Phone',
                    'address.country' => 'BillingCountry',
//            'metadata.region' => 'Geo__c',
//            'metadata.channel' => 'Contract_Relationship__c',
                    'address.postal_code' => 'BillingPostalCode',
                    'metadata.subsidiary' => 'Parent.Name',
                    // changed from Account_Owner_Name__c
                    'metadata.ae_rep_name' => 'Owner.Name',
//            'metadata.account_type' => 'Type__c',
                    'shipping.address.city' => 'ShippingCity',
//            'metadata.payment terms' => 'Payment_Terms__c',
                    'shipping.address.line1' => 'ShippingStreet',
                    'shipping.address.state' => 'ShippingState',
//            'metadata.cf_customer_id' => 'Customer_ID__c',
//            'metadata.collectability' => 'Litmos__Sla__c',
                    'metadata.sfdc_account_id' => 'Id',
                    'shipping.address.country' => 'ShippingCountry',
                    'metadata.sfdc_account_number' => 'AccountNumber',
                    'shipping.address.postal_code' => 'BillingPostalCode'
            }),
            'subscription' => new ConfigData('subscription', new Map<String, String> {
//            'metadata.Selling Method' => 'SBQQ__Quote__c.Selling_Method__c'
            }),
            'price_order_item' => new ConfigData('price_order_item', new Map<String, String> {
                    'metadata.Product' => 'Product2.Name',
                    'metadata.TaxCode' => 'Product2.SBQQ__TaxCode__c',
                    'metadata.Interval' => 'Product2.SBQQ__SubscriptionTerm__c',
                    'metadata.Usage Type' => 'Product2.SBQQ__BillingType__c',
//            'unit_amount_decimal' => 'Unit_Price_STRIPE__c',
                    'metadata.Description' => 'Product2.Description',
                    'metadata.Interval Count' => 'Product2.SBQQ__BillingFrequency__c',
                    'metadata.Unit Amount Decimal' => 'UnitPrice'
            }),
            'subscription_item' => new ConfigData('subscription_item', new Map<String, String> {
                    'quantity' => 'SBQQ__OrderedQuantity__c',
//            'metadata.price' => 'Unit_Price_Stripe__c',
                    'metadata.TaxCode' => 'SBQQ__TaxCode__c',
                    'metadata.quantity' => 'SBQQ__OrderedQuantity__c',
                    'metadata.tax code' => 'SBQQ__TaxCode__c',
                    'metadata.service_date' => 'SBQQ__QuoteLine__c.SBQQ__Quote__c.SBQQ__StartDate__c',
//            'metadata.end_customer_id' => 'Order.Account.Account_Number__c',
                    'metadata.item_list_price' => 'SBQQ__QuoteLine__c.SBQQ__ListPrice__c',
                    'metadata.sales_order_date' => 'Order.CreatedDate',
                    'metadata.billing frequency' => 'SBQQ__BillingFrequency__c',
                    'metadata.end_customer_name' => 'Order.Account.Name',
                    'metadata.item_contract_value' => 'SBQQ__QuoteLine__c.SBQQ__NetPrice__c',
                    'metadata.sales_order_line_id' => 'Id',
                    'metadata.sfdc_sales_order_id' => 'Order.Id',
//            'metadata.contract_free_period' => 'Order.Free_Period_Duration__c',
                    'metadata.sfdc_sales_order_num' => 'Order.OrderNumber',
//            'metadata.contract_tfc_duration' => 'Order.TFCMonths__c',
                    'metadata.sales_order_line_sequence' => 'SBQQ__QuoteLine__c.SBQQ__Number__c'
            }),
            'subscription_schedule' => new ConfigData('subscription_schedule', new Map<String, String> {
                    'metadata.SFDC ID' => 'AccountId.Id',
                    'metadata.Order Number' => 'OrderNumber',
//            'metadata.Selling Method' => 'SBQQ__Quote__c.Selling_Method__c',
//            'metadata.Contract Number' => 'Original_Order_ID__c',
//            'metadata.Agreement number' => 'Agreement_Number__c',
//            'metadata.Subscription Type' => 'Stripe_Subscription_Type__c',
//            'metadata.Company Signed Date' => 'Company_Signed_Date__c',
//            'metadata.Other Party Signed Date' => 'Other_Party_Signed_Date__c',
//            'metadata.contract_cf_signed_date' => 'Contract_Signed_Date__c',
                    'metadata.Contract Service End Date' => 'EndDate',
                    'metadata.Contract Service Start Date' => 'EffectiveDate'
            })
    };

    static Test_Tools testTools = new Test_Tools();

    @IsTest
    static void TestConfigData() {
        RelationshipMapper rm = new RelationshipMapper();

        for (String key : testData.keySet()) {
            ConfigData data = testData.get(key);
            SObjectType t = RelationshipMapper.SupportedTypes.get(key);
            for (String field : data.data.keySet()) {
                String path = data.get(field);
                List<RelationshipMapper.LookupChain> chains = rm.getValidChains(t, path);
                String testSignature = '[' + key + '] ' + field + ' => ' + path;
                System.assertEquals(false, chains.isEmpty(), 'Has at least one valid chain for ' +testSignature);
            }
        }
    }

    @IsTest
    static void TestLocalFieldLookupChainItems() {
        RelationshipMapper rm = new RelationshipMapper();

        SObjectType t = RelationshipMapper.SupportedTypes.get('price');
        String path = 'UnitPrice';

        Test.startTest();

        List<RelationshipMapper.LookupChainItem> items = rm.getLookupChainItems(t, path);

        Test.stopTest();

        System.assertEquals(1, items.size(), 'has one item');

        RelationshipMapper.LookupChainItem item = items.get(0);

        System.assertEquals(t, item.root);
        System.assertEquals(path, item.path);
        System.assertEquals(path, item.field);
        System.assertEquals(null, item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);
    }

    @IsTest
    static void TestRelationshipFieldLookupChainItems() {
        RelationshipMapper rm = new RelationshipMapper();

        SObjectType t = RelationshipMapper.SupportedTypes.get('price');
        String path = 'Product2.Name';

        Test.startTest();

        List<RelationshipMapper.LookupChainItem> items = rm.getLookupChainItems(t, path);

        Test.stopTest();

        System.assertEquals(2, items.size(), 'has two items');

        RelationshipMapper.LookupChainItem item = items.get(0);

        System.assertEquals(t, item.root);
        System.assertEquals('Product2.Name', item.path);
        System.assertEquals('Product2Id', item.field);
        System.assertEquals('Product2', item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);

        item = items.get(1);

        System.assertEquals(Product2.SObjectType, item.root);
        System.assertEquals('Name', item.path);
        System.assertEquals('Name', item.field);
        System.assertEquals(null, item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);
    }

    @IsTest
    static void TestRelationshipIdFieldLookupChainItems() {
        RelationshipMapper rm = new RelationshipMapper();

        SObjectType t = RelationshipMapper.SupportedTypes.get('price');
        String path = 'Product2Id.Name';

        Test.startTest();

        List<RelationshipMapper.LookupChainItem> items = rm.getLookupChainItems(t, path);

        Test.stopTest();

        System.assertEquals(2, items.size(), 'has two items');

        RelationshipMapper.LookupChainItem item = items.get(0);

        System.assertEquals(t, item.root);
        System.assertEquals('Product2Id.Name', item.path);
        System.assertEquals('Product2Id', item.field);
        System.assertEquals('Product2', item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);

        item = items.get(1);

        System.assertEquals(Product2.SObjectType, item.root);
        System.assertEquals('Name', item.path);
        System.assertEquals('Name', item.field);
        System.assertEquals(null, item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);
    }

    @IsTest
    static void TestDeepRelationshipFieldLookupChainItems() {
        RelationshipMapper rm = new RelationshipMapper();

        SObjectType t = RelationshipMapper.SupportedTypes.get('price');
        String path = 'Product2.CreatedBy.Name';

        Test.startTest();

        List<RelationshipMapper.LookupChainItem> items = rm.getLookupChainItems(t, path);

        Test.stopTest();

        System.assertEquals(3, items.size(), 'has three items');

        RelationshipMapper.LookupChainItem item = items.get(0);

        System.assertEquals(t, item.root);
        System.assertEquals('Product2.CreatedBy.Name', item.path);
        System.assertEquals('Product2Id', item.field);
        System.assertEquals('Product2', item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);

        item = items.get(1);

        System.assertEquals(Product2.SObjectType, item.root);
        System.assertEquals('CreatedBy.Name', item.path);
        System.assertEquals('CreatedById', item.field);
        System.assertEquals('CreatedBy', item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);

        item = items.get(2);

        System.assertEquals(User.SObjectType, item.root);
        System.assertEquals('Name', item.path);
        System.assertEquals('Name', item.field);
        System.assertEquals(null, item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);
    }

    @IsTest
    static void TestDeepRelationshipIdFieldLookupChainItems() {
        RelationshipMapper rm = new RelationshipMapper();

        SObjectType t = RelationshipMapper.SupportedTypes.get('price');
        String path = 'Product2Id.CreatedById.Name';

        Test.startTest();

        List<RelationshipMapper.LookupChainItem> items = rm.getLookupChainItems(t, path);

        Test.stopTest();

        System.assertEquals(3, items.size(), 'has three items');

        RelationshipMapper.LookupChainItem item = items.get(0);

        System.assertEquals(t, item.root);
        System.assertEquals('Product2Id.CreatedById.Name', item.path);
        System.assertEquals('Product2Id', item.field);
        System.assertEquals('Product2', item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);

        item = items.get(1);

        System.assertEquals(Product2.SObjectType, item.root);
        System.assertEquals('CreatedById.Name', item.path);
        System.assertEquals('CreatedById', item.field);
        System.assertEquals('CreatedBy', item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);

        item = items.get(2);

        System.assertEquals(User.SObjectType, item.root);
        System.assertEquals('Name', item.path);
        System.assertEquals('Name', item.field);
        System.assertEquals(null, item.relationshipName);
        System.assertNotEquals(null, item.fieldDesc);
    }

    @IsTest
    static void TestPolymorphicDetectionOnLookupChainItems() {
        RelationshipMapper rm = new RelationshipMapper();

        List<RelationshipMapper.LookupChainItem> polyItems = rm.getLookupChainItems(Case.SObjectType, 'OwnerId.Name');
        List<RelationshipMapper.LookupChainItem> nonPolyItems = rm.getLookupChainItems(Case.SObjectType, 'CreatedById.Name');

        RelationshipMapper.LookupChainItem polyItem = polyItems.get(0);
        RelationshipMapper.LookupChainItem nonPolyItem = nonPolyItems.get(0);

        System.assertEquals(true, polyItem.isPolymorphicReference, 'Polymorphic reference is counted as such');
        System.assertEquals(false, nonPolyItem.isPolymorphicReference, 'Non-polymorphic reference is counted as such');
    }

    @IsTest
    static void TestDeepRelationshipIdFieldLookupChains() {
        RelationshipMapper rm = new RelationshipMapper();

        SObjectType t = RelationshipMapper.SupportedTypes.get('price');
        String path = 'Product2Id.CreatedById.Name';

        Test.startTest();

        List<RelationshipMapper.LookupChain> chains = rm.getLookupChains(t, path);

        Test.stopTest();

        System.assertEquals(1, chains.size(), 'should only contain one chain');

        RelationshipMapper.LookupChain chain = chains.get(0);

        System.assertEquals(true, chain.isValid, 'should be valid');
    }

    @IsTest
    static void TestDeepInvalidRelationshipIdFieldLookupChains() {
        RelationshipMapper rm = new RelationshipMapper();

        SObjectType t = RelationshipMapper.SupportedTypes.get('price');
        String path = 'Product2Id.CreatedById.Nomo';

        Test.startTest();

        List<RelationshipMapper.LookupChain> chains = rm.getLookupChains(t, path);

        Test.stopTest();

        System.assertEquals(1, chains.size(), 'should only contain one chain');

        RelationshipMapper.LookupChain chain = chains.get(0);

        System.assertEquals(false, chain.isValid, 'should not be valid');
    }

    @IsTest
    static void TestMultiPathLookupChains() {
        RelationshipMapper rm = new RelationshipMapper();

        String path = 'OwnerId.Name';

        Test.startTest();

        List<RelationshipMapper.LookupChain> chains = rm.getLookupChains(Case.SObjectType, path);

        Test.stopTest();

        System.assertEquals(2, chains.size(), 'should only contain two chains');

        SObjectType firstType = Case.SObjectType;
        SObjectType lastType = null;

        for (Integer i = 0; i < chains.size(); i++) {
            RelationshipMapper.LookupChain chain = chains.get(i);
            chain.debug();
            System.assertEquals(true, chain.isValid, 'should be valid on chain ' + i);
            System.assertEquals(2, chain.items.size(), 'should contain two items on chain ' + i);
            RelationshipMapper.LookupChainItem firstItem = chain.items.get(0);
            System.assertEquals(firstType, firstItem.root, 'first item root is Case on chain ' + i);

            RelationshipMapper.LookupChainItem secondItem = chain.items.get(1);
            System.assertNotEquals(lastType, secondItem.root, 'second item is different root than other second item on chain ' + i);
            lastType = secondItem.root;
        }
    }

    @IsTest
    static void TestMultiPathLookupChainsWithFailure() {
        RelationshipMapper rm = new RelationshipMapper();

        String path = 'OwnerId.FirstName';

        Test.startTest();

        List<RelationshipMapper.LookupChain> chains = rm.getLookupChains(Case.SObjectType, path);

        Test.stopTest();

        System.assertEquals(2, chains.size(), 'should only contain two chains');

        SObjectType firstType = Case.SObjectType;
        SObjectType lastType = null;

        for (Integer i = 0; i < chains.size(); i++) {
            RelationshipMapper.LookupChain chain = chains.get(i);
            chain.debug();
            System.assertEquals(2, chain.items.size(), 'should contain two items on chain ' + i);
            RelationshipMapper.LookupChainItem firstItem = chain.items.get(0);
            System.assertEquals(firstType, firstItem.root, 'first item root is Case on chain ' + i);

            RelationshipMapper.LookupChainItem secondItem = chain.items.get(1);
            System.assertNotEquals(lastType, secondItem.root, 'second item is different root than other second item on chain ' + i);
            lastType = secondItem.root;

            if (secondItem.root == User.SObjectType) {
                System.assertEquals(true, secondItem.isValid, 'User field reference is valid on chain ' + i);
            } else if (secondItem.root == Group.SObjectType) {
                System.assertEquals(false, secondItem.isValid, 'Group field reference is invalid on chain ' + i);
            } else {
                System.assert(false, 'Unexpected type ' + secondItem.root + ' found on chain ' + i);
            }

            System.assertEquals(secondItem.isValid, chain.isValid, 'chain inherits is valid value on chain ' + i);
        }
    }

    @IsTest
    static void TestGetValidChains() {
        RelationshipMapper rm = new RelationshipMapper();

        String path = 'OwnerId.FirstName';
        List<RelationshipMapper.LookupChain> chains = rm.getLookupChains(Case.SObjectType, path);

        Test.startTest();

        List<RelationshipMapper.LookupChain> validChains = rm.getValidChains(chains);

        Test.stopTest();

        System.assertEquals(1, validChains.size(), 'should only contain 1 chain');

        RelationshipMapper.LookupChain chain = validChains.get(0);

        System.assertEquals(true, chain.isValid, 'chain is valid');
    }

    @IsTest
    static void TestAccessibility() {
        RelationshipMapper rm = new RelationshipMapper();
        String profile = Test_Tools.PROFILE_CHATTER_FREE_USER;
        String uniqueId = 'TestRMAccessibility';
        User unprivUser = testTools.createUser(profile, uniqueId);
        String path = 'Name';

        Test.startTest();

        List<RelationshipMapper.LookupChain> accessibleChains = rm.getLookupChains(Account.SObjectType, path);
        List<RelationshipMapper.LookupChain> inaccessibleChains = null;

        System.runAs(unprivUser) {
            inaccessibleChains = rm.getLookupChains(Account.SObjectType, path);
        }

        Test.stopTest();

        System.assertEquals(1, accessibleChains.size(), 'has only one accessible chain');
        System.assertEquals(1, inaccessibleChains.size(), 'has only one inaccessible chain');

        RelationshipMapper.LookupChain accessibleChain = accessibleChains.get(0);
        RelationshipMapper.LookupChain inaccessibleChain = inaccessibleChains.get(0);

        System.assertEquals(true, accessibleChain.isAccessible(), 'accessible chain is accessible');
        System.assertEquals(false, inaccessibleChain.isAccessible(), 'inaccessible chain is inaccessible');
    }

    @IsTest
    static void TestSoqlPath() {
        RelationshipMapper rm = new RelationshipMapper();

        SObjectType t = RelationshipMapper.SupportedTypes.get('price');

        Map<String, String> tests = new Map<String, String> {
            'Product2.CreatedBy.Name' => 'Product2.CreatedBy.Name',
            'Product2.CreatedById.Name' => 'Product2.CreatedBy.Name',
            'Product2Id.CreatedBy.Name' => 'Product2.CreatedBy.Name',
            'Product2Id.CreatedById.Name' => 'Product2.CreatedBy.Name',
            'Product2.CreatedById' => 'Product2.CreatedById',
            'Product2Id.CreatedById' => 'Product2.CreatedById',
            'Product2Id' => 'Product2Id'
        };

        Map<String, RelationshipMapper.LookupChain> results = new Map<String, RelationshipMapper.LookupChain>();

        Test.startTest();

        List<RelationshipMapper.LookupChain> items;

        for (String path : tests.keySet()) {
            items = rm.getLookupChains(t, path);

            System.assert(items.size() > 0, 'items found for "' + path + '" should not be empty');

            results.put(path, items.get(0));
        }

        Test.stopTest();

        System.assertEquals(tests.size(), results.size(), 'Should have a result for every test');

        for (String path : results.keySet()) {
            String expected = tests.get(path);
            RelationshipMapper.LookupChain result = results.get(path);
            String resultPath = result.getSOQLPath();
            System.assertEquals(expected, resultPath, 'Path "' + path + '" should resolve to "' + expected + '"');
        }
    }

    class ConfigData {
        public Map<String, String> data = null;
        public String name = null;

        public ConfigData(String name) {
            this.name = name;
            this.data = new Map<String, String>();
        }

        public ConfigData(String name, Map<String, String> data) {
            this.name = name;
            this.data = data;
        }

        public String get(String key) {
            return data.get(key);
        }
    }
}