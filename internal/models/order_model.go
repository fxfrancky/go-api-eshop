package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	PaymentMethod   string  `gorm:"not null"`
	TaxPrice        float64 `gorm:"not null;default:0.0"`
	ShippingPrice   float64 `gorm:"not null;default:0.0"`
	TotalPrice      float64 `gorm:"not null;default:0.0"`
	IsPaid          bool    `gorm:"not null;default:false"`
	PaidAt          time.Time
	IsDelivered     bool `gorm:"not null;default:false"`
	DeliveredAt     time.Time
	PaymentResult   PaymentResult
	PaymentResultID uint
	// OrderItems        []OrderItem `gorm:"many2many:order_orderitems;"`
	OrderItems        []OrderItem `gorm:"ForeignKey:OrderID"`
	ShippingAddress   ShippingAddress
	ShippingAddressID uint
	User              User
	UserID            uint
	gorm.Model
}

// type Order struct {
// 	User              User `gorm:"foreignKey:UserID"`
// 	UserID            uint
// 	OrderItems        []OrderItem     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// 	ShippingAddress   ShippingAddress `gorm:"foreignKey:ShippingAddressID"`
// 	ShippingAddressID uint
// 	PaymentMethod     string        `gorm:"not null"`
// 	PaymentResult     PaymentResult `gorm:"foreignKey:PaymentResultID"`
// 	PaymentResultID   uint
// 	TaxPrice          float64 `gorm:"not null;default:0.0"`
// 	ShippingPrice     float64 `gorm:"not null;default:0.0"`
// 	TotalPrice        float64 `gorm:"not null;default:0.0"`
// 	IsPaid            bool    `gorm:"not null;default:false"`
// 	PaidAt            time.Time
// 	IsDelivered       bool `gorm:"not null;default:false"`
// 	DeliveredAt       time.Time
// 	gorm.Model
// }

// type CreateOrderRequest struct {
// 	ID                int         `json:"id" validate:"required"`
// 	UserID            int         `json:"user_id" validate:"required"`
// 	OrderItems        []OrderItem `json:"orderItemList"`
// 	ShippingAddressID int         `json:"shipping_address_id" validate:"required"`
// 	PaymentMethod     string      `json:"payment_method" validate:"required"`
// 	PaymentResult     int         `json:"payment_result_id" validate:"required"`
// 	TaxPrice          float64     `json:"tax_price" validate:"required"`
// 	ShippingPrice     float64     `json:"shipping_price" validate:"required"`
// 	TotalPrice        float64     `json:"total_price" validate:"required"`
// 	IsPaid            bool        `json:"is_paid" validate:"required"`
// 	PaidAt            time.Time   `json:"paid_at" validate:"required"`
// 	IsDelivered       bool        `json:"is_delivered" validate:"required"`
// 	DeliveredAt       time.Time   `json:"delivered_at" validate:"required"`
// }

// type UpdateOrderRequest struct {
// 	ID                int         `json:"id" validate:"required"`
// 	UserID            int         `json:"user_id" validate:"required"`
// 	OrderItems        []OrderItem `json:"orderItemList"`
// 	ShippingAddressID int         `json:"shipping_address_id" validate:"required"`
// 	PaymentMethod     string      `json:"payment_method" validate:"required"`
// 	PaymentResult     int         `json:"payment_result_id" validate:"required"`
// 	TaxPrice          float64     `json:"tax_price" validate:"required"`
// 	ShippingPrice     float64     `json:"shipping_price" validate:"required"`
// 	TotalPrice        float64     `json:"total_price" validate:"required"`
// 	IsPaid            bool        `json:"is_paid" validate:"required"`
// 	PaidAt            time.Time   `json:"paid_at" validate:"required"`
// 	IsDelivered       bool        `json:"is_delivered" validate:"required"`
// 	DeliveredAt       time.Time   `json:"delivered_at" validate:"required"`
// }

// type OrderResponse struct {
// 	ID                uint        `json:"id,omitempty"`
// 	UserID            int         `json:"user_id,omitempty"`
// 	OrderItems        []OrderItem `json:"orderItemList,omitempty"`
// 	ShippingAddressID int         `json:"shipping_address_id,omitempty"`
// 	PaymentMethod     string      `json:"payment_method,omitempty"`
// 	PaymentResult     int         `json:"payment_result_id,omitempty"`
// 	TaxPrice          float64     `json:"tax_price,omitempty"`
// 	ShippingPrice     float64     `json:"shipping_price,omitempty"`
// 	TotalPrice        float64     `json:"total_price,omitempty"`
// 	IsPaid            bool        `json:"is_paid,omitempty"`
// 	PaidAt            time.Time   `json:"paid_at,omitempty"`
// 	IsDelivered       bool        `json:"is_delivered,omitempty"`
// 	DeliveredAt       time.Time   `json:"delivered_at"`
// 	CreatedAt         time.Time   `json:"created_at"`
// 	UpdatedAt         time.Time   `json:"updated_at"`
// 	DeletedAt         time.Time   `json:"deleted_at"`
// }
