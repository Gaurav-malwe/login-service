#!/bin/bash

# Set up AWS CLI to use LocalStack endpoint
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=us-east-1
export ENDPOINT_URL=http://localhost:4566

# Create Cognito User Pool
USER_POOL_ID=$(aws cognito-idp create-user-pool \
  --pool-name login-service-user-pool \
  --query 'UserPool.Id' \
  --output text \
  --endpoint-url $ENDPOINT_URL)

# Create Cognito User Pool Client
aws cognito-idp create-user-pool-client \
  --user-pool-id $USER_POOL_ID \
  --client-name login-service-client \
  --generate-secret \
  --query 'UserPoolClient.ClientId' \
  --output text \
  --endpoint-url $ENDPOINT_URL
