/**
 * Created by jmather-c on 10/10/23.
 */

import getCoreData from '@salesforce/apex/setupAssistant.getCoreData';
import getSetupData from '@salesforce/apex/setupAssistant.getSetupData';
import getDataMapperLayout from '@salesforce/apex/setupAssistant.getDataMapperLayout';
import getTranslationConfig from '@salesforce/apex/setupAssistant.getTranslationConfig';
import saveTranslationConfig from '@salesforce/apex/setupAssistant.saveTranslationConfig';
import { Manager as ServiceManager, ServiceEvents } from 'c/serviceManager';
import EventEmitter from './eventEmitter';
import Debugger from 'c/debugger';
const DebugLog = Debugger.withContext('SystemConfigManager');


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
 * @typedef {ResponseDataResults} CoreData
 * @property {string} packageVersion
 * @property {string} configDownloadUri
 * @property {boolean} isMultiCurrencyEnabled
 * @property {Array<{label: string, value: string}>} supportedISOCodes List of supported ISO Codes, with 'Select...' first.
 * @property {string} defaultCurrency
 * @property {MultiObjectPicklistValuesForMapper} fieldListByObjectMap
 * @property {boolean} isCpqInstalled
 * @property {boolean} isSandbox
 * @property {number} orgMaxApiLimit
 */

/**
 * @typedef {SaveMappingConfiguration} MappingConfiguration
 * @property {Object<string, Object<string, string>>} default_mappings
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
 * @typedef {ResponseDataResults} LayoutData
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
 * @typedef {Object} ResponseDataSaveResults
 * @property {boolean} isConfigSaved
 * @property {boolean} isValidationError
 * @property {boolean} isFiltersSaved
 * @property {Array<ValidationError>} ValidationErrors
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
 * @typedef {Object} ValidationError
 * @property {string} Error
 * @property {string} Object
 */

/**
 * @typedef {Object} TranslationData
 * @property {string} configurationHash
 * @property {string} api_percentage_limit
 * @property {string} configurationHash
 * @property {string} cpq_term_unit
 * @property {string} default_currency
 * @property {string} stripe_account_id
 * @property {String} sync_record_retention
 * @property {boolean} enabled
 * @property {string|null} last_synced
 * @property {boolean} polling_enabled
 * @property {string|null} sync_start_date
 * @property {Array<string>} hiddenSyncPrefsFields
 * @property {Array<string>} hiddenMapperFields
 * @property {MappingConfiguration} allMappingConfigurations
 * @property {string|null} pricebook_entry_filter
 * @property {string|null} account_filter
 * @property {string|null} order_filter
 * @property {string|null} product_filter
 * @property {boolean} isConfigSaved
 * @property {boolean} isConnected
 */

/**
 * @typedef {Object} SaveTranslationData
 * @property {string} [configurationHash]
 * @property {string} [api_percentage_limit]
 * @property {string} [cpq_term_unit]
 * @property {String} [sync_record_retention]
 * @property {boolean} [polling_enabled]
 * @property {Date|null} [sync_start_date]
 * @property {SaveMappingConfiguration} [allMappingConfigurations]
 * @property {string|null} [pricebook_entry_filter]
 * @property {string|null} [account_filter]
 * @property {string|null} [order_filter]
 * @property {string|null} [product_filter]
 */

/**
 * @typedef {ResponseDataSaveResults} SaveTranslationDataResults
 * @property {string} configurationHash
 */

/**
 * @typedef {Object} AllResults
 * @property {CoreData} core
 * @property {SetupData} setup
 * @property {LayoutData} layouts
 * @property {TranslationData} translation
 */

/**
 * @template T
 * @typedef {Event} SystemConfig<T>
 * @property {T} detail
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

class ConfigValidationError extends Error {
    /**
     *
     * @param {SaveTranslationDataResults} results
     */
    constructor(results) {
        super('Validation errors');
        this.results = results;
        this.errors = results.ValidationErrors;
    }

    getErrors() {
        return this.errors;
    }
}

class UnknownError extends Error {
    /**
     *
     * @param {ResponseData<SaveTranslationDataResults>} response
     */
    constructor(response) {
        super('Unknown error: ' + response.error);
        this.response = response;
    }
}

const DisconnectedPayload = () => { return { isConnected: false } };

class SystemConfigManager extends EventEmitter {
    constructor() {
        super(ConfigEvents, true);
        this.initialize();
    }

    initialize() {
        this.all_available = false;
        this.config = storeConfigData(null);
        this.promises = storeConfigData(null);
        this.save_translations = null;
        this.is_connected = false;
        Promise.all([
            this.getSetupData(),
            this.getCoreData(),
        ]).then(results => {
            DebugLog('ConfigManager initialized', results);
        }).catch(error => {
            Debugger.error('ConfigManager failed to initialize', error);
        });

        ServiceManager.on(ServiceEvents.core_functionality_established, this._coreFunctionalityEstablished.bind(this));
    }

    _coreFunctionalityEstablished() {
        this.is_connected = true;
        return Promise.all([
            this.getLayoutData(),
            this.getTranslationData(),
        ]).then(() => {
            DebugLog('ConfigManager _coreFunctionalityEstablished succeeded');
        }).catch(error => {
            Debugger.error('ConfigManager _coreFunctionalityEstablished failed', error);
        })
    }

    /**
     *
     * @return {Promise<CoreData>}
     */
    getCoreData() {
        if (this.promises.core === null) {
            this.promises.core = new Promise((resolve, reject) => {
                DebugLog('getCoreData called')
                return getCoreData()
                    .then((data) => {
                        this.promises.core = null;
                        DebugLog('getCoreData response', data)
                        /** @type {ResponseData<CoreData>} */
                        const response = JSON.parse(data);
                        DebugLog('getCoreData response decoded', response)
                        if (response.isSuccess) {
                            this._processUpdate(ConfigData.core, response.results);
                            resolve(response.results);
                            return response.results;
                        } else {
                            Debugger.error('getCoreData call error', response.error)
                            reject(response.error);
                            return new Error(response.error);
                        }
                    })
                    .catch((error) => {
                        Debugger.error('getCoreData error', error);
                        reject(error);
                        return error;
                    });
            });
        }
        return this.promises.core;
    }

    /**
     * @return {Promise<SetupData>}
     */
    getSetupData() {
        if (this.promises.setup === null) {
            this.promises.setup = new Promise((resolve, reject) => {
                DebugLog('getSetupData called')
                return getSetupData()
                    .then((data) => {
                        this.promises.setup = null;
                        DebugLog('getSetupData response', data)
                        /** @type {ResponseData<SetupDataResults>} */
                        const response = JSON.parse(data);
                        DebugLog('getSetupData response decoded', response)
                        const setupDataObject = response.results.setupData;
                        if (response.isSuccess) {
                            const setupData = {
                                id: setupDataObject.Id,
                                isSetupComplete: setupDataObject.isSetupComplete__c,
                                stepsCompleted: JSON.parse(setupDataObject.Steps_Completed__c),
                            }
                            this._processUpdate(ConfigData.setup, setupData);
                            resolve(setupData);
                            return setupData;
                        } else {
                            Debugger.error('getSetupData call error', response.error);
                            reject(response.error);
                            return new Error(response.error);
                        }
                    })
                    .catch((error) => {
                        Debugger.error('getSetupData error', error);
                        reject(error);
                        return error;
                    });
            });
        }

        return this.promises.setup;
    }

    /**
     * @return {Promise<LayoutData>}
     */
    getLayoutData() {
        if (this.promises.layouts === null) {
            this.promises.layouts = new Promise((resolve, reject) => {
                if (this.is_connected === false) {
                    DebugLog('getLayoutData called before connection established, waiting for connection');
                    resolve(DisconnectedPayload());
                    return DisconnectedPayload();
                }
                DebugLog('getLayoutData called')
                return getDataMapperLayout()
                    .then((data) => {
                        this.promises.layouts = null;
                        DebugLog('getLayoutData response', data)
                        /** @type {ResponseData<LayoutData>} */
                        const response = JSON.parse(data);
                        DebugLog('getLayoutData response decoded', response)
                        if (response.isSuccess) {
                            this._processUpdate(ConfigData.layouts, response.results);
                            resolve(response.results);
                            return response.results;
                        }

                        Debugger.error('getLayoutData call error', response.error);
                        reject(response.error);
                        return new Error(response.error);
                    })
                    .catch((error) => {
                        Debugger.error('getLayoutData error', error);
                        reject(error);
                        return error;
                    });
            });
        }

        return this.promises.layouts;
    }

    /**
     * @return {Promise<TranslationData>}
     */
    getTranslationData() {
        if (this.promises.translation === null) {
            this.promises.translation = new Promise((resolve, reject) => {
                if (this.is_connected === false) {
                    DebugLog('getTranslationData called before connection established, waiting for connection');
                    resolve(DisconnectedPayload());
                    return DisconnectedPayload();
                }
                DebugLog('getTranslationData called')
                return getTranslationConfig()
                    .then((data) => {
                        this.promises.translation = null;
                        DebugLog('getTranslationData response', data)
                        /** @type {ResponseData<TranslationData>} */
                        const response = JSON.parse(data);
                        DebugLog('getTranslationData response decoded', response)
                        if (response.isSuccess) {
                            this._processUpdate(ConfigData.translation, response.results);
                            resolve(response.results);
                            return response.results;
                        }

                        Debugger.error('getTranslationData call error', response.error);
                        reject(response.error);
                        return new Error(response.error);
                    })
                    .catch((error) => {
                        Debugger.error('getTranslationData error', error);
                        reject(error);
                        return error;
                    });
            });
        }

        return this.promises.translation;
    }

    /**
     *
     * @param {ConfigData} config_data
     * @return {Promise<TranslationData|LayoutData|SetupData|CoreData>}
     */
    getCached(config_data) {
        if (config_data === ConfigData.all) {
            DebugLog('getCached called with all')
            const pendingPromises = this._getPendingPromises();
            if (pendingPromises.length > 0) {
                DebugLog('getCached has pending promises')
                return Promise.all(pendingPromises).then(() => {
                    DebugLog('getCached all pending promises resolved')
                    return JSON.parse(JSON.stringify(this.config));
                });
            }
            DebugLog('getCached no pending promises')
            return Promise.resolve(JSON.parse(JSON.stringify(this.config)));
        }

        if (this.promises[config_data] === null) {
            DebugLog('getCached called with, have cached value', config_data)
            return Promise.resolve(JSON.parse(JSON.stringify(this.config[config_data])));
        }

        return new Promise((resolve, reject) => {
            DebugLog('getCached called with, no cached value', config_data)
            return this.promises[config_data].then(() => {
                DebugLog('getCached promise resolved', config_data)
                const val = JSON.parse(JSON.stringify(this.config[config_data]));
                resolve(val);
                return val;
            }).catch((error) => {
                Debugger.error('getCached promise rejected', config_data, error)
                reject(error);
                return error;
            });
        });
    }

    /**
     *
     * @return {Promise<SetupData>}
     */
    getCachedSetupData() {
        return this.getCached(ConfigData.setup);
    }

    /**
     *
     * @return {Promise<CoreData>}
     */
    getCachedCoreData() {
        return this.getCached(ConfigData.core);
    }

    /**
     *
     * @return {Promise<TranslationData>}
     */
    getCachedTranslationData() {
        if (this.is_connected === false) {
            return new Promise((resolve) => {
                DebugLog('getCachedTranslationData called before connection established, waiting for connection');
                resolve(DisconnectedPayload());
                return DisconnectedPayload();
            });
        }

        return this.getCached(ConfigData.translation);
    }

    /**
     *
     * @return {Promise<LayoutData>}
     */
    getCachedLayoutData() {
        return this.getCached(ConfigData.layouts);
    }

    /**
     * @return {Promise<TranslationData>}
     * @param {SaveTranslationData} config
     */
    saveTranslationConfig(config) {
        if (this.save_translations !== null) {
            DebugLog('Translation config already being saved')
            return Promise.reject('Translation config already being saved');
        }
        if (this.config.translation === null) {
            DebugLog('Translation config not loaded, cannot save until loaded')
            return Promise.reject('Translation config not loaded, cannot save until loaded');
        }
        const savePayload = this._formatTranslationConfigForSave(config);
        DebugLog('saveTranslationConfig called', savePayload)

        this.save_translations = saveTranslationConfig({ config: JSON.stringify(savePayload) })
            .then((data) => {
                this.save_translations = null;
                DebugLog('saveTranslationConfig response', data)
                /** @type {ResponseData<SaveTranslationDataResults>} */
                const response = JSON.parse(data);
                DebugLog('saveTranslationConfig response decoded', response)
                if (response.isSuccess && response.error === null && response.results.isValidationError !== true) {
                        savePayload.configurationHash = response.results.configurationHash;
                        savePayload.isConfigSaved = response.results.isConfigSaved;
                        savePayload.isValidationError = response.results.isValidationError;
                        savePayload.isFiltersSaved = response.results.isFiltersSaved;
                        this.config.translation = savePayload;
                        this._processUpdate(ConfigData.translation, savePayload);
                        return savePayload;
                } else if (response.error === "Validation errors") {
                    Debugger.warn('saveTranslationConfig call warning', response.error, response.results.errors);
                    throw new ConfigValidationError(response.results.ValidationErrors);
                }

                Debugger.error('saveTranslationConfig call error', response.error);
                throw new UnknownError(response.error);
            });

        return this.save_translations;
    }

    /**
     *
     * @param {SaveTranslationData} updates
     * @return {SaveTranslationData}
     * @private
     */

    _formatTranslationConfigForSave(updates) {
        DebugLog('formatTranslationConfigForSave called', updates)
        /** @type {SaveTranslationData} */
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
            save[key] = updates[key];
        });

        DebugLog('formatTranslationConfigForSave returning', save)
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
        this.promises[config_data] = null;
        DebugLog('processUpdate called', config_data, value)
        const current = this.config[config_data];
        const update_type = current === null ? ConfigStates.available : ConfigStates.updated;
        this.config[config_data] = value;

        // We need to make sure that we don't emit the original value, because it could be mutated.
        // If needed we can send each event a unique copy, but for now we'll just send a deep copy.
        const safeValue = JSON.parse(JSON.stringify(value));

        DebugLog('processUpdate firing event', `${config_data}.${update_type}`, value)
        this._fireEvent(`${config_data}.${update_type}`, safeValue);
        if (update_type === ConfigStates.available) {
            DebugLog('processUpdate firing event', `${config_data}.${ConfigStates.updated}`, value)
            this._fireEvent(`${config_data}.${ConfigStates.updated}`, safeValue);
        }

        if (this.all_available) {
            DebugLog('already sent all.available')
            return;
        }

        DebugLog('processUpdate checking for all available', {promises: this.promises})
        const still_pending = this._getPendingPromises(config_data);
        DebugLog('still_pending', still_pending);
        if (still_pending.length === 0) {
            this.all_available = true;
            DebugLog('processUpdate firing all.available', JSON.parse(JSON.stringify(this.config)))
            this._fireEvent('all.available', JSON.parse(JSON.stringify(this.config)));
        }
    }

    /**
     *
     * @return {string[]}
     * @private
     */
    _getPendingPromises() {
        return Object.keys(this.promises).filter((key) => this.promises[key] !== null);
    }
}

const ConfigManager = new SystemConfigManager();

export {ConfigManager, ConfigData, ConfigStates, ConfigValidationError};