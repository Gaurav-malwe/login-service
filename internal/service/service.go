package service

import (
	"github.com/Gaurav-malwe/login-service/config"
	"github.com/Gaurav-malwe/login-service/internal/repository"
)

type (
	Service interface {
		IAuthService
	}

	service struct {
		repo   repository.Repository
		config *config.Config
	}
)

func New(repository repository.Repository, config *config.Config) Service {
	return &service{
		repo:   repository,
		config: config,
	}
}
