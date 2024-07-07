package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const errorMessage = "amount must be greater than 0"

func TestCalculateTax(t *testing.T) {
	// Arange
	amount := 500.0
	expected := 5.0

	// Act
	tax, err := CalculateTax(amount)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expected, tax)

	tax, err = CalculateTax(0.0)
	assert.Error(t, err, errorMessage)
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestCalculateAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("Save", 10.0).Return(nil)
	repository.On("Save", 0.0).Return(errors.New(errorMessage))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.NotNil(t, err)
	assert.Error(t, err, errorMessage)

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "Save", 2)
}
