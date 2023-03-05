package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func(app *Config) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https//*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Accept","Authorization","X-CSRF-Token"},
		AllowCredentials: true,
		ExposedHeaders: []string{"Link"},
		MaxAge: 300,
	}))
	mux.Use(middleware.Heartbeat("/ping"))

	
	return mux
}