import companyLogo from '@salesforce/resourceUrl/companyLogo';
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
    @track steps = [
        {
            title: 'Connect Stripe and Salesforce',
            name: 'C-SYSTEM-CONNECTIONS-STEP',
            orderIndex: 1,
            isComplete: false,
            isActive: false
        },
        {
            title: 'Define Data Mapping',
            name: 'C-DATA-MAPPING-STEP',
            orderIndex: 2,
            isComplete: false,
            isActive: false
        },
        {
            title: 'Configure Sync Preferences',
            name: 'C-SYNC-PREFERENCES-STEP',
            orderIndex: 3,
            isComplete: false,
            isActive: false
        }
    ];

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

    completeSave() {
        this.contentLoading = false;
    }

    getmappingconfigurations() {
        this.template.querySelector('c-data-mapping-step').getPicklistValuesForMapper(true, '');
        this.template.querySelector('c-sync-preferences-step').connectedCallback();  
        this.nextDisabled = false;
    }

    async connectedCallback() {
        try {
            this.steps[this.activeStepIndex].isActive = true;
            const setupData = await getSetupData();
            this.data = JSON.parse(setupData);

            if (this.data.isSuccess && this.data.results.isConnected) {
                this.nextDisabled = false;
                this.setupComplete = this.data.results.setupData.isSetupComplete__c;
                let completedSteps = JSON.parse(this.data.results.setupData.Steps_Completed__c);
                
                if (Object.keys(completedSteps).length > 0) {
                    this.setupStarted = true;
                    this.steps[this.activeStepIndex].isComplete = true;
                    this.steps[this.activeStepIndex].isActive = false;

                    for (const step in completedSteps) {
                        this.steps[this.activeStepIndex].isComplete = true;

                        if (this.activeStepIndex < this.steps.length - 1) {
                            this.showNextStep();
                        }
                    }
                }

                const setOrganizationType = await setOrgType();
                let orgTyperResp =  JSON.parse(setOrganizationType);
                   
                if (!orgTyperResp.isSuccess) {
                    this.showSetupToast(orgTyperResp.error, 'error', 'sticky');
                }
            } else if (this.data.error) { 
                this.showSetupToast(this.data.error, 'error', 'sticky');
            }
        } catch (error) {
            this.showSetupToast(error.message, 'error', 'sticky');
        } finally {
            this.showContent();
            this.loading = false;
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
        if(this.activeSectionIndex == 1) {
            this.template.querySelector('c-data-mapping-step').saveCongfiguredMappings();
        } else if(this.activeSectionIndex == 2) {
            this.template.querySelector('c-sync-preferences-step').saveModifiedSyncPreferences();
        }
        this.contentLoading = true;
        this.saveDisabled = true;
    }

    async next(event) { 
        this.contentLoading = true;
        if(this.activeStepIndex <= (this.steps.length - 1)) {
            let stepName = this.steps[this.activeStepIndex].name; 
            let stepNumber = this.activeStepIndex + 1;
            let isLastStep = false;
            if (stepNumber === (this.steps.length)) {
                 isLastStep = true
                }
            try {
                const saveSetupData = await saveData({
                    newSetupDataRec: {
                        Steps_Completed__c: JSON.stringify({
                            [stepName]: stepNumber
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
            } catch (error) {
                this.showSetupToast(error.message, 'error', 'sticky');
            } finally {
                switch (stepName) {
                case 'C-DATA-MAPPING-STEP':
                    this.template.querySelector('c-data-mapping-step').saveCongfiguredMappings();
                    break; 
                case 'C-SYNC-PREFERENCES-STEP':
                    this.template.querySelector('c-sync-preferences-step').saveModifiedSyncPreferences();
                    this.steps[this.activeStepIndex].isActive = false;
                    this.steps[this.activeStepIndex].isComplete = true;
                    this.nextDisabled = true;     
                    this.setupComplete = true;
                    this.showSetupToast('Setup completed', 'success')
                    break;
                default:
                    break;
                }
                if(this.activeStepIndex < (this.steps.length - 1)) {
                    this.showNextStep();
                }
                this.contentLoading = false;
            }
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