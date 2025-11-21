package payments

import "github.com/stretchr/testify/mock"

type PaymentServiceMock struct {
	mock.Mock
}

func (m *PaymentServiceMock) Process() error {
	args := m.Called()

	return args.Error(0)
}

func (m *PaymentServiceMock) CheckPaymentStatus(id string) (bool, error) {
	args := m.Called(id)

	return args.Get(0).(bool), args.Error(0)
}
