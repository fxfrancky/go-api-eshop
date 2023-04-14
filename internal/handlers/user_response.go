package handlers

import (
	"time"

	"github.com/fxfrancky/go-api-eshop/internal/models"
)

// User Response
type userResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type userListResponse struct {
	Users      []*userResponse `json:"users"`
	UsersCount int64           `json:"usersCount"`
}

func newUserResponse(u *models.User) *userResponse {
	usr := new(userResponse)
	usr.ID = u.ID
	usr.Name = u.Name
	usr.Email = u.Email
	usr.Photo = u.Photo
	usr.Role = u.Role
	usr.Provider = u.Provider
	usr.IsAdmin = u.IsAdmin
	usr.CreatedAt = u.CreatedAt
	usr.UpdatedAt = u.UpdatedAt
	return usr
}

func newUserListResponse(users []models.User, count int64) *userListResponse {
	u := new(userListResponse)
	u.Users = make([]*userResponse, 0)
	for _, us := range users {
		ur := newUserResponse(&us)
		u.Users = append(u.Users, ur)
	}
	u.UsersCount = count
	return u
}
