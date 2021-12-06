import { LightningElement, track, api, wire } from 'lwc';
import saveData from '@salesforce/apex/setupAssistant.saveData';
import isCpqEnabled from '@salesforce/apex/setupAssistant.isCpqEnabled';
import validateConnectionStatus from '@salesforce/apex/setupAssistant.validateConnectionStatus';

export default class OutboundStep extends LightningElement {
    @api isAuthComplete = false;
    @track connectWindow;

    connectedCallback() {
        this.checkIfCPQIsEnabled();
    }

    stripeConnectedAppCallback(isCpqEnabled) {
        if(isCpqEnabled === true) {
            this.validateConnectionStatus(true);
            this.postMessageListener = (event) => {
                if(event.origin === 'https://stripe-force.herokuapp.com' && event.data === 'connectionSuccessful') {
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
        //let rubyAuthURI = 'https://stripe-force.herokuapp.com/auth/salesforce'; //production */
        let rubyAuthURI = 'https://stripe-force.herokuapp.com/auth/salesforcesandbox'; //sandbox
        this.connectWindow = window.open(rubyAuthURI, '"_blank"');
    }

    validateConnectionStatus(isConnectedCallback) {
        this.loading = true;
        validateConnectionStatus({
            isConnectedCallback : isConnectedCallback
        })
        .then(response => {
           let responseData = JSON.parse(response);
            if(responseData.isSuccess) {
                let isConnected = responseData.results.isConnected;
                this.domain = responseData.results.domain;
                if(isConnected === 'fresh') {
                    this.isAuthComplete = true;
                    this.showToast('Authorization successfully completed', 'success');
                } else if(isConnected === 'failed') {
                    this.isAuthComplete = false;
                    this.showToast('Please refresh the page there was an error when checking your connection status', 'error');
                } else if (isConnected === true) {
                    this.isAuthComplete = true;
                } 

            } else { 
                this.showToast(responseData.error, 'error');
            }
        }).catch(error => {
            this.showToast(error.body.message, 'error');
        }).finally(() => {
            this.loading = false;
        });
    }

    checkIfCPQIsEnabled() {
        isCpqEnabled()
        .then(response => {
           let responseData = JSON.parse(response);
            if(responseData.isSuccess) {
                 let isCpqInstalled = responseData.results.isCpqInstalled;
                if(isCpqInstalled === true) {
                    this.stripeConnectedAppCallback(true);
                } else {
                    this.showToast('The CPQ package is not installed in this org', 'error');
                }
            } else { 
                this.showToast(responseData.error, 'error');
            }
        }).catch(error => {
            this.showToast(error.body.message, 'error');
            this.loading = false;
        });
    }

    next(event) {
        if(this.isAuthComplete === true){
            event.stopPropagation();
            saveData({
                setupData: {
                    Steps_Completed__c: JSON.stringify({
                        'C-OUTBOUND-STEP': 1
                    })
                }
            }).then((responseDataString) => {
                let responseData = JSON.parse(responseDataString);

                if(responseData.isSuccess) {
                    this.setupData = responseData.results.setupData;
                    this.dispatchEvent(new CustomEvent('next', {
                        bubbles: true,
                        composed: true
                    }));

                } else {
                    this.showToast(responseData.error, 'error');
                }
            }).catch(error => {
                this.showToast(error, 'error');
            });
        }
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
