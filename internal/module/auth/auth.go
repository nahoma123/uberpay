package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant/errors"
	"template/internal/constant/model"
	"template/internal/module"
	"time"
)


type service struct {
	userPersistence user.UserPersistence
	jwtManager      JWTManager
	contextTimeout  time.Duration
}

func Initialize(userPersistence user.UserPersistence, jwtManager JWTManager, timeout time.Duration) module.LoginUseCase {
	return &service{
		userPersistence: userPersistence,
		jwtManager:      jwtManager,
		contextTimeout:  timeout,
	}
}
func (s service) GetClaims(token string) (*model.UserClaims, error) {
	claims, err := s.jwtManager.Verify(token)
	if err != nil {
		return nil, errors.ErrInvalidToken
	}
	return claims, nil
}

func (s service) Login(c context.Context, phoneNumber, password string) (*model.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	usr, err := s.userPersistence.UserByID(ctx, model.User{Phone: phoneNumber})
	if err != nil {
		return nil, err
	}
	//err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if password != usr.Password {
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
	companyUser, err := s.userPersistence.GetCompanyUserByID(ctx, usr.ID)
	if err != nil {
		return nil, err
	}
	claims.CompanyID = companyUser.CompanyID.String()
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
