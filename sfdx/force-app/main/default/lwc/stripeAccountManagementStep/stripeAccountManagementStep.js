import getStripeAuthorizationUri from '@salesforce/apex/setupAssistant.getStripeAuthorizationUri';
import addStripeAccount from '@salesforce/apex/setupAssistant.addStripeAccount';
import getStripeAccounts from '@salesforce/apex/setupAssistant.getStripeAccounts';
import setStripeAccountAsDefault from '@salesforce/apex/setupAssistant.setStripeAccountAsDefault';
import deleteStripeAccount from '@salesforce/apex/setupAssistant.deleteStripeAccount';
import getNamespacePrefix from '@salesforce/apex/setupAssistant.getNamespacePrefix';
import {LightningElement, track, api} from 'lwc';
import { getErrorMessage, createToast, openWindow } from 'c/utils'
import AddStripeAccountModal from 'c/addStripeAccountModal';
import Debugger from "c/debugger";
import {MessageListener, ConnectionStatus, ListenerEvents} from "c/systemStatusUtils";
import { Manager, ServiceEvents } from "c/serviceManager";
import LightningConfirm from "lightning/confirm";

const DebugLog = Debugger.withContext('StripeAccountManagementStep');

export default class StripeAccountManagementStep extends LightningElement {
    @track isConnected = false;
    @track loading = false;
    @track modalLoading = false;
    @track status = ConnectionStatus.loading;
    @track accountType;
    @track accounts = [];
    @track areMultipleStripeAcctsConnected = false;
    @api hideAction = false;
    @api isSetup = false;
    _connectWindow = null;
    _expectedUpdate = undefined;
    accountsByExternalId = {};
    contentLoading = new CustomEvent('contentloading');
    contentLoadingComplete = new CustomEvent('contentloadingcomplete');

    async handleRowAction(event) {
        const context = {
            action: event.detail.value,
            id: event.target.dataset.id,
        };
        context.account = this.accountsByExternalId[context.id];

        DebugLog('handleRowAction', context)

        if (context.account === undefined || context.account === null) {
            DebugLog('handleRowAction', 'account not found', context);
            this._showToast('Account not found', 'error', 'sticky');
            return;
        }

        context.functionArgs = {stripeAccountId: context.account.Stripe_ID__c, isLive: context.account.Is_Live_Mode__c};
        DebugLog('handleRowAction', 'before call', context);

        try {
            let fn = null;
            let confConfig = {
                theme: 'shade',
                label: 'Are you sure?',
            };
            if (context.action === 'SetPrimary') {
                fn = setStripeAccountAsDefault;
                confConfig.message = `Are you sure you want to set ${context.account.Stripe_ID__c} as the primary account?`;
            } else if (context.action === 'Delete') {
                fn = deleteStripeAccount;
                confConfig.message = `Deleting this Stripe account will disable you from syncing to this Stripe account. Are you sure you want to delete ${context.account.Stripe_ID__c}?`;
            } else {
                DebugLog('handleRowAction', 'unknown action', context.action);
                this._showToast('Unknown action', 'error', 'sticky');
                return;
            }

            const conf = await LightningConfirm.open(confConfig);
            if (conf === false) {
                return;
            }

            this.dispatchEvent(this.contentLoading);

            context.result = await fn(context.functionArgs);

            DebugLog('handleRowAction', 'after call', context);

            if (context.result !== true) {
                this._showToast('There was a problem updating the account. Please refresh and try again.', 'error', 'sticky');
                return;
            }

            await this.getStripeAccounts();
            this.dispatchEvent(this.contentLoadingComplete);
        } catch (e) {
            const msg = getErrorMessage(e);
            this._showToast(msg, 'error', 'sticky');
        }
    }

    async connectedCallback() {
        if (this._boundConnectionStatusUpdated) {
            DebugLog('connectedCallback', 'already connected');
            return;
        }

        DebugLog('connectedCallback', 'connecting');
        this.dispatchEvent(this.contentLoading);
        await this.getStripeAccounts();
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

    async getStripeAccounts() {
        try {
            const acctsByExtId = {};
            const namespacePrefix = await getNamespacePrefix();

            const stripeAccounts = await getStripeAccounts();
            stripeAccounts.forEach(acct => {
                acct.Is_Live_Mode__c = acct[namespacePrefix + 'Is_Live_Mode__c'];
                acct.livemode = acct.Is_Live_Mode__c === true ? 'live' : 'test';
                acct.Stripe_ID__c = acct[namespacePrefix + 'Stripe_ID__c'];
                acct.Is_Primary__c = acct[namespacePrefix + 'Is_Primary__c'];
                acct.External_ID__c = acct[namespacePrefix + 'External_ID__c'];
                acct.Connection_Status__c = acct[namespacePrefix + 'Connection_Status__c'];
                acct.CSAC_ID__c = acct[namespacePrefix + 'CSAC_ID__c'];
                acctsByExtId[acct.External_ID__c] = acct;
            });
            this.accounts = stripeAccounts;
            this.accountsByExternalId = acctsByExtId;
            this.areMultipleStripeAcctsConnected = this.accounts.length > 1;
            return this.accounts;
        } catch (e) {
            const err = getErrorMessage(e);
            this.showToast(err, 'error', 'sticky');
            return null;
        }
    }

    // connected to the "Authorize" button in the UI
    // main entry point for starting the authorization flow
    /**
     *
     * @return {Promise<void>}
     */
    async authorizeIntegrationUser() {
        DebugLog('authorizeIntegrationUser called');
        this.loading = true;
        const result = await AddStripeAccountModal.open({ size: 'small' });
        if (result === null || result === undefined) {
            return;
        }

        const authInfo = await getStripeAuthorizationUri({ liveMode: result === 'live' });
        /** @type {IntegrationUserAuthorizationUriPayload} */
        const responseData = JSON.parse(authInfo);
        DebugLog('authorizeIntegrationUser responseData', responseData);
        this._connectToMessageListener();
        MessageListener.listenFor(responseData.results.message_origin_uri);
        this._expectedUpdate = true;
        this.dispatchEvent(this.contentLoading);
        this._connectWindow = openWindow(responseData.results.authorization_uri);
    }

    get loginDisabled() {
        return !this.accountType;
    }

    updateAccountType(event) {
        this.accountType = event.detail.value;
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

        // const statuses = event.detail.statuses;
        // const stripeServices = Object.keys(statuses).filter(key => key.startsWith(ServiceManagerServices.stripe + SERVICE_DELIMINATOR));
        // const existingServices = {};
        // this.accounts.forEach(acct => existingServices[acct.id] = {account: acct, found: false});
        // stripeServices.forEach(svcName => {
        //     const [svcType, svcId, svcProps] = svcName.split(SERVICE_DELIMINATOR);
        //     const svcStatus = statuses[svcName];
        //     const svc = {
        //         id: svcName,
        //         account_id: svcId,
        //         type: svcType,
        //         props: svcProps,
        //         status: svcStatus,
        //         livemode: svcProps
        //     };
        //     if (existingServices[svcId]) {
        //         existingServices[svcId].found = true;
        //         existingServices[svcId].account.status = svc.status;
        //     } else {
        //         existingServices[svcId] = {found: true, account: svc};
        //     }
        // });
        //
        // const foundKeys = Object.keys(existingServices).filter(key => existingServices[key].found);
        // this.accounts = foundKeys.map(key => existingServices[key].account);

        this.dispatchEvent(this.contentLoadingComplete);
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

    async _onConnectionSuccess(event) {
        DebugLog('_onConnectionSuccess', 'got success', event.detail);
        await addStripeAccount({ state : event.detail.raw_state });
        await Manager.updateConnectionStatuses(event.detail.service);
        await this.getStripeAccounts();
        this.dispatchEvent(this.contentLoadingComplete);
    }

    _onConnectionError(error) {
        DebugLog('_onConnectionError', 'got error', error);
        this.dispatchEvent(this.contentLoadingComplete);
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
        DebugLog('_showErrorToast', 'showing error toast');
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