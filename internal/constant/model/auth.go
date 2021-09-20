package model

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	CompanyID string `json:"company_id,omitempty"`
}
type LoginResponse struct {
	Token           string `json:"token"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Role            string `json:"role"`
}