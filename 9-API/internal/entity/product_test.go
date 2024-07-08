package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const productName = "Product 1"

func TestNewProduct(t *testing.T) {
	expectedProductName := productName
	expectedProductPrice := 100

	p, err := NewProduct(expectedProductName, expectedProductPrice)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, p.Name, expectedProductName)
	assert.Equal(t, p.Price, expectedProductPrice)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 100)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct(productName, 0)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrPriceIsRequired)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct(productName, -1)
	assert.Nil(t, p)
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrInvalidPrice)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct(productName, 100)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
