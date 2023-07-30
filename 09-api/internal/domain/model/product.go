package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrInvalidId       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrNameIsRequired
	}
	if price == 0.0 {
		return nil, ErrPriceIsRequired
	}
	if price < 0.0 {
		return nil, ErrInvalidPrice
	}
	return &Product{
		ID:        uuid.New(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}, nil
}

func (p *Product) Validate() {

}
