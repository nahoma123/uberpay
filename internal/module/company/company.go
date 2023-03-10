package company

import (
	"context"
	"fmt"
	company "ride_plus/internal/adapter/storage/persistence"
	"ride_plus/internal/constant"
	"ride_plus/internal/constant/errors"
	custErr "ride_plus/internal/constant/errors"
	"ride_plus/internal/module"
	"time"

	model "ride_plus/internal/constant/model/dbmodel"
	utils "ride_plus/internal/constant/model/init"

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
func Initialize(companyPersist company.CompanyPersistence, utils utils.Utils) module.CompanyUsecase {
	return &service{
		companyPersist: companyPersist,
		validate:       utils.GoValidator,
		trans:          utils.Translator,
		contextTimeout: utils.Timeout,
	}
}
func (s *service) StoreCompanyImage(c context.Context, param model.CompanyImage) (*model.CompanyImage, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.StructValidator(param, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	isExist, err := s.companyPersist.ImageExists(model.Image{Hash: param.Image.Hash})
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, custErr.ErrFieldAlreadyExist
	}
	image, err := s.companyPersist.StoreCompanyImage(ctx, param)
	if err != nil {
		return nil, err
	}
	return image, nil

}
func (s *service) UpdateCompanyImage(c context.Context, param model.CompanyImage) (*model.CompanyImage, *errors.ErrorModel) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.VerifyInput(param, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	isExist, err := s.companyPersist.ImageExists(model.Image{Hash: param.Image.Hash})
	if err != nil {
		return nil, err
	}
	fmt.Println("isExist ", isExist)
	if isExist {
		return s.companyPersist.UpdateCompanyImage(ctx, param)
	}
	return nil, custErr.ErrRecordNotFound

}
func (s *service) CompanyImages(c context.Context) ([]model.CompanyImage, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.companyPersist.CompanyImages(ctx)
}

func (s *service) CompanyByID(c context.Context, param model.Company) (*model.Company, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.companyPersist.CompanyByID(ctx, param)
}

func (s *service) StoreCompany(c context.Context, param model.Company) (*model.Company, *errors.ErrorModel) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.VerifyInput(param, s.validate, s.trans)
	if errV != nil {
		return nil, errV
	}
	return s.companyPersist.StoreCompany(ctx, param)

}

func (s *service) UpdateCompany(c context.Context, param model.Company) (*model.Company, *errors.ErrorModel) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	errV := constant.VerifyInput(param, s.validate, s.trans)
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
