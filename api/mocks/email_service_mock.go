package mocks

import (
	"github.com/gumbo-millennium/thunderstruck-website/emails"
	"github.com/stretchr/testify/mock"
)

type EmailServiceMock struct {
	mock.Mock
}

func (m *EmailServiceMock) Send(options emails.EmailOptions) error {
	args := m.Called(options)

	return args.Error(0)
}
