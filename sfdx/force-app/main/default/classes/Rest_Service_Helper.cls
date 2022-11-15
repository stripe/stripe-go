/**
 * Created by jmather-c on 11/7/22.
 */

public with sharing class Rest_Service_Helper {
    public String fieldsToQueryString(Set<String> s) {
        return String.join((Iterable<String>)s, ',');
    }

    public String setToQueryString(Set<String> s) {
        return '(\'' + String.join((Iterable<String>)s, '\',\'') + '\')';
    }

    //get a set off all accessible fields to an object based on the objects api name passed in
    public Set<String> getAccessibleFields(String objectApiName, Set<Object> responseJsonPayloadErrors) {
        Set<String> accessibleFields = new Set<String>();
        //create error if we dont have access to this object

        if (Schema.describeSObjects(new List<String>{ objectApiName })[0].isAccessible()) {
            Schema.DescribeSObjectResult describedObj = Schema.describeSObjects(new List<String>{objectApiName}, SObjectDescribeOptions.DEFERRED).get(0);
            for(Schema.SObjectField fieldSchema : describedObj.fields.getMap().values()) {
                Schema.DescribeFieldResult describedField = fieldSchema.getDescribe();

                //create an error in the rest errors list if a field is not accessible
                if(!describedField.isAccessible()) {
                    Map<String,Object> errorResponseMap = new Map<String,Object>{
                            'error_type' => 'FieldNotAccessible',
                            'error_message' => 'Field Not Accessible',
                            'sobject_type' => objectApiName,
                            'sobject_field' => describedField.getName()
                    };
                    responseJsonPayloadErrors.add((Object)errorResponseMap);
                    continue;
                }
                //add all accessible fields to a list
                accessibleFields.add(describedField.getName());
            }
        } else {
            Map<String,Object> errorResponseMap = new Map<String,Object>{
                    'error_type' => 'ObjectNotAccessible',
                    'error_message' => 'Object Not Accessible',
                    'sobject_type' => objectApiName,
                    'sobject_field' => null
            };
            responseJsonPayloadErrors.add((Object)errorResponseMap);
        }
        return accessibleFields;
    }

}