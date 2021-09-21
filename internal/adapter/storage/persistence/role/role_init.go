package role

import (
	"gorm.io/gorm"
	"template/internal/adapter/storage/persistence"
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

func (r rolePersistence) Role(name string) (*model.Role, error) {
	conn := r.conn
	role := &model.Role{}

	err := conn.Model(&model.Role{}).Where(&model.Role{Name: name}).Find(role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r rolePersistence) Roles() ([]model.Role, error) {
	conn := r.conn

	roles := []model.Role{}

	err := conn.Model(&model.Role{}).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r rolePersistence) UpdateRole(role model.Role) (*model.Role, error) {

	conn := r.conn

	err := conn.Model(&model.Role{}).Where(&model.Role{Name: role.Name}).Updates(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r rolePersistence) DeleteRole(name string) error {
	conn := r.conn

	err := conn.Model(&model.Role{}).Where(&model.Role{Name: name}).Delete(&model.Role{Name: name}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r rolePersistence) StoreRole(role model.Role) (*model.Role, error) {
	conn := r.conn

	err := conn.Model(&model.Role{}).Create(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r rolePersistence) RoleExists(name string) (bool, error) {
	conn := r.conn
	var count int64

	err := conn.Model(&model.Role{}).Where(&model.Role{Name: name}).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
func (r rolePersistence) MigrateRole() error {
	db := r.conn

	err := db.Migrator().AutoMigrate(&model.Role{})
	if err != nil {
		return err
	}

	return nil
}
