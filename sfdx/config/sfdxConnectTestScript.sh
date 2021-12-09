#!/bin/bash 
#runfrom inside the config directory of project
#run => source sfdxConnectTestScript.sh; connectAndDeployToOrg
function connectAndDeployToOrg {
  echo "Check to see if sfdx is installed in sfdx directory:"
  isSfdxInstalled=`sfdx --version`
  if  [[ $isSfdxInstalled == sfdx-cli* ]] ;
    then 
      echo "SFDX is installed"
      echo "Is this for a Sandbox[s] or Production[p] org?"
      read orgType
      while [[ "$orgType" != "s" ]] && [[ "$orgType" != "Sandbox" ]] && [[ "$orgType" != "p" ]] && [[ "$orgType" != "Production" ]]
          do
          echo "I'm sorry. I didn't understand your response."
          echo "Is this for a Sandbox[s] or Production[p] org?"
          read orgType
      done
      echo "Please enter an Alias for this org:"
      read orgAlias
      
      if [[ "$orgType" == "s" ]] || [[ "$orgType" == "Sandbox" ]]
      #MIKE: Below is an example of how we would connect using JWT
      #sfdx auth:jwt:grant --clientid {clientId}--jwtkeyfile {pathToServer.key} --username {userName} --instanceurl https://login.salesforce.com --setalias $orgAlias --setdefaultdevhubusername
          then
          echo "Authenticating Sandbox org..."
          #sfdx force:auth:web:login --setalias $orgAlias --instanceurl https://test.salesforce.com --setdefaultusername
      else
          echo "Authenticating Production org..."
          #sfdx force:auth:web:login --setalias $orgAlias --instanceurl https://login.salesforce.com --setdefaultusername
      fi
      echo "Deploying code to connected org"
      sfdx force:source:deploy -p ../force-app/main/default/ -u $orgAlias
      echo "Successfully deployed local code to org"
      echo "Running all local tests in connected org"
      #sfdx force:apex:test:run --testlevel RunLocalTests -u $orgAlias --resultformat human
      sfdx force:org:open -u $orgAlias 

  else
    echo "SFDX is not installed please follow this link to install"
    echo "https://developer.salesforce.com/docs/atlas.en-us.234.0.sfdx_setup.meta/sfdx_setup/sfdx_setup_install_cli.htm"
  fi
}
