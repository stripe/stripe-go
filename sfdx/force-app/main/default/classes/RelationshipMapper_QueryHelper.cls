/**
 * Handles creating the query references to use for the results from the RelationshipMapper.
 */
public with sharing class RelationshipMapper_QueryHelper {
    public class QHException extends Sentry_Exception {}
    public static final String EXCEPTION_FIELD_INACCESSIBLE = 'Defined field mapping SOQL path is inaccessible.';
    public static final String EXCEPTION_POLY_RELATIONSHIP = 'References to custom fields with polymorphic types are not supported.';

    private Map<String, Set<String>> objectRelationshipFields = new Map<String, Set<String>>();
    private Map<String, Map<String, String>> objectFieldMappings = new Map<String, Map<String, String>>();

    private RelationshipMapper rm = new RelationshipMapper();

    enum Validity {
        ALL,
        SOME
    }

    // this maps from the stripe view of the metadata config to the salesforce view
    // as in: instead of customer/subscription we use Account/Order
    public void initializeRelationshipChains(Map<String, Map<String, String>> mappingData) {
        Map<SObjectType, Set<String>> reverseMapObjectNames = new Map<SObjectType, Set<String>>();
        Map<String, Map<String, String>> inaccessibleFields = new Map<String, Map<String, String>>();

        for (String stripeObj : RelationshipMapper.SupportedTypes.keySet()) {
            SObjectType objType = RelationshipMapper.SupportedTypes.get(stripeObj);

            if (reverseMapObjectNames.containsKey(objType)) {
                reverseMapObjectNames.get(objType).add(stripeObj);
            } else {
                reverseMapObjectNames.put(objType, new Set<String> { stripeObj });
            }
        }

        for (SObjectType objType : reverseMapObjectNames.keySet()) {
            String className = objType.getDescribe().getName();
            Set<String> fields = new Set<String>();
            objectRelationshipFields.put(className, fields);

            for (String stripeObj : reverseMapObjectNames.get(objType)) {
                if (mappingData.containsKey(stripeObj) == false) {
                    continue;
                }

                Map<String, String> fieldMappings = new Map<String, String>();
                objectFieldMappings.put(stripeObj, fieldMappings);
                Map<String, String> objMapping = mappingData.get(stripeObj);

                for (String field : objMapping.keySet()) {
                    String path = objMapping.get(field);
                    List<RelationshipMapper.LookupChain> foundChains = rm.getLookupChains(objType, path);

                    if (isSupportedResult(foundChains)) {
                        for (RelationshipMapper.LookupChain chain : foundChains) {
                            if (chain.isAccessible() == false) {
                                if (inaccessibleFields.containsKey(stripeObj)) {
                                    inaccessibleFields.get(stripeObj).put(field, chain.getSOQLPath());
                                } else {
                                    inaccessibleFields.put(stripeObj, new Map<String, String> {
                                        field => chain.getSOQLPath()
                                    });
                                }
                            }
                        }
                        String soqlPath = foundChains.get(0).getSOQLPath();
                        fields.add(soqlPath);
                        fieldMappings.put(field, soqlPath);
                    } else {
                        QHException ex = (QHException) Sentry_ExceptionFactory.build(QHException.class);
                        ex.setMessage(EXCEPTION_POLY_RELATIONSHIP);
                        ex.context.put('mapping_data', mappingData);
                        ex.context.put('stripe_obj', stripeObj);
                        ex.context.put('mapping_field', field);
                        ex.context.put('mapping_path', path);

                        throw ex;
                    }
                }
            }
        }

        if (inaccessibleFields.size() > 0) {
            QHException ex = (QHException) Sentry_ExceptionFactory.build(QHException.class);
            ex.setMessage(EXCEPTION_FIELD_INACCESSIBLE);
            ex.context.put('mapping_data', mappingData);
            ex.context.put('failed_fields', inaccessibleFields);
            throw ex;
        }
    }

    public void clear() {
        objectFieldMappings = new Map<String, Map<String, String>>();
        objectRelationshipFields = new Map<String, Set<String>>();
    }

    public Set<String> getRelatedFields(String objectName) {
        if (objectRelationshipFields.containsKey(objectName) == false) {
            return new Set<String>();
        }

        return objectRelationshipFields.get(objectName);
    }

    public Map<String, Map<String, String>> getFieldMappings() {
        return objectFieldMappings;
    }

    private Boolean isSupportedResult(List<RelationshipMapper.LookupChain> chains) {
        if (chains.isEmpty()) {
            return false;
        }

        Boolean allValid = areValid(chains, Validity.ALL);

        if (chains.size() < 2) {
            return allValid;
        }

        Boolean someValid = areValid(chains, Validity.SOME);
        Boolean hasPolyRef = hasPolymorphicReference(chains);

        if (hasPolyRef && someValid && allValid == false) {
            return false;
        }

        return allValid;
    }

    private Boolean hasPolymorphicReference(List<RelationshipMapper.LookupChain> chains) {
        for(RelationshipMapper.LookupChain chain : chains) {
            if (chain.hasPolymorphicReference()) {
                return true;
            }
        }

        return false;
    }

    private Boolean areValid(List<RelationshipMapper.LookupChain> chains, Validity requiredValidity) {
        Boolean isValid = false;
        for (RelationshipMapper.LookupChain chain : chains) {
            if (chain.isValid) {
                isValid = true;
            }

            if (chain.isValid == false && requiredValidity == Validity.ALL) {
                return false;
            }
        }

        return isValid;
    }
}