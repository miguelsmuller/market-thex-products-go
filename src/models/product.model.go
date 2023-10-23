package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var validate = validator.New()

type Product struct {
	ID          string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Description string    `gorm:"not null" json:"description,omitempty"`
	Price       float32   `gorm:"not null" json:"price"`
	Image       string    `json:"image,omitempty"`
	Category    string    `gorm:"varchar(100)" json:"category,omitempty"`
	Published   bool      `gorm:"default:false;not null" json:"published"`
	CreatedAt   time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
	UpdatedAt   time.Time `gorm:"not null;default:'1970-01-01 00:00:01';ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt,omitempty"`
}

type CreateProductSchema struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float32   `json:"price"`
	Image       string    `json:"image,omitempty"`
	Category    string    `json:"category,omitempty"`
	Published  bool      `json:"published"`
}

type UpdateProductSchema struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float32   `json:"price"`
	Image       string    `json:"image,omitempty"`
	Category    string    `json:"category,omitempty"`
	Published   bool      `json:"published"`
}

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}


func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New().String()
	return nil
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

