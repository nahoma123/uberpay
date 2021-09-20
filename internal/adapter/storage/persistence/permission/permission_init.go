package persistence

import (
	"template/internal/constant/model"
	"template/internal/adapter/storage/persistence"
	"gorm.io/gorm"
)


type permisionPersistence struct {
	conn *gorm.DB
}

func PermissionInit(conn *gorm.DB) storage.PermissionPersistence {
	return &permisionPersistence{
		conn: conn,
	}
}

func (r permisionPersistence) Persmision(id uint) (*model.CasbinRule, error) {
	conn := r.conn
	p := &model.CasbinRule{}

	err := conn.Model(&model.CasbinRule{}).Where(&model.CasbinRule{ID: id}).Find(p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r permisionPersistence) Persmisions() ([]model.CasbinRule, error) {
	conn := r.conn

	permisionsList := []model.CasbinRule{}

	err := conn.Model(&model.CasbinRule{}).Find(&permisionsList).Error
	if err != nil {
		return nil, err
	}
	return permisionsList, nil
}

func (r permisionPersistence) UpdatePersmision(casRule model.CasbinRule) (*model.CasbinRule, error) {

	conn := r.conn

	err := conn.Model(&model.CasbinRule{}).Where(&model.CasbinRule{ID: casRule.ID}).Updates(&casRule).Error
	if err != nil {
		return nil, err
	}
	return &casRule, nil
}

func (r permisionPersistence) DeletePersmision(id uint) error {
	conn := r.conn

	err := conn.Model(&model.CasbinRule{}).Where(
		&model.CasbinRule{ID: id}).
		Delete(&model.CasbinRule{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r permisionPersistence) StorePersmision(casRule model.CasbinRule) (*model.CasbinRule, error) {
	conn := r.conn

	err := conn.Model(&model.CasbinRule{}).Create(&casRule).Error
	if err != nil {
		return nil, err
	}
	return &casRule, nil
}


func (r permisionPersistence) MigratePersmision() error {
	db := r.conn

	err := db.Migrator().AutoMigrate(&model.CasbinRule{})
	if err != nil {
		return err
	}

	return nil
}
