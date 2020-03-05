#!/bin/bash
BASEDIR=$(dirname $0)

echo "== Fixing generated code"
sed -i '' 's/protected configuration: Configuration;/protected configuration: Configuration | undefined;/' $BASEDIR/orders/api.ts
sed -i '' 's/name: "RequiredError"/name = "RequiredError"/' $BASEDIR/orders/api.ts
sed -i '' 's/(\<any\>.*localVarRequestOptions/localVarRequestOptions/' $BASEDIR/orders/api.ts
sed -i '' 's/= \<any\> /= /' $BASEDIR/orders/api.ts
sed -i -e "/response\.status/a\\
if (options \&\& options.onResponseHeader) options.onResponseHeader(response);
" $BASEDIR/orders/api.ts

echo "== Removing garbage"
rm $BASEDIR/orders/.gitignore $BASEDIR/orders/git_push.sh $BASEDIR/orders/.swagger-codegen-ignore