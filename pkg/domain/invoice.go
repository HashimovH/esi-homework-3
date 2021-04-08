package domain

import (
	_ "github.com/jinzhu/gorm"
)

type Invoice struct {
	ID        int `gorm:"primary_key;type:int;"`
	Order []Order `gorm:"ForeignKey:ID"`
	Status    string
}