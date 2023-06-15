package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Notebook", 1200.99)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.Price, 1200.99)
}

func TestProductNameIsRequired(t *testing.T) {
	_, err := NewProduct("", 1200.99)
	assert.Error(t, err, "name is required")
}

func TestProductWithInvalidPrice(t *testing.T) {
	_, err := NewProduct("Notebook", 0)
	assert.Error(t, err, "invalid price")

	_, err = NewProduct("Notebook", -1)
	assert.Error(t, err, "invalid price")

	_, err = NewProduct("Notebook", 0.0)
	assert.Error(t, err, "price is required")
}
