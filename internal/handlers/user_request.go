package handlers

import (
	"github.com/fxfrancky/go-api-eshop/internal/models"
	"github.com/gofiber/fiber/v2"
)

// User Request
type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
	IsAdmin         bool   `json:"isAdmin"`
	Photo           string `json:"photo,omitempty"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

// User Request
type updateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Role     string `json:"role,omitempty"`
	Photo    string `json:"photo,omitempty"`
	Provider string `json:"provider"`
	IsAdmin  bool   `json:"isAdmin"`
	Password string `json:"password" validate:"required,min=8"`
}

func (u *updateUserRequest) bindUpdateUser(c *fiber.Ctx, usr *models.User) error {

	// Validate the user
	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Map the product
	usr.Name = u.Name
	usr.Email = u.Email
	usr.IsAdmin = u.IsAdmin
	usr.Password = u.Password
	usr.Provider = u.Provider
	usr.Role = u.Role
	usr.Photo = u.Photo

	return nil
}

func (s *updateUserRequest) populateUpdateUser(user *models.User) {
	s.Name = user.Name
	s.Email = user.Email
	s.IsAdmin = user.IsAdmin
	s.Password = user.Password
	s.Photo = user.Photo
	s.Role = user.Role
}
