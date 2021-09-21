package company

import (
	"template/internal/adapter/storage/persistence/company"
	"template/internal/constant/model"

	appErr "template/internal/constant/errors"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

// Usecase interface contains function of business logic for domian Company
type Usecase interface {
	Companies() ([]model.Company, error)
	CreateCompany(company *model.Company) (*model.Company, error)
	GetCompanyById(id uuid.UUID) (*model.Company, error)
	DeleteUser(id uuid.UUID) error
}

//Service defines all neccessary service for the domain Company
type service struct {
	companyPersist company.CompanyStorage
	validate       *validator.Validate
	trans          ut.Translator
}

// creates a new object with UseCase type
func Initialize(companyPersist company.CompanyStorage, validate *validator.Validate, trans ut.Translator) Usecase {
	return &service{
		companyPersist,
		validate,
		trans,
	}
}
func (s *service) Companies() ([]model.Company, error) {

	companies, err := s.companyPersist.Companies()
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (s *service) CreateCompany(comp *model.Company) (*model.Company, error) {

	valErr := s.validate.Struct(comp)

	if valErr != nil {
		errs := valErr.(validator.ValidationErrors)
		valErr := errs.Translate(s.trans)
		return nil, appErr.NewValErrResponse(valErr)
	}

	return s.companyPersist.CreateCompany(comp)
}

func (s *service) GetCompanyById(id uuid.UUID) (*model.Company, error) {
	return s.companyPersist.GetCompanyById(id)
}

func (s *service) DeleteUser(id uuid.UUID) error {
	return nil
}
