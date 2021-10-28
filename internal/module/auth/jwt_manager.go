package auth

import (
	"fmt"
	"ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"

	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{secretKey}
}

func (manager *JWTManager) Generate(userClaims *model.UserClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*model.UserClaims, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("empty token")
	}
	token, err := jwt.ParseWithClaims(
		accessToken,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.ErrInvalidToken
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, errors.ErrInvalidToken
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.ErrInvalidAccessToken
	}

	return claims, nil
}
