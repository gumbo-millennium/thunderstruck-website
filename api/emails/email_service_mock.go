package emails

import (
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/stretchr/testify/mock"
	"gopkg.in/gomail.v2"
)

type EmailServiceMock struct {
	mock.Mock
}

func (m *EmailServiceMock) Send(options EmailOptions) error {
	args := m.Called(options)

	return args.Error(0)
}

func (m *EmailServiceMock) SendTicketConfirmationEmail(ticket data.Ticket) error {
	args := m.Called(ticket)

	return args.Error(0)
}

type EmailServiceDialerMock struct {
	mock.Mock
}

func (m *EmailServiceDialerMock) DialAndSend(msg ...*gomail.Message) error {
	args := m.Called(msg)

	return args.Error(0)
}
