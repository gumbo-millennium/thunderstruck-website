package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v4/mollie"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/gumbo-millennium/thunderstruck-website/emails"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/gumbo-millennium/thunderstruck-website/migrations"
	"github.com/gumbo-millennium/thunderstruck-website/orders"
	"github.com/gumbo-millennium/thunderstruck-website/payments"
	"github.com/gumbo-millennium/thunderstruck-website/scanner"
	"github.com/gumbo-millennium/thunderstruck-website/tickets"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"gopkg.in/gomail.v2"
)

func main() {
	// Execute migrations on program boot
	err := migrations.Execute()
	if err != nil {
		panic(err)
	}

	conn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	// Grab SQLC queries
	ctx := context.Background()
	db, err := pgx.Connect(ctx, conn)
	if err != nil {
		panic(err)
	}
	defer db.Close(ctx)
	queries := data.New(db)

	// Setup mollie integration
	environment := os.Getenv("API_ENVIRONMENT")
	config := &mollie.Config{}
	if environment == "release" {
		config = mollie.NewAPIConfig(false)
	} else {
		config = mollie.NewAPITestingConfig(false)
	}
	client, err := mollie.NewClient(nil, config)

	// Setup smtp integration
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		panic(err)
	}
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)
	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// Instantiate services
	paymentService := payments.NewPaymentService(client)
	emailService := emails.NewEmailService("bandithardcore@gmail.com", dialer)
	ticketService := tickets.NewTicketService(queries, emailService)
	ticketController := tickets.NewTicketController(ticketService)
	orderService := orders.NewOrderService(queries, paymentService, ticketService)
	orderController := orders.NewOrderController(orderService)
	scannerController := scanner.NewScannerController(os.Getenv("SCANNER_TOKEN"))

	// Define global router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}))
	r.Use(middleware.Timeout(time.Second * 60))

	// Add routes to router
	r.Route("/orders", func(r chi.Router) {
		r.Post("/", orderController.NewOrder)
		r.Post("/confirm", orderController.ConfirmOrder)
		r.Get("/{id}", orderController.GetOrder)
	})

	r.Route("/tickets", func(r chi.Router) {
		r.Get("/{id}", ticketController.GetTicket)
	})

	r.Post("/scanner/token", scannerController.ValidateToken)

	// Print all defined routes
	docgen.PrintRoutes(r)

	// Boot up router
	http.ListenAndServe(":81", r)
}
