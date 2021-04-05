package domain

import (
	"time"
)

type Plantm struct {
	ID        interface{} `bson:"_id,omitempty" json:"ID"`
	Ident     string
	Name      string
	Price     float64
	Status    bool
	CreatedAt time.Time
}