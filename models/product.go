package models

import "time"

type Product struct {
	ID        string    `json:"id" gorm:"primaryKey;uuid;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	SKU       string    `json:"sku" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
