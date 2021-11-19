declare module "@salesforce/apex/setupAssistant.getPackageVersion" {
  export default function getPackageVersion(): Promise<any>;
}
declare module "@salesforce/apex/setupAssistant.saveData" {
  export default function saveData(param: {setupData: any}): Promise<any>;
}
declare module "@salesforce/apex/setupAssistant.getSetupData" {
  export default function getSetupData(): Promise<any>;
}
declare module "@salesforce/apex/setupAssistant.isCpqEnabled" {
  export default function isCpqEnabled(): Promise<any>;
}
declare module "@salesforce/apex/setupAssistant.validateConnectionStatus" {
  export default function validateConnectionStatus(param: {isConnectedCallback: any}): Promise<any>;
}
declare module "@salesforce/apex/setupAssistant.getMappingConfigurations" {
  export default function getMappingConfigurations(): Promise<any>;
}
declare module "@salesforce/apex/setupAssistant.saveMappingConfigurations" {
  export default function saveMappingConfigurations(param: {jsonMapObj: any}): Promise<any>;
}
declare module "@salesforce/apex/setupAssistant.manualRetry" {
  export default function manualRetry(param: {jsonMapObj: any}): Promise<any>;
}
