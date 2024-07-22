package models

import "github.com/golang-jwt/jwt"

type UserJWT struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}
