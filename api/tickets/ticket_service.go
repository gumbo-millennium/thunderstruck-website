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
	Create(ctx context.Context, arg data.CreateParams) (data.Ticket, error)
	Delete(ctx context.Context, id uuid.UUID) (data.Ticket, error)
	GetAll(ctx context.Context) ([]data.Ticket, error)
	GetOne(ctx context.Context, id uuid.UUID) (data.Ticket, error)
	Update(ctx context.Context, arg data.UpdateParams) (data.Ticket, error)
}

type EmailService interface {
	Send(options emails.EmailOptions) error
}

type PaymentService interface {
	Process() error
	CheckStatus(id string) (bool, error)
}

type TicketService struct {
	Repository     TicketRepository
	EmailService   EmailService
	PaymentService PaymentService
}

func NewTicketService(repository TicketRepository, emailService EmailService, paymentService PaymentService) TicketService {
	return TicketService{
		Repository:     repository,
		EmailService:   emailService,
		PaymentService: paymentService,
	}
}

func (s TicketService) GetAll() ([]data.Ticket, error) {
	return s.Repository.GetAll(context.Background())
}

func (s TicketService) NewTicket(mailAddress string) (data.Ticket, error) {
	if _, err := mail.ParseAddress(mailAddress); err != nil {
		return data.Ticket{}, ErrInvalidEmail
	}

	ticket, err := s.Repository.Create(context.Background(), data.CreateParams{
		Email: mailAddress,
		Type:  data.TicketTypeEntry,
		State: data.TicketStatePending,
		Value: s.NewTicketValue(),
	})
	if err != nil {
		return data.Ticket{}, err
	}

	if err := s.PaymentService.Process(); err != nil {
		return data.Ticket{}, err
	}

	err = s.EmailService.Send(emails.EmailOptions{
		To:      mailAddress,
		Title:   "Jouw Thunderstruck ticket staat klaar!",
		Message: "Lorem Ipsum",
	})
	if err != nil {
		return data.Ticket{}, err
	}

	return ticket, nil
}

func (s TicketService) NewTicketValue() string {
	return uuid.NewString()
}
