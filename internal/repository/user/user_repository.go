package repository

import "github.com/fxfrancky/go-api-eshop/internal/models"

type UserRepository interface {
	// User Repository
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	RegisterUser(user models.User) error
	UpdateUser(user *models.User) error
	AllUsers(offset, limit int) ([]models.User, int64, error)
	DeleteUser(user *models.User) error
}
