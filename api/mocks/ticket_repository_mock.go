package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/stretchr/testify/mock"
)

type TicketRepositoryMock struct {
	mock.Mock
}

func (m *TicketRepositoryMock) Create(ctx context.Context, arg data.CreateParams) (data.Ticket, error) {
	args := m.Called(ctx, arg)

	return args.Get(0).(data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) Delete(ctx context.Context, id uuid.UUID) (data.Ticket, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) GetAll(ctx context.Context) ([]data.Ticket, error) {
	args := m.Called(ctx)

	return args.Get(0).([]data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) GetOne(ctx context.Context, id uuid.UUID) (data.Ticket, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(data.Ticket), args.Error(1)
}

func (m *TicketRepositoryMock) Update(ctx context.Context, arg data.UpdateParams) (data.Ticket, error) {
	args := m.Called(ctx, arg)

	return args.Get(0).(data.Ticket), args.Error(1)
}
