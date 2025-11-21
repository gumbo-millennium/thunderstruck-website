package payments

import (
	"context"
	"fmt"
	"os"

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

func (s PaymentService) NewPayment(order data.Order) (string, string, error) {
	payment := mollie.CreatePayment{
		Amount: &mollie.Amount{
			Currency: "EUR",
			Value:    "5.40",
		},
		Description: "Thunderstruck Festival 2025 ticket",
		RedirectURL: fmt.Sprintf("%s/orders/%s", os.Getenv("MOLLIE_REDIRECT_URL"), order.ID),
		WebhookURL:  os.Getenv("MOLLIE_WEBHOOK_URL"),
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
