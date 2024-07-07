package tax

import "github.com/stretchr/testify/mock"

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) Save(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}
