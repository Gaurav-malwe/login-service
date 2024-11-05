package service

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type ICognitoService interface {
	CognitoRegisterUser(username, password string) error
	AuthenticateUser(username, password string) (string, error)
}

// RegisterUser registers a new user in Cognito User Pool
func (s *service) CognitoRegisterUser(username, password string) error {
	userAttributes := []*cognitoidentityprovider.AttributeType{
		{
			Name:  aws.String("email"),
			Value: aws.String(username), // Using username as email for this example
		},
	}

	input := &cognitoidentityprovider.SignUpInput{
		ClientId:       aws.String(os.Getenv("CLIENT_ID")),
		Username:       aws.String(username),
		Password:       aws.String(password),
		UserAttributes: userAttributes,
	}
	_, err := s.cp.SignUp(input)
	return err
}

// AuthenticateUser authenticates a user and returns a JWT token
func (s *service) AuthenticateUser(username, password string) (string, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		ClientId: aws.String(os.Getenv("CLIENT_ID")),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(username),
			"PASSWORD": aws.String(password),
		},
	}
	resp, err := s.cp.InitiateAuth(input)
	if err != nil {
		return "", err
	}

	return *resp.AuthenticationResult.IdToken, nil
}
