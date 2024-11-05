package service

import (
	"github.com/Gaurav-malwe/login-service/config"
	"github.com/Gaurav-malwe/login-service/internal/repository"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type (
	Service interface {
		IAuthService
		ICognitoService
	}

	service struct {
		repo   repository.Repository
		config *config.Config
		cp     *cognitoidentityprovider.CognitoIdentityProvider
	}
)

func New(repository repository.Repository, cprovider *cognitoidentityprovider.CognitoIdentityProvider, config *config.Config) Service {
	return &service{
		repo:   repository,
		config: config,
		cp:     cprovider,
	}
}
