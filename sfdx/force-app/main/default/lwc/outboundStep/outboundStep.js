import { LightningElement, track, api, wire } from 'lwc';
import saveData from '@salesforce/apex/setupAssistant.saveData';
import validateConnectionStatus from '@salesforce/apex/setupAssistant.validateConnectionStatus';

export default class OutboundStep extends LightningElement {
    @api isAuthComplete = false;
    @track connectWindow;
    @track nextDisabled = true;
    @track isSandbox;

    connectedCallback() {
        this.stripeConnectedAppCallback(true);
    }

    stripeConnectedAppCallback(isCpqEnabled) {
        const rubyServiceAuthOriginURI = 'https://stripe-force.herokuapp.com'
        if(isCpqEnabled === true) {
            this.validateConnectionStatus(true);
            this.postMessageListener = (event) => {
                if(event.origin === rubyServiceAuthOriginURI && event.data === 'connectionSuccessful') {
                    this.connectWindow.close()
                    this.validateConnectionStatus(false);
                }  
            } 
            window.addEventListener("message", this.postMessageListener.bind(this));
        }
    }

    disconnectedCallback() {
        window.removeEventListener('message', this.postMessageListener);
    }

    connectToStripe() {
        const rubyAuthURI = 'https://stripe-force.herokuapp.com/auth/'; 
        this.connectWindow = this.isSandbox ? window.open(rubyAuthURI+'salesforcesandbox', '"_blank"') : window.open(rubyAuthURI+'salesforce', '"_blank"');
    }

    async validateConnectionStatus(isConnectedCallback) {
        this.loading = true;
        try {
            const validateConnection = await validateConnectionStatus({
                isConnectedCallback : isConnectedCallback
            }); 
            this.data =  JSON.parse(validateConnection);
            if(this.data.isSuccess) {
                let isConnected = this.data.results.isConnected;
                this.isSandbox = this.data.results.isSandbox
                if(isConnected === 'fresh') {
                    this.isAuthComplete = true;
                    this.showToast('Authorization successfully completed', 'success');
                    this.dispatchEvent(new CustomEvent('fetchpicklistvalues'));
                } else if(isConnected === 'notConnected') {
                    this.isAuthComplete = false;
                    this.showToast('Your authorization expired, please reauthorize your stripe & salesforce accounts', 'error');
                }else if(isConnected === 'failed') {
                    this.isAuthComplete = false;
                    this.showToast('Please refresh the page there was an error when checking your connection status', 'error');
                } else if (isConnected === true) {
                    this.isAuthComplete = true;
                } 
            } else { 
                this.showToast(this.data.error, 'error');
            }
        } catch (error) {
            this.showToast(error.message, 'error');
        } finally {
            this.nextDisabled = !this.isAuthComplete
        }
    }

    async next(event) {
        if(this.isAuthComplete === true) {
            event.stopPropagation();
            try {
                const saveSetupData = await saveData({
                    newSetupDataRec: {
                        Steps_Completed__c: JSON.stringify({
                            'C-OUTBOUND-STEP': 1
                        })
                    },
                    isSetupComplete: false
                });
                this.data =  JSON.parse(saveSetupData);
                if(this.data.isSuccess) {
                    this.setupData = this.data.setupData;
                    this.dispatchEvent(new CustomEvent('next', {
                        bubbles: true,
                        composed: true
                    }));

                } else {
                    this.showToast(this.data.error, 'error');
                }
            } catch (error) {
                this.showToast(error.message, 'error');
            }
        }
    }

    back() {
        this.dispatchEvent(new CustomEvent('exit', {
            bubbles: true,
            composed: true
        }));
    }


    showToast(message, variant, mode) {
        this.dispatchEvent(new CustomEvent('showtoast', {
            bubbles: true,
            composed: true,
            detail: { 
                toast: {
                    message: message,
                    variant: variant,
                    mode: mode
                }
            }
        })); 
    }
}
