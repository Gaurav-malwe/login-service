package controller

import (
	"github.com/Gaurav-malwe/login-service/config"
	"github.com/Gaurav-malwe/login-service/internal/service"
)

type (
	Controller interface {
		IAuthController
	}

	controller struct {
		s   service.Service
		cfg *config.Config
	}
)

func New(svc service.Service, cfg *config.Config) Controller {
	return &controller{
		s:   svc,
		cfg: cfg,
	}
}
