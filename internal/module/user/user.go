package user

import (
	"template/internal/adapter/repository"
	"template/internal/adapter/storage/persistence/user"

	appErr "template/internal/constant/errors"
	"template/internal/constant/model"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

// Usecase interface contains function of business logic for domian USer
type Usecase interface {
	CreateUser(companyID uuid.UUID, user *model.User) (*model.User, error)
	CreateSystemUser(user *model.User) (*model.User, error)
	GetUserById(id uuid.UUID) (*model.User, error)
	DeleteUser(id uuid.UUID) error
	GetUsers() ([]model.User, error)
}

//Service defines all neccessary service for the domain User
type service struct {
	usrRepo    repository.UserRepository
	usrPersist user.UserStorage
	validate   *validator.Validate
	trans      ut.Translator
}

// creates a new object with UseCase type
func Initialize(usrRepo repository.UserRepository, usrPersist user.UserStorage, validate *validator.Validate, trans ut.Translator) Usecase {
	return &service{
		usrRepo,
		usrPersist,
		validate,
		trans,
	}
}

func (s *service) CreateUser(companyID uuid.UUID, user *model.User) (*model.User, error) {

	valErr := s.validate.Struct(user)

	if valErr != nil {
		errs := valErr.(validator.ValidationErrors)
		valErr := errs.Translate(s.trans)
		return nil, appErr.NewValErrResponse(valErr)
	}

	err := s.usrRepo.Encrypt(user)

	if err != nil {
		return nil, err
	}

	usr, err := s.usrPersist.CreateUser(companyID, user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *service) CreateSystemUser(user *model.User) (*model.User, error) {

	valErr := s.validate.Struct(user)

	if valErr != nil {
		errs := valErr.(validator.ValidationErrors)
		valErr := errs.Translate(s.trans)
		return nil, appErr.NewValErrResponse(valErr)
	}

	err := s.usrRepo.Encrypt(user)

	if err != nil {
		return nil, err
	}

	usr, err := s.usrPersist.CreateSystemUser(user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
func (s *service) GetUserById(id uuid.UUID) (*model.User, error) {
	usr, err := s.usrPersist.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *service) DeleteUser(id uuid.UUID) error {
	return s.usrPersist.DeleteUser(id)
}

func (s *service) GetUsers() ([]model.User, error) {
	return s.usrPersist.GetUsers()
}
