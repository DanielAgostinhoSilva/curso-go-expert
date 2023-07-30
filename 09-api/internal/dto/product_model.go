package dto

import (
	"github.com/google/uuid"
	"time"
)

type ProductModel struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
