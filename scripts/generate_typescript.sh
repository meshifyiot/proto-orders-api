#!/bin/bash
echo "== Delete old generated files"
rm -rf ./typescript

echo "== Generate client files from api.swagger.json"
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
    -i /local/swagger/api.swagger.json \
    -l typescript-fetch \
    -o /local/typescript

echo "== Fixing generated code"
sed -i '' 's/protected configuration: Configuration;/protected configuration: Configuration | undefined;/' typescript/api.ts
sed -i '' 's/name: "RequiredError"/name = "RequiredError"/' typescript/api.ts
sed -i '' 's/(\<any\>.*localVarRequestOptions/localVarRequestOptions/' typescript/api.ts
sed -i '' 's/= \<any\> /= /' typescript/api.ts
sed -i -e "/response\.status/a\\
if (options \&\& options.onResponseHeader) options.onResponseHeader(response);
" typescript/api.ts

echo "== Removing garbage"
rm ./typescript/.gitignore ./typescript/git_push.sh ./typescript/.swagger-codegen-ignore