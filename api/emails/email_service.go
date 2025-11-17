package emails

import (
	"errors"
	"net/mail"

	"gopkg.in/gomail.v2"
)

var (
	ErrToAddressInvalid error = errors.New("given to address is not a valid mail address")
	ErrToAddressEmpty   error = errors.New("given to address is empty")
	ErrTitleEmpty       error = errors.New("given title is empty")
	ErrMessageEmpty     error = errors.New("given message is empty")
)

type EmailService struct {
	FromAddress string
	Dialer      *gomail.Dialer
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

func NewEmailService(from string, dialer *gomail.Dialer) EmailService {
	return EmailService{
		FromAddress: from,
		Dialer:      dialer,
	}
}

func (s EmailService) Send(options EmailOptions) error {
	if err := options.Validate(); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.FromAddress)
	msg.SetHeader("To", options.To)
	msg.SetHeader("Subject", options.Title)
	msg.SetBody("text/plain", options.Message)

	return s.Dialer.DialAndSend(msg)
}
