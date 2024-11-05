package controller

import (
	"errors"
	"net/http"

	"github.com/Gaurav-malwe/login-service/internal/model"
	log "github.com/Gaurav-malwe/login-service/utils/logging"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
)

type IAuthController interface {
	Register(ginCtx *gin.Context)
	Login(ginCtx *gin.Context)
	Confirm(ginCtx *gin.Context)
}

func (c *controller) Register(ginCtx *gin.Context) {

	span, ctx := opentracing.StartSpanFromContext(ginCtx.Request.Context(), "Controller::Auth::Register")
	defer span.Finish()

	log := log.GetLogger(ctx)
	log.Info("Controller::Auth::Register")

	payload, err := validateRegisterUserRequest(ginCtx)
	if checkError(ginCtx, err) {
		return
	}

	err = c.s.RegisterUser(ctx, &payload)
	if checkError(ginCtx, err) {
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

}

func (c *controller) Login(ginCtx *gin.Context) {

	span, ctx := opentracing.StartSpanFromContext(ginCtx.Request.Context(), "Controller::Auth::Login")
	defer span.Finish()

	log := log.GetLogger(ctx)
	log.Info("Controller::Auth::Login")

	payload, err := validateLoginRequest(ginCtx)
	if checkError(ginCtx, err) {
		return
	}

	token, err := c.s.LoginUser(ctx, &payload)
	if checkError(ginCtx, err) {
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"message": "User signed in successfully", "auth_token": token})

}
func (c *controller) Confirm(ginCtx *gin.Context) {

	span, ctx := opentracing.StartSpanFromContext(ginCtx.Request.Context(), "Controller::Auth::Login")
	defer span.Finish()

	log := log.GetLogger(ctx)
	log.Info("Controller::Auth::Login")

	payload, err := validateConfirmRequest(ginCtx)
	if checkError(ginCtx, err) {
		return
	}

	err = c.s.ConfirmUser(ctx, &payload)
	if checkError(ginCtx, err) {
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{"message": "User confirmed successfully"})

}

func validateRegisterUserRequest(ginCtx *gin.Context) (model.RegisterUserRequest, error) {
	var payload model.RegisterUserRequest
	var err error

	// check binding
	if err := ginCtx.ShouldBind(&payload); err != nil {
		return payload, err
	}

	validate := validator.New()

	err = validate.Struct(payload)
	if err != nil {
		arr := listErrors(err)
		return payload, errors.New("Invalid/missing input parameters: " + arr)
	}
	return payload, nil
}

func validateLoginRequest(ginCtx *gin.Context) (model.LoginRequest, error) {
	var payload model.LoginRequest
	var err error

	// check binding
	if err := ginCtx.ShouldBind(&payload); err != nil {
		return payload, err
	}

	validate := validator.New()

	err = validate.Struct(payload)
	if err != nil {
		arr := listErrors(err)
		return payload, errors.New("Invalid/missing input parameters: " + arr)
	}
	return payload, nil
}

func validateConfirmRequest(ginCtx *gin.Context) (model.ConfirmRequest, error) {
	var payload model.ConfirmRequest
	var err error

	// check binding
	if err := ginCtx.ShouldBind(&payload); err != nil {
		return payload, err
	}

	validate := validator.New()

	err = validate.Struct(payload)
	if err != nil {
		arr := listErrors(err)
		return payload, errors.New("Invalid/missing input parameters: " + arr)
	}
	return payload, nil
}
