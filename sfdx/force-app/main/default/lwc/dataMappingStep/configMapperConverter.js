/**
 * Created by jmather-c on 10/20/23.
 */

import {ConfigManager, ConfigStates, ConfigData, ConfigValidationError} from 'c/systemConfigManager';

/**
 * @typedef {Object} GetPicklistValuesForMapperArguments
 * @property {boolean} isConnectedCallback
 * @property {string} ObjectApiName
 */

/**
 * @typedef {Object} SaveMappingConfigurationsArguments
 * @property {string} jsonMappingConfigurationsObject
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
 * @param {GetPicklistValuesForMapperArguments} args
 * @return {Promise<string>}
 */
async function getPicklistValuesForMapperCore(args) {
    /** @type {CoreData} */
    const result = await ConfigManager.getCachedCoreData();
    const response = createBaseResponse();

    if (args.isConnectedCallback) {
        response.results.fieldListByObjectMap = result.fieldListByObjectMap;
    } else {
        response.results.ObjectApiName = args.ObjectApiName;
        response.results.fieldListByObjectMap = result.fieldListByObjectMap[args.ObjectApiName];
    }

    return JSON.stringify(response);
}

/**
 *
 * @return {Promise<string>}
 */
async function getFormattedStripeObjectFieldsCore() {
    const result = await ConfigManager.getCachedLayoutData();
    const response = createBaseResponse();
    response.results = result;
    return JSON.stringify(response);
}

/**
 *
 * @return {Promise<string>}
 */
async function getMappingConfigurationsCore() {
    const result = await ConfigManager.getCachedTranslationData();
    const response = createBaseResponse();
    response.results.allMappingConfigurations = result.allMappingConfigurations;
    response.results.hiddenMapperFields = result.hiddenMapperFields;
    response.results.configurationHash = result.configurationHash;
    return JSON.stringify(response);
}

/**
 *
 * @param {SaveMappingConfigurationsArguments} args
 * @return {Promise<string>}
 */
async function saveMappingConfigurationsCore(args) {
    const data = JSON.parse(args.jsonMappingConfigurationsObject);
    const payload = {
        allMappingConfigurations: {
            field_mappings: data.field_mappings,
            field_defaults: data.field_defaults,
        }
    }
    const result =  await ConfigManager.saveTranslationConfig(payload);
    const response = createBaseResponse();
    response.results = result;
    return JSON.stringify(response);
}

/**
 *
 * @param {GetPicklistValuesForMapperArguments} args
 * @return {Promise<string>}
 */
export async function getPicklistValuesForMapper(args) {
    return wrapWithErrorHandling(getPicklistValuesForMapperCore)(args);
}

/**
 *
 * @return {Promise<string>}
 */
export async function getFormattedStripeObjectFields() {
    return wrapWithErrorHandling(getFormattedStripeObjectFieldsCore)();
}

/**
 *
 * @return {Promise<string>}
 */
export async function getMappingConfigurations() {
    return wrapWithErrorHandling(getMappingConfigurationsCore)();
}

/**
 * @param {SaveMappingConfigurationsArguments} args
 * @return {Promise<string>}
 */
export async function saveMappingConfigurations(args) {
    return wrapWithErrorHandling(saveMappingConfigurationsCore)(args);
}