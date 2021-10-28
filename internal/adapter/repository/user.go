package repository

import (
	"ride_plus/internal/constant/errors"
	model "ride_plus/internal/constant/model/dbmodel"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository contains the functions of data logic for domain user
type UserRepository interface {
	Encrypt(user *model.User) (err error)
	CheckPassword(password string, user *model.User) (bool, error)
}

type userRepository struct {
}

// UserInit initializes the data logic / repository for domain user
func UserInit() UserRepository {
	return &userRepository{}
}

// Encrypt encrypts a users password
func (ur *userRepository) Encrypt(user *model.User) (err error) {
	if user.Password != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
		if err != nil {
			return errors.ErrPasswordEncryption
		}
		user.Password = string(bytes)
	}

	return nil
}

// CheckPassword checks if the given password is natch with the hash saved in the database
func (ur *userRepository) CheckPassword(password string, user *model.User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil, errors.ErrInvalidUserPhoneOrPassword
}
