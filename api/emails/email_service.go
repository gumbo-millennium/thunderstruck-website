package emails

import (
	"errors"
	"net/mail"
)

var (
	ErrToAddressInvalid error = errors.New("given to address is not a valid mail address")
	ErrToAddressEmpty   error = errors.New("given to address is empty")
	ErrTitleEmpty       error = errors.New("given title is empty")
	ErrMessageEmpty     error = errors.New("given message is empty")
)

type EmailService struct {
	FromAddress string
}

type EmailOptions struct {
	To      string
	Title   string
	Message string
}

func (o EmailOptions) Validate() error {
	if o.To == "" {
		return ErrToAddressEmpty
	}

	if _, err := mail.ParseAddress(o.To); err != nil {
		return ErrToAddressInvalid
	}

	if o.Title == "" {
		return ErrTitleEmpty
	}

	if o.Message == "" {
		return ErrMessageEmpty
	}

	return nil
}

func NewEmailService(from string) EmailService {
	return EmailService{
		FromAddress: from,
	}
}

func (s EmailService) Send(options EmailOptions) error {
	if err := options.Validate(); err != nil {
		return err
	}

	return nil
}
