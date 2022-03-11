## Generate JSON for the Data Mapper from Open API Spec

```shell
# from root
cd sfdx/bin/openapi-spec-generator/

npm install
# jq isn't required, but nice for formatting
node index.js | jq > ../../../public/openapi.json
```
