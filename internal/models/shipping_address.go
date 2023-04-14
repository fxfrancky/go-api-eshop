package models

import (
	"gorm.io/gorm"
)

type ShippingAddress struct {
	Address    string `gorm:"not null"`
	City       string `gorm:"not null"`
	PostalCode string `gorm:"not null"`
	Country    string `gorm:"not null"`
	// OrderID    uint
	gorm.Model
}

// type CreateShippingAddressRequest struct {
// 	Address    string `json:"address" validate:"required"`
// 	City       string `json:"city" validate:"required"`
// 	PostalCode string `json:"postal_code" validate:"required"`
// 	Country    string `json:"country" validate:"required"`
// 	OrderID    uint   `json:"order_id" validate:"required"`
// }
// type UpdateShippingAddressRequest struct {
// 	ID         int     `json:"id" validate:"required"`
// 	Address    string  `json:"address,omitempty"`
// 	City       float64 `json:"city,omitempty"`
// 	PostalCode string  `json:"postal_code,omitempty"`
// 	Country    string  `json:"country,omitempty"`
// 	OrderID    uint    `json:"order_id" validate:"required"`
// }

// type ShippingAddressResponse struct {
// 	ID         uint      `json:"id,omitempty"`
// 	Address    string    `json:"address,omitempty"`
// 	City       string    `json:"city,omitempty"`
// 	PostalCode string    `json:"postal_code,omitempty"`
// 	Country    string    `json:"country,omitempty"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// 	DeletedAt  time.Time `json:"deleted_at"`
// }
