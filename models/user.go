package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        string `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}
