#!/bin/bash
set -e
BUILD_DIRECTORY="/build_output"
API_NAME="services-email"
HANDLER_ZIP_NAME="handler_email_post.zip"
LAMBDA_FUNCTION_NAME="services-email-email"
API_PATH="email"
# Must be all CAPS i.e. POST or GET
API_HTTP="POST"


cd $BUILD_DIRECTORY
echo ">> Building bootstrap"
make build-handler-email-post

echo ">> Changing the permissions on boostrap"
chmod +x "$BUILD_DIRECTORY/bin/handler_email_post/bootstrap"

echo ">> Zipping bootstrap"
make zip-handler-email-post

echo ">> Renaming handler zip file"
cp "$BUILD_DIRECTORY/$HANDLER_ZIP_NAME" "function.zip"

echo "🚀 Registering Lambda function"

awslocal lambda create-function \
  --function-name "$LAMBDA_FUNCTION_NAME" \
  --runtime provided.al2 \
  --handler bootstrap \
  --zip-file "fileb://$BUILD_DIRECTORY/function.zip" \
  --role arn:aws:iam::000000000000:role/lambda-role


awslocal apigateway create-rest-api --name "$API_NAME"

API_ID=$(awslocal apigateway get-rest-apis \
    --query "items[?name=='$API_NAME'].id" \
    --output text)

PARENT_ID=$(awslocal apigateway get-resources \
    --rest-api-id "$API_ID" \
    --query "items[0].id" \
    --output text)


awslocal apigateway create-resource \
  --rest-api-id "$API_ID" \
  --parent-id "$PARENT_ID" \
  --path-part "$API_PATH"


RESOURCE_ID=$(awslocal apigateway get-resources \
    --rest-api-id "$API_ID" \
    --query "items[?pathPart=='$API_PATH'].id" \
    --output text)

awslocal apigateway put-method \
  --rest-api-id "$API_ID" \
  --resource-id "$RESOURCE_ID" \
  --http-method "$API_HTTP" \
  --authorization-type "NONE"

awslocal apigateway put-integration \
  --rest-api-id "$API_ID" \
  --resource-id "$RESOURCE_ID" \
  --http-method "$API_HTTP" \
  --type AWS_PROXY \
  --integration-http-method POST \
  --uri "arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:000000000000:function:$LAMBDA_FUNCTION_NAME/invocations"

awslocal apigateway create-deployment \
  --rest-api-id "$API_ID" \
  --stage-name dev

echo "✅ Lambda and API Gateway ready"
echo "API_ID"
echo "$API_ID"
