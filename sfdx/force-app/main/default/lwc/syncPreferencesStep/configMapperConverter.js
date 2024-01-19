/**
 * Created by jmather-c on 10/20/23.
 */
import {ConfigManager, ConfigStates, ConfigData, ConfigValidationError} from 'c/systemConfigManager';

/** @typedef {Object} SaveSyncPreferencesArguments
 * @property {boolean} pollingEnabled
 * @property {string} defaultCurrency
 * @property {string} syncRecordRetention
 * @property {number} syncStartDate
 * @property {string} apiPercentageLimit
 * @property {string} cpqTermUnit
 * @property {string} cpqProratePrecision
 * @property {string} configurationHash
 */

/** @typedef {Object} SaveFilterSettingsArguments
 * @property {string} pricebookEntryFilter
 * @property {string} productFilter
 * @property {string} accountFilter
 * @property {string} orderFilter
 */

function createBaseResponse() {
    return {
        error: null,
        isSuccess: true,
        results: {
            isConnected: true,
        }
    };
}

function createErrorResponse(error) {
    return {
        error,
        isSuccess: false,
        results: {
            isConnected: true,
        }
    };
}

/**
 *
 * @param {(function(...[*]): Promise<string>)} fn
 * @return {(function(...[*]): Promise<string>)}
 */
function wrapWithErrorHandling(fn) {
    return async function (...args) {
        try {
            return await fn(...args);
        } catch (error) {
            return JSON.stringify(createErrorResponse(error.message));
        }
    }
}

/**
 *
 * @return {Promise<string>}
 */
async function getSyncPreferencesCore() {
    const core = await ConfigManager.getCachedCoreData();
    const config = await ConfigManager.getCachedTranslationData();
    const response = createBaseResponse();
    response.results.enabled = true;
    response.results.hiddenSyncPrefsFields = config.hiddenSyncPrefsFields;
    response.results.last_synced = config.last_synced;
    response.results.stripe_account_id = config.stripe_account_id;
    response.results.isCpqInstalled = core.isCpqInstalled;
    response.results.default_currency = core.defaultCurrency;
    response.results.polling_enabled = config.polling_enabled;
    response.results.cpq_term_unit = config.cpq_term_unit;
    response.results.api_percentage_limit = config.api_percentage_limit;
    response.results.sync_start_date = config.sync_start_date;
    response.results.sync_record_retention = config.sync_record_retention;
    response.results.configurationHash = config.configurationHash;
    response.results.isSandbox = core.isSandbox;
    return JSON.stringify(response);
}

/**
 *
 * @return {Promise<string>}
 */
async function getMulticurrencySelectionOptionsCore() {
    const core = await ConfigManager.getCachedCoreData();
    const response = createBaseResponse();
    response.results.multicurrencySelectionOptions = core.supportedISOCodes;
    response.results.orgMaxApiLimit = core.orgMaxApiLimit;
    response.results.isMultiCurrencyEnabled = core.isMultiCurrencyEnabled;
    response.results.supportedISOCodes = core.supportedISOCodes;
    return JSON.stringify(response);
}

/**
 *
 * @return {Promise<string>}
 */
async function getFilterSettingsCore() {
    const config = await ConfigManager.getCachedTranslationData();
    const response = createBaseResponse();
    response.results.pricebook_entry_filter = config.pricebook_entry_filter;
    response.results.product_filter = config.product_filter;
    response.results.account_filter = config.account_filter;
    response.results.order_filter = config.order_filter;
    return JSON.stringify(response);
}

/**
 *
 * @param {SaveSyncPreferencesArguments} args
 * @return {Promise<string>}
 */
async function saveSyncPreferencesCore(args) {
    const payload = {
        polling_enabled: args.pollingEnabled,
        sync_record_retention: args.syncRecordRetention,
        sync_start_date: args.syncStartDate,
        api_percentage_limit: args.apiPercentageLimit,
        cpq_term_unit: args.cpqTermUnit,
        cpq_prorate_precision: args.cpqProratePrecision,
    };
    const result =  await ConfigManager.saveTranslationConfig(payload);
    const response = createBaseResponse();
    response.results = result;
    return JSON.stringify(response);
}

/**
 *
 * @param {SaveFilterSettingsArguments} args
 * @return {Promise<string>}
 */
async function saveFilterSettingsCore(args) {
    const payload = {
        pricebook_entry_filter: args.pricebookEntryFilter,
        product_filter: args.productFilter,
        account_filter: args.accountFilter,
        order_filter: args.orderFilter,
    };
    const result =  await ConfigManager.saveTranslationConfig(payload);
    const response = createBaseResponse();
    response.results = result;
    return JSON.stringify(response);
}

/**
 *
 * @return {Promise<string>}
 */
export async function getSyncPreferences() {
    return wrapWithErrorHandling(getSyncPreferencesCore)();
}

/**
 *
 * @return {Promise<string>}
 */
export async function getMulticurrencySelectionOptions() {
    return wrapWithErrorHandling(getMulticurrencySelectionOptionsCore)();
}

/**
 *
 * @return {Promise<string>}
 */
export async function getFilterSettings() {
    return wrapWithErrorHandling(getFilterSettingsCore)();
}

/**
 *
 * @param {SaveSyncPreferencesArguments} args
 * @return {Promise<string>}
 */
export async function saveSyncPreferences(args) {
    return wrapWithErrorHandling(saveSyncPreferencesCore)(args);
}

/**
 *
 * @param {SaveFilterSettingsArguments} args
 * @return {Promise<string>}
 */
export async function saveFilterSettings(args) {
    return wrapWithErrorHandling(saveFilterSettingsCore)(args);
}