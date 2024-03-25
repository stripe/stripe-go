import validateConnectionStatus from '@salesforce/apex/setupAssistant.validateConnectionStatus';
import validateSharedState from '@salesforce/apex/setupAssistant.validateSharedState';
import { LightningElement, track, api } from 'lwc';
import { getErrorMessage } from 'c/utils'
import Debugger from "c/debugger";
import { ConnectionStatus, TextStrings } from "c/systemStatusUtils";
const DebugLog = Debugger.withContext('SystemConnectionsStep');

function failedCheckResponse() {
    return {
        salesforceStatus: ConnectionStatus.failed,
        stripeStatus: ConnectionStatus.failed,
        isComplete: false,
    };
}


let platformUri = '';
let salesforceSandboxUriFragment = '';
let salesforceProductionUriFragment = '';
let stripeLiveUriFragment = '';
let stripeTestUriFragment = '';
let isSandbox = false;
let salesforceNamespace = '';
let state = '';

export default class SystemConnectionsStep extends LightningElement {
    @track isSalesforceConnected = false;
    @track isStripeConnected = false;
    @track modalLoading = false;
    @track salesforceStatus = ConnectionStatus.loading;
    @track stripeStatus = ConnectionStatus.loading;
    @api hideAction = false;
    _listenerInitialized = false;
    _connectWindow = null;
    _listening = false;

    async connectedCallback() {
        this._initializeMessageListener();
        await this._checkConnectionStatus();
    }

    disconnectedCallback() {
        this._destroyMessageListener();
    }  

    // connected to the "Authorize" button in the UI
    // main entry point for starting the authorization flow
    connectToSalesforce() {
        let oauthConnectionURL = platformUri;

        if (isSandbox) {
            oauthConnectionURL += salesforceSandboxUriFragment;
        } else {
            oauthConnectionURL += salesforceProductionUriFragment;
        }

        // TODO should be removed since the namespace is defined via the post install ste[p]
        oauthConnectionURL += "?salesforceNamespace=" + salesforceNamespace + "&state=" + state;
        this._listening = true;

        this._openWindow(oauthConnectionURL);
    }

    connectToStripe() {
        this.showConnectToStripeModal();
    }

    connectToStripeLiveMode() {
        return this._connectToStripe(true);
    }

    connectToStripeTestMode() {
        return this._connectToStripe(false);
    }

    /**
     *
     * @param {boolean} isLive
     */
    _connectToStripe(isLive) {
        let oauthConnectionURL = platformUri;

        if (isLive) {
            oauthConnectionURL += stripeLiveUriFragment;
        } else {
            oauthConnectionURL += stripeTestUriFragment;
        }

        oauthConnectionURL += "?state=" + state;
        this._listening = true;
        this._openWindow(oauthConnectionURL);
    }

    get isComplete() {
        return this.isSalesforceConnected && this.isStripeConnected;
    }

    showConnectToStripeModal() {
        this.template.querySelector('.stripe-modal-connect-stripe-account').show();
    }

    hideConnectToStripeModal() {
        this.template.querySelector('.stripe-modal-connect-stripe-account').hide();
    }

    /**
     *
     * @param {string} [systemToCheck]
     * @return {Promise<SystemConnectionStatusResults>}
     */
    async _validateConnectionStatus(systemToCheck) {
        DebugLog('_fetchIntegrationUserStatus systemToCheck', systemToCheck);
        const validateConnection = await validateConnectionStatus({
            systemToCheck : systemToCheck
        });
        /** @type {SystemConnectionStatusesPayload} */
        const responseData = JSON.parse(validateConnection);

        if (responseData.isSuccess) {
            const results = responseData.results;
            const config = results.config;
            DebugLog('_fetchIntegrationUserStatus config', config);
            isSandbox = config.is_sandbox;
            salesforceNamespace = config.salesforce_namespace;
            salesforceSandboxUriFragment = config.salesforce_sandbox_uri_fragment;
            salesforceProductionUriFragment = config.salesforce_production_uri_fragment;
            platformUri = config.platform_uri;
            state = config.state;
            stripeTestUriFragment = config.stripe_test_uri_fragment;
            stripeLiveUriFragment = config.stripe_live_uri_fragment;
            delete results.config;
            return responseData.results;
        } else {
            throw new Error(responseData.error);
        }
    }

    async _checkConnectionStatus(systemToCheck) {
        this.loading = true;
        try {
            const results = await this._validateConnectionStatus(systemToCheck);
            DebugLog('_checkConnectionStatus', '_checkConnectionStatus results', results);
            this.stripeStatus = results.stripe;
            this.salesforceStatus = results.salesforce;

            if (this.stripeStatus === 'failed' && this.salesforceStatus === 'failed') {
                this.isSalesforceConnected = false;
                this.isStripeConnected = false;
                DebugLog('_checkConnectionStatus', 'both statuses failed');
                this._showToast('There was a problem checking connection status. Please refresh the page and try again.', 'error', 'sticky');
                this.loading = false;
                return failedCheckResponse();
            }

            const oldStripeState = this.isStripeConnected;
            const oldSalesforceState = this.isSalesforceConnected;

            this._checkConnectionResult(this.stripeStatus, TextStrings.stripe, TextStrings.salesforce, (v) => this.isStripeConnected = v);
            this._checkConnectionResult(this.salesforceStatus, TextStrings.salesforce, TextStrings.stripe, (v) => this.isSalesforceConnected = v);

            const hasChanged = oldStripeState !== this.isStripeConnected || oldSalesforceState !== this.isSalesforceConnected;

            const details = {
                stripeStatus: this.stripeStatus,
                salesforceStatus: this.salesforceStatus,
                hasChanged,
                isComplete: this.isComplete,
            }

            this.dispatchEvent(new CustomEvent('systemsconnected', { detail: details }));
            this.loading = false;
            return details;
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this._showToast(errorMessage, 'error', 'sticky');
        }

        this.loading = false;
        return failedCheckResponse();
    }

    async _postMessageListener(event) {
        if (!this._listening) {
            return;
        }
        this._listening = false;
        DebugLog('Got event: ' + event);
        if (event.origin !== platformUri) {
            DebugLog('_postMessageListener', "bad post message origin");
            return;
        }

        this._connectWindow.close();

        const [ result, identifier, state ] = event.data.split('-');
        DebugLog('_postMessageListener', 'got response', { result, identifier, state });
        const [ system, id ] = identifier.split('_');
        if (system === 'stripe') {
            this.hideConnectToStripeModal();
        }

        if (result === 'connection_successful') {
            return validateSharedState({ state: state })
                .then((response) => {
                    DebugLog('_postMessageListener', 'connectionSetupComplete', response);
                    return this._checkConnectionStatus(system);
                })
                .catch((error) => {
                    DebugLog('_postMessageListener', 'Caught error', error);
                    this._showToast('There was a problem checking connection status. Please refresh the page and try again.', 'error', 'sticky');
                });
        } else {
            DebugLog('_postMessageListener', 'Result was not successful', result);
            this._showToast('There was a problem checking connection status. Please refresh the page and try again.', 'error', 'sticky');
        }
    }

    _initializeMessageListener() {
        if (this._listenerInitialized) {
            return;
        }

        this._boundPostMessageListener = this._postMessageListener.bind(this);

        window.addEventListener("message", this._boundPostMessageListener);
        this._listenerInitialized = true;
    }

    _destroyMessageListener() {
        if (this._listenerInitialized === false) {
            return;
        }

        window.removeEventListener('message', this._boundPostMessageListener);
        delete this._boundPostMessageListener;
        this._listenerInitialized = false;
    }

    /**
     *
     * @param {ConnectionStatus} value
     * @param {string} from
     * @param {string} to
     * @param {function(boolean)} setBoolValue
     * @private
     */
    _checkConnectionResult (value, from, to, setBoolValue) {
        if (value === ConnectionStatus.fresh || value === ConnectionStatus.connected) {
            setBoolValue(true);

            if (value === ConnectionStatus.fresh) {
                this._showToast(`${from} to ${to} authorization successfully completed.`, 'success');
            }
        } else {
            setBoolValue(false);
        }
    }

    /**
     *
     * @param {string} url
     * @private
     */
    _openWindow(url) {
        this._connectWindow = window.open(url, '"_blank"');
    }

    _showToast(message, variant, mode) {
        this.dispatchEvent(new CustomEvent('showtoast', {
            bubbles: true,
            composed: true,
            detail: {
                toast: {
                    message: message,
                    variant: variant ? variant : 'info',
                    mode: mode ? mode : 'dismissible'
                }
            }
        }));
    }
}
