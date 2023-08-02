/**
 * Created by jmather-c on 4/26/23.
 */

public with sharing class ConfigPayloadHelper {
    @TestVisible
    Debug_Helper debugger = new Debug_Helper();

    /**
     * If a variable is prefixed with INPUT_ or OUTPUT_ it means
     * that it is only used for that portion of the translation.
     *
     * No prefix indicates that it is used in both the input and output.
     */

    public static final String INPUT_CONFIG_HASH = 'configuration_hash';
    public static final String OUTPUT_CONFIG_HASH = 'configurationHash';

    public static final String INPUT_CONNECTION_STATUS = 'connection_status';

    public static final String FIELD_MAPPINGS = 'field_mappings';
    public static final String DEFAULT_MAPPINGS = 'default_mappings';
    public static final String REQUIRED_MAPPINGS = 'required_mappings';
    public static final String FIELD_DEFAULTS = 'field_defaults';
    public static final String OUTPUT_ALL_MAPPING_CONFIGS = 'allMappingConfigurations';
    public static final List<String> MAPPING_KEYS = new List<String> {
            FIELD_MAPPINGS,
            DEFAULT_MAPPINGS,
            REQUIRED_MAPPINGS,
            FIELD_DEFAULTS
    };

    public static final String INPUT_HIDDEN_MAPPER_FIELDS = 'hidden_mapper_fields';
    public static final String OUTPUT_HIDDEN_MAPPER_FIELDS = 'hiddenMapperFields';

    public static final String INPUT_FILTERS = 'filters';
    public static final String INPUT_SETTINGS = 'settings';
    public static final String ENABLED = 'enabled';
    public static final String INPUT_HIDDEN_SYNC_PREFS_FIELDS = 'hidden_sync_pref_fields';
    public static final String OUTPUT_DEFAULT_CURRENCY = 'default_currency';
    public static final String OUTPUT_IS_CPQ_INSTALLED = 'isCpqInstalled';
    public static final String OUTPUT_HIDDEN_SYNC_PREFS_FIELDS = 'hiddenSyncPrefsFields';
    public static final String SYNC_RECORD_RETENTION = 'sync_record_retention';

    public static final List<String> SETTINGS_STRING_FIELDS = new List<String> {
            'sync_record_retention',
            'sync_start_date',
            'api_percentage_limit',
            'cpq_term_unit',
            'cpq_prorate_precision'
    };

    public static final List<String> SETTINGS_BOOLEAN_FIELDS = new List<String> {
            'polling_enabled'
    };

    public static final List<String> CONNECTION_READ_ONLY_FIELDS = new List<String> {
            'stripe_account_id',
            'last_synced'
    };

    Map<String, String> IOMap_FilterFields = new Map<String, String> {
            'Product2' => 'product_filter',
            'Order' => 'order_filter',
            'Account' => 'account_filter',
            'PricebookEntry' => 'pricebook_entry_filter'
    };

    Map<String, String> IOMap_Config = new Map<String, String> {
            INPUT_CONFIG_HASH => OUTPUT_CONFIG_HASH,
            INPUT_HIDDEN_SYNC_PREFS_FIELDS => OUTPUT_HIDDEN_SYNC_PREFS_FIELDS,
            INPUT_HIDDEN_MAPPER_FIELDS => OUTPUT_HIDDEN_MAPPER_FIELDS
    };

    public void extractConfigHash(Map<String, Object> configPayload, responseData response) {
        String mappedKey = IOMap_Config.get(INPUT_CONFIG_HASH);
        response.put(mappedKey, (String)configPayload.get(INPUT_CONFIG_HASH));
    }

    public Boolean extractConnectionStatus(Map<String, Object> configPayload, Stripe_Connection__c connection) {
        //get connection status values from response
        Map<String, Object> connectionStatus = (Map<String,Object>)configPayload.get(INPUT_CONNECTION_STATUS);
        connection.Stripe_Connected__c = (Boolean)connectionStatus.get(constants.SYSTEM_STRIPE);
        connection.Salesforce_Connected__c = (Boolean)connectionStatus.get(constants.SYSTEM_SALESFORCE);
        return connection.Salesforce_Connected__c && connection.Stripe_Connected__c;
    }

    public void extractFilterSettingsData(Map<String, Object> configPayload, responseData response) {
        Map<String, Object> settings = (Map<String,Object>)configPayload.get(INPUT_SETTINGS);
        Map<String, Object> filters = (Map<String,Object>)settings.get(INPUT_FILTERS);

        for (String key : IOMap_FilterFields.keySet()) {
            response.put(IOMap_FilterFields.get(key), String.valueOf(filters.get(key)));
        }
    }

    public void extractSettingsData(Map<String, Object> configPayload, responseData response, Boolean showInternalFields) {
        List<String> hiddenSyncPrefFields = getHiddenSyncPrefFields(configPayload);
        Map<String, Object> settings = (Map<String,Object>)configPayload.get(INPUT_SETTINGS);

        //gets non read only settings fields from response for frontend
        for (String stringField : SETTINGS_STRING_FIELDS) {
            // skip hidden field values
            if (hiddenSyncPrefFields.contains(stringField)) {
                continue;
            }
            response.put(stringField, String.valueOf(settings.get(stringField)));
        }

        //gets boolean connection fields from response for frontend
        for (String booleanField : SETTINGS_BOOLEAN_FIELDS) {
            if (settings.get(booleanField) == null) {
                response.put(booleanField, false);
            } else {
                response.put(booleanField, Boolean.valueOf(settings.get(booleanField)));
            }
        }

        if (showInternalFields) {
            response.put(OUTPUT_DEFAULT_CURRENCY, UserInfo.getDefaultCurrency());
            response.put(OUTPUT_IS_CPQ_INSTALLED, utilities.isCpqEnabled());

            //gets read only connection fields from response for frontend
            Map<String, Object> connectionStatus = (Map<String,Object>)configPayload.get(INPUT_CONNECTION_STATUS);
            for(String readOnlyField : CONNECTION_READ_ONLY_FIELDS) {
                response.put(readOnlyField, String.valueOf(connectionStatus.get(readOnlyField)));
            }

            String mappedKey = IOMap_Config.get(INPUT_HIDDEN_SYNC_PREFS_FIELDS);
            response.put(mappedKey, hiddenSyncPrefFields);
            response.put(ENABLED, (Boolean)configPayload.get(ENABLED));
        }
    }

    public void extractFieldMappingData(Map<String, Object> configPayload, responseData response, Boolean showInternalFields) {
        List<String> hiddenMapperFields = getHiddenMapperFields(configPayload);

        //used to store configuration maps from response
        Map<String, Object> allMappingConfigurations = new Map<String, Object>();

        Map<String, Object> valueMap;
        for (String key : MAPPING_KEYS) {
            valueMap = (Map<String,Object>)configPayload.get(key);
            if (valueMap != null && !valueMap.isEmpty()) {
                allMappingConfigurations.put(key, valueMap);
            }
        }

        HiddenMapperFieldAnalysis analysis = new HiddenMapperFieldAnalysis(hiddenMapperFields);
        stripHiddenMapperFields(analysis, (Map<String, Object>) allMappingConfigurations.get(FIELD_DEFAULTS));
        stripHiddenMapperFields(analysis, (Map<String, Object>) allMappingConfigurations.get(FIELD_MAPPINGS));

        if (showInternalFields) {
            String mappedKey = IOMap_Config.get(INPUT_HIDDEN_MAPPER_FIELDS);
            response.put(mappedKey, hiddenMapperFields);
        } else {
            allMappingConfigurations.remove(DEFAULT_MAPPINGS);
            allMappingConfigurations.remove(REQUIRED_MAPPINGS);
        }

        response.put(OUTPUT_ALL_MAPPING_CONFIGS, allMappingConfigurations);
    }

    /**
     * This method strips the fields defined by the analysis from the given data map.
     *
     * @param analysis the analysis object containing the processed lists of fields to remove
     * @param data the data to strip fields from
     */
    @TestVisible
    private void stripHiddenMapperFields(HiddenMapperFieldAnalysis analysis, Map<String, Object> data) {
        if (data == null) {
            return;
        }

        stripMapFields(data, analysis.topLevelRemovals);
        Set<String> allKeys = new Set<String>(analysis.subLevelRemovals.keySet());
        allKeys.addAll(analysis.partialLevelRemovals.keySet());

        for (String keyField : allKeys) {
            if (data.containsKey(keyField) == false) {
                continue;
            }

            Map<String, Object> subData = (Map<String, Object>) data.get(keyField);
            stripAllMapFields(subData, analysis.subLevelRemovals.get(keyField), analysis.partialLevelRemovals.get(keyField));
        }
    }

    /**
     * This just coordinates the calls to stripMapFields and stripPartialMapFields.
     *
     * @param data underlying object data
     * @param fields full field names
     * @param partials partial match field names
     */
    @TestVisible
    private void stripAllMapFields(Map<String, Object> data, List<String> fields, List<String> partials) {
        if (data == null) {
            return;
        }

        stripMapFields(data, fields);
        stripPartialMapFields(data, partials);
    }

    /**
     * Removes the specified fields from the data map.
     *
     * @param data underlying object data
     * @param fields full field names
     */
    @TestVisible
    private void stripMapFields(Map<String, Object> data, List<String> fields) {
        if (fields == null || fields.isEmpty() || data == null) {
            return;
        }

        for (String field : fields) {
            if (data.containsKey(field)) {
                data.remove(field);
            }
        }
    }

    /**
     * Removes the specified partial match fields from the data map.
     *
     * @param data underlying object data
     * @param partials partial match field names
     */
    @TestVisible
    private void stripPartialMapFields(Map<String, Object> data, List<String> partials) {
        if (partials == null || partials.isEmpty() || data == null) {
            return;
        }

        List<String> regexPartials = new List<String>();
        for (String partial : partials) {
            regexPartials.add(Pattern.quote(partial + '.'));
        }

        debugger.debug('ConfigPayloadHelper.stripPartialMapFields.regexPartials', regexPartials);

        String regex = '^(' + String.join(regexPartials, '|') + ')';

        debugger.debug('ConfigPayloadHelper.stripPartialMapFields.regex', regex);

        Pattern partialPattern = Pattern.compile(regex);
        for (String key : data.keySet()) {
            Boolean match = partialPattern.matcher(key).find();
            debugger.debug('ConfigPayloadHelper.stripPartialMapFields.check', key);
            debugger.debug('ConfigPayloadHelper.stripPartialMapFields.matches', match);

            if (match) {
                data.remove(key);
            }
        }
    }

    @TestVisible
    private List<String> getHiddenSyncPrefFields(Map<String, Object> configPayload) {
        List<String> hiddenSyncPrefFields = getListFieldFromPayload(INPUT_HIDDEN_SYNC_PREFS_FIELDS, configPayload, false);
        if (hiddenSyncPrefFields == null) {
            return new List<String> { 'api_percentage_limit', 'cpq_term_unit' };
        } else {
            return hiddenSyncPrefFields;
        }
    }

    @TestVisible
    private List<String> getHiddenMapperFields(Map<String, Object> configPayload) {
        List<String> hiddenMapperFields = getListFieldFromPayload(INPUT_HIDDEN_MAPPER_FIELDS, configPayload, false);
        if (hiddenMapperFields == null) {
            return new List<String> { 'coupon', 'subscription_schedule.prebilling', 'customer.standard.coupon' };
        } else {
            return hiddenMapperFields;
        }
    }

    public List<String> getListFieldFromPayload(String listKey, Map<String,Object> payload) {
        return getListFieldFromPayload(listKey, payload, true);
    }

    public List<String> getListFieldFromPayload(String listKey, Map<String,Object> payload, Boolean ensureList) {
        List<Object> respFields;
        List<String> respList = new List<String>();

        try {
            respFields = (List<Object>) payload.get(listKey);

            if (respFields == null) {
                return ensureList ? new List<String>() : null;
            }

            for (Object obj : respFields) {
                respList.add(String.valueOf(obj));
            }
        } catch (Exception e) {
            // data is not in the format we expect...
            return ensureList ? new List<String>() : null;
        }

        return respList;
    }

    @TestVisible
    class HiddenMapperFieldAnalysis {
        public List<String> topLevelRemovals = new List<String>();
        public Map<String, List<String>> partialLevelRemovals = new Map<String, List<String>>();
        public Map<String, List<String>> subLevelRemovals = new Map<String, List<String>>();

        /**
         * This class analyzes the hidden mapper fields and determines which fields, or patterns of fields,
         * need to be removed from the payload. The naming convention of the fields is a little sporadic, so
         * we need to be able to handle a few different scenarios. See the unit tests for examples.
         *
         * @param hiddenMapperFields the list of hidden mapper fields to analyze for permutations required for removal
         */
        public HiddenMapperFieldAnalysis(List<String> hiddenMapperFields) {
            for (String hiddenField : hiddenMapperFields) {
                if (hiddenField.contains('.')) {
                    Integer pos = hiddenField.indexOf('.');
                    String key = hiddenField.substring(0, pos);
                    String val = hiddenField.substring(pos + 1);

                    if (val.contains('.')) {
                        if (subLevelRemovals.containsKey(key)) {
                            subLevelRemovals.get(key).add(val);
                        } else {
                            subLevelRemovals.put(key, new List<String> { val });
                        }
                    } else {
                        if (partialLevelRemovals.containsKey(key)) {
                            partialLevelRemovals.get(key).add(val);
                        } else {
                            partialLevelRemovals.put(key, new List<String> { val });
                        }
                    }
                } else {
                    topLevelRemovals.add(hiddenField);
                }
            }
        }
    }
}