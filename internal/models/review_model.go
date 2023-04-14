package models

import (
	"gorm.io/gorm"
)

type Review struct {
	Name      string  `gorm:"not null"`
	Rating    float64 `gorm:"not null"`
	Comment   string  `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	UserID    uint
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	gorm.Model
}

// type CreateReviewRequest struct {
// 	Name      string  `json:"name" validate:"required"`
// 	Rating    float64 `json:"rating" validate:"required"`
// 	Comment   string  `json:"comment" validate:"required"`
// 	UserID    int     `json:"user_id" validate:"required"`
// 	ProductID int     `json:"product_id" validate:"required"`
// }
// type UpdateReviewRequest struct {
// 	ID        int     `json:"id" validate:"required"`
// 	Name      string  `json:"name" validate:"required"`
// 	Rating    float64 `json:"rating" validate:"required"`
// 	Comment   string  `json:"comment" validate:"required"`
// 	UserID    int     `json:"user_id" validate:"required"`
// 	ProductID int     `json:"product_id" validate:"required"`
// }

// type ReviewResponse struct {
// 	ID        uint      `json:"id,omitempty"`
// 	Name      string    `json:"name,omitempty"`
// 	Rating    float64   `json:"rating"`
// 	Comment   string    `json:"comment,omitempty"`
// 	UserID    int       `json:"user_id,omitempty"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	DeletedAt time.Time `json:"deleted_at"`
// }
