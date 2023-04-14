package models

import (
	"gorm.io/gorm"
)

type Product struct {
	// User         User `gorm:"foreignKey:UserID"`
	// UserID       uint
	Name         string `gorm:"uniqueIndex;not null"`
	Image        string `gorm:"not null"`
	Brand        string `gorm:"not null"`
	Category     string `gorm:"not null"`
	Description  string `gorm:"not null"`
	Reviews      []Review
	Rating       float64 `gorm:"not null;default:0"`
	NumReviews   int64   `gorm:"not null;default:0"`
	Price        float64 `gorm:"not null;default:0"`
	CountInStock int64   `gorm:"not null;default:0"`
	gorm.Model
}
