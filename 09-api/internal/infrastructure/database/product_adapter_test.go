package database

import (
	"fmt"
	"github.com/DanielAgostinhoSilva/curso-go-expert/09-api/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestProductAdapter_Save(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Product{})
	product, err := model.NewProduct("Product 1", 10.0)
	if err != nil {
		t.Error(err)
	}

	adapter := NewProductAdapter(db)
	err = adapter.Save(product)
	assert.Nil(t, err)

	var productFound model.Product
	err = db.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, product.ID, productFound.ID)
}

func TestFinalAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Product{})
	adapter := NewProductAdapter(db)
	for i := 1; i < 24; i++ {
		product, err := model.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		err = adapter.Save(product)
		assert.NoError(t, err)
	}
	products, err := adapter.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = adapter.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = adapter.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestProductAdapter_FindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Product{})
	product, err := model.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	adapter := NewProductAdapter(db)
	err = adapter.Save(product)
	assert.NoError(t, err)
	productFound, err := adapter.FindById(product.ID)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestProductAdapter_Update(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&model.Product{})
	adapter := NewProductAdapter(db)

	product, err := model.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	err = adapter.Save(product)
	assert.NoError(t, err)

	product.Name = "Product 2"
	product.Price = 20.0
	err = adapter.Update(product)
	assert.NoError(t, err)

	productFound, err := adapter.FindById(product.ID)
	assert.NoError(t, err)
	assert.Equal(t, productFound.Name, product.Name)
	assert.Equal(t, productFound.Price, product.Price)
}

func TestProductAdapter_Delete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&model.Product{})
	adapter := NewProductAdapter(db)

	product, err := model.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	err = adapter.Save(product)
	assert.NoError(t, err)

	err = adapter.Delete(product.ID)
	assert.NoError(t, err)

	_, err = adapter.FindById(product.ID)
	assert.Error(t, err)
}
