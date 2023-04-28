package models

import "github.com/golang-jwt/jwt/v4"

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Time     string `json:"time"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RecoveryPasswordRequest struct {
	Email string `json:"email"`
}

type AuthTokenClaims struct {
	jwt.StandardClaims
	UserID int    `json:"id"`
	Status string `json:"status"`
}
