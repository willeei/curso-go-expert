package tax

import "errors"

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than 0")
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5.0, nil
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}
	if amount >= 1000 && amount < 20000 {
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}

type Repository interface {
	Save(tax float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.Save(tax)
}
