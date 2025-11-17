package tickets

import (
	"context"
	"errors"
	"net/mail"

	"github.com/google/uuid"
	"github.com/gumbo-millennium/thunderstruck-website/emails"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
)

var (
	ErrInvalidEmail error = errors.New("invalid email given")
)

type TicketRepository interface {
	CreateTicket(ctx context.Context, arg data.CreateTicketParams) (data.Ticket, error)
	DeleteTicket(ctx context.Context, id uuid.UUID) (data.Ticket, error)
	GetAllTickets(ctx context.Context) ([]data.Ticket, error)
	GetOneTicket(ctx context.Context, id uuid.UUID) (data.Ticket, error)
	UpdateTicket(ctx context.Context, arg data.UpdateTicketParams) (data.Ticket, error)
}

type EmailService interface {
	Send(options emails.EmailOptions) error
	SendTicketConfirmationEmail(ticket data.Ticket) error
}

type PaymentService interface {
	Process() error
	CheckStatus(id string) (bool, error)
}

type TicketService struct {
	Repository   TicketRepository
	EmailService EmailService
}

func NewTicketService(repository TicketRepository, emailService EmailService) TicketService {
	return TicketService{
		Repository:   repository,
		EmailService: emailService,
	}
}

func (s TicketService) GetAll() ([]data.Ticket, error) {
	return s.Repository.GetAllTickets(context.Background())
}

func (s TicketService) NewTicket(mailAddress string) (data.Ticket, error) {
	if _, err := mail.ParseAddress(mailAddress); err != nil {
		return data.Ticket{}, ErrInvalidEmail
	}

	ticket, err := s.Repository.CreateTicket(context.Background(), data.CreateTicketParams{
		Email: mailAddress,
		Type:  data.TicketTypeEntry,
		State: data.TicketStateUnused,
		Value: s.NewTicketValue(),
	})
	if err != nil {
		return data.Ticket{}, err
	}

	err = s.EmailService.SendTicketConfirmationEmail(ticket)
	if err != nil {
		return data.Ticket{}, err
	}

	return ticket, nil
}

func (s TicketService) NewTicketValue() string {
	return uuid.NewString()
}
