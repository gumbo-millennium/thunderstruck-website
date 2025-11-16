package payments

import (
	"context"
	"fmt"

	"github.com/VictorAvelar/mollie-api-go/v4/mollie"
)

type PaymentService struct {
	Client *mollie.Client
}

func NewPaymentService(client *mollie.Client) PaymentService {
	return PaymentService{
		Client: client,
	}
}

func (s PaymentService) Process() error {
	payment := mollie.CreatePayment{
		Amount: &mollie.Amount{
			Currency: "EUR",
			Value:    "5.00",
		},
		Description: "Thunderstruck Festival 2025 ticket",
		RedirectURL: "https://thunderstruckfestival.nl/redirect?id=6ee3342f-1108-484a-8df0-910f003ec185",
		WebhookURL:  "https://6f872a82767a.ngrok-free.app/tickets/webhook",
	}

	_, createdPayment, err := s.Client.Payments.Create(context.Background(), payment, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Payment created with ID: %s\n", createdPayment.ID)
	fmt.Printf("Checkout URL: %s\n", createdPayment.Links.Checkout.Href)

	return nil
}

func (s PaymentService) CheckStatus(id string) (bool, error) {
	_, payment, err := s.Client.Payments.Get(context.Background(), id, nil)
	if err != nil {
		return false, err
	}

	fmt.Println(payment.Status)

	return false, nil
}
