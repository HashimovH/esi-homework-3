package domain

import (
	// "time"
	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model
	// ID        int `json:"id"`
	// Ident     string `json:"ident"`
	Name      string `json:"name"`
	Price     float64 `json:"price"`
	Status    bool `json:"status"`
	// CreatedAt time.Time `json:"createdAt"`
}