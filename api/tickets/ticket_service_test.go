package tickets

import (
	"context"
	"testing"

	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/gumbo-millennium/thunderstruck-website/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var ticketService TicketService

func TestMain(m *testing.M) {
	ticketRepository := new(mocks.TicketRepositoryMock)
	ticketRepository.
		On("Create", context.Background(), mock.AnythingOfType("data.CreateParams")).
		Return(data.Ticket{}, nil)

	emailService := new(mocks.EmailServiceMock)
	emailService.
		On("Send", mock.AnythingOfType("emails.EmailOptions")).
		Return(nil)

	paymentService := new(mocks.PaymentServiceMock)
	paymentService.
		On("Process").
		Return(nil)
	paymentService.
		On("CheckStatus", mock.AnythingOfType("string")).
		Return(false, nil)

	ticketService = NewTicketService(ticketRepository, emailService, paymentService)
	m.Run()
}

func TestNewTicketWithValidEmailShouldNotThrowError(t *testing.T) {
	// Arrange
	mail := "test@mail.com"

	// Act
	_, err := ticketService.NewTicket(mail)

	// Assert
	assert.NoError(t, err)
}

func TestNewTicketWithInvalidEmailShouldThrowError(t *testing.T) {
	// Arrange
	mail := "invalid"

	// Act
	_, err := ticketService.NewTicket(mail)

	// Assert
	assert.ErrorIs(t, err, ErrInvalidEmail)
}
