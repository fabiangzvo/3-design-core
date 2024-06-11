package models

import (
	"time"

	productModel "3-design-core/api/products/models"
)

//Category struct
type Category struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrementIncrement;"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
	Products    []productModel.Product `json:"products" gorm:"foreignKey:CategoryID"`
}

//TableName name of table in database
func (Category) TableName() string {
	return "categories"
}