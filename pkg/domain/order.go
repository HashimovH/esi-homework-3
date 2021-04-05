package domain

import (
	"time"
)

type Order struct {
	ID        int
	Ident     string
	Name      string
	Price     float64
	Status    bool
	Start     string
	End       string
	CreatedAt time.Time
}