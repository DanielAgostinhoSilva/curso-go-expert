package repository

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/google/uuid"
)

type ProductAdapter interface {
	Save(product *model.Product) error
	FindAll(page, limite int, sort string) ([]model.Product, error)
	FindById(id uuid.UUID) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id uuid.UUID) error
}
