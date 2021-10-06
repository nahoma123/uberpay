package company

import (
	"context"
	company "template/internal/adapter/storage/persistence"
	"template/internal/constant"
	"template/internal/constant/model"
	"template/internal/module"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)


//Service defines all necessary service for the domain Company
type service struct {
	companyPersist company.CompanyPersistence
	validate       *validator.Validate
	trans          ut.Translator
	contextTimeout time.Duration
}

//Initialize  creates a new object with UseCase type
func Initialize(companyPersist company.CompanyPersistence, validate *validator.Validate, trans ut.Translator, timeout time.Duration) module.CompanyUsecase {
	return &service{
		companyPersist: companyPersist,
		validate:       validate,
		trans:          trans,
		contextTimeout: timeout,
	}
}

func (s *service) CompanyByID(c context.Context, param model.Company) (*model.Company, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.companyPersist.CompanyByID(ctx, param)
}

func (s *service) StoreCompany(c context.Context, param model.Company) (*model.Company, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(param, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	return s.companyPersist.StoreCompany(ctx, param)

}

func (s *service) UpdateCompany(c context.Context, param model.Company) (*model.Company, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(param, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	return s.companyPersist.UpdateCompany(ctx, param)
}

func (s *service) DeleteCompany(c context.Context, param model.Company) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.companyPersist.DeleteCompany(ctx, param)
}

func (s *service) CompanyExists(c context.Context, param model.Company) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.companyPersist.CompanyExists(ctx, param)
}

func (s *service) Companies(c context.Context) ([]model.Company, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.companyPersist.Companies(ctx)
}
