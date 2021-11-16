/*** Strike DML Manager v2.0.0 ***/
public with sharing class dmlManager {
    public class accessException extends Exception {}

    public static Map<String, Schema.SObjectType> globalDescribe = Schema.getGlobalDescribe();

    // checkAccess example param
    //
    // Map<String, Map<String, Set<String>>> accessMapByObjectName = new Map<String, Map<String, Set<String>>> {
    //     'Account' => new Map<String, Set<String>> {
    //         'Create' => new Set<String> {
    //             'Name'
    //         },
    //         'Read' => new Set<String> {
    //             'Id', 'Name'
    //         },
    //         'Update' => new Set<String> {
    //             'Name'
    //         },
    //         'Delete' => new Set<String>()
    //     },
    //     'Contact' => new Map<String, Set<String>> {
    //         'Read' => new Set<String> {
    //             'Id', 'Name'
    //         }
    //     }
    // }

    public static void checkAccess(Map<String, Map<String, Set<String>>> accessMapByObjectName) {
        Set<String> objectNameSet = accessMapByObjectName.keySet();

        for (String objectName : objectNameSet) {
            checkAccess(accessMapByObjectName, objectName);
        }
    }

    public static void checkAccess(Map<String, Map<String, Set<String>>> accessMapByObjectName, String objectName) {
        Map<String, Set<String>> fieldNamesByAccess = accessMapByObjectName.get(objectName);
        Set<String> accessSet = fieldNamesByAccess.keySet();

        for (String access : accessSet) {
            Set<String> fieldNames = fieldNamesByAccess.get(access);

            switch on access {
                when 'Create', 'Read', 'Update', 'Delete' {
                    check(access, objectName, fieldNames);
                }
                when else {
                    dmlManager.throwException(String.format('Unknown access "{0}"', new List<String>{ access }));
                }
            }
        }
    }

    private static void check(String access, String objectName, Set<String> fieldNames) {
        DescribeSObjectResult mSObjectDescribe = getSObjectDescribeByObjectName(objectName);
        Boolean hasObjectAccess = false;

        switch on access {
            when 'Create' {
                hasObjectAccess = mSObjectDescribe.isCreateable();
            }
            when 'Read' {
                hasObjectAccess = mSObjectDescribe.isAccessible();
            }
            when 'Update' {
                hasObjectAccess = mSObjectDescribe.isUpdateable();
            }
            when 'Delete' {
                hasObjectAccess = mSObjectDescribe.isDeletable();
            }
        }

        if (!hasObjectAccess) {
            dmlManager.throwException(String.format('Insufficient "{0}" access on object "{1}"', new List<String>{ access, mSObjectDescribe.getLabel() }));
        }

        if (null != fieldNames) {
            for (String fieldName : fieldNames) {
                Schema.DescribeFieldResult mSObjectFieldDescribe = getFieldDescribeByFieldName(fieldName, mSObjectDescribe);
                Boolean hasFieldAccess = false;

                switch on access {
                    when 'Create' {
                        hasFieldAccess = mSObjectFieldDescribe.isCreateable();
                    }
                    when 'Read' {
                        hasFieldAccess = mSObjectFieldDescribe.isAccessible();
                    }
                    when 'Update' {
                        hasFieldAccess = mSObjectFieldDescribe.isUpdateable();
                    }
                }

                if (!hasFieldAccess) {
                    dmlManager.throwException(String.format('Insufficient "{0}" access on field "{1}" on object "{2}"', new List<String>{ access, mSObjectFieldDescribe.getLabel(), mSObjectDescribe.getLabel() }));
                }
            }
        }
    }

    private static DescribeSObjectResult getSObjectDescribeByObjectName(String objectName) {
        if (globalDescribe.containsKey(constants.NAMESPACE_API + objectName)) {
            return globalDescribe.get(constants.NAMESPACE_API + objectName).getDescribe();
        }

        if (!globalDescribe.containsKey(objectName)) {
            dmlManager.throwException(String.format('Unknown object "{0}"', new List<String> { objectName }));
        }

        return globalDescribe.get(objectName).getDescribe();
    }

    private static DescribeFieldResult getFieldDescribeByFieldName(String fieldName, DescribeSObjectResult mDescribeSObjectResult) {
        Map<String, Schema.SObjectField> mSObjectFieldDescribeByFieldName = mDescribeSObjectResult.fields.getMap();

        if (mSObjectFieldDescribeByFieldName.containsKey(constants.NAMESPACE_API + fieldName)) {
            return mSObjectFieldDescribeByFieldName.get(constants.NAMESPACE_API + fieldName).getDescribe();
        }

        if (!mSObjectFieldDescribeByFieldName.containsKey(fieldName)) {
            dmlManager.throwException(String.format('Unknown field "{0}" on object "{1}"', new List<String> { fieldName, mDescribeSObjectResult.getLabel() }));
        }

        return mSObjectFieldDescribeByFieldName.get(fieldName).getDescribe();
    }

    private static void throwException(String message) {
        throw new accessException(message);
    }
}