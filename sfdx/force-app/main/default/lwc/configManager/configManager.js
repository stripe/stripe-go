/**
 * Created by jmather-c on 10/3/23.
 */
import getCoreData from '@salesforce/apex/setupAssistant.getCoreData';
import getSetupData from '@salesforce/apex/setupAssistant.getSetupData';
import getDataMapperLayout from '@salesforce/apex/setupAssistant.getDataMapperLayout';
import getTranslationConfig from '@salesforce/apex/setupAssistant.getTranslationConfig';
import saveTranslationConfig from '@salesforce/apex/setupAssistant.saveTranslationConfig';
import EventEmitter from './eventEmitter';


/**
 * @typedef {Object} PicklistValuesForObject
 * @property {boolean} isFieldCustom
 * @property {boolean} isObjectCustom
 * @property {string} label The label of the field
 * @property {string} type The salesforce type of the object
 * @property {string} value The property on the object where the value is stored
 */

/**
 * @typedef {Object} MultiObjectPicklistValuesForMapper
 * @property {Array<PicklistValuesForObject>} Account
 * @property {Array<PicklistValuesForObject>} Order
 * @property {Array<PicklistValuesForObject>} OrderItem
 * @property {Array<PicklistValuesForObject>} Order_Stripe_Coupon__c
 * @property {Array<PicklistValuesForObject>} PricebookEntry
 * @property {Array<PicklistValuesForObject>} Product2
 */

/**
 * @typedef {ResponseDataResults} CoreDataResults
 * @property {string} packageVersion
 * @property {string} configDownloadUri
 * @property {boolean} isMultiCurrencyEnabled
 * @property {Array<{label: string, value: string}>} supportedISOCodes List of supported ISO Codes, with 'Select...' first.
 * @property {string} defaultCurrency
 * @property {MultiObjectPicklistValuesForMapper} fieldListByObjectMap
 * @property {boolean} isCpqInstalled
 * @property {boolean} isSandbox
 */

/**
 * @typedef {Object} MappingConfiguration
 * @property {Object<string, Object<string, string>>} default_mappings
 * @property {Object<string, Object<string, string>>} field_defaults
 * @property {Object<string, Object<string, string>>} field_mappings
 * @property {Object<string, Object<string, string>>} required_mappings
 */

/**
 * @typedef {Object} SaveMappingConfiguration
 * @property {Object<string, Object<string, string>>} field_defaults
 * @property {Object<string, Object<string, string>>} field_mappings
 */

/**
 * @typedef {ResponseDataResults} MappingConfigurationResults
 * @property {string} configurationHash
 * @property {Array<string>} hiddenMapperFields
 * @property {MappingConfiguration} allMappingConfigurations
 */

/**
 * @typedef {Object} FormattedStripeObjectField
 * @property {string} defaultValue
 * @property {string} description
 * @property {boolean} hasOverride
 * @property {boolean} hasRequiredValue
 * @property {boolean} hasSfValue
 * @property {string} name
 * @property {string} requiredValue
 * @property {string} sfValue
 * @property {string} sfValueType
 * @property {boolean} staticValue
 * @property {string} type
 * @property {string} value
 */

/**
 * @typedef {Object} FormattedStripeObjectSection
 * @property {string} description
 * @property {string} label
 * @property {string} name
 * @property {Array<FormattedStripeObjectField>} fields
 */

/**
 * @typedef {ResponseDataResults} FormattedStripeObjectFieldsResults
 * @property {Array<FormattedStripeObjectSection>} formattedStripeCouponFields
 * @property {Array<FormattedStripeObjectSection>} formattedStripeCustomerFields
 * @property {Array<FormattedStripeObjectSection>} formattedStripePriceFields
 * @property {Array<FormattedStripeObjectSection>} formattedStripePriceOrderItemFields
 * @property {Array<FormattedStripeObjectSection>} formattedStripeProductItemFields
 * @property {Array<FormattedStripeObjectSection>} formattedStripeSubscriptionFields
 * @property {Array<FormattedStripeObjectSection>} formattedStripeSubscriptionItemFields
 * @property {Array<FormattedStripeObjectSection>} formattedStripeSubscriptionSchedulePhaseFields
 */

/**
 * @typedef {ResponseDataResults} SyncPrefsResults
 * @property {string} api_percentage_limit
 * @property {string} configurationHash
 * @property {string} cpq_term_unit
 * @property {string} default_currency
 * @property {string} stripe_account_id
 * @property {String} sync_record_retention
 * @property {boolean} enabled
 * @property {boolean} isCpqInstalled
 * @property {boolean} isSandbox
 * @property {null} last_synced
 * @property {boolean} polling_enabled
 * @property {null} sync_start_date
 * @property {Array<string>} hiddenSyncPrefsFields
 */

/**
 * @typedef {ResponseDataResults} FilterSettingsResults
 * @property {string|null} pricebook_entry_filter
 * @property {string|null} account_filter
 * @property {string|null} order_filter
 * @property {string|null} product_filter
 */

/**
 * @typedef {ResponseDataResults} MultiCurrencySelectOptionsResults
 * @property {number} orgMaxApiLimit
 * @property {boolean} isMultiCurrencyEnabled
 * @property {Object<{label: string, value: string}>} [supportedISOCodes] List of supported ISO Codes, with 'Select...' first.
 */

/**
 * @typedef {{string: number}} StepsCompleted
 */

/**
 * @typedef {ResponseDataResults} SetupDataObject
 * @property {string} Id
 * @property {{type: string, url: string}} attributes
 * @property {boolean} isSetupComplete__c
 * @property {string} Steps_Completed__c This is a serialized StepsCompleted object.
 */

/**
 * @typedef {Object} SetupDataResults
 * @property {SetupDataObject} setupData
 */

/**
 * @typedef {Object} PermissionIssueMap
 * @property {boolean} isObjectPermissionMissing
 * @property {boolean} isSystemPermissionMissing
 * @property {boolean} isPermSetAssigned
 * @property {Array<string>} missingSystemPermissionList
 * @property {Array<string>} missingObjectPermissionList
 * @property {Array<string>} missingPermissionSets
 */

/**
 * @typedef {ResponseDataResults} CheckUserPermissionsResults
 * @property {PermissionIssueMap} permissionIssueMap
 */

/**
 * @typedef {Object} ResponseDataResults
 * @property {boolean} isConnected
 */

/**
 * @typedef {Object<type>} ResponseData
 * @property {string} error
 * @property {boolean} isSuccess
 * @property {type} results
 */

/**
 * @typedef {Object} SetupData
 * @property {boolean} isSetupComplete
 * @property {StepsCompleted} stepsCompleted
 */

/**
 * @typedef {Object} TranslationConfigResults
 * @property {string} configurationHash
 * @property {string} api_percentage_limit
 * @property {string} configurationHash
 * @property {string} cpq_term_unit
 * @property {string} default_currency
 * @property {string} stripe_account_id
 * @property {String} sync_record_retention
 * @property {boolean} enabled
 * @property {null} last_synced
 * @property {boolean} polling_enabled
 * @property {null} sync_start_date
 * @property {Array<string>} hiddenSyncPrefsFields
 * @property {Array<string>} hiddenMapperFields
 * @property {MappingConfiguration} allMappingConfigurations
 * @property {string|null} pricebook_entry_filter
 * @property {string|null} account_filter
 * @property {string|null} order_filter
 * @property {string|null} product_filter
 */

/**
 * @typedef {Object} SaveTranslationConfig
 * @property {string} configurationHash
 * @property {string} api_percentage_limit
 * @property {string} cpq_term_unit
 * @property {string} default_currency
 * @property {string} stripe_account_id
 * @property {String} sync_record_retention
 * @property {boolean} enabled
 * @property {boolean} polling_enabled
 * @property {null} sync_start_date
 * @property {SaveMappingConfiguration} allMappingConfigurations
 * @property {string|null} pricebook_entry_filter
 * @property {string|null} account_filter
 * @property {string|null} order_filter
 * @property {string|null} product_filter
 */

/**
 *
 * @enum {string}
 */
const ConfigStates = {
    available: 'available',
    updated: 'updated',
}

/**
 *
 * @enum {string}
 */
const ConfigData = {
    all: 'all', // 'all' is a special value that will return all the data in the config manager
    core: 'core',
    setup: 'setup',
    layouts: 'layouts',
    translation: 'translation',
};

const ConfigEvents = {};
Object.keys(ConfigData).forEach((key) => {
    if (key === 'all') {
        ConfigEvents[`${key}.${ConfigStates.available}`] = [`${key}.${ConfigStates.available}`];
        return;
    }
    Object.keys(ConfigStates).forEach((state) => {
        ConfigEvents[`${key}.${state}`] = [`${key}.${state}`];
    });
});

const storeConfigData = (defaultVal) => {
    const result = {};
    Object.keys(ConfigData).forEach((key) => {
        // don't need to process All.
        if (key === ConfigData.all) {
            return;
        }
        result[key] = defaultVal;
    });
    return result;
}

class ConfigManager extends EventEmitter {
    constructor() {
        super(ConfigEvents);
        this._initialize();
    }

    _initialize() {
        this.all_available = false;
        this.config = storeConfigData(null);
        this.promises = storeConfigData(null);
        this.save_translations = null;
        this.getSetupData();
        this.getCoreData();
        this.getLayouts();
        this.getTranslationConfig();
    }

    /**
     *
     * @return {Promise<CoreDataResults>}
     */
    getCoreData() {
        if (this.promises.core === null) {
            this.promises.core = new Promise((resolve, reject) => {
                return getCoreData()
                    .then((data) => {
                        /** @type {ResponseData<CoreDataResults>} */
                        const response = JSON.parse(data);
                        if (response.isSuccess) {
                            this._processUpdate(ConfigData.core, response.results);
                            resolve(response.results);
                            return response.results;
                        } else {
                            reject(response.error);
                        }
                    })
                    .catch((error) => {
                        reject(error);
                    });
            }).finally(() => { this.promises.core = null; });
        }
        return this.promises.core;
    }

    /**
     * @return {Promise<SetupData>}
     */
    getSetupData() {
        if (this.promises.setup === null) {
            this.promises.setup = new Promise((resolve, reject) => {
                return getSetupData()
                    .then((data) => {
                        /** @type {ResponseData<SetupDataResults>} */
                        const response = JSON.parse(data);
                        const setupDataObject = response.results.setupData;
                        if (response.isSuccess) {
                            const setupData = {
                                isSetupComplete: setupDataObject.isSetupComplete__c,
                                stepsCompleted: JSON.parse(setupDataObject.Steps_Completed__c),
                            }
                            this._processUpdate(ConfigData.setup, setupData);
                            resolve(setupData);
                            return setupData;
                        } else {
                            reject(response.error);
                        }
                    })
                    .catch((error) => {
                        reject(error);
                    });
            }).finally(() => { this.promises.setup = null; });
        }

        return this.promises.setup;
    }

    /**
     * @return {Promise<FormattedStripeObjectFieldsResults>}
     */
    getLayouts() {
        if (this.promises.layouts === null) {
            this.promises.layouts = new Promise((resolve, reject) => {
                return getDataMapperLayout()
                    .then((data) => {
                        /** @type {ResponseData<FormattedStripeObjectFieldsResults>} */
                        const response = JSON.parse(data);
                        if (response.isSuccess) {
                            this._processUpdate(ConfigData.layouts, response.results);
                            resolve(response.results);
                            return response.results;
                        }

                        reject(response.error);
                    })
                    .catch((error) => {
                        reject(error);
                    });
            }).finally(() => { this.promises.layouts = null; });
        }

        return this.promises.layouts;
    }

    /**
     * @return {Promise<TranslationConfigResults>}
     */
    getTranslationConfig() {
        if (this.promises.translation === null) {
            this.promises.translation = new Promise((resolve, reject) => {
                return getTranslationConfig()
                    .then((data) => {
                        /** @type {ResponseData<TranslationConfigResults>} */
                        const response = JSON.parse(data);
                        if (response.isSuccess) {
                            this._processUpdate(ConfigData.translation, response.results);
                            resolve(response.results);
                            return response.results;
                        }

                        reject(response.error);
                    })
                    .catch((error) => {
                        reject(error);
                    });
            }).finally(() => { this.promises.translation = null; });
        }

        return this.promises.translation;
    }

    /**
     *
     * @param config_data
     * @return {Promise<Object>}
     */
    getCached(config_data) {
        if (config_data === ConfigData.all) {
            const pendingPromises = this._getPendingPromises();
            if (pendingPromises.length > 0) {
                return Promise.all(pendingPromises).then(() => {
                    return JSON.parse(JSON.stringify(this.config));
                });
            }
            return Promise.resolve(JSON.parse(JSON.stringify(this.config)));
        }

        if (this.promises[config_data] === null) {
            return Promise.resolve(JSON.parse(JSON.stringify(this.config[config_data])));
        }

        return new Promise((resolve, reject) => {
            return this.promises[config_data].then(() => {
                resolve(JSON.parse(JSON.stringify(this.config[config_data])));
            }).catch((error) => {
                reject(error);
            });
        });
    }

    /**
     * @return {Promise<TranslationConfigResults>}
     * @param {SaveTranslationConfig} config
     */
    saveTranslationConfig(config) {
        if (this.save_translations !== null) {
            return Promise.reject('Translation config already being saved');
        }
        if (this.config.translation === null) {
            return Promise.reject('Translation config not loaded, cannot save until loaded');
        }
        this.save_translations = new Promise((resolve, reject) => {
            const savePayload = this._formatTranslationConfigForSave(config);
            debugger;
            return saveTranslationConfig({ config: JSON.stringify(savePayload) })
                .then((data) => {
                    /** @type {ResponseData<TranslationConfigResults>} */
                    const response = JSON.parse(data);
                    if (response.isSuccess) {
                        savePayload.configurationHash = response.results.configurationHash;
                        this.config.translation = savePayload;
                        this._processUpdate(ConfigData.translation, savePayload);
                        resolve(savePayload);
                        return savePayload;
                    }

                    reject(response.error);
                })
        }).finally(() => { this.save_translations = null; });
    }

    /**
     *
     * @param {SaveTranslationConfig} updates
     * @return {SaveTranslationConfig}
     * @private
     */

    _formatTranslationConfigForSave(updates) {
        /** @type {SaveTranslationConfig} */
        const save = JSON.parse(JSON.stringify(this.config.translation));
        let fieldDefaults = save.allMappingConfigurations.field_defaults;
        let fieldMappings = save.allMappingConfigurations.field_mappings;
        if (updates.allMappingConfigurations) {
            if (updates.allMappingConfigurations.field_defaults) {
                fieldDefaults = updates.allMappingConfigurations.field_defaults;
            }
            if (updates.allMappingConfigurations.field_mappings) {
                fieldMappings = updates.allMappingConfigurations.field_mappings;
            }
        }

        save.allMappingConfigurations = {
            field_defaults: fieldDefaults,
            field_mappings: fieldMappings,
        }


        // incorporate the config into the full payload. This should hopefully fix any more of the config hash issues...
        Object.keys(updates).forEach((key) => {
            if (key === 'allMappingConfigurations') {
                return;
            }
            if (save[key]) {
                save[key] = updates[key];
            }
        });

        return save;
    }

    /**
     *
     * @param {ConfigData} config_data
     * @param {ConfigStates} config_state
     * @param listener
     */
    register(config_data, config_state, listener) {
        this.on(`${config_data}.${config_state}`, listener);
    }

    /**
     *
     * @param {ConfigData} config_data
     * @param {ConfigStates} config_state
     * @param listener
     */
    unregister(config_data, config_state, listener) {
        this.off(`${config_data}.${config_state}`, listener);
    }

    /**
     *
     * @param {ConfigData} config_data
     * @param {Object} value
     * @private
     */
    _processUpdate(config_data, value) {
        const current = this.config[config_data];
        const update_type = current === null ? ConfigStates.available : ConfigStates.updated;
        this.config[config_data] = value;

        // We need to make sure that we don't emit the original value, because it could be mutated.
        // If needed we can send each event a unique copy, but for now we'll just send a deep copy.
        const safeValue = JSON.parse(JSON.stringify(value));

        this._fireEvent(`${config_data}.${update_type}`, safeValue);
        if (update_type === ConfigStates.available) {
            this._fireEvent(`${config_data}.${ConfigStates.updated}`, safeValue);
        }

        if (this.all_available) {
            return;
        }

        const still_pending = this._getPendingPromises();
        if (still_pending.length === 0) {
            this.all_available = true;
            this._fireEvent('all.available', JSON.parse(JSON.stringify(this.config)));
            this.all_available = true;
        }
    }

    /**
     *
     * @return {Promise<Object>[]}
     * @private
     */
    _getPendingPromises() {
        return Object.keys(this.promises).filter((key) => this.promises[key]).map(key => this.promises[key]);
    }
}

const Manager = new ConfigManager();

module.exports = {
    ConfigManager: Manager,
    ConfigData,
    ConfigStates,
}