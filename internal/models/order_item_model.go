package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	Name      string  `gorm:"not null"`
	Quantity  int64   `gorm:"not null"`
	Image     string  `gorm:"not null"`
	Price     float64 `gorm:"not null"`
	Product   Product
	ProductID uint
	OrderID   uint
	gorm.Model
}

// type CreateOrderItemRequest struct {
// 	Name      string  `json:"name" validate:"required"`
// 	Quantity  int64   `json:"quantity" validate:"required"`
// 	Price     float64 `json:"price" validate:"required"`
// 	ProductID uint    `json:"product_id" validate:"required"`
// 	OrderID   uint    `json:"order_id" validate:"required"`
// }

// type UpdateOrderItemRequest struct {
// 	ID        uint    `json:"id" validate:"required"`
// 	Name      string  `json:"name" validate:"required"`
// 	Quantity  int64   `json:"quantity" validate:"required"`
// 	Image     string  `json:"image" validate:"required"`
// 	Price     float64 `json:"price" validate:"required"`
// 	ProductID uint    `json:"product_id" validate:"required"`
// 	OrderID   uint    `json:"order_id" validate:"required"`
// }

// type OrderItemResponse struct {
// 	ID        uint      `json:"id,omitempty"`
// 	Name      string    `json:"name,omitempty"`
// 	Quantity  int64     `json:"quantity,omitempty"`
// 	Image     string    `json:"image,omitempty"`
// 	Price     float64   `json:"price,omitempty"`
// 	ProductID uint      `json:"product_id,omitempty"`
// 	OrderID   uint      `json:"order_id,omitempty"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	DeletedAt time.Time `json:"deleted_at"`
// }
