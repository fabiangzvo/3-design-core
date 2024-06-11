package models

import (
	"time"

	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Stock       int64     `json:"stock"`
	CategoryID  int64     `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

// TableName name of table in database
func (Product) TableName() string {
	return "products"
}
