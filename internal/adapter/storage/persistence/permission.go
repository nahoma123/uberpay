package storage

import "template/internal/constant/model"

type PermissionPersistence interface {
	Persmision(id uint) (*model.CasbinRule, error)
	Persmisions() ([]model.CasbinRule, error)
	UpdatePersmision(role model.CasbinRule) (*model.CasbinRule, error)
	DeletePersmision(id uint) error
	StorePersmision(role model.CasbinRule) (*model.CasbinRule, error)
	MigratePersmision() error
}