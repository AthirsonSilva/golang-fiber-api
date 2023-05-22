package models

import "time"

type Order struct {
	ID        string    `json:"id" gorm:"primaryKey;uuid;not null"`
	ProductID string    `json:"product_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	UserID    string    `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
