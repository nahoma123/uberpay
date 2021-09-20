package auth

import (
	"fmt"
	"log"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant/model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UseCase interface {
	Login(username, password string)  (*model.LoginResponse, error)
	GetClaims(token string) (*model.UserClaims, error)
}

type service struct {
	userPersistence    user.UserStorage
	jwtManager         JWTManager
}

func Initialize(
	userPersistence    user.UserStorage,
	jwtManager         JWTManager,
) UseCase {
	return &service{
		userPersistence:    userPersistence,
		jwtManager:         jwtManager,
	}
}
func (s service) GetClaims(token string) (*model.UserClaims, error) {
	claims, err := s.jwtManager.Verify(token)
	if err != nil {
		return nil, err
	}
	return claims, err
}

func (s service) Login(phoneNumber, password string) (*model.LoginResponse, error) {
	// if phoneNumber == "" || password == "" {

	// 	return nil,
	// }

	usr, err := s.userPersistence.User(model.User{Phone: phoneNumber})
	if err != nil {
		log.Println("53--", err)
		return nil, err

	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	claims := &model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			Subject: usr.ID.String(),
		},
		Phone: phoneNumber,
		Role:  usr.RoleName,
	}
	if usr.RoleName == "COMPANY-USER" {
		// companyID, _ := s.userPersistence.GetUserCompany(usr.ID)
		companyRole, err := s.userPersistence.UserCompanyRole(model.UserCompanyRole{UserID: usr.ID})

		if err != nil {
			return nil,err
		}
		claims.CompanyID = companyRole.CompanyID.String()
	}

	token, err := s.jwtManager.Generate(claims)
	if err != nil {
		return nil, err
	}
	return  &model.LoginResponse{
			Token:           token,
			Name:            fmt.Sprintf("%v %v %v", usr.FirstName, usr.MiddleName, usr.LastName),
			Email:           usr.Email,
			Phone:           usr.Phone,
			Role:            usr.RoleName,} ,nil
}