package domain

import (
	"time"
	_ "github.com/jinzhu/gorm"
	"gorm.io/gorm"
	
)

type Order struct {
	gorm.Model
	Price     float64
	Start     time.Time
	End       time.Time
	Plant Plant `gorm:"ForeignKey:ID"`
	PlantID int
	Status    bool
}