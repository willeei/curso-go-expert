package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	// Arange
	amount := 500.0
	expected := 5.0

	// Act
	result := CalculateTax(amount)

	// Assert
	if result != expected {
		t.Errorf("CalculateTax(%f): expected %f, got %f", amount, expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	// Arange
	type calcTax struct {
		amount, expected float64
	}

	// Act
	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	// Assert
	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("CalculateTax(%f): expected %f, got %f", item.amount, item.expected, result)
		}
	}
}
