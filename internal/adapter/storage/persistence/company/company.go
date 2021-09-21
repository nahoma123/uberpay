package company

import (
	"template/internal/constant/model"

	uuid "github.com/satori/go.uuid"
)

type CompanyStorage interface {
	Companies() ([]model.Company, error)
	CreateCompany(company *model.Company) (*model.Company, error)
	GetCompanyById(id uuid.UUID) (*model.Company, error)
	DeleteUser(id uuid.UUID) error
}
