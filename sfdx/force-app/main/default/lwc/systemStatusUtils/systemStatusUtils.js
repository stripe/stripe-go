/**
 * Created by jmather-c on 8/30/23.
 */

import { MessageListener, ListenerEvents } from "./postMessageListener";
import {EventEmitter} from "./eventEmitter";

class SystemStatusUtils {

}

/**
 * @enum {string}
 * @readonly
 */
const ConnectionStatus = {
    failed: 'failed',
    fresh: 'fresh',
    connected: 'connected',
    disconnected: 'disconnected',
    loading: 'loading'
}

/**
 * @enum {string}
 * @readonly
 */
const ServiceManagerServices = {
    salesforce: 'salesforce',
    stripe: 'stripe',
}


/** @typedef {Object} SystemConnectionStatusConfig
 * @property {boolean} is_sandbox
 * @property {string} salesforce_namespace
 * @property {string} salesforce_sandbox_uri_fragment
 * @property {string} salesforce_production_uri_fragment
 * @property {string} stripe_test_uri_fragment
 * @property {string} stripe_live_uri_fragment
 * @property {string} platform_uri
 * @property {string} state
 */

/** @typedef {Object} SystemConnectionStatusesPayload
 * @property {{services: {string: ConnectionStatus}}} results
 * @property {boolean} isSuccess
 * @property {string} error
 */

/** @typedef {Object} IntegrationUserAuthorizationUriResults
 * @property {string} authorization_uri
 * @property {string} message_origin_uri
 */

/** @typedef {Object} IntegrationUserAuthorizationUriPayload
 * @property {IntegrationUserAuthorizationUriResults} results
 * @property {boolean} isSuccess
 * @property {string} error
 */

/** @typedef {Object} SystemStatus
 * @property {ConnectionStatus} status
 * @property {string} service
 */

/** @typedef {Object} SystemStatusPayload
 * @property {SystemStatus} results
 * @property {boolean} isSuccess
 * @property {string} error
 */


const TextStrings = {
    salesforce: 'Salesforce',
    stripe: 'Stripe',
    connected: 'Connected',
    not_connected: 'Not Connected',
    authorize: 'Authorize',
    reauthorize: 'Reauthorize',
    success: 'success',
    offline: 'offline',
    icon_success: 'utility:success',
    icon_offline: 'utility:offline',
}

export {
    SystemStatusUtils,
    ConnectionStatus,
    TextStrings,
    MessageListener,
    ListenerEvents,
    EventEmitter,
    ServiceManagerServices,
}