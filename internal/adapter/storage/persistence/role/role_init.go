package role

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"template/internal/adapter/storage/persistence"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

type rolePersistence struct {
	conn *gorm.DB
}

func RoleInit(conn *gorm.DB) storage.RolePersistence {
	return &rolePersistence{
		conn: conn,
	}
}

func (r rolePersistence) Role(ctx context.Context, name string) (*model.Role, error) {
	conn := r.conn
	role := &model.Role{}
	err := conn.Model(&model.Role{}).Where(&model.Role{Name: name}).Find(role).Error
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return role, nil
}

func (r rolePersistence) Roles(ctx context.Context) ([]model.Role, error) {
	conn := r.conn
	roles := []model.Role{}
	err := conn.Model(&model.Role{}).Find(&roles).Error
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return roles, nil
}

func (r rolePersistence) UpdateRole(ctx context.Context, role model.Role) (*model.Role, error) {

	conn := r.conn

	err := conn.Model(&model.Role{}).Where(&model.Role{Name: role.Name}).Updates(&role).Error
	if err != nil {
		return nil, errors.ErrorUnableToFetch
	}
	return &role, nil
}

func (r rolePersistence) DeleteRole(ctx context.Context, name string) error {
	conn := r.conn
	err := conn.Model(&model.Role{}).Where(&model.Role{Name: name}).Delete(&model.Role{Name: name}).Error
	if err != nil {
		return errors.ErrUnableToDelete
	}
	return nil
}

func (r rolePersistence) StoreRole(ctx context.Context, role model.Role) (*model.Role, error) {
	conn := r.conn
	err := conn.Model(&model.Role{}).Create(&role).Error
	fmt.Println("error ", err)
	if err != nil {
		return nil, errors.ErrUnableToSave
	}
	return &role, nil
}

func (r rolePersistence) RoleExists(ctx context.Context, name string) (bool, error) {
	conn := r.conn
	var count int64
	err := conn.Model(&model.Role{}).Where(&model.Role{Name: name}).Count(&count).Error
	if err != nil {
		return false, errors.ErrRecordNotFound
	}

	return count > 0, nil
}
func (r rolePersistence) MigrateRole() error {
	db := r.conn
	err := db.Migrator().AutoMigrate(&model.Role{})
	if err != nil {
		return errors.ErrUnableToMigrate
	}

	return nil
}
