package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/osag1e/go-and-mongo/db"
	"github.com/osag1e/go-and-mongo/model"
)

type MovieTicketHandler struct {
	movieDataStore db.MovieDataStore
}

func NewMovieTicketHandler(movieDataStore db.MovieDataStore) *MovieTicketHandler {
	return &MovieTicketHandler{
		movieDataStore: movieDataStore,
	}
}

func (h *MovieTicketHandler) HandlePostMovie(w http.ResponseWriter, r *http.Request) {
	var params model.MovieTicket
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		render.JSON(w, r, map[string]string{"error": "Bad Request"})
		return
	}

	insertedMovie, err := h.movieDataStore.InsertMovie(r.Context(), &params)
	if err != nil {
		render.JSON(w, r, err)
		return
	}

	render.JSON(w, r, insertedMovie)
}

func (h *MovieTicketHandler) HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.movieDataStore.GetMovies(r.Context())
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Resource not found"})
		return
	}
	render.JSON(w, r, movies)
}

func (h *MovieTicketHandler) HandleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")
	if err := h.movieDataStore.DeleteMovie(r.Context(), movieID); err != nil {
		render.JSON(w, r, err)
		return
	}
	render.JSON(w, r, map[string]string{"deleted": movieID})
}
