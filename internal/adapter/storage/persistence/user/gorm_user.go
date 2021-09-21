package user

import (
	"log"

	"template/internal/constant/errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"template/internal/constant/model"
)

type userGormRepo struct {
	conn *gorm.DB
}

func UserInit(db *gorm.DB) UserStorage {
	return &userGormRepo{conn: db}
}

func (repo userGormRepo) User(param model.User) (*model.User, error) {
	conn := repo.conn
	user := &model.User{}

	err := conn.Model(&model.User{}).Where(&param).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repo userGormRepo) CreateSystemUser(user *model.User) (*model.User, error) {
	conn := repo.conn

	err := conn.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (repo userGormRepo) UserCompanyRole(param model.UserCompanyRole) (*model.UserCompanyRole, error) {
	conn := repo.conn
	userRole := &model.UserCompanyRole{}

	err := conn.Model(&model.UserCompanyRole{}).Where(&param).First(userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, nil
}
func (repo userGormRepo) CreateUser(companyID uuid.UUID, usr *model.User) (*model.User, error) {
	tx := repo.conn.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Error
	if err != nil {
		log.Printf("Error in message transaction: %v", err)
		return nil, err
	}

	company := &model.Company{}
	err = tx.First(company, companyID).Error

	if err != nil {
		tx.Rollback()
		log.Printf("Error encountered %v", err)
		return nil, errors.ErrUnknown
	}

	role := &model.Role{}
	err = tx.Where("name = ?", usr.RoleName).First(role).Error

	if err != nil {
		tx.Rollback()
		log.Printf("This is the error returned %v", err)
		return nil, errors.ErrUnknown
	}
	parentRole := ""
	childRole := ""
	if usr.RoleName == "COMPANY-ADMIN" || usr.RoleName == "COMPANY-CLERK" {
		parentRole = "COMPANY-USER"
		childRole = usr.RoleName
	}
	usr.RoleName = parentRole
	err = tx.Create(&usr).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return nil, errors.ErrUnknown
	}

	ucr := model.UserCompanyRole{
		UserID:    usr.ID,
		CompanyID: companyID,
		RoleName:  childRole,
	}

	err = tx.Create(ucr).Error

	if err != nil {
		tx.Rollback()
		log.Printf("This is the error returned %v", err)
		return nil, errors.ErrUnknown
	}

	err = tx.Commit().Error
	if err != nil {
		log.Printf("Error when commiting to db: %v", err)
		return nil, err
	}

	return usr, nil
}

func (repo userGormRepo) DeleteUser(id uuid.UUID) error {
	err := repo.conn.Delete(&model.User{}, id).Error
	if err != nil {
		log.Println(err)
		return errors.ErrUnknown
	}
	return nil
}

func (repo userGormRepo) GetUserById(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	err := repo.conn.First(user, id).Error
	if err != nil {
		log.Println(err)
		return nil, errors.ErrUnknown
	}
	return user, nil
}

func (repo userGormRepo) GetUsers() ([]model.User, error) {
	users := []model.User{}

	err := repo.conn.Find(&users).Error

	if err != nil {
		log.Println(err)
		return nil, errors.ErrUnknown
	}
	return users, nil
}
