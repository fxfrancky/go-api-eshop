package repository

import (
	"errors"
	"strings"

	"github.com/fxfrancky/go-api-eshop/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

// AllUsers implements UserRepository
func (u *UserRepositoryImpl) AllUsers(offset int, limit int) ([]models.User, int64, error) {
	var (
		users []models.User
		count int64
	)
	u.DB.Model(&users).Count(&count)
	u.DB.Offset(offset).Limit(limit).Find(&users)
	return users, count, nil
}

// DeleteUser implements UserRepository
func (u *UserRepositoryImpl) DeleteUser(user *models.User) error {
	return u.DB.Delete(user).Error
}

// GetUserByEmail implements UserRepository
func (u *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.DB.First(&user, "email = ?", strings.ToLower(email)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else if err != nil {
			return nil, err
		}

	}

	return &user, err
}

// GetUserById implements UserRepository
func (u *UserRepositoryImpl) GetUserById(id int) (*models.User, error) {
	var user models.User
	err := u.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, err

}

// RegisterUser implements UserRepository
func (u *UserRepositoryImpl) RegisterUser(user models.User) error {
	result := u.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateUser implements UserRepository
func (u *UserRepositoryImpl) UpdateUser(user *models.User) error {
	result := u.DB.Model(&user).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewUserRepositoryImpl(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}
