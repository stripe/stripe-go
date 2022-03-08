import getRelatedSyncRecords from '@salesforce/apex/relatedSyncRecordsController.getRelatedSyncRecords';
import { LightningElement, api, track } from "lwc";
import { NavigationMixin } from 'lightning/navigation';

export default class RelatedSyncRecords extends NavigationMixin(LightningElement) {
  @api recordId;
  @track addXMore = 10;
  @track listOfSyncRecords = [];
  @track listOfFailedSyncRecords = [];
  @track listOfDisplayFailedRecords = [];
  @track listOfDisplayRecords = [];
  @track loading = false;
  @track errorMessage = '';
  @track missingPermSet = false;

  async connectedCallback() {
        try {
            this.loading = true;
            const allRelatedSyncRecords = await getRelatedSyncRecords({
            recordId: this.recordId
            });
            this.data = JSON.parse(allRelatedSyncRecords);
            if(this.data.isSuccess && this.data.results.syncRecordList.length) {
                let syncRecords = this.data.results.syncRecordList;
                for (let i = 0; i < syncRecords.length; i++) {           
                    //created date is not writable in apex so we format the date here 
                    var formattedCreatedDate = new Date(syncRecords[i].CreatedDate);
                    syncRecords[i].CreatedDate = formattedCreatedDate;
                    Object.defineProperties(syncRecords[i], {
                        hasError: {
                            get() {return this.Resolution_Status__c === 'Error';}
                        },
                        statusClasslist: {
                            get() {
                                let classList = 'slds-badge';
                                if (this.Resolution_Status__c === 'Error') {
                                    classList+= ' slds-theme_error';
                                } else if (this.Resolution_Status__c === 'Success' || this.Resolution_Status__c === 'Resolved') {
                                    classList+= ' slds-theme_success';
                                } else if (this.Resolution_Status__c === 'Skipped') {
                                    classList+= ' slds-badge_inverse';
                                }
                                return classList;
                            } 
                        }
                    });
                } 
                this.listOfSyncRecords = syncRecords;
                this.listOfFailedSyncRecords = this.listOfSyncRecords.filter(syncRecords => syncRecords.Resolution_Status__c === 'Error')

                this.addRecordsToDisplayList(this.listOfDisplayRecords, this.listOfSyncRecords)   
                this.addRecordsToDisplayList(this.listOfDisplayFailedRecords, this.listOfFailedSyncRecords);    
            } else if (this.data.results.missingPermSet) {
                this.missingPermSet = true;
            } else if(this.data.error) { 
                this.errorMessage = this.data.error;
            } 
        } catch (error) {
            this.errorMessage = error.message;
        } finally {
            this.loading = false;
        }
    }
    loadMoreRecords() {
        this.addRecordsToDisplayList(this.listOfDisplayRecords, this.listOfSyncRecords); 
    }

    loadMoreErrorRecords() {
        this.addRecordsToDisplayList(this.listOfDisplayFailedRecords, this.listOfFailedSyncRecords);    
    }

    addRecordsToDisplayList(displayRecordsList, syncRecordList ) { 
        if(displayRecordsList.length < syncRecordList.length) {
            let startIndex = displayRecordsList.length;
            for (let i = startIndex; i < this.addXMore + startIndex; i++) {     
                if(syncRecordList[i]) {
                    displayRecordsList.push(syncRecordList[i]);
                } else {
                    break
                }
            }
        }
    }

    navToSyncRecord(event) {
        let index = event.currentTarget.dataset.index;
        this[NavigationMixin.Navigate]({
            type: 'standard__recordPage',
            attributes: {
                recordId: this.listOfSyncRecords[index].Id,
                actionName: 'view'
            }
        })
    }

    toggleRowDetails(event) {
        event.currentTarget.closest('tr.stripe-row').classList.toggle('slds-is-open');
    }

    resetListScroll(event) {
        const activeList = event.currentTarget.querySelector('.stripe-list-container');
        if(activeList) {
            activeList.scrollTop = 0;
        }
    }

    get allRecordsTabLabel() {
        return `All Records (${this.listOfSyncRecords ? this.listOfSyncRecords.length : '0'})`;
    }

    get issuesTabLabel() {
        return `Issues (${this.listOfDisplayFailedRecords ? this.listOfDisplayFailedRecords.length : '0'})`;
    }

    get showComponentMessage() {
        return this.errorMessage || this.missingPermSet;
    }
} 