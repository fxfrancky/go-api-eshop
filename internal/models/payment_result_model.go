package models

import (
	"gorm.io/gorm"
)

type PaymentResult struct {
	PaymentResID string
	Status       string
	UpdateTime   string
	EmailAddress string
	// OrderID      uint
	gorm.Model
}

// type CreatePaymentResultRequest struct {
// 	PaymentResultID string `json:"payment_result_id" validate:"required"`
// 	Status          string `json:"status" validate:"required"`
// 	UpdateTime      string `json:"update_time" validate:"required"`
// 	EmailAddress    string `json:"email_address" validate:"required"`
// }

// type PaymentResultResponse struct {
// 	ID              uint      `json:"id,omitempty"`
// 	PaymentResultID string    `json:"payment_result_id,omitempty"`
// 	Status          string    `json:"status,omitempty"`
// 	UpdateTime      string    `json:"update_time,omitempty"`
// 	EmailAddress    string    `json:"email_address,omitempty"`
// 	CreatedAt       time.Time `json:"created_at"`
// 	UpdatedAt       time.Time `json:"updated_at"`
// 	DeletedAt       time.Time `json:"deleted_at"`
// }
