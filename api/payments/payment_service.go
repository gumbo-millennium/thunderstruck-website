package payments

import (
	"context"

	"github.com/VictorAvelar/mollie-api-go/v4/mollie"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
)

type PaymentService struct {
	client *mollie.Client
}

func NewPaymentService(client *mollie.Client) PaymentService {
	return PaymentService{
		client: client,
	}
}

func (s PaymentService) NewPayment() (string, string, error) {
	payment := mollie.CreatePayment{
		Amount: &mollie.Amount{
			Currency: "EUR",
			Value:    "5.00",
		},
		Description: "Thunderstruck Festival 2025 ticket",
		RedirectURL: "https://thunderstruckfestival.nl/",
		WebhookURL:  "https://6f872a82767a.ngrok-free.app/orders/confirm",
	}

	_, createdPayment, err := s.client.Payments.Create(context.Background(), payment, nil)
	if err != nil {
		return "", "", err
	}

	return createdPayment.ID, createdPayment.Links.Checkout.Href, nil
}

func (s PaymentService) CheckPaymentStatus(id string) (data.OrderState, error) {
	_, payment, err := s.client.Payments.Get(context.Background(), id, nil)
	if err != nil {
		return data.OrderStatePending, err
	}

	switch payment.Status {
	case "paid":
		return data.OrderStatePaid, nil
	case "canceled":
		return data.OrderStateCancelled, nil
	default:
		return data.OrderStatePending, nil
	}
}
