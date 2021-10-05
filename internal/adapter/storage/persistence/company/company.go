package company

import (
	"context"
	"errors"
	"gorm.io/gorm"
	storage "template/internal/adapter/storage/persistence"
	custErr "template/internal/constant/errors"
	"template/internal/constant/model"
)

type companyPersistence struct {
	conn *gorm.DB
}

func CompanyInit(conn *gorm.DB) storage.CompanyPersistence {
	return &companyPersistence{
		conn: conn,
	}
}

func (r companyPersistence) CompanyByID(ctx context.Context, param model.Company) (*model.Company, error) {
	conn := r.conn.WithContext(ctx)
	company := &model.Company{}
	err := conn.Model(&model.Company{}).Where(&param).First(company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, custErr.ErrorUnableToFetch
	}
	return company, nil
}

func (r companyPersistence) Companies(ctx context.Context) ([]model.Company, error) {
	conn := r.conn.WithContext(ctx)
	companies := []model.Company{}
	err := conn.Model(&model.Company{}).Find(&companies).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, custErr.ErrorUnableToFetch
	}
	return companies, err
}

func (r companyPersistence) StoreCompany(ctx context.Context, company model.Company) (*model.Company, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Company{}).Create(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidTransaction) {
			return nil, gorm.ErrInvalidTransaction
		}
		return nil, custErr.ErrUnableToSave
	}
	return &company, nil
}

func (r companyPersistence) UpdateCompany(ctx context.Context, company model.Company) (*model.Company, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Company{}).Where(&model.Company{ID: company.ID}).Updates(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidTransaction) {
			return nil, gorm.ErrInvalidTransaction
		}
		return nil, custErr.ErrUnableToSave
	}
	return &company, nil
}

func (r companyPersistence) DeleteCompany(ctx context.Context, param model.Company) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.Company{}).Where(&param).Delete(&param).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return custErr.ErrorUnableToFetch
	}
	return nil
}

func (r companyPersistence) CompanyExists(ctx context.Context, param model.Company) (bool, error) {
	conn := r.conn.WithContext(ctx)
	var count int64
	err := conn.Model(&model.Company{}).Where(&param).Count(&count).Error
	if err != nil {
		return false, custErr.ErrRecordNotFound
	}
	return count > 0, nil
}
func (r companyPersistence) MigrateCompany(ctx context.Context) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Migrator().AutoMigrate(&model.Company{})
	if err != nil {
		return err
	}

	err = conn.Migrator().AutoMigrate(&model.CompanyUser{})
	if err != nil {
		return err
	}
	return nil
}
