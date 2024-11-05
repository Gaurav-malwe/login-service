package service

import (
	"context"

	"github.com/Gaurav-malwe/login-service/internal/model"
	log "github.com/Gaurav-malwe/login-service/utils/logging"
)

type IAuthService interface {
	RegisterUser(ctx context.Context, userRequest *model.RegisterUserRequest) error
	LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (string, error)
	ConfirmUser(ctx context.Context, loginRequest *model.ConfirmRequest) error
}

func (s *service) RegisterUser(ctx context.Context, userRequest *model.RegisterUserRequest) error {

	log := log.GetLogger(ctx)
	log.Debug("Service::RegisterUser")

	output := model.ToUserDetails(userRequest)

	err := s.CognitoRegisterUser(userRequest.Username, userRequest.Password)
	if err != nil {
		log.WithContext(ctx).WithFields(map[string]interface{}{
			"error": err,
		}).Error("Service::RegisterUser::Error while registering user")
		return err
	}

	err = s.repo.InsertUser(ctx, output)
	if err != nil {
		log.WithContext(ctx).WithFields(map[string]interface{}{
			"error": err,
		}).Error("Service::RegisterUser::Error while inserting user")
		return err
	}

	return nil

}

func (s *service) LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (string, error) {
	log := log.GetLogger(ctx)
	log.Debug("Service::LoginUser")

	token, err := s.AuthenticateUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		log.WithContext(ctx).WithFields(map[string]interface{}{
			"error": err,
		}).Error("Service::LoginUser::Error while logging in user")
		return "", err
	}

	return token, nil
}

func (s *service) ConfirmUser(ctx context.Context, loginRequest *model.ConfirmRequest) error {
	log := log.GetLogger(context.Background())
	log.Debug("Service::ConfirmUser")

	err := s.CognitoConfirmUser(loginRequest.Username, loginRequest.ConfirmationCode)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"error": err,
		}).Error("Service::LoginUser::Error while logging in user")
		return err
	}

	return nil
}
