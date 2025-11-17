package emails

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"net/mail"
	"text/template"

	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"gopkg.in/gomail.v2"
)

//go:embed email_ticket_template.html
var templates embed.FS

var (
	ErrToAddressInvalid error = errors.New("given to address is not a valid mail address")
	ErrToAddressEmpty   error = errors.New("given to address is empty")
	ErrTitleEmpty       error = errors.New("given title is empty")
	ErrMessageEmpty     error = errors.New("given message is empty")
)

type Dialer interface {
	DialAndSend(msg ...*gomail.Message) error
}

type EmailService struct {
	FromAddress string
	Dialer      Dialer
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

func NewEmailService(from string, dialer Dialer) EmailService {
	return EmailService{
		FromAddress: from,
		Dialer:      dialer,
	}
}

func (s EmailService) SendTicketConfirmationEmail(ticket data.Ticket) error {
	tmpl, err := template.ParseFS(templates, "email_ticket_template.html")
	if err != nil {
		return err
	}

	type Content struct {
		OrderURL string
	}

	content := Content{
		OrderURL: fmt.Sprintf("https://thunderstruckfestival.nl/tickets/%s", ticket.ID),
	}

	buf := bytes.Buffer{}
	err = tmpl.Execute(&buf, content)
	if err != nil {
		return err
	}

	return s.Send(EmailOptions{
		To:      ticket.Email,
		Title:   "Je Thunderstruck Festival ticket staat klaar!",
		Message: buf.String(),
	})
}

func (s EmailService) Send(options EmailOptions) error {
	if err := options.Validate(); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.FromAddress)
	msg.SetHeader("To", options.To)
	msg.SetHeader("Subject", options.Title)
	msg.SetBody("text/html", options.Message)

	return s.Dialer.DialAndSend(msg)
}
