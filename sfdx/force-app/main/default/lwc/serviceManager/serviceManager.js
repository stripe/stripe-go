/**
 * Created by jmather-c on 8/31/23.
 */
import getServiceConnectionStatuses from '@salesforce/apexContinuation/setupAssistant.getServiceConnectionStatuses';
import checkServiceConnectionStatuses from '@salesforce/apex/setupAssistant.checkServiceConnectionStatuses';
import {ConnectionStatus, ServiceManagerServices, EventEmitter} from "c/systemStatusUtils";
import Debugger from "c/debugger";
const DebugLog = Debugger.withContext('ServiceManager');
const SERVICE_DELIMINATOR = '|';

/** @typedef {Object} ConnectionStatusChangeDetail
 * @property {string, ConnectionStatus} statuses
 * @property {string, ConnectionStatus} changes
 */

/** @typedef {Object} CoreFunctionalityEstablishedDetail
 * @property {boolean} salesforce
 * @property {boolean} stripe
 */

/** @typedef {Event} ConnectionStatusChangeEvent
 * @property {ConnectionStatusChangeDetail} detail
 */

/** @typedef {Event} CoreFunctionalityEstablishedEvent
 * @property {CoreFunctionalityEstablishedDetail} detail
 */

/** @enum {string} */
const ServiceEvents = {
    connection_status_updated: 'connection_status_updated',
    core_functionality_established: 'core_functionality_established',
}

class ServiceManager extends EventEmitter {
    /** @type {{string: ConnectionStatus}} */
    statuses = {};
    updating = null;
    initialized = false;
    initializing = null;
    core_functionality_established = {
        [ServiceManagerServices.salesforce]: false,
        [ServiceManagerServices.stripe]: false,
    };

    constructor() {
        super(ServiceEvents);
        DebugLog('New ServiceManager instance created');
        this.updateConnectionStatuses();
    }

    _initialize() {
        if (this.initialized) {
            return Promise.resolve(this.statuses);
        }
        if (this.initializing) {
            DebugLog('ServiceManager initialization already in progress');
            return this.initializing;
        }
        this.initializing = checkServiceConnectionStatuses()
            .then((result) => {
                /** @type {SystemConnectionStatusesPayload} */
                const responseData = JSON.parse(result);
                DebugLog('ServiceManager initialization complete', { responseData });
                const statuses = this._processServiceStatusResults(responseData);
                this.initialized = true;
                this.initializing = null;
                return statuses;
            })
            .catch((error) => {
                this._processServiceStatusError(error);
                this.initializing = null;
                this.initialized = false;
            });
        return this.initializing;
    }

    /**
     *
     * @param {string} [serviceToCheck]
     * @return {null}
     */
    updateConnectionStatuses(serviceToCheck) {
        DebugLog('Updating connection statuses', { serviceToCheck });
        if (this.updating) {
            DebugLog('Connection status update already in progress');
            return this.updating;
        }
        this.updating = new Promise((resolve, reject) => {
            this._initialize().then(() => {
                return getServiceConnectionStatuses({ serviceToCheck })
                    .then(result => {
                        /** @type {SystemConnectionStatusesPayload} */
                        const responseData = JSON.parse(result);
                        DebugLog('Connection status update complete', { responseData });
                        return resolve(this._processServiceStatusResults(responseData))
                    })
                    .catch(reject);
            }).catch(reject);
        }).catch(error => this._processServiceStatusError)

        return this.updating;
    }

    _processServiceStatusError(error) {
        Debugger.error(error);
        this.updating = null;
    }

    /**
     *
     * @param {SystemConnectionStatusesPayload} responseData
     * @return {{string: ConnectionStatus}}
     * @private
     */
    _processServiceStatusResults(responseData) {
        DebugLog('Connection status update complete, processing results', responseData);

        const oldStatuses = this.statuses;
        const newStatuses = responseData.results.services;
        const changes = {};
        const checkCFE = {
            [ServiceManagerServices.salesforce]: this.core_functionality_established[ServiceManagerServices.salesforce],
            [ServiceManagerServices.stripe]: this.core_functionality_established[ServiceManagerServices.stripe],
        };
        let cfeChanged = false;
        DebugLog('current statuses', oldStatuses);
        DebugLog('new statuses', newStatuses);
        DebugLog('check cfe', checkCFE);

        for (const service in newStatuses) {
            const [service_type, service_id, service_props] = service.split(SERVICE_DELIMINATOR);
            if (this.statuses[service] !== newStatuses[service]) {
                changes[service] = newStatuses[service];
                // this._fireEvent(service, newStatuses[service]);
            }

            if (checkCFE[service_type] === false && this._isConnected(newStatuses[service])) {
                cfeChanged = true;
                checkCFE[service_type] = true;
            }
        }

        DebugLog('changes', changes);
        DebugLog('check cfe', checkCFE);

        this.core_functionality_established = checkCFE;
        this.statuses = newStatuses;

        if (Object.keys(changes).length > 0) {
            DebugLog('firing connection_status_updated event', changes);
            this._fireEvent(ServiceEvents.connection_status_updated, {
                statuses: this.statuses,
                changes,
            });
        }

        // for (const serviceType in checkCFE) {
        //     this._fireEvent(serviceType, checkCFE[serviceType]);
        // }

        if (cfeChanged) {
            DebugLog('firing core_functionality_established event', this.statuses);
            this._fireEvent(ServiceEvents.core_functionality_established, {
                isConnected: this.isCoreFunctionalityEnabled(),
                statuses: this.statuses,
                isFirstRun: Object.keys(oldStatuses).length === 0,
            });
        }

        this.updating = null;
        return this.statuses;
    }

    isCoreFunctionalityEnabled() {
        return this.core_functionality_established[ServiceManagerServices.salesforce] && this.core_functionality_established[ServiceManagerServices.stripe];
    }

    /**
     *
     * @param {string} status
     * @return {boolean}
     * @private
     */
    _isConnected(status) {
        return status === ConnectionStatus.connected || status === ConnectionStatus.fresh;
    }

    /**
     *
     * @param {ServiceEvents|string} event_name
     * @param {Function<Event>} callback
     */
    on(event_name, callback) {
        super.on(event_name, callback);

        if (event_name === ServiceEvents.connection_status_updated) {
            callback(this._createEvent(ServiceEvents.connection_status_updated, { statuses: this.statuses }));
        }

        if (event_name === ServiceEvents.core_functionality_established && this.isCoreFunctionalityEnabled()) {
            callback(this._createEvent(ServiceEvents.core_functionality_established, this.core_functionality_established));
        }

        if (this.statuses[event_name] !== undefined) {
            callback(this._createEvent(event_name, this.statuses[event_name]));
        }
    }

    /**
     *
     * @param {ServiceEvents|string} event_name
     * @param {Function<Event>} callback
     */
    off(event_name, callback) { super.off(event_name, callback); }

    /**
     *
     * @param {ServiceEvents|string} event_name
     * @param {Object} detail
     * @protected
     * @emits {Event}
     */
    _fireEvent(event_name, detail) { super._fireEvent(event_name, detail); }
}

const Manager = new ServiceManager();

export { Manager, ServiceEvents };