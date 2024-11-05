package model

import (
	"time"

	"github.com/Gaurav-malwe/login-service/utils"
)

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Mobile   string `json:"mobile"`
	RoleId   string `json:"role_id"`
	Admin    bool   `json:"admin"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ToUserDetails(req *RegisterUserRequest) *User {
	return &User{
		Email:     req.Email,
		Password:  req.Password,
		Username:  req.Username,
		Fullname:  req.Fullname,
		Mobile:    req.Mobile,
		RoleID:    req.RoleId,
		Admin:     req.Admin,
		CreatedAt: utils.GetCurrentTimeForDB(),
		UpdatedAt: time.Time{},
		Active:    true,
		DeletedAt: time.Time{},
	}
}
