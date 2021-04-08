package domain

import (
	"time"
	_ "github.com/jinzhu/gorm"
)

type Order struct {
	ID        int `gorm:"primary_key;type:int;"`
	Ident     string
	Name      string
	Price     float64
	Start     string
	End       string
	CreatedAt time.Time
	Plant []Plant `gorm:"ForeignKey:ID"`
	Status    bool
	Delivered bool
}