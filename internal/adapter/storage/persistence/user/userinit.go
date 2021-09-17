package persistence

import (
	"template/internal/adapter/storage"
	"template/internal/constant/errorconstant"
	"template/internal/constant/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type userGormRepo struct {
	conn *gorm.DB
}

func InitUserPersistence(db *gorm.DB) storage.UserStorage {
	return &userGormRepo{conn: db}
}

func (repo userGormRepo) CreateUser(usr *model.User) (*model.User, error) {
	err := repo.conn.Create(&usr).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorconstant.IDNotFound
		}

		return nil, errorconstant.ErrorUnableToSave
	}
	return usr, nil
}

func (repo userGormRepo) DeleteUser(id uuid.UUID) error {
	err := repo.conn.Delete(&model.User{}, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errorconstant.IDNotFound
		}
		return errorconstant.ErrorUnableToDelete
	}
	return nil
}

func (repo userGormRepo) GetUserById(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	err := repo.conn.First(user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorconstant.IDNotFound
		}
		return nil, errorconstant.ErrorUnableToFetch
	}
	return user, nil
}

func (repo userGormRepo) GetUsers() ([]model.User, error) {
	return nil, nil
}
