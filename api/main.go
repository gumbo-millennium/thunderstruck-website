package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/gumbo-millennium/thunderstruck-website/migrations"
	_ "github.com/lib/pq"
)

func main() {
	// Execute migrations on program boot
	err := migrations.Execute()
	if err != nil {
		panic(err)
	}

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
	r.Post("/tickets", func(w http.ResponseWriter, r *http.Request) {})

	// Print all defined routes
	docgen.PrintRoutes(r)

	// Boot up router
	http.ListenAndServe(":81", r)
}
