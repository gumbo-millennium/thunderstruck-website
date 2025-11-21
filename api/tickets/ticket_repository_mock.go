package tickets

import (
	"context"

	"github.com/google/uuid"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/stretchr/testify/mock"
)

type TicketRepositoryMock struct {
	mock.Mock
}

func (m *TicketRepositoryMock) CreateTicket(ctx context.Context, arg data.CreateTicketParams) (data.Ticket, error) {
	args := m.Called(ctx, arg)

	return args.Get(0).(data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) DeleteTicket(ctx context.Context, id uuid.UUID) (data.Ticket, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) GetAllTickets(ctx context.Context) ([]data.Ticket, error) {
	args := m.Called(ctx)

	return args.Get(0).([]data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) GetOneTicket(ctx context.Context, id uuid.UUID) (data.Ticket, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) GetOneTicketByValue(ctx context.Context, value string) (data.Ticket, error) {
	args := m.Called(ctx, value)

	return args.Get(0).(data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) UpdateTicket(ctx context.Context, arg data.UpdateTicketParams) (data.Ticket, error) {
	args := m.Called(ctx, arg)

	return args.Get(0).(data.Ticket), args.Error(1)
}
