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
  --region us-east-1 \
  --endpoint-url $ENDPOINT_URL)

# Create Cognito User Pool Client
CLIENT_ID=$(aws cognito-idp create-user-pool-client \
  --user-pool-id $USER_POOL_ID \
  --client-name login-service-client \
  --generate-secret \
  --query 'UserPoolClient.ClientId' \
  --output text \
  --region us-east-1 \
  --endpoint-url $ENDPOINT_URL)


echo "User Pool ID: $USER_POOL_ID"
echo "Client ID: $CLIENT_ID"


# Exporting them as environment variables for the application
echo "export USER_POOL_ID=$USER_POOL_ID" >> /root/.bashrc
echo "export CLIENT_ID=$CLIENT_ID" >> /root/.bashrc
