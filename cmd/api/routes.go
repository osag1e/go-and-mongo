package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/osag1e/go-and-mongo/db"
	"github.com/osag1e/go-and-mongo/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

func initializeRouter(client *mongo.Client) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	movieDataStore := db.NewMongoMovieDataStore(client)
	movieTicketHandler := handlers.NewMovieTicketHandler(movieDataStore)

	router.Post("/movie/ticket", movieTicketHandler.HandlePostMovie)
	router.Get("/movie/tickets", movieTicketHandler.HandleGetMovies)
	router.Delete("/movie/ticket/{id}", movieTicketHandler.HandleDeleteMovie)

	return router
}
