package cognitoClient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient interface {
	NewCognitoClient() (*cognitoidentityprovider.CognitoIdentityProvider, error)
}

func NewCognitoClient() (*cognitoidentityprovider.CognitoIdentityProvider, error) {
	// Set up AWS session with custom configuration for LocalStack
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:4566"), // LocalStack endpoint
	})
	if err != nil {
		return nil, err
	}

	// Create the CognitoIdentityProvider client
	cognitoClient := cognitoidentityprovider.New(sess)

	return cognitoClient, nil
}
