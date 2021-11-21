#!/bin/sh
#set -x
 
echo "🐼 > Run this script from the package root and make sure the path is correct in sfdx-project.json file"
echo "🐼 "
echo "🐼 ------------------------------------------"
echo "🐼 "
read -n 1 -s -r -p "🐼 Press any key to continue "
echo "🐼 "
echo "🐼 "
echo ===========Output CLI version===========
sfdx --version
echo =============Output org list===========
sfdx force:org:list
echo "🐼 "
echo "🐼 "
echo "🐼 Time to enter a good name for your scratch org..."
echo "🐼 Example: vishal-test"
echo "🐼 "
while [ ! -n "$ORG_NAME"  ]
do
    echo "🐼  Please enter a name for your scratch org:"
    echo " "
    read ORG_NAME
    echo " "
done
 
# START OF CREATING ORG
 
echo "🐼 "
echo "🐼 Building your org, please wait..."
RES=$(sfdx force:org:create -f config/project-scratch-def.json -d 30 -a "${ORG_NAME}" --json)
 
if [ "$?" = "1" ]
then
  echo "🐼 "
    echo "🐼 ERROR: Can't create your org."
  echo "🐼 "
    read -n 1 -s -r -p "🐼 Press any key to continue"
    exit
fi
 
echo "🐼 "
echo "🐼 Scratch org created successfully..."
 
# START OF PUSHING SOURCE CODE
 
echo "🐼 "
echo "🐼 Pushing the code, please wait. It may take a while."
 
sfdx force:source:push -u ${ORG_NAME}
 
if [ "$?" = "1" ]
then
  echo "🐼 "
    echo "🐼 ERROR: Can't push your source."
  echo "🐼 "
    read -n 1 -s -r -p "🐼 Press any key to continue"
    exit
fi
 
echo "🐼 "
echo "🐼 Code is pushed successfully."
 
echo "🐼 "
echo "🐼 BUILD AUTOMATION Finished."
echo "🐼 "
 
read -n 1 -s -r -p "🐼 Press any key to continue"
echo " "
