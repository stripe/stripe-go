import saveSyncPreferences from '@salesforce/apex/setupAssistant.saveSyncPreferences';
import getSyncPreferences from '@salesforce/apex/setupAssistant.getSyncPreferences';
import getMulticurrencySelectionOptions from '@salesforce/apex/setupAssistant.getMulticurrencySelectionOptions';
import { LightningElement, api, track} from 'lwc';

export default class SyncPreferencesStep extends LightningElement {
    @track stripeConnectRecord;
    @track stripeAccountId;
    @track lastSynced;
    @track defaultCurrency;
    @track syncRecordRetention;
    @track syncStartDate;
    @track apiPercentageLimit;
    @track cpqTermUnit;
    @track isMultiCurrencyEnabled;
    @track isPreferencesUpdated = false;
    @track totalApiCalls; 
    @track currencyOptions;
    @track isCpqInstalled;
    @track intervalOptions = [
        {
            label: 'Month',
            value: 'Month'
        },
        {
            label: 'Day',
            value: 'Day'
        }
    ];
 
    async connectedCallback() {
        try {
            const syncPreferences = await getSyncPreferences();
            this.data = JSON.parse(syncPreferences);
            if(this.data.isSuccess && this.data.results.isConnected) {
                this.stripeAccountId = this.data.results.stripe_account_id;
                this.lastSynced = new Date(this.data.results.last_synced * 1000).toLocaleString(undefined, {year:'numeric', month:'numeric', day: '2-digit', hour: 'numeric', minute:'2-digit', timeZoneName:'short'});
                this.defaultCurrency = this.data.results.default_currency;
                this.syncRecordRetention = this.data.results.sync_record_retention;
                this.syncStartDate = new Date(this.data.results.sync_start_date * 1000).toISOString();
                this.apiPercentageLimit = this.data.results.api_percentage_limit;
                this.cpqTermUnit = this.data.results.cpq_term_interval;
                this.isCpqInstalled = this.data.results.isCpqInstalled;

                const multiCurrencyCheck = await getMulticurrencySelectionOptions();
                this.data = JSON.parse(multiCurrencyCheck);
                if(this.data.isSuccess) {
                    this.totalApiCalls = this.data.results.orgMaxApiLimit;
                    if (this.data.results.isMultiCurrencyEnabled) {
                        this.currencyOptions = this.data.results.supportedISOCodes;
                    }
                } else if(this.data.error) { 
                    this.showToast(this.data.error, 'error', 'sticky');
                }
            } else if(this.data.error) { 
                this.showToast(this.data.error, 'error', 'sticky');
            }

            
        } catch (error) {
            this.showToast(error.message, 'error', 'sticky');
        } finally {
            this.loading = false;
        }
    }

    updateDefaultCurrency(event) {
        let updatedValue = this.enablesave(this.defaultCurrency, event.detail.value)
        this.defaultCurrency = updatedValue;
    }

    updateSyncRecordRetention(event) {
        let updatedValue = this.enablesave(this.syncRecordRetention, event.target.value)
        this.syncRecordRetention = updatedValue;
    }

    updateSyncStartDate(event) {
        let updatedValue = this.enablesave(this.syncStartDate, event.target.value)
        this.syncStartDate = updatedValue;
    }

    updateApiPercentageLimit(event) {
        let updatedValue = this.enablesave(this.apiPercentageLimit, event.target.value)
        this.apiPercentageLimit = updatedValue;
    }

    updateCpqTermInterval(event) {
        let updatedValue = this.enablesave(this.cpqTermUnit, event.detail.value)
        this.cpqTermUnit = updatedValue;
    }


    enablesave(savedValue, updatedValue) {
        if(savedValue !== updatedValue) {
            this.dispatchEvent(new CustomEvent('enablesave'));
            return updatedValue;
        }
        return savedValue;
    }

    get apiLimitCount() {
        return Math.floor(this.totalApiCalls * (this.apiPercentageLimit / 100));
    }

    @api async saveModifiedSyncPreferences() {
        try {
            const updatedSyncPreferences = await saveSyncPreferences({
                cpqTermUnit: this.cpqTermUnit,
                defaultCurrency: this.defaultCurrency,
                syncRecordRetention: this.syncRecordRetention,
                syncStartDate: (new Date(this.syncStartDate).getTime() / 1000),
                apiPercentageLimit: this.apiPercentageLimit
            });
            this.data =  JSON.parse(updatedSyncPreferences);
            if(this.data.isSuccess) {
                this.showToast('Changes were successfully saved', 'success');
            } else {
                this.showToast(this.data.error, 'error', 'sticky');
            }
        } catch (error) {
            this.showToast(error.message, 'error', 'sticky');
        } 
    }

    showToast(message, variant, mode) {
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