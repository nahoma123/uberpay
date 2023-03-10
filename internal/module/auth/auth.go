package auth

import (
	"context"
	"fmt"
	"ride_plus/internal/adapter/storage/persistence/user"
	"ride_plus/internal/constant/errors"
	"ride_plus/internal/module"
	"time"

	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type authservice struct {
	userPersistence user.UserPersistence
	jwtManager      JWTManager
	contextTimeout  time.Duration
}

func Initialize(userPersistence user.UserPersistence, jwtManager JWTManager, utils utils.Utils) module.LoginUseCase {
	return &authservice{
		userPersistence: userPersistence,
		jwtManager:      jwtManager,
		contextTimeout:  utils.Timeout,
	}
}
func (s authservice) GetClaims(token string) (*model.UserClaims, error) {
	claims, err := s.jwtManager.Verify(token)
	if err != nil {
		return nil, errors.ErrInvalidToken
	}
	return claims, nil
}

func (s authservice) Login(c context.Context, phoneNumber, password string) (*model.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	u := model.User{Phone: phoneNumber}
	usr, err := s.userPersistence.UserByID(ctx, u)
	u.Password = password
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		return nil, errors.ErrInvalidUserPhoneOrPassword
	}
	if usr.RoleName == "" {
		return nil, errors.ErrRequireApproval
	}
	claims := &model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			Subject: usr.ID.String(),
		},
		Phone: phoneNumber,
		Role:  usr.RoleName,
	}
	token, err := s.jwtManager.Generate(claims)
	if err != nil {
		return nil, errors.ErrGenerateToken
	}
	return &model.LoginResponse{
		Token: token,
		Name:  fmt.Sprintf("%v %v %v", usr.FirstName, usr.MiddleName, usr.LastName),
		Email: usr.Email,
		Phone: usr.Phone,
		Role:  usr.RoleName}, nil
}
