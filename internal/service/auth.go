package service

import (
	"context"
	"errors"
	"time"

	"github.com/Gaurav-malwe/login-service/internal/model"
	log "github.com/Gaurav-malwe/login-service/utils/logging"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthService interface {
	RegisterUser(ctx context.Context, userRequest *model.RegisterUserRequest) (string, error)
	LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (string, error)
	ValidateToken(ctx context.Context, signedToken string) (*model.Claims, error)
}

func (s *service) RegisterUser(ctx context.Context, userRequest *model.RegisterUserRequest) (string, error) {

	log := log.GetLogger(ctx)
	log.Debug("Service::RegisterUser")

	output := model.ToUserDetails(userRequest)

	_, err := s.repo.GetUserByEmail(ctx, output.Email)
	if err != nil && err == mongo.ErrNoDocuments {
		err := output.SetPassword(output.Password)
		if err != nil {
			log.WithContext(ctx).WithFields(map[string]interface{}{
				"error": err,
			}).Error("Service::RegisterUser::Error while setting password")
			return "", err
		}

		err = s.repo.InsertUser(ctx, output)
		if err != nil {
			log.WithContext(ctx).WithFields(map[string]interface{}{
				"error": err,
			}).Error("Service::RegisterUser::Error while inserting user")
			return "", err
		}

		token, err := s.generateJWT(output)
		if err != nil {
			log.WithContext(ctx).WithFields(map[string]interface{}{
				"error": err,
			}).Error("Service::RegisterUser::Error while generating token")
			return "", err
		}

		log.WithContext(ctx).Info("Service::RegisterUser::User registered successfully")
		return token, nil
	}

	log.WithContext(ctx).Info("Service::RegisterUser::User already exists")
	return "", errors.New("user already exists")

}

func (s *service) LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (string, error) {
	log := log.GetLogger(ctx)
	log.Debug("Service::LoginUser")

	user, err := s.repo.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		log.WithContext(ctx).Info("Service::LoginUser::Invalid credentials")
		return "", err
	}

	err = user.CheckPassword(loginRequest.Password)
	if err != nil {
		log.WithContext(ctx).Info("Service::LoginUser::Invalid Username/Password")
		return "", errors.New("invalid username/password")
	}

	token, err := s.generateJWT(user)
	if err != nil {
		log.WithContext(ctx).Info("Service::LoginUser::Error while generating token")
		return "", err
	}

	log.WithContext(ctx).Info("Service::LoginUser::User logged in successfully")
	return token, nil
}

func (s *service) ValidateToken(ctx context.Context, signedToken string) (*model.Claims, error) {

	log := log.GetLogger(ctx)
	log.Debug("Service::ValidateToken")

	jwtKey := s.config.GetString("JWT_KEY")

	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		log.WithContext(ctx).Info("Service::ValidateToken::Error while parsing token")
		return nil, err
	}

	if !token.Valid {
		log.WithContext(ctx).Info("Service::ValidateToken::Invalid token")
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (s *service) generateJWT(user *model.User) (string, error) {
	jwtKey := s.config.GetString("JWT_KEY")

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.Claims{
		UserId:   user.UserId,
		Admin:    user.Admin,
		Email:    user.Email,
		RoleId:   user.RoleID,
		Mobile:   user.Mobile,
		Fullname: user.Fullname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))
}
