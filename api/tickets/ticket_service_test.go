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
		On("CreateTicket", context.Background(), mock.AnythingOfType("data.CreateTicketParams")).
		Return(data.Ticket{}, nil)

	emailService := new(mocks.EmailServiceMock)
	emailService.
		On("Send", mock.AnythingOfType("emails.EmailOptions")).
		Return(nil)

	ticketService = NewTicketService(ticketRepository, emailService)
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
