package orders

import (
	"context"

	"github.com/google/uuid"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/gumbo-millennium/thunderstruck-website/payments"
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
}

func NewOrderService(repository OrderRepository, paymentService payments.PaymentService) OrderService {
	return OrderService{
		Repository:     repository,
		PaymentService: paymentService,
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

	id, checkout, err := s.PaymentService.NewPayment()
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
		return order, nil
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
