package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v4/mollie"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/gumbo-millennium/thunderstruck-website/emails"
	"github.com/gumbo-millennium/thunderstruck-website/internal/data"
	"github.com/gumbo-millennium/thunderstruck-website/migrations"
	"github.com/gumbo-millennium/thunderstruck-website/payments"
	"github.com/gumbo-millennium/thunderstruck-website/tickets"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
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

	// Instantiate services
	emailService := emails.NewEmailService("noreply@thunderstruckfestival.nl")
	paymentService := payments.NewPaymentService(client)
	ticketService := tickets.NewTicketService(queries, emailService, paymentService)
	ticketController := tickets.NewTicketController(ticketService)

	// Define global router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Accept", "Content-Type"},
	}))
	r.Use(middleware.Timeout(time.Second * 60))

	// Add routes to router
	r.Post("/tickets", ticketController.Purchase)
	r.Get("/tickets", ticketController.Index)
	r.Get("/tickets/{id}", ticketController.GetById)
	r.Post("/tickets/webhook", ticketController.Webhook)

	// Print all defined routes
	docgen.PrintRoutes(r)

	// Boot up router
	http.ListenAndServe(":81", r)
}
