package persistence

import (
	"context"
	storage "ride_plus/internal/adapter/storage/persistence"
	"ride_plus/internal/constant/errors"
	"ride_plus/internal/constant/model"

	"gorm.io/gorm"
)

type policyPersistence struct {
	conn *gorm.DB
}

func PolicyInit(conn *gorm.DB) storage.PermissionPersistence {
	return &policyPersistence{
		conn: conn,
	}
}

func (r policyPersistence) Policy(ctx context.Context, id uint) (*model.CasbinRule, error) {
	conn := r.conn.WithContext(ctx)
	p := &model.CasbinRule{}
	err := conn.Model(&model.CasbinRule{}).Where(&model.CasbinRule{ID: id}).Find(p).Error
	if err != nil {
		return nil, errors.ErrorUnableToFetch
	}
	return p, nil
}

func (r policyPersistence) Policies(ctx context.Context) ([]model.CasbinRule, error) {
	conn := r.conn.WithContext(ctx)
	permisionsList := []model.CasbinRule{}
	err := conn.Model(&model.CasbinRule{}).Find(&permisionsList).Error
	if err != nil {
		return nil, errors.ErrorUnableToFetch
	}
	return permisionsList, nil
}

func (r policyPersistence) UpdatePolicy(ctx context.Context, casRule model.CasbinRule) (*model.CasbinRule, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.CasbinRule{}).Where(&model.CasbinRule{ID: casRule.ID}).Updates(&casRule).Error
	if err != nil {
		return nil, errors.ErrUnableToSave
	}
	return &casRule, nil
}

func (r policyPersistence) RemovePolicy(ctx context.Context, id uint) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.CasbinRule{}).Where(
		&model.CasbinRule{ID: id}).
		Delete(&model.CasbinRule{ID: id}).Error
	if err != nil {
		return errors.ErrIDNotFound
	}
	return nil
}

func (r policyPersistence) StorePolicy(ctx context.Context, casRule model.CasbinRule) (*model.CasbinRule, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.CasbinRule{}).Create(&casRule).Error
	if err != nil {
		return nil, errors.ErrUnableToSave
	}
	return &casRule, nil
}
func (r policyPersistence) MigratePolicy(ctx context.Context) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Migrator().AutoMigrate(&model.CasbinRule{})
	if err != nil {
		return errors.ErrUnableToMigrate
	}
	return nil
}
