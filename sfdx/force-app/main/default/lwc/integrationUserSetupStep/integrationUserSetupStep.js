import getIntegrationUserAuthorizationUri from '@salesforce/apex/setupAssistant.getIntegrationUserAuthorizationUri';
import revokeIntegrationUserAuthorization from '@salesforce/apex/setupAssistant.revokeIntegrationUserAuthorization';
import {LightningElement, track, api} from 'lwc';
import {getErrorMessage, createToast, openWindow} from 'c/utils'
import Debugger from "c/debugger";
import {MessageListener, ConnectionStatus, ListenerEvents, ServiceManagerServices} from "c/systemStatusUtils";
import { Manager, ServiceEvents } from "c/serviceManager";
const DebugLog = Debugger.withContext('IntegrationUserSetupStep');
const SERVICE_DELIMINATOR = '|';

export default class IntegrationUserSetupStep extends LightningElement {
    @track isConnected = false;
    @track loading = false;
    @track modalLoading = false;
    @track status = ConnectionStatus.loading;

    @track integrationUserType = 'self';
    @api hideAction = false;
    @api isSetup = false;
    _connectWindow = null;
    _expectedUpdate = undefined;

    async connectedCallback() {
        if (this._boundConnectionStatusUpdated) {
            DebugLog('connectedCallback', 'already connected');
            return;
        }

        DebugLog('connectedCallback', 'connecting');
        this._boundConnectionStatusUpdated = this._connectionStatusUpdated.bind(this);
        Manager.on(ServiceEvents.connection_status_updated, this._boundConnectionStatusUpdated);
    }

    disconnectedCallback() {
        if (this._boundConnectionStatusUpdated === undefined) {
            DebugLog('disconnectedCallback', 'already disconnected');
            return;
        }

        DebugLog('disconnectedCallback', 'disconnecting');
        Manager.off(ServiceEvents.connection_status_updated, this._boundConnectionStatusUpdated);
    }

    async forceRefresh() {
        await Manager.updateConnectionStatuses();
    }

    // connected to the "Authorize" button in the UI
    // main entry point for starting the authorization flow
    async authorizeIntegrationUser() {
        DebugLog('authorizeIntegrationUser called');
        this.loading = true;
        const forceLogin = this.integrationUserType !== 'self';
        const authInfo = await getIntegrationUserAuthorizationUri({ forceLogin });
        /** @type {IntegrationUserAuthorizationUriPayload} */
        const responseData = JSON.parse(authInfo);
        DebugLog('authorizeIntegrationUser responseData', responseData);
        this._connectToMessageListener();
        MessageListener.listenFor(responseData.results.message_origin_uri);
        this._expectedUpdate = true;
        this._connectWindow = openWindow(responseData.results.authorization_uri);
    }

    updateAccountType(event) {
        this.integrationUserType = event.detail.value;
    }

    get loginDisabled() {
        return !this.integrationUserType;
    }

    showUnauthorizeModal() {
        this.template.querySelector('.stripe-modal_confirm-unauthorize').show();
    }

    hideUnauthorizeModal() {
        this.template.querySelector('.stripe-modal_confirm-unauthorize').hide();
    }

    async unauthorizeUser() {
        this.modalLoading = true;
        try {
            const revokedToken = await revokeIntegrationUserAuthorization();
            /** @type {SystemStatusPayload} */
            const responseData =  JSON.parse(revokedToken);
            if (!responseData.isSuccess) {
                this._showToast(responseData.error, 'error', 'sticky');
                return;
            }

            this._expectedUpdate = false;
            Manager.updateConnectionStatuses(responseData.results.service);
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this._showToast(errorMessage, 'error', 'sticky');
        } finally {
            this.modalLoading = false;
            this.hideUnauthorizeModal();
        }
    }

    /**
     *
     * @param {ConnectionStatusChangeEvent} event
     * @private
     */
    _connectionStatusUpdated(event) {
        DebugLog('connectionStatusUpdated', JSON.parse(JSON.stringify(event.detail)));
        DebugLog('connectionStatusUpdated', {expectedUpdate: this._expectedUpdate});
        if (this._expectedUpdate === null) {
            DebugLog('Not expecting an update, bailing...');
            return;
        }

        const statuses = event.detail.statuses;
        const salesforceService = Object.keys(statuses).find(key => key.startsWith(ServiceManagerServices.salesforce + SERVICE_DELIMINATOR));
        this.status = statuses[salesforceService];
        this.isConnected = this.status === ConnectionStatus.fresh || this.status === ConnectionStatus.connected;
        DebugLog('connectionStatusUpdated', 'salesforce status', {
            salesforceService,
            statue: this.status,
            isConnected: this.isConnected,
            isSetup: this.isSetup,
        });

        if (this._expectedUpdate !== undefined && this.status === 'failed') {
            DebugLog('connectionStatusUpdated', 'got failed status');
            this._showErrorToast();
        }

        if (this._expectedUpdate === true && this.status === ConnectionStatus.fresh) {
            DebugLog('connectionStatusUpdated', 'got fresh status');
            this._showToast('Integration User has been successfully connected.', 'success');
        } else if (this._expectedUpdate === false && this.status === ConnectionStatus.disconnected) {
            DebugLog('connectionStatusUpdated', 'got disconnected status');
            this._showToast('Integration User successfully unauthorized.', 'success');
        }

        if (this.isSetup && this.isConnected) {
            DebugLog('connectionStatusUpdated', 'dispatching stepcomplete');
            this.dispatchEvent(new CustomEvent('stepcomplete', { detail: { step: 'integration_user_connection' }} ));
        }

        this._expectedUpdate = undefined;
        this.loading = false;
    }

    _onResponseReceived(event) {
        DebugLog('_onResponseReceived', 'got response', event.detail);
        Debugger.oauthAutoCloseWindow.then((autoClose) => {
            DebugLog('_onResponseReceived', 'oauthAutoCloseWindow', {autoClose});
            if (autoClose) {
                DebugLog('_onResponseReceived', 'oauthAutoCloseWindow', 'closing window');
                this._connectWindow.close();
            } else {
                DebugLog('_onResponseReceived', 'oauthAutoCloseWindow', 'not closing window');
            }
        });
    }

    _onConnectionSuccess(event) {
        DebugLog('_onConnectionSuccess', 'got success', event.detail);
        return Manager.updateConnectionStatuses(event.detail.service);
    }

    _onConnectionError(error) {
        DebugLog('_onConnectionError', 'got error', error);
        this._showErrorToast();
    }

    _onMessageComplete(event) {
        DebugLog('_onConnectionComplete', 'got complete', event.detail);
        this._disconnectFromMessageListener();
    }

    /**
     *
     * @private
     */
    _connectToMessageListener() {
        if (this._boundOnResponseReceived) {
            DebugLog('_connectToMessageListener', 'already initialized');
            return;
        }

        this._boundOnResponseReceived = this._onResponseReceived.bind(this);
        this._boundOnConnectionSuccess = this._onConnectionSuccess.bind(this);
        this._boundOnConnectionError = this._onConnectionError.bind(this);
        this._boundOnMessageComplete = this._onMessageComplete.bind(this);

        MessageListener.on(ListenerEvents.response_received, this._boundOnResponseReceived);
        MessageListener.on(ListenerEvents.connection_successful, this._boundOnConnectionSuccess);
        MessageListener.on(ListenerEvents.error, this._boundOnConnectionError);
        MessageListener.on(ListenerEvents.complete, this._boundOnMessageComplete);
    }

    _disconnectFromMessageListener() {
        if (!this._boundOnResponseReceived) {
            DebugLog('_disconnectFromMessageListener', 'already uninitialized');
            return;
        }

        MessageListener.off(ListenerEvents.response_received, this._boundOnResponseReceived);
        MessageListener.off(ListenerEvents.connection_successful, this._boundOnConnectionSuccess);
        MessageListener.off(ListenerEvents.error, this._boundOnConnectionError);
        MessageListener.off(ListenerEvents.complete, this._boundOnMessageComplete);

        this._boundOnResponseReceived = undefined;
        this._boundOnConnectionSuccess = undefined;
        this._boundOnConnectionError = undefined;
        this._boundOnMessageComplete = undefined;
    }

    _showErrorToast() {
        DebugLog('_showErrorToast', 'Showing error toast');
        this._showToast('There was a problem checking connection status. Please refresh the page and try again.', 'error', 'sticky');
    }

    /**
     *
     * @param {string} url
     * @private
     */
    _openWindow(url) {
        this._connectWindow = openWindow(url);
    }

    _showToast(message, variant, mode) {
        this.dispatchEvent(createToast(message, variant, mode));
    }
}