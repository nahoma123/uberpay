package permission

import (
	"template/internal/adapter/storage/persistence"
	"template/internal/constant/model"
)

type UseCase interface {
	Persmision(id uint) (*model.CasbinRule, error)
	Persmisions() ([]model.CasbinRule, error)
	UpdatePersmision(role model.CasbinRule) (*model.CasbinRule, error)
	DeletePersmision(id uint) error
	StorePersmision(role model.CasbinRule) (*model.CasbinRule, error)
}
type service struct {
	permPersistence storage.PermissionPersistence
}

func PermissioInitialize(
	permPersistence storage.PermissionPersistence,
) UseCase {
	return &service{
		permPersistence: permPersistence,
	}
}

func (s service) Persmisions() ([]model.CasbinRule, error) {
	roles, err := s.permPersistence.Persmisions()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s service) Persmision(id uint) (*model.CasbinRule, error) {
	p, err := s.permPersistence.Persmision(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s service) DeletePersmision(id uint) error {

	err := s.permPersistence.DeletePersmision(id)
	if err != nil {

		return err
	}
	return nil
}
func (s service) UpdatePersmision(p model.CasbinRule) (*model.CasbinRule, error) {
	pr, err := s.permPersistence.UpdatePersmision(p)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func (s service) StorePersmision(p model.CasbinRule) (*model.CasbinRule, error) {
	pr, err := s.permPersistence.StorePersmision(p)
	if err != nil {
		return nil, err
	}
	return pr, nil
}
