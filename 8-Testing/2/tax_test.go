package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}
