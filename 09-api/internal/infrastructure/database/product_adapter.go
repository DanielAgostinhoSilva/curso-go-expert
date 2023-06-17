package database

import (
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductAdapter struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *ProductAdapter {
	return &ProductAdapter{DB: db}
}

func (adapter *ProductAdapter) Save(product *model.Product) error {
	return adapter.DB.Create(product).Error
}

func (adapter *ProductAdapter) FindById(id uuid.UUID) (*model.Product, error) {
	var product model.Product
	err := adapter.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (adapter *ProductAdapter) Update(product *model.Product) error {
	_, err := adapter.FindById(product.ID)
	if err != nil {
		return err
	}
	return adapter.DB.Save(product).Error
}

func (adapter *ProductAdapter) Delete(id uuid.UUID) error {
	product, err := adapter.FindById(id)
	if err != nil {
		return err
	}
	return adapter.DB.Delete(product).Error
}
