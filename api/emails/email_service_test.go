package emails

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var emailService EmailService

func TestMain(m *testing.M) {
	emailService = NewEmailService("test@mail.com")
	m.Run()
}

func TestMailShouldSendWhenAllOptionsValid(t *testing.T) {
	// Arrange
	options := EmailOptions{
		To:      "test@mail.com",
		Title:   "Hello, World!",
		Message: "Hello, World!",
	}

	// Act
	err := emailService.Send(options)

	// Assert
	assert.NoError(t, err)
}

func TestMailShouldNotSendWhenToAddressIsEmpty(t *testing.T) {
	// Arrange
	options := EmailOptions{
		Title:   "Hello, World!",
		Message: "Hello, World!",
	}

	// Act
	err := emailService.Send(options)

	// Assert
	assert.ErrorIs(t, err, ErrToAddressEmpty)
}

func TestMailShouldNotSendWhenToAddressIsInvalid(t *testing.T) {
	// Arrange
	options := EmailOptions{
		To:      "invalid",
		Title:   "Hello, World!",
		Message: "Hello, World!",
	}

	// Act
	err := emailService.Send(options)

	// Assert
	assert.ErrorIs(t, err, ErrToAddressInvalid)
}

func TestMailShouldNotSendWhenTitleIsEmpty(t *testing.T) {
	// Arrange
	options := EmailOptions{
		To:      "test@mail.com",
		Message: "Hello, World!",
	}

	// Act
	err := emailService.Send(options)

	// Assert
	assert.ErrorIs(t, err, ErrTitleEmpty)
}

func TestMailShouldNotSendWhenMessageIsEmpty(t *testing.T) {
	// Arrange
	options := EmailOptions{
		To:    "test@mail.com",
		Title: "Hello, World!",
	}

	// Act
	err := emailService.Send(options)

	// Assert
	assert.ErrorIs(t, err, ErrMessageEmpty)
}
