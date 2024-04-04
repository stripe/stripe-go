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
    public static final String V2_ACCOUNT_CONFIG = 'account_config';
    public static final String V2_INPUT_SETTINGS = 'connector_settings';
    public static final String ENABLED = 'enabled';
    public static final String INPUT_HIDDEN_SYNC_PREFS_FIELDS = 'hidden_sync_pref_fields';
    public static final String OUTPUT_DEFAULT_CURRENCY = 'default_currency';
    public static final String OUTPUT_IS_CPQ_INSTALLED = 'isCpqInstalled';
    public static final String OUTPUT_HIDDEN_SYNC_PREFS_FIELDS = 'hiddenSyncPrefsFields';
    public static final String SYNC_RECORD_RETENTION = 'sync_record_retention';
    public static final String SYNC_START_DATE = 'sync_start_date';
    public static final String API_PERCENTAGE_LIMIT = 'api_percentage_limit';
    public static final String CPQ_TERM_UNIT = 'cpq_term_unit';
    public static final String CPQ_PRORATE_PRECISION = 'cpq_prorate_precision';
    public static final String POLLING_ENABLED = 'polling_enabled';
    public static final String STRIPE_ACCOUNT_ID = 'stripe_account_id';
    public static final String V2_MERCHANT_ID = 'merchant';
    public static final String LAST_SYNCED = 'last_synced';
    public static final String DEFAULT_CURRENCY = 'default_currency';



    public static final List<String> SETTINGS_OBJECT_FIELDS = new List<String> {
            FIELD_MAPPINGS,
            FIELD_DEFAULTS
    };

    public static final List<String> SETTINGS_STRING_FIELDS = new List<String> {
            SYNC_RECORD_RETENTION,
            SYNC_START_DATE,
            API_PERCENTAGE_LIMIT,
            CPQ_TERM_UNIT,
            CPQ_PRORATE_PRECISION
    };

    public static final List<String> V2_SETTINGS_STRING_FIELDS = new List<String> {
            CPQ_TERM_UNIT,
            CPQ_PRORATE_PRECISION
    };

    public static final List<String> V2_SETTINGS_NUMERIC_FIELDS = new List<String> {
            SYNC_RECORD_RETENTION,
            SYNC_START_DATE,
            API_PERCENTAGE_LIMIT
    };

    public static final List<String> SETTINGS_BOOLEAN_FIELDS = new List<String> {
            POLLING_ENABLED
    };

    public static final List<String> CONNECTION_READ_ONLY_FIELDS = new List<String> {
            STRIPE_ACCOUNT_ID,
            LAST_SYNCED
    };

    public static final List<String> V2_CONNECTION_READ_ONLY_FIELDS = new List<String> {
            V2_MERCHANT_ID,
            LAST_SYNCED
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
        String field = (utilities.isUsingV2Middleware()) ? V2_INPUT_SETTINGS : INPUT_SETTINGS;
        Map<String, Object> settings = (Map<String,Object>)configPayload.get(field);
        Map<String, Object> filters = (Map<String,Object>)settings.get(INPUT_FILTERS);

        for (String key : IOMap_FilterFields.keySet()) {
            response.put(IOMap_FilterFields.get(key), String.valueOf(filters.get(key)));
        }
    }

    public void extractSettingsData(Map<String, Object> configPayload, responseData response, Boolean showInternalFields) {
        List<String> hiddenSyncPrefFields = getHiddenSyncPrefFields();
        Map<String, Object> settings = (Map<String,Object>)configPayload.get(INPUT_SETTINGS);

        //gets non read only settings fields from response for frontend
        for (String stringField : SETTINGS_STRING_FIELDS) {
            debugger.debug('extractSettingsData: stringField', stringField);
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

        response.put(OUTPUT_DEFAULT_CURRENCY, UserInfo.getDefaultCurrency());
        response.put(OUTPUT_IS_CPQ_INSTALLED, utilities.isCpqEnabled());

        //gets read only connection fields from response for frontend
        Map<String, Object> connectionStatus = (Map<String,Object>)configPayload.get(INPUT_CONNECTION_STATUS);
        for(String readOnlyField : CONNECTION_READ_ONLY_FIELDS) {
            response.put(readOnlyField, String.valueOf(connectionStatus.get(readOnlyField)));
        }

        if (showInternalFields) {
            String mappedKey = IOMap_Config.get(INPUT_HIDDEN_SYNC_PREFS_FIELDS);
            response.put(mappedKey, hiddenSyncPrefFields);
            response.put(ENABLED, (Boolean)configPayload.get(ENABLED));
        }
    }

    public void extractV2SettingsData(Map<String, Object> configPayload, responseData response, Boolean showInternalFields) {
        List<String> hiddenSyncPrefFields = getHiddenSyncPrefFields();
        Map<String, Object> settings = (Map<String,Object>)configPayload.get(V2_INPUT_SETTINGS);

        //gets non read only settings fields from response for frontend
        for (String stringField : V2_SETTINGS_STRING_FIELDS) {
            debugger.debug('extractV2SettingsData: stringField', stringField);
            // skip hidden field values
            if (hiddenSyncPrefFields.contains(stringField)) {
                continue;
            }
            response.put(stringField, String.valueOf(settings.get(stringField)));
        }

        for (String numericField : V2_SETTINGS_NUMERIC_FIELDS) {
            debugger.debug('extractV2SettingsData: numericField', numericField);
            // skip hidden field values
            if (hiddenSyncPrefFields.contains(numericField)) {
                continue;
            }
            response.put(numericField, Integer.valueOf(settings.get(numericField)));
        }

        //gets boolean connection fields from response for frontend
        for (String booleanField : SETTINGS_BOOLEAN_FIELDS) {
            if (settings.get(booleanField) == null) {
                response.put(booleanField, false);
            } else {
                response.put(booleanField, Boolean.valueOf(settings.get(booleanField)));
            }
        }

        response.put(OUTPUT_DEFAULT_CURRENCY, UserInfo.getDefaultCurrency());
        response.put(OUTPUT_IS_CPQ_INSTALLED, utilities.isCpqEnabled());

        //gets read only connection fields from response for frontend
        for(String readOnlyField : V2_CONNECTION_READ_ONLY_FIELDS) {
            // map the merchant id over to the expected name...
            if (readOnlyField.equals(V2_MERCHANT_ID)) {
                List<String> li = new List<String> {};
                List<Stripe_Account__c> stripeAccountList = setupAssistant.getStripeAccounts();
                for (Stripe_Account__c acct: stripeAccountList) {
                    li.add(acct.Stripe_Id__c);
                }
                response.put(STRIPE_ACCOUNT_ID, li);
                continue;
            }
            response.put(readOnlyField, String.valueOf(configPayload.get(readOnlyField)));
        }

        if (showInternalFields) {
            String mappedKey = IOMap_Config.get(INPUT_HIDDEN_SYNC_PREFS_FIELDS);
            response.put(mappedKey, hiddenSyncPrefFields);
            response.put(ENABLED, (Boolean)configPayload.get(ENABLED));
        }
    }

    public void extractFieldMappingData(Map<String, Object> configPayload, responseData response, Boolean showInternalFields) {
        debugger.debug('extractFieldMappingData: configPayload', configPayload);
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
        debugger.debug('stripHiddenMapperFields: allKeys', allKeys);

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
    private List<String> getHiddenSyncPrefFields() {
        return new List<String> { 'api_percentage_limit', 'cpq_term_unit' };
    }

    @TestVisible
    private List<String> getHiddenMapperFields(Map<String, Object> configPayload) {
        List<String> hiddenMapperFields = getListFieldFromPayload(INPUT_HIDDEN_MAPPER_FIELDS, configPayload, false);
        if (hiddenMapperFields == null) {
            return new List<String> { };
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

    public Boolean validateConfigHash(Map<String, Object> requestPayload, Map<String, Object> savePayload, ErrorTracker errors) {
        // this is static...
        debugger.debug('validateConfigHash requestPayload', requestPayload);

        if (requestPayload.containsKey(OUTPUT_CONFIG_HASH) == false) {
            errors.add(OUTPUT_CONFIG_HASH, 'No value sent for ' + OUTPUT_CONFIG_HASH);
            return true;
        }

        savePayload.put(INPUT_CONFIG_HASH, requestPayload.get(OUTPUT_CONFIG_HASH));
        debugger.debug('validateConfigHash savePayload', savePayload);
        return false;
    }

    public Boolean validateDataMappingConfig(Map<String, Object> requestPayload, Map<String, Object> savePayload, ErrorTracker errors) {
        List<String> required_fields = new List<String> {
                FIELD_DEFAULTS,
                FIELD_MAPPINGS
        };

        if (requestPayload.containsKey(OUTPUT_ALL_MAPPING_CONFIGS) == false) {
            errors.add(OUTPUT_ALL_MAPPING_CONFIGS, 'No value sent for ' + OUTPUT_ALL_MAPPING_CONFIGS);
            return true;
        }

        debugger.debug('validateConfigHash requestPayload', requestPayload);

        Map<String, Object> allMappingConfigurations = (Map<String, Object>) requestPayload.get(OUTPUT_ALL_MAPPING_CONFIGS);

        return validateRequiredFields(required_fields, allMappingConfigurations, savePayload, errors);
    }

    public Boolean validateSyncPreferences(Map<String, Object> requestPayload, Map<String, Object> savePayload, ErrorTracker errors) {
        List<String> required_fields = new List<String> {
                POLLING_ENABLED,
                SYNC_RECORD_RETENTION
        };

        List<String> optional_fields = new List<String> {
                SYNC_START_DATE,
                API_PERCENTAGE_LIMIT,
                CPQ_TERM_UNIT,
                CPQ_PRORATE_PRECISION
        };

        Map<String, Object> settings = (Map<String, Object>) savePayload.get('settings');
        if (settings == null) {
            settings = new Map<String, Object>();
        }

        settings.put(DEFAULT_CURRENCY, UserInfo.getDefaultCurrency());

        Boolean requiredError = validateRequiredFields(required_fields, requestPayload, settings, errors);
        Boolean optionalError = validateOptionalFields(optional_fields, requestPayload, settings, errors);

        debugger.debug('validateSyncPreferences settings', settings);
        savePayload.put('settings', settings);
        return requiredError || optionalError;
    }

    public Boolean validateRequiredFields(List<String> fields, Map<String, Object> requestPayload, Map<String, Object> saveTo, ErrorTracker errors) {
        Boolean hasError = false;

        for (String key : fields) {
            if (requestPayload.containsKey(key) == false) {
                hasError = true;
                errors.add(key, 'No value sent for ' + key);
            } else {
                hasError = hasError || extractValue(key, requestPayload, saveTo, errors);
            }
        }

        return hasError;
    }

    public Boolean validateOptionalFields(List<String> fields, Map<String, Object> requestPayload, Map<String, Object> saveTo, ErrorTracker errors) {
        Boolean hasError = false;

        for (String key : fields) {
            if (requestPayload.containsKey(key)) {
                hasError = hasError || extractValue(key, requestPayload, saveTo, errors);
            }
        }

        return hasError;
    }

    public Boolean extractValue(String key, Map<String, Object> requestPayload, Map<String, Object> saveTo, ErrorTracker errors) {
        if (utilities.isUsingV1Middleware()) {
            return extractValueV1(key, requestPayload, saveTo, errors);
        }

        return extractValueV2(key, requestPayload, saveTo, errors);
    }

    public Boolean extractValueV1(String key, Map<String, Object> requestPayload, Map<String, Object> saveTo, ErrorTracker errors) {
        if (SETTINGS_STRING_FIELDS.contains(key)) {
            saveTo.put(key, String.valueOf(requestPayload.get(key)));
        } else if (SETTINGS_BOOLEAN_FIELDS.contains(key)) {
            saveTo.put(key, Boolean.valueOf(requestPayload.get(key)));
        } else if (SETTINGS_OBJECT_FIELDS.contains(key)) {
            saveTo.put(key, requestPayload.get(key));
        } else {
            errors.add(key, 'No definition of how to translate ' + key);
            return true;
        }

        return false;
    }

    public Boolean extractValueV2(String key, Map<String, Object> requestPayload, Map<String, Object> saveTo, ErrorTracker errors) {
        if (V2_SETTINGS_STRING_FIELDS.contains(key)) {
            saveTo.put(key, String.valueOf(requestPayload.get(key)));
        } else if (V2_SETTINGS_NUMERIC_FIELDS.contains(key)) {
            saveTo.put(key, Integer.valueOf(requestPayload.get(key)));
        } else if (SETTINGS_BOOLEAN_FIELDS.contains(key)) {
            saveTo.put(key, Boolean.valueOf(requestPayload.get(key)));
        } else if (SETTINGS_OBJECT_FIELDS.contains(key)) {
            saveTo.put(key, requestPayload.get(key));
        } else {
            errors.add(key, 'No definition of how to translate ' + key);
            return true;
        }

        return false;
    }

    public Boolean validateFilterConfig(Map<String, Object> requestPayload, Map<String, Object> savePayload, ErrorTracker errors) {
        Boolean hadErrors = false;
        Map<String, String> filters = new Map<String, String>();
        for (String objectName : IOMap_FilterFields.keySet()) {
            try {
                String payloadPropertyName = IOMap_FilterFields.get(objectName);
                String value = String.valueOf(requestPayload.get(payloadPropertyName));
                if (String.isBlank(value)) {
                    value = null;
                }

                filters.put(objectName, value);

                if (value == null) {
                    continue;
                }

                // this is human input, so let's try to clean it a little...
                value = value.trim();

                /*
                Using limit zero in queries so there are no records are returned therefore it
                will not count against the total number of rows retrieved salesforce limitation.
                */
                String query = 'SELECT Id FROM ' + objectName + ' WHERE ' + value + ' WITH SECURITY_ENFORCED LIMIT 0';
                Database.query(query);
            } catch (Exception e) {
                hadErrors = true;
                errors.add(objectName, 'Encountered error processing filter: ' + e.getMessage());
                errorLogger.create('processing filters: ' + objectName, e);
            }
        }

        Map<String, Object> settings = (Map<String, Object>) savePayload.get('settings');
        if (settings == null) {
            settings = new Map<String, Object>();
        }
        settings.put('filters', filters);
        return hadErrors;
    }

    public class ErrorTracker {
        public final List<Map<String, Object>> errors = new List<Map<String, Object>>();

        public List<Map<String, Object>> getErrors() {
            return this.errors;
        }

        public Map<String, Object> add(String target, String msg) {
            Map<String, Object> err = new Map<String, Object> {
                    'Object' => target,
                    'Error' => msg
            };
            this.errors.add(err);
            return err;
        }
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