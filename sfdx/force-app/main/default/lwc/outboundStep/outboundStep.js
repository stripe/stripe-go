import { LightningElement, track, api, wire } from 'lwc';
import saveData from '@salesforce/apex/setupAssistant.saveData';
import validateConnectionStatus from '@salesforce/apex/setupAssistant.validateConnectionStatus';
import { NavigationMixin, CurrentPageReference } from 'lightning/navigation';

export default class OutboundStep extends NavigationMixin(LightningElement) {
    @api isAuthComplete = false;
    @track currentDomain;
    @wire(CurrentPageReference)
    getpageRef(pageRef) {
        console.log('data => ', JSON.stringify(pageRef));
        this.currentDomain =  pageRef;
    }
    

    connectedCallback() {
        this.template.addEventListener('next', this.next.bind(this));
        this.stripeConnectedAppCallback(true);
        console.log('IS this being hit');
        console.log(this.currentDomain);
        console.log( window.location.href);
        console.log( window.location.hostname);
        //console.log( window.parent.location.href);
        this.postMessageListener = (event) => {
            console.log('event')
            console.log(event)
            if(event.origin === 'https://stripe-force.herokuapp.com' ) {
                
                //this.stripeConnectedAppCallback(false);
            } 
        } 
       window.addEventListener('message', this.postMessageListener);
    }

    disconnectedCallback() {
        window.removeEventListener('message', this.postMessageListener);
    }
    connectToStripe() {
        let message = {hostName : 'https://'+window.location.hostname.replace("--c.visualforce", ".lightning.force")};
        //let rubyAuthURI = 'https://stripe-force.herokuapp.com/auth/salesforce'; //production
        let rubyAuthURI = 'https://stripe-force.herokuapp.com/auth/salesforcesandbox'; //sandbox
        let connectWindow = window.open(rubyAuthURI, '"_blank"');
        connectWindow.postMessage(message, rubyAuthURI);
        console.log('message');
        console.log(message);
    }

    stripeConnectedAppCallback(isConnectedCallback) {
        this.loading = true;
        validateConnectionStatus({
            isConnectedCallback : isConnectedCallback
        })
        .then(response => {
           let responseData = JSON.parse(response);
            if(responseData.isSuccess) {
                let isConnected = responseData.results.isConnected;
                if(isConnected === 'fresh') {
                    this.isAuthComplete = true;
                    this.showToast('Authorization successfully completed', 'success');
                }
                if(isConnected === true)this.isAuthComplete = true;
            } else { 
                this.showToast(responseData.error, 'error');
            }
        }).catch(error => {
            this.showToast(error.body.message, 'error');
        }).finally(() => {
            this.loading = false;
        });
    }

    next(event) {
        event.stopPropagation();
        saveData({
            setupData: {
                Steps_Completed__c: JSON.stringify({
                    'C-OUTBOUND-STEP': 1
                }),
                Live_Mode__c: this.setupData.Live_Mode__c
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