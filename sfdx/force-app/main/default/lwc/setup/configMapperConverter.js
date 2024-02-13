/**
 * Created by jmather-c on 11/13/23.
 *
 * File is WIP and not implemented in setup.js yet to minimize scope of change.
 */

import {ConfigManager} from 'c/systemConfigManager';


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
async function getSetupDataCore() {
    const result = await ConfigManager.getCachedSetupData();
    const response = createBaseResponse();
    response.results.Id = result.id;
    response.results.attributes = {
        type: "Setup_Data__c",
        url: "/services/data/v58.0/sobjects/Setup_Data__c/a1TDR0000020Xxf2AE"
    };
    response.results.isSetupComplete__c = result.isSetupComplete;
    response.results.Steps_Completed__c = JSON.stringify(result.stepsCompleted);
    return JSON.stringify(response);
}

/**
 *
 * @return {Promise<string>}
 */
async function getPackageVersionCore() {
    const result = await ConfigManager.getCachedCoreData();
    const pieces = result.packageVersion.split('.');

    const response = createBaseResponse();
    response.results.major = (pieces && pieces.length > 0) ? pieces[0] : '0';
    response.results.minor = (pieces && pieces.length > 1) ? pieces[1] : '0';
    return JSON.stringify(response);
}

/**
 *
 * @return {Promise<string>}
 */
export async function getSetupData() {
    return wrapWithErrorHandling(getSetupDataCore)();
}

/**
 *
 * @return {Promise<string>}
 */
export async function getPackageVersion() {
    return wrapWithErrorHandling(getPackageVersionCore)();
}
