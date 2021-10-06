package user

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"template/internal/constant"
	"template/internal/constant/errors"
	"template/internal/constant/model"
)

type UserPersistence interface {
	UserByID(ctx context.Context, param model.User) (*model.User, error)
	Users(ctx context.Context) ([]model.User, error)
	UpdateUser(ctx context.Context, user model.User) (*model.User, error)
	DeleteUser(ctx context.Context, param model.User) error
	StoreUser(ctx context.Context, user model.User) (*model.User, error)
	UserExists(ctx context.Context, param model.User) (bool, error)
	PhoneExists(ctx context.Context, phone string) (bool, error)
	EmailExists(ctx context.Context, email string) (bool, error)
	AddUserToCompany(ctx context.Context, parm model.CompanyUser) error
	RemoveUser(ctx context.Context, parm model.CompanyUser) error
	GetCompanyUsers(ctx context.Context, companyID uuid.UUID) ([]model.CompanyUser, error)
	GetCompanyUserByID(ctx context.Context, user_id uuid.UUID) (*model.CompanyUser, error)
	MigrateUser(ctx context.Context) error
}
type userPersistence struct {
	conn *gorm.DB
}

func UserInit(conn *gorm.DB) UserPersistence {
	return &userPersistence{
		conn: conn,
	}
}
func (r userPersistence) GetCompanyUserByID(ctx context.Context, user_id uuid.UUID) (*model.CompanyUser, error) {
	conn := r.conn.WithContext(ctx)
	cUser := &model.CompanyUser{}
	err := conn.Model(&model.CompanyUser{}).Where(model.CompanyUser{UserID: user_id}).First(&cUser).Error
	if err != nil {
		return nil, errors.ErrorUnableToFetch
	}
	return cUser, nil
}

func (r userPersistence) UserByID(ctx context.Context, param model.User) (*model.User, error) {
	conn := r.conn.WithContext(ctx)
	user := &model.User{}
	err := conn.Model(&model.User{}).Where(&param).First(&user).Error
	if err != nil {
		return nil, errors.ErrorUnableToFetch
	}
	return user, nil
}
func (r userPersistence) Users(ctx context.Context) ([]model.User, error) {
	conn := r.conn.WithContext(ctx)
	users := []model.User{}
	err := conn.Model(&model.User{}).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrRecordNotFound
		}
		return nil, errors.ErrorUnableToFetch
	}
	return users, nil
}

func (r userPersistence) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Where(&model.User{ID: user.ID}).Or(&model.User{Phone: user.Phone}).Or(&model.User{Email: user.Email}).Updates(&user).Error
	if err != nil {
		return nil, errors.ErrUnableToSave
	}
	return &user, nil
}

func (r userPersistence) DeleteUser(ctx context.Context, param model.User) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Where(model.User{ID: param.ID}).Delete(&param).Error
	if err != nil {
		return errors.ErrUnableToDelete
	}
	return nil
}

func (r userPersistence) StoreUser(ctx context.Context, user model.User) (*model.User, error) {
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return nil, errors.ErrUnableToSave
	}
	return &user, nil
}

func (r userPersistence) UserExists(ctx context.Context, param model.User) (bool, error) {
	var count int64 = 0
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Where(&model.User{ID: param.ID}).Count(&count).Error
	if err != nil {
		return false, errors.ErrRecordNotFound
	}
	return count > 0, nil
}

func (r userPersistence) PhoneExists(ctx context.Context, phone string) (bool, error) {
	var countCompany int64 = 0
	var count int64 = 0
	conn := r.conn.WithContext(ctx)
	err := conn.Transaction(func(tx *gorm.DB) error {
		err := conn.Model(&model.User{}).Where(&model.User{Phone: phone}).Count(&count).Error
		if err != nil {
			return errors.ErrRecordNotFound
		}
		err = conn.Model(&model.Company{}).Where(&model.Company{Phone: phone}).Count(&count).Error
		if err != nil {
			return errors.ErrRecordNotFound
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return count+countCompany > 0, nil
}

//EmailExists
func (r userPersistence) EmailExists(ctx context.Context, email string) (bool, error) {
	var count int64 = 0
	conn := r.conn.WithContext(ctx)
	err := conn.Model(&model.User{}).Where(&model.User{Email: email}).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

//AddUserToCompany add user to company
func (r userPersistence) AddUserToCompany(ctx context.Context, parm model.CompanyUser) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Transaction(func(tx *gorm.DB) error {
		if parm.Role == constant.CompanyAdmin {
			u := &model.CompanyUser{}
			_ = tx.Model(&model.CompanyUser{}).Where(&parm).Find(u).Error
			if u != nil {
				return errors.ErrConflictHappened
			}
		}

		err := tx.Model(&model.User{}).Where(&model.User{ID: parm.UserID}).Updates(model.User{RoleName: parm.Role, Status: "Active"}).Error
		if err != nil {
			return errors.ErrInvalidTransaction
		}

		err = tx.Model(model.CompanyUser{}).Create(parm).Error
		if err != nil {
			return errors.ErrUnableToSave
		}
		return nil
	})
	if err != nil {
		return errors.ErrInvalidTransaction
	}
	return nil
}
func (r userPersistence) GetCompanyUsers(ctx context.Context, companyID uuid.UUID) ([]model.CompanyUser, error) {
	conn := r.conn.WithContext(ctx)
	company_users := []model.CompanyUser{}
	err := conn.Model(&model.CompanyUser{}).Where(model.CompanyUser{CompanyID: companyID}).Find(&company_users).Error
fmt.Println("error persis get c-users ",err)
	if err != nil {
		return nil, errors.ErrRecordNotFound
	}
	return company_users, nil
}

func (r userPersistence) RemoveUser(ctx context.Context, parm model.CompanyUser) error {
	conn := r.conn.WithContext(ctx)
	err := conn.Transaction(func(tx *gorm.DB) error {
		err := conn.Model(model.CompanyUser{}).
			Where(model.CompanyUser{UserID: parm.UserID, CompanyID: parm.CompanyID}).
			Delete(model.CompanyUser{UserID: parm.UserID, CompanyID: parm.CompanyID}).Error
		if err != nil {
			return errors.ErrUnableToDelete
		}
		err = conn.Model(&model.User{}).Where(&model.User{ID: parm.UserID}).Updates(&model.User{RoleName: "", Status: "Pending"}).Error
		if err != nil {
			return errors.ErrUnableToDelete
		}
		return nil
	})
	if err != nil {
		return errors.ErrInvalidTransaction
	}
	return nil
}
func (r userPersistence) MigrateUser(ctx context.Context) error {
	db := r.conn.WithContext(ctx)
	err := db.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	return nil
}
