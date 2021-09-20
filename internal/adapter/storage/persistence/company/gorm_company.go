package company

import (
	"log"
	"template/internal/constant/errors"
	"template/internal/constant/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type companyGormRepo struct {
	conn *gorm.DB
}

func CompanyInit(db *gorm.DB) CompanyStorage {
	return &companyGormRepo{conn: db}
}

func (repo companyGormRepo) CreateCompany(company *model.Company) (*model.Company, error) {

	err := repo.conn.Create(company).Error

	if err != nil {
		log.Printf("Errror when saving  company to db %v", err)
		return nil, errors.ErrUnknown
	}
	return company, nil
}

func (repo companyGormRepo) GetCompanyById(id uuid.UUID) (*model.Company, error) {
	company := &model.Company{}
	err := repo.conn.First(company, id).Error
	if err != nil {
		log.Println(err)
		return nil, errors.ErrUnknown
	}
	return company, nil

}
func (repo companyGormRepo) Companies() ([]model.Company, error) {
	conn := repo.conn

	companies := []model.Company{}

	err := conn.Model(&model.Company{}).Find(&companies).Error
	if err != nil {
		return nil, err
	}
	return companies, nil

}
func (repo companyGormRepo) DeleteUser(id uuid.UUID) error {
	return nil
}
