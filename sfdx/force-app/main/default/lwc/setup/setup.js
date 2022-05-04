import companyLogo from '@salesforce/resourceUrl/companyLogo';
import checkUserPermissions from '@salesforce/apex/setupAssistant.checkUserPermissions';
import getSetupData from '@salesforce/apex/setupAssistant.getSetupData';
import setOrgType from '@salesforce/apex/utilities.setOrgType';
import saveData from '@salesforce/apex/setupAssistant.saveData';
import { LightningElement, track, api } from 'lwc';
import illustrations from '@salesforce/resourceUrl/stripe_setupIllustrations';

export default class FirstTimeSetup extends LightningElement {
    systemConnectionsIllustration = illustrations + '/stripe_illustration_systemConnections.svg';
    dataMappingIllustration = illustrations + '/stripe_illustration_dataMapping.svg';
    syncPreferencesIllustration = illustrations + '/stripe_illustration_syncPreferences.svg';
    @track contentShown = false;
    @track logoUrl = companyLogo;
    @track activeStepIndex = 0;
    @track activeSectionIndex = undefined;
    @track showIntro = false;
    @track setupComplete = false;
    @track setupStarted = false;
    @track nextDisabled = true;
    @track saveDisabled = true;
    @track loading = true;
    @track contentLoading = false;
    @track stepName;
    @track steps = [
        {
            title: 'Update Page Layouts',
            name: 'C-ORG-SETTINGS-STEP',
            orderIndex: 1,
            isComplete: false,
            isActive: false
        },
        {
            title: 'Connect Stripe and Salesforce',
            name: 'C-SYSTEM-CONNECTIONS-STEP',
            orderIndex: 2,
            isComplete: false,
            isActive: false
        },
        {
            title: 'Define Data Mapping',
            name: 'C-DATA-MAPPING-STEP',
            orderIndex: 3,
            isComplete: false,
            isActive: false
        },
        {
            title: 'Configure Sync Preferences',
            name: 'C-SYNC-PREFERENCES-STEP',
            orderIndex: 4,
            isComplete: false,
            isActive: false
        }
    ];
    @track missingPermissions = {};
    @track hasMissingPermissions = false;

    @track _showSetupLanding = true;
    get showSetupLanding() {
        return this.setupComplete ? false : this._showSetupLanding;
    }

    set showSetupLanding(value) {
        this._showSetupLanding = value;
    }

    get showSettingsLanding() {
        return this.activeSectionIndex == undefined;
    }

    get backDisabled() {
        return this.activeStepIndex === 0;
    }

    get isFinalStep() {
        return this.activeStepIndex === (this.steps.length - 1);
    }

    get nextLabel() {
        return this.isFinalStep ? 'Finish' : 'Next';
    }

    get sidebarClasslist() {
        return 'slds-col slds-size_1-of-4' + (this.setupComplete ? ' stripe-sidebar_navigation' : ' stripe-sidebar_progress');
    }

    get contentClasslist() {
        return ('stripe-page-container slds-container slds-container_center' + (this.loading ? ' slds-hide' : ''));
    }

    get stepContentClasslist() {
        return this.showSetupLanding ? 'slds-hide' : '';
    }

    get settingsContentClasslist() {
        return this.showSettingsLanding ? 'slds-hide' : '';
    }

    get landingButtonLabel() {
        return this.setupStarted ? 'Continue' : 'Get Started';
    }

    enableNext() {
        this.nextDisabled = false;
    }

    enableSave() {
        this.saveDisabled = false;
    }

    completeSave(event) {
        this.contentLoading = false;
        if(event.detail.saveSuccess && this.setupComplete === true) {
            this.saveDisabled = true;
        } else if (event.detail.saveSuccess && this.setupComplete === false) {
            this.nextNavigate();
        } else {
            this.showSetupToast('There was a problem saving your changes. Make sure that all updated values are valid before saving.', 'error', 'sticky');
        }
    }

    getmappingconfigurations() {
        this.template.querySelector('c-data-mapping-step').getPicklistValuesForMapper(true, '');
        this.template.querySelector('c-sync-preferences-step').connectedCallback();  
        this.nextDisabled = false;
    }

    async connectedCallback() { 
        try {
            const userPermissionCheck = await checkUserPermissions();
            const userPermissionResponseData = JSON.parse(userPermissionCheck);
            if (userPermissionResponseData.error) { 
                this.showSetupToast(userPermissionResponseData.error, 'error', 'sticky');
                return;
            }

            const permissionIssueMap = userPermissionResponseData.results.permissionIssueMap;

            if (!permissionIssueMap.isPermSetAssigned || permissionIssueMap.isSystemPermissionMissing || permissionIssueMap.isObjectPermissionMissing) {
                this.missingPermissions = permissionIssueMap;
                this.hasMissingPermissions = true;
                return;               
            } 
            this.fetchSetupData();

            const setOrganizationType = await setOrgType();
            const orgTypeResponseData =  JSON.parse(setOrganizationType);
                
            if (orgTypeResponseData.error) {
                this.showSetupToast(orgTypeResponseData.error, 'error', 'sticky');
                return;
            }
            
        } catch (error) {
            this.showSetupToast(error.message, 'error', 'sticky');
        } finally {
            this.loading = false;
        }
    }

    async fetchSetupData() {
        this.steps[this.activeStepIndex].isActive = true;
        try {
            const setupData = await getSetupData();
            const responseData = JSON.parse(setupData);
            if (responseData.error) { 
                this.showSetupToast(responseData.error, 'error', 'sticky');
                return;
            }

            if (!responseData.results.isConnected) { 
                return;
            }

            this.nextDisabled = false;
            this.setupComplete = responseData.results.setupData.isSetupComplete__c;
            let completedSteps = JSON.parse(responseData.results.setupData.Steps_Completed__c);
            if (Object.keys(completedSteps).length <= 0) {
                return;
            }
            
            this.setupStarted = true;
            this.steps[this.activeStepIndex].isComplete = true;
            this.steps[this.activeStepIndex].isActive = false;

            for (const step in completedSteps) {
                this.steps[this.activeStepIndex].isComplete = true;

                if (this.activeStepIndex < this.steps.length - 1) {
                    this.showNextStep();
                }
            }

        } catch (error) {
            this.showSetupToast(error.message, 'error', 'sticky');
        } finally {
            this.showContent();
        }
    }

    showContent() {
        if(this.setupComplete) {
            if(this.activeSectionIndex) {
                const activeSection = this.template.querySelector('c-step[data-index="' + this.activeSectionIndex + '"]');
                if(activeSection) {
                    activeSection.classList.remove('slds-hide');
                }
            }
            this.contentShown = true;
        } else {
            const activeStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
            if(activeStep) {
                activeStep.classList.remove('slds-hide');
            }
            this.contentShown = true;
        }
    }    

    save() {
        this.contentLoading = true;
        if(this.activeSectionIndex == 1) {
            this.template.querySelector('c-data-mapping-step').saveCongfiguredMappings();
        } else if(this.activeSectionIndex == 2) {
            this.template.querySelector('c-sync-preferences-step').saveModifiedSyncPreferences();
        }
    }

    next() { 
        this.contentLoading = true;
        this.stepName = this.steps[this.activeStepIndex].name;
        if(this.stepName === 'C-DATA-MAPPING-STEP') {
            this.template.querySelector('c-data-mapping-step').saveCongfiguredMappings();
        } else if(this.stepName === 'C-SYNC-PREFERENCES-STEP') {
            this.template.querySelector('c-sync-preferences-step').saveModifiedSyncPreferences();
        } else {
            this.nextNavigate();
        }       
    }

    async nextNavigate() {
        try {
            if(this.activeStepIndex <= (this.steps.length - 1)) {
                 
                let stepNumber = this.activeStepIndex + 1;
                let isLastStep = false;
                if (stepNumber === (this.steps.length)) {
                    isLastStep = true
                    this.steps[this.activeStepIndex].isActive = false;
                    this.steps[this.activeStepIndex].isComplete = true;
                    this.nextDisabled = true;     
                    this.setupComplete = true;
                    this.showSetupToast('Setup completed', 'success');
                }
            
                const saveSetupData = await saveData({
                    newSetupDataRec: {
                        Steps_Completed__c: JSON.stringify({
                            [this.stepName]: stepNumber
                        })
                    },
                    isSetupComplete: isLastStep
                });
                this.data =  JSON.parse(saveSetupData);
                if(this.data.isSuccess) {
                    this.setupData = this.data.setupData;
                } else {
                    this.showSetupToast(this.data.error, 'error', 'sticky');
                }
            }
        } catch (error) {
            this.showSetupToast(error.message, 'error', 'sticky');
        } finally {
            if(this.activeStepIndex < (this.steps.length - 1)) {
                this.showNextStep();
            }
            this.contentLoading = false;
        }
    }

    back() {
        this.contentLoading = true;
        if(this.activeStepIndex > 0) {
            const lastActiveStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
            this.activeStepIndex--;
            const newActiveStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
            if(lastActiveStep && newActiveStep) {
                this.steps[lastActiveStep.dataset.index].isActive = false;
                this.steps[lastActiveStep.dataset.index].isComplete = false;
                this.steps[newActiveStep.dataset.index].isActive = true;
                this.steps[newActiveStep.dataset.index].isComplete = false;
                lastActiveStep.classList.add('slds-hide');
                newActiveStep.classList.remove('slds-hide');
            }
        }
        this.contentLoading = false;
    }

    sectionNavigate(event) {
        if(this.saveDisabled) {
            let targetSection = event.currentTarget;
            if(targetSection.dataset.index !== this.activeSectionIndex) {
                const currentActive = this.template.querySelector('.stripe-navigation-item_active');
                if(currentActive) {
                    // Clear previous active section
                    currentActive.classList.remove('stripe-navigation-item_active');
                    const previousSectionContent = this.template.querySelector('c-step[data-index="' + this.activeSectionIndex + '"]');
                    if(previousSectionContent) {
                        previousSectionContent.classList.add('slds-hide');
                    }
                }
                // Make selected section active 
                targetSection.classList.add('stripe-navigation-item_active');
                this.activeSectionIndex = targetSection.dataset.index;
                const targetSectionContent = this.template.querySelector('c-step[data-index="' + this.activeSectionIndex + '"]');
                if(targetSectionContent) {
                    targetSectionContent.classList.remove('slds-hide');
                }
            }
        } else {
            this.showSetupToast('Save your changes before navigating to another tab', 'error', 'sticky')
        }
    }

    showNextStep() {
        const lastActiveStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
        this.activeStepIndex++;
        const newActiveStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
        if (lastActiveStep && newActiveStep) {
            this.steps[lastActiveStep.dataset.index].isActive = false;
            this.steps[lastActiveStep.dataset.index].isComplete = true;
            this.steps[newActiveStep.dataset.index].isActive = true;
            lastActiveStep.classList.add('slds-hide');
            newActiveStep.classList.remove('slds-hide');
        }
    }

    continueSetup() {
        this.showSetupLanding = false;
    }

    showSetupToast(message, variant, mode) {
        this.template.querySelector('c-toast').show(message, variant, mode);
    }
    
    showToast(event) {
        event.stopPropagation();
        let toast = event.detail.toast;
        this.template.querySelector('c-toast').show(toast.message, toast.variant, toast.mode);
    }
}