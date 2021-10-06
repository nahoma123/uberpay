package user

import (
	"context"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/user"
	"template/internal/constant"
	"template/internal/constant/model"
	"template/internal/module"
	"time"
)



//Service defines all necessary service for the domain User
type service struct {
	usrRepo        repository.UserRepository
	usrPersist     user.UserPersistence
	validate       *validator.Validate
	trans          ut.Translator
	contextTimeout time.Duration
}

//Initialize creates a new object with LoginUseCase type
func Initialize(usrRepo repository.UserRepository, usrPersist user.UserPersistence, validate *validator.Validate, trans ut.Translator, timeout time.Duration) module.UserUsecase {
	return &service{
		usrRepo:        usrRepo,
		usrPersist:     usrPersist,
		validate:       validate,
		trans:          trans,
		contextTimeout: timeout,
	}
}

func (s *service) GetCompanyUserByID(c context.Context, user_id uuid.UUID) (*model.CompanyUser, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.GetCompanyUserByID(ctx, user_id)
}
func (s *service) UserByID(c context.Context, param model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.UserByID(ctx, param)
}

func (s *service) Users(c context.Context) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.Users(ctx)
}

func (s *service) UpdateUser(c context.Context, user model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.UpdateUser(ctx, user)
}

func (s *service) DeleteUser(c context.Context, param model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.DeleteUser(ctx, param)
}

func (s *service) StoreUser(c context.Context, user model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(user, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	err := s.usrRepo.Encrypt(&user)
	if err != nil {
		return nil, err
	}
	return s.usrPersist.StoreUser(ctx, user)

}

func (s *service) UserExists(c context.Context, param model.User) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.UserExists(ctx, param)

}

func (s *service) PhoneExists(c context.Context, phone string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.PhoneExists(ctx, phone)
}

func (s *service) EmailExists(c context.Context, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.EmailExists(ctx, email)

}

func (s *service) AddUserToCompany(c context.Context, parm model.CompanyUser) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(parm, s.validate, s.trans)
	if errV != nil {
		return  errV
	}
	return s.usrPersist.AddUserToCompany(ctx, parm)

}

func (s *service) RemoveUser(c context.Context, parm model.CompanyUser) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.RemoveUser(ctx, parm)

}

func (s *service) GetCompanyUsers(c context.Context, companyID uuid.UUID) ([]model.CompanyUser, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.usrPersist.GetCompanyUsers(ctx, companyID)

}
