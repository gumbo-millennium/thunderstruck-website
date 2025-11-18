package orders

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/gumbo-millennium/thunderstruck-website/payments"
	"github.com/gumbo-millennium/thunderstruck-website/tickets"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrOrderAlreadyPaid error = errors.New("given order has already been paid in full")
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, arg data.CreateOrderParams) (data.Order, error)
	DeleteOrder(ctx context.Context, id uuid.UUID) (data.Order, error)
	GetAllOrders(ctx context.Context) ([]data.Order, error)
	GetOneOrder(ctx context.Context, id uuid.UUID) (data.Order, error)
	GetOneOrderByReference(ctx context.Context, reference string) (data.Order, error)
	UpdateOrder(ctx context.Context, arg data.UpdateOrderParams) (data.Order, error)
}

type OrderService struct {
	Repository     OrderRepository
	PaymentService payments.PaymentService
	TicketService  tickets.TicketService
}

func NewOrderService(repository OrderRepository, paymentService payments.PaymentService, ticketService tickets.TicketService) OrderService {
	return OrderService{
		Repository:     repository,
		PaymentService: paymentService,
		TicketService:  ticketService,
	}
}

func (s OrderService) NewOrder(mailAddress string) (data.Order, error) {
	order, err := s.Repository.CreateOrder(context.Background(), data.CreateOrderParams{
		State: data.OrderStatePending,
		Email: mailAddress,
	})
	if err != nil {
		return data.Order{}, err
	}

	id, checkout, err := s.PaymentService.NewPayment(order)
	if err != nil {
		return data.Order{}, err
	}

	order, err = s.Repository.UpdateOrder(context.Background(), data.UpdateOrderParams{
		ID:        order.ID,
		TicketID:  order.TicketID,
		Reference: id,
		Checkout:  checkout,
		State:     data.OrderStatePending,
		Email:     mailAddress,
	})
	if err != nil {
		return data.Order{}, err
	}

	return order, nil
}

func (s OrderService) ValidateOrder(reference string) (data.Order, error) {
	order, err := s.Repository.GetOneOrderByReference(context.Background(), reference)
	if err != nil {
		return data.Order{}, err
	}

	if order.State == data.OrderStatePaid {
		return order, ErrOrderAlreadyPaid
	}

	state, err := s.PaymentService.CheckPaymentStatus(reference)
	if err != nil {
		return data.Order{}, err
	}

	if state == order.State {
		return order, nil
	}

	order, err = s.Repository.UpdateOrder(context.Background(), data.UpdateOrderParams{
		ID:        order.ID,
		TicketID:  order.TicketID,
		Reference: order.Reference,
		Checkout:  order.Checkout,
		State:     state,
		Email:     order.Email,
	})
	if err != nil {
		return data.Order{}, err
	}

	return order, nil
}

func (s OrderService) ConfirmOrderByReference(reference string) (data.Order, error) {
	order, err := s.ValidateOrder(reference)
	if err != nil {
		return data.Order{}, err
	}

	if order.State != data.OrderStatePaid {
		return data.Order{}, err
	}

	ticket, err := s.TicketService.NewTicket(order.Email)
	if err != nil {
		return data.Order{}, err
	}

	order, err = s.AddTicketToOrder(ticket, order)
	if err != nil {
		return data.Order{}, err
	}

	return order, nil
}

func (s OrderService) AddTicketToOrder(ticket data.Ticket, order data.Order) (data.Order, error) {
	ticketId := pgtype.UUID{}
	err := ticketId.Scan(ticket.ID.String())
	if err != nil {
		return data.Order{}, err
	}

	order, err = s.Repository.UpdateOrder(context.Background(), data.UpdateOrderParams{
		ID:        order.ID,
		TicketID:  ticketId,
		Reference: order.Reference,
		Checkout:  order.Checkout,
		State:     order.State,
		Email:     order.Email,
	})
	if err != nil {
		return data.Order{}, err
	}

	return order, nil
}

func (s OrderService) GetOrderByID(id uuid.UUID) (data.Order, error) {
	return s.Repository.GetOneOrder(context.Background(), id)
}
