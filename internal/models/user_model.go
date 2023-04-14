package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex:idx_email;not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Role     string `gorm:"type:varchar(50);default:'user';not null"`
	Provider string `gorm:"type:varchar(50);default:'local';not null"`
	Photo    string `gorm:"not null;default:'default.png'"`
	Verified bool   `gorm:"not null;default:false"`
	IsAdmin  bool   `gorm:"not null;default:false"`
	Orders   []Order
	gorm.Model
}

// User Response
type UserResponse struct {
	ID uint `json:"id,omitempty"`
	// ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FilterUserRecord(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Photo:     user.Photo,
		Provider:  user.Provider,
		IsAdmin:   user.IsAdmin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
