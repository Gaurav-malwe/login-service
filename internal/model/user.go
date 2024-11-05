package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    string    `json:"user_id" bson:"_id,omitempty"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	Username  string    `json:"username" bson:"username"`
	Fullname  string    `json:"fullname" bson:"fullname"`
	Mobile    string    `json:"mobile" bson:"mobile"`
	RoleID    string    `json:"role_id" bson:"role_id"`
	Admin     bool      `json:"admin" bson:"admin"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at"`
	Active    bool      `json:"active" bson:"active"`
}

type Claims struct {
	UserId   string `json:"user_id"`
	Admin    bool   `json:"admin"`
	Email    string `json:"email"`
	RoleId   string `json:"role_id"`
	Mobile   string `json:"mobile"`
	Fullname string `json:"fullname"`
	jwt.RegisteredClaims
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
